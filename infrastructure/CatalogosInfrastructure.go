package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func EstadosGetAsync(db *gorm.DB, model *models.EstadosObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.EstadosObtenerEntity

		exec := "EXEC dbo.sp_csr_CatalogoEstados_Obtener @IdEstado = ?"

		db.Raw(exec, model.IdEstado).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron estados"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func EstatusTareaPendienteGetAsync(db *gorm.DB, model *models.EstatusTareaPendienteObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.EstatusTareasPendienteObtenerEntity

		exec := "EXEC dbo.sp_csr_CatalogoEstatusTareasPendientes_Obtener @IdEstatusTareaPendiente = ?"

		db.Raw(exec, model.IdEstatusTareaPendiente).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron estatus de tareas pendientes"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}
