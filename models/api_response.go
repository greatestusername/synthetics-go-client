// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIResponse A success or error message
//
// swagger:model api_response
type APIResponse struct {

	// A collection of error messages
	Errors []*APIResponseErrorsItems0 `json:"errors"`

	// A high-level status message
	Message string `json:"message,omitempty"`

	// result
	// Enum: [success error]
	Result string `json:"result,omitempty"`
}

// Validate validates this api response
func (m *APIResponse) Validate(formats strfmt.Registry) error {
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

func (m *APIResponse) validateErrors(formats strfmt.Registry) error {
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

var apiResponseTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiResponseTypeResultPropEnum = append(apiResponseTypeResultPropEnum, v)
	}
}

const (

	// APIResponseResultSuccess captures enum value "success"
	APIResponseResultSuccess string = "success"

	// APIResponseResultError captures enum value "error"
	APIResponseResultError string = "error"
)

// prop value enum
func (m *APIResponse) validateResultEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apiResponseTypeResultPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *APIResponse) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this api response based on the context it is used
func (m *APIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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
func (m *APIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResponse) UnmarshalBinary(b []byte) error {
	var res APIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// APIResponseErrorsItems0 API response errors items0
//
// swagger:model APIResponseErrorsItems0
type APIResponseErrorsItems0 struct {

	// Additional detail about the error, if available
	Description string `json:"description,omitempty"`

	// A summary of the error
	Title string `json:"title,omitempty"`
}

// Validate validates this API response errors items0
func (m *APIResponseErrorsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this API response errors items0 based on context it is used
func (m *APIResponseErrorsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIResponseErrorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResponseErrorsItems0) UnmarshalBinary(b []byte) error {
	var res APIResponseErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
