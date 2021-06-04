// Code generated by go-swagger; DO NOT EDIT.

package system_events

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

// NewUpdateCheckSystemEventParams creates a new UpdateCheckSystemEventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCheckSystemEventParams() *UpdateCheckSystemEventParams {
	return &UpdateCheckSystemEventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCheckSystemEventParamsWithTimeout creates a new UpdateCheckSystemEventParams object
// with the ability to set a timeout on a request.
func NewUpdateCheckSystemEventParamsWithTimeout(timeout time.Duration) *UpdateCheckSystemEventParams {
	return &UpdateCheckSystemEventParams{
		timeout: timeout,
	}
}

// NewUpdateCheckSystemEventParamsWithContext creates a new UpdateCheckSystemEventParams object
// with the ability to set a context for a request.
func NewUpdateCheckSystemEventParamsWithContext(ctx context.Context) *UpdateCheckSystemEventParams {
	return &UpdateCheckSystemEventParams{
		Context: ctx,
	}
}

// NewUpdateCheckSystemEventParamsWithHTTPClient creates a new UpdateCheckSystemEventParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCheckSystemEventParamsWithHTTPClient(client *http.Client) *UpdateCheckSystemEventParams {
	return &UpdateCheckSystemEventParams{
		HTTPClient: client,
	}
}

/* UpdateCheckSystemEventParams contains all the parameters to send to the API endpoint
   for the update check system event operation.

   Typically these are written to a http.Request.
*/
type UpdateCheckSystemEventParams struct {

	/* CheckID.

	   ID of check whose system event you wish to update

	   Format: int32
	*/
	CheckID int32

	/* ID.

	   ID of system event to fetch

	   Format: int32
	*/
	ID int32

	// SystemEvent.
	SystemEvent UpdateCheckSystemEventBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update check system event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCheckSystemEventParams) WithDefaults() *UpdateCheckSystemEventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update check system event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCheckSystemEventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update check system event params
func (o *UpdateCheckSystemEventParams) WithTimeout(timeout time.Duration) *UpdateCheckSystemEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update check system event params
func (o *UpdateCheckSystemEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update check system event params
func (o *UpdateCheckSystemEventParams) WithContext(ctx context.Context) *UpdateCheckSystemEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update check system event params
func (o *UpdateCheckSystemEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update check system event params
func (o *UpdateCheckSystemEventParams) WithHTTPClient(client *http.Client) *UpdateCheckSystemEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update check system event params
func (o *UpdateCheckSystemEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckID adds the checkID to the update check system event params
func (o *UpdateCheckSystemEventParams) WithCheckID(checkID int32) *UpdateCheckSystemEventParams {
	o.SetCheckID(checkID)
	return o
}

// SetCheckID adds the checkId to the update check system event params
func (o *UpdateCheckSystemEventParams) SetCheckID(checkID int32) {
	o.CheckID = checkID
}

// WithID adds the id to the update check system event params
func (o *UpdateCheckSystemEventParams) WithID(id int32) *UpdateCheckSystemEventParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update check system event params
func (o *UpdateCheckSystemEventParams) SetID(id int32) {
	o.ID = id
}

// WithSystemEvent adds the systemEvent to the update check system event params
func (o *UpdateCheckSystemEventParams) WithSystemEvent(systemEvent UpdateCheckSystemEventBody) *UpdateCheckSystemEventParams {
	o.SetSystemEvent(systemEvent)
	return o
}

// SetSystemEvent adds the systemEvent to the update check system event params
func (o *UpdateCheckSystemEventParams) SetSystemEvent(systemEvent UpdateCheckSystemEventBody) {
	o.SystemEvent = systemEvent
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCheckSystemEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.SystemEvent); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
