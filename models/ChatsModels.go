package models

type ChatRegistrarInputModel struct {
	IdUsuarioEmisor   *int `json:"idUsuarioEmisor" binding:"required"`
	IdUsuarioReceptor *int `json:"idUsuarioReceptor" binding:"required"`
}

type ChatRemoverInputModel struct {
	IdChat    *int `json:"idChat" binding:"required"`
	IdUsuario *int `json:"idUsuario" binding:"required"`
}

type ChatMensajeRegistrarInputModel struct {
	IdChat          *int    `json:"idChat" binding:"required"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" binding:"required"`
	Mensaje         *string `json:"mensaje" binding:"required"`
	Archivo         *string `json:"archivo" binding:"required"`
}

type ChatMensajeRemoverInputModel struct {
	IdChat          *int `json:"idChat" binding:"required"`
	IdUsuarioEmisor *int `json:"idUsuarioEmisor" binding:"required"`
	IdMensaje       *int `json:"idMensaje" binding:"required"`
}

type ChatMensajesObtenerInputModel struct {
	IdChat *int    `json:"idChat" binding:"required"`
	Ultimo *string `json:"idMensaje"`
}

type ChatsBuscarInputModel struct {
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" binding:"required"`
	Nombre          *string `json:"nombre" binding:"required"`
	Paterno         *string `json:"paterno" binding:"required"`
	Materno         *string `json:"materno" binding:"required"`
}

type ChatsObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" binding:"required"`
}
