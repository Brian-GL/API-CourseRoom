package main

import (
	"api-courseroom/controllers"
	"api-courseroom/middleware"
	"fmt"
	"net/http"
	"time"
)

func main() {

	//Creación de middleware:
	middleware := middleware.NewMiddleware()

	// Controladores
	avisosController := controllers.NewAvisosController(middleware)
	catalogoController := controllers.NewCatalogoController(middleware)
	chatController := controllers.NewChatController(middleware)
	// cursoController := controllers.NewCursoController(middleware)
	gruposController := controllers.NewGruposController(middleware)
	usuariosController := controllers.NewUsuariosController(middleware)

	// #region Avisos Endpoints

	http.HandleFunc("/api/avisos/actualizar", avisosController.AvisoActualizar)
	http.HandleFunc("/api/avisos/registrar", avisosController.AvisoRegistrar)
	http.HandleFunc("/api/avisos/remover", avisosController.AvisoRemover)
	http.HandleFunc("/api/avisos/detalle", avisosController.AvisoDetalleObtener)
	http.HandleFunc("/api/avisos/plagioprofesor", avisosController.AvisoPlagioProfesorRegistrar)
	http.HandleFunc("/api/avisos/obtener", avisosController.AvisosObtener)
	http.HandleFunc("/api/avisos/validar", avisosController.AvisosValidar)

	// // #endregion

	// // #region Catalogos Endpoints

	http.HandleFunc("/api/catalogos/estados", catalogoController.EstadosObtener)
	http.HandleFunc("/api/catalogos/estatustareapendiente", catalogoController.EstatusTareaPendiente)
	http.HandleFunc("/api/catalogos/cursoestatus", catalogoController.CursoEstatus)
	http.HandleFunc("/api/catalogos/localidades", catalogoController.Localidades)
	http.HandleFunc("/api/catalogos/preguntarespuesta", catalogoController.PreguntaRespuesta)
	http.HandleFunc("/api/catalogos/preguntascuestionario", catalogoController.PreguntasCuestionario)
	http.HandleFunc("/api/catalogos/tematicas", catalogoController.Tematicas)
	http.HandleFunc("/api/catalogos/tiposusuario", catalogoController.TiposUsuario)

	// // #endregion

	// // #region Chats Endpoints

	http.HandleFunc("/api/chats/registrar", chatController.ChatRegistrar)
	http.HandleFunc("/api/chats/remover", chatController.ChatRemover)
	http.HandleFunc("/api/chats/mensajeregistrar", chatController.ChatMensajeRegistrar)
	http.HandleFunc("/api/chats/mensajeremover", chatController.ChatMensajeRemover)
	http.HandleFunc("/api/chats/mensajesobtener", chatController.ChatMensajesObtener)
	http.HandleFunc("/api/chats/buscar", chatController.ChatsBuscar)
	http.HandleFunc("/api/chats/obtener", chatController.ChatsObtener)

	// // #endregion

	// // #region Cursos Endpoints

	// http.HandleFunc("/api/cursos/registrar", cursoController.CursoRegistrar)
	// http.HandleFunc("/api/cursos/remover", cursoController.CursoRemover)
	// http.HandleFunc("/api/cursos/grupos", cursoController.CursoGruposObtener)

	// // #endregion

	// #region Grupos Endpoints

	http.HandleFunc("/api/grupos/actualizar", gruposController.GrupoActualizar)
	http.HandleFunc("/api/grupos/archivoscompartidos", gruposController.GrupoArchivosCompartidosObtener)
	http.HandleFunc("/api/grupos/archivocompartido", gruposController.GrupoArchivoCompartidoRegistrar)

	// #endregion

	// #region Usuarios Endpoints

	http.HandleFunc("/api/usuarios/registrar", usuariosController.UsuarioRegistrar)
	// http.HandleFunc("/api/avisos/detalle", avisosController.AvisoDetalleObtener)
	// http.HandleFunc("/api/avisos/plagioprofesor", avisosController.AvisoPlagioProfesorRegistrar)
	// http.HandleFunc("/api/avisos/obtener", avisosController.AvisosObtener)
	// http.HandleFunc("/api/avisos/validar", avisosController.AvisosValidar)

	// // #endregion

	fmt.Println("\nCourseRoom API Opened At " + time.Now().Format("2006-01-02 15:04:05 Monday"))

	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		panic(err)
	}

}
