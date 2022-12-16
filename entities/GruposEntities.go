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

type GrupoRegistrarEntity struct {
	IdCurso     *int    `json:"idCurso"`
	Nombre      *string `json:"nombre"`
	Descripcion *string `json:"descripcion"`
	Imagen      *string `json:"Imagen"`
	Codigo      *int    `json:"Codigo"`
	Mensaje     *string `json:"Mensaje"`
}

type GrupoRemoverEntity struct {
	IdGrupo    *int    `json:"idGrupo"`
	IdCurso    *int    `json:"idCurso"`
	IdProfesor *int    `json:"idProfesor"`
	Codigo     *int    `json:"Codigo"`
	Mensaje    *string `json:"Mensaje"`
}

type GrupoAbandonarActualizarEntity struct {
	IdArchivoCompartido *int    `json:"idArchivoCompartido"`
	IdGrupo             *int    `json:"idGrupo"`
	IdUsuario           *int    `json:"idUsuario"`
	Codigo              *int    `json:"Codigo"`
	Mensaje             *string `json:"Mensaje"`
}

type GrupoDetalleObtenerEntity struct {
	IdGrupo            *int       `json:"idGrupo"`
	Nombre             string     `json:"nombre"`
	Descripcion        *string    `json:"descripcion"`
	Imagen             *string    `json:"imagen"`
	IdCurso            *int       `json:"idCurso"`
	Curso              string     `json:"curso"`
	ImagenCurso        *string    `json:"imagenCurso"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type GrupoMensajeRegistrarEntity struct {
	IdGrupo         *int    `json:"idGrupo"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor"`
	Mensaje         *string `json:"mensaje"`
	Arvhivo         *string `json:"archivo"`
	Codigo          *int    `json:"codigo"`
}

type GrupoMensajeRemoverEntity struct {
	IdGrupo         *int    `json:"idGrupo"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor"`
	IdMensaje       *string `json:"idMensaje"`
	Arvhivo         *string `json:"archivo"`
	Codigo          *int    `json:"codigo"`
}
