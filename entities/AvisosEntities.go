package entities

import "time"

type AvisosObtenerEntity struct {
	IdAviso       int       `json:"idAviso"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	Aviso         string    `json:"aviso"`
	Estatus       string    `json:"estatus"`
	TipoAviso     string    `json:"tipoAviso"`
}

type AvisoDetalleObtenerEntity struct {
	Aviso              string     `json:"aviso"`
	Descripcion        string     `json:"descripcion"`
	Estatus            string     `json:"estatus"`
	TipoAviso          string     `json:"tipoAviso"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type AvisosValidarEntity struct {
	IdAviso       int       `json:"idAviso"`
	TipoAviso     string    `json:"tipoAviso"`
	Aviso         string    `json:"aviso"`
	FechaRegistro time.Time `json:"fechaRegistro"`
}
