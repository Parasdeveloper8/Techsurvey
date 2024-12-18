package auth

import (
	"database/sql"
	"fmt"
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

var (
	jwtKey []byte
	store  *sessions.CookieStore
	dsn    string
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
	jwtKey = secretKey

	// Initialize session store with secret key
	store = sessions.NewCookieStore(secretKey)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false, // Use `true` in production
		SameSite: http.SameSiteLaxMode,
	}
	dsn = os.Getenv("DB_URL")
}

// Middleware to add session store to context
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "login-session")
		if err != nil {
			log.Println("Failed to get session:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
			c.Abort()
			return
		}
		// Store session in context
		c.Set("session", session)
		c.Next()
	}
}

// HandleLogin authenticates the user and sets session information
func HandleLogin(c *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	email := c.PostForm("email")
	pass := c.PostForm("password")

	var username, password string
	query := "SELECT username, password FROM techsurvey.users WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&username, &password)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error"})
		log.Printf("Database query error: %v", err)
		return
	}

	if !CheckPassword(password, pass) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		return
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		Issuer:    "techsurvey",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	session := c.MustGet("session").(*sessions.Session) // Retrieve the session from context

	session.Values["username"] = username
	session.Values["email"] = email
	session.Values["token"] = tokenString
	//fmt.Println(session.Values["email"])
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/afterlog?user logged in successfully")
	fmt.Println(session)
}

// CheckPassword compares a hashed password with a plain password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// HandleLogout handles user logout
func HandleLogout(c *gin.Context) {
	session := c.MustGet("session").(*sessions.Session) // Use the session from context

	// Clear the session data
	session.Values = make(map[interface{}]interface{}) // Reset the session data
	err := session.Save(c.Request, c.Writer)           // Save the cleared session
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out"})
		return
	}

	fmt.Println("Session destroyed")
	c.Redirect(http.StatusSeeOther, "/") // Redirect to home
}

// middleware to check if email exists in session
func CheckEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "login-session")
		sessionEmail, ok := session.Values["email"].(string)
		if !ok || sessionEmail == "" {
			c.Redirect(http.StatusFound, "/logpage") // Redirect to login page
			c.Abort()
			return
		}
		c.Next()
	}
}
