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

// DuogroupModifyReader is a Reader for the DuogroupModify structure.
type DuogroupModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DuogroupModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDuogroupModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDuogroupModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDuogroupModifyOK creates a DuogroupModifyOK with default headers values
func NewDuogroupModifyOK() *DuogroupModifyOK {
	return &DuogroupModifyOK{}
}

/*
DuogroupModifyOK describes a response with status code 200, with default header values.

OK
*/
type DuogroupModifyOK struct {
}

// IsSuccess returns true when this duogroup modify o k response has a 2xx status code
func (o *DuogroupModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this duogroup modify o k response has a 3xx status code
func (o *DuogroupModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this duogroup modify o k response has a 4xx status code
func (o *DuogroupModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this duogroup modify o k response has a 5xx status code
func (o *DuogroupModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this duogroup modify o k response a status code equal to that given
func (o *DuogroupModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the duogroup modify o k response
func (o *DuogroupModifyOK) Code() int {
	return 200
}

func (o *DuogroupModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroupModifyOK", 200)
}

func (o *DuogroupModifyOK) String() string {
	return fmt.Sprintf("[PATCH /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroupModifyOK", 200)
}

func (o *DuogroupModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDuogroupModifyDefault creates a DuogroupModifyDefault with default headers values
func NewDuogroupModifyDefault(code int) *DuogroupModifyDefault {
	return &DuogroupModifyDefault{
		_statusCode: code,
	}
}

/*
DuogroupModifyDefault describes a response with status code -1, with default header values.

Error
*/
type DuogroupModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this duogroup modify default response has a 2xx status code
func (o *DuogroupModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this duogroup modify default response has a 3xx status code
func (o *DuogroupModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this duogroup modify default response has a 4xx status code
func (o *DuogroupModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this duogroup modify default response has a 5xx status code
func (o *DuogroupModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this duogroup modify default response a status code equal to that given
func (o *DuogroupModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the duogroup modify default response
func (o *DuogroupModifyDefault) Code() int {
	return o._statusCode
}

func (o *DuogroupModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroup_modify default %s", o._statusCode, payload)
}

func (o *DuogroupModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/authentication/duo/groups/{owner.uuid}/{name}][%d] duogroup_modify default %s", o._statusCode, payload)
}

func (o *DuogroupModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DuogroupModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
