// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// ExportRuleCollectionGetReader is a Reader for the ExportRuleCollectionGet structure.
type ExportRuleCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportRuleCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportRuleCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportRuleCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportRuleCollectionGetOK creates a ExportRuleCollectionGetOK with default headers values
func NewExportRuleCollectionGetOK() *ExportRuleCollectionGetOK {
	return &ExportRuleCollectionGetOK{}
}

/*
ExportRuleCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ExportRuleCollectionGetOK struct {
	Payload *models.ExportRuleResponse
}

// IsSuccess returns true when this export rule collection get o k response has a 2xx status code
func (o *ExportRuleCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export rule collection get o k response has a 3xx status code
func (o *ExportRuleCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export rule collection get o k response has a 4xx status code
func (o *ExportRuleCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export rule collection get o k response has a 5xx status code
func (o *ExportRuleCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export rule collection get o k response a status code equal to that given
func (o *ExportRuleCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export rule collection get o k response
func (o *ExportRuleCollectionGetOK) Code() int {
	return 200
}

func (o *ExportRuleCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules][%d] exportRuleCollectionGetOK %s", 200, payload)
}

func (o *ExportRuleCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules][%d] exportRuleCollectionGetOK %s", 200, payload)
}

func (o *ExportRuleCollectionGetOK) GetPayload() *models.ExportRuleResponse {
	return o.Payload
}

func (o *ExportRuleCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportRuleCollectionGetDefault creates a ExportRuleCollectionGetDefault with default headers values
func NewExportRuleCollectionGetDefault(code int) *ExportRuleCollectionGetDefault {
	return &ExportRuleCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	ExportRuleCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262194     | The expression is missing a value |
| 262197     | The value provided is invalid for the field |
| 2621462    | The specified SVM name does not exist |
| 2621519    | SVM name is invalid. The SVM name must begin with a letter or an underscore. If the SVM is of type \"sync-source\", the maximum supported length is 41. Otherwise, the maximum supported length is 47 |
| 2621643    | The specified SVM name is too long |
| 2621685    | SVM name length cannot be zero |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name |
| 6691623    | User is not authorized |
*/
type ExportRuleCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this export rule collection get default response has a 2xx status code
func (o *ExportRuleCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this export rule collection get default response has a 3xx status code
func (o *ExportRuleCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this export rule collection get default response has a 4xx status code
func (o *ExportRuleCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this export rule collection get default response has a 5xx status code
func (o *ExportRuleCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this export rule collection get default response a status code equal to that given
func (o *ExportRuleCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the export rule collection get default response
func (o *ExportRuleCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ExportRuleCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules][%d] export_rule_collection_get default %s", o._statusCode, payload)
}

func (o *ExportRuleCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/nfs/export-policies/{policy.id}/rules][%d] export_rule_collection_get default %s", o._statusCode, payload)
}

func (o *ExportRuleCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportRuleCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
