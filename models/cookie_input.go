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

// CookieInput An array of cookies to add to the check. Set to an empty array (`[]`) to clear.
//
// swagger:model cookie_input
type CookieInput struct {

	// The domain of the requests to add the cookie to. When no domain is provided,
	//               the domain and all subdomains of the starting URL are used.
	// Example: .example.com
	Domain string `json:"domain,omitempty"`

	// The name of the cookie
	// Example: cookie-name
	// Required: true
	Key *string `json:"key"`

	// The path of the requests to add the cookie to. When no path is provided,
	//               the root path ("/") is used.
	Path *string `json:"path,omitempty"`

	// The value of the cookie
	// Example: cookie-value
	// Required: true
	Value *string `json:"value"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CookieInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Domain string `json:"domain,omitempty"`

		Key *string `json:"key"`

		Path *string `json:"path,omitempty"`

		Value *string `json:"value"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Domain = dataAO0.Domain

	m.Key = dataAO0.Key

	m.Path = dataAO0.Path

	m.Value = dataAO0.Value

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CookieInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Domain string `json:"domain,omitempty"`

		Key *string `json:"key"`

		Path *string `json:"path,omitempty"`

		Value *string `json:"value"`
	}

	dataAO0.Domain = m.Domain

	dataAO0.Key = m.Key

	dataAO0.Path = m.Path

	dataAO0.Value = m.Value

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this cookie input
func (m *CookieInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
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

func (m *CookieInput) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *CookieInput) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cookie input based on context it is used
func (m *CookieInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CookieInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CookieInput) UnmarshalBinary(b []byte) error {
	var res CookieInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
