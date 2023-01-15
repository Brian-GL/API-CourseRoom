package controllers

import (
	"api-courseroom/middleware"

	jsoniter "github.com/json-iterator/go"
)

type PreguntasController struct {
	Middleware *middleware.Middleware
	JsonIter   jsoniter.API
}

func NewPreguntasController(middleware *middleware.Middleware) PreguntasController {
	return PreguntasController{
		Middleware: middleware,
		JsonIter:   jsoniter.ConfigCompatibleWithStandardLibrary}
}
