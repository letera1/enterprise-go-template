# Go-Training - Full Stack Web Application

A modern full-stack web application built with a **Go (Golang)** backend and a **Next.js (React)** frontend. This project demonstrates authentication (JWT & OAuth), database interaction with GORM, and a responsive UI with Tailwind CSS.

## 🚀 Tech Stack

### Backend
- **Language:** Go (1.25+)
- **Framework:** [Gin](https://gin-gonic.com/) - High-performance HTTP web framework
- **ORM:** [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- **Database:** PostgreSQL
- **Authentication:** JWT (JSON Web Tokens) & OAuth2 (Google, GitHub)

### Frontend
- **Framework:** [Next.js 16+](https://nextjs.org/) (App Router)
- **Library:** React 19
- **Styling:** [Tailwind CSS v4](https://tailwindcss.com/)
- **Language:** TypeScript

---

## 🛠️ Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://go.dev/dl/) (v1.23 or newer)
- [Node.js](https://nodejs.org/) (v18 or newer)
- [PostgreSQL](https://www.postgresql.org/)

---

## ⚙️ Setup Instructions

### 1. Database Setup
Create a PostgreSQL database for the project. You can name it `gocodetest` as per the default configuration.

```sql
CREATE DATABASE gocodetest;
```

### 2. Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the `backend/` directory with the following variables:
   ```env
   PORT=9000
   
   # Database Configuration
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=gocodetest
   DB_PORT=5432 # or 3036 if using custom port

   # JWT Secret
   JWT_SECRET=your_super_secret_key

   # OAuth Configuration (Optional for local testing if not using OAuth)
   GITHUB_CLIENT_ID=your_github_client_id
   GITHUB_CLIENT_SECRET=your_github_client_secret
   GITHUB_REDIRECT_URL=http://localhost:9000/auth/github/callback

   GOOGLE_CLIENT_ID=your_google_client_id
   GOOGLE_CLIENT_SECRET=your_google_client_secret
   GOOGLE_REDIRECT_URL=http://localhost:9000/auth/google/callback
   ```

4. Run the backend server:
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:9000`.

### 3. Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install Node dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```
   The frontend will start on `http://localhost:3050`.

---

## 📂 Project Structure

```
├── backend/
│   ├── controllers/      # Route handlers (Auth, User logic)
│   ├── database/         # Database connection & init
│   ├── models/           # GORM data models (User struct)
│   ├── main.go           # Entry point & router setup
│   └── go.mod            # Go module definitions
│
├── frontend/
│   ├── app/              # Next.js App Router pages
│   │   ├── login/        # Login page
│   │   ├── signup/       # Signup page
│   │   └── dashboard/    # Protected dashboard page
│   ├── public/           # Static assets
│   └── package.json      # Frontend dependencies
└── README.md             # Project documentation
```

## ✨ Features

- **User Registration & Login**: Secure email/password authentication using bcrypt hashing.
- **Social Login**: Sign in with **GitHub** or **Google**.
- **Protected Routes**: Dashboard accessible only to authenticated users via JWT.
- **RESTful API**: specific endpoints for auth and user management.
- **Responsive Design**: Mobile-friendly UI built with Tailwind CSS.

## 📝 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/signup` | Register a new user |
| POST | `/login` | Authenticate user & get JWT |
| GET | `/auth/google` | Start Google OAuth flow |
| GET | `/auth/github` | Start GitHub OAuth flow |
| GET | `/dashboard` | Protected route example |

## 🤝 Contributing

Feel free to fork this project and submit pull requests.

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
