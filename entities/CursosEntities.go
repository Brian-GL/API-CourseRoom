package entities

import "time"

type CursoGruposObtenerEntity struct {
	IdGrupo           int     `json:"idGrupo"`
	Nombre            string  `json:"nombre"`
	Imagen            *string `json:"imagen"`
	NumeroIntegrantes *int    `json:"numeroIntegrantes"`
}

type CursoActualizarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoRegistrarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoRemoverEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoAbandonarActualizarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoCuestionarioContestarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoEstudianteFinalizarActualizarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoCuestionarioRespuestaRegistrarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoCuestionarioContestarValidarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoDesempenoObtenerEntity struct {
	IdDesempeño                int    `json:"idDesempeño"`
	IdUsuario                  string `json:"idUsuario"`
	NombreCompleto             int    `json:"nombreCompleto"`
	Imagen                     string `json:"imagen"`
	IdTarea                    int    `json:"idTarea"`
	Tarea                      string `json:"tarea"`
	Calificacion               int    `json:"calificacion"`
	PromedioCurso              string `json:"promedioCurso"`
	PrediccionPromedioCurso    int    `json:"prediccionPromedioCurso"`
	RumboPromedioCurso         string `json:"rumboPrediccionCurso"`
	PuntualidadCurso           int    `json:"puntualidaCurso"`
	PrediccionPuntualidadCurso string `json:"prediccionPuntualidadCurso"`
	RumboPuntualidadCurso      int    `json:"rumboPuntualidadCurso"`
	FechaRegistro              string `json:"fechaRegistro"`
}

type CursoEstudianteRegistrarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoEstudianteDetalleObtenerEntity struct {
	Nombre                  string    `json:"nombre"`
	Descripcion             string    `json:"descripcion"`
	Imagen                  string    `json:"imagen"`
	IdProfesor              int       `json:"idProfesor"`
	NombreProfesor          string    `json:"nombreProfesor"`
	ImagenProfesor          string    `json:"imagenProfesor"`
	FechaRegistroCurso      time.Time `json:"fechaRegistroCurso"`
	FechaActualizacionCurso string    `json:"fechaActulizacionCurso"`
	Finalizado              bool      `json:"finalizado"`
	FechaRegistro           time.Time `json:"fechaRegistro"`
	FechaActualizacion      time.Time `json:"fechaActualizacion"`
	Estatus                 string    `json:"estatus"`
	DescripcionEstatus      string    `json:"descripcionEstatus"`
}

type CursoCursoFinalizarActualizarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoGrupoaObtenerEntity struct {
	IdGrupo           int    `json:"idGrupo"`
	Nombre            string `json:"nombre"`
	Imagen            string `json:"imagen"`
	NumeroIntegrantes int    `json:"numeroIntegrantes"`
}

type CursoMaterialRegistrarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoMaterialRemoverEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoMaterialObtenerEntity struct {
	IdMaterialSubido    int    `json:"idMaterialsubido"`
	Nombre              string `json:"nombre"`
	Archivo             string `json:"archivo"`
	FechaActualizacion  string `json:"fechaActualizacion"`
	IdUsuarioemisor     int    `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string `json:"nombreUsuarioEmisor"`
	FechaRegistro       int    `json:"fechaRegistro"`
}

type CursoMensajeRegistrarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoMensajeRemoverEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoMensajeObtenerEntity struct {
	IdMensaje           int    `json:"idmensaje"`
	Mensaje             string `json:"mensaje"`
	Archivo             string `json:"archivo"`
	IdUsuarioemisor     int    `json:"idUsuarioEmisor"`
	NombreUsuarioEmisor string `json:"nombreUsuarioEmisor"`
	FechaRegistro       int    `json:"fechaRegistro"`
}

type CursoEstudianteRemoverEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoEstudianteObtenerEntity struct {
	IdUsuario          int    `json:"idUsuario"`
	NombreCompleto     string `json:"nombreCompleto"`
	Imagen             string `json:"imagen"`
	FechaRegistro      string `json:"fechaRegistro"`
	FechaActualizacion string `json:"fechaActializacion"`
	Estatus            string `json:"estatus"`
}

type CursoProfesorDetalleObtenerEntity struct {
	Nombre             string  `json:"nombre"`
	Descripcion        string  `json:"descripcion"`
	Imagen             string  `json:"imagen"`
	IdProfesor         int     `json:"idProfesor"`
	NombreProfesor     string  `json:"nombreProfesor"`
	ImagenProfesor     int     `json:"imagenProfesor"`
	FechaRegistro      int     `json:"fechaRegistro"`
	FechaActualizacion int     `json:"fechaActualizacion"`
	Puntaje            float32 `json:"puntaje"`
	Finalizado         bool    `json:"finalizado"`
}

type CursoProfesorTareaObtenerEntity struct {
	IdTarea        int    `json:"idtarea"`
	Tarea          string `json:"tarea"`
	FechaRegistro  string `json:"fechaRegistro"`
	FechaEntrega   string `json:"fechaentrega"`
	EstatusEntrega string `json:"estatusEntrega"`
}

type CursoPromedioObtenerEntity struct {
	PromedioCurso float32 `json:"PromedioCurso"`
}

type CursoBuscarEntity struct {
	IdCurso        int       `json:"idCurso"`
	Curso          string    `json:"curso"`
	ImagenCurso    string    `json:"imagenCurso"`
	IdProfesor     int       `json:"idProfesor"`
	Profesor       string    `json:"profesor"`
	ImagenProfesor string    `json:"imagenProfesor"`
	ListaTematica  string    `json:"listaTematica"`
	FechaRegistro  time.Time `json:"fechaRegistro"`
	Puntaje        float32   `json:"puntaje"`
	FechaIngreso   time.Time `json:"fechaIngreso"`
	Estatus        string    `json:"estatus"`
}

type CursoObtenerEntity struct {
	IdCurso        int       `json:"idCurso"`
	Curso          string    `json:"curso"`
	ImagenCurso    string    `json:"imagenCurso"`
	IdProfesor     int       `json:"idProfesor"`
	Profesor       string    `json:"profesor"`
	ImagenProfesor string    `json:"imagenProfesor"`
	ListaTematica  string    `json:"listaTematica"`
	FechaRegistro  time.Time `json:"fechaRegistro"`
	Puntaje        float32   `json:"puntaje"`
	FechaIngreso   time.Time `json:"fechaIngreso"`
	Estatus        string    `json:"estatus"`
}

type CursoNuevosObtenerEntity struct {
	IdCurso        int       `json:"idCurso"`
	Curso          string    `json:"curso"`
	ImagenCurso    string    `json:"imagenCurso"`
	IdProfesor     int       `json:"idProfesor"`
	Profesor       string    `json:"profesor"`
	ImagenProfesor string    `json:"imagenProfesor"`
	ListaTematica  string    `json:"listaTematica"`
	FechaRegistro  time.Time `json:"fechaRegistro"`
	Puntaje        float32   `json:"puntaje"`
}

type CursoProfesorObtenerEntity struct {
	IdCurso        int       `json:"idCurso"`
	Curso          string    `json:"curso"`
	Imagen         string    `json:"imagen"`
	ListaTematicas string    `json:"listaTematica"`
	Estatus        string    `json:"estatus"`
	FechaRegistro  time.Time `json:"fechaRegistro"`
	Puntaje        float32   `json:"puntaje"`
}

type CursoTareaEstudianteObtenerEntity struct {
	IdTarea           int       `json:"idTarea"`
	Nombre            string    `json:"nombre"`
	FechaRegistro     string    `json:"fechaRegistro"`
	FechaEntrega      string    `json:"fechaEntrega"`
	FechaEntregada    string    `json:"fechaEntregada"`
	FechaCalificacion time.Time `json:"fechaCalificacion"`
	Calificacion      float32   `json:"calificacion"`
	Puntualidad       float32   `json:"puntualidad"`
	Estatus           string    `json:"estatus"`
}

type CursoTematicaRegistrarEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoTematicaRemoverEntity struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

type CursoTematicasObtenerEntity struct {
	IdTematica int    `json:"idTematica"`
	Tematica   string `json:"tematica"`
}

type CursoEstudianteDesempenoObtenerEntity struct {
	IdDesempeño                int     `json:"idDesempeño"`
	IdTarea                    int     `json:"idTarea"`
	Tarea                      string  `json:"tarea"`
	Calificacion               float32 `json:"calificacion"`
	PromedioCurso              float32 `json:"promediCurso"`
	PrediccionPromedioCurso    float32 `json:"prediccionPromedioCurso"`
	PuntualidadCurso           float32 `json:"puntualidadCurso"`
	PrediccionPuntualidadCurso float32 `json:"prediccionPuntualidadcurso"`
	RumboPuntualidadCurso      string  `json:"rumboPuntualidadCurso"`
	FechaRegistro              string  `json:"fechaRegistro"`
}

type CursoEstudiantesSinGrupoObtenerEntity struct {
	IdUsuario     int    `json:"idUsuario"`
	Estudiante    string `json:"estudiante"`
	FechaRegistro string `json:"fechaRegistro"`
}
