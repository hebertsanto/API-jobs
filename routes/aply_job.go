package routes

import "github.com/gin-gonic/gin"

func AplyJobRoutes(router *gin.Engine) {

	company := router.Group("/aply-job")
	{
		company.GET("/")
		company.GET("/:id")
		company.POST("/")
		company.DELETE("/:id")
	}
}