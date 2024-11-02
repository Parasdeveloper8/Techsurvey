package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	JwtKey []byte
	store  *sessions.CookieStore
)

func init() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := []byte(os.Getenv("SECRETKEY"))
	if secretKey == nil {
		log.Fatal("Secret key not found in environment")
	}
	JwtKey = secretKey

	// Initialize session store with secret key
	store = sessions.NewCookieStore(secretKey)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false, // Use `true` in production
		SameSite: http.SameSiteLaxMode,
	}
}

func HandleHome(c *gin.Context) {

	// Retrieve the session from the context or directly from the store
	session, err := store.Get(c.Request, "login-session")
	if err != nil {
		// Handle session retrieval error if needed
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Session error"})
		return
	}

	// Get the email from session
	sessionEmail, ok := session.Values["email"].(string)
	if !ok || sessionEmail == "" {
		// If no email found, render the index page
		c.HTML(http.StatusOK, "index.html", nil)
		return
	}

	// If email exists, render the page for logged-in users
	c.HTML(http.StatusOK, "afterlog.html", nil)
}
