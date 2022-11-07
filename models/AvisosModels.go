package models

import "time"

type AvisosObtenerInputModel struct {
	IdUsuario *int  `json:"idUsuario" binding:"required"`
	Leido     *bool `json:"leido"`
}

type AvisoInputModel struct {
	IdAviso *int `json:"idAviso" binding:"required"`
}

type AvisoAccionInputModel struct {
	IdAviso   *int `json:"idAviso" binding:"required"`
	IdUsuario *int `json:"idUsuario" binding:"required"`
}

type AvisoRegistrarInputModel struct {
	IdUsuario   *int    `json:"idUsuario" binding:"required"`
	Aviso       *string `json:"aviso" binding:"required"`
	Descripcion *string `json:"descripcion" binding:"required"`
	IdTipoAviso *int    `json:"idTipoAviso" binding:"required"`
}

type AvisosValidarInputModel struct {
	IdUsuario          *int       `json:"idUsuario" binding:"required"`
	FechaVisualizacion *time.Time `json:"fechaVisualizacion"`
}

type AvisoPlagioProfesorRegistrarInputModel struct {
	IdProfesor    *int    `json:"idProfesor" binding:"required"`
	IdUsuario     *int    `json:"idUsuario" binding:"required"`
	IdTarea       *int    `json:"idTarea" binding:"required"`
	NombreArchivo *string `json:"nombreArchivo" binding:"required"`
}
