package main

import (
	"api-courseroom/controllers"
	"api-courseroom/database"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	db := database.GetDatabase()

	// #region Controllers

	catalogoController := controllers.NewCatalogoController(db)

	avisosController := controllers.NewAvisosController(db)

	// #endregion

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// #region Swagger Documentacion

	v1 := router.Group("/api/v1")
	{
		avisos := v1.Group("/avisos")
		{
			avisos.POST("/api/avisos/obtener", avisosController.AvisosObtener)
			avisos.PUT("/api/avisos/actualizar", avisosController.AvisoActualizar)
			avisos.POST("/api/avisos/detalle", avisosController.AvisoDetalleObtener)
			avisos.POST("/api/avisos/registrar", avisosController.AvisoRegistrar)
			avisos.DELETE("/api/avisos/remover", avisosController.AvisoRemover)
			avisos.POST("/api/avisos/validar", avisosController.AvisosValidar)
		}

		catalogos := v1.Group("/catalogos")
		{
			catalogos.POST("/api/catalogos/estados", catalogoController.Estados)
			catalogos.POST("/api/catalogos/estatustareapendiente", catalogoController.EstatusTareasPendientes)
		}

	}

	// #endregion

	// #region Avisos Endpoints

	router.POST("/api/avisos/obtener", avisosController.AvisosObtener)
	router.PUT("/api/avisos/actualizar", avisosController.AvisoActualizar)
	router.POST("/api/avisos/detalle", avisosController.AvisoDetalleObtener)
	router.POST("/api/avisos/registrar", avisosController.AvisoRegistrar)
	router.DELETE("/api/avisos/remover", avisosController.AvisoRemover)
	router.POST("/api/avisos/validar", avisosController.AvisosValidar)

	// #endregion

	// #region Catalogos Endpoints

	router.POST("/api/catalogos/estados", catalogoController.Estados)
	router.POST("/api/catalogos/estatustareapendiente", catalogoController.EstatusTareasPendientes)

	// #endregion

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("Running API...")
	router.Run(":1313")

}
