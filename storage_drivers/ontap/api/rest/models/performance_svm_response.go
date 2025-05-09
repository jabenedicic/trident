// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PerformanceSvmResponse performance svm response
//
// swagger:model performance_svm_response
type PerformanceSvmResponse struct {

	// links
	Links *PerformanceSvmResponseInlineLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// performance svm response inline records
	PerformanceSvmResponseInlineRecords []*PerformanceSvmResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`
}

// Validate validates this performance svm response
func (m *PerformanceSvmResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceSvmResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponse) validatePerformanceSvmResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceSvmResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.PerformanceSvmResponseInlineRecords); i++ {
		if swag.IsZero(m.PerformanceSvmResponseInlineRecords[i]) { // not required
			continue
		}

		if m.PerformanceSvmResponseInlineRecords[i] != nil {
			if err := m.PerformanceSvmResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance svm response based on the context it is used
func (m *PerformanceSvmResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceSvmResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponse) contextValidatePerformanceSvmResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerformanceSvmResponseInlineRecords); i++ {

		if m.PerformanceSvmResponseInlineRecords[i] != nil {
			if err := m.PerformanceSvmResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmResponse) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmResponseInlineLinks performance svm response inline links
//
// swagger:model performance_svm_response_inline__links
type PerformanceSvmResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance svm response inline links
func (m *PerformanceSvmResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance svm response inline links based on the context it is used
func (m *PerformanceSvmResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmResponseInlineRecordsInlineArrayItem Performance numbers, such as IOPS latency and throughput, for SVM protocols.
//
// swagger:model performance_svm_response_inline_records_inline_array_item
type PerformanceSvmResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: ["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]
	Duration *string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency `json:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: ["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25 11:20:13
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance svm response inline records inline array item
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

var performanceSvmResponseInlineRecordsInlineArrayItemTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceSvmResponseInlineRecordsInlineArrayItemTypeDurationPropEnum = append(performanceSvmResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT15S captures enum value "PT15S"
	PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT4M captures enum value "PT4M"
	PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT30M captures enum value "PT30M"
	PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT2H captures enum value "PT2H"
	PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemDurationP1D captures enum value "P1D"
	PerformanceSvmResponseInlineRecordsInlineArrayItemDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT5M captures enum value "PT5M"
	PerformanceSvmResponseInlineRecordsInlineArrayItemDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceSvmResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceSvmResponseInlineRecordsInlineArrayItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceSvmResponseInlineRecordsInlineArrayItemTypeStatusPropEnum = append(performanceSvmResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// ok
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusOk captures enum value "ok"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusOk string = "ok"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// error
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusError captures enum value "error"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusError string = "error"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialNoData captures enum value "partial_no_data"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusNegativeDelta captures enum value "negative_delta"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// not_found
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusNotFound captures enum value "not_found"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusNotFound string = "not_found"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusBackfilledData captures enum value "backfilled_data"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusInconsistentOldData string = "inconsistent_old_data"

	// BEGIN DEBUGGING
	// performance_svm_response_inline_records_inline_array_item
	// PerformanceSvmResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceSvmResponseInlineRecordsInlineArrayItemStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceSvmResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance svm response inline records inline array item based on the context it is used
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {
		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {
		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_svm_response_inline_records_inline_array_item_inline_iops
type PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance svm response inline records inline array item inline iops
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm response inline records inline array item inline iops based on the context it is used
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmResponseInlineRecordsInlineArrayItemInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_svm_response_inline_records_inline_array_item_inline_latency
type PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance svm response inline records inline array item inline latency
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm response inline records inline array item inline latency based on the context it is used
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks performance svm response inline records inline array item inline links
//
// swagger:model performance_svm_response_inline_records_inline_array_item_inline__links
type PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance svm response inline records inline array item inline links
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance svm response inline records inline array item inline links based on the context it is used
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_svm_response_inline_records_inline_array_item_inline_throughput
type PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Performance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance svm response inline records inline array item inline throughput
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance svm response inline records inline array item inline throughput based on the context it is used
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceSvmResponseInlineRecordsInlineArrayItemInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
