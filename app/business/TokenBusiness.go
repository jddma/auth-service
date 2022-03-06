package business

import (
	"auth-service/app/DTOs"
	"auth-service/app/util"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func GenerateTokenBusiness(ctx *gin.Context, generateTokenRequestDTO DTOs.GenerateTokenRequestDTO) {
	token := util.GenerateToken(generateTokenRequestDTO.IdUser, generateTokenRequestDTO.IdRole)
	generateTokenResponseDTO := DTOs.NewGenerateTokenResponseDTO(token)

	log.Printf("Objeto respuesta GenerateTokenBusiness -> %s", util.DtoToJson(generateTokenResponseDTO))
	ctx.JSON(http.StatusOK, generateTokenResponseDTO)
}

func ValidateTokenBusiness(ctx *gin.Context, headersDTO DTOs.HeadersWithIdentityDTO) {
	tokenArray := strings.Split(headersDTO.Token, "Bearer ")
	if len(tokenArray) != 2 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Header Authorization invalid"))
		return
	}
	if util.ValidateToken(tokenArray[1]) {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}
