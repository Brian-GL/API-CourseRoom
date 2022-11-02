package models

type AvisosObtenerInputModel struct {
	IdUsuario *int  `json:"idUsuario" binding:"required"`
	Leido     *bool `json:"leido"`
}
