package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func GrupoActualizarPutAsync(db *gorm.DB, model *models.GrupoActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Grupo_Actualizar @IdGrupo = ?, @IdCurso = ?, @Nombre = ?, @Descripcion = ?, @Imagen = ?"

		db.Raw(exec, model.IdGrupo, model.IdCurso, model.Nombre, model.Descripcion, model.Imagen).Scan(&resultado)

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

func GrupoArchivosCompartidosObtenerGetAsync(db *gorm.DB, model *models.GrupoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.GrupoArchivosCompartidosObtenerEntity

		exec := "EXEC dbo.GrupoArchivosCompartidos_Obtener @IdGrupo = ?"

		db.Raw(exec, model.IdGrupo).Scan(&resultado)

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

func GrupoArchivoCompartidoRegistrarPostAsync(db *gorm.DB, model *models.GrupoArchivoCompartidoRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.GrupoArchivoCompartido_Registrar @IdGrupo = ?, @IdUsuario = ?, @NombreArchivo = ?, @Archivo = ?"

		db.Raw(exec, model.IdGrupo, model.IdUsuario, model.NombreArchivo, model.Archivo).Scan(&resultado)

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

func GruposMensajesObtenerGetAsync(db *gorm.DB, model *models.GruposMensajesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.GruposMensajesObtenerEntity

		exec := "EXEC dbo.GruposMensajes_Obtener @IdGrupo = ?, @UltimoMensaje = ?"

		db.Raw(exec, model.IdGrupo, model.UltimoMensaje).Scan(&resultado)

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

func GruposObtenerGetAsync(db *gorm.DB, model *models.GruposObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.GruposObtenerEntity

		exec := "EXEC dbo.Grupos_Obtener @IdUsuario = ?"

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

func GrupoMiembrosObtenerGetAsync(db *gorm.DB, model *models.GrupoMiembrosObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.GrupoMiembrosObtenerEntity

		exec := "EXEC dbo.GrupoMiembros_Obtener @IdGrupo = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdGrupo, model.IdUsuario).Scan(&resultado)

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

func GrupoTareasPendientesObtenerGetAsync(db *gorm.DB, model *models.GrupoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.GrupoTareasPendientesObtenerEntity

		exec := "EXEC dbo.GrupoMiembros_Obtener @IdGrupo = ?"

		db.Raw(exec, model.IdGrupo).Scan(&resultado)

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

func GrupoTareaPendienteDetalleObtenerGetAsync(db *gorm.DB, model *models.GrupoTareaPendienteDetalleObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.GrupoTareaPendienteDetalleObtenerEntity

		exec := "EXEC dbo.GrupoTareaPendienteDetalle_Obtener @IdTareaPendiente = ?"

		db.Raw(exec, model.IdTareaPendiente).Scan(&resultado)

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

func GrupoTareaPendienteEstatusActualizarPutAsync(db *gorm.DB, model *models.GrupoTareaPendienteEstatusActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.GrupoTareaPendienteEstatus_Actualizar @IdGrupo = ?, @IdTareaPendiente = ?, @IdUsuarioReceptor = ?, @IdEstatusTareaPendiente = ?"

		db.Raw(exec, model.IdGrupo, model.IdTareaPendiente, model.IdUsuarioReceptor, model.IdEstatusTareaPendiente).Scan(&resultado)

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

func GrupoMiembroRemoverDeleteAsync(db *gorm.DB, model *models.GrupoMiembroRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.GrupoMiembro_Remover @IdGrupo = ?, @IdProfesor = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdGrupo, model.IdProfesor, model.IdUsuario).Scan(&resultado)

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

func GrupoMiembroRegistrarPostAsync(db *gorm.DB, model *models.GrupoMiembroRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.GrupoMiembro_Registrar @IdGrupo = ?, @IdProfesor = ?, @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdGrupo, model.IdProfesor, model.IdCurso, model.IdUsuario).Scan(&resultado)

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

func GrupoTareaPendienteActualizarPutAsync(db *gorm.DB, model *models.GrupoTareaPendienteActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.GrupoTareaPendiente_Actualizar @IdGrupo = ?, @IdUsuario = ?, @IdTareaPendiente = ?, @Nombre = ?, @Descripcion = ?"

		db.Raw(exec, model.IdGrupo, model.IdUsuario, model.IdTareaPendiente, model.Nombre, model.Descripcion).Scan(&resultado)

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

func GrupoTareaPendienteRegistrarPostAsync(db *gorm.DB, model *models.GrupoTareaPendienteRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.GrupoTareaPendiente_Registrar @IdGrupo = ?, @IdUsuarioEmisor = ?, @IdUsuarioReceptor = ?, @Nombre = ?, @Descripcion = ?, @FechaFinalizacion"

		db.Raw(exec, model.IdGrupo, model.IdUsuarioEmisor, model.IdUsuarioReceptor, model.Nombre, model.Descripcion, model.FechaFinalizacion).Scan(&resultado)

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