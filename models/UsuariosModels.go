package models

import "time"

type UsuarioRegistrarInputModel struct {
	Nombre            *string    `json:"nombre" validate:"required"`
	Paterno           *string    `json:"paterno" validate:"required"`
	Materno           *string    `json:"materno"`
	FechaNacimiento   *time.Time `json:"fechaNacimiento"`
	Genero            *string    `json:"genero"`
	Descripcion       *string    `json:"descripcion"`
	IdLocalidad       *int       `json:"idLocalidad" validate:"required"`
	IdTipoUsuario     *int       `json:"idTipoUsuario" validate:"required"`
	CorreoElectronico *string    `json:"correoElectronico" validate:"required,email"`
	Contrasena        *string    `json:"contrasena" validate:"required,base64"`
	ChatsConmigo      *bool      `json:"chatsConmigo" validate:"required"`
	MostrarAvisos     *bool      `json:"mostrarAvisos" validate:"required"`
	Imagen            *string    `json:"imagen"`
}
