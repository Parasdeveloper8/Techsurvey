package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Cookiestore *sessions.CookieStore

func init() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := []byte(os.Getenv("SECRETKEY"))
	if secretKey == nil {
		log.Fatal("Secret key not found in environment")
	}
	// Initialize session store with secret key
	Cookiestore = sessions.NewCookieStore(secretKey)
	Cookiestore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false, // Use `true` in production
		SameSite: http.SameSiteLaxMode,
	}
}
func HandleFeedSubmission(c *gin.Context) {
	session, err := Cookiestore.Get(c.Request, "login-session") // Use Store here
	if err != nil {
		log.Printf("Failed to get session: %v", err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get session"})
		return
	}
	// Get the email from session
	sessionEmail, ok := session.Values["email"].(string)
	if !ok || sessionEmail == "" {
		log.Println("Email not found in session") // Log the error
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email not found in session"})
		return
	}

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
	feedback := c.PostForm("feedback")
	currentTime := time.Now()

	// Insert data into the database
	query := "INSERT INTO techsurvey.feedback(emailofuser,comment,time) VALUES (?,?,?)"
	result, err := db.Exec(query, sessionEmail, feedback, currentTime)
	if err != nil {
		log.Printf("Failed to send feedback into database: %v", err) // Log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send feedback into database"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/afterlog?message=Thanks for suggestion")
	fmt.Println(result)
}
