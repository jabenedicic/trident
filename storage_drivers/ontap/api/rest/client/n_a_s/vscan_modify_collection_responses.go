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

// VscanModifyCollectionReader is a Reader for the VscanModifyCollection structure.
type VscanModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanModifyCollectionOK creates a VscanModifyCollectionOK with default headers values
func NewVscanModifyCollectionOK() *VscanModifyCollectionOK {
	return &VscanModifyCollectionOK{}
}

/*
VscanModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type VscanModifyCollectionOK struct {
}

// IsSuccess returns true when this vscan modify collection o k response has a 2xx status code
func (o *VscanModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan modify collection o k response has a 3xx status code
func (o *VscanModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan modify collection o k response has a 4xx status code
func (o *VscanModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan modify collection o k response has a 5xx status code
func (o *VscanModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan modify collection o k response a status code equal to that given
func (o *VscanModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan modify collection o k response
func (o *VscanModifyCollectionOK) Code() int {
	return 200
}

func (o *VscanModifyCollectionOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/vscan][%d] vscanModifyCollectionOK", 200)
}

func (o *VscanModifyCollectionOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/vscan][%d] vscanModifyCollectionOK", 200)
}

func (o *VscanModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanModifyCollectionDefault creates a VscanModifyCollectionDefault with default headers values
func NewVscanModifyCollectionDefault(code int) *VscanModifyCollectionDefault {
	return &VscanModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	VscanModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027015   | Attempting to enable a Vscan but no active scanner-pool exists for the specified SVM
| 10027009   | Attempting to enable a Vscan for an SVM for which it's already enabled
| 10027010   | Attempting to disable a Vscan for an SVM for which it's already disabled
| 10027011   | Attempting to enable a Vscan for an SVM for which no CIFS server exists
| 10027023   | Attempting to enable a Vscan for an SVM for which no active Vscan On-Access policy exists
| 10027012   | Cannot enable Vscan on an administrative SVM.
*/
type VscanModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan modify collection default response has a 2xx status code
func (o *VscanModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan modify collection default response has a 3xx status code
func (o *VscanModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan modify collection default response has a 4xx status code
func (o *VscanModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan modify collection default response has a 5xx status code
func (o *VscanModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan modify collection default response a status code equal to that given
func (o *VscanModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan modify collection default response
func (o *VscanModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *VscanModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/vscan][%d] vscan_modify_collection default %s", o._statusCode, payload)
}

func (o *VscanModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/vscan][%d] vscan_modify_collection default %s", o._statusCode, payload)
}

func (o *VscanModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VscanModifyCollectionBody vscan modify collection body
swagger:model VscanModifyCollectionBody
*/
type VscanModifyCollectionBody struct {

	// links
	Links *models.VscanInlineLinks `json:"_links,omitempty"`

	// Discards the cached information of the files that have been successfully scanned. Once the cache is cleared, files are scanned again when they are accessed. PATCH only
	CacheClear *bool `json:"cache_clear,omitempty"`

	// Specifies whether or not Vscan is enabled on the SVM.
	Enabled *bool `json:"enabled,omitempty"`

	// svm
	Svm *models.VscanInlineSvm `json:"svm,omitempty"`

	// vscan inline on access policies
	VscanInlineOnAccessPolicies []*models.VscanOnAccessPolicy `json:"on_access_policies,omitempty"`

	// vscan inline on demand policies
	VscanInlineOnDemandPolicies []*models.VscanOnDemandPolicy `json:"on_demand_policies,omitempty"`

	// vscan inline scanner pools
	VscanInlineScannerPools []*models.ScannerPool `json:"scanner_pools,omitempty"`

	// vscan response inline records
	VscanResponseInlineRecords []*models.Vscan `json:"records,omitempty"`
}

// Validate validates this vscan modify collection body
func (o *VscanModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVscanInlineOnAccessPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVscanInlineOnDemandPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVscanInlineScannerPools(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVscanResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
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

func (o *VscanModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

func (o *VscanModifyCollectionBody) validateVscanInlineOnAccessPolicies(formats strfmt.Registry) error {
	if swag.IsZero(o.VscanInlineOnAccessPolicies) { // not required
		return nil
	}

	for i := 0; i < len(o.VscanInlineOnAccessPolicies); i++ {
		if swag.IsZero(o.VscanInlineOnAccessPolicies[i]) { // not required
			continue
		}

		if o.VscanInlineOnAccessPolicies[i] != nil {
			if err := o.VscanInlineOnAccessPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "on_access_policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *VscanModifyCollectionBody) validateVscanInlineOnDemandPolicies(formats strfmt.Registry) error {
	if swag.IsZero(o.VscanInlineOnDemandPolicies) { // not required
		return nil
	}

	for i := 0; i < len(o.VscanInlineOnDemandPolicies); i++ {
		if swag.IsZero(o.VscanInlineOnDemandPolicies[i]) { // not required
			continue
		}

		if o.VscanInlineOnDemandPolicies[i] != nil {
			if err := o.VscanInlineOnDemandPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "on_demand_policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *VscanModifyCollectionBody) validateVscanInlineScannerPools(formats strfmt.Registry) error {
	if swag.IsZero(o.VscanInlineScannerPools) { // not required
		return nil
	}

	for i := 0; i < len(o.VscanInlineScannerPools); i++ {
		if swag.IsZero(o.VscanInlineScannerPools[i]) { // not required
			continue
		}

		if o.VscanInlineScannerPools[i] != nil {
			if err := o.VscanInlineScannerPools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "scanner_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *VscanModifyCollectionBody) validateVscanResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.VscanResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.VscanResponseInlineRecords); i++ {
		if swag.IsZero(o.VscanResponseInlineRecords[i]) { // not required
			continue
		}

		if o.VscanResponseInlineRecords[i] != nil {
			if err := o.VscanResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vscan modify collection body based on the context it is used
func (o *VscanModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVscanInlineOnAccessPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVscanInlineOnDemandPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVscanInlineScannerPools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVscanResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (o *VscanModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (o *VscanModifyCollectionBody) contextValidateVscanInlineOnAccessPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VscanInlineOnAccessPolicies); i++ {

		if o.VscanInlineOnAccessPolicies[i] != nil {
			if err := o.VscanInlineOnAccessPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "on_access_policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *VscanModifyCollectionBody) contextValidateVscanInlineOnDemandPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VscanInlineOnDemandPolicies); i++ {

		if o.VscanInlineOnDemandPolicies[i] != nil {
			if err := o.VscanInlineOnDemandPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "on_demand_policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *VscanModifyCollectionBody) contextValidateVscanInlineScannerPools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VscanInlineScannerPools); i++ {

		if o.VscanInlineScannerPools[i] != nil {
			if err := o.VscanInlineScannerPools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "scanner_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *VscanModifyCollectionBody) contextValidateVscanResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VscanResponseInlineRecords); i++ {

		if o.VscanResponseInlineRecords[i] != nil {
			if err := o.VscanResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *VscanModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res VscanModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VscanInlineLinks vscan inline links
swagger:model vscan_inline__links
*/
type VscanInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this vscan inline links
func (o *VscanInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this vscan inline links based on the context it is used
func (o *VscanInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *VscanInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanInlineLinks) UnmarshalBinary(b []byte) error {
	var res VscanInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VscanInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model vscan_inline_svm
*/
type VscanInlineSvm struct {

	// links
	Links *models.VscanInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this vscan inline svm
func (o *VscanInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this vscan inline svm based on the context it is used
func (o *VscanInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *VscanInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanInlineSvm) UnmarshalBinary(b []byte) error {
	var res VscanInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VscanInlineSvmInlineLinks vscan inline svm inline links
swagger:model vscan_inline_svm_inline__links
*/
type VscanInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this vscan inline svm inline links
func (o *VscanInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this vscan inline svm inline links based on the context it is used
func (o *VscanInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *VscanInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res VscanInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
