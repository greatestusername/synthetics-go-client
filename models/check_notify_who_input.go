// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CheckNotifyWhoInput Where to send notifications
//
// swagger:model check_notify_who_input
type CheckNotifyWhoInput struct {

	// A custom email to notify. Other fields can be left blank when setting this.
	CustomEmail string `json:"custom_email,omitempty"`

	// The id of the user, group, or alert webhook
	ID int32 `json:"id,omitempty"`

	// The type of recipient. Can be either `user`, `group`, or `alert_webhook`.
	Type string `json:"type,omitempty"`

	NotifyViaMethods
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CheckNotifyWhoInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		CustomEmail string `json:"custom_email,omitempty"`

		ID int32 `json:"id,omitempty"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.CustomEmail = dataAO0.CustomEmail

	m.ID = dataAO0.ID

	m.Type = dataAO0.Type

	// AO1
	var aO1 NotifyViaMethods
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NotifyViaMethods = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CheckNotifyWhoInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		CustomEmail string `json:"custom_email,omitempty"`

		ID int32 `json:"id,omitempty"`

		Type string `json:"type,omitempty"`
	}

	dataAO0.CustomEmail = m.CustomEmail

	dataAO0.ID = m.ID

	dataAO0.Type = m.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.NotifyViaMethods)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this check notify who input
func (m *CheckNotifyWhoInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NotifyViaMethods
	if err := m.NotifyViaMethods.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this check notify who input based on the context it is used
func (m *CheckNotifyWhoInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NotifyViaMethods
	if err := m.NotifyViaMethods.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CheckNotifyWhoInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckNotifyWhoInput) UnmarshalBinary(b []byte) error {
	var res CheckNotifyWhoInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}