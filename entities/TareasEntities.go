package entities

import "time"

type TareaArchivosAdjuntosObtenerEntity struct {
	IdArchivoAdjunto   int        `json:"idArchivoAdjunto"`
	Nombre             string     `json:"nombre"`
	Archivo            string     `json:"archivo"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type TareaArchivosEntregadosObtenerEntity struct {
	IdArchivoEntregado int        `json:"idArchivoEntregado"`
	Nombre             string     `json:"nombre"`
	Archivo            string     `json:"archivo"`
	FechaRegistro      time.Time  `json:"fechaRegistro"`
	FechaActualizacion *time.Time `json:"fechaActualizacion"`
}

type TareaEstudianteDetalleObtenerEntity struct {
	IdCurso           int        `json:"idCurso"`
	Curso             string     `json:"curso"`
	ImagenCurso       *string    `json:"imagenCurso"`
	IdProfesor        int        `json:"idProfesor"`
	NombreProfesor    string     `json:"nombreProfesor"`
	ImagenProfesor    *string    `json:"imagenProfesor"`
	Tarea             string     `json:"tarea"`
	Descripcion       string     `json:"descripcion"`
	FechaRegistro     time.Time  `json:"fechaRegistro"`
	FechaEntrega      time.Time  `json:"fechaEntrega"`
	Calificacion      float64    `json:"calificacion"`
	Puntualidad       float64    `json:"puntualidad"`
	FechaCalificacion *time.Time `json:"fechaCalificacion"`
	FechaEntregada    *time.Time `json:"fechaEntregada"`
	Estatus           *string    `json:"estatus"`
}

type TareasMesObtenerEntity struct {
	IdTarea       int       `json:"idTarea"`
	Nombre        string    `json:"nombre"`
	Estatus       string    `json:"estatus"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	FechaEntrega  time.Time `json:"fechaEntrega"`
}

type TareasEstudianteObtenerEntity struct {
	IdTarea           int        `json:"idTarea"`
	Nombre            string     `json:"nombre"`
	IdCurso           int        `json:"idCurso"`
	NombreCurso       string     `json:"nombreCurso"`
	ImagenCurso       *string    `json:"imagenCurso"`
	FechaRegistro     time.Time  `json:"fechaRegistro"`
	FechaEntrega      time.Time  `json:"fechaEntrega"`
	FechaEntregada    *time.Time `json:"fechaEntregada"`
	Puntualidad       float32    `json:"puntualidad"`
	Calificacion      float32    `json:"calificacion"`
	FechaCalificacion *time.Time `json:"fechaCalificacion"`
	Estatus           string     `json:"estatus"`
}

type TareasCreadasProfesorObtenerEntity struct {
	IdTarea       int       `json:"idTarea"`
	Nombre        string    `json:"nombre"`
	IdCurso       int       `json:"idCurso"`
	NombreCurso   string    `json:"nombreCurso"`
	ImagenCurso   *string   `json:"imagenCurso"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	FechaEntrega  time.Time `json:"fechaEntrega"`
	Estatus       string    `json:"estatus"`
}

type TareaProfesorDetalleObtenerEntity struct {
	IdCurso       int       `json:"idCurso"`
	Curso         string    `json:"curso"`
	ImagenCurso   *string   `json:"imagenCurso"`
	Nombre        string    `json:"nombre"`
	Descripcion   string    `json:"descripcion"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	FechaEntrega  time.Time `json:"fechaEntrega"`
	Estatus       *string   `json:"estatus"`
}

type TareaImagenesEntregadasObtenerEntity struct {
	Nombre    string  `json:"nombre"`
	Archivo   string  `json:"archivo"`
	Extension *string `json:"extension"`
	Estatus   string  `json:"estatus"`
	IdUsuario int     `json:"idUsuario"`
}

type TareaReatroalimentacionesObtenerEntity struct {
	IdRetroalimentacion int       `json:"idRetroalimentacion"`
	Nombre              string    `json:"nombre"`
	Retroalimentacion   string    `json:"retroalimentacion"`
	FechaRegistro       time.Time `json:"fechaRegistro"`
}

type TareaRetroalimentacionDetalleObtenerEntity struct {
	Nombre            int       `json:"nombre"`
	Retroalimentacion int       `json:"retroalimentacion"`
	NombreArchivo     *string   `json:"nombreArchivo"`
	Archivo           *string   `json:"archivo"`
	FechaRegistro     time.Time `json:"fechaRegistro"`
}

type TareasCalificarObtenerEntity struct {
	IdTarea          int        `json:"idTarea"`
	Nombre           string     `json:"nombre"`
	IdCurso          int        `json:"idCurso"`
	NombreCurso      string     `json:"nombreCurso"`
	ImagenCurso      *string    `json:"imagenCurso"`
	IdUsuario        int        `json:"idUsuario"`
	NombreEstudiante string     `json:"nombreEstudiante"`
	ImagenEstudiante *string    `json:"imagenEstudiante"`
	FechaRegistro    time.Time  `json:"fechaRegistro"`
	FechaEntrega     time.Time  `json:"fechaEntrega"`
	FechaEntregada   *time.Time `json:"fechaEntregada"`
	Estatus          string     `json:"estatus"`
}

type TareaCalificarActualizarEntity struct {
	Codigo      int      `json:"codigo"`
	Puntualidad *float32 `json:"puntualidad"`
	NombreTarea *string  `json:"nombreTarea"`
	Mensaje     string   `json:"mensaje"`
}
