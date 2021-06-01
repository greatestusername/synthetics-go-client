// Code generated by go-swagger; DO NOT EDIT.

package real_browser_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/greatestusername/synthetics-go-client/models"
)

// GetRealBrowserCheckRunReader is a Reader for the GetRealBrowserCheckRun structure.
type GetRealBrowserCheckRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRealBrowserCheckRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRealBrowserCheckRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRealBrowserCheckRunOK creates a GetRealBrowserCheckRunOK with default headers values
func NewGetRealBrowserCheckRunOK() *GetRealBrowserCheckRunOK {
	return &GetRealBrowserCheckRunOK{}
}

/* GetRealBrowserCheckRunOK describes a response with status code 200, with default header values.

Real Browser check run response
*/
type GetRealBrowserCheckRunOK struct {
	Payload *models.RealBrowserCheckRun
}

func (o *GetRealBrowserCheckRunOK) Error() string {
	return fmt.Sprintf("[GET /v2/checks/real_browsers/{check_id}/runs/{id}][%d] getRealBrowserCheckRunOK  %+v", 200, o.Payload)
}
func (o *GetRealBrowserCheckRunOK) GetPayload() *models.RealBrowserCheckRun {
	return o.Payload
}

func (o *GetRealBrowserCheckRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RealBrowserCheckRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}