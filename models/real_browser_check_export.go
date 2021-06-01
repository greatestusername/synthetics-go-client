// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RealBrowserCheckExport A Real Browser check converted to a Selenium SIDE script
//
// swagger:model real_browser_check_export
type RealBrowserCheckExport struct {

	// The endpoint to get the converted script
	// Example: /v2/checks/real_browsers/1/export.side
	Path string `json:"path,omitempty"`

	// A list of warnings about the exported script
	Warnings []string `json:"warnings"`
}

// Validate validates this real browser check export
func (m *RealBrowserCheckExport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this real browser check export based on context it is used
func (m *RealBrowserCheckExport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RealBrowserCheckExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RealBrowserCheckExport) UnmarshalBinary(b []byte) error {
	var res RealBrowserCheckExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}