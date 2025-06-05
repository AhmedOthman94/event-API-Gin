package main

import (
	"github.com/AhmedOthman94/REST-API-Gin/db"
	"github.com/AhmedOthman94/REST-API-Gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
