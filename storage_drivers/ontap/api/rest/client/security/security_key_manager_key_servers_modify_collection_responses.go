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
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityKeyManagerKeyServersModifyCollectionReader is a Reader for the SecurityKeyManagerKeyServersModifyCollection structure.
type SecurityKeyManagerKeyServersModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerKeyServersModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerKeyServersModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerKeyServersModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerKeyServersModifyCollectionOK creates a SecurityKeyManagerKeyServersModifyCollectionOK with default headers values
func NewSecurityKeyManagerKeyServersModifyCollectionOK() *SecurityKeyManagerKeyServersModifyCollectionOK {
	return &SecurityKeyManagerKeyServersModifyCollectionOK{}
}

/*
SecurityKeyManagerKeyServersModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerKeyServersModifyCollectionOK struct {
}

// IsSuccess returns true when this security key manager key servers modify collection o k response has a 2xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager key servers modify collection o k response has a 3xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager key servers modify collection o k response has a 4xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager key servers modify collection o k response has a 5xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager key servers modify collection o k response a status code equal to that given
func (o *SecurityKeyManagerKeyServersModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security key manager key servers modify collection o k response
func (o *SecurityKeyManagerKeyServersModifyCollectionOK) Code() int {
	return 200
}

func (o *SecurityKeyManagerKeyServersModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers][%d] securityKeyManagerKeyServersModifyCollectionOK", 200)
}

func (o *SecurityKeyManagerKeyServersModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers][%d] securityKeyManagerKeyServersModifyCollectionOK", 200)
}

func (o *SecurityKeyManagerKeyServersModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityKeyManagerKeyServersModifyCollectionDefault creates a SecurityKeyManagerKeyServersModifyCollectionDefault with default headers values
func NewSecurityKeyManagerKeyServersModifyCollectionDefault(code int) *SecurityKeyManagerKeyServersModifyCollectionDefault {
	return &SecurityKeyManagerKeyServersModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerKeyServersModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536600 | Cannot modify a key server while a node is out quorum. |
| 65536824 | Multitenant key management is not supported in MetroCluster configurations. |
| 65536828 | External key management is not enabled for the SVM. |
| 65536843 | The key management server is not configured for the SVM. |
| 65536845 | Missing username. |
| 65536846 | Missing password. |
| 65536880 | One or more of the following values must be provided \"timeout\", \"username\", \"password\", \"secondary_key_servers\", \"create_remove_timeout\". |
| 65536921 | Unable to execute the command on the KMIP server. |
| 65537400 | Exceeded maximum number of secondary key servers. |
| 65538407 | A secondary key server is a duplicate of the associated primary key server. |
| 65538408 | The list of secondary key servers contains duplicates. |
| 65538413 | A secondary key server address is not formatted correctly. |
| 65538502 | A secondary key server is also a primary key server. |
| 65538503 | Support for adding secondary key servers requires an ECV of ONTAP 9.11.1 or later. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeyManagerKeyServersModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security key manager key servers modify collection default response has a 2xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager key servers modify collection default response has a 3xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager key servers modify collection default response has a 4xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager key servers modify collection default response has a 5xx status code
func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager key servers modify collection default response a status code equal to that given
func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security key manager key servers modify collection default response
func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers][%d] security_key_manager_key_servers_modify_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-managers/{uuid}/key-servers][%d] security_key_manager_key_servers_modify_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecurityKeyManagerKeyServersModifyCollectionBody security key manager key servers modify collection body
swagger:model SecurityKeyManagerKeyServersModifyCollectionBody
*/
type SecurityKeyManagerKeyServersModifyCollectionBody struct {

	// links
	Links *models.KeyServerInlineLinks `json:"_links,omitempty"`

	// connectivity
	Connectivity *models.KeyServerInlineConnectivity `json:"connectivity,omitempty"`

	// The key server timeout for create and remove operations.
	// -1 indicates that the server will wait indefinitely for the event to occur. 0 indicates that the server will not wait and will immediately timeout if it does not receive a response.
	//
	// Example: 60
	// Maximum: 60
	// Minimum: -1
	CreateRemoveTimeout *int64 `json:"create_remove_timeout,omitempty"`

	// An array of key servers specified to add multiple key servers to a key manager in a single API call. Valid in POST only and not valid if `server` is provided.
	//
	// Max Items: 4
	KeyServerInlineRecords []*models.KeyServerInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// A list of the secondary key servers associated with the primary key server.
	// Example: ["secondary1.com","10.1.2.3"]
	KeyServerInlineSecondaryKeyServers []*string `json:"secondary_key_servers,omitempty"`

	// Password credentials for connecting with the key server. This is not audited.
	// Example: password
	// Format: password
	Password *strfmt.Password `json:"password,omitempty"`

	// External key server for key management. If no port is provided, a default port of 5696 is used. Not valid in POST if `records` is provided.
	// Example: keyserver1.com:5698
	Server *string `json:"server,omitempty"`

	// I/O timeout in seconds for communicating with the key server.
	// -1 indicates that the server will wait indefinitely for the event to occur. 0 indicates that the server will not wait and will immediately timeout if it does not receive a response.
	//
	// Example: 60
	// Maximum: 60
	// Minimum: -1
	Timeout *int64 `json:"timeout,omitempty"`

	// KMIP username credentials for connecting with the key server.
	// Example: username
	Username *string `json:"username,omitempty"`
}

