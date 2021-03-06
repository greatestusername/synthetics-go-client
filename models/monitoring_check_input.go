// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MonitoringCheckInput A monitoring check
//
// swagger:model monitoring_check_input
type MonitoringCheckInput struct {
	CheckInput

	// When enabled, the check will retry up to two times from the same location after a failed run. Ensure your account plan supports this feature before enabling.
	AutoRetry *bool `json:"auto_retry,omitempty"`

	// True if the check is not paused
	Enabled *bool `json:"enabled,omitempty"`

	// http request headers
	HTTPRequestHeaders interface{} `json:"http_request_headers,omitempty"`

	// The integrations to send metrics to
	Integrations []int32 `json:"integrations"`

	// The locations to run the check from
	Locations []int32 `json:"locations"`

	// notifications
	Notifications *NotificationsInput `json:"notifications,omitempty"`

	// Mark a run as a failure if the total response time
	//                                    is above this threshold (in milliseconds)
	// Maximum: 60000
	// Minimum: 0
	ResponseTimeMonitorMilliseconds *int32 `json:"response_time_monitor_milliseconds,omitempty"`

	// When enabled, the check cycles through locations round-robin style with each run.Ensure your account plan supports concurrent checks before disabling.
	RoundRobin *bool `json:"round_robin,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MonitoringCheckInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CheckInput
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CheckInput = aO0

	// AO1
	var dataAO1 struct {
		AutoRetry *bool `json:"auto_retry,omitempty"`

		Enabled *bool `json:"enabled,omitempty"`

		HTTPRequestHeaders interface{} `json:"http_request_headers,omitempty"`

		Integrations []int32 `json:"integrations"`

		Locations []int32 `json:"locations"`

		Notifications *NotificationsInput `json:"notifications,omitempty"`

		ResponseTimeMonitorMilliseconds *int32 `json:"response_time_monitor_milliseconds,omitempty"`

		RoundRobin *bool `json:"round_robin,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AutoRetry = dataAO1.AutoRetry

	m.Enabled = dataAO1.Enabled

	m.HTTPRequestHeaders = dataAO1.HTTPRequestHeaders

	m.Integrations = dataAO1.Integrations

	m.Locations = dataAO1.Locations

	m.Notifications = dataAO1.Notifications

	m.ResponseTimeMonitorMilliseconds = dataAO1.ResponseTimeMonitorMilliseconds

	m.RoundRobin = dataAO1.RoundRobin

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MonitoringCheckInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CheckInput)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AutoRetry *bool `json:"auto_retry,omitempty"`

		Enabled *bool `json:"enabled,omitempty"`

		HTTPRequestHeaders interface{} `json:"http_request_headers,omitempty"`

		Integrations []int32 `json:"integrations"`

		Locations []int32 `json:"locations"`

		Notifications *NotificationsInput `json:"notifications,omitempty"`

		ResponseTimeMonitorMilliseconds *int32 `json:"response_time_monitor_milliseconds,omitempty"`

		RoundRobin *bool `json:"round_robin,omitempty"`
	}

	dataAO1.AutoRetry = m.AutoRetry

	dataAO1.Enabled = m.Enabled

	dataAO1.HTTPRequestHeaders = m.HTTPRequestHeaders

	dataAO1.Integrations = m.Integrations

	dataAO1.Locations = m.Locations

	dataAO1.Notifications = m.Notifications

	dataAO1.ResponseTimeMonitorMilliseconds = m.ResponseTimeMonitorMilliseconds

	dataAO1.RoundRobin = m.RoundRobin

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this monitoring check input
func (m *MonitoringCheckInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CheckInput
	if err := m.CheckInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTimeMonitorMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitoringCheckInput) validateNotifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *MonitoringCheckInput) validateResponseTimeMonitorMilliseconds(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseTimeMonitorMilliseconds) { // not required
		return nil
	}

	if err := validate.MinimumInt("response_time_monitor_milliseconds", "body", int64(*m.ResponseTimeMonitorMilliseconds), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("response_time_monitor_milliseconds", "body", int64(*m.ResponseTimeMonitorMilliseconds), 60000, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this monitoring check input based on the context it is used
func (m *MonitoringCheckInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CheckInput
	if err := m.CheckInput.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitoringCheckInput) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifications != nil {
		if err := m.Notifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MonitoringCheckInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitoringCheckInput) UnmarshalBinary(b []byte) error {
	var res MonitoringCheckInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
