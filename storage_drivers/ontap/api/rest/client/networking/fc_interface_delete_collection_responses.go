// Code generated by go-swagger; DO NOT EDIT.

package networking

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcInterfaceDeleteCollectionReader is a Reader for the FcInterfaceDeleteCollection structure.
type FcInterfaceDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcInterfaceDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcInterfaceDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcInterfaceDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcInterfaceDeleteCollectionOK creates a FcInterfaceDeleteCollectionOK with default headers values
func NewFcInterfaceDeleteCollectionOK() *FcInterfaceDeleteCollectionOK {
	return &FcInterfaceDeleteCollectionOK{}
}

/*
FcInterfaceDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type FcInterfaceDeleteCollectionOK struct {
}

// IsSuccess returns true when this fc interface delete collection o k response has a 2xx status code
func (o *FcInterfaceDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fc interface delete collection o k response has a 3xx status code
func (o *FcInterfaceDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fc interface delete collection o k response has a 4xx status code
func (o *FcInterfaceDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fc interface delete collection o k response has a 5xx status code
func (o *FcInterfaceDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fc interface delete collection o k response a status code equal to that given
func (o *FcInterfaceDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fc interface delete collection o k response
func (o *FcInterfaceDeleteCollectionOK) Code() int {
	return 200
}

func (o *FcInterfaceDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /network/fc/interfaces][%d] fcInterfaceDeleteCollectionOK", 200)
}

func (o *FcInterfaceDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /network/fc/interfaces][%d] fcInterfaceDeleteCollectionOK", 200)
}

func (o *FcInterfaceDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFcInterfaceDeleteCollectionDefault creates a FcInterfaceDeleteCollectionDefault with default headers values
func NewFcInterfaceDeleteCollectionDefault(code int) *FcInterfaceDeleteCollectionDefault {
	return &FcInterfaceDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	FcInterfaceDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 53280992 | The FC interface could not be deleted because it is enabled. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type FcInterfaceDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this fc interface delete collection default response has a 2xx status code
func (o *FcInterfaceDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fc interface delete collection default response has a 3xx status code
func (o *FcInterfaceDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fc interface delete collection default response has a 4xx status code
func (o *FcInterfaceDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fc interface delete collection default response has a 5xx status code
func (o *FcInterfaceDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fc interface delete collection default response a status code equal to that given
func (o *FcInterfaceDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the fc interface delete collection default response
func (o *FcInterfaceDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *FcInterfaceDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/fc/interfaces][%d] fc_interface_delete_collection default %s", o._statusCode, payload)
}

func (o *FcInterfaceDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /network/fc/interfaces][%d] fc_interface_delete_collection default %s", o._statusCode, payload)
}

func (o *FcInterfaceDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcInterfaceDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
FcInterfaceDeleteCollectionBody fc interface delete collection body
swagger:model FcInterfaceDeleteCollectionBody
*/
type FcInterfaceDeleteCollectionBody struct {

	// fc interface response inline records
	FcInterfaceResponseInlineRecords []*models.FcInterface `json:"records,omitempty"`

	// recommend
	Recommend *models.FcInterfaceResponseInlineRecommend `json:"recommend,omitempty"`
}

// Validate validates this fc interface delete collection body
func (o *FcInterfaceDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFcInterfaceResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecommend(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FcInterfaceDeleteCollectionBody) validateFcInterfaceResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.FcInterfaceResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.FcInterfaceResponseInlineRecords); i++ {
		if swag.IsZero(o.FcInterfaceResponseInlineRecords[i]) { // not required
			continue
		}

		if o.FcInterfaceResponseInlineRecords[i] != nil {
			if err := o.FcInterfaceResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *FcInterfaceDeleteCollectionBody) validateRecommend(formats strfmt.Registry) error {
	if swag.IsZero(o.Recommend) { // not required
		return nil
	}

	if o.Recommend != nil {
		if err := o.Recommend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "recommend")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc interface delete collection body based on the context it is used
func (o *FcInterfaceDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFcInterfaceResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRecommend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FcInterfaceDeleteCollectionBody) contextValidateFcInterfaceResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FcInterfaceResponseInlineRecords); i++ {

		if o.FcInterfaceResponseInlineRecords[i] != nil {
			if err := o.FcInterfaceResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *FcInterfaceDeleteCollectionBody) contextValidateRecommend(ctx context.Context, formats strfmt.Registry) error {

	if o.Recommend != nil {
		if err := o.Recommend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "recommend")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FcInterfaceDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FcInterfaceDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res FcInterfaceDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FcInterfaceResponseInlineRecommend Response properties specific to the FC interface placement functionality. See the _Interface placement recommendations_ section of [`DOC /network/fc/interfaces`](#docs-networking-network_fc_interfaces)
//
swagger:model fc_interface_response_inline_recommend
*/
type FcInterfaceResponseInlineRecommend struct {

	// Messages describing the results of a FC network interface placement operation or evaluation of caller-proposed locations.
	//
	// Read Only: true
	Messages []*models.FcInterfaceRecommendMessage `json:"messages,omitempty"`
}

// Validate validates this fc interface response inline recommend
func (o *FcInterfaceResponseInlineRecommend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FcInterfaceResponseInlineRecommend) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(o.Messages) { // not required
		return nil
	}

	for i := 0; i < len(o.Messages); i++ {
		if swag.IsZero(o.Messages[i]) { // not required
			continue
		}

		if o.Messages[i] != nil {
			if err := o.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "recommend" + "." + "messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fc interface response inline recommend based on the context it is used
func (o *FcInterfaceResponseInlineRecommend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FcInterfaceResponseInlineRecommend) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"recommend"+"."+"messages", "body", []*models.FcInterfaceRecommendMessage(o.Messages)); err != nil {
		return err
	}

	for i := 0; i < len(o.Messages); i++ {

		if o.Messages[i] != nil {
			if err := o.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "recommend" + "." + "messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *FcInterfaceResponseInlineRecommend) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FcInterfaceResponseInlineRecommend) UnmarshalBinary(b []byte) error {
	var res FcInterfaceResponseInlineRecommend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
