package main

import (
	"example.com/first-app/database"
	"example.com/first-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Initdatabse()
	server := gin.Default()
	routes.Registerevent(server)
	server.Run(":8080")
}
