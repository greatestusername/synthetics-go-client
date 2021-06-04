// Code generated by go-swagger; DO NOT EDIT.

package comparison_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetComparisonReportsReader is a Reader for the GetComparisonReports structure.
type GetComparisonReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComparisonReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComparisonReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComparisonReportsOK creates a GetComparisonReportsOK with default headers values
func NewGetComparisonReportsOK() *GetComparisonReportsOK {
	var (
		// initialize headers with default values
		xRateLimitLimitDefault = int64("5000")

		xRateLimitResetDefault = int64("1621968845")
	)

	return &GetComparisonReportsOK{

		XRateLimitLimit: xRateLimitLimitDefault,
		XRateLimitReset: xRateLimitResetDefault,
	}
}

/* GetComparisonReportsOK describes a response with status code 200, with default header values.

OK
*/
type GetComparisonReportsOK struct {

	/* The number of requests a user is allowed per hour. Users are identified by IP address.

	   Default: "5000"
	*/
	XRateLimitLimit int64

	/* The number of requests remaining in the current rate limit window
	 */
	XRateLimitRemaining int64

	/* When the current rate limit window resets (in UTC epoch seconds).

	   Default: "1621968845"
	*/
	XRateLimitReset int64

	Payload *GetComparisonReportsOKBody
}

func (o *GetComparisonReportsOK) Error() string {
	return fmt.Sprintf("[GET /v2/comparison_reports][%d] getComparisonReportsOK  %+v", 200, o.Payload)
}
func (o *GetComparisonReportsOK) GetPayload() *GetComparisonReportsOKBody {
	return o.Payload
}

