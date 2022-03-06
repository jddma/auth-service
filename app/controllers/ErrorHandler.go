package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return true
	}
	return false
}
