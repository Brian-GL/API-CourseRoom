package entities

import "time"

type Estado struct {
	IdEstado int    `json:"idEstado"`
	Estado   string `json:"estado"`
}

type EstatusTareaPendiente struct {
	IdEstatus int    `json:"idEstatus"`
	Estatus   string `json:"estatus"`
}

type AvisosObtener struct {
	IdAviso       int       `json:"idAviso"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	Aviso         string    `json:"aviso"`
	Estatus       string    `json:"estatus"`
	TipoAviso     string    `json:"tipoAviso"`
}
