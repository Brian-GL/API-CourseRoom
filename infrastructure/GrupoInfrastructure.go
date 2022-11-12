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
