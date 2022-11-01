package main

import (
	"api-courseroom/controllers"
	"api-courseroom/database"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.GetDatabase()

	catalogoController := controllers.NewCatalogoController(db)

	router := gin.Default()

	router.POST("/api/catalogos/estados", go catalogoController.Estados)

	router.Run(":1313")

}
