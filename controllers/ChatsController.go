package controllers

import (
	"api-courseroom/async"
	"api-courseroom/infrastructure"
	"api-courseroom/middleware"
	"api-courseroom/models"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type ChatController struct {
	Middleware *middleware.Middleware
	JsonIter   jsoniter.API
}

func NewChatController(middleware *middleware.Middleware) ChatController {
	return ChatController{
		Middleware: middleware,
		JsonIter:   jsoniter.ConfigCompatibleWithStandardLibrary}
}

func (controller *ChatController) ChatRegistrar(res http.ResponseWriter, req *http.Request) {

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

				var modelo *models.ChatRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ChatRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *ChatController) ChatRemover(res http.ResponseWriter, req *http.Request) {

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
				//Registrar aviso:

				var modelo *models.ChatRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ChatRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *ChatController) ChatMensajeRegistrar(res http.ResponseWriter, req *http.Request) {

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

				var modelo *models.ChatMensajeRegistrarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ChatMensajeRegistrarPostAsync(controller.Middleware.DB, modelo)
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

func (controller *ChatController) ChatMensajeRemover(res http.ResponseWriter, req *http.Request) {

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
				//Registrar aviso:

				var modelo *models.ChatMensajeRemoverInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ChatMensajeRemoverDeleteAsync(controller.Middleware.DB, modelo)
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

func (controller *ChatController) ChatMensajesObtener(res http.ResponseWriter, req *http.Request) {

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

				var modelo *models.ChatMensajesObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ChatMensajesObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *ChatController) ChatsObtener(res http.ResponseWriter, req *http.Request) {

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

				var modelo *models.ChatsObtenerInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ChatsObtenerGetAsync(controller.Middleware.DB, modelo)
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

func (controller *ChatController) ChatsBuscar(res http.ResponseWriter, req *http.Request) {

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

				var modelo *models.ChatsBuscarInputModel

				err := controller.JsonIter.NewDecoder(req.Body).Decode(&modelo)

				if err == nil {

					err = controller.Middleware.ValidateModel(modelo)

					if err == nil {

						future := async.Exec(func() interface{} {
							return infrastructure.ChatsBuscarGetAsync(controller.Middleware.DB, modelo)
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
