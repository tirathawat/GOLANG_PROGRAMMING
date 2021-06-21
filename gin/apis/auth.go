package apis

import (
	Models "gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleRegisterRequest(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "register success"})
}

func handleLoginRequest(c *gin.Context) {
	var form Models.LoginForm
	if c.ShouldBind(&form) == nil {
		if form.Username == "admin" && form.Password == "1234" {
			c.JSON(http.StatusOK, map[string]interface{}{"status": "You are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{"status": "Unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "Bad request"})
	}
}
