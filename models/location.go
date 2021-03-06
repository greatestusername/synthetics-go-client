// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Location location
// Example: {"id":1,"name":"N. Virginia","region_code":"na-us-virginia","world_region":"NA"}
//
// swagger:model location
type Location struct {

	// The unique ID for the location
	ID int32 `json:"id,omitempty"`

	// The name of the location
	Name string `json:"name,omitempty"`

	// A readable code representing the location
	RegionCode string `json:"region_code,omitempty"`

	// The region the location is in
	WorldRegion string `json:"world_region,omitempty"`
}

// Validate validates this location
func (m *Location) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this location based on context it is used
func (m *Location) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Location) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Location) UnmarshalBinary(b []byte) error {
	var res Location
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
