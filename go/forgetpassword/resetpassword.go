package forgetpassword

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResetPassword(c *gin.Context) {
	c.HTML(http.StatusOK, "newpasspage.html", nil)
}
