package business

import (
	"auth-service/app/DTOs"
	"auth-service/app/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GenerateTokenBusiness(ctx *gin.Context, generateTokenRequestDTO DTOs.GenerateTokenRequestDTO) {
	token := util.GenerateToken(generateTokenRequestDTO.IdUser, generateTokenRequestDTO.IdRole)
	generateTokenResponseDTO := DTOs.NewGenerateTokenResponseDTO(token)

	log.Printf("Objeto respuesta GenerateTokenBusiness -> %s", util.DtoToJson(generateTokenResponseDTO))
	ctx.JSON(http.StatusOK, generateTokenResponseDTO)
}

func ValidateTokenBusiness(ctx *gin.Context, validateTokenRequestDTO DTOs.TokenDTO) {
	if util.ValidateToken(validateTokenRequestDTO.Token) {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}
