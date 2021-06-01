// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotifyViaMethods notify via methods
//
// swagger:model notify_via_methods
type NotifyViaMethods struct {

	// Notify via phone call (requires that the recipient has a valid phone number and accepts phone call alerts)
	Call *bool `json:"call,omitempty"`

	// Notify via email
	Email *bool `json:"email,omitempty"`

	// Notify via SMS (requires that the recipient has a valid phone number and accepts SMS alerts)
	Sms *bool `json:"sms,omitempty"`
}

// Validate validates this notify via methods
func (m *NotifyViaMethods) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notify via methods based on context it is used
func (m *NotifyViaMethods) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotifyViaMethods) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifyViaMethods) UnmarshalBinary(b []byte) error {
	var res NotifyViaMethods
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
