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

type CatalogoController struct {
	DB *gorm.DB
}

func NewCatalogoController(db *gorm.DB) CatalogoController {
	return CatalogoController{DB: db}
}

func (controller *CatalogoController) Estados(c *gin.Context) {

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

		var modelo *models.EstadosObtenerInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.EstadosGetAsync(controller.DB, modelo)
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

func (controller *CatalogoController) EstatusTareasPendientes(c *gin.Context) {

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

		var modelo *models.EstatusTareaPendienteObtenerInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.EstatusTareaPendienteGetAsync(controller.DB, modelo)
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

func (controller *CatalogoController) CursoEstatus(c *gin.Context) {

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

		var modelo *models.CursoEstatusObtenerInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.CursoEstatusObtenerGetAsync(controller.DB, modelo)
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
