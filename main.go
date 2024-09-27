package main

import (
	"77gsi_mynt.com/cv_tagging/db"
	"77gsi_mynt.com/cv_tagging/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db.InitDB()

	routes.Routes(server)

	server.Run()
}
