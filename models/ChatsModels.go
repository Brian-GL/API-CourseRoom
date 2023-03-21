package models

type ChatRegistrarInputModel struct {
	IdUsuarioEmisor   *int `json:"idUsuarioEmisor" validate:"required"`
	IdUsuarioReceptor *int `json:"idUsuarioReceptor" validate:"required"`
}

type ChatRemoverInputModel struct {
	IdChat    *int `json:"idChat" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type ChatMensajeRegistrarInputModel struct {
	IdChat          *int    `json:"idChat" validate:"required"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" validate:"required"`
	Mensaje         *string `json:"mensaje" validate:"required"`
	Archivo         *string `json:"archivo"`
}

type ChatMensajeRemoverInputModel struct {
	IdChat          *int `json:"idChat" validate:"required"`
	IdUsuarioEmisor *int `json:"idUsuarioEmisor" validate:"required"`
	IdMensaje       *int `json:"idMensaje" validate:"required"`
}

type ChatMensajesObtenerInputModel struct {
	IdChat          *int  `json:"idChat" validate:"required"`
	IdUsuarioLector *int  `json:"idUsuarioLector" validate:"required"`
	Leidos          *bool `json:"leidos"`
}

type ChatsBuscarInputModel struct {
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" validate:"required"`
	Nombre          *string `json:"nombre"`
	Paterno         *string `json:"paterno"`
	Materno         *string `json:"materno"`
}

type ChatsObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
}
