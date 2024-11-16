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

var Store *sessions.CookieStore

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
	Store = sessions.NewCookieStore(secretKey)
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false, // Use `true` in production
		SameSite: http.SameSiteLaxMode,
	}
}
func HandleFavProSurveySubmission(c *gin.Context) {
	// Retrieve the session from the context or directly from the store
	session, err := Store.Get(c.Request, "login-session") // Use Store here
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
	number := c.PostForm("numofp")
	favlang := c.PostForm("favp")
	currentDate := time.Now()
	// Insert data into the database
	query := "INSERT INTO techsurvey.faveprogramlang (email, vote, date, number) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, sessionEmail, favlang, currentDate, number) // Include number in query
	if err != nil {
		log.Printf("Failed to insert data into database: %v", err) // Log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into database or giving survey for more than 1 time"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/afterlog?message=Thanks for taking part in survey")
	fmt.Println(result)
}
