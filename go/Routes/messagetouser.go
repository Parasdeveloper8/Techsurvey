package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

}
func HandleUpdate(c *gin.Context) {

	// Get DSN from env file
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found in environment variables")
	}

	// Database connection setup
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Log fatal error
	}
	defer db.Close()
	// Check the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err) // Log fatal error
	}
	fmt.Println("Successfully connected to MySQL database!")

	// Get form data
	update := c.PostForm("update")
	currentTime := time.Now()

	// Insert data into the database
	query := "INSERT INTO techsurvey.updates(message,time) VALUES (?,?)"
	result, err := db.Exec(query, update, currentTime)
	if err != nil {
		log.Printf("Failed to send message into database: %v", err) // Log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send feedback into database"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin?message=update sent")
	fmt.Println(result)
}
