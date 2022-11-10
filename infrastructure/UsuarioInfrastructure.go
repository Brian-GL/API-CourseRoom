package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"strings"
	"time"
)

func UsuarioRegistrarPostAsync(middleware *middleware.Middleware, model *models.UsuarioRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	// validar existencia email:
	responseAPI := middleware.EmailVerificatorAPI(model.CorreoElectronico)

	if responseAPI.Codigo > 0 {

		db := middleware.DB

		if db != nil {

			var resultado *entities.AccionEntity

			exec := "EXEC dbo.Usuario_Registrar @Nombre = ?, @Paterno = ?, @Materno = ?, @FechaNacimiento = ?, @Genero = ?, @Descripcion = ?, @IdLocalidad = ?, @IdTipoUsuario = ?, @CorreoElectronico = ?, @Contrasena = ?, @ChatsConmigo = ?, @MostrarAvisos = ?, @Imagen = ?"

			db.Raw(exec, model.Nombre, model.Paterno, model.Materno, model.FechaNacimiento, model.Genero, model.Descripcion, model.IdLocalidad, model.IdTipoUsuario, strings.ToUpper(*model.CorreoElectronico), model.Contrasena, model.ChatsConmigo, model.MostrarAvisos, model.Imagen).Scan(&resultado)

			if resultado != nil {

				if resultado.Codigo > 0 {
					response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}

					dataBienvenidaEmail := models.BienvenidaEmail{
						CorreoElectronico: *model.CorreoElectronico,
						NombreCompleto:    *model.Nombre + " " + *model.Paterno + " " + *model.Materno,
						Nombre:            *model.Nombre,
						Anio:              time.Now().Year()}

					go middleware.SendBienvenidaEmail(&dataBienvenidaEmail)

				} else {
					response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
				}

			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió realizar la acción"}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
		}
	} else {
		response = models.ResponseInfrastructure{Status: models.ALERT, Data: responseAPI.Mensaje}
	}

	return response

}
