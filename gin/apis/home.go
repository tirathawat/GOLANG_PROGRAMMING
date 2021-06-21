package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHomeRequest(c *gin.Context) {
	username, age := c.Query("username"), c.Query("age")
	result := map[string]interface{}{"result": "ok", "username": username, "age": age}
	c.JSON(http.StatusOK, result)
}

func handleProfileRequest(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf8", []byte("PROFILE"))
}

func handleBookRequest(c *gin.Context) {
	from, to := c.Param("from"), c.Param("to")
	result := map[string]interface{}{"result": "ok", "from": from, "to": to}
	c.JSON(http.StatusOK, result)
}
