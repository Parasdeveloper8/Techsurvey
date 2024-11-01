package auth

import (
	"database/sql"
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

	secretKey := os.Getenv("SECRETKEY")
	if secretKey == "" {
		log.Fatal("Secret key not found in environment")
	}
	jwtKey = []byte(secretKey)

	// Initialize session store with secret key
	store = sessions.NewCookieStore(jwtKey)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1 week
		HttpOnly: true,
		Secure:   false, // Use `true` in production
		SameSite: http.SameSiteLaxMode,
	}
	dsn = os.Getenv("DSN")
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
	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Get form values
	email := c.PostForm("email")
	pass := c.PostForm("password")

	var (
		username string
		password string
	)
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

	session, exists := c.Get("session")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": session, "err": exists})
		return
	}

	sess, ok := session.(*sessions.Session)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Session type assertion failed"})
		return
	}

	sess.Values["username"] = username
	sess.Values["email"] = email
	sess.Values["token"] = tokenString

	err = sess.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/afterlog?user logged in successfully")
	//c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// CheckPassword compares a hashed password with a plain password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
