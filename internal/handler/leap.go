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
	year := p.Year
	switch {
	case year == 0:
		return operations.NewLeapYearBadRequest().WithPayload(&models.Error{
			Message: "the Gregorian calendar does not have a year \"0\"",
		})
	case year < 1:
		year += 1
		fallthrough
	default:
		leap := year%4 == 0 && year%100 != 0 || year%400 == 0
		return operations.NewLeapYearOK().WithPayload(&models.Success{
			IsLeap: &leap,
		})
	}
}
