// Code generated by go-swagger; DO NOT EDIT.

package checks

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

// NewGetCheckLocationMetricDataParams creates a new GetCheckLocationMetricDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCheckLocationMetricDataParams() *GetCheckLocationMetricDataParams {
	return &GetCheckLocationMetricDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCheckLocationMetricDataParamsWithTimeout creates a new GetCheckLocationMetricDataParams object
// with the ability to set a timeout on a request.
func NewGetCheckLocationMetricDataParamsWithTimeout(timeout time.Duration) *GetCheckLocationMetricDataParams {
	return &GetCheckLocationMetricDataParams{
		timeout: timeout,
	}
}

// NewGetCheckLocationMetricDataParamsWithContext creates a new GetCheckLocationMetricDataParams object
// with the ability to set a context for a request.
func NewGetCheckLocationMetricDataParamsWithContext(ctx context.Context) *GetCheckLocationMetricDataParams {
	return &GetCheckLocationMetricDataParams{
		Context: ctx,
	}
}

// NewGetCheckLocationMetricDataParamsWithHTTPClient creates a new GetCheckLocationMetricDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCheckLocationMetricDataParamsWithHTTPClient(client *http.Client) *GetCheckLocationMetricDataParams {
	return &GetCheckLocationMetricDataParams{
		HTTPClient: client,
	}
}

/* GetCheckLocationMetricDataParams contains all the parameters to send to the API endpoint
   for the get check location metric data operation.

   Typically these are written to a http.Request.
*/
type GetCheckLocationMetricDataParams struct {

	/* From.

	   The start time for the timeframe

	   Format: date-time
	*/
	From *strfmt.DateTime

	/* ID.

	   ID of check to fetch

	   Format: int32
	*/
	ID int32

	/* IncludeSummary.

	   If true, include aggregate data for each location (defaults to false)
	*/
	IncludeSummary *bool

	/* Locations.

	   A list of location IDs used to filter the results.
	*/
	Locations []int32

	/* Names.

	   The metric names to get data for
	*/
	Names []string

	/* Range.

	   A predefined timeframe to be used instead of `from` and `to`. Defaults to the last hour.
	*/
	Range *string

	/* To.

	   The end time for the timeframe

	   Format: date-time
	*/
	To *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get check location metric data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCheckLocationMetricDataParams) WithDefaults() *GetCheckLocationMetricDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get check location metric data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCheckLocationMetricDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithTimeout(timeout time.Duration) *GetCheckLocationMetricDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithContext(ctx context.Context) *GetCheckLocationMetricDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithHTTPClient(client *http.Client) *GetCheckLocationMetricDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithFrom(from *strfmt.DateTime) *GetCheckLocationMetricDataParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetFrom(from *strfmt.DateTime) {
	o.From = from
}

// WithID adds the id to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithID(id int32) *GetCheckLocationMetricDataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetID(id int32) {
	o.ID = id
}

// WithIncludeSummary adds the includeSummary to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithIncludeSummary(includeSummary *bool) *GetCheckLocationMetricDataParams {
	o.SetIncludeSummary(includeSummary)
	return o
}

// SetIncludeSummary adds the includeSummary to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetIncludeSummary(includeSummary *bool) {
	o.IncludeSummary = includeSummary
}

// WithLocations adds the locations to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithLocations(locations []int32) *GetCheckLocationMetricDataParams {
	o.SetLocations(locations)
	return o
}

// SetLocations adds the locations to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetLocations(locations []int32) {
	o.Locations = locations
}

// WithNames adds the names to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithNames(names []string) *GetCheckLocationMetricDataParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetNames(names []string) {
	o.Names = names
}

// WithRange adds the rangeVar to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithRange(rangeVar *string) *GetCheckLocationMetricDataParams {
	o.SetRange(rangeVar)
	return o
}

// SetRange adds the range to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetRange(rangeVar *string) {
	o.Range = rangeVar
}

// WithTo adds the to to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) WithTo(to *strfmt.DateTime) *GetCheckLocationMetricDataParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get check location metric data params
func (o *GetCheckLocationMetricDataParams) SetTo(to *strfmt.DateTime) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetCheckLocationMetricDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom strfmt.DateTime

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom.String()
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.IncludeSummary != nil {

		// query param include_summary
		var qrIncludeSummary bool

		if o.IncludeSummary != nil {
			qrIncludeSummary = *o.IncludeSummary
		}
		qIncludeSummary := swag.FormatBool(qrIncludeSummary)
		if qIncludeSummary != "" {

			if err := r.SetQueryParam("include_summary", qIncludeSummary); err != nil {
				return err
			}
		}
	}

	if o.Locations != nil {

		// binding items for locations
		joinedLocations := o.bindParamLocations(reg)

		// query array param locations
		if err := r.SetQueryParam("locations", joinedLocations...); err != nil {
			return err
		}
	}

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if o.Range != nil {

		// query param range
		var qrRange string

		if o.Range != nil {
			qrRange = *o.Range
		}
		qRange := qrRange
		if qRange != "" {

			if err := r.SetQueryParam("range", qRange); err != nil {
				return err
			}
		}
	}

	if o.To != nil {

		// query param to
		var qrTo strfmt.DateTime

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo.String()
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetCheckLocationMetricData binds the parameter locations
func (o *GetCheckLocationMetricDataParams) bindParamLocations(formats strfmt.Registry) []string {
	locationsIR := o.Locations

	var locationsIC []string
	for _, locationsIIR := range locationsIR { // explode []int32

		locationsIIV := swag.FormatInt32(locationsIIR) // int32 as string
		locationsIC = append(locationsIC, locationsIIV)
	}

	// items.CollectionFormat: ""
	locationsIS := swag.JoinByFormat(locationsIC, "")

	return locationsIS
}

// bindParamGetCheckLocationMetricData binds the parameter names
func (o *GetCheckLocationMetricDataParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: ""
	namesIS := swag.JoinByFormat(namesIC, "")

	return namesIS
}
