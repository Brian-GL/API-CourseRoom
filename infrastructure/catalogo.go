package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func EstadosGetAsync(db *gorm.DB, IdEstado *int) models.ResposeInfrastruture {

	var response models.ResposeInfrastruture

	if db != nil {

		var estados []entities.Estado

		exec := "EXEC dbo.sp_csr_CatalogoEstados_Obtener @IdEstado = ?"

		db.Raw(exec, IdEstado).Scan(&estados)

		if len(estados) > 0 {
			response = models.ResposeInfrastruture{Status: models.SUCCESS, Data: estados}
		} else {
			response = models.ResposeInfrastruture{Status: models.ALERT, Data: "No se encontraron estados"}
		}

	} else {
		response = models.ResposeInfrastruture{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}
