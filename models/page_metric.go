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

// PageMetric page metric
//
// swagger:model page_metric
type PageMetric struct {

	// A short description of the metric.
	// Example: Time until document has fully loaded and parsed the HTML document
	Description string `json:"description,omitempty"`

	// A titleized label for the metric.
	// Example: DOM Load Time
	// Required: true
	Label *string `json:"label"`

	// The metric name. If rollup data is returned,
	//              the metric will be prefixed by `median_`, `average_`, `max_`, or `min_`.
	// Example: dom_load_time
	// Required: true
	// Enum: [server_time dom_load_time start_render onload_time visually_complete fully_loaded_time speed_index request_count content_size html_count html_size image_count image_size javascript_count javascript_size css_count css_size video_count video_size font_count font_size other_count other_size client_error_count connection_error_count server_error_count error_count]
	Name *string `json:"name"`

	// The units for the metric value.
	// Example: milliseconds
	// Required: true
	// Enum: [milliseconds count bytes]
	Unit *string `json:"unit"`
}

// Validate validates this page metric
func (m *PageMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageMetric) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var pageMetricTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["server_time","dom_load_time","start_render","onload_time","visually_complete","fully_loaded_time","speed_index","request_count","content_size","html_count","html_size","image_count","image_size","javascript_count","javascript_size","css_count","css_size","video_count","video_size","font_count","font_size","other_count","other_size","client_error_count","connection_error_count","server_error_count","error_count"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pageMetricTypeNamePropEnum = append(pageMetricTypeNamePropEnum, v)
	}
}

const (

	// PageMetricNameServerTime captures enum value "server_time"
	PageMetricNameServerTime string = "server_time"

	// PageMetricNameDomLoadTime captures enum value "dom_load_time"
	PageMetricNameDomLoadTime string = "dom_load_time"

	// PageMetricNameStartRender captures enum value "start_render"
	PageMetricNameStartRender string = "start_render"

	// PageMetricNameOnloadTime captures enum value "onload_time"
	PageMetricNameOnloadTime string = "onload_time"

	// PageMetricNameVisuallyComplete captures enum value "visually_complete"
	PageMetricNameVisuallyComplete string = "visually_complete"

	// PageMetricNameFullyLoadedTime captures enum value "fully_loaded_time"
	PageMetricNameFullyLoadedTime string = "fully_loaded_time"

	// PageMetricNameSpeedIndex captures enum value "speed_index"
	PageMetricNameSpeedIndex string = "speed_index"

	// PageMetricNameRequestCount captures enum value "request_count"
	PageMetricNameRequestCount string = "request_count"

	// PageMetricNameContentSize captures enum value "content_size"
	PageMetricNameContentSize string = "content_size"

	// PageMetricNameHTMLCount captures enum value "html_count"
	PageMetricNameHTMLCount string = "html_count"

	// PageMetricNameHTMLSize captures enum value "html_size"
	PageMetricNameHTMLSize string = "html_size"

	// PageMetricNameImageCount captures enum value "image_count"
	PageMetricNameImageCount string = "image_count"

	// PageMetricNameImageSize captures enum value "image_size"
	PageMetricNameImageSize string = "image_size"

	// PageMetricNameJavascriptCount captures enum value "javascript_count"
	PageMetricNameJavascriptCount string = "javascript_count"

	// PageMetricNameJavascriptSize captures enum value "javascript_size"
	PageMetricNameJavascriptSize string = "javascript_size"

	// PageMetricNameCSSCount captures enum value "css_count"
	PageMetricNameCSSCount string = "css_count"

	// PageMetricNameCSSSize captures enum value "css_size"
	PageMetricNameCSSSize string = "css_size"

	// PageMetricNameVideoCount captures enum value "video_count"
	PageMetricNameVideoCount string = "video_count"

	// PageMetricNameVideoSize captures enum value "video_size"
	PageMetricNameVideoSize string = "video_size"

	// PageMetricNameFontCount captures enum value "font_count"
	PageMetricNameFontCount string = "font_count"

	// PageMetricNameFontSize captures enum value "font_size"
	PageMetricNameFontSize string = "font_size"

	// PageMetricNameOtherCount captures enum value "other_count"
	PageMetricNameOtherCount string = "other_count"

	// PageMetricNameOtherSize captures enum value "other_size"
	PageMetricNameOtherSize string = "other_size"

	// PageMetricNameClientErrorCount captures enum value "client_error_count"
	PageMetricNameClientErrorCount string = "client_error_count"

	// PageMetricNameConnectionErrorCount captures enum value "connection_error_count"
	PageMetricNameConnectionErrorCount string = "connection_error_count"

	// PageMetricNameServerErrorCount captures enum value "server_error_count"
	PageMetricNameServerErrorCount string = "server_error_count"

	// PageMetricNameErrorCount captures enum value "error_count"
	PageMetricNameErrorCount string = "error_count"
)

// prop value enum
func (m *PageMetric) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pageMetricTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PageMetric) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

var pageMetricTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["milliseconds","count","bytes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pageMetricTypeUnitPropEnum = append(pageMetricTypeUnitPropEnum, v)
	}
}

const (

	// PageMetricUnitMilliseconds captures enum value "milliseconds"
	PageMetricUnitMilliseconds string = "milliseconds"

	// PageMetricUnitCount captures enum value "count"
	PageMetricUnitCount string = "count"

	// PageMetricUnitBytes captures enum value "bytes"
	PageMetricUnitBytes string = "bytes"
)

// prop value enum
func (m *PageMetric) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pageMetricTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PageMetric) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this page metric based on context it is used
func (m *PageMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PageMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageMetric) UnmarshalBinary(b []byte) error {
	var res PageMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}