package business

import (
	"auth-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func validateConnectionWithDatabase(ctx *gin.Context, err error) bool {
	if err != nil {
		log.Printf("Error de conexión con base de datos")
		ctx.JSON(http.StatusInternalServerError, util.DatabaseConnectionError)
		return true
	}
	return false
}
