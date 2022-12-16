package models

import "time"

type GrupoActualizarInputModel struct {
	IdGrupo     *int    `json:"idGrupo" validate:"required"`
	IdCurso     *int    `json:"idCurso" validate:"required"`
	Nombre      *string `json:"nombre" validate:"required"`
	Descripcion *string `json:"descripcion"`
	Imagen      *string `json:"imagen"`
}

type GrupoInputModel struct {
	IdGrupo *int `json:"idGrupo" validate:"required"`
}

type GrupoArchivoCompartidoRegistrarInputModel struct {
	IdGrupo       *int    `json:"idGrupo" validate:"required"`
	IdUsuario     *int    `json:"idUsuario" validate:"required"`
	NombreArchivo *string `json:"nombreArchivo" validate:"required"`
	Archivo       *string `json:"archivo" validate:"required"`
}

type GruposMensajesObtenerInputModel struct {
	IdGrupo       *int  `json:"idGrupo" validate:"required"`
	UltimoMensaje *bool `json:"ultimoMensaje" validate:"required"`
}

type GruposObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type GrupoMiembrosObtenerInputModel struct {
	IdGrupo   *int `json:"idGrupo" validate:"required"`
	IdUsuario *int `json:"idUsuario"`
}

type GrupoTareaPendienteDetalleObtenerInputModel struct {
	IdTareaPendiente *int `json:"idTareaPendiente" validate:"required"`
}

type GrupoTareaPendienteEstatusActualizarInputModel struct {
	IdGrupo                 *int `json:"idGrupo" validate:"required"`
	IdTareaPendiente        *int `json:"idTareaPendiente" validate:"required"`
	IdUsuarioReceptor       *int `json:"idUsuarioReceptor" validate:"required"`
	IdEstatusTareaPendiente *int `json:"idEstatusTareaPnediente" validate:"required"`
}

type GrupoMiembroRemoverInputModel struct {
	IdGrupo    *int `json:"idGrupo" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
	IdUsuario  *int `json:"idUsuario" validate:"required"`
}

type GrupoMiembroRegistrarInputModel struct {
	IdGrupo    *int `json:"idGrupo" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdUsuario  *int `json:"idUsuario" validate:"required"`
}

type GrupoTareaPendienteActualizarInputModel struct {
	IdGrupo          *int    `json:"idGrupo" validate:"required"`
	IdUsuario        *int    `json:"idUsuario" validate:"required"`
	IdTareaPendiente *int    `json:"idTareaPendiente" validate:"required"`
	Nombre           *string `json:"nombre" validate:"required"`
	Descripcion      *string `json:"descripcion"`
}

type GrupoTareaPendienteRegistrarInputModel struct {
	IdGrupo           *int       `json:"idGrupo" validate:"required"`
	IdUsuarioEmisor   *int       `json:"idUsuarioEmisor" validate:"required"`
	IdUsuarioReceptor *int       `json:"idUsuarioReceptor" validate:"required"`
	Nombre            *string    `json:"nombre" validate:"required"`
	Descripcion       *string    `json:"descripcion"`
	FechaFinalizacion *time.Time `json:"fechaFinalizacion"`
}

type GrupoRegistrarInputModel struct {
	IdCurso     *int    `json:"idCurso" validate:"required"`
	Nombre      *string `json:"nombre" validate:"required"`
	Descripcion *string `json:"descripcion"`
	Imagen      *string `json:"Imagen"`
}

type GrupoRemoverInputModel struct {
	IdGrupo    *int `json:"idGrupo" validate:"required"`
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type GrupoAbandonarActualizarInputModel struct {
	IdGrupo   *int `json:"idGrupo" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type GrupoArchivoCompartidoRemoverInputModel struct {
	IdArchivoCompartido *int `json:"idArchivocompartido" validate:"required"`
	IdGrupo             *int `json:"idGrupo" validate:"required"`
	IdUsuario           *int `json:"idUsuario" validate:"required"`
}

type GrupoDetalleObtenerInputModel struct {
	IdGrupo *int `json:"idGrupo" validate:"required"`
}

type GrupoMensajeRegistrarInputModel struct {
	IdGrupo         *int    `json:"idGrupo" validate:"required"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" validate:"required"`
	Mensaje         *string `json:"mensaje" validate:"required"`
	Archivo         *string `json:"archivo"`
}

type GrupoMensajeRemoverInputModel struct {
	IdGrupo         *int `json:"idGrupo" validate:"required"`
	IdUsuarioEmisor *int `json:"idUsuarioEmisor" validate:"required"`
	IdMensaje       *int `json:"idMensaje" validate:"required"`
}
