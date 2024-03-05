package routes

import (
	handlers "vagas/handlers/users"

	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(router *gin.Engine) {

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", handlers.CreateUser)
	}
}
