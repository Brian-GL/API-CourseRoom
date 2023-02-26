package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/rpc"
	"os"
	"path/filepath"
	"time"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func TareaArchivosAdjuntosObtenerGetAsync(db *gorm.DB, model *models.TareaArchivosAdjuntosObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaArchivosAdjuntosObtenerEntity

		exec := "EXEC dbo.TareaArchivosAdjuntos_Obtener @IdTarea = ?"

		db.Raw(exec, model.IdTarea).Scan(&resultado)

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

func TareaArchivosEntregadosObtenerGetAsync(db *gorm.DB, model *models.TareaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaArchivosEntregadosObtenerEntity

		exec := "EXEC dbo.TareaArchivosEntregados_Obtener @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario).Scan(&resultado)

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

func TareaEstudianteDetalleObtenerGetAsync(db *gorm.DB, model *models.TareaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaEstudianteDetalleObtenerEntity

		exec := "EXEC dbo.TareaEstudianteDetalle_Obtener @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario).Scan(&resultado)

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

func TareasMesObtenerGetAsync(db *gorm.DB, model *models.TareasMesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareasMesObtenerEntity

		exec := "EXEC dbo.TareasMes_Obtener @IdUsuario = ?, @Mes = ?"

		db.Raw(exec, model.IdUsuario, model.Mes).Scan(&resultado)

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

func TareasEstudianteObtenerGetAsync(db *gorm.DB, model *models.TareasEstudianteObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareasEstudianteObtenerEntity

		exec := "EXEC dbo.TareasEstudiante_Obtener @IdUsuario = ?"

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

func TareasCreadasProfesorObtenerGetAsync(db *gorm.DB, model *models.TareasCreadasProfesorObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareasCreadasProfesorObtenerEntity

		exec := "EXEC dbo.TareasCreadasProfesor_Obtener @IdProfesor = ?"

		db.Raw(exec, model.IdProfesor).Scan(&resultado)

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

func TareaProfesorDetalleObtenerGetAsync(db *gorm.DB, model *models.TareaProfesorDetalleObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaProfesorDetalleObtenerEntity

		exec := "EXEC dbo.TareaProfesorDetalle_Obtener @IdTarea = ?, @IdProfesor = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor).Scan(&resultado)

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

func TareaImagenesEntregadasObtenerGetAsync(db *gorm.DB, model *models.TareaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaImagenesEntregadasObtenerEntity

		exec := "EXEC dbo.TareaImagenesEntregadas_Obtener @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario).Scan(&resultado)

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

func TareaReatroalimentacionesObtenerGetAsync(db *gorm.DB, model *models.TareaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaReatroalimentacionesObtenerEntity

		exec := "EXEC dbo.TareaReatroalimentaciones_Obtener @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario).Scan(&resultado)

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

func TareaRetroalimentacionDetalleObtenerGetAsync(db *gorm.DB, model *models.TareaRetroalimentacionDetalleObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareaRetroalimentacionDetalleObtenerEntity

		exec := "EXEC dbo.TareaRetroalimentacionDetalle_Obtener @IdRetroalimentacion = ?"

		db.Raw(exec, model.IdRetroalimentacion).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontró información del registro"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response
}

func TareasCalificarObtenerGetAsync(db *gorm.DB, model *models.TareasCalificarObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.TareasCalificarObtenerEntity

		exec := "EXEC dbo.TareasCalificar_Obtener @IdProfesor = ?"

		db.Raw(exec, model.IdProfesor).Scan(&resultado)

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

func TareaActualizarPutAsync(db *gorm.DB, model *models.TareaActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Tarea_Actualizar @IdTarea = ?, @IdProfesor = ?, @Nombre = ?, @Descripcion = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor, model.Nombre, model.Descripcion).Scan(&resultado)

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

func TareaEntregarActualizarPutAsync(db *gorm.DB, model *models.TareaInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaEntregar_Actualizar @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario).Scan(&resultado)

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

func TareaArchivoEntregadoRemoverDeleteAsync(db *gorm.DB, model *models.TareaArchivoEntregadoRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaArchivoEntregado_Remover @IdTarea = ?, @IdUsuario = ?, @IdArchivoEntregado"

		db.Raw(exec, model.IdTarea, model.IdUsuario, model.IdArchivoEntregado).Scan(&resultado)

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

func TareaCalificarActualizarPutAsync(db *gorm.DB, emailConfiguration *models.EmailConfiguration, COURSEROOM_CALCULATOR *string, SECRET_TOKEN *string, model *models.TareaCalificarActualizarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.TareaCalificarActualizarEntity

		exec := "EXEC dbo.TareaCalificar_Actualizar @IdTarea = ?, @IdProfesor = ?, @IdUsuario = ?, @Calificacion = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor, model.IdUsuario, model.Calificacion).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {

				var usuario *entities.UsuarioCuentaObtenerEntity

				exec = "EXEC dbo.UsuarioCuenta_Obtener @IdUsuario = ?"

				db.Raw(exec, model.IdUsuario).Scan(&usuario)

				if usuario != nil {

					modelEmail := models.CalificacionEmail{
						CorreoElectronico:    usuario.CorreoElectronico,
						NombreTarea:          *resultado.NombreTarea,
						FechaCalificacion:    time.Now().Format("10-10-2022 12:00 p.m"),
						CalificacionObtenida: *model.Calificacion,
						Anio:                 time.Now().Year()}

					go SendCalificacionEmail(emailConfiguration, &modelEmail)
				}

				// Send request to rpc courseroom calculator:

				rpc_client, err := rpc.DialHTTP("tcp", *COURSEROOM_CALCULATOR)

				if err == nil {

					modelCalculator := models.CourseRoomCalculator{
						IdUsuario:    *model.IdUsuario,
						IdDesempeno:  resultado.Codigo,
						SECRET_TOKEN: *SECRET_TOKEN}

					go func() {
						err := rpc_client.Call("Server.Calificacion", modelCalculator, nil)
						if err != nil {
							fmt.Println(err.Error())
						}
					}()

				}

				defer rpc_client.Close()

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

func TareaArchivoEntregadoRegistrarPostAsync(db *gorm.DB, model *models.TareaArchivoEntregadoRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaArchivoEntregado_Registrar @IdTarea = ?, @IdUsuario = ?, @NombreArchivo = ?, @Archivo = ?"

		db.Raw(exec, model.IdTarea, model.IdUsuario, model.NombreArchivo, model.Archivo).Scan(&resultado)

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

func TareaArchivoAdjuntoRemoverDeleteAsync(db *gorm.DB, model *models.TareaArchivoAdjuntoRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaArchivoAdjunto_Remover @IdTarea = ?, @IdProfesor = ?, @IdArchivoAdjunto"

		db.Raw(exec, model.IdTarea, model.IdProfesor, model.IdArchivoAdjunto).Scan(&resultado)

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

func TareaArchivoAdjuntoRegistrarPostAsync(db *gorm.DB, model *models.TareaArchivoAdjuntoRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaArchivoAdjunto_Registrar @IdTarea = ?, @IdProfesor = ?, @NombreArchivo = ?, @Archivo = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor, model.NombreArchivo, model.Archivo).Scan(&resultado)

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

func TareaRemoverDeleteAsync(db *gorm.DB, model *models.TareaRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Tarea_Remover @IdTarea = ?, @IdProfesor = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor).Scan(&resultado)

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

func TareaRegistrarPostAsync(db *gorm.DB, model *models.TareaRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Tarea_Registrar @IdCurso = ?, @IdProfesor = ?, @Nombre = ?, @Descripcion = ?, @FechaEntrega = ?"

		db.Raw(exec, model.IdCurso, model.IdProfesor, model.Nombre, model.Descripcion, model.FechaEntrega).Scan(&resultado)

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

func TareaRetroalimentacionRegistrarPostAsync(db *gorm.DB, model *models.TareaRetroalimentacionRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.TareaRetroalimentacion_Registrar @Idtarea = ?, @IdProfesor = ?, @IdUsuario = ?, @Nombre = ?, @Retroalimentacion = ?, @NombreArchivo = ?, @Archivo = ?"

		db.Raw(exec, model.IdTarea, model.IdProfesor, model.IdUsuario, model.Nombre, model.Retroalimentacion, model.NombreArchivo, model.Archivo).Scan(&resultado)

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

func SendCalificacionEmail(emailConfiguration *models.EmailConfiguration, data *models.CalificacionEmail) error {

	smtpPass := emailConfiguration.EMAIL_CREDENTIALS
	smtpUser := emailConfiguration.EMAIL_ADDRESS
	smtpHost := emailConfiguration.EMAIL_SERVER
	smtpPort := emailConfiguration.EMAIL_PORT

	var body bytes.Buffer

	template, err := ParseTemplateDir("app_data")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "calificacion.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", data.CorreoElectronico)
	m.SetHeader("Subject", "Nueva tarea calificada")
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

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
