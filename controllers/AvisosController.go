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

type AvisosController struct {
	DB *gorm.DB
}

func NewAvisosController(db *gorm.DB) AvisosController {
	return AvisosController{DB: db}
}

// AvisosObtener godoc
//
//	@Summary      Obtener avisos
//	@Description  Obtiene los avisos de un usuario
//	@Accept       json
//	@Produce      json
//	@Success      200  {[]entities.AvisosObtenerEntity}
//	@Failure      400  {string}
//	@Failure      401  {string}
//	@Failure      404  {string}
//	@Failure      500  {string}
//	@Security     Basic Token
//
// @Router /api/avisos/obtener [post]
func (controller *AvisosController) AvisosObtener(c *gin.Context) {

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

		var modelo *models.AvisosObtenerInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.AvisosObtenerGetAsync(controller.DB, modelo)
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
