// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NotificationsInput Configure how and when alerts are sent
//
// swagger:model notifications_input
type NotificationsInput struct {

	// escalations
	Escalations []*CheckEscalationInput `json:"escalations"`

	// Muted checks do not send any alert notifications
	Muted *bool `json:"muted,omitempty"`

	// Alert once the number of failed runs reaches this threshold.
	//                                    Recommended threshold is 2.
	// Maximum: 10
	// Minimum: 1
	NotifyAfterFailureCount int32 `json:"notify_after_failure_count,omitempty"`

	// Alert if the check is failing from only one location
	NotifyOnLocationFailure *bool `json:"notify_on_location_failure,omitempty"`

	// notify who
	// Unique: true
	NotifyWho []*CheckNotifyWhoInput `json:"notify_who"`

	NotifyViaMethods
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NotificationsInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Escalations []*CheckEscalationInput `json:"escalations"`

		Muted *bool `json:"muted,omitempty"`

		NotifyAfterFailureCount int32 `json:"notify_after_failure_count,omitempty"`

		NotifyOnLocationFailure *bool `json:"notify_on_location_failure,omitempty"`

		NotifyWho []*CheckNotifyWhoInput `json:"notify_who"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Escalations = dataAO0.Escalations

	m.Muted = dataAO0.Muted

	m.NotifyAfterFailureCount = dataAO0.NotifyAfterFailureCount

	m.NotifyOnLocationFailure = dataAO0.NotifyOnLocationFailure

	m.NotifyWho = dataAO0.NotifyWho

	// AO1
	var aO1 NotifyViaMethods
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NotifyViaMethods = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NotificationsInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Escalations []*CheckEscalationInput `json:"escalations"`

		Muted *bool `json:"muted,omitempty"`

		NotifyAfterFailureCount int32 `json:"notify_after_failure_count,omitempty"`

		NotifyOnLocationFailure *bool `json:"notify_on_location_failure,omitempty"`

		NotifyWho []*CheckNotifyWhoInput `json:"notify_who"`
	}

	dataAO0.Escalations = m.Escalations

	dataAO0.Muted = m.Muted

	dataAO0.NotifyAfterFailureCount = m.NotifyAfterFailureCount

	dataAO0.NotifyOnLocationFailure = m.NotifyOnLocationFailure

	dataAO0.NotifyWho = m.NotifyWho

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

// Validate validates this notifications input
func (m *NotificationsInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEscalations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyAfterFailureCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifyWho(formats); err != nil {
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

func (m *NotificationsInput) validateEscalations(formats strfmt.Registry) error {

	if swag.IsZero(m.Escalations) { // not required
		return nil
	}

	for i := 0; i < len(m.Escalations); i++ {
		if swag.IsZero(m.Escalations[i]) { // not required
			continue
		}

		if m.Escalations[i] != nil {
			if err := m.Escalations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("escalations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NotificationsInput) validateNotifyAfterFailureCount(formats strfmt.Registry) error {

	if swag.IsZero(m.NotifyAfterFailureCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("notify_after_failure_count", "body", int64(m.NotifyAfterFailureCount), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("notify_after_failure_count", "body", int64(m.NotifyAfterFailureCount), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *NotificationsInput) validateNotifyWho(formats strfmt.Registry) error {

	if swag.IsZero(m.NotifyWho) { // not required
		return nil
	}

	if err := validate.UniqueItems("notify_who", "body", m.NotifyWho); err != nil {
		return err
	}

	for i := 0; i < len(m.NotifyWho); i++ {
		if swag.IsZero(m.NotifyWho[i]) { // not required
			continue
		}

		if m.NotifyWho[i] != nil {
			if err := m.NotifyWho[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notify_who" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this notifications input based on the context it is used
func (m *NotificationsInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEscalations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifyWho(ctx, formats); err != nil {
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

func (m *NotificationsInput) contextValidateEscalations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Escalations); i++ {

		if m.Escalations[i] != nil {
			if err := m.Escalations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("escalations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NotificationsInput) contextValidateNotifyWho(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NotifyWho); i++ {

		if m.NotifyWho[i] != nil {
			if err := m.NotifyWho[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notify_who" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationsInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationsInput) UnmarshalBinary(b []byte) error {
	var res NotificationsInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}