package auth

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// Package-level variables
var (
	jwtKey []byte
	store  *sessions.CookieStore
	dsn    string // Define your DSN here or load it from the .env file
)

// Initialize environment variables, JWT key, and session store
func init() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRETKEY")
	if secretKey == "" {
		log.Fatal("Secret key not found in environment")
	}
	jwtKey = []byte(secretKey)

	// Initialize session store with secret key
	store = sessions.NewCookieStore(jwtKey)
	dsn = os.Getenv("DSN")
}

// Middleware to add session store to context
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "login-session")
		if err != nil {
			log.Println("Failed to get session:", err) // Log the error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
			c.Abort() // Stop processing
			return
		}
		// Store session in context
		c.Set("session", session)
		c.Next()
	}
}

// HandleLogin authenticates the user and sets session information
func HandleLogin(c *gin.Context) {
	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Get form values
	email := c.PostForm("email")
	pass := c.PostForm("password")

	// Query the user
	var (
		username string
		password string
	)
	query := "SELECT username, password FROM techsurvey.users WHERE email = ?"
	if err := db.QueryRow(query, email).Scan(&username, &password); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		} else {
			log.Printf("Database query error: %v", err) // Log the actual error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		}
		return
	}

	// Validate the password
	if !CheckPassword(password, pass) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		return
	}

	// Create JWT claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		Issuer:    "techsurvey",
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Retrieve the session and set values
	session, exists := c.Get("session")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Session not found"})
		return
	}

	sess := session.(*sessions.Session)
	sess.Values["username"] = username
	sess.Values["email"] = email
	sess.Values["token"] = tokenString

	if err := sess.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	// Redirect after successful login
	c.Redirect(http.StatusSeeOther, "/afterlog?message=User logged in successfully!")
}

// CheckPassword compares a hashed password with a plain password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
