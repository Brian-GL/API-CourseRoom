package models

type PreguntasRespuestaActualizarInputModel struct {
	IdUsuario           *int    `json:"idUsuario" validate:"required"`
	IdPreguntaRespuesta *int    `json:"idPreguntaRespuesta" validate:"required"`
	Pregunta            *string `json:"pregunta" validate:"required"`
	Descripcion         *string `json:"descripcion" validate:"required"`
}

type PreguntasRespuestaRegistrarInputModel struct {
	IdUsuario   *int    `json:"idUsuario" validate:"required"`
	Pregunta    *string `json:"pregunta" validate:"required"`
	Descripcion *string `json:"descripcion" validate:"required"`
}

type PreguntasRespuestaRemoverInputModel struct {
	IdUsuario           *int `json:"idUsuario" validate:"required"`
	IdPreguntaRespuesta *int `json:"idPreguntaRespuesta" validate:"required"`
}

type PreguntasRespuestaDetalleObtenerInputModel struct {
	IdPreguntaRespuesta *int `json:"idPreguntaRespuesta" validate:"required"`
}

type PreguntasRespuestaEstatusActualizarInputModel struct {
	IdUsuario           *int `json:"idUsuario" validate:"required"`
	IdPreguntaRespuesta *int `json:"idPreguntaRespuesta" validate:"required"`
	IdEstatusPregunta   *int `json:"idEstatusPregunta" validate:"required"`
}

type PreguntasRespuestaMensajeRegistrarInputModel struct {
	IdPreguntaRespuesta *int    `json:"idPreguntaRespuesta" validate:"required"`
	IdUsuarioEmisor     *int    `json:"idUsuarioEmisor" validate:"required"`
	Mensaje             *string `json:"mensaje" validate:"required"`
	Archivo             *string `json:"archivo"`
}

type PreguntasRespuestaMensajeRemoverInputModel struct {
	IdPreguntaRespuesta *int    `json:"idPreguntaRespuesta" validate:"required"`
	IdUsuarioEmisor     *int    `json:"idUsuarioEmisor" validate:"required"`
	IdMensaje           *string `json:"idMensaje" validate:"required"`
}

type PreguntasRespuestaMensajesObtenerInputModel struct {
	IdPreguntaRespuesta *int `json:"idPreguntaRespuesta" validate:"required"`
}

type PreguntasRespuestasBuscarInputModel struct {
	Busqueda *string `json:"busqueda" validate:"required"`
}

type PreguntasRespuestasObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
}
