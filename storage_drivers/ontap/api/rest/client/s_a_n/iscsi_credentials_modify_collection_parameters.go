// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewIscsiCredentialsModifyCollectionParams creates a new IscsiCredentialsModifyCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiCredentialsModifyCollectionParams() *IscsiCredentialsModifyCollectionParams {
	return &IscsiCredentialsModifyCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiCredentialsModifyCollectionParamsWithTimeout creates a new IscsiCredentialsModifyCollectionParams object
// with the ability to set a timeout on a request.
func NewIscsiCredentialsModifyCollectionParamsWithTimeout(timeout time.Duration) *IscsiCredentialsModifyCollectionParams {
	return &IscsiCredentialsModifyCollectionParams{
		timeout: timeout,
	}
}

// NewIscsiCredentialsModifyCollectionParamsWithContext creates a new IscsiCredentialsModifyCollectionParams object
// with the ability to set a context for a request.
func NewIscsiCredentialsModifyCollectionParamsWithContext(ctx context.Context) *IscsiCredentialsModifyCollectionParams {
	return &IscsiCredentialsModifyCollectionParams{
		Context: ctx,
	}
}

// NewIscsiCredentialsModifyCollectionParamsWithHTTPClient creates a new IscsiCredentialsModifyCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiCredentialsModifyCollectionParamsWithHTTPClient(client *http.Client) *IscsiCredentialsModifyCollectionParams {
	return &IscsiCredentialsModifyCollectionParams{
		HTTPClient: client,
	}
}

/*
IscsiCredentialsModifyCollectionParams contains all the parameters to send to the API endpoint

	for the iscsi credentials modify collection operation.

	Typically these are written to a http.Request.
*/
type IscsiCredentialsModifyCollectionParams struct {

	/* AddInitiatorAddresses.

	   If _true_, the initiator addresses in the body merge into the existing addresses in the iSCSI security object rather than replace the existing addresses.

	*/
	AddInitiatorAddresses *bool

	/* AuthenticationType.

	   Filter by authentication_type
	*/
	AuthenticationType *string

	/* ChapInboundUser.

	   Filter by chap.inbound.user
	*/
	ChapInboundUser *string

	/* ChapOutboundUser.

	   Filter by chap.outbound.user
	*/
	ChapOutboundUser *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info IscsiCredentialsModifyCollectionBody

	/* Initiator.

	   Filter by initiator
	*/
	Initiator *string

	/* InitiatorAddressMasksAddress.

	   Filter by initiator_address.masks.address
	*/
	InitiatorAddressMasksAddress *string

	/* InitiatorAddressMasksFamily.

	   Filter by initiator_address.masks.family
	*/
	InitiatorAddressMasksFamily *string

	/* InitiatorAddressMasksNetmask.

	   Filter by initiator_address.masks.netmask
	*/
	InitiatorAddressMasksNetmask *string

	/* InitiatorAddressRangesEnd.

	   Filter by initiator_address.ranges.end
	*/
	InitiatorAddressRangesEnd *string

	/* InitiatorAddressRangesFamily.

	   Filter by initiator_address.ranges.family
	*/
	InitiatorAddressRangesFamily *string

	/* InitiatorAddressRangesStart.

	   Filter by initiator_address.ranges.start
	*/
	InitiatorAddressRangesStart *string

	/* RemoveInitiatorAddresses.

	   If _true_, the initiator addresses in the body are removed from the existing addresses in the iSCSI security object rather than replace the existing addresses.

	*/
	RemoveInitiatorAddresses *bool

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi credentials modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsModifyCollectionParams) WithDefaults() *IscsiCredentialsModifyCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi credentials modify collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsModifyCollectionParams) SetDefaults() {
	var (
		addInitiatorAddressesDefault = bool(false)

		continueOnFailureDefault = bool(false)

		removeInitiatorAddressesDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := IscsiCredentialsModifyCollectionParams{
		AddInitiatorAddresses:    &addInitiatorAddressesDefault,
		ContinueOnFailure:        &continueOnFailureDefault,
		RemoveInitiatorAddresses: &removeInitiatorAddressesDefault,
		ReturnRecords:            &returnRecordsDefault,
		ReturnTimeout:            &returnTimeoutDefault,
		SerialRecords:            &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithTimeout(timeout time.Duration) *IscsiCredentialsModifyCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithContext(ctx context.Context) *IscsiCredentialsModifyCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithHTTPClient(client *http.Client) *IscsiCredentialsModifyCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddInitiatorAddresses adds the addInitiatorAddresses to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithAddInitiatorAddresses(addInitiatorAddresses *bool) *IscsiCredentialsModifyCollectionParams {
	o.SetAddInitiatorAddresses(addInitiatorAddresses)
	return o
}

// SetAddInitiatorAddresses adds the addInitiatorAddresses to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetAddInitiatorAddresses(addInitiatorAddresses *bool) {
	o.AddInitiatorAddresses = addInitiatorAddresses
}

// WithAuthenticationType adds the authenticationType to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithAuthenticationType(authenticationType *string) *IscsiCredentialsModifyCollectionParams {
	o.SetAuthenticationType(authenticationType)
	return o
}

// SetAuthenticationType adds the authenticationType to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetAuthenticationType(authenticationType *string) {
	o.AuthenticationType = authenticationType
}

