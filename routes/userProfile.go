package routes

import (
	controllers "vagas/controllers/profile"

	"github.com/gin-gonic/gin"
)

func UserProfileRoutes(router *gin.Engine) {

	profile := router.Group("/profile")
	{
		profile.GET("/:id", controllers.GetProfile)
		profile.POST("/", controllers.CreateProfile)
	}
}
