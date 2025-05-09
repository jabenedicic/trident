// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AntiRansomwareSuspectCollectionGetReader is a Reader for the AntiRansomwareSuspectCollectionGet structure.
type AntiRansomwareSuspectCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AntiRansomwareSuspectCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAntiRansomwareSuspectCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAntiRansomwareSuspectCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAntiRansomwareSuspectCollectionGetOK creates a AntiRansomwareSuspectCollectionGetOK with default headers values
func NewAntiRansomwareSuspectCollectionGetOK() *AntiRansomwareSuspectCollectionGetOK {
	return &AntiRansomwareSuspectCollectionGetOK{}
}

/*
AntiRansomwareSuspectCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type AntiRansomwareSuspectCollectionGetOK struct {
	Payload *models.AntiRansomwareSuspectResponse
}

// IsSuccess returns true when this anti ransomware suspect collection get o k response has a 2xx status code
func (o *AntiRansomwareSuspectCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware suspect collection get o k response has a 3xx status code
func (o *AntiRansomwareSuspectCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware suspect collection get o k response has a 4xx status code
func (o *AntiRansomwareSuspectCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware suspect collection get o k response has a 5xx status code
func (o *AntiRansomwareSuspectCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware suspect collection get o k response a status code equal to that given
func (o *AntiRansomwareSuspectCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the anti ransomware suspect collection get o k response
func (o *AntiRansomwareSuspectCollectionGetOK) Code() int {
	return 200
}

func (o *AntiRansomwareSuspectCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/anti-ransomware/suspects][%d] antiRansomwareSuspectCollectionGetOK %s", 200, payload)
}

func (o *AntiRansomwareSuspectCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/anti-ransomware/suspects][%d] antiRansomwareSuspectCollectionGetOK %s", 200, payload)
}

func (o *AntiRansomwareSuspectCollectionGetOK) GetPayload() *models.AntiRansomwareSuspectResponse {
	return o.Payload
}

func (o *AntiRansomwareSuspectCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntiRansomwareSuspectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAntiRansomwareSuspectCollectionGetDefault creates a AntiRansomwareSuspectCollectionGetDefault with default headers values
func NewAntiRansomwareSuspectCollectionGetDefault(code int) *AntiRansomwareSuspectCollectionGetDefault {
	return &AntiRansomwareSuspectCollectionGetDefault{
		_statusCode: code,
	}
}

/*
AntiRansomwareSuspectCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type AntiRansomwareSuspectCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this anti ransomware suspect collection get default response has a 2xx status code
func (o *AntiRansomwareSuspectCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this anti ransomware suspect collection get default response has a 3xx status code
func (o *AntiRansomwareSuspectCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this anti ransomware suspect collection get default response has a 4xx status code
func (o *AntiRansomwareSuspectCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this anti ransomware suspect collection get default response has a 5xx status code
func (o *AntiRansomwareSuspectCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this anti ransomware suspect collection get default response a status code equal to that given
func (o *AntiRansomwareSuspectCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the anti ransomware suspect collection get default response
func (o *AntiRansomwareSuspectCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *AntiRansomwareSuspectCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/anti-ransomware/suspects][%d] anti_ransomware_suspect_collection_get default %s", o._statusCode, payload)
}

func (o *AntiRansomwareSuspectCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security/anti-ransomware/suspects][%d] anti_ransomware_suspect_collection_get default %s", o._statusCode, payload)
}

func (o *AntiRansomwareSuspectCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AntiRansomwareSuspectCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
