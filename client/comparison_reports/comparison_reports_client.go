// Code generated by go-swagger; DO NOT EDIT.

package comparison_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new comparison reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for comparison reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateComparisonReport(params *CreateComparisonReportParams, opts ...ClientOption) (*CreateComparisonReportOK, error)

	DeleteComparisonReport(params *DeleteComparisonReportParams, opts ...ClientOption) (*DeleteComparisonReportOK, error)

	GetComparisonReport(params *GetComparisonReportParams, opts ...ClientOption) (*GetComparisonReportOK, error)

	GetComparisonReports(params *GetComparisonReportsParams, opts ...ClientOption) (*GetComparisonReportsOK, error)

	UpdateComparisonReport(params *UpdateComparisonReportParams, opts ...ClientOption) (*UpdateComparisonReportOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateComparisonReport creates a comparison report

  Creates a new Comparison Report and returns the report configuration
*/
func (a *Client) CreateComparisonReport(params *CreateComparisonReportParams, opts ...ClientOption) (*CreateComparisonReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateComparisonReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createComparisonReport",
		Method:             "POST",
		PathPattern:        "/v2/comparison_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateComparisonReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateComparisonReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createComparisonReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteComparisonReport deletes a comparison report

  Deletes a Comparison Report and returns a confirmation message
*/
func (a *Client) DeleteComparisonReport(params *DeleteComparisonReportParams, opts ...ClientOption) (*DeleteComparisonReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComparisonReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteComparisonReport",
		Method:             "DELETE",
		PathPattern:        "/v2/comparison_reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteComparisonReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteComparisonReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteComparisonReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetComparisonReport gets comparison report detail

  Returns a Comparison Report based on a single ID
*/
func (a *Client) GetComparisonReport(params *GetComparisonReportParams, opts ...ClientOption) (*GetComparisonReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComparisonReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComparisonReport",
		Method:             "GET",
		PathPattern:        "/v2/comparison_reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComparisonReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComparisonReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComparisonReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetComparisonReports gets comparison reports

  Returns all Comparison Reports in the account
*/
func (a *Client) GetComparisonReports(params *GetComparisonReportsParams, opts ...ClientOption) (*GetComparisonReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComparisonReportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComparisonReports",
		Method:             "GET",
		PathPattern:        "/v2/comparison_reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComparisonReportsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComparisonReportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComparisonReports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateComparisonReport updates a comparison report

  Updates a Comparison Report and returns the report configuration
*/
func (a *Client) UpdateComparisonReport(params *UpdateComparisonReportParams, opts ...ClientOption) (*UpdateComparisonReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateComparisonReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateComparisonReport",
		Method:             "PUT",
		PathPattern:        "/v2/comparison_reports/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateComparisonReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateComparisonReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateComparisonReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
