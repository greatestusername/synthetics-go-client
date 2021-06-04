// Code generated by go-swagger; DO NOT EDIT.

package checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCheckLocationMetricDataReader is a Reader for the GetCheckLocationMetricData structure.
type GetCheckLocationMetricDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCheckLocationMetricDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCheckLocationMetricDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCheckLocationMetricDataOK creates a GetCheckLocationMetricDataOK with default headers values
func NewGetCheckLocationMetricDataOK() *GetCheckLocationMetricDataOK {
	return &GetCheckLocationMetricDataOK{}
}

/* GetCheckLocationMetricDataOK describes a response with status code 200, with default header values.

Check location data
*/
type GetCheckLocationMetricDataOK struct {
	Payload *GetCheckLocationMetricDataOKBody
}

func (o *GetCheckLocationMetricDataOK) Error() string {
	return fmt.Sprintf("[GET /v2/checks/{id}/locations/metrics/data][%d] getCheckLocationMetricDataOK  %+v", 200, o.Payload)
}
func (o *GetCheckLocationMetricDataOK) GetPayload() *GetCheckLocationMetricDataOKBody {
	return o.Payload
}

func (o *GetCheckLocationMetricDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCheckLocationMetricDataOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCheckLocationMetricDataOKBody Check data segmented by location
swagger:model GetCheckLocationMetricDataOKBody
*/
type GetCheckLocationMetricDataOKBody struct {

	// The start time for the timeframe (UTC)
	// Example: 2021-05-25T16:54:05Z
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from"`

	// The unique ID for the check
	// Example: 1
	ID int32 `json:"id,omitempty"`

	// An array of location IDs representing where the data was gathered.
	//               If a location filter was provided, this will match that parameter. Otherwise, this includes
	//               all locations where the check ran in this timeframe.
	Locations []int32 `json:"locations"`

	// A list of the metric names included in series
	// Example: ["percentage_uptime"]
	// Required: true
	Names []string `json:"names"`

	// The predefined timeframe, if provided
	// Enum: [last_hour last_4_hours last_8_hours last_12_hours last_24_hours yesterday today last_7_days last_30_days this_week last_week this_month month_to_date last_month last_3_months last_6_months last_12_months]
	Range string `json:"range,omitempty"`

	// An array of data sets, one for each location
	// Required: true
	Series []*GetCheckLocationMetricDataOKBodySeriesItems0 `json:"series"`

	// Aggregate data for each location, over the given timeframe
	// Example: {"percentage_uptime":99.3}
	Summary interface{} `json:"summary,omitempty"`

	// The end time for the timeframe (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Required: true
	// Format: date-time
	To *strfmt.DateTime `json:"to"`
}

// Validate validates this get check location metric data o k body
func (o *GetCheckLocationMetricDataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSeries(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBody) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("getCheckLocationMetricDataOK"+"."+"from", "body", o.From); err != nil {
		return err
	}

	if err := validate.FormatOf("getCheckLocationMetricDataOK"+"."+"from", "body", "date-time", o.From.String(), formats); err != nil {
		return err
	}

	return nil
}

var getCheckLocationMetricDataOKBodyNamesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["average_response_time","max_response_time","min_response_time","standard_deviation","run_count","error_count","percentage_uptime","percentage_downtime","sla_percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCheckLocationMetricDataOKBodyNamesItemsEnum = append(getCheckLocationMetricDataOKBodyNamesItemsEnum, v)
	}
}

