package controllers

import (
	"api-courseroom/async"
	"api-courseroom/infrastructure"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type TareasController struct {
	Middleware *middleware.Middleware
	JsonIter   jsoniter.API
}

func NewTareasController(middleware *middleware.Middleware) TareasController {
	return TareasController{
		Middleware: middleware,
		JsonIter:   jsoniter.ConfigCompatibleWithStandardLibrary}
}

func (controller *TareasController) TareaArchivosAdjuntosObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaArchivosAdjuntosObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaArchivosAdjuntosObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaEstudianteDetalleObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaEstudianteDetalleObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareasMesObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareasMesObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareasMesObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaImagenesEntregadasObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaImagenesEntregadasObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaRetroalimentacionDetalleObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaRetroalimentacionDetalleObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaRetroalimentacionDetalleObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaActualizar(res http.ResponseWriter, req *http.Request) {

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

		case "PUT":
			{
				//Actualizar grupo:
				var modelo *models.TareaActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaActualizarPutAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaArchivoEntregadoRegistrar(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaArchivoEntregadoRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaArchivoEntregadoRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaRemover(res http.ResponseWriter, req *http.Request) {

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

		case "DELETE":
			{
				//Actualizar grupo:
				var modelo *models.TareaRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaRegistrar(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaRetroalimentacionRegistrar(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaRetroalimentacionRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaRetroalimentacionRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaCalificarActualizar(res http.ResponseWriter, req *http.Request) {

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

		case "PUT":
			{
				//Actualizar grupo:
				var modelo *models.TareaCalificarActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaCalificarActualizarPutAsync(controller.Middleware, modelo)
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

func (controller *TareasController) TareaArchivosEntregadosObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaArchivosEntregadosObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaArchivosEntregadosObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaEstudianteObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaEstudianteObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaEstudianteObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaCreadaProfesorObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaCreadaProfesorObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaCreadaProfesorObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaProfesorDetalleObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaProfesorDetalleObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaProfesorDetalleObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaReatroalimentacionesObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaReatroalimentacionesObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaReatroalimentacionesObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaCalificarObtener(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaCalificarObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaCalificarObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaEntregarActualizar(res http.ResponseWriter, req *http.Request) {

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

		case "PUT":
			{
				//Actualizar grupo:
				var modelo *models.TareaEntregarActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaEntregarActualizarPutAsync(controller.Middleware, modelo)
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

func (controller *TareasController) TareaArchivoEntregadoRemover(res http.ResponseWriter, req *http.Request) {

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

		case "DELETE":
			{
				//Actualizar grupo:
				var modelo *models.TareaArchivoEntregadoRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaArchivoEntregadoRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaArchivoAdjuntoRemover(res http.ResponseWriter, req *http.Request) {

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

		case "DELETE":
			{
				//Actualizar grupo:
				var modelo *models.TareaArchivoAdjuntoRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaArchivoAdjuntoRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *TareasController) TareaArchivoAdjuntoRegistrar(res http.ResponseWriter, req *http.Request) {

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
				//Actualizar grupo:
				var modelo *models.TareaArchivoAdjuntoRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.TareaArchivoAdjuntoRegistrarPostAsync(controller.Middleware.DB, modelo)
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
