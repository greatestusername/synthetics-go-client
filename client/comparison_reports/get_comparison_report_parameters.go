// Code generated by go-swagger; DO NOT EDIT.

package comparison_reports

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

// NewGetComparisonReportParams creates a new GetComparisonReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComparisonReportParams() *GetComparisonReportParams {
	return &GetComparisonReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComparisonReportParamsWithTimeout creates a new GetComparisonReportParams object
// with the ability to set a timeout on a request.
func NewGetComparisonReportParamsWithTimeout(timeout time.Duration) *GetComparisonReportParams {
	return &GetComparisonReportParams{
		timeout: timeout,
	}
}

// NewGetComparisonReportParamsWithContext creates a new GetComparisonReportParams object
// with the ability to set a context for a request.
func NewGetComparisonReportParamsWithContext(ctx context.Context) *GetComparisonReportParams {
	return &GetComparisonReportParams{
		Context: ctx,
	}
}

// NewGetComparisonReportParamsWithHTTPClient creates a new GetComparisonReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComparisonReportParamsWithHTTPClient(client *http.Client) *GetComparisonReportParams {
	return &GetComparisonReportParams{
		HTTPClient: client,
	}
}

/* GetComparisonReportParams contains all the parameters to send to the API endpoint
   for the get comparison report operation.

   Typically these are written to a http.Request.
*/
type GetComparisonReportParams struct {

	/* ID.

	   The unique ID of the Comparison Report to fetch

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get comparison report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComparisonReportParams) WithDefaults() *GetComparisonReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get comparison report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComparisonReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get comparison report params
func (o *GetComparisonReportParams) WithTimeout(timeout time.Duration) *GetComparisonReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get comparison report params
func (o *GetComparisonReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get comparison report params
func (o *GetComparisonReportParams) WithContext(ctx context.Context) *GetComparisonReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get comparison report params
func (o *GetComparisonReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get comparison report params
func (o *GetComparisonReportParams) WithHTTPClient(client *http.Client) *GetComparisonReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get comparison report params
func (o *GetComparisonReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get comparison report params
func (o *GetComparisonReportParams) WithID(id int32) *GetComparisonReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get comparison report params
func (o *GetComparisonReportParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetComparisonReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
