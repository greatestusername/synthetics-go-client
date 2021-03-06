// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SuccessCriterion A criterion for the check to pass
//
// swagger:model success_criterion
type SuccessCriterion struct {

	// action type
	// Required: true
	// Enum: [presence_of_text absence_of_text matches_regular_expression does_not_contain_regular_expression response_code goes_to_url]
	ActionType *string `json:"action_type"`

	// comparison string
	// Required: true
	ComparisonString *string `json:"comparison_string"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this success criterion
func (m *SuccessCriterion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComparisonString(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var successCriterionTypeActionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["presence_of_text","absence_of_text","matches_regular_expression","does_not_contain_regular_expression","response_code","goes_to_url"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		successCriterionTypeActionTypePropEnum = append(successCriterionTypeActionTypePropEnum, v)
	}
}

const (

	// SuccessCriterionActionTypePresenceOfText captures enum value "presence_of_text"
	SuccessCriterionActionTypePresenceOfText string = "presence_of_text"

	// SuccessCriterionActionTypeAbsenceOfText captures enum value "absence_of_text"
	SuccessCriterionActionTypeAbsenceOfText string = "absence_of_text"

	// SuccessCriterionActionTypeMatchesRegularExpression captures enum value "matches_regular_expression"
	SuccessCriterionActionTypeMatchesRegularExpression string = "matches_regular_expression"

	// SuccessCriterionActionTypeDoesNotContainRegularExpression captures enum value "does_not_contain_regular_expression"
	SuccessCriterionActionTypeDoesNotContainRegularExpression string = "does_not_contain_regular_expression"

	// SuccessCriterionActionTypeResponseCode captures enum value "response_code"
	SuccessCriterionActionTypeResponseCode string = "response_code"

	// SuccessCriterionActionTypeGoesToURL captures enum value "goes_to_url"
	SuccessCriterionActionTypeGoesToURL string = "goes_to_url"
)

// prop value enum
func (m *SuccessCriterion) validateActionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, successCriterionTypeActionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SuccessCriterion) validateActionType(formats strfmt.Registry) error {

	if err := validate.Required("action_type", "body", m.ActionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionTypeEnum("action_type", "body", *m.ActionType); err != nil {
		return err
	}

	return nil
}

func (m *SuccessCriterion) validateComparisonString(formats strfmt.Registry) error {

	if err := validate.Required("comparison_string", "body", m.ComparisonString); err != nil {
		return err
	}

	return nil
}

func (m *SuccessCriterion) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SuccessCriterion) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this success criterion based on context it is used
func (m *SuccessCriterion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuccessCriterion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessCriterion) UnmarshalBinary(b []byte) error {
	var res SuccessCriterion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
