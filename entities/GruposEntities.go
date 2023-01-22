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

type GrupoDetalleObtenerEntity struct {
	Nombre             string     `json:"nombre"`
	Descripcion        string     `json:"descripcion"`
	Imagen             *string    `json:"imagen"`
	IdCurso            int        `json:"idCurso"`
	Curso              string     `json:"curso"`
	ImagenCurso        *string    `json:"imagenCurso"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type GruposMensajesObtenerEntity struct {
	IdMensaje           int       `json:"idMensaje"`
	Mensaje             string    `json:"mensaje"`
	Archivo             *string   `json:"archivo"`
	IdUsuarioEmisor     int       `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string    `json:"nombreUsuarioEmisor"`
	FechaRegistro       time.Time `json:"fechaRegistro"`
}

type GruposObtenerEntity struct {
	IdGrupo       int       `json:"idGrupo"`
	Nombre        string    `json:"nombre"`
	ImagenGrupo   *string   `json:"imagenGrupo"`
	IdCurso       int       `json:"idCurso"`
	NombreCurso   string    `json:"nombreCurso"`
	ImagenCurso   *string   `json:"imagenCurso"`
	FechaRegistro time.Time `json:"fechaRegistro"`
}

type GrupoMiembrosObtenerEntity struct {
	IdUsuario          int        `json:"idUsuario"`
	NombreCompleto     string     `json:"nombreCompleto"`
	Imagen             *string    `json:"imagen"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type GrupoTareasPendientesObtenerEntity struct {
	IdTareaPendiente         int        `json:"idTareaPendiente"`
	Nombre                   string     `json:"nombre"`
	IdUsuarioCreador         int        `json:"idUsuarioCreador"`
	NombreUsuarioCreador     string     `json:"nombreUsuarioCreador"`
	ImagenUsuarioCreador     *string    `json:"imagenUsuarioCreador"`
	IdUsuarioResponsable     int        `json:"idUsuarioResponsable"`
	NombreUsuarioResponsable string     `json:"nombreUsuarioResponsable"`
	ImagenUsuarioResponsable *string    `json:"imagenUsuarioResponsable"`
	FechaRegistro            time.Time  `json:"fechaRegistro"`
	FechaFinalizacion        *time.Time `json:"fechaFinalizacion"`
	Estatus                  string     `json:"estatus"`
}

type GrupoTareaPendienteDetalleObtenerEntity struct {
	Nombre                   string     `json:"nombre"`
	Descripcion              *string    `json:"descripcion"`
	IdUsuarioCreador         int        `json:"idUsuarioCreador"`
	NombreUsuarioCreador     string     `json:"nombreUsuarioCreador"`
	ImagenUsuarioCreador     *string    `json:"imagenUsuarioCreador"`
	IdUsuarioResponsable     int        `json:"idUsuarioResponsable"`
	NombreUsuarioResponsable string     `json:"nombreUsuarioResponsable"`
	ImagenUsuarioResponsable *string    `json:"imagenUsuarioResponsable"`
	FechaRegistro            time.Time  `json:"fechaRegistro"`
	FechaFinalizacion        *time.Time `json:"fechaFinalizacion"`
	FechaActualizacion       *time.Time `json:"fechaActualizacion"`
	Estatus                  string     `json:"estatus"`
}
