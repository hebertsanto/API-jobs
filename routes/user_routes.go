package routes

import (
	controllers "vagas/controllers/users"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/all", controllers.GetUsers)
	}
}
