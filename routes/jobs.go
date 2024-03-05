package routes

import (
	handlers "vagas/handlers/jobs"

	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine) {
	jobRoutes := router.Group("/jobs")
	{
		jobRoutes.POST("/", handlers.PublicJob)
		jobRoutes.GET("/:id", handlers.GetJobById)
		jobRoutes.DELETE("/:id", handlers.DeleteJob)
		jobRoutes.PUT("/:id", handlers.UpdateJobById)
	}
}
