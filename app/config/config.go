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

func GetCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("uuid", "Authorization")
	config.AllowAllOrigins = true
	return cors.New(config)
}

func SetLogger(router *gin.RouterGroup) {
	router.Use(gin.Logger())
}
