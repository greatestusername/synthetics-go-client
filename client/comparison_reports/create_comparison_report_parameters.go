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

	"github.com/greatestusername/synthetics-go-client/models"
)

// NewCreateComparisonReportParams creates a new CreateComparisonReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateComparisonReportParams() *CreateComparisonReportParams {
	return &CreateComparisonReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateComparisonReportParamsWithTimeout creates a new CreateComparisonReportParams object
// with the ability to set a timeout on a request.
func NewCreateComparisonReportParamsWithTimeout(timeout time.Duration) *CreateComparisonReportParams {
	return &CreateComparisonReportParams{
		timeout: timeout,
	}
}

// NewCreateComparisonReportParamsWithContext creates a new CreateComparisonReportParams object
// with the ability to set a context for a request.
func NewCreateComparisonReportParamsWithContext(ctx context.Context) *CreateComparisonReportParams {
	return &CreateComparisonReportParams{
		Context: ctx,
	}
}

// NewCreateComparisonReportParamsWithHTTPClient creates a new CreateComparisonReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateComparisonReportParamsWithHTTPClient(client *http.Client) *CreateComparisonReportParams {
	return &CreateComparisonReportParams{
		HTTPClient: client,
	}
}

/* CreateComparisonReportParams contains all the parameters to send to the API endpoint
   for the create comparison report operation.

   Typically these are written to a http.Request.
*/
type CreateComparisonReportParams struct {

	/* ComparisonReportDetail.

	   The configuration for the new Comparison Report
	*/
	ComparisonReportDetail *models.ComparisonReportInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create comparison report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateComparisonReportParams) WithDefaults() *CreateComparisonReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create comparison report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateComparisonReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create comparison report params
func (o *CreateComparisonReportParams) WithTimeout(timeout time.Duration) *CreateComparisonReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create comparison report params
func (o *CreateComparisonReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create comparison report params
func (o *CreateComparisonReportParams) WithContext(ctx context.Context) *CreateComparisonReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create comparison report params
func (o *CreateComparisonReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create comparison report params
func (o *CreateComparisonReportParams) WithHTTPClient(client *http.Client) *CreateComparisonReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create comparison report params
func (o *CreateComparisonReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComparisonReportDetail adds the comparisonReportDetail to the create comparison report params
func (o *CreateComparisonReportParams) WithComparisonReportDetail(comparisonReportDetail *models.ComparisonReportInput) *CreateComparisonReportParams {
	o.SetComparisonReportDetail(comparisonReportDetail)
	return o
}

// SetComparisonReportDetail adds the comparisonReportDetail to the create comparison report params
func (o *CreateComparisonReportParams) SetComparisonReportDetail(comparisonReportDetail *models.ComparisonReportInput) {
	o.ComparisonReportDetail = comparisonReportDetail
}

// WriteToRequest writes these params to a swagger request
func (o *CreateComparisonReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ComparisonReportDetail != nil {
		if err := r.SetBodyParam(o.ComparisonReportDetail); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
