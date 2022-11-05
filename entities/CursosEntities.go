package entities

type CursoGruposObtenerEntity struct {
	IdGrupo           int     `json:"idGrupo"`
	Nombre            string  `json:"nombre"`
	Imagen            *string `json:"imagen"`
	NumeroIntegrantes *int    `json:"numeroIntegrantes"`
}
