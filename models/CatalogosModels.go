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
