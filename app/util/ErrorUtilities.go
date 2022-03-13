package util

type ErrorMsg struct {
	Code        string
	Description string
}

var (
	DatabaseConnectionError = ErrorMsg{
		Code:        "EDB01",
		Description: "Error de conexión con la base de datos",
	}
	EmailDuplicateError = ErrorMsg{
		Code:        "EU01",
		Description: "Ya existe un usuario registrado con el correo ingresado",
	}
	UserDuplicateError = ErrorMsg{
		Code:        "EU02",
		Description: "Ya existe un usuario registrado para la persona y rol ingresado",
	}
	OrchestrationError = ErrorMsg{
		Code:        "EU03",
		Description: "Ocurrio un problema al orquestar los servicios. intente mas tarde",
	}
)
