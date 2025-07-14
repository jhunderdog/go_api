package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jhunderdog/go_api/db"
	"github.com/jhunderdog/go_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080

}

