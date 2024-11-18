package forgetpassword

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var (
	jwtKey []byte
)

func init() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load secret key from environment
	secretKey := os.Getenv("SECRETKEY")
	if secretKey == "" {
		log.Fatal("SECRETKEY not found in environment variables")
	}
	jwtKey = []byte(secretKey)
}

func GenToken(c *gin.Context) {
	var passData struct {
		Email string `json:"email" binding:"required,email"`
	}

	// Bind JSON data from the request body
	if err := c.ShouldBindJSON(&passData); err != nil {
		log.Printf("Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Please provide a valid email."})
		return
	}

	// Database connection setup
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found in environment variables")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	defer db.Close()

	// Check if the email exists
	var emailExists string
	query := "SELECT email FROM techsurvey.users WHERE email = ?"
	err = db.QueryRow(query, passData.Email).Scan(&emailExists)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
		return
	} else if err != nil {
		log.Printf("Database query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		return
	}

	// Create token
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		Issuer:    "techsurvey",
		Subject:   passData.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Insert the token into the database
	insertQuery := "INSERT INTO techsurvey.reset_tokens (email, token, expires_at) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, passData.Email, tokenString, time.Unix(claims.ExpiresAt, 0))
	if err != nil {
		log.Printf("Failed to save token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reset token"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Password reset link generated successfully", "token": tokenString})

}
