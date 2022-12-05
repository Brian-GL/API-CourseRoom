package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/libraries"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"encoding/base64"
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

func UsuarioCredencialObtenerPostAsync(middleware *middleware.Middleware, model *models.UsuarioCredencialObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	db := middleware.DB

	if db != nil {

		var resultado *string

		exec := "EXEC dbo.UsuarioCredencial_Obtener @CorreoElectronico = ?"

		db.Raw(exec, strings.ToUpper(*model.CorreoElectronico)).Scan(&resultado)

		if resultado != nil {

			decodificacion, err := base64.StdEncoding.DecodeString(*resultado)
			if err != nil {
				response = models.ResponseInfrastructure{Status: models.ERROR, Data: err.Error()}
			} else {

				query := libraries.FormatString(middleware.QR_SERVER_API, decodificacion)

				dataCredencialesEmail := models.CredencialesEmail{
					CorreoElectronico: query,
					Anio:              time.Now().Year()}

				go middleware.SendCredencialesEmail(&dataCredencialesEmail)

				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: "Se ha enviado el correo electrónico de recuperación de credenciales"}

			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió realizar la acción"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioCuentaActualizarPutAsync(middleware *middleware.Middleware, model *models.UsuarioCuentaActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	// validar existencia email:
	responseAPI := middleware.EmailVerificatorAPI(model.CorreoElectronico)

	if responseAPI.Codigo > 0 {

		db := middleware.DB

		if db != nil {

			var resultado *entities.AccionEntity

			exec := "EXEC dbo.UsuarioCuenta_Actualizar @IdUsuario = ?, @CorreoElectronico = ?, @Contrasena = ?, @ChatsConmigo = ?, @MostrarAvisos = ?, @Imagen = ?"

			db.Raw(exec, model.IdUsuario, strings.ToUpper(*model.CorreoElectronico), model.Contrasena, model.ChatsConmigo, model.MostrarAvisos, model.Imagen).Scan(&resultado)

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
	} else {
		response = models.ResponseInfrastructure{Status: models.ALERT, Data: responseAPI.Mensaje}
	}

	return response

}
