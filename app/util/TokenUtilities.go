package util

import (
	"auth-service/app/model"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type authCustomClaims struct {
	IdUser   int            `json:"idUser""`
	IdRole   int            `json:"idRole"`
	Accesses []model.Access `json:"accesses"`
	jwt.StandardClaims
}

func getSecretKey() string {
	return os.Getenv("token_secret_key")
}

func GenerateToken(idUser int, idRole int, accesess []model.Access) string {
	claims := &authCustomClaims{
		idUser,
		idRole,
		accesess,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Service",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(getSecretKey()))
	return t
}

func ValidateToken(tokenStr string) bool {
	claims := new(authCustomClaims)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecretKey()), nil
	})
	if err != nil {
		log.Printf("Error al validar el Token -> %s, Descripcion -> %s", tokenStr, err.Error())
		return false
	}
	if !token.Valid {
		log.Printf("Token-> %s no valido", tokenStr)
		return false
	}
	return true
}
