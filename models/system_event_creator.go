// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemEventCreator system event creator
//
// swagger:model system_event_creator
type SystemEventCreator struct {

	// The unique id for the creator
	ID int32 `json:"id,omitempty"`

	// A URL to the user's gravatar image, if available
	// Example: https://secure.gravatar.com/avatar/0.jpg?r=g\u0026d=mm\u0026s=60
	ImageURL string `json:"image_url,omitempty"`

	// The name of the creator
	// Example: Jane Doe
	Name string `json:"name,omitempty"`

	// The type of the creator
	// Example: User
	Type string `json:"type,omitempty"`
}

// Validate validates this system event creator
func (m *SystemEventCreator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system event creator based on context it is used
func (m *SystemEventCreator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemEventCreator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemEventCreator) UnmarshalBinary(b []byte) error {
	var res SystemEventCreator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
