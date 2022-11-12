package models

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
