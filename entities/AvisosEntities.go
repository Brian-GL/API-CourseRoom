package entities

import "time"

type AvisosObtenerEntity struct {
	IdAviso       int       `json:"idAviso"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	Aviso         string    `json:"aviso"`
	Estatus       string    `json:"estatus"`
	TipoAviso     string    `json:"tipoAviso"`
}
