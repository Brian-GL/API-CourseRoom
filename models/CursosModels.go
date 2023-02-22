package models

import "time"

type CursoRemoverInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type CursoRegistrarInputModel struct {
	Nombre      *string `json:"nombre" validate:"required"`
	Descripcion *string `json:"descripcion" validate:"required"`
	Imagen      *string `json:"imagen"`
	IdProfesor  *int    `json:"idProfesor" validate:"required"`
}

type CursoGruposObtenerInputModel struct {
	IdCurso *int  `json:"idCurso" validate:"required"`
	Activo  *bool `json:"activo" validate:"required"`
}

type CursoActualizarInputModel struct {
	IdCurso     *int    `json:"idCurso" validate:"required"`
	IdProfesor  *int    `json:"idProfesor"`
	Nombre      *string `json:"nombre" validate:"required"`
	Descripcion *string `json:"descripcion" validate:"required"`
	Imagen      *string `json:"imagen" validate:"required"`
}

type CursoAbandonarActualizarInputModel struct {
	IdCurso   *int    `json:"idCurso" validate:"required"`
	IdUsuario *int    `json:"idUsuario" validate:"required"`
	Codigo    *int    `json:"codigo"`
	Mensaje   *string `json:"mensaje"`
}

type CursoCuestionarioContestarInputModel struct {
	IdCurso   *int    `json:"idCurso" validate:"required"`
	IdUsuario *int    `json:"idUsuario" validate:"required"`
	Codigo    *int    `json:"codigo"`
	Mensaje   *string `json:"mensaje"`
}

type CursoCuestionarioAbandonarActualizarInputModel struct {
	IdCurso   *int    `json:"idCurso" validate:"required"`
	IdUsuario *int    `json:"idUsuario" validate:"required"`
	Codigo    *int    `json:"codigo"`
	Mensaje   *string `json:"mensaje"`
}

type CursoDesempenoObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoEstudianteRegistrarInputModel struct {
	IdCurso   *int `json:"idCurso" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
	Codigo    *int `json:"codigo"`
	Mensaje   *int `json:"mensaje"`
}

type CursoEstudianteDetalleObtenerInputModel struct {
	IdCurso                 *int       `json:"idCurso" validate:"required"`
	IdUsuario               *int       `json:"idUsuario" validate:"required"`
	Nombre                  *string    `json:"nombre"`
	Descripcion             *string    `json:"descripcion"`
	Imagen                  *string    `json:"imagen"`
	IdProfesor              *int       `json:"idProfesor"`
	NombreProfesor          *string    `json:"nombreProfesor"`
	ImagenProfesor          *string    `json:"imagenProfesor"`
	FechaRegistroCurso      *time.Time `json:"fechaRegistrocurso"`
	FechaActualizacionCurso *time.Time `json:"fechaActualizacionCurso"`
	Finalizado              *bool      `json:"finalizado"`
	FechaRegistro           *time.Time `json:"fechaRegistro"`
	FechaActualizacion      *string    `json:"fechaActualizacion"`
	Estatus                 *string    `json:"estatus"`
	DescripcionEstatus      *string    `json:"descripcionEstatus"`
}

type CursoFinalizarActualizarInputModel struct {
	IdCurso    *int    `json:"idCurso" validate:"required"`
	IdProfesor *int    `json:"idProfesor" validate:"required"`
	Codigo     *int    `json:"Codigo"`
	Mensaje    *string `json:"mensaje"`
}

type CursoMaterialRegistrarInputModel struct {
	IdCurso       *int    `json:"idCurso" validate:"required"`
	IdUsuario     *int    `json:"idUsuario" validate:"required"`
	NombreArchivo *string `json:"nombreArchivo" validate:"required"`
	Archivo       *string `json:"archivo" validate:"required"`
	Codigo        *int    `json:"codigo"`
	Mensaje       *string `json:"mensaje"`
}

type CursoMaterialRemoverInputModel struct {
	IdMaterial *int    `json:"idMaterial" validate:"required"`
	IdCurso    *int    `json:"idCurso" validate:"required"`
	IdUsuario  *int    `json:"idUsuario" validate:"required"`
	Codigo     *int    `json:"codigo"`
	Mensaje    *string `json:"mensaje"`
}

type CursoMaterialesObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoMensajeRegistrarInputModel struct {
	IdCurso         *int    `json:"idCurso" validate:"required"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" validate:"required"`
	Mensaje         *string `json:"mensaje" validate:"required"`
	Archivo         *string `json:"archivo" validate:"required"`
	Codigo          *int    `json:"codigo"`
}

type CursoMensajeRemoverInputModel struct {
	IdCurso         *int    `json:"idCurso" validate:"required"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" validate:"required"`
	IdMensaje       *int    `json:"idmensaje" validate:"required"`
	Codigo          *int    `json:"codigo"`
	Mensaje         *string `json:"mensaje"`
}

