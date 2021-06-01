// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Page page
//
// swagger:model page
type Page struct {

	// The hostname of the page.
	// Example: example.com
	Hostname string `json:"hostname,omitempty"`

	// The unique ID for the page.
	// Example: 1
	ID int32 `json:"id,omitempty"`

	// The full URL of the page.
	// Example: http://example.com
	URL string `json:"url,omitempty"`
}

// Validate validates this page
func (m *Page) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this page based on context it is used
func (m *Page) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Page) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Page) UnmarshalBinary(b []byte) error {
	var res Page
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
