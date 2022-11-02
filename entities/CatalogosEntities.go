package entities

type EstadosObtenerEntity struct {
	IdEstado int    `json:"idEstado"`
	Estado   string `json:"estado"`
}

type EstatusTareasPendienteObtenerEntity struct {
	IdEstatus int    `json:"idEstatus"`
	Estatus   string `json:"estatus"`
}
