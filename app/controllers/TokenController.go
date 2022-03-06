package controllers

import (
	"auth-service/app/DTOs"
	"auth-service/app/business"
	"auth-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GenerateTokenController(ctx *gin.Context) {
	headersDTO, err := getAnonymousHeaders(ctx)
	if handleError(ctx, err) {
		return
	}
	generateTokenRequestDTO := new(DTOs.GenerateTokenRequestDTO)
	err = ctx.ShouldBindJSON(generateTokenRequestDTO)
	if handleError(ctx, err) {
		return
	}

	log.Printf("Objeto request GenerateTokenController body -> %s, headers -> %s",
		util.DtoToJson(generateTokenRequestDTO), util.DtoToJson(headersDTO))
	business.GenerateTokenBusiness(ctx, *generateTokenRequestDTO)
}

func ValidateTokenController(ctx *gin.Context) {
	headersDTO, err := getHeadersWithIdentityDTO(ctx)
	if handleError(ctx, err) {
		return
	}

	log.Printf("Headers request ValidateTokenController -> %s", util.DtoToJson(headersDTO))
	business.ValidateTokenBusiness(ctx, *headersDTO)
}
