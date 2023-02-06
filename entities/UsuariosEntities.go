package entities

import "time"

type UsuarioAccesoObtenerEntity struct {
	IdUsuario     int `json:"idUsuario"`
	IdTipoUsuario int `json:"idTipoUsuario"`
}

type UsuarioCuentaObtenerEntity struct {
	CorreoElectronico string  `json:"correoElectronico"`
	Contrasena        string  `json:"contrasena"`
	Imagen            *string `json:"imagen"`
	ChatsConmigo      bool    `json:"chatsConmigo"`
	MostrarAvisos     bool    `json:"mostrarAvisos"`
}

type UsuarioDesempenoObtenerEntity struct {
	IdCurso                     int       `json:"idCurso"`
	Curso                       string    `json:"curso"`
	ImagenCurso                 *string   `json:"imagenCurso"`
	IdTarea                     int       `json:"idTarea"`
	Tarea                       string    `json:"tarea"`
	Calificacion                float32   `json:"calificacion"`
	PromedioCalificacionCurso   float32   `json:"promedioCalificacionCurso"`
	MedianaCalificacionCurso    float32   `json:"medianaCalificacionCurso"`
	ResultadoCalificacionCurso  float32   `json:"resultadoCalificacionCurso"`
	PrediccionCalificacionCurso *float32  `json:"prediccionCalificacionCurso"`
	Puntualidad                 float32   `json:"puntualidad"`
	PromedioPuntualidadCurso    float32   `json:"promedioPuntualidadCurso"`
	MedianaPuntualidadCurso     float32   `json:"medianaPuntualidadCurso"`
	ResultadoPuntualidadCurso   float32   `json:"resultadoPuntualidadCurso"`
	PrediccionPuntualidadCurso  *float32  `json:"prediccionPuntualidadCurso"`
	FechaRegistro               time.Time `json:"fechaRegistro"`
}

type UsuarioDetalleObtenerEntity struct {
	Nombre             string     `json:"nombre"`
	Paterno            string     `json:"paterno"`
	Materno            *string    `json:"materno"`
	Descripcion        *string    `json:"descripcion"`
	FechaNacimiento    *time.Time `json:"fechaNacimiento"`
	Genero             *string    `json:"genero"`
	IdLocalidad        *int       `json:"idLocalidad"`
	IdEstado           *int       `json:"idEstado"`
	TipoUsuario        string     `json:"tipoUsuario"`
	PromedioGeneral    *float64   `json:"promedioGeneral"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type UsuarioPuntualidadEntity struct {
	NuevaPuntualidad *float64 `json:"nuevaPuntualidad"`
}

type UsuarioPromedioEntity struct {
	NuevoPromedio *float64 `json:"nuevoPromedio"`
}

type UsuariosBuscarEntity struct {
	IdUsuario         int     `json:"idUsuario"`
	NombreCompleto    string  `json:"nombreCompleto"`
	Imagen            *string `json:"imagen"`
	CorreoElectronico string  `json:"correoElectronico"`
	TipoUsuario       string  `json:"tipoUsuario"`
}

type UsuarioSesionValidarEntity struct {
	Activo *bool `json:"activo"`
}

type UsuarioSesionesObtenerEntity struct {
	IdSesion           int        `json:"idSesion"`
	Dispositivo        *string    `json:"dispositivo"`
	Fabricante         *string    `json:"fabricante"`
	DireccionIP        *string    `json:"direccionIP"`
	DireccionMAC       *string    `json:"direccionMAC"`
	UserAgent          *string    `json:"userAgent"`
	Navegador          *string    `json:"navegador"`
	Estatus            string     `json:"estatus"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type UsuarioTematicasObtenerEntity struct {
	IdTematica int    `json:"idTematica"`
	Tematica   string `json:"tematica"`
}
