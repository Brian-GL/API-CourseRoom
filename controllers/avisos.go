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

type AvisosController struct {
	DB *gorm.DB
}

func NewAvisosController(db *gorm.DB) AvisosController {
	return AvisosController{DB: db}
}

func (controller *AvisosController) AvisosObtener(c *gin.Context) {

	var modelo *models.AvisosObtenerInput

	err := json.NewDecoder(c.Request.Body).Decode(&modelo)

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

}
