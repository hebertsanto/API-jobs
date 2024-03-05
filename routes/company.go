package routes

import "github.com/gin-gonic/gin"

func CompanyRoutes(router *gin.Engine) {

	company := router.Group("/company")
	{
		company.GET("/")
		company.GET("/:id")
		company.POST("/")
		company.PUT("/:id")
		company.DELETE("/:id")
	}
}
