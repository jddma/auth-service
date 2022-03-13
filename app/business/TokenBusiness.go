package business

import (
	"auth-service/app/DTOs"
	"auth-service/app/repository"
	"auth-service/app/util"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func GenerateTokenBusiness(ctx *gin.Context, generateTokenRequestDTO DTOs.GenerateTokenRequestDTO) {
	accesses, err := repository.FindAccessRoleByIdRole(generateTokenRequestDTO.IdRole)
	if validateConnectionWithDatabase(ctx, err) {
		return
	}
	log.Printf("Accesos concdidos para id role # %d -> %s", generateTokenRequestDTO.IdRole,
		util.DtoToJson(accesses))

	token := util.GenerateToken(generateTokenRequestDTO.IdUser, generateTokenRequestDTO.IdRole, accesses)
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
