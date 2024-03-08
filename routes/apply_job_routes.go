package routes

import (
	controller "vagas/controllers/apply_job"

	"github.com/gin-gonic/gin"
)

func AplyJobRoutes(router *gin.Engine) {

	company := router.Group("/aply-job")
	{
		company.GET("/:id", controller.GetAplyJob)
		company.POST("/", controller.CreateAplyJob)
		company.DELETE("/:id", controller.DeleteAplyJob)
	}
}
