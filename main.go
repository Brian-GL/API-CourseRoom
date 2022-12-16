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
	cursosController := controllers.NewCursosController(middleware)
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

	http.HandleFunc("/api/cursos/registrar", cursosController.CursoRegistrar)
	http.HandleFunc("/api/cursos/remover", cursosController.CursoRemover)
	http.HandleFunc("/api/cursos/grupos", cursosController.CursoGruposObtener)
	http.HandleFunc("/api/cursos/actualizar", cursosController.CursoActualizar)
	http.HandleFunc("/api/cursos/abandonaractualizar", cursosController.CursoAbandonarActualizar)
	http.HandleFunc("/api/cursos/cursocuestionariocontestar", cursosController.CursoCuestionarioContestar)
	http.HandleFunc("/api/cursos/cursocuestionarioabandonaractualizar", cursosController.CursoCuestionarioAbandonarActualizar)
	http.HandleFunc("/api/cursos/cursodesempenoobtener", cursosController.CursoDesempenoObtener)
	http.HandleFunc("/api/cursos/cursoestudianteregistrar", cursosController.CursoEstudianteRegistrar)
	http.HandleFunc("/api/cursos/cursoestudiantedetalleobtener", cursosController.CursoEstudianteDetalleObtener)
	http.HandleFunc("/api/cursos/cursofinalizaractualizar", cursosController.CursoFinalizarActualizar)
	http.HandleFunc("/api/cursos/cursomaterialregistrar", cursosController.CursoMaterialRegistrar)
	http.HandleFunc("/api/cursos/cursomaterialremover", cursosController.CursoMaterialRemover)
	http.HandleFunc("/api/cursos/cursomensajeregistrar", cursosController.CursoMensajeRegistrar)
	http.HandleFunc("/api/cursos/cursomensajeremover", cursosController.CursoMensajeRemover)
	http.HandleFunc("/api/cursos/cursomensajesobtener", cursosController.CursoMensajesObtener)
	http.HandleFunc("/api/cursos/cursoestudianteobtener", cursosController.CursoEstudianteObtener)
	http.HandleFunc("/api/cursos/cursoprofesordetalleobtener", cursosController.CursoProfesorDetalleObtener)
	http.HandleFunc("/api/cursos/cursoprofesortareasobtener", cursosController.CursoProfesorTareasObtener)
	http.HandleFunc("/api/cursos/cursopromedioobtener", cursosController.CursoPromedioObtener)
	http.HandleFunc("/api/cursos/cursobuscarobtener", cursosController.CursoBuscarObtener)
	http.HandleFunc("/api/cursos/cursoobtener", cursosController.CursoObtener)
	http.HandleFunc("/api/cursos/cursonuevoobtener", cursosController.CursoNuevoObtener)
	http.HandleFunc("/api/cursos/cursoprofesorobtener", cursosController.CursoProfesorObtener)
	http.HandleFunc("/api/cursos/cursotareasestudianteobtener", cursosController.CursoTareasEstudianteObtener)
	http.HandleFunc("/api/cursos/cursotematicaregistrar", cursosController.CursoTematicaRegistrar)
	http.HandleFunc("/api/cursos/cursotematicaremover", cursosController.CursoTematicaRemover)
	http.HandleFunc("/api/cursos/cursotematicaobtener", cursosController.CursoTematicaObtener)
	http.HandleFunc("/api/cursos/cursoestudiantedesempenoobtener", cursosController.CursoEstudianteDesempenoObtener)
	http.HandleFunc("/api/cursos/cursoestudiantessingrupoobtener", cursosController.CursoEstudiantesSinGrupoObtener)
	http.HandleFunc("/api/cursos/cursoestudiantefinalizaractualizar", cursosController.CursoEstudianteFinalizarActualizar)
	http.HandleFunc("/api/cursos/cursocuestionariorespuestaregistrar", cursosController.CursoCuestionarioRespuestaRegistrar)

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
	http.HandleFunc("/api/grupos/gruporegistrar", gruposController.GrupoRegistrar)
	http.HandleFunc("/api/grupos/gruporemover", gruposController.GrupoRemover)
	http.HandleFunc("/api/grupos/grupoabandonaractualizar", gruposController.GrupoAbandonarActualizar)
	http.HandleFunc("/api/grupos/grupoarchivocompartidoremover", gruposController.GrupoArchivoCompartidoRemover)
	http.HandleFunc("/api/grupos/grupodetalleobtener", gruposController.GrupoDetalleObtener)
	http.HandleFunc("/api/grupos/grupomensajeregistrar", gruposController.GrupoMensajeRegistrar)
	http.HandleFunc("/api/grupos/grupomensajeremover", gruposController.GrupoMensajeRemover)

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
	http.HandleFunc("/api/tareas/calificar", tareasController.TareaCalificarActualizar) //PUT

	// #endregion

	// #region Usuarios Endpoints

	http.HandleFunc("/api/usuarios/actualizar", usuariosController.UsuarioActualizar)
	http.HandleFunc("/api/usuarios/registrar", usuariosController.UsuarioRegistrar) //POST
	http.HandleFunc("/api/usuarios/remover", usuariosController.UsuarioRemover)
	http.HandleFunc("/api/usuarios/acceso", usuariosController.UsuarioAccesoObtener)
	http.HandleFunc("/api/usuarios/credenciales", usuariosController.UsuarioCredencialesObtener) //POST
	http.HandleFunc("/api/usuarios/cuenta", usuariosController.UsuarioCuentaActualizar)          //PUT
	http.HandleFunc("/api/usuarios/cuentaobtener", usuariosController.UsuarioCuentaObtener)
	http.HandleFunc("/api/usuarios/desempeno", usuariosController.UsuarioDesempenoObtener)
	http.HandleFunc("/api/usuarios/desempenoregistrar", usuariosController.UsuarioDesempenoRegistrar)
	http.HandleFunc("/api/usuarios/detalle", usuariosController.UsuarioDetalleObtener)
	http.HandleFunc("/api/usuarios/nuevapuntualidad", usuariosController.UsuarioNuevaPuntualidadCursoObtener)
	http.HandleFunc("/api/usuarios/nuevapuntualidadgeneral", usuariosController.UsuarioNuevaPuntualidadGeneralObtener)
	http.HandleFunc("/api/usuarios/nuevopromedio", usuariosController.UsuarioNuevoPromedioCursoObtener)
	http.HandleFunc("/api/usuarios/nuevopromediogeneral", usuariosController.UsuarioNuevoPromedioGeneralObtener)
	http.HandleFunc("/api/usuarios/buscar", usuariosController.UsuariosBuscar)
	http.HandleFunc("/api/usuarios/sesion", usuariosController.UsuarioSesionActualizar)
	http.HandleFunc("/api/usuarios/sesionregistrar", usuariosController.UsuarioSesionRegistrar)
	http.HandleFunc("/api/usuarios/sesionvalidar", usuariosController.UsuarioSesionValidar)
	http.HandleFunc("/api/usuarios/sesiones", usuariosController.UsuarioSesionesObtener)
	http.HandleFunc("/api/usuarios/tematica", usuariosController.UsuarioTematicaRegistrar)
	http.HandleFunc("/api/usuarios/tematicaremover", usuariosController.UsuarioTematicaRemover)

	// // #endregion

	fmt.Println("\nCourseRoom API Opened At " + time.Now().Format("2006-01-02 15:04:05 Monday"))

	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		panic(err)
	}

}