func (o *GetComparisonReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetComparisonReportsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetComparisonReportsOKBody A page of Comparison Reports
swagger:model GetComparisonReportsOKBody
*/
type GetComparisonReportsOKBody struct {
	ComparisonReports []GetComparisonReportsOKBodyComparisonReportsItems0 `json:"comparison_reports"`

	// Current page number
	// Example: 2
	CurrentPage int32 `json:"current_page,omitempty"`

	// Next page number (null if none)
	// Example: 3
	NextPage int32 `json:"next_page,omitempty"`

	// Number of results for each page
	// Example: 50
	PerPage int32 `json:"per_page,omitempty"`

	// Previous page number (null if none)
	// Example: 1
	PreviousPage int32 `json:"previous_page,omitempty"`

	// Total number of results across all pages
	// Example: 105
	TotalCount int32 `json:"total_count,omitempty"`

	// Total number of pages
	// Example: 3
	TotalPages int32 `json:"total_pages,omitempty"`
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (o *GetComparisonReportsOKBody) UnmarshalJSON(raw []byte) error {
	var data struct {
		ComparisonReports json.RawMessage `json:"comparison_reports"`

		// Current page number
		// Example: 2
		CurrentPage int32 `json:"current_page,omitempty"`

		// Next page number (null if none)
		// Example: 3
		NextPage int32 `json:"next_page,omitempty"`

		// Number of results for each page
		// Example: 50
		PerPage int32 `json:"per_page,omitempty"`

		// Previous page number (null if none)
		// Example: 1
		PreviousPage int32 `json:"previous_page,omitempty"`

		// Total number of results across all pages
		// Example: 105
		TotalCount int32 `json:"total_count,omitempty"`

		// Total number of pages
		// Example: 3
		TotalPages int32 `json:"total_pages,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var allOfComparisonReports []GetComparisonReportsOKBodyComparisonReportsItems0
	if string(data.ComparisonReports) != "null" {
		comparisonReports, err := UnmarshalGetComparisonReportsOKBodyComparisonReportsItems0Slice(bytes.NewBuffer(data.ComparisonReports), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		allOfComparisonReports = comparisonReports
	}

	var result GetComparisonReportsOKBody

	result.comparisonReportsField = allOfComparisonReports

	result.CurrentPage = data.CurrentPage
	result.NextPage = data.NextPage
	result.PerPage = data.PerPage
	result.PreviousPage = data.PreviousPage
	result.TotalCount = data.TotalCount
	result.TotalPages = data.TotalPages

	*o = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (o GetComparisonReportsOKBody) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Current page number
		// Example: 2
		CurrentPage int32 `json:"current_page,omitempty"`

		// Next page number (null if none)
		// Example: 3
		NextPage int32 `json:"next_page,omitempty"`

		// Number of results for each page
		// Example: 50
		PerPage int32 `json:"per_page,omitempty"`

		// Previous page number (null if none)
		// Example: 1
		PreviousPage int32 `json:"previous_page,omitempty"`

		// Total number of results across all pages
		// Example: 105
		TotalCount int32 `json:"total_count,omitempty"`

		// Total number of pages
		// Example: 3
		TotalPages int32 `json:"total_pages,omitempty"`
	}{

		CurrentPage: o.CurrentPage,

		NextPage: o.NextPage,

		PerPage: o.PerPage,

		PreviousPage: o.PreviousPage,

		TotalCount: o.TotalCount,

		TotalPages: o.TotalPages,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ComparisonReports []GetComparisonReportsOKBodyComparisonReportsItems0 `json:"comparison_reports"`
	}{

		ComparisonReports: o.ComparisonReports(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this get comparison reports o k body
func (o *GetComparisonReportsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComparisonReports(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComparisonReportsOKBody) validateComparisonReports(formats strfmt.Registry) error {

	if swag.IsZero(o.ComparisonReports()) { // not required
		return nil
	}

	for i := 0; i < len(o.ComparisonReports()); i++ {

		if err := o.ComparisonReports[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getComparisonReportsOK" + "." + "comparison_reports" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this get comparison reports o k body based on the context it is used
func (o *GetComparisonReportsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateComparisonReports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComparisonReportsOKBody) contextValidateComparisonReports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ComparisonReports()); i++ {

		if err := o.ComparisonReports[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getComparisonReportsOK" + "." + "comparison_reports" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetComparisonReportsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetComparisonReportsOKBody) UnmarshalBinary(b []byte) error {
	var res GetComparisonReportsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetComparisonReportsOKBodyComparisonReportsItems0 A single Comparison Report
swagger:model GetComparisonReportsOKBodyComparisonReportsItems0
*/
type GetComparisonReportsOKBodyComparisonReportsItems0 interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The function used to aggregate data for the Comparison Report
	// Example: mean
	// Enum: [mean median max min sum stddev percentile_90 percentile_95 percentile_99]
	Aggregate() string
	SetAggregate(string)

	// Used to aggregate and display data at this interval. It must be shorter
	//                               than the duration of the report timeframe.
	// Example: 1m
	// Enum: [1m 5m 10m 15m 30m 1h 3h 6h 12h 24h 2d 3d 1w 2w 30d]
	DataPointInterval() string
	SetDataPointInterval(string)

	// When set to `true`, excludes failed runs from the Comparison Report results
	// Example: false
	ExcludeFailures() bool
	SetExcludeFailures(bool)

	// The beginning of the report timeframe (if range is "custom")
	// Example: 2021-05-24T17:54:05Z
	From() string
	SetFrom(string)

	// The metric to compare on the Comparison Report
	// Example: first_byte_time_ms
	// Enum: [first_byte_time_ms dom_interactive_time_ms dom_load_time_ms dom_complete_time_ms start_render_ms onload_time_ms visually_complete_ms fully_loaded_time_ms first_paint_time_ms first_contentful_paint_time_ms first_meaningful_paint_time_ms first_interactive_time_ms first_cpu_idle_time_ms first_request_dns_time_ms first_request_connect_time_ms first_request_ssl_time_ms first_request_send_time_ms first_request_wait_time_ms first_request_receive_time_ms speed_index requests content_bytes html_files html_bytes image_files image_bytes javascript_files javascript_bytes css_files css_bytes video_files video_bytes font_files font_bytes other_files other_bytes client_errors connection_errors server_errors errors run_count success_count failure_count lighthouse_performance_score availability downtime total_blocking_time_ms largest_contentful_paint_time_ms cumulative_layout_shift]
	Metric() string
	SetMetric(string)

	// The unique name for the Comparison Report
	// Example: Example Comparison Report
	Name() string
	SetName(string)

	// The time range that the Comparison Report spans over. Set to `custom`
	//                               when the report is set to a static timeframe.
	// Example: custom
	// Enum: [last_hour last_4_hours last_8_hours last_12_hours last_24_hours yesterday today last_7_days last_30_days this_week last_week this_month last_month last_3_months last_6_months custom]
	Range() string
	SetRange(string)

	// The end of the report timeframe (if range is "custom")
	// Example: 2021-05-25T17:54:05Z
	To() string
	SetTo(string)

	// account
	Account() *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account
	SetAccount(*GetComparisonReportsOKBodyComparisonReportsItems0AO1Account)

	// An array of Real Browser checks selected for the Comparison Report. Needs to have
	//                                   at least two checks per comparison report
	Checks() []*GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0
	SetChecks([]*GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0)

	// When the Comparison Report was created (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	CreatedAt() strfmt.DateTime
	SetCreatedAt(strfmt.DateTime)

	// The unique ID of the Comparison Report
	// Example: 1
	ID() int32
	SetID(int32)

	// An array of location IDs selected for the Comparison Report. If this array is empty,
	//                                   the report will include data from all locations.
	Locations() []*GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0
	SetLocations([]*GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0)

	// A shareable link that can be viewed by anyone
	// Example: https://monitoring.rigor.com/share/64ec88ca00b268e5ba1a35678a1b5316d212f4f366b2477232534a8aeca37f3c*OzE7Mg==
	ShareLink() string
	SetShareLink(string)

	// When the Comparison Report was last updated (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	UpdatedAt() strfmt.DateTime
	SetUpdatedAt(strfmt.DateTime)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type getComparisonReportsOKBodyComparisonReportsItems0 struct {

	// The function used to aggregate data for the Comparison Report
	// Example: mean
	// Enum: [mean median max min sum stddev percentile_90 percentile_95 percentile_99]
	Aggregate string `json:"aggregate,omitempty"`

	// Used to aggregate and display data at this interval. It must be shorter
	//                               than the duration of the report timeframe.
	// Example: 1m
	// Enum: [1m 5m 10m 15m 30m 1h 3h 6h 12h 24h 2d 3d 1w 2w 30d]
	DataPointInterval string `json:"data_point_interval,omitempty"`

	// When set to `true`, excludes failed runs from the Comparison Report results
	// Example: false
	ExcludeFailures bool `json:"exclude_failures,omitempty"`

	// The beginning of the report timeframe (if range is "custom")
	// Example: 2021-05-24T17:54:05Z
	From string `json:"from,omitempty"`

	// The metric to compare on the Comparison Report
	// Example: first_byte_time_ms
	// Enum: [first_byte_time_ms dom_interactive_time_ms dom_load_time_ms dom_complete_time_ms start_render_ms onload_time_ms visually_complete_ms fully_loaded_time_ms first_paint_time_ms first_contentful_paint_time_ms first_meaningful_paint_time_ms first_interactive_time_ms first_cpu_idle_time_ms first_request_dns_time_ms first_request_connect_time_ms first_request_ssl_time_ms first_request_send_time_ms first_request_wait_time_ms first_request_receive_time_ms speed_index requests content_bytes html_files html_bytes image_files image_bytes javascript_files javascript_bytes css_files css_bytes video_files video_bytes font_files font_bytes other_files other_bytes client_errors connection_errors server_errors errors run_count success_count failure_count lighthouse_performance_score availability downtime total_blocking_time_ms largest_contentful_paint_time_ms cumulative_layout_shift]
	Metric string `json:"metric,omitempty"`

	// The unique name for the Comparison Report
	// Example: Example Comparison Report
	Name string `json:"name,omitempty"`

	// The time range that the Comparison Report spans over. Set to `custom`
	//                               when the report is set to a static timeframe.
	// Example: custom
	// Enum: [last_hour last_4_hours last_8_hours last_12_hours last_24_hours yesterday today last_7_days last_30_days this_week last_week this_month last_month last_3_months last_6_months custom]
	Range string `json:"range,omitempty"`

	// The end of the report timeframe (if range is "custom")
	// Example: 2021-05-25T17:54:05Z
	To string `json:"to,omitempty"`

	// account
	Account *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account `json:"account,omitempty"`

	// An array of Real Browser checks selected for the Comparison Report. Needs to have
	//                                   at least two checks per comparison report
	Checks []*GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0 `json:"checks"`

	// When the Comparison Report was created (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The unique ID of the Comparison Report
	// Example: 1
	ID int32 `json:"id,omitempty"`

	// An array of location IDs selected for the Comparison Report. If this array is empty,
	//                                   the report will include data from all locations.
	Locations []*GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0 `json:"locations"`

	// A shareable link that can be viewed by anyone
	// Example: https://monitoring.rigor.com/share/64ec88ca00b268e5ba1a35678a1b5316d212f4f366b2477232534a8aeca37f3c*OzE7Mg==
	ShareLink string `json:"share_link,omitempty"`

	// When the Comparison Report was last updated (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// UnmarshalGetComparisonReportsOKBodyComparisonReportsItems0Slice unmarshals polymorphic slices of GetComparisonReportsOKBodyComparisonReportsItems0
func UnmarshalGetComparisonReportsOKBodyComparisonReportsItems0Slice(reader io.Reader, consumer runtime.Consumer) ([]GetComparisonReportsOKBodyComparisonReportsItems0, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []GetComparisonReportsOKBodyComparisonReportsItems0
	for _, element := range elements {
		obj, err := unmarshalGetComparisonReportsOKBodyComparisonReportsItems0(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalGetComparisonReportsOKBodyComparisonReportsItems0 unmarshals polymorphic GetComparisonReportsOKBodyComparisonReportsItems0
func UnmarshalGetComparisonReportsOKBodyComparisonReportsItems0(reader io.Reader, consumer runtime.Consumer) (GetComparisonReportsOKBodyComparisonReportsItems0, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalGetComparisonReportsOKBodyComparisonReportsItems0(data, consumer)
}

func unmarshalGetComparisonReportsOKBodyComparisonReportsItems0(data []byte, consumer runtime.Consumer) (GetComparisonReportsOKBodyComparisonReportsItems0, error) {
	buf := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the  property.
	var getType struct {
		Empty string `json:""`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("", "body", getType.Empty); err != nil {
		return nil, err
	}

	// The value of  is used to determine which type to create and unmarshal the data into
	switch getType.Empty {
	}
	return nil, errors.New(422, "invalid  value: %q", getType.Empty)
}

// Validate validates this get comparison reports o k body comparison reports items0
func (o *GetComparisonReportsOKBodyComparisonReportsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDataPointInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getComparisonReportsOKBodyComparisonReportsItems0TypeAggregatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mean","median","max","min","sum","stddev","percentile_90","percentile_95","percentile_99"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getComparisonReportsOKBodyComparisonReportsItems0TypeAggregatePropEnum = append(getComparisonReportsOKBodyComparisonReportsItems0TypeAggregatePropEnum, v)
	}
}

// property enum
func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateAggregateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getComparisonReportsOKBodyComparisonReportsItems0TypeAggregatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateAggregate(formats strfmt.Registry) error {

	if swag.IsZero(o.Aggregate) { // not required
		return nil
	}

	// value enum
	if err := o.validateAggregateEnum("aggregate", "body", o.Aggregate); err != nil {
		return err
	}

	return nil
}

var getComparisonReportsOKBodyComparisonReportsItems0TypeDataPointIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1m","5m","10m","15m","30m","1h","3h","6h","12h","24h","2d","3d","1w","2w","30d"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getComparisonReportsOKBodyComparisonReportsItems0TypeDataPointIntervalPropEnum = append(getComparisonReportsOKBodyComparisonReportsItems0TypeDataPointIntervalPropEnum, v)
	}
}

// property enum
func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateDataPointIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getComparisonReportsOKBodyComparisonReportsItems0TypeDataPointIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateDataPointInterval(formats strfmt.Registry) error {

	if swag.IsZero(o.DataPointInterval) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataPointIntervalEnum("data_point_interval", "body", o.DataPointInterval); err != nil {
		return err
	}

	return nil
}

var getComparisonReportsOKBodyComparisonReportsItems0TypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["first_byte_time_ms","dom_interactive_time_ms","dom_load_time_ms","dom_complete_time_ms","start_render_ms","onload_time_ms","visually_complete_ms","fully_loaded_time_ms","first_paint_time_ms","first_contentful_paint_time_ms","first_meaningful_paint_time_ms","first_interactive_time_ms","first_cpu_idle_time_ms","first_request_dns_time_ms","first_request_connect_time_ms","first_request_ssl_time_ms","first_request_send_time_ms","first_request_wait_time_ms","first_request_receive_time_ms","speed_index","requests","content_bytes","html_files","html_bytes","image_files","image_bytes","javascript_files","javascript_bytes","css_files","css_bytes","video_files","video_bytes","font_files","font_bytes","other_files","other_bytes","client_errors","connection_errors","server_errors","errors","run_count","success_count","failure_count","lighthouse_performance_score","availability","downtime","total_blocking_time_ms","largest_contentful_paint_time_ms","cumulative_layout_shift"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getComparisonReportsOKBodyComparisonReportsItems0TypeMetricPropEnum = append(getComparisonReportsOKBodyComparisonReportsItems0TypeMetricPropEnum, v)
	}
}

// property enum
func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getComparisonReportsOKBodyComparisonReportsItems0TypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateMetric(formats strfmt.Registry) error {

	if swag.IsZero(o.Metric) { // not required
		return nil
	}

	// value enum
	if err := o.validateMetricEnum("metric", "body", o.Metric); err != nil {
		return err
	}

	return nil
}

var getComparisonReportsOKBodyComparisonReportsItems0TypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["last_hour","last_4_hours","last_8_hours","last_12_hours","last_24_hours","yesterday","today","last_7_days","last_30_days","this_week","last_week","this_month","last_month","last_3_months","last_6_months","custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getComparisonReportsOKBodyComparisonReportsItems0TypeRangePropEnum = append(getComparisonReportsOKBodyComparisonReportsItems0TypeRangePropEnum, v)
	}
}

// property enum
func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateRangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getComparisonReportsOKBodyComparisonReportsItems0TypeRangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(o.Range) { // not required
		return nil
	}

	// value enum
	if err := o.validateRangeEnum("range", "body", o.Range); err != nil {
		return err
	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(o.Account) { // not required
		return nil
	}

	if o.Account != nil {
		if err := o.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateChecks(formats strfmt.Registry) error {

	if swag.IsZero(o.Checks) { // not required
		return nil
	}

	for i := 0; i < len(o.Checks); i++ {
		if swag.IsZero(o.Checks[i]) { // not required
			continue
		}

		if o.Checks[i] != nil {
			if err := o.Checks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(o.Locations) { // not required
		return nil
	}

	for i := 0; i < len(o.Locations); i++ {
		if swag.IsZero(o.Locations[i]) { // not required
			continue
		}

		if o.Locations[i] != nil {
			if err := o.Locations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get comparison reports o k body comparison reports items0 based on the context it is used
func (o *GetComparisonReportsOKBodyComparisonReportsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if o.Account != nil {
		if err := o.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) contextValidateChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Checks); i++ {

		if o.Checks[i] != nil {
			if err := o.Checks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0) contextValidateLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Locations); i++ {

		if o.Locations[i] != nil {
			if err := o.Locations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

/*GetComparisonReportsOKBodyComparisonReportsItems0AO1Account get comparison reports o k body comparison reports items0 a o1 account
swagger:model GetComparisonReportsOKBodyComparisonReportsItems0AO1Account
*/
type GetComparisonReportsOKBodyComparisonReportsItems0AO1Account struct {

	// The unique ID of the account
	// Example: 1
	// Required: true
	ID *int32 `json:"id"`

	// The name of the account
	// Example: Example Company
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this get comparison reports o k body comparison reports items0 a o1 account
func (o *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account) validateID(formats strfmt.Registry) error {

	if err := validate.Required("account"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account) validateName(formats strfmt.Registry) error {

	if err := validate.Required("account"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get comparison reports o k body comparison reports items0 a o1 account based on context it is used
func (o *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetComparisonReportsOKBodyComparisonReportsItems0AO1Account) UnmarshalBinary(b []byte) error {
	var res GetComparisonReportsOKBodyComparisonReportsItems0AO1Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0 get comparison reports o k body comparison reports items0 checks items0
swagger:model GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0
*/
type GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0 struct {

	// When set to `true`, the metrics from this check are used as a baseline for the other checks
	// Example: true
	Baseline bool `json:"baseline,omitempty"`

	// The unique ID of the check attached to the Comparison Report
	// Example: 1
	// Required: true
	ID *int32 `json:"id"`

	// An optional alias to use in place of the check name in the Comparison Report
	// Example: Example Nickname
	Nickname string `json:"nickname,omitempty"`

	// The name of the check attached to the Comparison Report
	// Example: Example Check
	Name string `json:"name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Baseline bool `json:"baseline,omitempty"`

		ID *int32 `json:"id"`

		Nickname string `json:"nickname,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.Baseline = dataAO0.Baseline

	o.ID = dataAO0.ID

	o.Nickname = dataAO0.Nickname

	// AO1
	var dataAO1 struct {
		Name string `json:"name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Baseline bool `json:"baseline,omitempty"`

		ID *int32 `json:"id"`

		Nickname string `json:"nickname,omitempty"`
	}

	dataAO0.Baseline = o.Baseline

	dataAO0.ID = o.ID

	dataAO0.Nickname = o.Nickname

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		Name string `json:"name,omitempty"`
	}

	dataAO1.Name = o.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get comparison reports o k body comparison reports items0 checks items0
func (o *GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get comparison reports o k body comparison reports items0 checks items0 based on context it is used
func (o *GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0) UnmarshalBinary(b []byte) error {
	var res GetComparisonReportsOKBodyComparisonReportsItems0ChecksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0 get comparison reports o k body comparison reports items0 locations items0
swagger:model GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0
*/
type GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0 struct {

	// The unique ID of the location
	// Example: 1
	// Required: true
	ID *int32 `json:"id"`

	// The name of the location
	// Example: Example Location
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this get comparison reports o k body comparison reports items0 locations items0
func (o *GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get comparison reports o k body comparison reports items0 locations items0 based on context it is used
func (o *GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0) UnmarshalBinary(b []byte) error {
	var res GetComparisonReportsOKBodyComparisonReportsItems0LocationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