// WithChapInboundUser adds the chapInboundUser to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithChapInboundUser(chapInboundUser *string) *IscsiCredentialsModifyCollectionParams {
	o.SetChapInboundUser(chapInboundUser)
	return o
}

// SetChapInboundUser adds the chapInboundUser to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetChapInboundUser(chapInboundUser *string) {
	o.ChapInboundUser = chapInboundUser
}

// WithChapOutboundUser adds the chapOutboundUser to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithChapOutboundUser(chapOutboundUser *string) *IscsiCredentialsModifyCollectionParams {
	o.SetChapOutboundUser(chapOutboundUser)
	return o
}

// SetChapOutboundUser adds the chapOutboundUser to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetChapOutboundUser(chapOutboundUser *string) {
	o.ChapOutboundUser = chapOutboundUser
}

// WithContinueOnFailure adds the continueOnFailure to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *IscsiCredentialsModifyCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInfo(info IscsiCredentialsModifyCollectionBody) *IscsiCredentialsModifyCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInfo(info IscsiCredentialsModifyCollectionBody) {
	o.Info = info
}

// WithInitiator adds the initiator to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInitiator(initiator *string) *IscsiCredentialsModifyCollectionParams {
	o.SetInitiator(initiator)
	return o
}

// SetInitiator adds the initiator to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInitiator(initiator *string) {
	o.Initiator = initiator
}

// WithInitiatorAddressMasksAddress adds the initiatorAddressMasksAddress to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInitiatorAddressMasksAddress(initiatorAddressMasksAddress *string) *IscsiCredentialsModifyCollectionParams {
	o.SetInitiatorAddressMasksAddress(initiatorAddressMasksAddress)
	return o
}

// SetInitiatorAddressMasksAddress adds the initiatorAddressMasksAddress to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInitiatorAddressMasksAddress(initiatorAddressMasksAddress *string) {
	o.InitiatorAddressMasksAddress = initiatorAddressMasksAddress
}

// WithInitiatorAddressMasksFamily adds the initiatorAddressMasksFamily to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInitiatorAddressMasksFamily(initiatorAddressMasksFamily *string) *IscsiCredentialsModifyCollectionParams {
	o.SetInitiatorAddressMasksFamily(initiatorAddressMasksFamily)
	return o
}

// SetInitiatorAddressMasksFamily adds the initiatorAddressMasksFamily to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInitiatorAddressMasksFamily(initiatorAddressMasksFamily *string) {
	o.InitiatorAddressMasksFamily = initiatorAddressMasksFamily
}

// WithInitiatorAddressMasksNetmask adds the initiatorAddressMasksNetmask to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInitiatorAddressMasksNetmask(initiatorAddressMasksNetmask *string) *IscsiCredentialsModifyCollectionParams {
	o.SetInitiatorAddressMasksNetmask(initiatorAddressMasksNetmask)
	return o
}

// SetInitiatorAddressMasksNetmask adds the initiatorAddressMasksNetmask to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInitiatorAddressMasksNetmask(initiatorAddressMasksNetmask *string) {
	o.InitiatorAddressMasksNetmask = initiatorAddressMasksNetmask
}

// WithInitiatorAddressRangesEnd adds the initiatorAddressRangesEnd to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInitiatorAddressRangesEnd(initiatorAddressRangesEnd *string) *IscsiCredentialsModifyCollectionParams {
	o.SetInitiatorAddressRangesEnd(initiatorAddressRangesEnd)
	return o
}

// SetInitiatorAddressRangesEnd adds the initiatorAddressRangesEnd to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInitiatorAddressRangesEnd(initiatorAddressRangesEnd *string) {
	o.InitiatorAddressRangesEnd = initiatorAddressRangesEnd
}

