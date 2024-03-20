package main

import (
	"github.com/gin-gonic/gin"
	"example.com/REST-API/routes"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
