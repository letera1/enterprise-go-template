package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var (
	githubOauthConfig *oauth2.Config
	googleOauthConfig *oauth2.Config
	jwtKey            []byte
)

// User struct to hold user data
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// Claims (JWT Payload)
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found (using system env vars)")
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		jwtKey = []byte("default_secret_key_change_me")
	}

	// GitHub OAuth Config
	githubOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}

	// Google OAuth Config
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func main() {
	http.HandleFunc("/auth/github/login", handleGithubLogin)
	http.HandleFunc("/auth/github/callback", handleGithubCallback)

	http.HandleFunc("/auth/google/login", handleGoogleLogin)
	http.HandleFunc("/auth/google/callback", handleGoogleCallback)

	http.HandleFunc("/auth/logout", handleLogout)
	http.HandleFunc("/api/me", authMiddleware(handleMe))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server started at http://127.0.0.1:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(http.DefaultServeMux)))
}

// Global CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from Next.js frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Redirect to GitHub Login
func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL("random_state_string", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Redirect to Google Login
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("random_state_string", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Helper: Handle callback and generate JWT
func handleCallbackCommon(w http.ResponseWriter, r *http.Request, conf *oauth2.Config, userInfoURL string, provider string) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := conf.Client(context.Background(), token)
	resp, err := client.Get(userInfoURL)
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user User
	// Parse based on provider if structure differs slightly, but basic fields usually align or we can use map[string]interface{}
	// For simplicity, we use a generic decoder and map relevant fields if needed.
	// GitHub returns: login, id, avatar_url, name, email (sometimes private)
	// Google returns: id, email, verified_email, name, given_name, family_name, picture, locale

	var rawData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawData); err != nil {
		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
		return
	}

	// Normalize User Data
	if provider == "github" {
		user.ID = fmt.Sprintf("%.0f", rawData["id"].(float64)) // JSON numbers are float64
		if name, ok := rawData["name"].(string); ok {
			user.Name = name
		} else {
			user.Name = rawData["login"].(string)
		}
		if email, ok := rawData["email"].(string); ok {
			user.Email = email
		}
		if avatar, ok := rawData["avatar_url"].(string); ok {
			user.AvatarURL = avatar
		}
	} else if provider == "google" {
		user.ID = rawData["id"].(string)
		user.Email = rawData["email"].(string)
		user.Name = rawData["name"].(string)
		user.AvatarURL = rawData["picture"].(string)
	}

	// Generate JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "go-auth-app",
		},
	}

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenJwt.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to sign token", http.StatusInternalServerError)
		return
	}

	// Set Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true, // Secure, not accessible by JS
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		// Secure:   true, // Uncomment in production with HTTPS
	})

	// Redirect to Frontend Home
	http.Redirect(w, r, "http://localhost:3000/dashboard", http.StatusSeeOther)
}

func handleGithubCallback(w http.ResponseWriter, r *http.Request) {
	handleCallbackCommon(w, r, githubOauthConfig, "https://api.github.com/user", "github")
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	handleCallbackCommon(w, r, googleOauthConfig, "https://www.googleapis.com/oauth2/v2/userinfo", "google")
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out"))
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("auth_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		tokenStr := c.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !tkn.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add context or pass user data if needed
		next(w, r)
	}
}

func handleMe(w http.ResponseWriter, r *http.Request) {
	// In a real app, fetch user from DB using claims.UserID
	// Here we just return the claims for demo
	c, _ := r.Cookie("auth_token")
	claims := &Claims{}
	jwt.ParseWithClaims(c.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome back!",
		"user_id": claims.UserID,
		"email":   claims.Email,
	})
}
