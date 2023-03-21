package models

type CursoInputModel struct {
	IdCurso   *int `json:"idCurso" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type CursoActualizarInputModel struct {
	IdCurso     *int    `json:"idCurso" validate:"required"`
	IdProfesor  *int    `json:"idProfesor" validate:"required"`
	Nombre      *string `json:"nombre" validate:"required"`
	Descripcion *string `json:"descripcion" validate:"required"`
	Imagen      *string `json:"imagen"`
}

type CursoRegistrarInputModel struct {
	Nombre      *string `json:"nombre" validate:"required"`
	Descripcion *string `json:"descripcion" validate:"required"`
	Imagen      *string `json:"imagen"`
	IdProfesor  *int    `json:"idProfesor" validate:"required"`
}

type CursoRemoverInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type CursoDesempenoObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoDetalleObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoFinalizarActualizarInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type CursoGruposObtenerInputModel struct {
	IdCurso *int  `json:"idCurso" validate:"required"`
	Activo  *bool `json:"activo"`
}

type CursoMaterialRegistrarInputModel struct {
	IdCurso       *int    `json:"idCurso" validate:"required"`
	IdUsuario     *int    `json:"idUsuario" validate:"required"`
	NombreArchivo *string `json:"nombreArchivo" validate:"required"`
	Archivo       *string `json:"archivo"`
}

type CursoMaterialRemoverInputModel struct {
	IdMaterial *int `json:"idMaterial" validate:"required"`
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdUsuario  *int `json:"idUsuario" validate:"required"`
}

type CursoMaterialesObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoMensajeRegistrarInputModel struct {
	IdCurso         *int    `json:"idCurso" validate:"required"`
	IdUsuarioEmisor *int    `json:"idUsuarioEmisor" validate:"required"`
	Mensaje         *string `json:"mensaje" validate:"required"`
	Archivo         *string `json:"archivo"`
}

type CursoMensajeRemoverInputModel struct {
	IdCurso         *int `json:"idCurso" validate:"required"`
	IdUsuarioEmisor *int `json:"idUsuarioEmisor" validate:"required"`
	IdMensaje       *int `json:"idMensaje" validate:"required"`
}

type CursoMensajesObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoEstudianteRemoverInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
	IdUsuario  *int `json:"idUsuario" validate:"required"`
}

type CursoEstudiantesObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoProfesorDetalleObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoProfesorTareasObtenerInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type CursoPromedioObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursosBuscarInputModel struct {
	Busqueda  *string `json:"busqueda" validate:"required"`
	IdUsuario *int    `json:"idUsuario" validate:"required"`
}

type CursosObtenerInputModel struct {
	IdUsuario        *int `json:"idUsuario" validate:"required"`
	IdEstatusUsuario *int `json:"idEstatusUsuario"`
}

type CursosNuevosObtenerInputModel struct {
	IdUsuario        *int `json:"idUsuario" validate:"required"`
	NumeroResultados *int `json:"numeroResultados" validate:"required"`
}

type CursosProfesorObtenerInputModel struct {
	IdProfesor *int  `json:"idProfesor" validate:"required"`
	Finalizado *bool `json:"finalizado" validate:"required"`
}

type CursoTematicaInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdTematica *int `json:"idTematica" validate:"required"`
}

type CursoTematicasObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoEstudiantesSinGrupoObtenerInputModel struct {
	IdCurso *int `json:"idCurso" validate:"required"`
}

type CursoCuestionarioRespuestaRegistrarInputModel struct {
	IdCurso    *int `json:"idCurso" validate:"required"`
	IdUsuario  *int `json:"idUsuario" validate:"required"`
	IdPregunta *int `json:"idPregunta" validate:"required"`
	Puntaje    *int `json:"puntaje" validate:"required"`
}
