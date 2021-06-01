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

// ComparisonCheck comparison check
//
// swagger:model comparison_check
type ComparisonCheck struct {
	ComparisonCheckGeneric

	// The name of the check attached to the Comparison Report
	// Example: Example Check
	Name string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ComparisonCheck) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ComparisonCheckGeneric
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ComparisonCheckGeneric = aO0

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

	aO0, err := swag.WriteJSON(m.ComparisonCheckGeneric)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
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

	// validation for a type composition with ComparisonCheckGeneric
	if err := m.ComparisonCheckGeneric.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this comparison check based on the context it is used
func (m *ComparisonCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ComparisonCheckGeneric
	if err := m.ComparisonCheckGeneric.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
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
