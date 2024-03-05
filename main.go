package main

import (
	"vagas/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.SetupJobRoutes(r)

	routes.SetupUsersRoutes(r)

	r.Run(":3000")
}
