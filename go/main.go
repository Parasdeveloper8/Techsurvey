package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		// Pass data to the template
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": "User",
		})
	})

	router.GET("/surveypage", serveSurveyPageRoute)
	router.Run(":4700")
}
func serveSurveyPageRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "surveypages.html", nil)
}
