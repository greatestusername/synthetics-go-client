// Code generated by go-swagger; DO NOT EDIT.

package real_browser_checks

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
)

// NewSeleniumSideExportParams creates a new SeleniumSideExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSeleniumSideExportParams() *SeleniumSideExportParams {
	return &SeleniumSideExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSeleniumSideExportParamsWithTimeout creates a new SeleniumSideExportParams object
// with the ability to set a timeout on a request.
func NewSeleniumSideExportParamsWithTimeout(timeout time.Duration) *SeleniumSideExportParams {
	return &SeleniumSideExportParams{
		timeout: timeout,
	}
}

// NewSeleniumSideExportParamsWithContext creates a new SeleniumSideExportParams object
// with the ability to set a context for a request.
func NewSeleniumSideExportParamsWithContext(ctx context.Context) *SeleniumSideExportParams {
	return &SeleniumSideExportParams{
		Context: ctx,
	}
}

// NewSeleniumSideExportParamsWithHTTPClient creates a new SeleniumSideExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewSeleniumSideExportParamsWithHTTPClient(client *http.Client) *SeleniumSideExportParams {
	return &SeleniumSideExportParams{
		HTTPClient: client,
	}
}

/* SeleniumSideExportParams contains all the parameters to send to the API endpoint
   for the selenium side export operation.

   Typically these are written to a http.Request.
*/
type SeleniumSideExportParams struct {

	/* ID.

	   ID of the Real Browser check to export
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the selenium side export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeleniumSideExportParams) WithDefaults() *SeleniumSideExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the selenium side export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeleniumSideExportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the selenium side export params
func (o *SeleniumSideExportParams) WithTimeout(timeout time.Duration) *SeleniumSideExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the selenium side export params
func (o *SeleniumSideExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the selenium side export params
func (o *SeleniumSideExportParams) WithContext(ctx context.Context) *SeleniumSideExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the selenium side export params
func (o *SeleniumSideExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the selenium side export params
func (o *SeleniumSideExportParams) WithHTTPClient(client *http.Client) *SeleniumSideExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the selenium side export params
func (o *SeleniumSideExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the selenium side export params
func (o *SeleniumSideExportParams) WithID(id string) *SeleniumSideExportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the selenium side export params
func (o *SeleniumSideExportParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SeleniumSideExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
