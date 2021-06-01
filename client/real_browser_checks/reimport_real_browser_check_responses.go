// Code generated by go-swagger; DO NOT EDIT.

package real_browser_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/greatestusername/synthetics-go-client/models"
)

// ReimportRealBrowserCheckReader is a Reader for the ReimportRealBrowserCheck structure.
type ReimportRealBrowserCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReimportRealBrowserCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReimportRealBrowserCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReimportRealBrowserCheckOK creates a ReimportRealBrowserCheckOK with default headers values
func NewReimportRealBrowserCheckOK() *ReimportRealBrowserCheckOK {
	var (
		// initialize headers with default values
		xRateLimitLimitDefault = int64(5000)

		xRateLimitResetDefault = int64(1621968845)
	)

	return &ReimportRealBrowserCheckOK{

		XRateLimitLimit: xRateLimitLimitDefault,
		XRateLimitReset: xRateLimitResetDefault,
	}
}

/* ReimportRealBrowserCheckOK describes a response with status code 200, with default header values.

Real Browser check response
*/
type ReimportRealBrowserCheckOK struct {

	/* The number of requests a user is allowed per hour. Users are identified by IP address.

	   Default: 5000
	*/
	XRateLimitLimit int64

	/* The number of requests remaining in the current rate limit window
	 */
	XRateLimitRemaining int64

	/* When the current rate limit window resets (in UTC epoch seconds).

	   Default: 1621968845
	*/
	XRateLimitReset int64

	Payload *models.RealBrowserCheck
}

func (o *ReimportRealBrowserCheckOK) Error() string {
	return fmt.Sprintf("[PUT /v2/checks/real_browsers/{id}/import][%d] reimportRealBrowserCheckOK  %+v", 200, o.Payload)
}
func (o *ReimportRealBrowserCheckOK) GetPayload() *models.RealBrowserCheck {
	return o.Payload
}

func (o *ReimportRealBrowserCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-Reset
	hdrXRateLimitReset := response.GetHeader("X-RateLimit-Reset")

	if hdrXRateLimitReset != "" {
		valxRateLimitReset, err := swag.ConvertInt64(hdrXRateLimitReset)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Reset", "header", "int64", hdrXRateLimitReset)
		}
		o.XRateLimitReset = valxRateLimitReset
	}

	o.Payload = new(models.RealBrowserCheck)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}