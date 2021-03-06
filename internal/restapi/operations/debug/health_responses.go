// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"http-api/internal/models"
)

// HealthOKCode is the HTTP code returned for type HealthOK
const HealthOKCode int = 200

/*HealthOK Success

swagger:response healthOK
*/
type HealthOK struct {

	/*
	  In: Body
	*/
	Payload *models.Success `json:"body,omitempty"`
}

// NewHealthOK creates HealthOK with default headers values
func NewHealthOK() *HealthOK {

	return &HealthOK{}
}

// WithPayload adds the payload to the health o k response
func (o *HealthOK) WithPayload(payload *models.Success) *HealthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the health o k response
func (o *HealthOK) SetPayload(payload *models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
