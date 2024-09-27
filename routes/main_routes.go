package routes

import (
	middleware "77gsi_mynt.com/cv_tagging/middlewares"
	"77gsi_mynt.com/cv_tagging/services"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {

	server.POST("/signup", services.Signup)
	server.POST("/login", services.Login)

	private := server.Group("/")
	private.Use(middleware.Authenticate)
	private.POST("/keyword", services.SaveKeyword)
	private.POST("/scan", services.Scan)
}