// Validate validates this security key manager key servers modify collection body
func (o *SecurityKeyManagerKeyServersModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConnectivity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreateRemoveTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKeyServerInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) validateConnectivity(formats strfmt.Registry) error {
	if swag.IsZero(o.Connectivity) { // not required
		return nil
	}

	if o.Connectivity != nil {
		if err := o.Connectivity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "connectivity")
			}
			return err
		}
	}

	return nil
}

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) validateCreateRemoveTimeout(formats strfmt.Registry) error {
	if swag.IsZero(o.CreateRemoveTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("info"+"."+"create_remove_timeout", "body", *o.CreateRemoveTimeout, -1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("info"+"."+"create_remove_timeout", "body", *o.CreateRemoveTimeout, 60, false); err != nil {
		return err
	}

	return nil
}

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) validateKeyServerInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.KeyServerInlineRecords) { // not required
		return nil
	}

	iKeyServerInlineRecordsSize := int64(len(o.KeyServerInlineRecords))

	if err := validate.MaxItems("info"+"."+"records", "body", iKeyServerInlineRecordsSize, 4); err != nil {
		return err
	}

	for i := 0; i < len(o.KeyServerInlineRecords); i++ {
		if swag.IsZero(o.KeyServerInlineRecords[i]) { // not required
			continue
		}

		if o.KeyServerInlineRecords[i] != nil {
			if err := o.KeyServerInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(o.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("info"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(o.Timeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("info"+"."+"timeout", "body", *o.Timeout, -1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("info"+"."+"timeout", "body", *o.Timeout, 60, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security key manager key servers modify collection body based on the context it is used
func (o *SecurityKeyManagerKeyServersModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateConnectivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateKeyServerInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) contextValidateConnectivity(ctx context.Context, formats strfmt.Registry) error {

	if o.Connectivity != nil {
		if err := o.Connectivity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "connectivity")
			}
			return err
		}
	}

	return nil
}

func (o *SecurityKeyManagerKeyServersModifyCollectionBody) contextValidateKeyServerInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.KeyServerInlineRecords); i++ {

		if o.KeyServerInlineRecords[i] != nil {
			if err := o.KeyServerInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *SecurityKeyManagerKeyServersModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeyManagerKeyServersModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res SecurityKeyManagerKeyServersModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KeyServerInlineLinks key server inline links
swagger:model key_server_inline__links
*/
type KeyServerInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this key server inline links
func (o *KeyServerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this key server inline links based on the context it is used
func (o *KeyServerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *KeyServerInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KeyServerInlineLinks) UnmarshalBinary(b []byte) error {
	var res KeyServerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KeyServerInlineConnectivity This property contains the key server connectivity state of all nodes in the cluster.
// This is an advanced property; there is an added computational cost to retrieving its value. The property is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter or GET for all advanced properties is enabled.
//
swagger:model key_server_inline_connectivity
*/
type KeyServerInlineConnectivity struct {

	// Set to true when key server connectivity state is available on all nodes of the cluster.
	// Read Only: true
	ClusterAvailability *bool `json:"cluster_availability,omitempty"`

	// An array of key server connectivity states for each node.
	//
	// Read Only: true
	NodeStates []*models.KeyServerState `json:"node_states,omitempty"`
}

// Validate validates this key server inline connectivity
func (o *KeyServerInlineConnectivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineConnectivity) validateNodeStates(formats strfmt.Registry) error {
	if swag.IsZero(o.NodeStates) { // not required
		return nil
	}

	for i := 0; i < len(o.NodeStates); i++ {
		if swag.IsZero(o.NodeStates[i]) { // not required
			continue
		}

		if o.NodeStates[i] != nil {
			if err := o.NodeStates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "connectivity" + "." + "node_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this key server inline connectivity based on the context it is used
func (o *KeyServerInlineConnectivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateClusterAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNodeStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineConnectivity) contextValidateClusterAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"connectivity"+"."+"cluster_availability", "body", o.ClusterAvailability); err != nil {
		return err
	}

	return nil
}

func (o *KeyServerInlineConnectivity) contextValidateNodeStates(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"connectivity"+"."+"node_states", "body", []*models.KeyServerState(o.NodeStates)); err != nil {
		return err
	}

	for i := 0; i < len(o.NodeStates); i++ {

		if o.NodeStates[i] != nil {
			if err := o.NodeStates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "connectivity" + "." + "node_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *KeyServerInlineConnectivity) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KeyServerInlineConnectivity) UnmarshalBinary(b []byte) error {
	var res KeyServerInlineConnectivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KeyServerInlineRecordsInlineArrayItem key server inline records inline array item
swagger:model key_server_inline_records_inline_array_item
*/
type KeyServerInlineRecordsInlineArrayItem struct {

	// links
	Links *models.KeyServerInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// connectivity
	Connectivity *models.KeyServerInlineRecordsInlineArrayItemInlineConnectivity `json:"connectivity,omitempty"`

	// Password credentials for connecting with the key server. This is not audited.
	// Example: password
	// Format: password
	Password *strfmt.Password `json:"password,omitempty"`

	// External key server for key management. If no port is provided, a default port of 5696 is used. Not valid in POST if `records` is provided.
	// Example: bulkkeyserver.com:5698
	Server *string `json:"server,omitempty"`

	// I/O timeout in seconds for communicating with the key server.
	// Example: 60
	// Maximum: 60
	// Minimum: 1
	Timeout *int64 `json:"timeout,omitempty"`

	// KMIP username credentials for connecting with the key server.
	// Example: username
	Username *string `json:"username,omitempty"`
}

// Validate validates this key server inline records inline array item
func (o *KeyServerInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConnectivity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItem) validateConnectivity(formats strfmt.Registry) error {
	if swag.IsZero(o.Connectivity) { // not required
		return nil
	}

	if o.Connectivity != nil {
		if err := o.Connectivity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity")
			}
			return err
		}
	}

	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItem) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(o.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItem) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(o.Timeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout", "body", *o.Timeout, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timeout", "body", *o.Timeout, 60, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this key server inline records inline array item based on the context it is used
func (o *KeyServerInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateConnectivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItem) contextValidateConnectivity(ctx context.Context, formats strfmt.Registry) error {

	if o.Connectivity != nil {
		if err := o.Connectivity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KeyServerInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KeyServerInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res KeyServerInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KeyServerInlineRecordsInlineArrayItemInlineLinks key server inline records inline array item inline links
swagger:model key_server_inline_records_inline_array_item_inline__links
*/
type KeyServerInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this key server inline records inline array item inline links
func (o *KeyServerInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this key server inline records inline array item inline links based on the context it is used
func (o *KeyServerInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *KeyServerInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KeyServerInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res KeyServerInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
KeyServerInlineRecordsInlineArrayItemInlineConnectivity This property contains the key server connectivity state of all nodes in the cluster.
// This is an advanced property; there is an added computational cost to retrieving its value. The property is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter or GET for all advanced properties is enabled.
//
swagger:model key_server_inline_records_inline_array_item_inline_connectivity
*/
type KeyServerInlineRecordsInlineArrayItemInlineConnectivity struct {

	// Set to true when key server connectivity state is available on all nodes of the cluster.
	// Read Only: true
	ClusterAvailability *bool `json:"cluster_availability,omitempty"`

	// An array of key server connectivity states for each node.
	//
	// Read Only: true
	NodeStates []*models.KeyServerState `json:"node_states"`
}

// Validate validates this key server inline records inline array item inline connectivity
func (o *KeyServerInlineRecordsInlineArrayItemInlineConnectivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItemInlineConnectivity) validateNodeStates(formats strfmt.Registry) error {
	if swag.IsZero(o.NodeStates) { // not required
		return nil
	}

	for i := 0; i < len(o.NodeStates); i++ {
		if swag.IsZero(o.NodeStates[i]) { // not required
			continue
		}

		if o.NodeStates[i] != nil {
			if err := o.NodeStates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectivity" + "." + "node_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this key server inline records inline array item inline connectivity based on the context it is used
func (o *KeyServerInlineRecordsInlineArrayItemInlineConnectivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateClusterAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNodeStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItemInlineConnectivity) contextValidateClusterAvailability(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connectivity"+"."+"cluster_availability", "body", o.ClusterAvailability); err != nil {
		return err
	}

	return nil
}

func (o *KeyServerInlineRecordsInlineArrayItemInlineConnectivity) contextValidateNodeStates(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connectivity"+"."+"node_states", "body", []*models.KeyServerState(o.NodeStates)); err != nil {
		return err
	}

	for i := 0; i < len(o.NodeStates); i++ {

		if o.NodeStates[i] != nil {
			if err := o.NodeStates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectivity" + "." + "node_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *KeyServerInlineRecordsInlineArrayItemInlineConnectivity) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KeyServerInlineRecordsInlineArrayItemInlineConnectivity) UnmarshalBinary(b []byte) error {
	var res KeyServerInlineRecordsInlineArrayItemInlineConnectivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
