package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var StoreofTransPoints *sessions.CookieStore

func init() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := []byte(os.Getenv("SECRETKEY"))
	if len(secretKey) == 0 {
		log.Fatal("Secret key not found in environment")
	}
	// Initialize session store with secret key
	StoreofTransPoints = sessions.NewCookieStore(secretKey)
	StoreofTransPoints.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false, // Use `true` in production
		SameSite: http.SameSiteLaxMode,
	}
}

func TransferPoints(c *gin.Context) {
	var transferData struct {
		Points int `json:"points"`
	}

	// Bind JSON data from the request body
	if err := c.ShouldBindJSON(&transferData); err != nil {
		log.Printf("Failed to bind JSON data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Retrieve the session from the context or directly from the store
	session, err := StoreofTransPoints.Get(c.Request, "login-session")
	if err != nil {
		log.Printf("Failed to get session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get session"})
		return
	}
	// Get the email from session
	sessionEmail, ok := session.Values["email"].(string)
	if !ok || sessionEmail == "" {
		log.Println("Email not found in session")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email not found in session"})
		return
	}
	// Get DSN from env file
	dsn := os.Getenv("DSN")
	if len(dsn) == 0 {
		log.Println("DSN not found in environment variables")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Configuration error"})
		return
	}
	// Database connection setup
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()
	// Check the connection
	if err := db.Ping(); err != nil {
		log.Printf("Database connection failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	fmt.Println("Successfully connected to MySQL database!")

	query := `UPDATE techsurvey.users 
	SET points = ? 
	WHERE email = ?;`
	_, err = db.Exec(query, transferData.Points, sessionEmail)
	if err != nil {
		log.Printf("Failed to update points: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Points updated successfully"})
}