// WithInitiatorAddressRangesFamily adds the initiatorAddressRangesFamily to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInitiatorAddressRangesFamily(initiatorAddressRangesFamily *string) *IscsiCredentialsModifyCollectionParams {
	o.SetInitiatorAddressRangesFamily(initiatorAddressRangesFamily)
	return o
}

// SetInitiatorAddressRangesFamily adds the initiatorAddressRangesFamily to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInitiatorAddressRangesFamily(initiatorAddressRangesFamily *string) {
	o.InitiatorAddressRangesFamily = initiatorAddressRangesFamily
}

// WithInitiatorAddressRangesStart adds the initiatorAddressRangesStart to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithInitiatorAddressRangesStart(initiatorAddressRangesStart *string) *IscsiCredentialsModifyCollectionParams {
	o.SetInitiatorAddressRangesStart(initiatorAddressRangesStart)
	return o
}

// SetInitiatorAddressRangesStart adds the initiatorAddressRangesStart to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetInitiatorAddressRangesStart(initiatorAddressRangesStart *string) {
	o.InitiatorAddressRangesStart = initiatorAddressRangesStart
}

// WithRemoveInitiatorAddresses adds the removeInitiatorAddresses to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithRemoveInitiatorAddresses(removeInitiatorAddresses *bool) *IscsiCredentialsModifyCollectionParams {
	o.SetRemoveInitiatorAddresses(removeInitiatorAddresses)
	return o
}

// SetRemoveInitiatorAddresses adds the removeInitiatorAddresses to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetRemoveInitiatorAddresses(removeInitiatorAddresses *bool) {
	o.RemoveInitiatorAddresses = removeInitiatorAddresses
}

