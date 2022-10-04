// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatusDescriptionError Error status details of a volume group
//
// swagger:model StatusDescriptionError
type StatusDescriptionError struct {

	// Indicates the volume group error key
	Key string `json:"key,omitempty"`

	// Failure message providing more details about the error key
	Message string `json:"message,omitempty"`

	// List of volume IDs, which failed to be added/removed to/from the volume-group, with the given error.
	VolIDs []string `json:"volIDs"`
}

// Validate validates this status description error
func (m *StatusDescriptionError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status description error based on context it is used
func (m *StatusDescriptionError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusDescriptionError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusDescriptionError) UnmarshalBinary(b []byte) error {
	var res StatusDescriptionError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
