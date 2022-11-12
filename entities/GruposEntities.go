package entities

import "time"

type GrupoArchivosCompartidosObtenerEntity struct {
	IdArchivoCompartido int       `json:"idArchivoCompartido"`
	Nombre              string    `json:"nombre"`
	Archivo             string    `json:"archivo"`
	IdUsuarioEmisor     int       `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string    `json:"nombreUsuarioEmisor"`
	FechaRegistro       time.Time `json:"fechaRegistro"`
}
