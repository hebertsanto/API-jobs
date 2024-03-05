package main

import (
	"vagas/database"
	"vagas/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDb()
	router := gin.Default()

	routes.SetupJobRoutes(router)

	routes.SetupUsersRoutes(router)

	router.Run(":3000")
}
