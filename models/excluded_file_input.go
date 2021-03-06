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

// ExcludedFileInput excluded file input
//
// swagger:model excluded_file_input
type ExcludedFileInput struct {
	ExcludedFileBasics

	// A regular expression to match against all URLs visited during the check run
	// Example: static\\.chartbeat\\.com
	Pattern string `json:"pattern,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ExcludedFileInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ExcludedFileBasics
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ExcludedFileBasics = aO0

	// AO1
	var dataAO1 struct {
		Pattern string `json:"pattern,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Pattern = dataAO1.Pattern

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ExcludedFileInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ExcludedFileBasics)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Pattern string `json:"pattern,omitempty"`
	}

	dataAO1.Pattern = m.Pattern

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this excluded file input
func (m *ExcludedFileInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ExcludedFileBasics
	if err := m.ExcludedFileBasics.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this excluded file input based on the context it is used
func (m *ExcludedFileInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ExcludedFileBasics
	if err := m.ExcludedFileBasics.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExcludedFileInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExcludedFileInput) UnmarshalBinary(b []byte) error {
	var res ExcludedFileInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
