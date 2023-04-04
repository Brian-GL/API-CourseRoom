package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func CursoActualizarPutAsync(db *gorm.DB, model *models.CursoActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Curso_Actualizar @IdCurso = ?, @IdProfesor = ?, @Nombre = ?, @Descripcion = ?, @Imagen = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor, model.Nombre, model.Descripcion, model.Imagen).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió actualizar el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoRegistrarPostAsync(db *gorm.DB, model *models.CursoRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Curso_Registrar @Nombre = ?, @Descripcion = ?, @Imagen = ?, @IdProfesor = ?"

		db.Raw(exec, model.Nombre, model.Descripcion, model.Imagen, model.IdProfesor).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoRemoverDeleteAsync(db *gorm.DB, model *models.CursoRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Curso_Remover @IdCurso = ?, @IdProfesor = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor).Scan(&resultado)

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

func CursoAbandonarActualizarPutAsync(db *gorm.DB, model *models.CursoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoAbandonar_Actualizar @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió abandonar el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoCuestionarioContestarValidarGetAsync(db *gorm.DB, model *models.CursoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoCuestionarioContestar_Validar @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió validar las respuestar del cuestionario"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoDesempenoObtenerGetAsync(db *gorm.DB, model *models.CursoDesempenoObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoDesempenoObtenerEntity

		exec := "EXEC dbo.CursoDesempeno_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron desempeños de cursos"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoEstudianteRegistrarPostAsync(db *gorm.DB, model *models.CursoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoEstudiante_Registrar @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar al estudiante al curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoEstudianteDetalleObtenerGetAsync(db *gorm.DB, model *models.CursoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.CursoEstudianteDetalleObtenerEntity

		exec := "EXEC dbo.CursoEstudianteDetalle_Obtener @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró detalles del estudiante en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func CursoDetalleObtenerGetAsync(db *gorm.DB, model *models.CursoDetalleObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.CursoDetalleObtenerEntity

		exec := "EXEC dbo.CursoDetalle_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoFinalizarActualizarPutAsync(db *gorm.DB, model *models.CursoFinalizarActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoFinalizar_Actualizar @IdCurso = ?, @IdProfesor = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió finalizar el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoGruposObtenerGetAsync(db *gorm.DB, model *models.CursoGruposObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoGruposObtenerEntity

		exec := "EXEC dbo.CursoGrupos_Obtener @IdCurso = ?, @Activo = ?"

		db.Raw(exec, model.IdCurso, model.Activo).Scan(&resultado)

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

func CursoMaterialRegistrarPostAsync(db *gorm.DB, model *models.CursoMaterialRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoMaterial_Registrar @IdCurso = ?, @IdUsuario = ?, @NombreArchivo = ?, @Archivo = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario, model.NombreArchivo, model.Archivo).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar el material en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoMaterialRemoverDeleteAsync(db *gorm.DB, model *models.CursoMaterialRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoMaterial_Remover @IdMaterial = ?, @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdMaterial, model.IdCurso, model.IdUsuario).Scan(&resultado)

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

func CursoMaterialesObtenerGetAsync(db *gorm.DB, model *models.CursoMaterialesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoMaterialesObtenerEntity

		exec := "EXEC dbo.CursoMateriales_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron materiales en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoMensajeRegistrarPostAsync(db *gorm.DB, model *models.CursoMensajeRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoMensaje_Registrar @IdCurso = ?, @IdUsuarioEmisor = ?, @Mensaje = ?, @Archivo = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuarioEmisor, model.Mensaje, model.Archivo).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar mensajes en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoMensajeRemoverDeleteAsync(db *gorm.DB, model *models.CursoMensajeRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoMensaje_Remover @IdCurso = ?, @IdUsuarioEmisor = ?, @IdMensaje = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuarioEmisor, model.IdMensaje).Scan(&resultado)

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

func CursoMensajesObtenerGetAsync(db *gorm.DB, model *models.CursoMensajesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoMensajesObtenerEntity

		exec := "EXEC dbo.CursoMensajes_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron mensajes en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoEstudianteRemoverDeleteAsync(db *gorm.DB, model *models.CursoEstudianteRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoEstudiante_Remover @IdCurso = ?, @IdProfesor = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor, model.IdUsuario).Scan(&resultado)

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

func CursoEstudiantesObtenerGetAsync(db *gorm.DB, model *models.CursoEstudiantesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoEstudiantesObtenerEntity

		exec := "EXEC dbo.CursoEstudiantes_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron estudiantes en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoProfesorDetalleObtenerGetAsync(db *gorm.DB, model *models.CursoProfesorDetalleObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.CursoProfesorDetalleObtenerEntity

		exec := "EXEC dbo.CursoProfesorDetalle_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del profesor en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoProfesorTareasObtenerGetAsync(db *gorm.DB, model *models.CursoProfesorTareasObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoProfesorTareasObtenerEntity

		exec := "EXEC dbo.CursoProfesorTareas_Obtener @IdCurso = ?, @IdProfesor = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron tareas del profesor en el curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursosBuscarGetAsync(db *gorm.DB, model *models.CursosBuscarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursosBuscarEntity

		exec := "EXEC dbo.Cursos_Buscar @Busqueda = ?, @IdUsuario = ?"

		db.Raw(exec, model.Busqueda, model.IdUsuario).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se logro burscar el curso solicitado"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursosObtenerGetAsync(db *gorm.DB, model *models.CursosObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursosObtenerEntity

		exec := "EXEC dbo.Cursos_Obtener @IdUsuario = ?, @IdEstatusCurso = ?"

		db.Raw(exec, model.IdUsuario, model.IdEstatusCurso).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron cursos"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursosNuevosObtenerGetAsync(db *gorm.DB, model *models.CursosNuevosObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursosNuevosObtenerEntity

		exec := "EXEC dbo.CursosNuevos_Obtener @IdUsuario = ?, @NumeroResultados = ?"

		db.Raw(exec, model.IdUsuario, model.NumeroResultados).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron nuevos cursos"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursosProfesorObtenerGetAsync(db *gorm.DB, model *models.CursosProfesorObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursosProfesorObtenerEntity

		exec := "EXEC dbo.CursosProfesor_Obtener @IdProfesor = ?, @Finalizado = ?"

		db.Raw(exec, model.IdProfesor, model.Finalizado).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron cursos de profesores"}
		}
	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoTareasEstudianteObtenerGetAsync(db *gorm.DB, model *models.CursoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoTareasEstudianteObtenerEntity

		exec := "EXEC dbo.CursoTareasEstudiante_Obtener @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron tareas entregadas por los estudiantes"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoTematicaRegistrarPostAsync(db *gorm.DB, model *models.CursoTematicaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoTematica_Registrar @IdCurso = ?, @IdTematica = ?"

		db.Raw(exec, model.IdCurso, model.IdTematica).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar la nueva tematica"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoTematicaRemoverDeleteAsync(db *gorm.DB, model *models.CursoTematicaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoTematica_Remover @IdCurso = ?, @IdTematica = ?"

		db.Raw(exec, model.IdCurso, model.IdTematica).Scan(&resultado)

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

func CursoTematicasObtenerGetAsync(db *gorm.DB, model *models.CursoTematicasObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoTematicasObtenerEntity

		exec := "EXEC dbo.CursoTematicas_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron tematicas en cursos"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoEstudianteDesempenoObtenerGetAsync(db *gorm.DB, model *models.CursoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoEstudianteDesempenoObtenerEntity

		exec := "EXEC dbo.CursoEstudianteDesempeno_Obtener @IdCurso = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros del desempeno de los estudiantes"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoEstudiantesSinGrupoObtenerGetAsync(db *gorm.DB, model *models.CursoEstudiantesSinGrupoObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoEstudiantesSinGrupoObtenerEntity

		exec := "EXEC dbo.CursoEstudiantesSinGrupo_Obtener @IdCurso = ?"

		db.Raw(exec, model.IdCurso).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron estudiantes sin grupos"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoEstudianteFinalizarActualizarPutAsync(db *gorm.DB, model *models.CursoInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoEstudianteFinalizar_Actualizar @IdCurso = ?, IdUsuario = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado.Mensaje}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió finalizar al estudiante del curso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func CursoCuestionarioRespuestaRegistrarPostAsync(db *gorm.DB, model *models.CursoCuestionarioRespuestaRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.CursoCuestionarioRespuesta_Registrar @IdCurso = ?, @IdUsuario = ?, @IdPregunta = ?, @Puntaje = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario, model.IdPregunta, model.Puntaje).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar las respuestas"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}
