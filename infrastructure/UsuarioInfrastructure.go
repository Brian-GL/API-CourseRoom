package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/libraries"
	"api-courseroom/models"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func UsuarioActualizarPutAsync(db *gorm.DB, model *models.UsuarioActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Usuario_Actualizar @IdUsuario = ?, @Nombre = ?, @Paterno = ?, @Materno = ?, @FechaNacimiento = ?, @Genero = ?, @Descripcion = ?, @IdLocalidad = ?"

		db.Raw(exec, model.IdUsuario, model.Nombre, model.Paterno, model.Materno, model.FechaNacimiento, model.Genero, model.Descripcion, model.IdLocalidad).Scan(&resultado)

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

func UsuarioRegistrarPostAsync(db *gorm.DB, EMAIL_VERIFICATOR_API *string, emailConfiguration *models.EmailConfiguration, model *models.UsuarioRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	// validar existencia email:
	responseAPI := EmailVerificatorAPI(EMAIL_VERIFICATOR_API, model.CorreoElectronico)

	if responseAPI.Codigo > 0 {

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

					go SendBienvenidaEmail(&dataBienvenidaEmail, emailConfiguration)

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

func EmailVerificatorAPI(EMAIL_VERIFICATOR_API *string, email *string) entities.AccionEntity {

	var response entities.AccionEntity

	jsonIter := jsoniter.ConfigCompatibleWithStandardLibrary

	query := libraries.FormatString(*EMAIL_VERIFICATOR_API, *email)

	resp, err := http.Get(query)

	if err != nil {
		response = entities.AccionEntity{
			Codigo:  -1,
			Mensaje: err.Error()}
	} else {

		if resp.StatusCode == 200 {

			var modelo *models.EmailVerificatorAPISuccess

			err := jsonIter.NewDecoder(resp.Body).Decode(&modelo)

			if err != nil {
				response = entities.AccionEntity{
					Codigo:  -1,
					Mensaje: err.Error()}
			} else {

				success := modelo.Status

				if success {
					response = entities.AccionEntity{
						Codigo:  1,
						Mensaje: "Ok"}
				} else {
					response = entities.AccionEntity{
						Codigo:  -1,
						Mensaje: "El correo electrónico al que se hace referencia no existe"}
				}

			}

		} else {

			var modelo *models.EmailVerificatorAPIError

			err := jsonIter.NewDecoder(resp.Body).Decode(&modelo)

			if err != nil {
				response = entities.AccionEntity{
					Codigo:  -1,
					Mensaje: err.Error()}
			} else {
				response = entities.AccionEntity{
					Codigo:  -1,
					Mensaje: modelo.Error.Message}
			}
		}
	}

	return response
}

func SendBienvenidaEmail(data *models.BienvenidaEmail, emailConfiguration *models.EmailConfiguration) error {

	smtpPass := emailConfiguration.EMAIL_CREDENTIALS
	smtpUser := emailConfiguration.EMAIL_ADDRESS
	smtpHost := emailConfiguration.EMAIL_SERVER
	smtpPort := emailConfiguration.EMAIL_PORT

	var body bytes.Buffer

	template, err := ParseTemplateDir("app_data")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "bienvenida.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", data.CorreoElectronico)
	m.SetHeader("Subject", "Bienvenid@ a la comunidad de CourseRoom®")
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func UsuarioRemoverDeleteAsync(db *gorm.DB, model *models.UsuarioRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Usuario_Remover @IdUsuario = ?, IdTipoUsuario = ?"

		db.Raw(exec, model.IdUsuario, model.IdTipoUsuario).Scan(&resultado)

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

func UsuarioAccesoObtenerGetAsync(db *gorm.DB, model *models.UsuarioAccesoObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioAccesoObtenerEntity

		exec := "EXEC dbo.UsuarioAcceso_Obtener @CorreoElectronico = ?, @Contrasena = ?"

		db.Raw(exec, model.CorreoElectronico, model.Contrasena).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioCredencialObtenerPostAsync(db *gorm.DB, QR_SERVER_API *string, emailConfiguration *models.EmailConfiguration, model *models.UsuarioCredencialObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *string

		exec := "EXEC dbo.UsuarioCredencial_Obtener @CorreoElectronico = ?"

		db.Raw(exec, strings.ToUpper(*model.CorreoElectronico)).Scan(&resultado)

		if resultado != nil {

			decodificacion, err := base64.StdEncoding.DecodeString(*resultado)
			if err != nil {
				response = models.ResponseInfrastructure{Status: models.ERROR, Data: err.Error()}
			} else {

				password := string(decodificacion)
				query := libraries.FormatString(*QR_SERVER_API, password)

				dataCredencialesEmail := models.CredencialesEmail{
					CorreoElectronico: *model.CorreoElectronico,
					QR_URL:            query,
					Anio:              time.Now().Year()}

				go SendCredencialesEmail(&dataCredencialesEmail, emailConfiguration)

				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: "Se ha enviado el correo electrónico de recuperación de credenciales"}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "El correo al que hace referencia no se encuentra registrado"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func SendCredencialesEmail(data *models.CredencialesEmail, emailConfiguration *models.EmailConfiguration) error {

	smtpPass := emailConfiguration.EMAIL_CREDENTIALS
	smtpUser := emailConfiguration.EMAIL_ADDRESS
	smtpHost := emailConfiguration.EMAIL_SERVER
	smtpPort := emailConfiguration.EMAIL_PORT

	var body bytes.Buffer

	template, err := ParseTemplateDir("app_data")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "credenciales.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", data.CorreoElectronico)
	m.SetHeader("Subject", "Recuperación de credenciales")
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func UsuarioCuentaActualizarPutAsync(db *gorm.DB, EMAIL_VERIFICATOR_API *string, model *models.UsuarioCuentaActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	// validar existencia email:
	responseAPI := EmailVerificatorAPI(EMAIL_VERIFICATOR_API, model.CorreoElectronico)

	if responseAPI.Codigo > 0 {

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

func UsuarioCuentaObtenerGetAsync(db *gorm.DB, model *models.UsuarioInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioCuentaObtenerEntity

		exec := "EXEC dbo.UsuarioCuenta_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioDesempenoObtenerGetAsync(db *gorm.DB, model *models.UsuarioInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.UsuarioDesempenoObtenerEntity

		exec := "EXEC dbo.UsuarioDesempeno_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

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

func UsuarioDetalleObtenerGetAsync(db *gorm.DB, model *models.UsuarioInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioDetalleObtenerEntity

		exec := "EXEC dbo.UsuarioDetalle_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del registro"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuariosBuscarGetAsync(db *gorm.DB, model *models.UsuariosBuscarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.UsuariosBuscarEntity

		exec := "EXEC dbo.Usuarios_Buscar @Nombre = ?, @Paterno = ?, @Materno = ?"

		db.Raw(exec, model.Nombre, model.Paterno, model.Materno).Scan(&resultado)

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

func UsuarioSesionActualizarPutAsync(db *gorm.DB, model *models.UsuarioSesionInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.UsuarioSesion_Actualizar @IdUsuario = ?, @IdSesion = ?"

		db.Raw(exec, model.IdUsuario, model.IdSesion).Scan(&resultado)

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

func UsuarioSesionRegistrarPostAsync(db *gorm.DB, model *models.UsuarioSesionRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.UsuarioSesion_Registrar @IdUsuario = ?, @Dispositivo = ?, @Fabricante = ?, @DireccionIP = ?, @DireccionMAC = ?, @UserAgent = ?, @Navegador = ?"

		db.Raw(exec, model.IdUsuario, model.Dispositivo, model.Fabricante, model.DireccionIP, model.DireccionMAC, model.UserAgent, model.Navegador).Scan(&resultado)

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

func UsuarioSesionValidarGetAsync(db *gorm.DB, model *models.UsuarioSesionInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *bool

		exec := "EXEC dbo.UsuarioSesion_Validar @IdUsuario = ?, @IdSesion = ?"

		db.Raw(exec, model.IdUsuario, model.IdSesion).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió realizar la acción"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func UsuarioSesionesObtenerGetAsync(db *gorm.DB, model *models.UsuarioSesionesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.UsuarioSesionesObtenerEntity

		exec := "EXEC dbo.UsuarioSesiones_Obtener @IdUsuario = ?, @Activa = ?"

		db.Raw(exec, model.IdUsuario, model.Activa).Scan(&resultado)

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

func UsuarioTematicaRegistrarPostAsync(db *gorm.DB, model *models.UsuarioTematicaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.UsuarioTematica_Registrar @IdUsuario = ?, @IdTematica = ?"

		db.Raw(exec, model.IdUsuario, model.IdTematica).Scan(&resultado)

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

func UsuarioTematicaRemoverDeleteAsync(db *gorm.DB, model *models.UsuarioTematicaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.UsuarioTematica_Remover @IdUsuario = ?, @IdTematica = ?"

		db.Raw(exec, model.IdUsuario, model.IdTematica).Scan(&resultado)

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

func UsuarioTematicasObtenerGetAsync(db *gorm.DB, model *models.UsuarioTematicasObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.UsuarioTematicasObtenerEntity

		exec := "EXEC dbo.UsuarioTematicas_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if resultado != nil {

			if len(resultado) > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros"}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func UsuarioCalculatorInformacionObtenerGetAsync(db *gorm.DB, model *models.UsuarioCalculatorInformacionObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.UsuarioCalculatorInformacionObtenerEntity

		exec := "EXEC dbo.CalculatorInformacionDesempeno_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if resultado != nil {

			if len(resultado) > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros"}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron registros"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}
