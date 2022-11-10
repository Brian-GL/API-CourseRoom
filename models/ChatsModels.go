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
	Archivo         *string `json:"archivo" validate:"required"`
}

type ChatMensajeRemoverInputModel struct {
	IdChat          *int `json:"idChat" validate:"required"`
	IdUsuarioEmisor *int `json:"idUsuarioEmisor" validate:"required"`
	IdMensaje       *int `json:"idMensaje" validate:"required"`
}

type ChatMensajesObtenerInputModel struct {
	IdChat *int    `json:"idChat" validate:"required"`
	Ultimo *string `json:"idMensaje"`
}

type ChatsBuscarInputModel struct {
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" validate:"required"`
	Nombre          *string `json:"nombre" validate:"required"`
	Paterno         *string `json:"paterno" validate:"required"`
	Materno         *string `json:"materno" validate:"required"`
}

type ChatsObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
}
