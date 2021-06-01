// Code generated by go-swagger; DO NOT EDIT.

package http_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/greatestusername/synthetics-go-client/models"
)

// NewUpdateHTTPCheckParams creates a new UpdateHTTPCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHTTPCheckParams() *UpdateHTTPCheckParams {
	return &UpdateHTTPCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHTTPCheckParamsWithTimeout creates a new UpdateHTTPCheckParams object
// with the ability to set a timeout on a request.
func NewUpdateHTTPCheckParamsWithTimeout(timeout time.Duration) *UpdateHTTPCheckParams {
	return &UpdateHTTPCheckParams{
		timeout: timeout,
	}
}

// NewUpdateHTTPCheckParamsWithContext creates a new UpdateHTTPCheckParams object
// with the ability to set a context for a request.
func NewUpdateHTTPCheckParamsWithContext(ctx context.Context) *UpdateHTTPCheckParams {
	return &UpdateHTTPCheckParams{
		Context: ctx,
	}
}

// NewUpdateHTTPCheckParamsWithHTTPClient creates a new UpdateHTTPCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHTTPCheckParamsWithHTTPClient(client *http.Client) *UpdateHTTPCheckParams {
	return &UpdateHTTPCheckParams{
		HTTPClient: client,
	}
}

/* UpdateHTTPCheckParams contains all the parameters to send to the API endpoint
   for the update Http check operation.

   Typically these are written to a http.Request.
*/
type UpdateHTTPCheckParams struct {

	// CheckDetail.
	CheckDetail *models.HTTPCheckInput

	/* ID.

	   The unique ID of the HTTP check to update

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Http check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHTTPCheckParams) WithDefaults() *UpdateHTTPCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Http check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHTTPCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Http check params
func (o *UpdateHTTPCheckParams) WithTimeout(timeout time.Duration) *UpdateHTTPCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Http check params
func (o *UpdateHTTPCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Http check params
func (o *UpdateHTTPCheckParams) WithContext(ctx context.Context) *UpdateHTTPCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Http check params
func (o *UpdateHTTPCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Http check params
func (o *UpdateHTTPCheckParams) WithHTTPClient(client *http.Client) *UpdateHTTPCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Http check params
func (o *UpdateHTTPCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckDetail adds the checkDetail to the update Http check params
func (o *UpdateHTTPCheckParams) WithCheckDetail(checkDetail *models.HTTPCheckInput) *UpdateHTTPCheckParams {
	o.SetCheckDetail(checkDetail)
	return o
}

// SetCheckDetail adds the checkDetail to the update Http check params
func (o *UpdateHTTPCheckParams) SetCheckDetail(checkDetail *models.HTTPCheckInput) {
	o.CheckDetail = checkDetail
}

// WithID adds the id to the update Http check params
func (o *UpdateHTTPCheckParams) WithID(id int32) *UpdateHTTPCheckParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update Http check params
func (o *UpdateHTTPCheckParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHTTPCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CheckDetail != nil {
		if err := r.SetBodyParam(o.CheckDetail); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
