package routes

import (
	controllers "vagas/controllers/jobs"

	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine) {
	jobRoutes := router.Group("/jobs")
	{
		jobRoutes.POST("/", controllers.PublishJob)
		jobRoutes.GET("/:id", controllers.GetJobById)
		jobRoutes.DELETE("/:id", controllers.DeleteJob)
		jobRoutes.PUT("/:id", controllers.UpdateJobById)
	}
}
