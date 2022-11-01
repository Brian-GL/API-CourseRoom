package models

type AvisosObtenerInput struct {
	IdUsuario int     `json:"idUsuario"`
	Leido     *bool   `json:"leido"`
	Token     *string `json:"token"`
}
