package infrastructure

import (
	"api-courseroom/entities"
	"api-courseroom/models"

	"gorm.io/gorm"
)

func ChatRegistrarPostAsync(db *gorm.DB, model *models.ChatRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Chat_Registrar @IdUsuarioEmisor = ?, @IdUsuarioReceptor = ?"

		db.Raw(exec, model.IdUsuarioEmisor, model.IdUsuarioReceptor).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar un nuevo chat"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func ChatRemoverDeleteAsync(db *gorm.DB, model *models.ChatRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.Chat_Remover @IdChat = ?, @IdUsuario = ?"

		db.Raw(exec, model.IdChat, model.IdUsuario).Scan(&resultado)

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

func ChatMensajeRegistrarPostAsync(db *gorm.DB, model *models.ChatMensajeRegistrarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "dbo.ChatMensaje_Registrar @IdChat = ?, @IdUsuarioEmisor = ?, @Mensaje = ?, @Archivo = ?"

		db.Raw(exec, model.IdChat, model.IdUsuarioEmisor, model.Mensaje, model.Archivo).Scan(&resultado)

		if resultado != nil {

			if resultado.Codigo > 0 {
				response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
			} else {
				response = models.ResponseInfrastructure{Status: models.ALERT, Data: resultado.Mensaje}
			}

		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se consiguió registrar un mensaje en el chat"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func ChatMensajeRemoverDeleteAsync(db *gorm.DB, model *models.ChatMensajeRemoverInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado *entities.AccionEntity

		exec := "EXEC dbo.ChatMensaje_Remover @IdChat = ?, @IdUsuarioEmisor = ?, @IdMensaje = ?"

		db.Raw(exec, model.IdChat, model.IdUsuarioEmisor, model.IdMensaje).Scan(&resultado)

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

func ChatMensajesObtenerGetAsync(db *gorm.DB, model *models.ChatMensajesObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.ChatMensajesObtenerEntity

		exec := "EXEC dbo.ChatMensajes_Obtener @IdChat = ?, @IdUsuarioLector = ?, @Leidos = ?"

		db.Raw(exec, model.IdChat, model.IdUsuarioLector, model.Leidos).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron mensajes de los chats"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func ChatsBuscarGetAsync(db *gorm.DB, model *models.ChatsBuscarInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.ChatsBuscarEntity

		exec := "EXEC dbo.Chats_Buscar @IdUsuarioEmisor = ?, @Nombre = ?, @Paterno = ?, @Materno = ?"

		db.Raw(exec, model.IdUsuarioEmisor, model.Nombre, model.Paterno, model.Materno).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontro la busqueda solicitada"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}

func ChatsObtenerGetAsync(db *gorm.DB, model *models.ChatsObtenerInputModel) models.ResponseInfrastructure {

	var response models.ResponseInfrastructure

	if db != nil {

		var resultado []entities.ChatsObtenerEntity

		exec := "EXEC dbo.Chats_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		if len(resultado) > 0 {
			response = models.ResponseInfrastructure{Status: models.SUCCESS, Data: resultado}
		} else {
			response = models.ResponseInfrastructure{Status: models.ALERT, Data: "No se encontraron chats"}
		}

	} else {
		response = models.ResponseInfrastructure{Status: models.ERROR, Data: "No se ha podido conectar a la base de datos"}
	}

	return response

}
