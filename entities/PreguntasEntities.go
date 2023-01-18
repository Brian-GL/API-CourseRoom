package entities

import "time"

type PreguntasRespuestaActualizarEntity struct {
	IdUsuario           int    `json:"idUsuario"`
	IdPreguntaRespuesta int    `json:"idPreguntaRespuesta"`
	Pregunta            string `json:"pregunta"`
	Descripcion         string `json:"descripcion"`
	Codigo              int    `json:"codigo"`
	Mensaje             string `json:"mensaje"`
}

type PreguntasRespuestaRegistarEntity struct {
	IdUsuario   int    `json:"idUsuario"`
	Pregunta    string `json:"pregunta"`
	Descripcion string `json:"descripcion"`
	Codigo      int    `json:"codigo"`
	Mensaje     string `json:"mensaje"`
}

type PreguntasRespuestaRemoverEntity struct {
	IdUsuario           int    `json:"idUsuario"`
	IdPreguntaRespuesta int    `json:"idPreguntaRespuesta"`
	Codigo              int    `json:"codigo"`
	Mensaje             string `json:"mensaje"`
}

type PreguntasRespuestaDetalleObtenerEntity struct {
	IdPreguntaRespuesta int        `json:"idPreguntaRespuesta"`
	IdUsuario           int        `json:"idUsuario"`
	NombreUsuario       string     `json:"nombreUsuario"`
	ImagenUsuario       string     `json:"imagenUsuario"`
	Pregunta            string     `json:"pregunta"`
	Descripcion         string     `json:"descripcion"`
	FechaRegistro       *time.Time `json:"fechaRegistro"`
	FechaActualizacion  *time.Time `json:"fechaActualizacion"`
	EstatusPregunta     string     `json:"estatusPregunta"`
}

type PreguntasRespuestaEstatusActualizarEntity struct {
	IdUsuario           int    `json:"idUsuario"`
	IdPreguntaRespuesta int    `json:"idPreguntaRespuesta"`
	IdEstatusPregunta   int    `json:"idEstatusPregunta"`
	Codigo              int    `json:"codigo"`
	Mensaje             string `json:"mensaje"`
}

type PreguntasRespuestaMensajeRegistrarEntity struct {
	IdPreguntaRespuesta int    `json:"idPreguntaRespuesta"`
	IdUsuarioEmisor     int    `json:"idUsuarioEmisor"`
	Mensaje             string `json:"mensaje"`
	Archivo             string `json:"archivo"`
	Codigo              int    `json:"codigo"`
}

type PreguntasRespuestaMensajeRemoverEntity struct {
	IdPreguntaRespuesta int    `json:"idPreguntaRespuesta"`
	IdUsuarioEmisor     int    `json:"idUsuarioEmisor"`
	IdMensaje           int    `json:"idMensaje"`
	Codigo              int    `json:"codigo"`
	Mensaje             string `json:"mensaje"`
}

type PreguntasRespuestaMensajesObtenerEntity struct {
	IdPreguntaRespuesta int        `json:"idPreguntaRespuesta"`
	UltimoMensaje       bool       `json:"ultimoMensaje"`
	IdMensaje           int        `json:"idMensaje"`
	Mensaje             string     `json:"mensaje"`
	Archivo             string     `json:"archivo"`
	IdUsuarioEmisor     int        `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string     `json:"nombreUsuarioEmisor"`
	FechaRegistro       *time.Time `json:"fechaRegistro"`
}

type PreguntasRespuestasBuscarEntity struct {
	Busqueda        string     `json:"busqueda"`
	IdPregunta      int        `json:"idPregunta"`
	IdUsuario       int        `json:"idUsaurio"`
	ImagenUsuario   string     `json:"imagenUsuario"`
	NombreUsuario   string     `json:"nombreUsuario"`
	Pregunta        string     `json:"pregunta"`
	FechaRegistro   *time.Time `json:"fechaRegistro"`
	EstatusPregunta string     `json:"estatusPregunta"`
}

type PreguntasRespuestasObtenerEntity struct {
	IdUsuario       int        `json:"idUsaurio"`
	IdPregunta      int        `json:"idPregunta"`
	Pregunta        string     `json:"pregunta"`
	FechaRegistro   *time.Time `json:"fechaRegistro"`
	EstatusPregunta string     `json:"estatusPregunta"`
}
