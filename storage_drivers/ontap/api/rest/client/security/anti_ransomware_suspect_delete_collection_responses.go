// Code generated by go-swagger; DO NOT EDIT.

package security

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

// AntiRansomwareSuspectDeleteCollectionReader is a Reader for the AntiRansomwareSuspectDeleteCollection structure.
type AntiRansomwareSuspectDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AntiRansomwareSuspectDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAntiRansomwareSuspectDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAntiRansomwareSuspectDeleteCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /security/anti-ransomware/suspects] anti_ransomware_suspect_delete_collection", response, response.Code())
	}
}

// NewAntiRansomwareSuspectDeleteCollectionOK creates a AntiRansomwareSuspectDeleteCollectionOK with default headers values
func NewAntiRansomwareSuspectDeleteCollectionOK() *AntiRansomwareSuspectDeleteCollectionOK {
	return &AntiRansomwareSuspectDeleteCollectionOK{}
}

/*
AntiRansomwareSuspectDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type AntiRansomwareSuspectDeleteCollectionOK struct {
	Payload *models.AntiRansomwareSuspectJobLinkResponse
}

// IsSuccess returns true when this anti ransomware suspect delete collection o k response has a 2xx status code
func (o *AntiRansomwareSuspectDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware suspect delete collection o k response has a 3xx status code
func (o *AntiRansomwareSuspectDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware suspect delete collection o k response has a 4xx status code
func (o *AntiRansomwareSuspectDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware suspect delete collection o k response has a 5xx status code
func (o *AntiRansomwareSuspectDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware suspect delete collection o k response a status code equal to that given
func (o *AntiRansomwareSuspectDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the anti ransomware suspect delete collection o k response
func (o *AntiRansomwareSuspectDeleteCollectionOK) Code() int {
	return 200
}

func (o *AntiRansomwareSuspectDeleteCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects][%d] antiRansomwareSuspectDeleteCollectionOK %s", 200, payload)
}

func (o *AntiRansomwareSuspectDeleteCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects][%d] antiRansomwareSuspectDeleteCollectionOK %s", 200, payload)
}

func (o *AntiRansomwareSuspectDeleteCollectionOK) GetPayload() *models.AntiRansomwareSuspectJobLinkResponse {
	return o.Payload
}

func (o *AntiRansomwareSuspectDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntiRansomwareSuspectJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAntiRansomwareSuspectDeleteCollectionAccepted creates a AntiRansomwareSuspectDeleteCollectionAccepted with default headers values
func NewAntiRansomwareSuspectDeleteCollectionAccepted() *AntiRansomwareSuspectDeleteCollectionAccepted {
	return &AntiRansomwareSuspectDeleteCollectionAccepted{}
}

/*
AntiRansomwareSuspectDeleteCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AntiRansomwareSuspectDeleteCollectionAccepted struct {
	Payload *models.AntiRansomwareSuspectJobLinkResponse
}

// IsSuccess returns true when this anti ransomware suspect delete collection accepted response has a 2xx status code
func (o *AntiRansomwareSuspectDeleteCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware suspect delete collection accepted response has a 3xx status code
func (o *AntiRansomwareSuspectDeleteCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware suspect delete collection accepted response has a 4xx status code
func (o *AntiRansomwareSuspectDeleteCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware suspect delete collection accepted response has a 5xx status code
func (o *AntiRansomwareSuspectDeleteCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware suspect delete collection accepted response a status code equal to that given
func (o *AntiRansomwareSuspectDeleteCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the anti ransomware suspect delete collection accepted response
func (o *AntiRansomwareSuspectDeleteCollectionAccepted) Code() int {
	return 202
}

func (o *AntiRansomwareSuspectDeleteCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects][%d] antiRansomwareSuspectDeleteCollectionAccepted %s", 202, payload)
}

func (o *AntiRansomwareSuspectDeleteCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects][%d] antiRansomwareSuspectDeleteCollectionAccepted %s", 202, payload)
}

func (o *AntiRansomwareSuspectDeleteCollectionAccepted) GetPayload() *models.AntiRansomwareSuspectJobLinkResponse {
	return o.Payload
}

func (o *AntiRansomwareSuspectDeleteCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntiRansomwareSuspectJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AntiRansomwareSuspectDeleteCollectionBody anti ransomware suspect delete collection body
swagger:model AntiRansomwareSuspectDeleteCollectionBody
*/
type AntiRansomwareSuspectDeleteCollectionBody struct {

	// anti ransomware suspect response inline records
	AntiRansomwareSuspectResponseInlineRecords []*models.AntiRansomwareSuspect `json:"records,omitempty"`
}

// Validate validates this anti ransomware suspect delete collection body
func (o *AntiRansomwareSuspectDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAntiRansomwareSuspectResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AntiRansomwareSuspectDeleteCollectionBody) validateAntiRansomwareSuspectResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.AntiRansomwareSuspectResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.AntiRansomwareSuspectResponseInlineRecords); i++ {
		if swag.IsZero(o.AntiRansomwareSuspectResponseInlineRecords[i]) { // not required
			continue
		}

		if o.AntiRansomwareSuspectResponseInlineRecords[i] != nil {
			if err := o.AntiRansomwareSuspectResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this anti ransomware suspect delete collection body based on the context it is used
func (o *AntiRansomwareSuspectDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAntiRansomwareSuspectResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AntiRansomwareSuspectDeleteCollectionBody) contextValidateAntiRansomwareSuspectResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AntiRansomwareSuspectResponseInlineRecords); i++ {

		if o.AntiRansomwareSuspectResponseInlineRecords[i] != nil {
			if err := o.AntiRansomwareSuspectResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *AntiRansomwareSuspectDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AntiRansomwareSuspectDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res AntiRansomwareSuspectDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
