package controllers

import (
	"api-courseroom/async"
	"api-courseroom/infrastructure"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type GruposController struct {
	Middleware *middleware.Middleware
	JsonIter   jsoniter.API
}

func NewGruposController(middleware *middleware.Middleware) GruposController {
	return GruposController{
		Middleware: middleware,
		JsonIter:   jsoniter.ConfigCompatibleWithStandardLibrary}
}

func (controller *GruposController) GrupoActualizar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoActualizarPutAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoArchivosCompartidosObtener(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoArchivosCompartidosObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoArchivoCompartidoRegistrar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoArchivoCompartidoRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoArchivoCompartidoRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GruposMensajesObtener(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GruposMensajesObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GruposMensajesObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GruposObtener(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GruposObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GruposObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoMiembrosObtener(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoMiembrosObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoMiembrosObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoTareasPendientesObtener(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoTareasPendientesObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoTareaPendienteDetalleObtener(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoTareaPendienteDetalleObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoTareaPendienteDetalleObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoTareaPendienteEstatusActualizar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoTareaPendienteEstatusActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoTareaPendienteEstatusActualizarPutAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoMiembroRemover(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoMiembroRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoMiembroRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoMiembroRegistrar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoMiembroRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoMiembroRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoTareaPendienteActualizar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoTareaPendienteActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoTareaPendienteActualizarPutAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoTareaPendienteRegistrar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoTareaPendienteRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoTareaPendienteRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoRegistrar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoRemover(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoAbandonarActualizar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoAbandonarActualizarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoAbandonarActualizarPutAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoArchivoCompartidoRemover(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoArchivoCompartidoRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoArchivoCompartidoRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoDetalleObtener(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoDetalleObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoDetalleObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoMensajeRegistrar(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoMensajeRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoMensajeRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *GruposController) GrupoMensajeRemover(res http.ResponseWriter, req *http.Request) {

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
				var modelo *models.GrupoMensajeRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.GrupoMensajeRemoverDeleteAsync(controller.Middleware.DB, modelo)
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
