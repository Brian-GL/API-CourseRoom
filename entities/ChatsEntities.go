package entities

import "time"

type ChatMensajesObtenerEntity struct {
	IdMensaje           int       `json:"idMensaje"`
	Mensaje             string    `json:"mensaje"`
	Archivo             *string   `json:"archivo"`
	IdUsuarioEmisor     int       `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string    `json:"nombreUsuarioEmisor"`
	FechaRegistro       time.Time `json:"fechaRegistro"`
}

type ChatsBuscarEntity struct {
	IdChat            int        `json:"idChat"`
	IdUsuarioReceptor int        `json:"idUsuarioReceptor"`
	Receptor          string     `json:"receptor"`
	ImagenReceptor    *string    `json:"imagenReceptor"`
	Mensaje           *string    `json:"mensaje"`
	FechaRegistro     time.Time  `json:"fechaRegistro"`
	FechaEnvio        *time.Time `json:"fechaEnvio"`
}

type ChatsObtenerEntity struct {
	IdChat            int        `json:"idChat"`
	IdUsuarioReceptor int        `json:"idUsuarioReceptor"`
	Receptor          string     `json:"receptor"`
	ImagenReceptor    *string    `json:"imagenReceptor"`
	FechaRegistro     time.Time  `json:"fechaRegistro"`
	Mensaje           *string    `json:"mensaje"`
	FechaEnvio        *time.Time `json:"fechaEnvio"`
}
