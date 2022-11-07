package entities

import "time"

type ChatMensajesObtenerEntity struct {
	Ultimo              bool      `json:"ultimo"`
	IdMensaje           int       `json:"idMensaje"`
	Mensaje             string    `json:"mensaje"`
	Archivo             *string   `json:"archivo"`
	IdUsuarioEmisor     int       `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string    `json:"nombreUsuarioEmisor"`
	FechaRegistro       time.Time `json:"fechaRegistro"`
}

type ChatsBuscarEntity struct {
	IdChat            int
	IdUsuarioReceptor int
	Receptor          string
	ImagenReceptor    *string
	Mensaje           *string
	FechaRegistro     time.Time  `json:"fechaRegistro"`
	FechaEnvio        *time.Time `json:"fechaEnvio"`
}

type ChatsObtenerEntity struct {
	IdChat            int
	IdUsuarioReceptor int
	Receptor          string
	ImagenReceptor    *string
	Mensaje           *string
	FechaRegistro     time.Time  `json:"fechaRegistro"`
	FechaEnvio        *time.Time `json:"fechaEnvio"`
}
