package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
}
