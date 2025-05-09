// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// QosPolicyCreateReader is a Reader for the QosPolicyCreate structure.
type QosPolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewQosPolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewQosPolicyCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyCreateCreated creates a QosPolicyCreateCreated with default headers values
func NewQosPolicyCreateCreated() *QosPolicyCreateCreated {
	return &QosPolicyCreateCreated{}
}

/*
QosPolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type QosPolicyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.QosPolicyJobLinkResponse
}

// IsSuccess returns true when this qos policy create created response has a 2xx status code
func (o *QosPolicyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy create created response has a 3xx status code
func (o *QosPolicyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy create created response has a 4xx status code
func (o *QosPolicyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy create created response has a 5xx status code
func (o *QosPolicyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy create created response a status code equal to that given
func (o *QosPolicyCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the qos policy create created response
func (o *QosPolicyCreateCreated) Code() int {
	return 201
}

func (o *QosPolicyCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qosPolicyCreateCreated %s", 201, payload)
}

func (o *QosPolicyCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qosPolicyCreateCreated %s", 201, payload)
}

func (o *QosPolicyCreateCreated) GetPayload() *models.QosPolicyJobLinkResponse {
	return o.Payload
}

func (o *QosPolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.QosPolicyJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyCreateAccepted creates a QosPolicyCreateAccepted with default headers values
func NewQosPolicyCreateAccepted() *QosPolicyCreateAccepted {
	return &QosPolicyCreateAccepted{}
}

/*
QosPolicyCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QosPolicyCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.QosPolicyJobLinkResponse
}

// IsSuccess returns true when this qos policy create accepted response has a 2xx status code
func (o *QosPolicyCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy create accepted response has a 3xx status code
func (o *QosPolicyCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy create accepted response has a 4xx status code
func (o *QosPolicyCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy create accepted response has a 5xx status code
func (o *QosPolicyCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy create accepted response a status code equal to that given
func (o *QosPolicyCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the qos policy create accepted response
func (o *QosPolicyCreateAccepted) Code() int {
	return 202
}

func (o *QosPolicyCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qosPolicyCreateAccepted %s", 202, payload)
}

func (o *QosPolicyCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qosPolicyCreateAccepted %s", 202, payload)
}

func (o *QosPolicyCreateAccepted) GetPayload() *models.QosPolicyJobLinkResponse {
	return o.Payload
}

func (o *QosPolicyCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.QosPolicyJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyCreateDefault creates a QosPolicyCreateDefault with default headers values
func NewQosPolicyCreateDefault(code int) *QosPolicyCreateDefault {
	return &QosPolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	QosPolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8454147 | The maximum limit for QoS policies has been reached. |
| 8454154 | The name specified for creating conflicts with an existing QoS policy name. |
| 8454194 | The minimum throughput value for the policy group must be less than or equal to the maximum throughput value. |
| 8454260 | Invalid value for maximum and minimum fields. Valid values for max_throughput_iops and max_throughput_mbps combination is for the ratio of max_throughput_mbps and max_throughput_iops to be within 1 to 4096. |
| 8454273 | Invalid value for an adaptive field. Value should be non-zero. |
| 8454274 | Invalid value for an adaptive field. Value for expected_iops must be less than or equal to the value for peak_iops. |
| 8454277 | The name specified for creating an adaptive QoS policy conflicts with an existing fixed QoS policy name. |
| 8454278 | The name specified for creating a fixed QoS policy conflicts with an existing adaptive QoS policy name. |
| 8454379 | The name specified for creating a fixed QoS policy already exists. |
| 8454380 | The name specified for creating an adaptive QoS policy already exists. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QosPolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qos policy create default response has a 2xx status code
func (o *QosPolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qos policy create default response has a 3xx status code
func (o *QosPolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qos policy create default response has a 4xx status code
func (o *QosPolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qos policy create default response has a 5xx status code
func (o *QosPolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qos policy create default response a status code equal to that given
func (o *QosPolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qos policy create default response
func (o *QosPolicyCreateDefault) Code() int {
	return o._statusCode
}

func (o *QosPolicyCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qos_policy_create default %s", o._statusCode, payload)
}

func (o *QosPolicyCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qos_policy_create default %s", o._statusCode, payload)
}

func (o *QosPolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
