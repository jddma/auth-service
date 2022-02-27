package controllers

import (
	"auth-service/app/DTOs"
	"auth-service/app/business"
	"auth-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GenerateTokenController(ctx *gin.Context) {
	generateTokenRequestDTO := new(DTOs.GenerateTokenRequestDTO)
	err := ctx.Bind(generateTokenRequestDTO)
	handleError(ctx, err)

	log.Printf("Objeto request GenerateTokenController -> %s", util.DtoToJson(generateTokenRequestDTO))
	business.GenerateTokenBusiness(ctx, *generateTokenRequestDTO)
}

func ValidateTokenController(ctx *gin.Context) {
	validateTokenRequestDTO := new(DTOs.TokenDTO)
	err := ctx.Bind(validateTokenRequestDTO)
	handleError(ctx, err)

	log.Printf("Objeto request ValidateTokenController -> %s", util.DtoToJson(validateTokenRequestDTO))
	business.ValidateTokenBusiness(ctx, *validateTokenRequestDTO)
}
