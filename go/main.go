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

	//Middlewares start
	router.Use(auth.SessionMiddleware())
	//Middlewares end

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/favpro", auth.CheckEmail(), serveFavProForm)
	router.GET("/logpage", serveLogPage)
	router.GET("/regpage", serveRegPage)
	router.GET("/surveypage", auth.CheckEmail(), serveSurveyPageRoute)
	router.POST("/register", auth.HandleRegister)
	router.POST("/login", auth.HandleLogin)                   // Removed redundant SessionMiddleware here
	router.GET("/afterlog", auth.CheckEmail(), serveAfterLog) // Moved before Run()
	router.POST("/logout", auth.HandleLogout)
	router.Run(":4700")
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

func serveAfterLog(c *gin.Context) {
	c.HTML(http.StatusOK, "afterlog.html", nil)
}
func serveFavProForm(c *gin.Context) {
	c.HTML(http.StatusOK, "favproform.html", nil)
}
