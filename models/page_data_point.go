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

// PageDataPoint page data point
//
// swagger:model page_data_point
type PageDataPoint struct {

	// The start timestamp for the data point (UTC)
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from"`

	// The unique id for this page and run (run-level page data only).
	PageHistoryID int32 `json:"page_history_id,omitempty"`

	// The unique id for this run (run-level data only).
	RunID int32 `json:"run_id,omitempty"`

	// The end timestamp for the data point (UTC)
	// Required: true
	// Format: date-time
	To *strfmt.DateTime `json:"to"`

	// The value of the metric.
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this page data point
func (m *PageDataPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageDataPoint) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PageDataPoint) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PageDataPoint) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this page data point based on context it is used
func (m *PageDataPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PageDataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageDataPoint) UnmarshalBinary(b []byte) error {
	var res PageDataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}