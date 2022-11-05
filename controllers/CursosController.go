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

type CursoController struct {
	DB *gorm.DB
}

func NewCursoController(db *gorm.DB) CursoController {
	return CursoController{DB: db}
}

func (controller *CursoController) CursoRemover(c *gin.Context) {

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

		var modelo *models.CursoRemoverInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.CursoRemoverDeleteAsync(controller.DB, modelo)
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

func (controller *CursoController) CursoRegistrar(c *gin.Context) {

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

		var modelo *models.CursoRegistrarInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.CursoRegistrarPostAsync(controller.DB, modelo)
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

func (controller *CursoController) CursoGruposObtener(c *gin.Context) {

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

		var modelo *models.CursoGruposObtenerInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.CursoGruposObtenerGetAsync(controller.DB, modelo)
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
