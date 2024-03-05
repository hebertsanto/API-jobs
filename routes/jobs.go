package routes

import (
	handlers "vagas/handlers/jobs"

	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine) {
	jobRoutes := router.Group("/jobs")
	{
		jobRoutes.GET("/", handlers.ListJob)
		jobRoutes.GET("/:id", handlers.ListJob)
		jobRoutes.POST("/", handlers.PublicJob)
		jobRoutes.PUT("/:id", handlers.ListJob)
		jobRoutes.DELETE("/:id", handlers.ListJob)
	}
}
