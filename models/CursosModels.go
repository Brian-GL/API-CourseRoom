package models

type CursoRemoverInputModel struct {
	IdCurso    *int `json:"idCurso" binding:"required"`
	IdProfesor *int `json:"idProfesor" binding:"required"`
}
