package handler

import (
	"http-api/internal/models"
	"http-api/internal/restapi/operations/debug"

	"github.com/go-openapi/runtime/middleware"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (Health) Handle(_ debug.HealthParams) middleware.Responder {
	return debug.NewHealthOK().WithPayload(&models.Success{
		Message: "ok",
	})
}
