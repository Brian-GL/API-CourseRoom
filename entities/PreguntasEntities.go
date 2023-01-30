package entities

import "time"

type PreguntaRespuestaDetalleObtenerEntity struct {
	IdUsuario          int        `json:"idUsuario"`
	NombreUsuario      string     `json:"nombreUsuario"`
	ImagenUsuario      *string    `json:"imagenUsuario"`
	Pregunta           string     `json:"pregunta"`
	Descripcion        string     `json:"descripcion"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
	EstatusPregunta    string     `json:"estatusPregunta"`
}

type PreguntaRespuestaMensajesObtenerEntity struct {
	IdMensaje           int       `json:"idMensaje"`
	Mensaje             string    `json:"mensaje"`
	Archivo             *string   `json:"archivo"`
	IdUsuarioEmisor     int       `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string    `json:"nombreUsuarioEmisor"`
	ImagenEmisor        string    `json:"imagenEmisor"`
	FechaRegistro       time.Time `json:"fechaRegistro"`
}

type PreguntasRespuestasBuscarEntity struct {
	IdPregunta      int       `json:"idPregunta"`
	IdUsuario       int       `json:"idUsuario"`
	ImagenUsuario   *string   `json:"imagenUsuario"`
	NombreUsuario   string    `json:"nombreUsuario"`
	Pregunta        string    `json:"pregunta"`
	FechaRegistro   time.Time `json:"fechaRegistro"`
	EstatusPregunta string    `json:"estatusPregunta"`
}

type PreguntasRespuestasObtenerEntity struct {
	IdPregunta      int       `json:"idPregunta"`
	Pregunta        string    `json:"pregunta"`
	FechaRegistro   time.Time `json:"fechaRegistro"`
	EstatusPregunta string    `json:"estatusPregunta"`
}
