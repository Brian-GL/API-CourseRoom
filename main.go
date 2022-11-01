package main

import (
	"api-courseroom/controllers"
	"api-courseroom/database"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.GetDatabase()

	// #region Controllers

	catalogoController := controllers.NewCatalogoController(db)

	avisosController := controllers.NewAvisosController(db)

	// #endregion

	router := gin.Default()

	router.POST("/api/catalogos/estados", catalogoController.Estados)
	router.POST("/api/catalogos/estatustareaspendiente", catalogoController.EstatusTareasPendientes)

	router.POST("/api/avisos/obtener", avisosController.AvisosObtener)

	router.Run(":1313")

}
