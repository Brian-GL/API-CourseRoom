package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func PreguntasRespuestaActualizarPutAsync(db *gorm.DB, model *models.PreguntasRespuestaActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.PreguntaRespuesta_Actualizar @IdUsuario = ?, @IdPreguntaRespuesta = ?, @Pregunta = ?, @Descripcion = ?"

		db.Raw(exec, model.IdUsuario, model.IdPreguntaRespuesta, model.Pregunta, model.Descripcion).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió actualizar las preguntas"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaRegistrarPostAsync(db *gorm.DB, model *models.PreguntasRespuestaRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.PreguntaRespuesta_Registrar @IdUsuario = ?, @Pregunta = ?, @Descripcion = ?"

		db.Raw(exec, model.IdUsuario, model.Pregunta, model.Descripcion).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar las preguntas"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaRemoverDeleteAsync(db *gorm.DB, model *models.PreguntasRespuestaRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.PreguntaRespuesta_Remover @IdUsuario = ?, @IdPreguntaRespuesta = ?"

		db.Raw(exec, model.IdUsuario, model.IdPreguntaRespuesta).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió generar la acción"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaDetalleObtenerGetAsync(db *gorm.DB, model *models.PreguntasRespuestaDetalleObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.PreguntaRespuestaDetalleObtenerEntity

		exec := "EXEC dbo.PreguntaRespuestaDetalle_Obtener @IdPreguntaRespuesta = ?"

		db.Raw(exec, model.IdPreguntaRespuesta).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información de la pregunta"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaEstatusActualizarPutAsync(db *gorm.DB, model *models.PreguntasRespuestaEstatusActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.PreguntaRespuestaEstatus_Actualizar @IdUsuario = ?, @IdPreguntaRespuesta = ?, @IdEstatusPregunta = ?"

		db.Raw(exec, model.IdUsuario, model.IdPreguntaRespuesta, model.IdEstatusPregunta).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió actualizar el estatus de la pregunta"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaMensajeRegistrarPostAsync(db *gorm.DB, model *models.PreguntasRespuestaMensajeRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.PreguntaRespuestaMensaje_Registrar @IdPreguntaRespuesta = ?, @IdUsuarioEmisor = ?, @Mensaje = ?, @Archivo = ?"

		db.Raw(exec, model.IdPreguntaRespuesta, model.IdUsuarioEmisor, model.Mensaje, model.Archivo).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar mensaje en la pregunta"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaMensajeRemoverDeleteAsync(db *gorm.DB, model *models.PreguntasRespuestaMensajeRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.PreguntaRespuestaMensaje_Remover @IdPreguntaRespuesta = ?, @IdUsuarioEmisor = ?, @IdMensaje = ?"

		db.Raw(exec, model.IdPreguntaRespuesta, model.IdUsuarioEmisor, model.IdMensaje).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió generar la acción"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaMensajesObtenerGetAsync(db *gorm.DB, model *models.PreguntasRespuestaMensajesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.PreguntaRespuestaMensajesObtenerEntity

		exec := "EXEC dbo.PreguntaRespuestaMensajes_Obtener @IdPreguntaRespuesta = ?"

		db.Raw(exec, model.IdPreguntaRespuesta).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron mensajes"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestasBuscarGetAsync(db *gorm.DB, model *models.PreguntasRespuestasBuscarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.PreguntasRespuestasBuscarEntity

		exec := "EXEC dbo.PreguntasRespuestas_Buscar @Busqueda = ?"

		db.Raw(exec, model.Busqueda).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros de preguntas"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func PreguntasRespuestasObtenerGetAsync(db *gorm.DB, model *models.PreguntasRespuestasObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.PreguntasRespuestasObtenerEntity

		exec := "EXEC dbo.PreguntasRespuestas_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron preguntas"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}
