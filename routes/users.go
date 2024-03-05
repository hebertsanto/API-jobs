package routes

import (
	handlers "vagas/handlers/users"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.POST("/:id", handlers.GetUserById)
	}
}
