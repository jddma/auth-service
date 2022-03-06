package DTOs

type HeadersWithIdentityDTO struct {
	Token string `header:"Authorization" binding:"required"`
	Uuid  string `header:"uuid" binding:"required"`
}

type AnonymousHeadersDTO struct {
	Uuid string `header:"uuid" binding:"required"`
}
