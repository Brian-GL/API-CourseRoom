package controllers

import (
	"api-courseroom/async"
	"api-courseroom/infrastructure"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatController struct {
	DB *gorm.DB
}

func NewChatController(db *gorm.DB) ChatController {
	return ChatController{DB: db}
}

func (controller *ChatController) ChatRegistrar(c *gin.Context) {

	// Obtener token
	token := c.GetHeader("Authorization")

	// Validar que el token no se encuentre vacío:
	if token == "" {
		c.IndentedJSON(http.StatusUnauthorized, "El token es necesario para acceder a este recurso")
		return
	}

	// Validar que el token sea el correcto:
	validarToken := middleware.ValidateToken(&token)

	if validarToken {

		var modelo *models.ChatRegistrarInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.ChatRegistrarPostAsync(controller.DB, modelo)
			})

			response := future.Await().(models.ResponseInfrastructure)

			switch response.Status {
			case models.SUCCESS:
				{
					c.IndentedJSON(http.StatusOK, response.Data)
				}
			case models.ALERT:
				{
					c.IndentedJSON(http.StatusNotFound, response.Data)
				}
			default:
				{
					c.IndentedJSON(http.StatusInternalServerError, response.Data)
				}
			}
		} else {
			c.IndentedJSON(http.StatusBadRequest, "El parámetro de entrada no cuenta con un formato adecuado")
		}
	} else {
		c.IndentedJSON(http.StatusUnauthorized, "Token inválido")
	}

}
