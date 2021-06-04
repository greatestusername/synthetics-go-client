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

// ExcludedFile excluded file
//
// swagger:model excluded_file
type ExcludedFile struct {

	// The type of exclusion.
	//               "preset": exclude a preset URL (default).
	//               "custom": exclude a custom URL.
	//               "all_except": whitelist a custom URL. Whitelisted URLs override all others.
	// Enum: [preset custom all_except]
	ExclusionType *string `json:"exclusion_type,omitempty"`

	// The name of the excluded preset URL. Null if the `exclusion_type` is not "preset".
	// Enum: [chartbeat clicktale comscore coremetrics crazy-egg eloqua gomez google-analytics hubspot liveperson mixpanel omniture optimizely pardot quantcast spectate tealium white-ops]
	PresetName string `json:"preset_name,omitempty"`

	// When the excluded file was created (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// When the excluded file was last updated (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// A regular expression to match against all URLs visited during the check run
	URL string `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ExcludedFile) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ExclusionType *string `json:"exclusion_type,omitempty"`

		PresetName string `json:"preset_name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ExclusionType = dataAO0.ExclusionType

	m.PresetName = dataAO0.PresetName

	// AO1
	var dataAO1 struct {
		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

		URL string `json:"url,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CreatedAt = dataAO1.CreatedAt

	m.UpdatedAt = dataAO1.UpdatedAt

	m.URL = dataAO1.URL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ExcludedFile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ExclusionType *string `json:"exclusion_type,omitempty"`

		PresetName string `json:"preset_name,omitempty"`
	}

	dataAO0.ExclusionType = m.ExclusionType

	dataAO0.PresetName = m.PresetName

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

		UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

		URL string `json:"url,omitempty"`
	}

	dataAO1.CreatedAt = m.CreatedAt

	dataAO1.UpdatedAt = m.UpdatedAt

	dataAO1.URL = m.URL

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this excluded file
func (m *ExcludedFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresetName(formats); err != nil {
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

var excludedFileTypeExclusionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["preset","custom","all_except"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		excludedFileTypeExclusionTypePropEnum = append(excludedFileTypeExclusionTypePropEnum, v)
	}
}

// property enum
func (m *ExcludedFile) validateExclusionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, excludedFileTypeExclusionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExcludedFile) validateExclusionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ExclusionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExclusionTypeEnum("exclusion_type", "body", *m.ExclusionType); err != nil {
		return err
	}

	return nil
}

var excludedFileTypePresetNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chartbeat","clicktale","comscore","coremetrics","crazy-egg","eloqua","gomez","google-analytics","hubspot","liveperson","mixpanel","omniture","optimizely","pardot","quantcast","spectate","tealium","white-ops"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		excludedFileTypePresetNamePropEnum = append(excludedFileTypePresetNamePropEnum, v)
	}
}

// property enum
func (m *ExcludedFile) validatePresetNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, excludedFileTypePresetNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExcludedFile) validatePresetName(formats strfmt.Registry) error {

	if swag.IsZero(m.PresetName) { // not required
		return nil
	}

	// value enum
	if err := m.validatePresetNameEnum("preset_name", "body", m.PresetName); err != nil {
		return err
	}

	return nil
}

func (m *ExcludedFile) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExcludedFile) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this excluded file based on context it is used
func (m *ExcludedFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExcludedFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExcludedFile) UnmarshalBinary(b []byte) error {
	var res ExcludedFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
