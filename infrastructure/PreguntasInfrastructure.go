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

		exec := "EXEC dbo.PreguntasRespuesta_Actualizar @IdUsuario = ?, @IdPreguntaRespuesta = ?, @IdPregunta = ?, @Descripcion = ?"

		db.Raw(exec, model.IdPregunta, model.IdUsuario, model.IdPreguntaRespuesta, model.Descripcion).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió realizar la acción"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func PreguntasRespuestaRegistarPostAsync(db *gorm.DB, model *models.PreguntasRespuestaRegistarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.PreguntasRespuesta_Registar @IdUsuario = ?, @IdPregunta = ?, @Descripcion = ?"

		db.Raw(exec, model.IdUsuario, model.IdPregunta, model.Descripcion).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió realizar la acción"}
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

		exec := "EXEC dbo.PreguntasRespuesta_Remover @IdUsuario = ?, @IdPreguntaRespuesta = ?"

		db.Raw(exec, model.IdUsuario, model.IdPreguntaRespuesta).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado}
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

		var resultado []entities.PreguntasRespuestaDetalleObtenerEntity

		exec := "EXEC dbo.PreguntasRespuestaDetalle_Obtener @IdPreguntaRespuesta = ?"

		db.Raw(exec, model.IdPreguntaRespuesta).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron grupos de curso"}
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

		exec := "EXEC dbo.PreguntasRespuestaEstatus_Actualizar @IdUsuario = ?, @IdPreguntaRespuesta = ?, @IdEstatusPregunta = ?"

		db.Raw(exec, model.IdPregunta, model.IdPreguntaRespuesta, model.IdEstatusPregunta).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió realizar la acción"}
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

		exec := "EXEC dbo.PreguntasRespuestaMensaje_Registrar @IdPreguntaRespuesta = ?, @IdUsuarioEmisor = ?, @Mensaje = ?, @Archivo = ?"

		db.Raw(exec, model.IdPreguntaRespuesta, model.IdUsuarioEmisor, model.Mensaje, model.Archivo).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió realizar la acción"}
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

		exec := "EXEC dbo.PreguntasRespuestaMensaje_Remover @IdPreguntaRespuesta = ?, @IdUsuarioEmisor = ?, @IdMensaje = ?"

		db.Raw(exec, model.IdPreguntaRespuesta, model.IdUsuarioEmisor, model.IdMensaje).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado}
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

		var resultado []entities.PreguntasRespuestaMensajesObtenerEntity

		exec := "EXEC dbo.PreguntasRespuestaMensajes_Obtener @IdPreguntaRespuesta = ?, @UltimoMensaje = ?"

		db.Raw(exec, model.IdPreguntaRespuesta, model.UltimoMensaje).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron grupos de curso"}
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
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros"}
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
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron grupos de curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}
