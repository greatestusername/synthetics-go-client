// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// EnforceSslValidation When true, the check will fail if the browser encounters invalid security certificates.
//
// swagger:model enforce_ssl_validation
type EnforceSslValidation bool

// Validate validates this enforce ssl validation
func (m EnforceSslValidation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this enforce ssl validation based on context it is used
func (m EnforceSslValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
