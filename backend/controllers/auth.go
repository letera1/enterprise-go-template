package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"backend/database"
	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "", // Set in Init
	ClientSecret: "", // Set in Init
	RedirectURL:  os.Getenv("GITHUB_REDIRECT_URL"),
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
}

func InitOAuth() {
	githubOauthConfig.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	githubOauthConfig.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
}

// Signup
func Signup(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Provider: "email",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

// GitHub Login Redirect
func GitHubLogin(c *gin.Context) {
	url := githubOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GitHub Callback
func GitHubCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify code"})
		return
	}

	client := githubOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	var githubUser struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		ID        int    `json:"id"`
		AvatarURL string `json:"avatar_url"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&githubUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	// If email is missing (private), fetch it specifically
	if githubUser.Email == "" {
		// Fetch emails logic... simplified for now, assuming public email or scope works
		// In production, implement /user/emails call here if needed
	}

	var user models.User
	if err := database.DB.Where("email = ?", githubUser.Email).First(&user).Error; err != nil {
		// Create new user
		user = models.User{
			Name:      githubUser.Name,
			Email:     githubUser.Email,
			AvatarURL: githubUser.AvatarURL,
			Provider:  "github",
		}
		database.DB.Create(&user)
	} else {
		// Update existing user info in case it changed on GitHub
		user.Name = githubUser.Name
		user.AvatarURL = githubUser.AvatarURL
		database.DB.Save(&user)
	}

	jwtToken, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Redirect to frontend with token
	// Assuming frontend is at localhost:3000
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://localhost:3000/login?token=%s", jwtToken))
}

func generateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
