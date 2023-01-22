package entities

import "time"

type AvisoDetalleObtenerEntity struct {
	Aviso              string     `json:"aviso"`
	Descripcion        string     `json:"descripcion"`
	Estatus            string     `json:"estatus"`
	TipoAviso          string     `json:"tipoAviso"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type AvisosObtenerEntity struct {
	IdAviso       int       `json:"idAviso"`
	Aviso         string    `json:"aviso"`
	Estatus       string    `json:"estatus"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	TipoAviso     string    `json:"tipoAviso"`
}

type AvisosValidarEntity struct {
	IdAviso       int       `json:"idAviso"`
	TipoAviso     string    `json:"tipoAviso"`
	Aviso         string    `json:"aviso"`
	FechaRegistro time.Time `json:"fechaRegistro"`
}
