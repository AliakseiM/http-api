// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewLeapYearParams creates a new LeapYearParams object
//
// There are no default values defined in the spec.
func NewLeapYearParams() LeapYearParams {

	return LeapYearParams{}
}

// LeapYearParams contains all the bound params for the leap year operation
// typically these are obtained from a http.Request
//
// swagger:parameters leapYear
type LeapYearParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Year
	  Required: true
	  Maximum: 9999
	  Minimum: -9999
	  In: query
	*/
	Year int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewLeapYearParams() beforehand.
func (o *LeapYearParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qYear, qhkYear, _ := qs.GetOK("year")
	if err := o.bindYear(qYear, qhkYear, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindYear binds and validates parameter Year from query.
func (o *LeapYearParams) bindYear(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("year", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("year", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("year", "query", "int64", raw)
	}
	o.Year = value

	if err := o.validateYear(formats); err != nil {
		return err
	}

	return nil
}

// validateYear carries on validations for parameter Year
func (o *LeapYearParams) validateYear(formats strfmt.Registry) error {

	if err := validate.MinimumInt("year", "query", o.Year, -9999, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("year", "query", o.Year, 9999, false); err != nil {
		return err
	}

	return nil
}