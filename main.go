package main

import (
	"github.com/gin-gonic/gin"
	"go.com/rest-api/db"
	"go.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
