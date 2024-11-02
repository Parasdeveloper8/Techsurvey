package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
