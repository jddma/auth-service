package config

import (
	"auth-service/app/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefinePaths(router *gin.RouterGroup) {
	router.POST("/token", controllers.GenerateTokenController)
	router.POST("/validate-token", controllers.ValidateTokenController)
}

func SetCors(router *gin.RouterGroup) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
}

func SetLogger(router *gin.RouterGroup) {
	router.Use(gin.Logger())
}
