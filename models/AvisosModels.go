package models

import "time"

type AvisosObtenerInputModel struct {
	IdUsuario *int  `json:"idUsuario" validate:"required"`
	Leido     *bool `json:"leido"`
}

type AvisoInputModel struct {
	IdAviso *int `json:"idAviso" validate:"required"`
}

type AvisoAccionInputModel struct {
	IdAviso   *int `json:"idAviso" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type AvisoRegistrarInputModel struct {
	IdUsuario   *int    `json:"idUsuario" validate:"required"`
	Aviso       *string `json:"aviso" validate:"required"`
	Descripcion *string `json:"descripcion" validate:"required"`
	IdTipoAviso *int    `json:"idTipoAviso" validate:"required"`
}

type AvisosValidarInputModel struct {
	IdUsuario          *int       `json:"idUsuario" validate:"required"`
	FechaVisualizacion *time.Time `json:"fechaVisualizacion" validate:"required"`
}

type AvisoPlagioProfesorRegistrarInputModel struct {
	IdProfesor    *int    `json:"idProfesor" validate:"required"`
	IdUsuario     *int    `json:"idUsuario" validate:"required"`
	IdTarea       *int    `json:"idTarea" validate:"required"`
	NombreArchivo *string `json:"nombreArchivo" validate:"required"`
}
