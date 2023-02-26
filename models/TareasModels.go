package models

import "time"

type TareaInputModel struct {
	IdTarea   *int `json:"idTarea" validate:"required"`
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type TareaArchivosAdjuntosObtenerInputModel struct {
	IdTarea *int `json:"idTarea" validate:"required"`
}

type TareasMesObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
	Mes       *int `json:"mes" validate:"required"`
}

type TareasEstudianteObtenerInputModel struct {
	IdUsuario *int `json:"idUsuario" validate:"required"`
}

type TareasCreadasProfesorObtenerInputModel struct {
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type TareaProfesorDetalleObtenerInputModel struct {
	IdTarea    *int `json:"idTarea" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type TareaRetroalimentacionDetalleObtenerInputModel struct {
	IdRetroalimentacion *int `json:"idRetroalimentacion" validate:"required"`
}

type TareasCalificarObtenerInputModel struct {
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type TareaActualizarInputModel struct {
	IdTarea     *int    `json:"idTarea" validate:"required"`
	IdProfesor  *int    `json:"idProfesor" validate:"required"`
	Nombre      *string `json:"nombre" validate:"required"`
	Descripcion *string `json:"descripcion" validate:"required"`
}

type TareaArchivoEntregadoRemoverInputModel struct {
	IdTarea            *int `json:"idTarea" validate:"required"`
	IdUsuario          *int `json:"idUsuario" validate:"required"`
	IdArchivoEntregado *int `json:"idArchivoEntregado" validate:"required"`
}

type TareaCalificarActualizarInputModel struct {
	IdTarea      *int     `json:"idTarea" validate:"required"`
	IdCurso      *int     `json:"idCurso" validate:"required"`
	IdProfesor   *int     `json:"idProfesor" validate:"required"`
	IdUsuario    *int     `json:"idUsuario" validate:"required"`
	Calificacion *float32 `json:"calificacion" validate:"required"`
}

type TareaArchivoEntregadoRegistrarInputModel struct {
	IdTarea       *int    `json:"idTarea" validate:"required"`
	IdUsuario     *int    `json:"idUsuario" validate:"required"`
	NombreArchivo *string `json:"nombreArchivo" validate:"required"`
	Archivo       *string `json:"archivo" validate:"required"`
}

type TareaArchivoAdjuntoRemoverInputModel struct {
	IdTarea          *int `json:"idTarea" validate:"required"`
	IdProfesor       *int `json:"idProfesor" validate:"required"`
	IdArchivoAdjunto *int `json:"idArchivoAdjunto" validate:"required"`
}

type TareaArchivoAdjuntoRegistrarInputModel struct {
	IdTarea       *int    `json:"idTarea" validate:"required"`
	IdProfesor    *int    `json:"idProfesor" validate:"required"`
	NombreArchivo *string `json:"nombreArchivo" validate:"required"`
	Archivo       *string `json:"Archivo" validate:"required"`
}

type TareaRemoverInputModel struct {
	IdTarea    *int `json:"idTarea" validate:"required"`
	IdProfesor *int `json:"idProfesor" validate:"required"`
}

type TareaRegistrarInputModel struct {
	IdCurso      *int       `json:"idCurso" validate:"required"`
	IdProfesor   *int       `json:"idProfesor" validate:"required"`
	Nombre       *string    `json:"nombre" validate:"required"`
	Descripcion  *string    `json:"descripcion" validate:"required"`
	FechaEntrega *time.Time `json:"fechaEntrega" validate:"required"`
}

type TareaRetroalimentacionRegistrarInputModel struct {
	IdTarea           *int    `json:"idTarea" validate:"required"`
	IdProfesor        *int    `json:"idProfesor" validate:"required"`
	IdUsuario         *int    `json:"idUsuario" validate:"required"`
	Nombre            *string `json:"nombre" validate:"required"`
	Retroalimentacion *string `json:"retroalimentacion" validate:"required"`
	NombreArchivo     *string `json:"nombreArchivo"`
	Archivo           *string `json:"archivo"`
}
