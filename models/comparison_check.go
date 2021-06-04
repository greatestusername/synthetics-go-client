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

// ComparisonCheck comparison check
//
// swagger:model comparison_check
type ComparisonCheck struct {

	// When set to `true`, the metrics from this check are used as a baseline for the other checks
	// Example: true
	Baseline bool `json:"baseline,omitempty"`

	// The unique ID of the check attached to the Comparison Report
	// Example: 1
	// Required: true
	ID *int32 `json:"id"`

	// An optional alias to use in place of the check name in the Comparison Report
	// Example: Example Nickname
	Nickname string `json:"nickname,omitempty"`

	// The name of the check attached to the Comparison Report
	// Example: Example Check
	Name string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ComparisonCheck) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Baseline bool `json:"baseline,omitempty"`

		ID *int32 `json:"id"`

		Nickname string `json:"nickname,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Baseline = dataAO0.Baseline

	m.ID = dataAO0.ID

	m.Nickname = dataAO0.Nickname

	// AO1
	var dataAO1 struct {
		Name string `json:"name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ComparisonCheck) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Baseline bool `json:"baseline,omitempty"`

		ID *int32 `json:"id"`

		Nickname string `json:"nickname,omitempty"`
	}

	dataAO0.Baseline = m.Baseline

	dataAO0.ID = m.ID

	dataAO0.Nickname = m.Nickname

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		Name string `json:"name,omitempty"`
	}

	dataAO1.Name = m.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this comparison check
func (m *ComparisonCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComparisonCheck) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this comparison check based on context it is used
func (m *ComparisonCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComparisonCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComparisonCheck) UnmarshalBinary(b []byte) error {
	var res ComparisonCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
