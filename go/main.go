package main

import (
	"net/http"

	routes "github.com/Parasdeveloper8/myexpgoweb/Routes"
	"github.com/Parasdeveloper8/myexpgoweb/auth"
	"github.com/Parasdeveloper8/myexpgoweb/cors"
	"github.com/Parasdeveloper8/myexpgoweb/forgetpassword"
	"github.com/Parasdeveloper8/myexpgoweb/limiter"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	//Middlewares start
	router.Use(limiter.RateLimit())

	router.Use(auth.SessionMiddleware())
	//Middlewares end
	// Set up CORS
	cors.SetupCORS(router)

	router.GET("/", routes.HandleHome)
	router.GET("/logpage", serveLogPage)
	router.GET("/regpage", serveRegPage)
	router.GET("/admin", serveAdminPage)
	router.GET("/surveypage", auth.CheckEmail(), serveSurveyPageRoute)
	router.POST("/register", auth.HandleRegister)
	router.POST("/login", auth.HandleLogin)
	router.GET("/leaderboard", auth.CheckEmail(), serveLeaderboard)
	router.GET("/afterlog", auth.CheckEmail(), serveAfterLog)
	router.POST("/logout", auth.HandleLogout)
	router.GET("/feedback", serveGiveFeedBack)

	//points related routes start
	router.GET("/getpoints", routes.GetPoint)
	router.POST("/transferpoints", routes.TransferPoints)
	//points related routes end

	//survey submission routes start
	router.POST("/submitfav", routes.HandleFavProSurveySubmission)
	router.POST("/submitfavframe", routes.HandleFavFrameSurveySubmission)
	router.POST("/submitfavdev", routes.HandleFavDevSurveySubmission)
	//survey submission routes end

	router.POST("/subfeed", auth.CheckEmail(), routes.HandleFeedSubmission)
	router.GET("/points", auth.CheckEmail(), servePoints)
	router.GET("/admincomment", serveAdminComment)
	router.GET("/adminsurvey", serveAdminSurvey)
	router.GET("/updatesforuser", auth.CheckEmail(), serveCommunity)
	router.GET("/adminmessage", serveAdminMessage)
	router.POST("/messagetouser", routes.HandleUpdate)

	//Survey Pages routes start
	router.GET("/favpro", auth.CheckEmail(), serveFavProForm)
	router.GET("/favframe", auth.CheckEmail(), serveFavFrame)
	router.GET("/favdev", auth.CheckEmail(), serveFavDev)
	//Survey Pages routes end

	//Forget password related routes start
	router.POST("/gentoken", forgetpassword.GenToken)
	//Forget password related routes end

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
func serveAdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}
func serveGiveFeedBack(c *gin.Context) {
	c.HTML(http.StatusOK, "givefeedback.html", nil)
}
func serveLeaderboard(c *gin.Context) {
	c.HTML(http.StatusOK, "leaderboard.html", nil)
}
func servePoints(c *gin.Context) {
	c.HTML(http.StatusOK, "point.html", nil)
}
func serveAdminComment(c *gin.Context) {
	c.HTML(http.StatusOK, "admincomment.html", nil)
}
func serveAdminSurvey(c *gin.Context) {
	c.HTML(http.StatusOK, "adminsurvey.html", nil)
}
func serveCommunity(c *gin.Context) {
	c.HTML(http.StatusOK, "community.html", nil)
}
func serveAdminMessage(c *gin.Context) {
	c.HTML(http.StatusOK, "adminmessage.html", nil)
}

func serveFavDev(c *gin.Context) {
	c.HTML(http.StatusOK, "favdev.html", nil)
}

func serveFavFrame(c *gin.Context) {
	c.HTML(http.StatusOK, "favframe.html", nil)
}
