// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Success success
//
// swagger:model Success
type Success struct {

	// message
	// Example: OK
	Message string `json:"message,omitempty"`
}

// Validate validates this success
func (m *Success) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this success based on context it is used
func (m *Success) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Success) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Success) UnmarshalBinary(b []byte) error {
	var res Success
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
