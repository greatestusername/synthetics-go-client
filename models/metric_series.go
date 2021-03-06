// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MetricSeries The data for the metric
//
// swagger:model metric_series
type MetricSeries struct {
	MetricName

	// An array of data points
	// Min Items: 0
	Points []*PointWithLocations `json:"points"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MetricSeries) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MetricName
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MetricName = aO0

	// AO1
	var dataAO1 struct {
		Points []*PointWithLocations `json:"points"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Points = dataAO1.Points

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MetricSeries) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MetricName)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Points []*PointWithLocations `json:"points"`
	}

	dataAO1.Points = m.Points

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this metric series
func (m *MetricSeries) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetricName
	if err := m.MetricName.Validate(formats); err != nil {
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

func (m *MetricSeries) validatePoints(formats strfmt.Registry) error {

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

// ContextValidate validate this metric series based on the context it is used
func (m *MetricSeries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetricName
	if err := m.MetricName.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricSeries) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MetricSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricSeries) UnmarshalBinary(b []byte) error {
	var res MetricSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
