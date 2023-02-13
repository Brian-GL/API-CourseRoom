package entities

import "time"

type CursoDesempenoObtenerEntity struct {
	IdUsuario                   string   `json:"idUsuario"`
	NombreCompleto              int      `json:"nombreCompleto"`
	Imagen                      *string  `json:"imagen"`
	IdTarea                     int      `json:"idTarea"`
	Tarea                       string   `json:"tarea"`
	Calificacion                int      `json:"calificacion"`
	PromedioCalificacionCurso   float32  `json:"promedioCalificacionCurso"`
	PrediccionCalificacionCurso *float32 `json:"prediccionCalificacionCurso"`
	Puntualidad                 float32  `json:"puntualidad"`
	PromedioPuntualidadCurso    float32  `json:"promedioPuntualidadCurso"`
	FechaRegistro               string   `json:"fechaRegistro"`
}

type CursoEstudianteDetalleObtenerEntity struct {
	Nombre                  string     `json:"nombre"`
	Descripcion             string     `json:"descripcion"`
	Imagen                  *string    `json:"imagen"`
	IdProfesor              int        `json:"idProfesor"`
	NombreProfesor          string     `json:"nombreProfesor"`
	ImagenProfesor          *string    `json:"imagenProfesor"`
	FechaRegistroCurso      time.Time  `json:"fechaRegistroCurso"`
	FechaActualizacionCurso *time.Time `json:"fechaActualizacionCurso"`
	Finalizado              bool       `json:"finalizado"`
	FechaRegistro           time.Time  `json:"fechaRegistro"`
	FechaActualizacion      *time.Time `json:"fechaActualizacion"`
	Estatus                 string     `json:"estatus"`
	DescripcionEstatus      *string    `json:"descripcionEstatus"`
}

type CursoGruposObtenerEntity struct {
	IdGrupo           int     `json:"idGrupo"`
	Nombre            string  `json:"nombre"`
	Imagen            *string `json:"imagen"`
	NumeroIntegrantes *int    `json:"numeroIntegrantes"`
}

type CursoMaterialesObtenerEntity struct {
	IdMaterialSubido    int        `json:"idMaterialsubido"`
	Nombre              string     `json:"nombre"`
	Archivo             string     `json:"archivo"`
	FechaActualizacion  *time.Time `json:"fechaActualizacion"`
	IdUsuarioemisor     int        `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string     `json:"nombreUsuarioEmisor"`
	FechaRegistro       time.Time  `json:"fechaRegistro"`
}

type CursoMensajesObtenerEntity struct {
	IdMensaje           int       `json:"idMensaje"`
	Mensaje             string    `json:"mensaje"`
	Archivo             *string   `json:"archivo"`
	IdUsuarioemisor     int       `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string    `json:"nombreUsuarioEmisor"`
	FechaRegistro       time.Time `json:"fechaRegistro"`
}

type CursoEstudiantesObtenerEntity struct {
	IdUsuario          int        `json:"idUsuario"`
	NombreCompleto     string     `json:"nombreCompleto"`
	Imagen             *string    `json:"imagen"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
	Estatus            string     `json:"estatus"`
}

type CursoProfesorDetalleObtenerEntity struct {
	Nombre             string     `json:"nombre"`
	Descripcion        string     `json:"descripcion"`
	Imagen             *string    `json:"imagen"`
	IdProfesor         int        `json:"idProfesor"`
	NombreProfesor     string     `json:"nombreProfesor"`
	ImagenProfesor     *int       `json:"imagenProfesor"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
	Puntaje            float32    `json:"puntaje"`
	Finalizado         bool       `json:"finalizado"`
}

type CursoProfesorTareasObtenerEntity struct {
	IdTarea        int        `json:"idTarea"`
	Tarea          string     `json:"tarea"`
	FechaRegistro  time.Time  `json:"fechaRegistro"`
	FechaEntrega   *time.Time `json:"fechaEntrega"`
	EstatusEntrega string     `json:"estatusEntrega"`
}

