package models

import "time"

type UsuarioActualizarInputModel struct {
	IdUsuario       *int       `json:"idUsuario" validate:"required"`
	Nombre          *string    `json:"nombre" validate:"required"`
	Paterno         *string    `json:"paterno" validate:"required"`
	Materno         *string    `json:"materno"`
	FechaNacimiento *time.Time `json:"fechaNacimiento"`
	Genero          *string    `json:"genero"`
	Descripcion     *string    `json:"descripcion"`
	IdLocalidad     *int       `json:"idLocalidad" validate:"required"`
}

type UsuarioRegistrarInputModel struct {
	Nombre            *string    `json:"nombre" validate:"required"`
	Paterno           *string    `json:"paterno" validate:"required"`
	Materno           *string    `json:"materno"`
	FechaNacimiento   *time.Time `json:"fechaNacimiento" validate:"required"`
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

type UsuarioRemoverInputModel struct {
	IdUsuario     *int `json:"idUsuario" validate:"required"`
	IdTipoUsuario *int `json:"idTipoUsuario" validate:"required"`
}

type UsuarioInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type UsuarioAccesoObtenerInputModel struct {
	CorreoElectronico *string `json:"correoElectronico" validate:"required,email"`
	Contrasena        *string `json:"contrasena" validate:"required,base64"`
}

type UsuarioCredencialObtenerInputModel struct {
	CorreoElectronico *string `json:"correoElectronico" validate:"required,email"`
}

type UsuarioCuentaActualizarInputModel struct {
	IdUsuario         *int    `json:"idUsuario" validate:"required"`
	CorreoElectronico *string `json:"correoElectronico" validate:"required,email"`
	Contrasena        *string `json:"contrasena" validate:"required,base64"`
	ChatsConmigo      *bool   `json:"chatsConmigo" validate:"required"`
	MostrarAvisos     *bool   `json:"mostrarAvisos" validate:"required"`
	Imagen            *string `json:"imagen"`
}

type UsuarioDesempenoRegistrarInputModel struct {
	IdUsuario                    *int     `json:"idUsuario" validate:"required"`
	IdTarea                      *int     `json:"idTarea" validate:"required"`
	Calificacion                 *float64 `json:"calificacion" validate:"required"`
	PromedioCurso                *float64 `json:"promedioCurso" validate:"required"`
	PrediccionPromedioCurso      *float64 `json:"prediccionPromedioCurso"`
	RumboPromedioCurso           *string  `json:"rumboPromedioCurso"`
	PromedioGeneral              *float64 `json:"promedioGeneral" validate:"required"`
	PrediccionPromedioGeneral    *float64 `json:"prediccionPromedioGeneral"`
	RumboPromedioGeneral         *string  `json:"rumboPromedioGeneral"`
	PuntualidadCurso             *float64 `json:"puntualidadCurso" validate:"required"`
	PrediccionPuntualidadCurso   *float64 `json:"prediccionPuntualidadCurso"`
	RumboPuntualidadCurso        *string  `json:"rumboPuntualidadCurso"`
	PuntualidadGeneral           *float64 `json:"puntualidadGeneral" validate:"required"`
	PrediccionPuntualidadGeneral *float64 `json:"prediccionPuntualidadGeneral"`
	RumboPuntualidadGeneral      *string  `json:"rumboPuntualidadGeneral"`
}

type UsuarioNuevaPuntualidadCursoObtenerInputModel struct {
	IdCurso     *int     `json:"idCurso" validate:"required"`
	IdUsuario   *int     `json:"idUsuario" validate:"required"`
	Puntualidad *float64 `json:"puntualidad" validate:"required"`
}

type UsuarioNuevaPuntualidadGeneralObtenerInputModel struct {
	IdUsuario   *int     `json:"idUsuario" validate:"required"`
	Puntualidad *float64 `json:"puntualidad" validate:"required"`
}

type UsuarioNuevoPromedioCursoObtenerInputModel struct {
	IdCurso      *int     `json:"idCurso" validate:"required"`
	IdUsuario    *int     `json:"idUsuario" validate:"required"`
	Calificacion *float64 `json:"puntualidad" validate:"required"`
}

type UsuarioNuevoPromedioGeneralObtenerInputModel struct {
	IdUsuario    *int     `json:"idUsuario" validate:"required"`
	Calificacion *float64 `json:"puntualidad" validate:"required"`
}

type UsuariosBuscarInputModel struct {
	Nombre  *string `json:"nombre"`
	Paterno *string `json:"paterno"`
	Materno *string `json:"materno"`
}

type UsuarioSesionInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
	IdSesion  *int `json:"idSesion" validate:"required"`
}

type UsuarioSesionRegistrarInputModel struct {
	IdUsuario    *int    `json:"idUsuario" validate:"required"`
	Dispositivo  *string `json:"dispositivo"`
	Fabricante   *string `json:"fabricante"`
	DireccionIP  *string `json:"direccionIP"`
	DireccionMAC *string `json:"direccionMAC"`
	UserAgent    *string `json:"userAgent"`
	Navegador    *string `json:"navegador"`
}

type UsuarioSesionesObtenerInputModel struct {
	IdUsuario *int  `json:"idUsuario" validate:"required"`
	Activa    *bool `json:"activa"`
}

type UsuarioTematicaInputModel struct {
	IdUsuario  *int `json:"idUsuario" validate:"required"`
	IdTematica *int `json:"idTematica" validate:"required"`
}

type UsuarioTematicasObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
}
