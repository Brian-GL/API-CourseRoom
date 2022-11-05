package models

type ChatRegistrarInputModel struct {
	IdUsuarioEmisor   *int `json:"idUsuarioEmisor" binding:"required"`
	IdUsuarioReceptor *int `json:"idUsuarioReceptor" binding:"required"`
}
