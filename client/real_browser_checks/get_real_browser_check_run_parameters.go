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
	"github.com/go-openapi/swag"
)

// NewGetRealBrowserCheckRunParams creates a new GetRealBrowserCheckRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRealBrowserCheckRunParams() *GetRealBrowserCheckRunParams {
	return &GetRealBrowserCheckRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRealBrowserCheckRunParamsWithTimeout creates a new GetRealBrowserCheckRunParams object
// with the ability to set a timeout on a request.
func NewGetRealBrowserCheckRunParamsWithTimeout(timeout time.Duration) *GetRealBrowserCheckRunParams {
	return &GetRealBrowserCheckRunParams{
		timeout: timeout,
	}
}

// NewGetRealBrowserCheckRunParamsWithContext creates a new GetRealBrowserCheckRunParams object
// with the ability to set a context for a request.
func NewGetRealBrowserCheckRunParamsWithContext(ctx context.Context) *GetRealBrowserCheckRunParams {
	return &GetRealBrowserCheckRunParams{
		Context: ctx,
	}
}

// NewGetRealBrowserCheckRunParamsWithHTTPClient creates a new GetRealBrowserCheckRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRealBrowserCheckRunParamsWithHTTPClient(client *http.Client) *GetRealBrowserCheckRunParams {
	return &GetRealBrowserCheckRunParams{
		HTTPClient: client,
	}
}

/* GetRealBrowserCheckRunParams contains all the parameters to send to the API endpoint
   for the get real browser check run operation.

   Typically these are written to a http.Request.
*/
type GetRealBrowserCheckRunParams struct {

	/* CheckID.

	   ID of the Real Browser check to fetch

	   Format: int32
	*/
	CheckID int32

	/* ID.

	   ID of the run to fetch

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get real browser check run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRealBrowserCheckRunParams) WithDefaults() *GetRealBrowserCheckRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get real browser check run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRealBrowserCheckRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) WithTimeout(timeout time.Duration) *GetRealBrowserCheckRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) WithContext(ctx context.Context) *GetRealBrowserCheckRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) WithHTTPClient(client *http.Client) *GetRealBrowserCheckRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckID adds the checkID to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) WithCheckID(checkID int32) *GetRealBrowserCheckRunParams {
	o.SetCheckID(checkID)
	return o
}

// SetCheckID adds the checkId to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) SetCheckID(checkID int32) {
	o.CheckID = checkID
}

// WithID adds the id to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) WithID(id int32) *GetRealBrowserCheckRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get real browser check run params
func (o *GetRealBrowserCheckRunParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRealBrowserCheckRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param check_id
	if err := r.SetPathParam("check_id", swag.FormatInt32(o.CheckID)); err != nil {
		return err
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
