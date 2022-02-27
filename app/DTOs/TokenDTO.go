package DTOs

type TokenDTO struct {
	Token string `json:"token"`
}

func NewGenerateTokenResponseDTO(token string) *TokenDTO {
	return &TokenDTO{Token: token}
}
