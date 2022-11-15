package controllers

import (
	"api-courseroom/async"
	"api-courseroom/infrastructure"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type UsuariosController struct {
	Middleware *middleware.Middleware
	JsonIter   jsoniter.API
}

func NewUsuariosController(middleware *middleware.Middleware) UsuariosController {
	return UsuariosController{
		Middleware: middleware,
		JsonIter:   jsoniter.ConfigCompatibleWithStandardLibrary}
}

func (controller *UsuariosController) UsuarioRegistrar(res http.ResponseWriter, req *http.Request) {

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

	if token == controller.Middleware.SECRET_TOKEN {

		switch req.Method {
		case "POST":
			{
				//Registrar aviso:

				var modelo *models.UsuarioRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.UsuarioRegistrarPostAsync(controller.Middleware, modelo)
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

func (controller *UsuariosController) UsuarioCredencialesObtener(res http.ResponseWriter, req *http.Request) {

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

	if token == controller.Middleware.SECRET_TOKEN {

		switch req.Method {
		case "POST":
			{
				//Registrar aviso:

				var modelo *models.UsuarioCredencialObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.UsuarioCredencialObtenerPostAsync(controller.Middleware, modelo)
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
