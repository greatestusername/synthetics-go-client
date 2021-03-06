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

// NewGetRealBrowserCheckRunsParams creates a new GetRealBrowserCheckRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRealBrowserCheckRunsParams() *GetRealBrowserCheckRunsParams {
	return &GetRealBrowserCheckRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRealBrowserCheckRunsParamsWithTimeout creates a new GetRealBrowserCheckRunsParams object
// with the ability to set a timeout on a request.
func NewGetRealBrowserCheckRunsParamsWithTimeout(timeout time.Duration) *GetRealBrowserCheckRunsParams {
	return &GetRealBrowserCheckRunsParams{
		timeout: timeout,
	}
}

// NewGetRealBrowserCheckRunsParamsWithContext creates a new GetRealBrowserCheckRunsParams object
// with the ability to set a context for a request.
func NewGetRealBrowserCheckRunsParamsWithContext(ctx context.Context) *GetRealBrowserCheckRunsParams {
	return &GetRealBrowserCheckRunsParams{
		Context: ctx,
	}
}

// NewGetRealBrowserCheckRunsParamsWithHTTPClient creates a new GetRealBrowserCheckRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRealBrowserCheckRunsParamsWithHTTPClient(client *http.Client) *GetRealBrowserCheckRunsParams {
	return &GetRealBrowserCheckRunsParams{
		HTTPClient: client,
	}
}

/* GetRealBrowserCheckRunsParams contains all the parameters to send to the API endpoint
   for the get real browser check runs operation.

   Typically these are written to a http.Request.
*/
type GetRealBrowserCheckRunsParams struct {

	/* CheckID.

	   ID of the Real Browser check to fetch

	   Format: int32
	*/
	CheckID int32

	/* From.

	   The start time for the timeframe

	   Format: date-time
	*/
	From *strfmt.DateTime

	/* Page.

	   Page number, starting at 1

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* PerPage.

	   The number of checks returned in each page

	   Format: int32
	   Default: 50
	*/
	PerPage *int32

	/* Range.

	   A predefined timeframe to be used instead of `from` and `to`
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

// WithDefaults hydrates default values in the get real browser check runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRealBrowserCheckRunsParams) WithDefaults() *GetRealBrowserCheckRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get real browser check runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRealBrowserCheckRunsParams) SetDefaults() {
	var (
		pageDefault = int32(1)

		perPageDefault = int32(50)
	)

	val := GetRealBrowserCheckRunsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithTimeout(timeout time.Duration) *GetRealBrowserCheckRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithContext(ctx context.Context) *GetRealBrowserCheckRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithHTTPClient(client *http.Client) *GetRealBrowserCheckRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckID adds the checkID to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithCheckID(checkID int32) *GetRealBrowserCheckRunsParams {
	o.SetCheckID(checkID)
	return o
}

// SetCheckID adds the checkId to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetCheckID(checkID int32) {
	o.CheckID = checkID
}

// WithFrom adds the from to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithFrom(from *strfmt.DateTime) *GetRealBrowserCheckRunsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetFrom(from *strfmt.DateTime) {
	o.From = from
}

// WithPage adds the page to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithPage(page *int32) *GetRealBrowserCheckRunsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithPerPage(perPage *int32) *GetRealBrowserCheckRunsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithRange adds the rangeVar to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithRange(rangeVar *string) *GetRealBrowserCheckRunsParams {
	o.SetRange(rangeVar)
	return o
}

// SetRange adds the range to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetRange(rangeVar *string) {
	o.Range = rangeVar
}

// WithTo adds the to to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) WithTo(to *strfmt.DateTime) *GetRealBrowserCheckRunsParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get real browser check runs params
func (o *GetRealBrowserCheckRunsParams) SetTo(to *strfmt.DateTime) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetRealBrowserCheckRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param check_id
	if err := r.SetPathParam("check_id", swag.FormatInt32(o.CheckID)); err != nil {
		return err
	}

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

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
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
