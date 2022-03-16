package handler

import (
	"http-api/internal/models"
	"http-api/internal/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type IsLeap struct{}

func NewIsLeap() *IsLeap {
	return &IsLeap{}
}

func (IsLeap) Handle(p operations.LeapYearParams) middleware.Responder {
	if p.Year == 0 {
		return operations.NewLeapYearInternalServerError()
	}

	var leap bool
	if p.Year%4 == 0 && p.Year%100 != 0 || p.Year%400 == 0 {
		leap = true
	}

	return operations.NewLeapYearOK().WithPayload(&models.Success{
		IsLeap: &leap,
	})
}
