package controllers

import (
	"api-courseroom/async"
	"api-courseroom/infrastructure"
	"api-courseroom/models"
	"encoding/json"
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

	var modelo *models.EstadoInputModel

	err := json.NewDecoder(c.Request.Body).Decode(&modelo)

	if err == nil {

		future := async.Exec(func() interface{} {
			return infrastructure.EstadosGetAsync(controller.DB, modelo.IdEstado)
		})

		response := future.Await().(models.ResposeInfrastruture)

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

}
