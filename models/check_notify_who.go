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

// CheckNotifyWho Where to send notifications
//
// swagger:model check_notify_who
type CheckNotifyWho struct {

	// The recipient's email, if notifying a custom email address
	CustomUserEmail string `json:"custom_user_email,omitempty"`

	// The id of the user, group, or alert webhook
	ID int32 `json:"id,omitempty"`

	// links
	Links *CheckNotifyWhoAO0Links `json:"links,omitempty"`

	// The type of recipient. Can be either `user`, `group`, or `alert_webhook`.
	Type string `json:"type,omitempty"`

	NotifyViaMethods
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CheckNotifyWho) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		CustomUserEmail string `json:"custom_user_email,omitempty"`

		ID int32 `json:"id,omitempty"`

		Links *CheckNotifyWhoAO0Links `json:"links,omitempty"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.CustomUserEmail = dataAO0.CustomUserEmail

	m.ID = dataAO0.ID

	m.Links = dataAO0.Links

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
func (m CheckNotifyWho) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		CustomUserEmail string `json:"custom_user_email,omitempty"`

		ID int32 `json:"id,omitempty"`

		Links *CheckNotifyWhoAO0Links `json:"links,omitempty"`

		Type string `json:"type,omitempty"`
	}

	dataAO0.CustomUserEmail = m.CustomUserEmail

	dataAO0.ID = m.ID

	dataAO0.Links = m.Links

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

// Validate validates this check notify who
func (m *CheckNotifyWho) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with NotifyViaMethods
	if err := m.NotifyViaMethods.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckNotifyWho) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this check notify who based on the context it is used
func (m *CheckNotifyWho) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with NotifyViaMethods
	if err := m.NotifyViaMethods.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckNotifyWho) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckNotifyWho) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckNotifyWho) UnmarshalBinary(b []byte) error {
	var res CheckNotifyWho
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CheckNotifyWhoAO0Links check notify who a o0 links
//
// swagger:model CheckNotifyWhoAO0Links
type CheckNotifyWhoAO0Links struct {

	// The html view for this recipient, if available
	SelfHTML string `json:"self_html,omitempty"`
}

// Validate validates this check notify who a o0 links
func (m *CheckNotifyWhoAO0Links) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check notify who a o0 links based on context it is used
func (m *CheckNotifyWhoAO0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckNotifyWhoAO0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckNotifyWhoAO0Links) UnmarshalBinary(b []byte) error {
	var res CheckNotifyWhoAO0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
