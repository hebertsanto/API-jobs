package routes

import "github.com/gin-gonic/gin"

func UserProfileRoutes(router *gin.Engine) {

	profile := router.Group("/profile")
	{
		profile.GET("/")
		profile.GET("/:id")
	}
}
