package util

import (
	"auth-service/app/DTOs"
	"encoding/json"
)

func DtoToJson(dto DTOs.DTO) string {
	jsonString, _ := json.Marshal(dto)
	return string(jsonString)
}
