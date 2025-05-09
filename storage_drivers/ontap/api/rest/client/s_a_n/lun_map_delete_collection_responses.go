// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// LunMapDeleteCollectionReader is a Reader for the LunMapDeleteCollection structure.
type LunMapDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunMapDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapDeleteCollectionOK creates a LunMapDeleteCollectionOK with default headers values
func NewLunMapDeleteCollectionOK() *LunMapDeleteCollectionOK {
	return &LunMapDeleteCollectionOK{}
}

/*
LunMapDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type LunMapDeleteCollectionOK struct {
}

// IsSuccess returns true when this lun map delete collection o k response has a 2xx status code
func (o *LunMapDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun map delete collection o k response has a 3xx status code
func (o *LunMapDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun map delete collection o k response has a 4xx status code
func (o *LunMapDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun map delete collection o k response has a 5xx status code
func (o *LunMapDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun map delete collection o k response a status code equal to that given
func (o *LunMapDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lun map delete collection o k response
func (o *LunMapDeleteCollectionOK) Code() int {
	return 200
}

func (o *LunMapDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps][%d] lunMapDeleteCollectionOK", 200)
}

func (o *LunMapDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps][%d] lunMapDeleteCollectionOK", 200)
}

func (o *LunMapDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLunMapDeleteCollectionDefault creates a LunMapDeleteCollectionDefault with default headers values
func NewLunMapDeleteCollectionDefault(code int) *LunMapDeleteCollectionDefault {
	return &LunMapDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	LunMapDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN does not exist or is not accessible to the caller. |
| 5374878 | The specified initiator group does not exist, is not accessible to the caller, or is not in the same SVM as the specified LUN. |
| 5374922 | The specified LUN map does not exist. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type LunMapDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this lun map delete collection default response has a 2xx status code
func (o *LunMapDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun map delete collection default response has a 3xx status code
func (o *LunMapDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun map delete collection default response has a 4xx status code
func (o *LunMapDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun map delete collection default response has a 5xx status code
func (o *LunMapDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun map delete collection default response a status code equal to that given
func (o *LunMapDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lun map delete collection default response
func (o *LunMapDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *LunMapDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps][%d] lun_map_delete_collection default %s", o._statusCode, payload)
}

func (o *LunMapDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/san/lun-maps][%d] lun_map_delete_collection default %s", o._statusCode, payload)
}

func (o *LunMapDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
LunMapDeleteCollectionBody lun map delete collection body
swagger:model LunMapDeleteCollectionBody
*/
type LunMapDeleteCollectionBody struct {

	// lun map response inline records
	LunMapResponseInlineRecords []*models.LunMap `json:"records,omitempty"`
}

// Validate validates this lun map delete collection body
func (o *LunMapDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLunMapResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LunMapDeleteCollectionBody) validateLunMapResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.LunMapResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.LunMapResponseInlineRecords); i++ {
		if swag.IsZero(o.LunMapResponseInlineRecords[i]) { // not required
			continue
		}

		if o.LunMapResponseInlineRecords[i] != nil {
			if err := o.LunMapResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this lun map delete collection body based on the context it is used
func (o *LunMapDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLunMapResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LunMapDeleteCollectionBody) contextValidateLunMapResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.LunMapResponseInlineRecords); i++ {

		if o.LunMapResponseInlineRecords[i] != nil {
			if err := o.LunMapResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *LunMapDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LunMapDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res LunMapDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
