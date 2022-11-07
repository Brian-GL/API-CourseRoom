package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func CursoRemoverDeleteAsync(db *gorm.DB, model *models.CursoRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Curso_Remover @IdCurso = ?, @IdProfesor = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor).Scan(&resultado)

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

func CursoRegistrarPostAsync(db *gorm.DB, model *models.CursoRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Curso_Registrar  @Nombre = ?, @Descripcion = ?, @Imagen = ?, @IdProfesor = ?"

		db.Raw(exec, model.Nombre, model.Descripcion, model.Imagen, model.IdProfesor).Scan(&resultado)

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
