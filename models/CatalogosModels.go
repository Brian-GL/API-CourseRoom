package models

type EstadosObtenerInputModel struct {
	IdEstado *int `json:"idEstado"`
}

type EstatusTareaPendienteObtenerInputModel struct {
	IdEstatusTareaPendiente *int `json:"idEstatusTareaPendiente"`
}

type CursoEstatusObtenerInputModel struct {
	IdEstatus *int `json:"idEstatus"`
}

type LocalidadesObtenerInputModel struct {
	IdEstado    *int `json:"idEstado"`
	IdLocalidad *int `json:"idLocalidad"`
}

type PreguntaRespuestaEstatusObtenerInputModel struct {
	IdEstatusPreguntaRespuesta *int `json:"idEstatusPreguntaRespuesta"`
}

type PreguntasCuestionarioObtenerInputModel struct {
	IdCuestionario *int `json:"idCuestionario"`
}

type TematicasObtenerInputModel struct {
	IdTematica *int `json:"idTematica"`
}

type TiposUsuarioObtenerInputModel struct {
	IdTipoUsuario *int `json:"idTipoUsuario"`
}

type TiposArchivoObtenerInputModel struct {
	IdTipoArchivo *int `json:"idTipoArchivo"`
}
