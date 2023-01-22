package entities

type CursoEstatusObtenerEntity struct {
	IdEstatus   int    `json:"idEstatus"`
	Estatus     string `json:"estatus"`
	Descripcion string `json:"descripcion"`
}

type EstadosObtenerEntity struct {
	IdEstado int    `json:"idEstado"`
	Estado   string `json:"estado"`
}

type EstatusTareaPendienteObtenerEntity struct {
	IdEstatus int    `json:"idEstatus"`
	Estatus   string `json:"estatus"`
}

type LocalidadesObtenerEntity struct {
	IdLocalidad int    `json:"idLocalidad"`
	Localidad   string `json:"localidad"`
	Estado      string `json:"estado"`
}

type PreguntaRespuestaEstatusObtenerEntity struct {
	IdEstatus int    `json:"idEstatus"`
	Estatus   string `json:"estatus"`
}

type PreguntasCuestionarioObtenerEntity struct {
	IdPregunta int    `json:"idPregunta"`
	Pregunta   string `json:"pregunta"`
}

type TematicasObtenerEntity struct {
	IdTematica int    `json:"idTematica"`
	Tematica   string `json:"tematica"`
}

type TiposUsuarioObtenerEntity struct {
	IdTipoUsuario int    `json:"idTipoUsuario"`
	TipoUsuario   string `json:"tipoUsuario"`
}

type TiposArchivoObtenerEntity struct {
	IdTipoArchivo int    `json:"idTipoArchivo"`
	TipoArchivo   string `json:"tipoArchivo"`
}
