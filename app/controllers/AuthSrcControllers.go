package controllers

import (
	"auth-service/app/DTOs"
	"github.com/gin-gonic/gin"
)

func getAnonymousHeaders(ctx *gin.Context) (*DTOs.AnonymousHeadersDTO, error) {
	headersDTO := new(DTOs.AnonymousHeadersDTO)
	err := ctx.ShouldBindHeader(&headersDTO)
	return headersDTO, err
}

func getHeadersWithIdentityDTO(ctx *gin.Context) (*DTOs.HeadersWithIdentityDTO, error) {
	headersDTO := new(DTOs.HeadersWithIdentityDTO)
	err := ctx.ShouldBindHeader(&headersDTO)
	return headersDTO, err
}
