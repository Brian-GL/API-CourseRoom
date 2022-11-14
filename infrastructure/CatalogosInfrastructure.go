package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func CursoEstatusGetAsync(db *gorm.DB, model *models.CursoEstatusObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.CursoEstatusObtenerEntity

		exec := "EXEC dbo.CatalogoCursoEstatus_Obtener @IdEstatusCurso = ?"

		db.Raw(exec, model.IdEstatus).Scan(&resultado)

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

func EstadosGetAsync(db *gorm.DB, model *models.EstadosObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.EstadosObtenerEntity

		exec := "EXEC dbo.CatalogoEstados_Obtener @IdEstado = ?"

		db.Raw(exec, model.IdEstado).Scan(&resultado)

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

func EstatusTareaPendienteGetAsync(db *gorm.DB, model *models.EstatusTareaPendienteObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.EstatusTareasPendienteObtenerEntity

		exec := "EXEC dbo.CatalogoTareaPendienteEstatus_Obtener @IdEstatusTareaPendiente = ?"

		db.Raw(exec, model.IdEstatusTareaPendiente).Scan(&resultado)

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

func LocalidadesGetAsync(db *gorm.DB, model *models.LocalidadesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.LocalidadesObtenerEntity

		exec := "EXEC dbo.CatalogoLocalidades_Obtener @IdEstado = ?, @IdLocalidad = ?"

		db.Raw(exec, model.IdEstado, model.IdLocalidad).Scan(&resultado)

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

func PreguntaRespuestaEstatusGetAsync(db *gorm.DB, model *models.PreguntaRespuestaEstatusObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.PreguntaRespuestaEstatusObtenerEntity

		exec := "EXEC dbo.CatalogoPreguntaRespuestaEstatus_Obtener @IdEstatusPreguntaRespuesta = ?"

		db.Raw(exec, model.IdEstatusPreguntaRespuesta).Scan(&resultado)

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

func PreguntasCuestionarioGetAsync(db *gorm.DB, model *models.PreguntasCuestionarioObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.PreguntasCuestionarioObtenerEntity

		exec := "EXEC dbo.CatalogoPreguntasCuestionario_Obtener @IdCuestionario = ?"

		db.Raw(exec, model.IdCuestionario).Scan(&resultado)

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

func TematicasGetAsync(db *gorm.DB, model *models.TematicasObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TematicasObtenerEntity

		exec := "EXEC dbo.CatalogoTematicas_Obtener @IdTematica = ?"

		db.Raw(exec, model.IdTematica).Scan(&resultado)

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

func TiposUsuarioGetAsync(db *gorm.DB, model *models.TiposUsuarioObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TiposUsuarioObtenerEntity

		exec := "EXEC dbo.CatalogoTiposUsuario_Obtener @IdTipoUsuario = ?"

		db.Raw(exec, model.IdTipoUsuario).Scan(&resultado)

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
