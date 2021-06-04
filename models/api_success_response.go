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

// APISuccessResponse A success message
//
// swagger:model api_success_response
type APISuccessResponse struct {

	// A collection of error messages
	Errors []*APISuccessResponseErrorsItems0 `json:"errors"`

	// A high-level status message
	Message string `json:"message,omitempty"`

	// result
	// Enum: [success error]
	Result string `json:"result,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *APISuccessResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Errors []*APISuccessResponseErrorsItems0 `json:"errors"`

		Message string `json:"message,omitempty"`

		Result string `json:"result,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Errors = dataAO0.Errors

	m.Message = dataAO0.Message

	m.Result = dataAO0.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m APISuccessResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Errors []*APISuccessResponseErrorsItems0 `json:"errors"`

		Message string `json:"message,omitempty"`

		Result string `json:"result,omitempty"`
	}

	dataAO0.Errors = m.Errors

	dataAO0.Message = m.Message

	dataAO0.Result = m.Result

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this api success response
func (m *APISuccessResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISuccessResponse) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var apiSuccessResponseTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiSuccessResponseTypeResultPropEnum = append(apiSuccessResponseTypeResultPropEnum, v)
	}
}

// property enum
func (m *APISuccessResponse) validateResultEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiSuccessResponseTypeResultPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APISuccessResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api success response based on the context it is used
func (m *APISuccessResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISuccessResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APISuccessResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APISuccessResponse) UnmarshalBinary(b []byte) error {
	var res APISuccessResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// APISuccessResponseErrorsItems0 API success response errors items0
//
// swagger:model APISuccessResponseErrorsItems0
type APISuccessResponseErrorsItems0 struct {

	// Additional detail about the error, if available
	Description string `json:"description,omitempty"`

	// A summary of the error
	Title string `json:"title,omitempty"`
}

// Validate validates this API success response errors items0
func (m *APISuccessResponseErrorsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this API success response errors items0 based on context it is used
func (m *APISuccessResponseErrorsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APISuccessResponseErrorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APISuccessResponseErrorsItems0) UnmarshalBinary(b []byte) error {
	var res APISuccessResponseErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
