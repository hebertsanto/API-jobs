package main

import (
	"vagas/database"
	"vagas/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDb()
	router := gin.Default()

	routes.JobRoutes(router)

	routes.UsersRoutes(router)

	routes.CompanyRoutes(router)

	routes.AplyJobRoutes(router)

	routes.UserProfileRoutes(router)

	router.Run(":3000")
}
