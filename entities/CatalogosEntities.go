package entities

type EstadosObtenerEntity struct {
	IdEstado int    `json:"idEstado"`
	Estado   string `json:"estado"`
}

type EstatusTareasPendienteObtenerEntity struct {
	IdEstatus int    `json:"idEstatus"`
	Estatus   string `json:"estatus"`
}

type CursoEstatusObtenerEntity struct {
	IdEstatusCurso int    `json:"idEstatusCurso"`
	Estatus        string `json:"estatus"`
	Descripcion    string `json:"descripcion"`
}
