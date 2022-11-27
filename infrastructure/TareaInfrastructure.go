package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func TareaArchivosAdjuntosObtenerGetAsync(db *gorm.DB, model *models.TareaArchivosAdjuntosObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaArchivosAdjuntosObtenerEntity

		exec := "EXEC dbo.TareaArchivosAdjuntos_Obtener @IdTarea = ?"

		db.Raw(exec, model.IdTarea).Scan(&resultado)

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

func TareaEstudianteDetalleObtenerGetAsync(db *gorm.DB, model *models.TareaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaEstudianteDetalleObtenerEntity

		exec := "EXEC dbo.TareaEstudianteDetalle_Obtener @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario).Scan(&resultado)

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

func TareasMesObtenerGetAsync(db *gorm.DB, model *models.TareasMesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareasMesObtenerEntity

		exec := "EXEC dbo.TareasMes_Obtener @IdUsuario = ?, @Mes = ?"

		db.Raw(exec, model.IdUsuario, model.Mes).Scan(&resultado)

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

func TareaImagenesEntregadasObtenerGetAsync(db *gorm.DB, model *models.TareaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaImagenesEntregadasObtenerEntity

		exec := "EXEC dbo.TareaImagenesEntregadas_Obtener @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario).Scan(&resultado)

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

func TareaRetroalimentacionDetalleObtenerGetAsync(db *gorm.DB, model *models.TareaRetroalimentacionDetalleObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaImagenesEntregadasObtenerEntity

		exec := "EXEC dbo.TareaRetroalimentacionDetalle_Obtener @IdRetroalimentacion = ?"

		db.Raw(exec, model.IdRetroalimentacion).Scan(&resultado)

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

func TareaActualizarPutAsync(db *gorm.DB, model *models.TareaActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Tarea_Actualizar @IdTarea = ?, @IdProfesor = ?, @Nombre = ?, @Descripcion = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor, model.Nombre, model.Descripcion).Scan(&resultado)

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

func TareaArchivoEntregadoRegistrarPostAsync(db *gorm.DB, model *models.TareaArchivoEntregadoRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaArchivoEntregado_Registrar @IdTarea = ?, @IdUsuario = ?, @NombreArchivo = ?, @Archivo = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario, model.NombreArchivo, model.Archivo).Scan(&resultado)

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

func TareaRemoverDeleteAsync(db *gorm.DB, model *models.TareaRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Tarea_Remover @IdTarea = ?, @IdProfesor = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor).Scan(&resultado)

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

func TareaRegistrarPostAsync(db *gorm.DB, model *models.TareaRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Tarea_Registrar @IdCurso = ?, @IdProfesor = ?, @Nombre = ?, @Descripcion = ?, @FechaEntrega = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor, model.Nombre, model.Descripcion, model.FechaEntrega).Scan(&resultado)

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

func TareaRetroalimentacionRegistrarPostAsync(db *gorm.DB, model *models.TareaRetroalimentacionRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaRetroalimentacion_Registrar @Idtarea = ?, @IdProfesor = ?, @IdUsuario = ?, @Nombre = ?, @Retroalimentacion = ?, @NombreArchivo = ?, @Archivo = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor, model.IdUsuario, model.Nombre, model.Retroalimentacion, model.NombreArchivo, model.Archivo).Scan(&resultado)

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
