package DTOs

type GenerateTokenRequestDTO struct {
	IdUser int `json:"idUser" binding:"required"`
	IdRole int `json:"idRole" binding:"required"`
}
