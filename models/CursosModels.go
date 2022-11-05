package models

type CursoRemoverInputModel struct {
	IdCurso    *int `json:"idCurso" binding:"required"`
	IdProfesor *int `json:"idProfesor" binding:"required"`
}

type CursoRegistrarInputModel struct {
	Nombre      *string `json:"nombre" binding:"required"`
	Descripcion *string `json:"descripcion" binding:"required"`
	Imagen      *string `json:"imagen"`
	IdProfesor  *int    `json:"idProfesor" binding:"required"`
}

type CursoGruposObtenerInputModel struct {
	IdCurso *int  `json:"idCurso" binding:"required"`
	Activo  *bool `json:"activo"`
}
