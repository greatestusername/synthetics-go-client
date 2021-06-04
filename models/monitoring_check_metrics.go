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

// MonitoringCheckMetrics A list of available metrics for the check
//
// swagger:model monitoring_check_metrics
type MonitoringCheckMetrics struct {

	// The unique ID for the check
	// Example: 1
	ID int32 `json:"id,omitempty"`

	// A list of metrics
	Metrics []*MonitoringCheckMetricsMetricsItems0 `json:"metrics"`
}

// Validate validates this monitoring check metrics
func (m *MonitoringCheckMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitoringCheckMetrics) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {
		if swag.IsZero(m.Metrics[i]) { // not required
			continue
		}

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this monitoring check metrics based on the context it is used
func (m *MonitoringCheckMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitoringCheckMetrics) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metrics); i++ {

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MonitoringCheckMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitoringCheckMetrics) UnmarshalBinary(b []byte) error {
	var res MonitoringCheckMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MonitoringCheckMetricsMetricsItems0 An available metric for the check
//
// swagger:model MonitoringCheckMetricsMetricsItems0
type MonitoringCheckMetricsMetricsItems0 struct {

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

	// links
	Links *MonitoringCheckMetricsMetricsItems0AO1Links `json:"links,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MonitoringCheckMetricsMetricsItems0) UnmarshalJSON(raw []byte) error {
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
		Links *MonitoringCheckMetricsMetricsItems0AO1Links `json:"links,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Links = dataAO1.Links

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MonitoringCheckMetricsMetricsItems0) MarshalJSON() ([]byte, error) {
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
		Links *MonitoringCheckMetricsMetricsItems0AO1Links `json:"links,omitempty"`
	}

	dataAO1.Links = m.Links

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this monitoring check metrics metrics items0
func (m *MonitoringCheckMetricsMetricsItems0) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var monitoringCheckMetricsMetricsItems0TypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["milliseconds","count","percent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		monitoringCheckMetricsMetricsItems0TypeFormatPropEnum = append(monitoringCheckMetricsMetricsItems0TypeFormatPropEnum, v)
	}
}

// property enum
func (m *MonitoringCheckMetricsMetricsItems0) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, monitoringCheckMetricsMetricsItems0TypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MonitoringCheckMetricsMetricsItems0) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

var monitoringCheckMetricsMetricsItems0TypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Mean Response Time","Maximum Response Time","Minimum Response Time","Response Time Standard Deviation","Run Count","Error Count","Percentage Uptime","Percentage Downtime","SLA Percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		monitoringCheckMetricsMetricsItems0TypeLabelPropEnum = append(monitoringCheckMetricsMetricsItems0TypeLabelPropEnum, v)
	}
}

// property enum
func (m *MonitoringCheckMetricsMetricsItems0) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, monitoringCheckMetricsMetricsItems0TypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MonitoringCheckMetricsMetricsItems0) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var monitoringCheckMetricsMetricsItems0TypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["average_response_time","max_response_time","min_response_time","standard_deviation","run_count","error_count","percentage_uptime","percentage_downtime","sla_percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		monitoringCheckMetricsMetricsItems0TypeNamePropEnum = append(monitoringCheckMetricsMetricsItems0TypeNamePropEnum, v)
	}
}

// property enum
func (m *MonitoringCheckMetricsMetricsItems0) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, monitoringCheckMetricsMetricsItems0TypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MonitoringCheckMetricsMetricsItems0) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MonitoringCheckMetricsMetricsItems0) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this monitoring check metrics metrics items0 based on the context it is used
func (m *MonitoringCheckMetricsMetricsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitoringCheckMetricsMetricsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MonitoringCheckMetricsMetricsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitoringCheckMetricsMetricsItems0) UnmarshalBinary(b []byte) error {
	var res MonitoringCheckMetricsMetricsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MonitoringCheckMetricsMetricsItems0AO1Links monitoring check metrics metrics items0 a o1 links
//
// swagger:model MonitoringCheckMetricsMetricsItems0AO1Links
type MonitoringCheckMetricsMetricsItems0AO1Links struct {

	// A link to the data for this metric
	// Example: https://monitoring-api.rigor.com/v2/checks/1/metrics?name=average_response_time
	Self string `json:"self,omitempty"`
}

// Validate validates this monitoring check metrics metrics items0 a o1 links
func (m *MonitoringCheckMetricsMetricsItems0AO1Links) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this monitoring check metrics metrics items0 a o1 links based on context it is used
func (m *MonitoringCheckMetricsMetricsItems0AO1Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MonitoringCheckMetricsMetricsItems0AO1Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitoringCheckMetricsMetricsItems0AO1Links) UnmarshalBinary(b []byte) error {
	var res MonitoringCheckMetricsMetricsItems0AO1Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
