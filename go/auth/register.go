package auth

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func HandleRegister(c *gin.Context) {
	// Get DSN from environment variables
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found in environment variables")
	}

	// Database connection setup
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Check the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Successfully connected to MySQL database!")

	// Retrieve form data
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	currentDate := time.Now()

	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert data into the database
	query := "INSERT INTO techsurvey.users (username, password, created, email) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, name, hashedPassword, currentDate, email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to insert data into database"})
		return
	}

	// Retrieve the last inserted ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve last insert ID"})
		return
	}

	// Return success response
	c.JSON(200, gin.H{
		"message": "User registered successfully!",
		"user_id": lastInsertID,
	})
}

// HashPassword hashes a password with bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
