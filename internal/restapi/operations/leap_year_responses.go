// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"http-api/internal/models"
)

// LeapYearOKCode is the HTTP code returned for type LeapYearOK
const LeapYearOKCode int = 200

/*LeapYearOK Success

swagger:response leapYearOK
*/
type LeapYearOK struct {

	/*
	  In: Body
	*/
	Payload *models.Success `json:"body,omitempty"`
}

// NewLeapYearOK creates LeapYearOK with default headers values
func NewLeapYearOK() *LeapYearOK {

	return &LeapYearOK{}
}

// WithPayload adds the payload to the leap year o k response
func (o *LeapYearOK) WithPayload(payload *models.Success) *LeapYearOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the leap year o k response
func (o *LeapYearOK) SetPayload(payload *models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LeapYearOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LeapYearBadRequestCode is the HTTP code returned for type LeapYearBadRequest
const LeapYearBadRequestCode int = 400

/*LeapYearBadRequest bad request

swagger:response leapYearBadRequest
*/
type LeapYearBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLeapYearBadRequest creates LeapYearBadRequest with default headers values
func NewLeapYearBadRequest() *LeapYearBadRequest {

	return &LeapYearBadRequest{}
}

// WithPayload adds the payload to the leap year bad request response
func (o *LeapYearBadRequest) WithPayload(payload *models.Error) *LeapYearBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the leap year bad request response
func (o *LeapYearBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LeapYearBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LeapYearInternalServerErrorCode is the HTTP code returned for type LeapYearInternalServerError
const LeapYearInternalServerErrorCode int = 500

/*LeapYearInternalServerError internal server error

swagger:response leapYearInternalServerError
*/
type LeapYearInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLeapYearInternalServerError creates LeapYearInternalServerError with default headers values
func NewLeapYearInternalServerError() *LeapYearInternalServerError {

	return &LeapYearInternalServerError{}
}

// WithPayload adds the payload to the leap year internal server error response
func (o *LeapYearInternalServerError) WithPayload(payload *models.Error) *LeapYearInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the leap year internal server error response
func (o *LeapYearInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LeapYearInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}