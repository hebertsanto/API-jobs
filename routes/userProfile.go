package routes

import (
	handlers "vagas/handlers/profile"

	"github.com/gin-gonic/gin"
)

func UserProfileRoutes(router *gin.Engine) {

	profile := router.Group("/profile")
	{
		profile.GET("/:id", handlers.GetProfile)
		profile.POST("/", handlers.CreateProfile)
	}
}
