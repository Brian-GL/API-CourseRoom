package models

type ResponseStatus int

const (
	SUCCESS = 1
	ALERT   = 2
	ERROR   = 3
)

type ResponseInfrastructure struct {
	Status ResponseStatus `json:"status"`
	Data   any            `json:"data"`
}

type BienvenidaEmail struct {
	CorreoElectronico string
	NombreCompleto    string
	Nombre            string
	Anio              int
}

type CalificacionEmail struct {
	CorreoElectronico    string
	NombreTarea          string
	FechaCalificacion    string
	CalificacionObtenida float32
	PuntualidadObtenida  float32
	Anio                 int
}

type CredencialesEmail struct {
	CorreoElectronico string
	QR_URL            string
	Anio              int
}

type EmailVerificatorAPIError struct {
	Status bool `json:"status"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type EmailVerificatorAPISuccess struct {
	Status bool   `json:"status"`
	Email  string `json:"email"`
	Domain string `json:"domain"`
}

type CourseRoomCalculator struct {
	IdUsuario  int
	IdTarea    int
	IdProfesor int
}

type EmailConfiguration struct {
	EMAIL_SERVER      string
	EMAIL_PORT        int
	EMAIL_ADDRESS     string
	EMAIL_CREDENTIALS string
}
