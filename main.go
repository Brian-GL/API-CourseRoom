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
	tareasController := controllers.NewTareasController(middleware)
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
	http.HandleFunc("/api/grupos/mensajes", gruposController.GruposMensajesObtener)
	http.HandleFunc("/api/grupos/obtener", gruposController.GruposObtener)
	http.HandleFunc("/api/grupos/miembros", gruposController.GrupoMiembrosObtener)
	http.HandleFunc("/api/grupos/tareaspendientes", gruposController.GrupoTareasPendientesObtener)
	http.HandleFunc("/api/grupos/tareapendientedetalle", gruposController.GrupoTareaPendienteDetalleObtener)
	http.HandleFunc("/api/grupos/tareapendienteestatus", gruposController.GrupoTareaPendienteEstatusActualizar)
	http.HandleFunc("/api/grupos/miembro", gruposController.GrupoMiembroRemover)
	http.HandleFunc("/api/grupos/miembro", gruposController.GrupoMiembroRegistrar)
	http.HandleFunc("/api/grupos/tareapendiente", gruposController.GrupoTareaPendienteActualizar)
	http.HandleFunc("/api/grupos/tareapendiente", gruposController.GrupoTareaPendienteRegistrar)

	// #endregion

	// #region Tareas Endpoints

	http.HandleFunc("/api/tareas/archivosadjuntos", tareasController.TareaArchivosAdjuntosObtener)
	http.HandleFunc("/api/tareas/estudiantedetalle", tareasController.TareaEstudianteDetalleObtener)
	http.HandleFunc("/api/tareas/mes", tareasController.TareasMesObtener)
	http.HandleFunc("/api/tareas/imagenesentregadas", tareasController.TareaImagenesEntregadasObtener)
	http.HandleFunc("/api/tareas/retroalimentaciondetalle", tareasController.TareaRetroalimentacionDetalleObtener)
	http.HandleFunc("/api/tareas/actualizar", tareasController.TareaActualizar)
	http.HandleFunc("/api/tareas/archivoentregado", tareasController.TareaArchivoEntregadoRegistrar)
	http.HandleFunc("/api/tareas/remover", tareasController.TareaRemover)
	http.HandleFunc("/api/tareas/registrar", tareasController.TareaRemover)
	http.HandleFunc("/api/tareas/retroalimentacion", tareasController.TareaRetroalimentacionRegistrar)

	// #endregion

	// #region Usuarios Endpoints

	http.HandleFunc("/api/usuarios/registrar", usuariosController.UsuarioRegistrar)              //POST
	http.HandleFunc("/api/usuarios/credenciales", usuariosController.UsuarioCredencialesObtener) //POST
	http.HandleFunc("/api/usuarios/cuenta", usuariosController.UsuarioCuentaActualizar)          //PUT
	// // #endregion

	fmt.Println("\nCourseRoom API Opened At " + time.Now().Format("2006-01-02 15:04:05 Monday"))

	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		panic(err)
	}

}
