package entities

type UsuarioAccesoObtenerEntity struct {
	IdUsuario     *int `json:"idUsuario"`
	IdTipoUsuario *int `json:"idTipoUsuario"`
}

type UsuarioCuentaObtenerEntity struct {
	CorreoElectronico string  `json:"correoElectronico"`
	Contrasena        string  `json:"contrasena"`
	Imagen            *string `json:"imagen"`
	ChatsConmigo      bool    `json:"chatsConmigo"`
	MostrarAvisos     bool    `json:"mostrarAvisos"`
}
