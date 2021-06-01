// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPCheck An HTTP check
//
// swagger:model http_check
type HTTPCheck struct {
	UptimeCheck

	// Network connection settings to simulate different types of networks.
	Connection *Connection `json:"connection,omitempty"`

	// http method
	// Enum: [GET HEAD POST POST/XML PUT DELETE]
	HTTPMethod *string `json:"http_method,omitempty"`

	// Request body to send (e.g. post data).
	HTTPRequestBody string `json:"http_request_body,omitempty"`

	// success criteria
	SuccessCriteria []*SuccessCriterion `json:"success_criteria"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HTTPCheck) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UptimeCheck
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UptimeCheck = aO0

	// AO1
	var dataAO1 struct {
		Connection *Connection `json:"connection,omitempty"`

		HTTPMethod *string `json:"http_method,omitempty"`

		HTTPRequestBody string `json:"http_request_body,omitempty"`

		SuccessCriteria []*SuccessCriterion `json:"success_criteria"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Connection = dataAO1.Connection

	m.HTTPMethod = dataAO1.HTTPMethod

	m.HTTPRequestBody = dataAO1.HTTPRequestBody

	m.SuccessCriteria = dataAO1.SuccessCriteria

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HTTPCheck) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.UptimeCheck)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Connection *Connection `json:"connection,omitempty"`

		HTTPMethod *string `json:"http_method,omitempty"`

		HTTPRequestBody string `json:"http_request_body,omitempty"`

		SuccessCriteria []*SuccessCriterion `json:"success_criteria"`
	}

	dataAO1.Connection = m.Connection

	dataAO1.HTTPMethod = m.HTTPMethod

	dataAO1.HTTPRequestBody = m.HTTPRequestBody

	dataAO1.SuccessCriteria = m.SuccessCriteria

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this http check
func (m *HTTPCheck) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UptimeCheck
	if err := m.UptimeCheck.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccessCriteria(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPCheck) validateConnection(formats strfmt.Registry) error {

	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

var httpCheckTypeHTTPMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GET","HEAD","POST","POST/XML","PUT","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpCheckTypeHTTPMethodPropEnum = append(httpCheckTypeHTTPMethodPropEnum, v)
	}
}

// property enum
func (m *HTTPCheck) validateHTTPMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpCheckTypeHTTPMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HTTPCheck) validateHTTPMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPMethodEnum("http_method", "body", *m.HTTPMethod); err != nil {
		return err
	}

	return nil
}

func (m *HTTPCheck) validateSuccessCriteria(formats strfmt.Registry) error {

	if swag.IsZero(m.SuccessCriteria) { // not required
		return nil
	}

	for i := 0; i < len(m.SuccessCriteria); i++ {
		if swag.IsZero(m.SuccessCriteria[i]) { // not required
			continue
		}

		if m.SuccessCriteria[i] != nil {
			if err := m.SuccessCriteria[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("success_criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this http check based on the context it is used
func (m *HTTPCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UptimeCheck
	if err := m.UptimeCheck.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccessCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPCheck) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {
		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *HTTPCheck) contextValidateSuccessCriteria(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SuccessCriteria); i++ {

		if m.SuccessCriteria[i] != nil {
			if err := m.SuccessCriteria[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("success_criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPCheck) UnmarshalBinary(b []byte) error {
	var res HTTPCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
