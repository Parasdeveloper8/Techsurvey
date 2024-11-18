package forgetpassword

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Parasdeveloper8/myexpgoweb/email"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/joho/godotenv"
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
}

func ResetLink(c *gin.Context) {
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
	dsn := os.Getenv("DB_URL")
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

	// Success response
	c.JSON(http.StatusOK, gin.H{"message": "Password reset link generated successfully"})
	subject := "Password Reset Link"
	body := fmt.Sprintf("This is your password reset link: http://localhost:4700/resetpasspage")
	err = email.SendMail(passData.Email, subject, body)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}
}
