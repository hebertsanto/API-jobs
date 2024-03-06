package routes

import (
	controllers "vagas/controllers/company"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(router *gin.Engine) {

	company := router.Group("/company")
	{
		company.GET("/:id", controllers.GetCompany)
		company.POST("/", controllers.PulishCompany)
		company.PUT("/:id", controllers.UpdateCompany)
		company.DELETE("/:id", controllers.DeleteCompany)
	}
}
