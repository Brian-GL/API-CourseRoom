package models

type ArchivoActualizarInputModel struct {
	Archivo       *string `json:"archivo" validate:"required"`
	IdTipoArchivo *int    `json:"idTipoArchivo" validate:"required"`
	IdRegistro    *int    `json:"idRegistro"`
}
