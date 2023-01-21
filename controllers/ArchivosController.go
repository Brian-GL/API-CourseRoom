package controllers

import (
	"api-courseroom/async"
	"api-courseroom/infrastructure"
	"api-courseroom/models"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type ArchivosController struct {
	DB           *gorm.DB
	Validator    *validator.Validate
	SECRET_TOKEN string
	JsonIter     jsoniter.API
}

func NewArchivosController() ArchivosController {

	//godotenv.Load(".env")

	server := os.Getenv("SERVER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DATABASE")
	secretToken := os.Getenv("SECRET_TOKEN")

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName

	db, _ := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	return ArchivosController{
		SECRET_TOKEN: secretToken,
		DB:           db,
		Validator:    validator.New(),
		JsonIter:     jsoniter.ConfigCompatibleWithStandardLibrary}
}

func (controller *ArchivosController) ValidateModel(data interface{}) error {
	return controller.Validator.Struct(data)
}

func (controller *ArchivosController) ArchivoActualizar(res http.ResponseWriter, req *http.Request) {

	// Cabecera de respuesta:
	res.Header().Add("Content-Type", "application/json")

	// Obtener token
	token := req.Header.Get("Authorization")

	// Validar que el token no se encuentre vacío:
	if token == "" {

		jsonBytes, err := controller.JsonIter.Marshal("El token es necesario para acceder a este recurso")

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
		} else {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write(jsonBytes)
		}

		return
	}

	// Validar que el token sea el correcto:

	if token == controller.SECRET_TOKEN {

		switch req.Method {

		case "PUT":
			{
				//Actualizar aviso:
				var modelo *models.ArchivoActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ArchivoActualizarPutAsync(controller.DB, modelo)
						})

						response := future.Await().(models.ResponseInfrastructure)

						switch response.Status {
						case models.SUCCESS:
							{
								jsonBytes, err := controller.JsonIter.Marshal(response.Data)
								if err != nil {
									res.WriteHeader(http.StatusInternalServerError)
									res.Write([]byte(err.Error()))
								} else {
									res.WriteHeader(http.StatusOK)
									res.Write(jsonBytes)
								}
							}
						case models.ALERT:
							{
								jsonBytes, err := controller.JsonIter.Marshal(response.Data)
								if err != nil {
									res.WriteHeader(http.StatusInternalServerError)
									res.Write([]byte(err.Error()))
								} else {
									res.WriteHeader(http.StatusConflict)
									res.Write(jsonBytes)
								}
							}
						default:
							{
								jsonBytes, err := controller.JsonIter.Marshal(response.Data)
								if err != nil {
									res.WriteHeader(http.StatusInternalServerError)
									res.Write([]byte(err.Error()))
								} else {
									res.WriteHeader(http.StatusInternalServerError)
									res.Write(jsonBytes)
								}
							}
						}
					} else {

						jsonBytes, err := controller.JsonIter.Marshal("El parámetro de entrada no cuenta con un formato adecuado")

						if err != nil {
							res.WriteHeader(http.StatusInternalServerError)
							res.Write([]byte(err.Error()))
						} else {
							res.WriteHeader(http.StatusBadRequest)
							res.Write(jsonBytes)
						}
					}
				} else {

					jsonBytes, err := controller.JsonIter.Marshal("El parámetro de entrada no cuenta con un formato adecuado")

					if err != nil {
						res.WriteHeader(http.StatusInternalServerError)
						res.Write([]byte(err.Error()))
					} else {
						res.WriteHeader(http.StatusConflict)
						res.Write(jsonBytes)
					}
				}

			}

		default:
			{
				jsonBytes, err := controller.JsonIter.Marshal("Ruta inválida")

				if err != nil {
					res.WriteHeader(http.StatusInternalServerError)
					res.Write([]byte(err.Error()))
				} else {
					res.WriteHeader(http.StatusNotImplemented)
					res.Write(jsonBytes)
				}
			}
		}

	} else {

		jsonBytes, err := controller.JsonIter.Marshal("Token inválido")

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
		} else {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write(jsonBytes)
		}
	}
}
