package main

import (
	"log"
	"strconv"
	"vagas/config"
	"vagas/database"
	"vagas/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	database.InitDb()
	router := gin.Default()
	routes.JobRoutes(router)
	routes.UsersRoutes(router)
	routes.CompanyRoutes(router)
	routes.AplyJobRoutes(router)
	routes.UserProfileRoutes(router)

	port := config.Port

	portStr := strconv.Itoa(port)

	if err := router.Run(":" + portStr); err != nil {
		log.Fatal(err)
	}

}
