package main

import (
	"os"
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