// WithReturnRecords adds the returnRecords to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithReturnRecords(returnRecords *bool) *IscsiCredentialsModifyCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithReturnTimeout(returnTimeout *int64) *IscsiCredentialsModifyCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithSerialRecords(serialRecords *bool) *IscsiCredentialsModifyCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithSvmName(svmName *string) *IscsiCredentialsModifyCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) WithSvmUUID(svmUUID *string) *IscsiCredentialsModifyCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the iscsi credentials modify collection params
func (o *IscsiCredentialsModifyCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiCredentialsModifyCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddInitiatorAddresses != nil {

		// query param add_initiator_addresses
		var qrAddInitiatorAddresses bool

		if o.AddInitiatorAddresses != nil {
			qrAddInitiatorAddresses = *o.AddInitiatorAddresses
		}
		qAddInitiatorAddresses := swag.FormatBool(qrAddInitiatorAddresses)
		if qAddInitiatorAddresses != "" {

			if err := r.SetQueryParam("add_initiator_addresses", qAddInitiatorAddresses); err != nil {
				return err
			}
		}
	}

	if o.AuthenticationType != nil {

		// query param authentication_type
		var qrAuthenticationType string

		if o.AuthenticationType != nil {
			qrAuthenticationType = *o.AuthenticationType
		}
		qAuthenticationType := qrAuthenticationType
		if qAuthenticationType != "" {

			if err := r.SetQueryParam("authentication_type", qAuthenticationType); err != nil {
				return err
			}
		}
	}

	if o.ChapInboundUser != nil {

		// query param chap.inbound.user
		var qrChapInboundUser string

		if o.ChapInboundUser != nil {
			qrChapInboundUser = *o.ChapInboundUser
		}
		qChapInboundUser := qrChapInboundUser
		if qChapInboundUser != "" {

			if err := r.SetQueryParam("chap.inbound.user", qChapInboundUser); err != nil {
				return err
			}
		}
	}

	if o.ChapOutboundUser != nil {

		// query param chap.outbound.user
		var qrChapOutboundUser string

		if o.ChapOutboundUser != nil {
			qrChapOutboundUser = *o.ChapOutboundUser
		}
		qChapOutboundUser := qrChapOutboundUser
		if qChapOutboundUser != "" {

			if err := r.SetQueryParam("chap.outbound.user", qChapOutboundUser); err != nil {
				return err
			}
		}
	}

	if o.ContinueOnFailure != nil {

		// query param continue_on_failure
		var qrContinueOnFailure bool

		if o.ContinueOnFailure != nil {
			qrContinueOnFailure = *o.ContinueOnFailure
		}
		qContinueOnFailure := swag.FormatBool(qrContinueOnFailure)
		if qContinueOnFailure != "" {

			if err := r.SetQueryParam("continue_on_failure", qContinueOnFailure); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.Initiator != nil {

		// query param initiator
		var qrInitiator string

		if o.Initiator != nil {
			qrInitiator = *o.Initiator
		}
		qInitiator := qrInitiator
		if qInitiator != "" {

			if err := r.SetQueryParam("initiator", qInitiator); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressMasksAddress != nil {

		// query param initiator_address.masks.address
		var qrInitiatorAddressMasksAddress string

		if o.InitiatorAddressMasksAddress != nil {
			qrInitiatorAddressMasksAddress = *o.InitiatorAddressMasksAddress
		}
		qInitiatorAddressMasksAddress := qrInitiatorAddressMasksAddress
		if qInitiatorAddressMasksAddress != "" {

			if err := r.SetQueryParam("initiator_address.masks.address", qInitiatorAddressMasksAddress); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressMasksFamily != nil {

		// query param initiator_address.masks.family
		var qrInitiatorAddressMasksFamily string

		if o.InitiatorAddressMasksFamily != nil {
			qrInitiatorAddressMasksFamily = *o.InitiatorAddressMasksFamily
		}
		qInitiatorAddressMasksFamily := qrInitiatorAddressMasksFamily
		if qInitiatorAddressMasksFamily != "" {

			if err := r.SetQueryParam("initiator_address.masks.family", qInitiatorAddressMasksFamily); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressMasksNetmask != nil {

		// query param initiator_address.masks.netmask
		var qrInitiatorAddressMasksNetmask string

		if o.InitiatorAddressMasksNetmask != nil {
			qrInitiatorAddressMasksNetmask = *o.InitiatorAddressMasksNetmask
		}
		qInitiatorAddressMasksNetmask := qrInitiatorAddressMasksNetmask
		if qInitiatorAddressMasksNetmask != "" {

			if err := r.SetQueryParam("initiator_address.masks.netmask", qInitiatorAddressMasksNetmask); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressRangesEnd != nil {

		// query param initiator_address.ranges.end
		var qrInitiatorAddressRangesEnd string

		if o.InitiatorAddressRangesEnd != nil {
			qrInitiatorAddressRangesEnd = *o.InitiatorAddressRangesEnd
		}
		qInitiatorAddressRangesEnd := qrInitiatorAddressRangesEnd
		if qInitiatorAddressRangesEnd != "" {

			if err := r.SetQueryParam("initiator_address.ranges.end", qInitiatorAddressRangesEnd); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressRangesFamily != nil {

		// query param initiator_address.ranges.family
		var qrInitiatorAddressRangesFamily string

		if o.InitiatorAddressRangesFamily != nil {
			qrInitiatorAddressRangesFamily = *o.InitiatorAddressRangesFamily
		}
		qInitiatorAddressRangesFamily := qrInitiatorAddressRangesFamily
		if qInitiatorAddressRangesFamily != "" {

			if err := r.SetQueryParam("initiator_address.ranges.family", qInitiatorAddressRangesFamily); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressRangesStart != nil {

		// query param initiator_address.ranges.start
		var qrInitiatorAddressRangesStart string

		if o.InitiatorAddressRangesStart != nil {
			qrInitiatorAddressRangesStart = *o.InitiatorAddressRangesStart
		}
		qInitiatorAddressRangesStart := qrInitiatorAddressRangesStart
		if qInitiatorAddressRangesStart != "" {

			if err := r.SetQueryParam("initiator_address.ranges.start", qInitiatorAddressRangesStart); err != nil {
				return err
			}
		}
	}

	if o.RemoveInitiatorAddresses != nil {

		// query param remove_initiator_addresses
		var qrRemoveInitiatorAddresses bool

		if o.RemoveInitiatorAddresses != nil {
			qrRemoveInitiatorAddresses = *o.RemoveInitiatorAddresses
		}
		qRemoveInitiatorAddresses := swag.FormatBool(qrRemoveInitiatorAddresses)
		if qRemoveInitiatorAddresses != "" {

			if err := r.SetQueryParam("remove_initiator_addresses", qRemoveInitiatorAddresses); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SerialRecords != nil {

		// query param serial_records
		var qrSerialRecords bool

		if o.SerialRecords != nil {
			qrSerialRecords = *o.SerialRecords
		}
		qSerialRecords := swag.FormatBool(qrSerialRecords)
		if qSerialRecords != "" {

			if err := r.SetQueryParam("serial_records", qSerialRecords); err != nil {
				return err
			}
		}
	}

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
