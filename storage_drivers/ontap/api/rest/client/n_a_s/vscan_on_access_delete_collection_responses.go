// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// VscanOnAccessDeleteCollectionReader is a Reader for the VscanOnAccessDeleteCollection structure.
type VscanOnAccessDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnAccessDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnAccessDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnAccessDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnAccessDeleteCollectionOK creates a VscanOnAccessDeleteCollectionOK with default headers values
func NewVscanOnAccessDeleteCollectionOK() *VscanOnAccessDeleteCollectionOK {
	return &VscanOnAccessDeleteCollectionOK{}
}

/*
VscanOnAccessDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnAccessDeleteCollectionOK struct {
}

// IsSuccess returns true when this vscan on access delete collection o k response has a 2xx status code
func (o *VscanOnAccessDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on access delete collection o k response has a 3xx status code
func (o *VscanOnAccessDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on access delete collection o k response has a 4xx status code
func (o *VscanOnAccessDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on access delete collection o k response has a 5xx status code
func (o *VscanOnAccessDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on access delete collection o k response a status code equal to that given
func (o *VscanOnAccessDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan on access delete collection o k response
func (o *VscanOnAccessDeleteCollectionOK) Code() int {
	return 200
}

func (o *VscanOnAccessDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessDeleteCollectionOK", 200)
}

func (o *VscanOnAccessDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessDeleteCollectionOK", 200)
}

func (o *VscanOnAccessDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanOnAccessDeleteCollectionDefault creates a VscanOnAccessDeleteCollectionDefault with default headers values
func NewVscanOnAccessDeleteCollectionDefault(code int) *VscanOnAccessDeleteCollectionDefault {
	return &VscanOnAccessDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	VscanOnAccessDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027034   | An On-Access policy associated with a data SVM which was created by SVM owned by the cluster cannot be deleted. |
| 10027040   | An On-Access policy with a status enabled cannot be deleted. Disable the policy and then delete the policy. |
*/
type VscanOnAccessDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan on access delete collection default response has a 2xx status code
func (o *VscanOnAccessDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on access delete collection default response has a 3xx status code
func (o *VscanOnAccessDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on access delete collection default response has a 4xx status code
func (o *VscanOnAccessDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on access delete collection default response has a 5xx status code
func (o *VscanOnAccessDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on access delete collection default response a status code equal to that given
func (o *VscanOnAccessDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan on access delete collection default response
func (o *VscanOnAccessDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *VscanOnAccessDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_delete_collection default %s", o._statusCode, payload)
}

func (o *VscanOnAccessDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_delete_collection default %s", o._statusCode, payload)
}

func (o *VscanOnAccessDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnAccessDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VscanOnAccessDeleteCollectionBody vscan on access delete collection body
swagger:model VscanOnAccessDeleteCollectionBody
*/
type VscanOnAccessDeleteCollectionBody struct {

	// vscan on access response inline records
	VscanOnAccessResponseInlineRecords []*models.VscanOnAccess `json:"records,omitempty"`
}

// Validate validates this vscan on access delete collection body
func (o *VscanOnAccessDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVscanOnAccessResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessDeleteCollectionBody) validateVscanOnAccessResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.VscanOnAccessResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.VscanOnAccessResponseInlineRecords); i++ {
		if swag.IsZero(o.VscanOnAccessResponseInlineRecords[i]) { // not required
			continue
		}

		if o.VscanOnAccessResponseInlineRecords[i] != nil {
			if err := o.VscanOnAccessResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vscan on access delete collection body based on the context it is used
func (o *VscanOnAccessDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVscanOnAccessResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanOnAccessDeleteCollectionBody) contextValidateVscanOnAccessResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VscanOnAccessResponseInlineRecords); i++ {

		if o.VscanOnAccessResponseInlineRecords[i] != nil {
			if err := o.VscanOnAccessResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *VscanOnAccessDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanOnAccessDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res VscanOnAccessDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
