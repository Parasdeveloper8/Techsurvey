package forgetpassword

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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
func ResetPassword(c *gin.Context) {
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
	email := c.PostForm("email")
	newPassword := c.PostForm("new-password")

	// Hash the password
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	query := "UPDATE techsurvey.users SET password = ? WHERE email = ?"
	result, err := db.Exec(query, hashedPassword, email)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
	c.Redirect(http.StatusSeeOther, "/logpage?message=password change successfully")
}

// HashPassword hashes a password with bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
