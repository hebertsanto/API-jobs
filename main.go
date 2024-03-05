package main

import (
	"vagas/database"
	"vagas/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDb()

	r := gin.Default()

	routes.SetupJobRoutes(r)

	routes.SetupUsersRoutes(r)

	r.Run(":3000")
}
