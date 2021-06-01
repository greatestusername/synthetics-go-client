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

// ExcludedFileBasics excluded file basics
//
// swagger:model excluded_file_basics
type ExcludedFileBasics struct {

	// The type of exclusion.
	//               "preset": exclude a preset URL (default).
	//               "custom": exclude a custom URL.
	//               "all_except": whitelist a custom URL. Whitelisted URLs override all others.
	// Enum: [preset custom all_except]
	ExclusionType *string `json:"exclusion_type,omitempty"`

	// The name of the excluded preset URL. Null if the `exclusion_type` is not "preset".
	// Enum: [chartbeat clicktale comscore coremetrics crazy-egg eloqua gomez google-analytics hubspot liveperson mixpanel omniture optimizely pardot quantcast spectate tealium white-ops]
	PresetName string `json:"preset_name,omitempty"`
}

// Validate validates this excluded file basics
func (m *ExcludedFileBasics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresetName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var excludedFileBasicsTypeExclusionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["preset","custom","all_except"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		excludedFileBasicsTypeExclusionTypePropEnum = append(excludedFileBasicsTypeExclusionTypePropEnum, v)
	}
}

const (

	// ExcludedFileBasicsExclusionTypePreset captures enum value "preset"
	ExcludedFileBasicsExclusionTypePreset string = "preset"

	// ExcludedFileBasicsExclusionTypeCustom captures enum value "custom"
	ExcludedFileBasicsExclusionTypeCustom string = "custom"

	// ExcludedFileBasicsExclusionTypeAllExcept captures enum value "all_except"
	ExcludedFileBasicsExclusionTypeAllExcept string = "all_except"
)

// prop value enum
func (m *ExcludedFileBasics) validateExclusionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, excludedFileBasicsTypeExclusionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExcludedFileBasics) validateExclusionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ExclusionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExclusionTypeEnum("exclusion_type", "body", *m.ExclusionType); err != nil {
		return err
	}

	return nil
}

var excludedFileBasicsTypePresetNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chartbeat","clicktale","comscore","coremetrics","crazy-egg","eloqua","gomez","google-analytics","hubspot","liveperson","mixpanel","omniture","optimizely","pardot","quantcast","spectate","tealium","white-ops"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		excludedFileBasicsTypePresetNamePropEnum = append(excludedFileBasicsTypePresetNamePropEnum, v)
	}
}

const (

	// ExcludedFileBasicsPresetNameChartbeat captures enum value "chartbeat"
	ExcludedFileBasicsPresetNameChartbeat string = "chartbeat"

	// ExcludedFileBasicsPresetNameClicktale captures enum value "clicktale"
	ExcludedFileBasicsPresetNameClicktale string = "clicktale"

	// ExcludedFileBasicsPresetNameComscore captures enum value "comscore"
	ExcludedFileBasicsPresetNameComscore string = "comscore"

	// ExcludedFileBasicsPresetNameCoremetrics captures enum value "coremetrics"
	ExcludedFileBasicsPresetNameCoremetrics string = "coremetrics"

	// ExcludedFileBasicsPresetNameCrazyDashEgg captures enum value "crazy-egg"
	ExcludedFileBasicsPresetNameCrazyDashEgg string = "crazy-egg"

	// ExcludedFileBasicsPresetNameEloqua captures enum value "eloqua"
	ExcludedFileBasicsPresetNameEloqua string = "eloqua"

	// ExcludedFileBasicsPresetNameGomez captures enum value "gomez"
	ExcludedFileBasicsPresetNameGomez string = "gomez"

	// ExcludedFileBasicsPresetNameGoogleDashAnalytics captures enum value "google-analytics"
	ExcludedFileBasicsPresetNameGoogleDashAnalytics string = "google-analytics"

	// ExcludedFileBasicsPresetNameHubspot captures enum value "hubspot"
	ExcludedFileBasicsPresetNameHubspot string = "hubspot"

	// ExcludedFileBasicsPresetNameLiveperson captures enum value "liveperson"
	ExcludedFileBasicsPresetNameLiveperson string = "liveperson"

	// ExcludedFileBasicsPresetNameMixpanel captures enum value "mixpanel"
	ExcludedFileBasicsPresetNameMixpanel string = "mixpanel"

	// ExcludedFileBasicsPresetNameOmniture captures enum value "omniture"
	ExcludedFileBasicsPresetNameOmniture string = "omniture"

	// ExcludedFileBasicsPresetNameOptimizely captures enum value "optimizely"
	ExcludedFileBasicsPresetNameOptimizely string = "optimizely"

	// ExcludedFileBasicsPresetNamePardot captures enum value "pardot"
	ExcludedFileBasicsPresetNamePardot string = "pardot"

	// ExcludedFileBasicsPresetNameQuantcast captures enum value "quantcast"
	ExcludedFileBasicsPresetNameQuantcast string = "quantcast"

	// ExcludedFileBasicsPresetNameSpectate captures enum value "spectate"
	ExcludedFileBasicsPresetNameSpectate string = "spectate"

	// ExcludedFileBasicsPresetNameTealium captures enum value "tealium"
	ExcludedFileBasicsPresetNameTealium string = "tealium"

	// ExcludedFileBasicsPresetNameWhiteDashOps captures enum value "white-ops"
	ExcludedFileBasicsPresetNameWhiteDashOps string = "white-ops"
)

// prop value enum
func (m *ExcludedFileBasics) validatePresetNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, excludedFileBasicsTypePresetNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExcludedFileBasics) validatePresetName(formats strfmt.Registry) error {
	if swag.IsZero(m.PresetName) { // not required
		return nil
	}

	// value enum
	if err := m.validatePresetNameEnum("preset_name", "body", m.PresetName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this excluded file basics based on context it is used
func (m *ExcludedFileBasics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExcludedFileBasics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExcludedFileBasics) UnmarshalBinary(b []byte) error {
	var res ExcludedFileBasics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
