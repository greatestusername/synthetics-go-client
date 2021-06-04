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

// LocationMetrics The metric data for the location
//
// swagger:model location_metrics
type LocationMetrics struct {

	// The format of the data for this metric.
	// Example: percent
	// Enum: [milliseconds count percent]
	Format string `json:"format,omitempty"`

	// A readable label for the metric, in Title Case.
	// Example: Percentage Uptime
	// Enum: [Mean Response Time Maximum Response Time Minimum Response Time Response Time Standard Deviation Run Count Error Count Percentage Uptime Percentage Downtime SLA Percentage]
	Label string `json:"label,omitempty"`

	// The name of the metric, in snake_case.
	// Example: percentage_uptime
	// Enum: [average_response_time max_response_time min_response_time standard_deviation run_count error_count percentage_uptime percentage_downtime sla_percentage]
	Name string `json:"name,omitempty"`

	// An array of data points
	// Min Items: 0
	Points []*LocationMetricsPointsItems0 `json:"points"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LocationMetrics) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Format string `json:"format,omitempty"`

		Label string `json:"label,omitempty"`

		Name string `json:"name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Format = dataAO0.Format

	m.Label = dataAO0.Label

	m.Name = dataAO0.Name

	// AO1
	var dataAO1 struct {
		Points []*LocationMetricsPointsItems0 `json:"points"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Points = dataAO1.Points

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LocationMetrics) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Format string `json:"format,omitempty"`

		Label string `json:"label,omitempty"`

		Name string `json:"name,omitempty"`
	}

	dataAO0.Format = m.Format

	dataAO0.Label = m.Label

	dataAO0.Name = m.Name

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		Points []*LocationMetricsPointsItems0 `json:"points"`
	}

	dataAO1.Points = m.Points

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this location metrics
func (m *LocationMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var locationMetricsTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["milliseconds","count","percent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationMetricsTypeFormatPropEnum = append(locationMetricsTypeFormatPropEnum, v)
	}
}

// property enum
func (m *LocationMetrics) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, locationMetricsTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocationMetrics) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

var locationMetricsTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Mean Response Time","Maximum Response Time","Minimum Response Time","Response Time Standard Deviation","Run Count","Error Count","Percentage Uptime","Percentage Downtime","SLA Percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationMetricsTypeLabelPropEnum = append(locationMetricsTypeLabelPropEnum, v)
	}
}

// property enum
func (m *LocationMetrics) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, locationMetricsTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocationMetrics) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var locationMetricsTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["average_response_time","max_response_time","min_response_time","standard_deviation","run_count","error_count","percentage_uptime","percentage_downtime","sla_percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationMetricsTypeNamePropEnum = append(locationMetricsTypeNamePropEnum, v)
	}
}

// property enum
func (m *LocationMetrics) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, locationMetricsTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocationMetrics) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LocationMetrics) validatePoints(formats strfmt.Registry) error {

	if swag.IsZero(m.Points) { // not required
		return nil
	}

	iPointsSize := int64(len(m.Points))

	if err := validate.MinItems("points", "body", iPointsSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(m.Points); i++ {
		if swag.IsZero(m.Points[i]) { // not required
			continue
		}

		if m.Points[i] != nil {
			if err := m.Points[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this location metrics based on the context it is used
func (m *LocationMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationMetrics) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Points); i++ {

		if m.Points[i] != nil {
			if err := m.Points[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocationMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationMetrics) UnmarshalBinary(b []byte) error {
	var res LocationMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LocationMetricsPointsItems0 A single data point. May be a single run or the mean across multiple runs.
//
// swagger:model LocationMetricsPointsItems0
type LocationMetricsPointsItems0 struct {

	// The start timestamp for the data point (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	From strfmt.DateTime `json:"from,omitempty"`

	// The end timestamp for the data point (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty"`

	// The value for the data point. May be run-level or aggregate data.
	// Example: 99.3
	Value float64 `json:"value,omitempty"`
}

// Validate validates this location metrics points items0
func (m *LocationMetricsPointsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationMetricsPointsItems0) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.From) { // not required
		return nil
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LocationMetricsPointsItems0) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this location metrics points items0 based on context it is used
func (m *LocationMetricsPointsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocationMetricsPointsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationMetricsPointsItems0) UnmarshalBinary(b []byte) error {
	var res LocationMetricsPointsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
