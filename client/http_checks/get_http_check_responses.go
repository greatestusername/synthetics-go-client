// Code generated by go-swagger; DO NOT EDIT.

package http_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/greatestusername/synthetics-go-client/models"
)

// GetHTTPCheckReader is a Reader for the GetHTTPCheck structure.
type GetHTTPCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHTTPCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHTTPCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHTTPCheckOK creates a GetHTTPCheckOK with default headers values
func NewGetHTTPCheckOK() *GetHTTPCheckOK {
	return &GetHTTPCheckOK{}
}

/* GetHTTPCheckOK describes a response with status code 200, with default header values.

HTTP check response
*/
type GetHTTPCheckOK struct {
	Payload *models.HTTPCheck
}

func (o *GetHTTPCheckOK) Error() string {
	return fmt.Sprintf("[GET /v2/checks/http/{id}][%d] getHttpCheckOK  %+v", 200, o.Payload)
}
func (o *GetHTTPCheckOK) GetPayload() *models.HTTPCheck {
	return o.Payload
}

func (o *GetHTTPCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPCheck)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
