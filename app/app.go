package app

import (
	"auth-service/app/config"
	"github.com/gin-gonic/gin"
	"os"
)

func Run() {
	r := gin.Default()
	r.Use(config.GetCors())

	g := r.Group("/api/v1")
	config.DefinePaths(g)
	config.SetLogger(g)

	r.Run(os.Getenv("port"))
}
