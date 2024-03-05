package routes

import (
	handlers "vagas/handlers/users"

	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", handlers.Users)
		userRoutes.GET("/:id", handlers.ListUserById)
		userRoutes.POST("/", handlers.CreateUser)

	}
}
