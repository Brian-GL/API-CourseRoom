package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/libraries"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"encoding/base64"
	"strings"
	"time"

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
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
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

func UsuarioCuentaObtenerGetAsync(db *gorm.DB, model *models.UsuarioInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioCuentaObtenerEntity

		exec := "EXEC dbo.UsuarioCuenta_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioDesempenoObtenerGetAsync(db *gorm.DB, model *models.UsuarioInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioDesempenoObtenerEntity

		exec := "EXEC dbo.UsuarioDesempeno_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioDesempenoRegistrarPostAsync(db *gorm.DB, model *models.UsuarioDesempenoRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.UsuarioDesempeno_Registrar @IdUsuario = ?, @IdTarea = ?, @Calificacion = ?, @PromedioCurso = ?, @PrediccionPromedioCurso = ?, @RumboPromedioCurso = ?, @PromedioGeneral = ?, @PrediccionPromedioGeneral = ?, @RumboPromedioGeneral = ?, @PuntualidadCurso = ?, @PrediccionPuntualidadCurso = ?, @RumboPuntualidadCurso = ?, @PuntualidadGeneral = ?, @PrediccionPuntualidadGeneral = ?, @RumboPuntualidadGeneral = ?"

		db.Raw(exec, model.IdUsuario, model.IdTarea, model.Calificacion, model.PromedioCurso, model.PrediccionPromedioCurso, model.RumboPromedioCurso, model.PromedioGeneral, model.PrediccionPromedioGeneral, model.RumboPromedioGeneral, model.PuntualidadCurso, model.PrediccionPuntualidadCurso, model.RumboPuntualidadCurso, model.PuntualidadGeneral, model.PrediccionPuntualidadGeneral, model.RumboPuntualidadGeneral).Scan(&resultado)

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

func UsuarioDetalleObtenerGetAsync(db *gorm.DB, model *models.UsuarioInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioDetalleObtenerEntity

		exec := "EXEC dbo.UsuarioDetalle_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioNuevaPuntualidadCursoObtenerGetAsync(db *gorm.DB, model *models.UsuarioNuevaPuntualidadCursoObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioPuntualidadEntity

		exec := "EXEC dbo.UsuarioNuevaPuntualidadCurso_Obtener @IdCurso = ?, @IdUsuario = ?, @Puntualidad = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario, model.Puntualidad).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioNuevaPuntualidadGeneralObtenerGetAsync(db *gorm.DB, model *models.UsuarioNuevaPuntualidadGeneralObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioPuntualidadEntity

		exec := "EXEC dbo.UsuarioNuevaPuntualidadCurso_Obtener @IdUsuario = ?, @Puntualidad = ?"

		db.Raw(exec, model.IdUsuario, model.Puntualidad).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioNuevoPromedioCursoObtenerGetAsync(db *gorm.DB, model *models.UsuarioNuevoPromedioCursoObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioPromedioEntity

		exec := "EXEC dbo.UsuarioNuevoPromedioCurso_Obtener @IdCurso = ?, @IdUsuario = ?, @Calificacion = ?"

		db.Raw(exec, model.IdCurso, model.IdUsuario, model.Calificacion).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func UsuarioNuevoPromedioGeneralObtenerGetAsync(db *gorm.DB, model *models.UsuarioNuevoPromedioGeneralObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.UsuarioPromedioEntity

		exec := "EXEC dbo.UsuarioNuevoPromedioGeneral_Obtener @IdUsuario = ?, @Calificacion = ?"

		db.Raw(exec, model.IdUsuario, model.Calificacion).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
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

		var resultado *entities.UsuarioSesionesObtenerEntity

		exec := "EXEC dbo.UsuarioSesiones_Obtener @IdUsuario = ?, @Activa = ?"

		db.Raw(exec, model.IdUsuario, model.Activa).Scan(&resultado)

		if resultado != nil {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del aviso"}
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
