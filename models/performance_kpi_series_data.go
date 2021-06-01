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

// PerformanceKpiSeriesData A single data point containing key-value pairs of metric names and their values
// Example: {"custom_timings":[{"entry_type":"mark","name":"myCustomMark","value":150.5},{"entry_type":"mark","name":"myCustomMark","value":321}],"first_request_server_ip":"104.28.17.124","requests":123,"run_count":1,"run_id":1,"start_render_ms":1000,"success_count":1,"time":"2021-05-25T17:54:05Z","url":"https://rigor.com"}
//
// swagger:model performance_kpi_series_data
type PerformanceKpiSeriesData struct {

	// custom timings
	CustomTimings []*CustomTiming `json:"custom_timings"`

	// The IP address of the first request for the page if grouping by `page` or the IP address of the first request for the first page of the run if not grouping
	// Example: 104.28.17.124
	FirstRequestServerIP string `json:"first_request_server_ip,omitempty"`

	// The number of runs aggregated in this data point
	// Example: 1
	// Required: true
	RunCount *int32 `json:"run_count"`

	// The unique ID for this run (if the interval provided was blank)
	RunID int32 `json:"run_id,omitempty"`

	// The number of successful runs aggregated in this data point
	// Example: 1
	// Required: true
	SuccessCount *int32 `json:"success_count"`

	// The timestamp for this run (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Required: true
	// Format: date-time
	Time *strfmt.DateTime `json:"time"`

	// The URL for the page if grouping by `page` or the URL for the first page of the run if not grouping
	// Example: https://rigor.com/
	URL string `json:"url,omitempty"`
}

// Validate validates this performance kpi series data
func (m *PerformanceKpiSeriesData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomTimings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccessCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceKpiSeriesData) validateCustomTimings(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomTimings) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomTimings); i++ {
		if swag.IsZero(m.CustomTimings[i]) { // not required
			continue
		}

		if m.CustomTimings[i] != nil {
			if err := m.CustomTimings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_timings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PerformanceKpiSeriesData) validateRunCount(formats strfmt.Registry) error {

	if err := validate.Required("run_count", "body", m.RunCount); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceKpiSeriesData) validateSuccessCount(formats strfmt.Registry) error {

	if err := validate.Required("success_count", "body", m.SuccessCount); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceKpiSeriesData) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance kpi series data based on the context it is used
func (m *PerformanceKpiSeriesData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomTimings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceKpiSeriesData) contextValidateCustomTimings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomTimings); i++ {

		if m.CustomTimings[i] != nil {
			if err := m.CustomTimings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_timings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceKpiSeriesData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceKpiSeriesData) UnmarshalBinary(b []byte) error {
	var res PerformanceKpiSeriesData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}