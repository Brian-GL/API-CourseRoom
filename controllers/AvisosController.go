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

func (controller *AvisosController) AvisoActualizar(c *gin.Context) {

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

		var modelo *models.AvisoAccionInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.AvisoActualizarPutAsync(controller.DB, modelo)
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

func (controller *AvisosController) AvisoRegistrar(c *gin.Context) {

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

		var modelo *models.AvisoRegistrarInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.AvisoRegistrarPostAsync(controller.DB, modelo)
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

func (controller *AvisosController) AvisoRemover(c *gin.Context) {

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

		var modelo *models.AvisoAccionInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.AvisoRemoverDeleteAsync(controller.DB, modelo)
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

func (controller *AvisosController) AvisoDetalleObtener(c *gin.Context) {

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

		var modelo *models.AvisoInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.AvisoDetalleObtenerGetAsync(controller.DB, modelo)
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

func (controller *AvisosController) AvisoPlagioProfesorRegistrar(c *gin.Context) {

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

		var modelo *models.AvisoPlagioProfesorRegistrarInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.AvisoPlagioProfesorRegistrarPostAsync(controller.DB, modelo)
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

func (controller *AvisosController) AvisosValidar(c *gin.Context) {

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

		var modelo *models.AvisosValidarInputModel

		err := c.ShouldBindJSON(&modelo)

		if err == nil {

			future := async.Exec(func() interface{} {
				return infrastructure.AvisoValidarGetAsync(controller.DB, modelo)
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
