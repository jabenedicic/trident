// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetgroupsSettingsModifyCollectionReader is a Reader for the NetgroupsSettingsModifyCollection structure.
type NetgroupsSettingsModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetgroupsSettingsModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetgroupsSettingsModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetgroupsSettingsModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetgroupsSettingsModifyCollectionOK creates a NetgroupsSettingsModifyCollectionOK with default headers values
func NewNetgroupsSettingsModifyCollectionOK() *NetgroupsSettingsModifyCollectionOK {
	return &NetgroupsSettingsModifyCollectionOK{}
}

/*
NetgroupsSettingsModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type NetgroupsSettingsModifyCollectionOK struct {
}

// IsSuccess returns true when this netgroups settings modify collection o k response has a 2xx status code
func (o *NetgroupsSettingsModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this netgroups settings modify collection o k response has a 3xx status code
func (o *NetgroupsSettingsModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this netgroups settings modify collection o k response has a 4xx status code
func (o *NetgroupsSettingsModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this netgroups settings modify collection o k response has a 5xx status code
func (o *NetgroupsSettingsModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this netgroups settings modify collection o k response a status code equal to that given
func (o *NetgroupsSettingsModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the netgroups settings modify collection o k response
func (o *NetgroupsSettingsModifyCollectionOK) Code() int {
	return 200
}

func (o *NetgroupsSettingsModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/cache/netgroup/settings][%d] netgroupsSettingsModifyCollectionOK", 200)
}

func (o *NetgroupsSettingsModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /name-services/cache/netgroup/settings][%d] netgroupsSettingsModifyCollectionOK", 200)
}

func (o *NetgroupsSettingsModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetgroupsSettingsModifyCollectionDefault creates a NetgroupsSettingsModifyCollectionDefault with default headers values
func NewNetgroupsSettingsModifyCollectionDefault(code int) *NetgroupsSettingsModifyCollectionDefault {
	return &NetgroupsSettingsModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	NetgroupsSettingsModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 23724055 | Internal error. Configuration for Vserver failed. Verify that the cluster is healthy, then try the command again. For further assistance, contact technical support. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetgroupsSettingsModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this netgroups settings modify collection default response has a 2xx status code
func (o *NetgroupsSettingsModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this netgroups settings modify collection default response has a 3xx status code
func (o *NetgroupsSettingsModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this netgroups settings modify collection default response has a 4xx status code
func (o *NetgroupsSettingsModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this netgroups settings modify collection default response has a 5xx status code
func (o *NetgroupsSettingsModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this netgroups settings modify collection default response a status code equal to that given
func (o *NetgroupsSettingsModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the netgroups settings modify collection default response
func (o *NetgroupsSettingsModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *NetgroupsSettingsModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/cache/netgroup/settings][%d] netgroups_settings_modify_collection default %s", o._statusCode, payload)
}

func (o *NetgroupsSettingsModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /name-services/cache/netgroup/settings][%d] netgroups_settings_modify_collection default %s", o._statusCode, payload)
}

func (o *NetgroupsSettingsModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetgroupsSettingsModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetgroupsSettingsModifyCollectionBody netgroups settings modify collection body
swagger:model NetgroupsSettingsModifyCollectionBody
*/
type NetgroupsSettingsModifyCollectionBody struct {

	// links
	Links *models.NetgroupsSettingsInlineLinks `json:"_links,omitempty"`

	// Indicates whether or not the cache is enabled.
	//
	Enabled *bool `json:"enabled,omitempty"`

	// Indicates whether or not the negative cache is enabled.
	//
	NegativeCacheEnabledByhost *bool `json:"negative_cache_enabled_byhost,omitempty"`

	// Specifies negative Time to Live by host, in ISO 8601 format.
	//
	// Example: PT2M5S
	NegativeTTLByhost *string `json:"negative_ttl_byhost,omitempty"`

	// netgroups settings response inline records
	NetgroupsSettingsResponseInlineRecords []*models.NetgroupsSettings `json:"records,omitempty"`

	// svm
	Svm *models.NetgroupsSettingsInlineSvm `json:"svm,omitempty"`

	// Specifies Time to Live by host, in ISO 8601 format.
	//
	// Example: PT24H
	TTLByhost *string `json:"ttl_byhost,omitempty"`

	// Specifies Time to Live for netgroup members, in ISO 8601 format.
	//
	// Example: PT2M
	TTLForMembers *string `json:"ttl_for_members,omitempty"`
}

// Validate validates this netgroups settings modify collection body
func (o *NetgroupsSettingsModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetgroupsSettingsResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *NetgroupsSettingsModifyCollectionBody) validateNetgroupsSettingsResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.NetgroupsSettingsResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.NetgroupsSettingsResponseInlineRecords); i++ {
		if swag.IsZero(o.NetgroupsSettingsResponseInlineRecords[i]) { // not required
			continue
		}

		if o.NetgroupsSettingsResponseInlineRecords[i] != nil {
			if err := o.NetgroupsSettingsResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetgroupsSettingsModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(o.Svm) { // not required
		return nil
	}

	if o.Svm != nil {
		if err := o.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this netgroups settings modify collection body based on the context it is used
func (o *NetgroupsSettingsModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNetgroupsSettingsResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *NetgroupsSettingsModifyCollectionBody) contextValidateNetgroupsSettingsResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.NetgroupsSettingsResponseInlineRecords); i++ {

		if o.NetgroupsSettingsResponseInlineRecords[i] != nil {
			if err := o.NetgroupsSettingsResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *NetgroupsSettingsModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if o.Svm != nil {
		if err := o.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetgroupsSettingsModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetgroupsSettingsModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res NetgroupsSettingsModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetgroupsSettingsInlineLinks netgroups settings inline links
swagger:model netgroups_settings_inline__links
*/
type NetgroupsSettingsInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this netgroups settings inline links
func (o *NetgroupsSettingsInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this netgroups settings inline links based on the context it is used
func (o *NetgroupsSettingsInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetgroupsSettingsInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetgroupsSettingsInlineLinks) UnmarshalBinary(b []byte) error {
	var res NetgroupsSettingsInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetgroupsSettingsInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model netgroups_settings_inline_svm
*/
type NetgroupsSettingsInlineSvm struct {

	// links
	Links *models.NetgroupsSettingsInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this netgroups settings inline svm
func (o *NetgroupsSettingsInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this netgroups settings inline svm based on the context it is used
func (o *NetgroupsSettingsInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetgroupsSettingsInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetgroupsSettingsInlineSvm) UnmarshalBinary(b []byte) error {
	var res NetgroupsSettingsInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetgroupsSettingsInlineSvmInlineLinks netgroups settings inline svm inline links
swagger:model netgroups_settings_inline_svm_inline__links
*/
type NetgroupsSettingsInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this netgroups settings inline svm inline links
func (o *NetgroupsSettingsInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this netgroups settings inline svm inline links based on the context it is used
func (o *NetgroupsSettingsInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetgroupsSettingsInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetgroupsSettingsInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetgroupsSettingsInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res NetgroupsSettingsInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
