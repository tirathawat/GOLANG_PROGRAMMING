package apis

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	authAPI := r.Group("/authen")
	{
		authAPI.POST("/login", handleLoginRequest)
		authAPI.GET("/register", handleLoginRequest)
	}
	homeAPI := r.Group("/home")
	{
		homeAPI.GET("/profile", handleProfileRequest)
		homeAPI.GET("/home", handleHomeRequest)
		homeAPI.GET("/book/:from/:to", handleBookRequest)
	}
}
