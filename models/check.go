// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Check A single check
//
// swagger:discriminator check type
type Check interface {
	runtime.Validatable
	runtime.ContextValidatable

	// When the check was created (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	CreatedAt() strfmt.DateTime
	SetCreatedAt(strfmt.DateTime)

	// Run the check at this interval (in minutes)
	// Example: 5
	Frequency() int64
	SetFrequency(int64)

	// The unique ID for the check
	// Example: 1
	// Required: true
	ID() *int32
	SetID(*int32)

	// links
	Links() *CheckLinks
	SetLinks(*CheckLinks)

	// If notifications for this check are muted or not
	// Example: false
	Muted() bool
	SetMuted(bool)

	// The unique name for the check
	// Example: Example Check
	Name() string
	SetName(string)

	// If the check is paused or not
	// Example: false
	Paused() bool
	SetPaused(bool)

	// status
	Status() *Status
	SetStatus(*Status)

	// An array of tags applied to the check
	Tags() []*Tag
	SetTags([]*Tag)

	// The check type
	// Required: true
	// Enum: [http http_multi_step port real_browser benchmark content uptime monitoring api]
	Type() string
	SetType(string)

	// When the check was last updated (UTC)
	// Example: 2021-05-25T17:54:05Z
	// Format: date-time
	UpdatedAt() strfmt.DateTime
	SetUpdatedAt(strfmt.DateTime)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type check struct {
	createdAtField strfmt.DateTime

	frequencyField int64

	idField *int32

	linksField *CheckLinks

	mutedField bool

	nameField string

	pausedField bool

	statusField *Status

	tagsField []*Tag

	typeField string

	updatedAtField strfmt.DateTime
}

// CreatedAt gets the created at of this polymorphic type
func (m *check) CreatedAt() strfmt.DateTime {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this polymorphic type
func (m *check) SetCreatedAt(val strfmt.DateTime) {
	m.createdAtField = val
}

// Frequency gets the frequency of this polymorphic type
func (m *check) Frequency() int64 {
	return m.frequencyField
}

// SetFrequency sets the frequency of this polymorphic type
func (m *check) SetFrequency(val int64) {
	m.frequencyField = val
}

// ID gets the id of this polymorphic type
func (m *check) ID() *int32 {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *check) SetID(val *int32) {
	m.idField = val
}

// Links gets the links of this polymorphic type
func (m *check) Links() *CheckLinks {
	return m.linksField
}

// SetLinks sets the links of this polymorphic type
func (m *check) SetLinks(val *CheckLinks) {
	m.linksField = val
}

// Muted gets the muted of this polymorphic type
func (m *check) Muted() bool {
	return m.mutedField
}

// SetMuted sets the muted of this polymorphic type
func (m *check) SetMuted(val bool) {
	m.mutedField = val
}

// Name gets the name of this polymorphic type
func (m *check) Name() string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *check) SetName(val string) {
	m.nameField = val
}

// Paused gets the paused of this polymorphic type
func (m *check) Paused() bool {
	return m.pausedField
}

// SetPaused sets the paused of this polymorphic type
func (m *check) SetPaused(val bool) {
	m.pausedField = val
}

// Status gets the status of this polymorphic type
func (m *check) Status() *Status {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *check) SetStatus(val *Status) {
	m.statusField = val
}

// Tags gets the tags of this polymorphic type
func (m *check) Tags() []*Tag {
	return m.tagsField
}

// SetTags sets the tags of this polymorphic type
func (m *check) SetTags(val []*Tag) {
	m.tagsField = val
}

// Type gets the type of this polymorphic type
func (m *check) Type() string {
	return "check"
}

// SetType sets the type of this polymorphic type
func (m *check) SetType(val string) {
}

// UpdatedAt gets the updated at of this polymorphic type
func (m *check) UpdatedAt() strfmt.DateTime {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this polymorphic type
func (m *check) SetUpdatedAt(val strfmt.DateTime) {
	m.updatedAtField = val
}

// UnmarshalCheckSlice unmarshals polymorphic slices of Check
func UnmarshalCheckSlice(reader io.Reader, consumer runtime.Consumer) ([]Check, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Check
	for _, element := range elements {
		obj, err := unmarshalCheck(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCheck unmarshals polymorphic Check
func UnmarshalCheck(reader io.Reader, consumer runtime.Consumer) (Check, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCheck(data, consumer)
}

func unmarshalCheck(data []byte, consumer runtime.Consumer) (Check, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "check":
		var result check
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "monitoring_check":
		var result MonitoringCheck
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)
}

// Validate validates this check
func (m *check) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *check) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt()) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *check) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	return nil
}

func (m *check) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links()) { // not required
		return nil
	}

	if m.Links() != nil {
		if err := m.Links().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *check) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status()) { // not required
		return nil
	}

	if m.Status() != nil {
		if err := m.Status().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *check) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags()) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags()); i++ {
		if swag.IsZero(m.tagsField[i]) { // not required
			continue
		}

		if m.tagsField[i] != nil {
			if err := m.tagsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var checkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","http_multi_step","port","real_browser","benchmark","content","uptime","monitoring","all","api"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checkTypeTypePropEnum = append(checkTypeTypePropEnum, v)
	}
}

const (

	// CheckTypeHTTP captures enum value "http"
	CheckTypeHTTP string = "http"

	// CheckTypeHTTPMultiStep captures enum value "http_multi_step"
	CheckTypeHTTPMultiStep string = "http_multi_step"

	// CheckTypePort captures enum value "port"
	CheckTypePort string = "port"

	// CheckTypeRealBrowser captures enum value "real_browser"
	CheckTypeRealBrowser string = "real_browser"

	// CheckTypeBenchmark captures enum value "benchmark"
	CheckTypeBenchmark string = "benchmark"

	// CheckTypeContent captures enum value "content"
	CheckTypeContent string = "content"

	// CheckTypeUptime captures enum value "uptime"
	CheckTypeUptime string = "uptime"

	// CheckTypeMonitoring captures enum value "monitoring"
	CheckTypeMonitoring string = "monitoring"

	// CheckTypeAPI captures enum value "all"
	CheckTypeAll string = "all"

	// CheckTypeAPI captures enum value "api"
	CheckTypeAPI string = "api"
)

// prop value enum
func (m *check) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, checkTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *check) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt()) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt().String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this check based on the context it is used
func (m *check) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *check) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links() != nil {
		if err := m.Links().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *check) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status() != nil {
		if err := m.Status().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *check) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags()); i++ {

		if m.tagsField[i] != nil {
			if err := m.tagsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// CheckLinks check links
//
// swagger:model CheckLinks
type CheckLinks struct {

	// The URL for the last run of this check
	// Example: https://monitoring.rigor.com/checks/1/runs/1
	LastRun string `json:"last_run,omitempty"`

	// The URL for the available metrics for this check
	// Example: https://monitoring-api.rigor.com/v2/checks/1/metrics
	Metrics string `json:"metrics,omitempty"`

	// The URL for the check detail
	// Example: https://monitoring-api.rigor.com/v2/checks/1
	Self string `json:"self,omitempty"`

	// The URL for the HTML view for this check
	// Example: https://monitoring.rigor.com/checks/http/1
	SelfHTML string `json:"self_html,omitempty"`
}

// Validate validates this check links
func (m *CheckLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check links based on context it is used
func (m *CheckLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckLinks) UnmarshalBinary(b []byte) error {
	var res CheckLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