type CursoPromedioObtenerEntity struct {
	PromedioCurso *float32 `json:"promedioCurso"`
}

type CursosBuscarEntity struct {
	IdCurso        int        `json:"idCurso"`
	Curso          string     `json:"curso"`
	ImagenCurso    *string    `json:"imagenCurso"`
	IdProfesor     int        `json:"idProfesor"`
	Profesor       string     `json:"profesor"`
	ImagenProfesor *string    `json:"imagenProfesor"`
	ListaTematicas *string    `json:"listaTematicas"`
	FechaRegistro  time.Time  `json:"fechaRegistro"`
	Puntaje        float32    `json:"puntaje"`
	FechaIngreso   *time.Time `json:"fechaIngreso"`
	Estatus        *string    `json:"estatus"`
}

type CursosObtenerEntity struct {
	IdCurso        int       `json:"idCurso"`
	Curso          string    `json:"curso"`
	ImagenCurso    *string   `json:"imagenCurso"`
	IdProfesor     int       `json:"idProfesor"`
	Profesor       string    `json:"profesor"`
	ImagenProfesor *string   `json:"imagenProfesor"`
	ListaTematicas *string   `json:"listaTematicas"`
	FechaRegistro  time.Time `json:"fechaRegistro"`
	Puntaje        float32   `json:"puntaje"`
	FechaIngreso   time.Time `json:"fechaIngreso"`
	Estatus        string    `json:"estatus"`
}

type CursosNuevosObtenerEntity struct {
	IdCurso        int       `json:"idCurso"`
	Curso          string    `json:"curso"`
	ImagenCurso    *string   `json:"imagenCurso"`
	IdProfesor     int       `json:"idProfesor"`
	Profesor       string    `json:"profesor"`
	ImagenProfesor *string   `json:"imagenProfesor"`
	ListaTematicas *string   `json:"listaTematicas"`
	FechaRegistro  time.Time `json:"fechaRegistro"`
	Puntaje        float32   `json:"puntaje"`
}

type CursosProfesorObtenerEntity struct {
	IdCurso        int       `json:"idCurso"`
	Curso          string    `json:"curso"`
	Imagen         *string   `json:"imagen"`
	ListaTematicas *string   `json:"listaTematicas"`
	Estatus        string    `json:"estatus"`
	FechaRegistro  time.Time `json:"fechaRegistro"`
	Puntaje        float32   `json:"puntaje"`
}

type CursoTareasEstudianteObtenerEntity struct {
	IdTarea           int        `json:"idTarea"`
	Nombre            string     `json:"nombre"`
	FechaRegistro     time.Time  `json:"fechaRegistro"`
	FechaEntrega      time.Time  `json:"fechaEntrega"`
	FechaEntregada    *time.Time `json:"fechaEntregada"`
	FechaCalificacion *time.Time `json:"fechaCalificacion"`
	Calificacion      *float32   `json:"calificacion"`
	Puntualidad       *float32   `json:"puntualidad"`
	Estatus           string     `json:"estatus"`
}

type CursoTematicasObtenerEntity struct {
	IdTematica int    `json:"idTematica"`
	Tematica   string `json:"tematica"`
}

type CursoEstudianteDesempenoObtenerEntity struct {
	IdTarea                     int       `json:"idTarea"`
	Tarea                       string    `json:"tarea"`
	Calificacion                float32   `json:"calificacion"`
	PromedioCalificacionCurso   float32   `json:"promedioCalificacionCurso"`
	PrediccionCalificacionCurso *float32  `json:"prediccionCalificacionCurso"`
	Puntualidad                 float32   `json:"puntualidad"`
	PromedioPuntualidadCurso    float32   `json:"promedioPuntualidadCurso"`
	FechaRegistro               time.Time `json:"fechaRegistro"`
}

type CursoEstudiantesSinGrupoObtenerEntity struct {
	IdUsuario    int       `json:"idUsuario"`
	Estudiante   string    `json:"estudiante"`
	FechaIngreso time.Time `json:"fechaIngreso"`
}
