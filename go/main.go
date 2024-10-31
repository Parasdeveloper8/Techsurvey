package main

import (
	"net/http"

	"github.com/Parasdeveloper8/myexpgoweb/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/afterlog", serveAfterLog)

	router.GET("/logpage", serveLogPage)

	router.GET("/regpage", serveRegPage)

	router.GET("/surveypage", serveSurveyPageRoute)

	router.POST("/register", auth.HandleRegister)

	router.POST("/login", auth.HandleLogin)

	router.Run(":4700")
}
func serveAfterLog(c *gin.Context) {
	c.HTML(http.StatusOK, "afterlog.html", nil)
}
func serveLogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "loginpage.html", nil)
}
func serveRegPage(c *gin.Context) {
	c.HTML(http.StatusOK, "registerpage.html", nil)
}
func serveSurveyPageRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "surveypages.html", nil)
}
