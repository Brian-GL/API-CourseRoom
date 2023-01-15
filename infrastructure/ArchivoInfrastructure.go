package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func ArchivoActualizarPutAsync(db *gorm.DB, model *models.ArchivoActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Archivo_Actualizar @Archivo = ?, @IdTipoArchivo = ?, @IdRegistro"

		db.Raw(exec, model.Archivo, model.IdTipoArchivo, model.IdRegistro).Scan(&resultado)

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
