package routes

import (
	handlers "vagas/handlers/company"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(router *gin.Engine) {

	company := router.Group("/company")
	{
		company.GET("/:id", handlers.GetCompany)
		company.POST("/", handlers.PulishCompany)
		company.PUT("/:id", handlers.UpdateCompany)
		company.DELETE("/:id", handlers.DeleteCompany)
	}
}