func (o *GetCheckLocationMetricDataOKBody) validateNamesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCheckLocationMetricDataOKBodyNamesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBody) validateNames(formats strfmt.Registry) error {

	if err := validate.Required("getCheckLocationMetricDataOK"+"."+"names", "body", o.Names); err != nil {
		return err
	}

	for i := 0; i < len(o.Names); i++ {

		// value enum
		if err := o.validateNamesItemsEnum("getCheckLocationMetricDataOK"+"."+"names"+"."+strconv.Itoa(i), "body", o.Names[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCheckLocationMetricDataOKBodyTypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["last_hour","last_4_hours","last_8_hours","last_12_hours","last_24_hours","yesterday","today","last_7_days","last_30_days","this_week","last_week","this_month","month_to_date","last_month","last_3_months","last_6_months","last_12_months"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCheckLocationMetricDataOKBodyTypeRangePropEnum = append(getCheckLocationMetricDataOKBodyTypeRangePropEnum, v)
	}
}

const (

	// GetCheckLocationMetricDataOKBodyRangeLastHour captures enum value "last_hour"
	GetCheckLocationMetricDataOKBodyRangeLastHour string = "last_hour"

	// GetCheckLocationMetricDataOKBodyRangeLast4Hours captures enum value "last_4_hours"
	GetCheckLocationMetricDataOKBodyRangeLast4Hours string = "last_4_hours"

	// GetCheckLocationMetricDataOKBodyRangeLast8Hours captures enum value "last_8_hours"
	GetCheckLocationMetricDataOKBodyRangeLast8Hours string = "last_8_hours"

	// GetCheckLocationMetricDataOKBodyRangeLast12Hours captures enum value "last_12_hours"
	GetCheckLocationMetricDataOKBodyRangeLast12Hours string = "last_12_hours"

	// GetCheckLocationMetricDataOKBodyRangeLast24Hours captures enum value "last_24_hours"
	GetCheckLocationMetricDataOKBodyRangeLast24Hours string = "last_24_hours"

	// GetCheckLocationMetricDataOKBodyRangeYesterday captures enum value "yesterday"
	GetCheckLocationMetricDataOKBodyRangeYesterday string = "yesterday"

	// GetCheckLocationMetricDataOKBodyRangeToday captures enum value "today"
	GetCheckLocationMetricDataOKBodyRangeToday string = "today"

	// GetCheckLocationMetricDataOKBodyRangeLast7Days captures enum value "last_7_days"
	GetCheckLocationMetricDataOKBodyRangeLast7Days string = "last_7_days"

	// GetCheckLocationMetricDataOKBodyRangeLast30Days captures enum value "last_30_days"
	GetCheckLocationMetricDataOKBodyRangeLast30Days string = "last_30_days"

	// GetCheckLocationMetricDataOKBodyRangeThisWeek captures enum value "this_week"
	GetCheckLocationMetricDataOKBodyRangeThisWeek string = "this_week"

	// GetCheckLocationMetricDataOKBodyRangeLastWeek captures enum value "last_week"
	GetCheckLocationMetricDataOKBodyRangeLastWeek string = "last_week"

	// GetCheckLocationMetricDataOKBodyRangeThisMonth captures enum value "this_month"
	GetCheckLocationMetricDataOKBodyRangeThisMonth string = "this_month"

	// GetCheckLocationMetricDataOKBodyRangeMonthToDate captures enum value "month_to_date"
	GetCheckLocationMetricDataOKBodyRangeMonthToDate string = "month_to_date"

	// GetCheckLocationMetricDataOKBodyRangeLastMonth captures enum value "last_month"
	GetCheckLocationMetricDataOKBodyRangeLastMonth string = "last_month"

	// GetCheckLocationMetricDataOKBodyRangeLast3Months captures enum value "last_3_months"
	GetCheckLocationMetricDataOKBodyRangeLast3Months string = "last_3_months"

	// GetCheckLocationMetricDataOKBodyRangeLast6Months captures enum value "last_6_months"
	GetCheckLocationMetricDataOKBodyRangeLast6Months string = "last_6_months"

	// GetCheckLocationMetricDataOKBodyRangeLast12Months captures enum value "last_12_months"
	GetCheckLocationMetricDataOKBodyRangeLast12Months string = "last_12_months"
)

// prop value enum
func (o *GetCheckLocationMetricDataOKBody) validateRangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCheckLocationMetricDataOKBodyTypeRangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBody) validateRange(formats strfmt.Registry) error {
	if swag.IsZero(o.Range) { // not required
		return nil
	}

	// value enum
	if err := o.validateRangeEnum("getCheckLocationMetricDataOK"+"."+"range", "body", o.Range); err != nil {
		return err
	}

	return nil
}

func (o *GetCheckLocationMetricDataOKBody) validateSeries(formats strfmt.Registry) error {

	if err := validate.Required("getCheckLocationMetricDataOK"+"."+"series", "body", o.Series); err != nil {
		return err
	}

	for i := 0; i < len(o.Series); i++ {
		if swag.IsZero(o.Series[i]) { // not required
			continue
		}

		if o.Series[i] != nil {
			if err := o.Series[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCheckLocationMetricDataOK" + "." + "series" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetCheckLocationMetricDataOKBody) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("getCheckLocationMetricDataOK"+"."+"to", "body", o.To); err != nil {
		return err
	}

	if err := validate.FormatOf("getCheckLocationMetricDataOK"+"."+"to", "body", "date-time", o.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get check location metric data o k body based on the context it is used
func (o *GetCheckLocationMetricDataOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSeries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBody) contextValidateSeries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Series); i++ {

		if o.Series[i] != nil {
			if err := o.Series[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCheckLocationMetricDataOK" + "." + "series" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBody) UnmarshalBinary(b []byte) error {
	var res GetCheckLocationMetricDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCheckLocationMetricDataOKBodySeriesItems0 The data for one location
swagger:model GetCheckLocationMetricDataOKBodySeriesItems0
*/
type GetCheckLocationMetricDataOKBodySeriesItems0 struct {

	// The unique ID for the location
	ID int32 `json:"id,omitempty"`

	// The name of the location
	Name string `json:"name,omitempty"`

	// A readable code representing the location
	RegionCode string `json:"region_code,omitempty"`

	// The region the location is in
	WorldRegion string `json:"world_region,omitempty"`

	// metrics
	Metrics []*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0 `json:"metrics"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCheckLocationMetricDataOKBodySeriesItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID int32 `json:"id,omitempty"`

		Name string `json:"name,omitempty"`

		RegionCode string `json:"region_code,omitempty"`

		WorldRegion string `json:"world_region,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.ID = dataAO0.ID

	o.Name = dataAO0.Name

	o.RegionCode = dataAO0.RegionCode

	o.WorldRegion = dataAO0.WorldRegion

	// AO1
	var dataAO1 struct {
		Metrics []*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0 `json:"metrics"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.Metrics = dataAO1.Metrics

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCheckLocationMetricDataOKBodySeriesItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID int32 `json:"id,omitempty"`

		Name string `json:"name,omitempty"`

		RegionCode string `json:"region_code,omitempty"`

		WorldRegion string `json:"world_region,omitempty"`
	}

	dataAO0.ID = o.ID

	dataAO0.Name = o.Name

	dataAO0.RegionCode = o.RegionCode

	dataAO0.WorldRegion = o.WorldRegion

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		Metrics []*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0 `json:"metrics"`
	}

	dataAO1.Metrics = o.Metrics

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get check location metric data o k body series items0
func (o *GetCheckLocationMetricDataOKBodySeriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(o.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(o.Metrics); i++ {
		if swag.IsZero(o.Metrics[i]) { // not required
			continue
		}

		if o.Metrics[i] != nil {
			if err := o.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get check location metric data o k body series items0 based on the context it is used
func (o *GetCheckLocationMetricDataOKBodySeriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Metrics); i++ {

		if o.Metrics[i] != nil {
			if err := o.Metrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBodySeriesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBodySeriesItems0) UnmarshalBinary(b []byte) error {
	var res GetCheckLocationMetricDataOKBodySeriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0 The metric data for the location
swagger:model GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0
*/
type GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0 struct {

	// The format of the data for this metric.
	// Example: percent
	// Enum: [milliseconds count percent]
	Format string `json:"format,omitempty"`

	// A readable label for the metric, in Title Case.
	// Example: Percentage Uptime
	// Enum: [Mean Response Time Maximum Response Time Minimum Response Time Response Time Standard Deviation Run Count Error Count Percentage Uptime Percentage Downtime SLA Percentage]
	Label string `json:"label,omitempty"`

	// The name of the metric, in snake_case.
	// Example: percentage_uptime
	// Enum: [average_response_time max_response_time min_response_time standard_deviation run_count error_count percentage_uptime percentage_downtime sla_percentage]
	Name string `json:"name,omitempty"`

	// An array of data points
	// Min Items: 0
	Points []*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0 `json:"points"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Format string `json:"format,omitempty"`

		Label string `json:"label,omitempty"`

		Name string `json:"name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.Format = dataAO0.Format

	o.Label = dataAO0.Label

	o.Name = dataAO0.Name

	// AO1
	var dataAO1 struct {
		Points []*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0 `json:"points"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.Points = dataAO1.Points

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Format string `json:"format,omitempty"`

		Label string `json:"label,omitempty"`

		Name string `json:"name,omitempty"`
	}

	dataAO0.Format = o.Format

	dataAO0.Label = o.Label

	dataAO0.Name = o.Name

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	var dataAO1 struct {
		Points []*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0 `json:"points"`
	}

	dataAO1.Points = o.Points

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get check location metric data o k body series items0 metrics items0
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["milliseconds","count","percent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeFormatPropEnum = append(getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeFormatPropEnum, v)
	}
}

// property enum
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(o.Format) { // not required
		return nil
	}

	// value enum
	if err := o.validateFormatEnum("format", "body", o.Format); err != nil {
		return err
	}

	return nil
}

var getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Mean Response Time","Maximum Response Time","Minimum Response Time","Response Time Standard Deviation","Run Count","Error Count","Percentage Uptime","Percentage Downtime","SLA Percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeLabelPropEnum = append(getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeLabelPropEnum, v)
	}
}

// property enum
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(o.Label) { // not required
		return nil
	}

	// value enum
	if err := o.validateLabelEnum("label", "body", o.Label); err != nil {
		return err
	}

	return nil
}

var getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["average_response_time","max_response_time","min_response_time","standard_deviation","run_count","error_count","percentage_uptime","percentage_downtime","sla_percentage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeNamePropEnum = append(getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeNamePropEnum, v)
	}
}

// property enum
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCheckLocationMetricDataOKBodySeriesItems0MetricsItems0TypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	// value enum
	if err := o.validateNameEnum("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) validatePoints(formats strfmt.Registry) error {

	if swag.IsZero(o.Points) { // not required
		return nil
	}

	iPointsSize := int64(len(o.Points))

	if err := validate.MinItems("points", "body", iPointsSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(o.Points); i++ {
		if swag.IsZero(o.Points[i]) { // not required
			continue
		}

		if o.Points[i] != nil {
			if err := o.Points[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get check location metric data o k body series items0 metrics items0 based on the context it is used
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Points); i++ {

		if o.Points[i] != nil {
			if err := o.Points[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0) UnmarshalBinary(b []byte) error {
	var res GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0 A single data point. May be a single run or the mean across multiple runs.
swagger:model GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0
*/
type GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0 struct {

	// The start timestamp for the data point (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	From strfmt.DateTime `json:"from,omitempty"`

	// The end timestamp for the data point (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty"`

	// The value for the data point. May be run-level or aggregate data.
	// Example: 99.3
	Value float64 `json:"value,omitempty"`
}

// Validate validates this get check location metric data o k body series items0 metrics items0 points items0
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(o.From) { // not required
		return nil
	}

	if err := validate.FormatOf("from", "body", "date-time", o.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(o.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", o.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get check location metric data o k body series items0 metrics items0 points items0 based on context it is used
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0) UnmarshalBinary(b []byte) error {
	var res GetCheckLocationMetricDataOKBodySeriesItems0MetricsItems0PointsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
