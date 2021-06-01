// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemEventDescription system event description
//
// swagger:model system_event_description
type SystemEventDescription struct {

	// A detailed list describing the event
	// Example: ["Name changed","Steps changed"]
	Detail []string `json:"detail"`

	// A summary of the event
	// Example: Jane Doe made the following changes:
	Summary string `json:"summary,omitempty"`
}

// Validate validates this system event description
func (m *SystemEventDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system event description based on context it is used
func (m *SystemEventDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemEventDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemEventDescription) UnmarshalBinary(b []byte) error {
	var res SystemEventDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
