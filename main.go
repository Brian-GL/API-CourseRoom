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

	avisosController := controllers.NewAvisosController(db)

	catalogoController := controllers.NewCatalogoController(db)

	chatController := controllers.NewChatController(db)

	cursoController := controllers.NewCursoController(db)

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
			avisos.PUT("/api/avisos/actualizar", avisosController.AvisoActualizar)
			avisos.POST("/api/avisos/registrar", avisosController.AvisoRegistrar)
			avisos.DELETE("/api/avisos/remover", avisosController.AvisoRemover)
			avisos.POST("/api/avisos/detalle", avisosController.AvisoDetalleObtener)
			avisos.POST("/api/avisos/plagioprofesor", avisosController.AvisoPlagioProfesorRegistrar)
			avisos.POST("/api/avisos/obtener", avisosController.AvisosObtener)
			avisos.POST("/api/avisos/validar", avisosController.AvisosValidar)
		}

		catalogos := v1.Group("/catalogos")
		{
			catalogos.POST("/api/catalogos/estados", catalogoController.Estados)
			catalogos.POST("/api/catalogos/estatustareapendiente", catalogoController.EstatusTareasPendientes)
			catalogos.POST("/api/catalogos/cursoestatus", catalogoController.CursoEstatus)
			catalogos.POST("/api/catalogos/localidades", catalogoController.Localidades)
			catalogos.POST("/api/catalogos/preguntarespuesta", catalogoController.PreguntaRespuesta)
			catalogos.POST("/api/catalogos/preguntascuestionario", catalogoController.PreguntasCuestionario)
			catalogos.POST("/api/catalogos/tematicas", catalogoController.Tematicas)
			catalogos.POST("/api/catalogos/tiposusuario", catalogoController.TiposUsuario)
		}

		chats := v1.Group("/chat")
		{
			chats.POST("/api/chats/registrar", chatController.ChatRegistrar)
			chats.POST("/api/chats/remover", chatController.ChatRemover)
			chats.POST("/api/chats/mensajeregistrar", chatController.ChatMensajeRegistrar)
			chats.POST("/api/chats/mensajeremover", chatController.ChatMensajeRemover)
			chats.POST("/api/chats/mensajesobtener", chatController.ChatMensajesObtener)
			chats.POST("/api/chats/buscar", chatController.ChatsBuscar)
			chats.POST("/api/chats/obtener", chatController.ChatsObtener)
		}

		cursos := v1.Group("/cursos")
		{
			cursos.DELETE("/api/cursos/remover", cursoController.CursoRemover)
			cursos.POST("/api/cursos/registrar", cursoController.CursoRegistrar)
			cursos.POST("/api/cursos/grupos", cursoController.CursoGruposObtener)
		}

	}

	// #endregion

	// #region Avisos Endpoints

	router.PUT("/api/avisos/actualizar", avisosController.AvisoActualizar)
	router.POST("/api/avisos/registrar", avisosController.AvisoRegistrar)
	router.DELETE("/api/avisos/remover", avisosController.AvisoRemover)
	router.POST("/api/avisos/detalle", avisosController.AvisoDetalleObtener)
	router.POST("/api/avisos/plagioprofesor", avisosController.AvisoPlagioProfesorRegistrar)
	router.POST("/api/avisos/obtener", avisosController.AvisosObtener)
	router.POST("/api/avisos/validar", avisosController.AvisosValidar)

	// #endregion

	// #region Catalogos Endpoints

	router.POST("/api/catalogos/estados", catalogoController.Estados)
	router.POST("/api/catalogos/estatustareapendiente", catalogoController.EstatusTareasPendientes)
	router.POST("/api/catalogos/cursoestatus", catalogoController.CursoEstatus)
	router.POST("/api/catalogos/localidades", catalogoController.Localidades)
	router.POST("/api/catalogos/preguntarespuesta", catalogoController.PreguntaRespuesta)
	router.POST("/api/catalogos/preguntascuestionario", catalogoController.PreguntasCuestionario)
	router.POST("/api/catalogos/tematicas", catalogoController.Tematicas)
	router.POST("/api/catalogos/tiposusuario", catalogoController.TiposUsuario)

	// #endregion

	// #region Chats Endpoints

	router.POST("/api/chats/registrar", chatController.ChatRegistrar)
	router.POST("/api/chats/remover", chatController.ChatRemover)
	router.POST("/api/chats/mensajeregistrar", chatController.ChatMensajeRegistrar)
	router.POST("/api/chats/mensajeremover", chatController.ChatMensajeRemover)
	router.POST("/api/chats/mensajesobtener", chatController.ChatMensajesObtener)
	router.POST("/api/chats/buscar", chatController.ChatsBuscar)
	router.POST("/api/chats/obtener", chatController.ChatsObtener)

	// #endregion

	// #region Cursos Endpoints

	router.POST("/api/cursos/registrar", cursoController.CursoRegistrar)
	router.DELETE("/api/cursos/remover", cursoController.CursoRemover)
	router.POST("/api/cursos/grupos", cursoController.CursoGruposObtener)

	// #endregion

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("Running API...")
	router.Run(":1313")

}
