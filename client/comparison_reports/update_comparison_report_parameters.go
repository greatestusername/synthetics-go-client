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

	"github.com/greatestusername/synthetics-go-client/models"
)

// NewUpdateComparisonReportParams creates a new UpdateComparisonReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateComparisonReportParams() *UpdateComparisonReportParams {
	return &UpdateComparisonReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateComparisonReportParamsWithTimeout creates a new UpdateComparisonReportParams object
// with the ability to set a timeout on a request.
func NewUpdateComparisonReportParamsWithTimeout(timeout time.Duration) *UpdateComparisonReportParams {
	return &UpdateComparisonReportParams{
		timeout: timeout,
	}
}

// NewUpdateComparisonReportParamsWithContext creates a new UpdateComparisonReportParams object
// with the ability to set a context for a request.
func NewUpdateComparisonReportParamsWithContext(ctx context.Context) *UpdateComparisonReportParams {
	return &UpdateComparisonReportParams{
		Context: ctx,
	}
}

// NewUpdateComparisonReportParamsWithHTTPClient creates a new UpdateComparisonReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateComparisonReportParamsWithHTTPClient(client *http.Client) *UpdateComparisonReportParams {
	return &UpdateComparisonReportParams{
		HTTPClient: client,
	}
}

/* UpdateComparisonReportParams contains all the parameters to send to the API endpoint
   for the update comparison report operation.

   Typically these are written to a http.Request.
*/
type UpdateComparisonReportParams struct {

	/* ComparisonReportDetail.

	   The updated configuration for the Comparison Report
	*/
	ComparisonReportDetail *models.ComparisonReportInput

	/* ID.

	   The unique ID of the Comparison Report to update

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update comparison report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateComparisonReportParams) WithDefaults() *UpdateComparisonReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update comparison report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateComparisonReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update comparison report params
func (o *UpdateComparisonReportParams) WithTimeout(timeout time.Duration) *UpdateComparisonReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update comparison report params
func (o *UpdateComparisonReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update comparison report params
func (o *UpdateComparisonReportParams) WithContext(ctx context.Context) *UpdateComparisonReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update comparison report params
func (o *UpdateComparisonReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update comparison report params
func (o *UpdateComparisonReportParams) WithHTTPClient(client *http.Client) *UpdateComparisonReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update comparison report params
func (o *UpdateComparisonReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComparisonReportDetail adds the comparisonReportDetail to the update comparison report params
func (o *UpdateComparisonReportParams) WithComparisonReportDetail(comparisonReportDetail *models.ComparisonReportInput) *UpdateComparisonReportParams {
	o.SetComparisonReportDetail(comparisonReportDetail)
	return o
}

// SetComparisonReportDetail adds the comparisonReportDetail to the update comparison report params
func (o *UpdateComparisonReportParams) SetComparisonReportDetail(comparisonReportDetail *models.ComparisonReportInput) {
	o.ComparisonReportDetail = comparisonReportDetail
}

// WithID adds the id to the update comparison report params
func (o *UpdateComparisonReportParams) WithID(id int32) *UpdateComparisonReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update comparison report params
func (o *UpdateComparisonReportParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateComparisonReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ComparisonReportDetail != nil {
		if err := r.SetBodyParam(o.ComparisonReportDetail); err != nil {
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
