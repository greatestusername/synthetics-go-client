// Code generated by go-swagger; DO NOT EDIT.

package system_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/greatestusername/synthetics-go-client/models"
)

// DeleteCheckSystemEventReader is a Reader for the DeleteCheckSystemEvent structure.
type DeleteCheckSystemEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCheckSystemEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCheckSystemEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCheckSystemEventOK creates a DeleteCheckSystemEventOK with default headers values
func NewDeleteCheckSystemEventOK() *DeleteCheckSystemEventOK {
	return &DeleteCheckSystemEventOK{}
}

/* DeleteCheckSystemEventOK describes a response with status code 200, with default header values.

OK
*/
type DeleteCheckSystemEventOK struct {
	Payload *models.SystemEvent
}

func (o *DeleteCheckSystemEventOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/checks/{check_id}/system_events/{id}][%d] deleteCheckSystemEventOK  %+v", 200, o.Payload)
}
func (o *DeleteCheckSystemEventOK) GetPayload() *models.SystemEvent {
	return o.Payload
}

func (o *DeleteCheckSystemEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
