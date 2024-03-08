package routes

import (
	controller "vagas/controllers/aply_job"

	"github.com/gin-gonic/gin"
)

func AplyJobRoutes(router *gin.Engine) {

	company := router.Group("/aply-job")
	{
		company.GET("/")
		company.GET("/:id")
		company.POST("/", controller.CreateAplyJob)
		company.DELETE("/:id", controller.DeleteAplyJob)
	}
}
