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
	archivoController := controllers.NewArchivoController(middleware)

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
	http.HandleFunc("/api/catalogos/tiposarchivo", catalogoController.TiposArchivo)

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
	http.HandleFunc("/api/cursos/cuestionariocontestar", cursosController.CursoCuestionarioContestar)
	http.HandleFunc("/api/cursos/cuestionarioabandonaractualizar", cursosController.CursoCuestionarioAbandonarActualizar)
	http.HandleFunc("/api/cursos/desempenoobtener", cursosController.CursoDesempenoObtener)
	http.HandleFunc("/api/cursos/estudianteregistrar", cursosController.CursoEstudianteRegistrar)
	http.HandleFunc("/api/cursos/estudiantedetalleobtener", cursosController.CursoEstudianteDetalleObtener)
	http.HandleFunc("/api/cursos/finalizaractualizar", cursosController.CursoFinalizarActualizar)
	http.HandleFunc("/api/cursos/materialregistrar", cursosController.CursoMaterialRegistrar)
	http.HandleFunc("/api/cursos/materialremover", cursosController.CursoMaterialRemover)
	http.HandleFunc("/api/cursos/mensajeregistrar", cursosController.CursoMensajeRegistrar)
	http.HandleFunc("/api/cursos/mensajeremover", cursosController.CursoMensajeRemover)
	http.HandleFunc("/api/cursos/mensajesobtener", cursosController.CursoMensajesObtener)
	http.HandleFunc("/api/cursos/estudianteobtener", cursosController.CursoEstudianteObtener)
	http.HandleFunc("/api/cursos/profesordetalleobtener", cursosController.CursoProfesorDetalleObtener)
	http.HandleFunc("/api/cursos/profesortareasobtener", cursosController.CursoProfesorTareasObtener)
	http.HandleFunc("/api/cursos/promedioobtener", cursosController.CursoPromedioObtener)
	http.HandleFunc("/api/cursos/buscarobtener", cursosController.CursoBuscarObtener)
	http.HandleFunc("/api/cursos/obtener", cursosController.CursoObtener)
	http.HandleFunc("/api/cursos/nuevoobtener", cursosController.CursoNuevoObtener)
	http.HandleFunc("/api/cursos/profesorobtener", cursosController.CursoProfesorObtener)
	http.HandleFunc("/api/cursos/tareasestudianteobtener", cursosController.CursoTareasEstudianteObtener)
	http.HandleFunc("/api/cursos/tematicaregistrar", cursosController.CursoTematicaRegistrar)
	http.HandleFunc("/api/cursos/tematicaremover", cursosController.CursoTematicaRemover)
	http.HandleFunc("/api/cursos/tematicaobtener", cursosController.CursoTematicaObtener)
	http.HandleFunc("/api/cursos/estudiantedesempenoobtener", cursosController.CursoEstudianteDesempenoObtener)
	http.HandleFunc("/api/cursos/estudiantessingrupoobtener", cursosController.CursoEstudiantesSinGrupoObtener)
	http.HandleFunc("/api/cursos/estudiantefinalizaractualizar", cursosController.CursoEstudianteFinalizarActualizar)
	http.HandleFunc("/api/cursos/cuestionariorespuestaregistrar", cursosController.CursoCuestionarioRespuestaRegistrar)

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
	http.HandleFunc("/api/grupos/registrar", gruposController.GrupoRegistrar)
	http.HandleFunc("/api/grupos/remover", gruposController.GrupoRemover)
	http.HandleFunc("/api/grupos/abandonaractualizar", gruposController.GrupoAbandonarActualizar)
	http.HandleFunc("/api/grupos/archivocompartidoremover", gruposController.GrupoArchivoCompartidoRemover)
	http.HandleFunc("/api/grupos/detalleobtener", gruposController.GrupoDetalleObtener)
	http.HandleFunc("/api/grupos/mensajeregistrar", gruposController.GrupoMensajeRegistrar)
	http.HandleFunc("/api/grupos/mensajeremover", gruposController.GrupoMensajeRemover)

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
	http.HandleFunc("/api/tareas/archivosentregados", tareasController.TareaArchivosEntregadosObtener)
	http.HandleFunc("/api/tareas/estudiante", tareasController.TareaEstudianteObtener)
	http.HandleFunc("/api/tareas/creadaprofesor", tareasController.TareaCreadaProfesorObtener)
	http.HandleFunc("/api/tareas/profesordetalle", tareasController.TareaProfesorDetalleObtener)
	http.HandleFunc("/api/tareas/retroalimentaciones", tareasController.TareaReatroalimentacionesObtener)
	http.HandleFunc("/api/tareas/calificarobtener", tareasController.TareaCalificarObtener)
	http.HandleFunc("/api/tareas/entregar", tareasController.TareaEntregarActualizar)
	http.HandleFunc("/api/tareas/archivoentregado", tareasController.TareaArchivoEntregadoRemover)
	http.HandleFunc("/api/tareas/archivoadjunto", tareasController.TareaArchivoAdjuntoRemover)
	http.HandleFunc("/api/tareas/archivoadjuntoregistrar", tareasController.TareaArchivoAdjuntoRegistrar)

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

	// #region Archivo Endpoints

	http.HandleFunc("/api/archivo/actualizar", archivoController.ArchivoActualizar)

	// endregion

	// #region preguntasrespuestas Endpoints

	http.HandleFunc("/api/preguntas/actualizar", preguntasController.PreguntasRespuestaActualizar)
	http.HandleFunc("/api/preguntas/registrar", preguntasController.PreguntasRespuestaRegistar)
	http.HandleFunc("/api/preguntas/remover", preguntasController.PreguntasRespuestaRemover)
	http.HandleFunc("/api/preguntas/detalle", preguntasController.PreguntasRespuestaDetalleObtener)
	http.HandleFunc("/api/preguntas/estatus", preguntasController.PreguntasRespuestaEstatusActualizar)
	http.HandleFunc("/api/preguntas/mensajeregistrar", preguntasController.PreguntasRespuestaMensajeRegistrar)
	http.HandleFunc("/api/preguntas/mensajeremover", preguntasController.PreguntasRespuestaMensajeRemover)
	http.HandleFunc("/api/preguntas/mensajeobtener", preguntasController.PreguntasRespuestaMensajesObtener)
	http.HandleFunc("/api/preguntas/buscar", preguntasController.PreguntasRespuestasBuscar)
	http.HandleFunc("/api/preguntas/obtener", preguntasController.PreguntasRespuestasObtener)

	// endregion

	// // #endregion

	fmt.Println("\nCourseRoom API Opened At " + time.Now().Format("2006-01-02 15:04:05 Monday"))

	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		panic(err)
	}

}