type CursoMensajesObtenerInputModel struct {
	IdCurso             *int       `json:"idCurso" validate:"required"`
	UltimoMensaje       *bool      `json:"ultimoMensaje" validate:"required"`
	IdMensaje           *int       `json:"idmensaje"`
	Mensaje             *string    `json:"mensaje"`
	Archivo             *string    `json:"archivo"`
	IdUsuarioEmisor     *int       `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor *string    `json:"nombreUsuarioEmisor"`
	FechaRegistro       *time.Time `json:"fechaRegistro"`
}

type CursoEstudianteRemoverInputModel struct {
	IdCurso    *int    `json:"idCurso" validate:"required"`
	IdProfesor *int    `json:"idProfesor" validate:"required"`
	IdUsuario  *int    `json:"idUsuario" validate:"required"`
	Codigo     *int    `json:"codigo"`
	Mensaje    *string `json:"mensaje"`
}

type CursoEstudianteObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoProfesorDetalleObtenerInputModel struct {
	IdCurso            *int       `json:"idCurso" validate:"required"`
	Nombre             *string    `json:"nombre"`
	Descripcion        *string    `json:"descripcion"`
	Imagen             *string    `json:"imagen"`
	IdProfesor         *int       `json:"idProfesor"`
	NombreProfesor     *string    `json:"nombreProfesor"`
	ImagenProfesor     *string    `json:"imagenProfesor"`
	FechaRegistro      *time.Time `json:"fechaResgistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
	Puntaje            *float32   `json:"puntaje"`
	Finalizado         *bool      `json:"finalizacion"`
}

type CursoProfesorTareasObtenerInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type CursoPromedioObtenerInputModel struct {
	IdCurso       *int     `json:"idCurso" validate:"required"`
	PromedioCurso *float32 `json:"promedioCurso"`
}

type CursoBuscarInputModel struct {
	Busqueda       *string    `json:"busqueda" validate:"required"`
	IdUsuario      *int       `json:"idUsuario" validate:"required"`
	IdCurso        *int       `json:"idCurso"`
	Curso          *string    `json:"curso"`
	ImagenCurso    *string    `json:"imagencurso"`
	IdProfesor     *int       `json:"idProfesor"`
	Profesor       *string    `json:"profesor"`
	ImagenProfesor *string    `json:"imagenProfesor"`
	ListaTematicas *string    `json:"listaTematica"`
	FechaRegistro  *time.Time `json:"fechaRegistro"`
	Puntaje        *float32   `json:"puntaje"`
	FechaIngreso   *time.Time `json:"fechaIngreso"`
	Estatus        *string    `json:"estatus"`
}

type CursoObtenerInputModel struct {
	IdUsuario        *int `json:"idUsuario" validate:"required"`
	IdEstatusUsuario *int `json:"idEstatusUsuario" validate:"required"`
}

type CursoNuevoObtenerInputModel struct {
	IdUsuario        *int `json:"idUsuario" validate:"required"`
	NumeroResultados *int `json:"numeroResultados" validate:"required"`
}

type CursoProfesorObtenerInputModel struct {
	IdProfesor *int  `json:"idProfesor" validate:"required"`
	Finalizado *bool `json:"finalizado"`
}

type CursoTareasEstudianteObtenerInputModel struct {
	IdCurso   *int `json:"idCurso" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type CursoTematicaRegistrarInputModel struct {
	IdCurso    *int    `json:"idCurso" validate:"required"`
	IdTematica *int    `json:"idTematica" validate:"required"`
	Codigo     *int    `json:"codigo"`
	Mensaje    *string `json:"mensaje"`
}

type CursoTematicaRemoverInputModel struct {
	IdCurso    *int    `json:"idCurso" validate:"required"`
	IdTematica *int    `json:"idTematica" validate:"required"`
	Codigo     *int    `json:"codigo"`
	Mensaje    *string `json:"mensaje"`
}

type CursoTematicaObtenerInputModel struct {
	IdCurso    *int    `json:"idCurso" validate:"required"`
	IdTematica *int    `json:"idTematica"`
	Tematica   *string `json:"tematica"`
}

type CursoEstudianteDesempenoObtenerInputModel struct {
	IdCurso   *int `json:"idCurso" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type CursoEstudiantesSinGrupoObtenerInputModel struct {
	IdCurso      *int       `json:"idCurso" validate:"required"`
	IdUsuario    *int       `json:"idUsuario"`
	Estudiante   *string    `json:"estudiante"`
	FechaIngreso *time.Time `json:"fechaIngreso"`
}

type CursoEstudianteFinalizarActualizarInputModel struct {
	IdCurso   *int    `json:"idCurso" validate:"required"`
	IdUsuario *int    `json:"idUsuario" validate:"required"`
	Codigo    *int    `json:"codigo"`
	Mensaje   *string `json:"mensaje"`
}

type CursoCuestionarioRespuestaRegistrarInputModel struct {
	IdCurso    *int    `json:"idCurso" validate:"required"`
	IdUsuario  *int    `json:"idUsuario" validate:"required"`
	IdPregunta *int    `json:"idPregunta" validate:"required"`
	Puntaje    *int    `json:"puntaje" validate:"required"`
	Codigo     *int    `json:"codigo"`
	Mensaje    *string `json:"mensaje"`
}
