// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewFlexcacheDeleteCollectionParams creates a new FlexcacheDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFlexcacheDeleteCollectionParams() *FlexcacheDeleteCollectionParams {
	return &FlexcacheDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFlexcacheDeleteCollectionParamsWithTimeout creates a new FlexcacheDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewFlexcacheDeleteCollectionParamsWithTimeout(timeout time.Duration) *FlexcacheDeleteCollectionParams {
	return &FlexcacheDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewFlexcacheDeleteCollectionParamsWithContext creates a new FlexcacheDeleteCollectionParams object
// with the ability to set a context for a request.
func NewFlexcacheDeleteCollectionParamsWithContext(ctx context.Context) *FlexcacheDeleteCollectionParams {
	return &FlexcacheDeleteCollectionParams{
		Context: ctx,
	}
}

// NewFlexcacheDeleteCollectionParamsWithHTTPClient creates a new FlexcacheDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewFlexcacheDeleteCollectionParamsWithHTTPClient(client *http.Client) *FlexcacheDeleteCollectionParams {
	return &FlexcacheDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
FlexcacheDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the flexcache delete collection operation.

	Typically these are written to a http.Request.
*/
type FlexcacheDeleteCollectionParams struct {

	/* AggregatesName.

	   Filter by aggregates.name
	*/
	AggregatesName *string

	/* AggregatesUUID.

	   Filter by aggregates.uuid
	*/
	AggregatesUUID *string

	/* AtimeScrubEnabled.

	   Filter by atime_scrub.enabled
	*/
	AtimeScrubEnabled *bool

	/* AtimeScrubPeriod.

	   Filter by atime_scrub.period
	*/
	AtimeScrubPeriod *int64

	/* CifsChangeNotifyEnabled.

	   Filter by cifs_change_notify.enabled
	*/
	CifsChangeNotifyEnabled *bool

	/* ConstituentsPerAggregate.

	   Filter by constituents_per_aggregate
	*/
	ConstituentsPerAggregate *int64

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* DrCache.

	   Filter by dr_cache
	*/
	DrCache *bool

	/* GlobalFileLockingEnabled.

	   Filter by global_file_locking_enabled
	*/
	GlobalFileLockingEnabled *bool

	/* GuaranteeType.

	   Filter by guarantee.type
	*/
	GuaranteeType *string

	/* Info.

	   Info specification
	*/
	Info FlexcacheDeleteCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OriginsClusterName.

	   Filter by origins.cluster.name
	*/
	OriginsClusterName *string

	/* OriginsClusterUUID.

	   Filter by origins.cluster.uuid
	*/
	OriginsClusterUUID *string

	/* OriginsCreateTime.

	   Filter by origins.create_time
	*/
	OriginsCreateTime *string

	/* OriginsIPAddress.

	   Filter by origins.ip_address
	*/
	OriginsIPAddress *string

	/* OriginsSize.

	   Filter by origins.size
	*/
	OriginsSize *int64

	/* OriginsState.

	   Filter by origins.state
	*/
	OriginsState *string

	/* OriginsSvmName.

	   Filter by origins.svm.name
	*/
	OriginsSvmName *string

	/* OriginsSvmUUID.

	   Filter by origins.svm.uuid
	*/
	OriginsSvmUUID *string

	/* OriginsVolumeName.

	   Filter by origins.volume.name
	*/
	OriginsVolumeName *string

	/* OriginsVolumeUUID.

	   Filter by origins.volume.uuid
	*/
	OriginsVolumeUUID *string

	/* OverrideEncryption.

	   Filter by override_encryption
	*/
	OverrideEncryption *bool

	/* Path.

	   Filter by path
	*/
	Path *string

	/* RelativeSizeEnabled.

	   Filter by relative_size.enabled
	*/
	RelativeSizeEnabled *bool

	/* RelativeSizePercentage.

	   Filter by relative_size.percentage
	*/
	RelativeSizePercentage *int64

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

	/* Size.

	   Filter by size
	*/
	Size *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* UseTieredAggregate.

	   Filter by use_tiered_aggregate
	*/
	UseTieredAggregate *bool

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	/* WritebackEnabled.

	   Filter by writeback.enabled
	*/
	WritebackEnabled *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the flexcache delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlexcacheDeleteCollectionParams) WithDefaults() *FlexcacheDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the flexcache delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlexcacheDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := FlexcacheDeleteCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithTimeout(timeout time.Duration) *FlexcacheDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithContext(ctx context.Context) *FlexcacheDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithHTTPClient(client *http.Client) *FlexcacheDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAggregatesName adds the aggregatesName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithAggregatesName(aggregatesName *string) *FlexcacheDeleteCollectionParams {
	o.SetAggregatesName(aggregatesName)
	return o
}

// SetAggregatesName adds the aggregatesName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetAggregatesName(aggregatesName *string) {
	o.AggregatesName = aggregatesName
}

// WithAggregatesUUID adds the aggregatesUUID to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithAggregatesUUID(aggregatesUUID *string) *FlexcacheDeleteCollectionParams {
	o.SetAggregatesUUID(aggregatesUUID)
	return o
}

// SetAggregatesUUID adds the aggregatesUuid to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetAggregatesUUID(aggregatesUUID *string) {
	o.AggregatesUUID = aggregatesUUID
}

// WithAtimeScrubEnabled adds the atimeScrubEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithAtimeScrubEnabled(atimeScrubEnabled *bool) *FlexcacheDeleteCollectionParams {
	o.SetAtimeScrubEnabled(atimeScrubEnabled)
	return o
}

// SetAtimeScrubEnabled adds the atimeScrubEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetAtimeScrubEnabled(atimeScrubEnabled *bool) {
	o.AtimeScrubEnabled = atimeScrubEnabled
}

// WithAtimeScrubPeriod adds the atimeScrubPeriod to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithAtimeScrubPeriod(atimeScrubPeriod *int64) *FlexcacheDeleteCollectionParams {
	o.SetAtimeScrubPeriod(atimeScrubPeriod)
	return o
}

// SetAtimeScrubPeriod adds the atimeScrubPeriod to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetAtimeScrubPeriod(atimeScrubPeriod *int64) {
	o.AtimeScrubPeriod = atimeScrubPeriod
}

// WithCifsChangeNotifyEnabled adds the cifsChangeNotifyEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithCifsChangeNotifyEnabled(cifsChangeNotifyEnabled *bool) *FlexcacheDeleteCollectionParams {
	o.SetCifsChangeNotifyEnabled(cifsChangeNotifyEnabled)
	return o
}

// SetCifsChangeNotifyEnabled adds the cifsChangeNotifyEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetCifsChangeNotifyEnabled(cifsChangeNotifyEnabled *bool) {
	o.CifsChangeNotifyEnabled = cifsChangeNotifyEnabled
}

// WithConstituentsPerAggregate adds the constituentsPerAggregate to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithConstituentsPerAggregate(constituentsPerAggregate *int64) *FlexcacheDeleteCollectionParams {
	o.SetConstituentsPerAggregate(constituentsPerAggregate)
	return o
}

// SetConstituentsPerAggregate adds the constituentsPerAggregate to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetConstituentsPerAggregate(constituentsPerAggregate *int64) {
	o.ConstituentsPerAggregate = constituentsPerAggregate
}

// WithContinueOnFailure adds the continueOnFailure to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *FlexcacheDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithDrCache adds the drCache to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithDrCache(drCache *bool) *FlexcacheDeleteCollectionParams {
	o.SetDrCache(drCache)
	return o
}

// SetDrCache adds the drCache to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetDrCache(drCache *bool) {
	o.DrCache = drCache
}

// WithGlobalFileLockingEnabled adds the globalFileLockingEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithGlobalFileLockingEnabled(globalFileLockingEnabled *bool) *FlexcacheDeleteCollectionParams {
	o.SetGlobalFileLockingEnabled(globalFileLockingEnabled)
	return o
}

// SetGlobalFileLockingEnabled adds the globalFileLockingEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetGlobalFileLockingEnabled(globalFileLockingEnabled *bool) {
	o.GlobalFileLockingEnabled = globalFileLockingEnabled
}

// WithGuaranteeType adds the guaranteeType to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithGuaranteeType(guaranteeType *string) *FlexcacheDeleteCollectionParams {
	o.SetGuaranteeType(guaranteeType)
	return o
}

// SetGuaranteeType adds the guaranteeType to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetGuaranteeType(guaranteeType *string) {
	o.GuaranteeType = guaranteeType
}

// WithInfo adds the info to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithInfo(info FlexcacheDeleteCollectionBody) *FlexcacheDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetInfo(info FlexcacheDeleteCollectionBody) {
	o.Info = info
}

// WithName adds the name to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithName(name *string) *FlexcacheDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithOriginsClusterName adds the originsClusterName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsClusterName(originsClusterName *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsClusterName(originsClusterName)
	return o
}

// SetOriginsClusterName adds the originsClusterName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsClusterName(originsClusterName *string) {
	o.OriginsClusterName = originsClusterName
}

// WithOriginsClusterUUID adds the originsClusterUUID to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsClusterUUID(originsClusterUUID *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsClusterUUID(originsClusterUUID)
	return o
}

// SetOriginsClusterUUID adds the originsClusterUuid to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsClusterUUID(originsClusterUUID *string) {
	o.OriginsClusterUUID = originsClusterUUID
}

// WithOriginsCreateTime adds the originsCreateTime to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsCreateTime(originsCreateTime *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsCreateTime(originsCreateTime)
	return o
}

// SetOriginsCreateTime adds the originsCreateTime to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsCreateTime(originsCreateTime *string) {
	o.OriginsCreateTime = originsCreateTime
}

// WithOriginsIPAddress adds the originsIPAddress to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsIPAddress(originsIPAddress *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsIPAddress(originsIPAddress)
	return o
}

// SetOriginsIPAddress adds the originsIpAddress to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsIPAddress(originsIPAddress *string) {
	o.OriginsIPAddress = originsIPAddress
}

// WithOriginsSize adds the originsSize to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsSize(originsSize *int64) *FlexcacheDeleteCollectionParams {
	o.SetOriginsSize(originsSize)
	return o
}

// SetOriginsSize adds the originsSize to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsSize(originsSize *int64) {
	o.OriginsSize = originsSize
}

// WithOriginsState adds the originsState to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsState(originsState *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsState(originsState)
	return o
}

// SetOriginsState adds the originsState to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsState(originsState *string) {
	o.OriginsState = originsState
}

// WithOriginsSvmName adds the originsSvmName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsSvmName(originsSvmName *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsSvmName(originsSvmName)
	return o
}

// SetOriginsSvmName adds the originsSvmName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsSvmName(originsSvmName *string) {
	o.OriginsSvmName = originsSvmName
}

// WithOriginsSvmUUID adds the originsSvmUUID to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsSvmUUID(originsSvmUUID *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsSvmUUID(originsSvmUUID)
	return o
}

// SetOriginsSvmUUID adds the originsSvmUuid to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsSvmUUID(originsSvmUUID *string) {
	o.OriginsSvmUUID = originsSvmUUID
}

// WithOriginsVolumeName adds the originsVolumeName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsVolumeName(originsVolumeName *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsVolumeName(originsVolumeName)
	return o
}

// SetOriginsVolumeName adds the originsVolumeName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsVolumeName(originsVolumeName *string) {
	o.OriginsVolumeName = originsVolumeName
}

// WithOriginsVolumeUUID adds the originsVolumeUUID to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOriginsVolumeUUID(originsVolumeUUID *string) *FlexcacheDeleteCollectionParams {
	o.SetOriginsVolumeUUID(originsVolumeUUID)
	return o
}

// SetOriginsVolumeUUID adds the originsVolumeUuid to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOriginsVolumeUUID(originsVolumeUUID *string) {
	o.OriginsVolumeUUID = originsVolumeUUID
}

// WithOverrideEncryption adds the overrideEncryption to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithOverrideEncryption(overrideEncryption *bool) *FlexcacheDeleteCollectionParams {
	o.SetOverrideEncryption(overrideEncryption)
	return o
}

// SetOverrideEncryption adds the overrideEncryption to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetOverrideEncryption(overrideEncryption *bool) {
	o.OverrideEncryption = overrideEncryption
}

// WithPath adds the path to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithPath(path *string) *FlexcacheDeleteCollectionParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetPath(path *string) {
	o.Path = path
}

// WithRelativeSizeEnabled adds the relativeSizeEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithRelativeSizeEnabled(relativeSizeEnabled *bool) *FlexcacheDeleteCollectionParams {
	o.SetRelativeSizeEnabled(relativeSizeEnabled)
	return o
}

// SetRelativeSizeEnabled adds the relativeSizeEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetRelativeSizeEnabled(relativeSizeEnabled *bool) {
	o.RelativeSizeEnabled = relativeSizeEnabled
}

// WithRelativeSizePercentage adds the relativeSizePercentage to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithRelativeSizePercentage(relativeSizePercentage *int64) *FlexcacheDeleteCollectionParams {
	o.SetRelativeSizePercentage(relativeSizePercentage)
	return o
}

// SetRelativeSizePercentage adds the relativeSizePercentage to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetRelativeSizePercentage(relativeSizePercentage *int64) {
	o.RelativeSizePercentage = relativeSizePercentage
}

// WithReturnRecords adds the returnRecords to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *FlexcacheDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *FlexcacheDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *FlexcacheDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSize adds the size to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithSize(size *int64) *FlexcacheDeleteCollectionParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetSize(size *int64) {
	o.Size = size
}

// WithSvmName adds the svmName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithSvmName(svmName *string) *FlexcacheDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithSvmUUID(svmUUID *string) *FlexcacheDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUseTieredAggregate adds the useTieredAggregate to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithUseTieredAggregate(useTieredAggregate *bool) *FlexcacheDeleteCollectionParams {
	o.SetUseTieredAggregate(useTieredAggregate)
	return o
}

// SetUseTieredAggregate adds the useTieredAggregate to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetUseTieredAggregate(useTieredAggregate *bool) {
	o.UseTieredAggregate = useTieredAggregate
}

// WithUUID adds the uuid to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithUUID(uuid *string) *FlexcacheDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithWritebackEnabled adds the writebackEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) WithWritebackEnabled(writebackEnabled *bool) *FlexcacheDeleteCollectionParams {
	o.SetWritebackEnabled(writebackEnabled)
	return o
}

// SetWritebackEnabled adds the writebackEnabled to the flexcache delete collection params
func (o *FlexcacheDeleteCollectionParams) SetWritebackEnabled(writebackEnabled *bool) {
	o.WritebackEnabled = writebackEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *FlexcacheDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AggregatesName != nil {

		// query param aggregates.name
		var qrAggregatesName string

		if o.AggregatesName != nil {
			qrAggregatesName = *o.AggregatesName
		}
		qAggregatesName := qrAggregatesName
		if qAggregatesName != "" {

			if err := r.SetQueryParam("aggregates.name", qAggregatesName); err != nil {
				return err
			}
		}
	}

	if o.AggregatesUUID != nil {

		// query param aggregates.uuid
		var qrAggregatesUUID string

		if o.AggregatesUUID != nil {
			qrAggregatesUUID = *o.AggregatesUUID
		}
		qAggregatesUUID := qrAggregatesUUID
		if qAggregatesUUID != "" {

			if err := r.SetQueryParam("aggregates.uuid", qAggregatesUUID); err != nil {
				return err
			}
		}
	}

	if o.AtimeScrubEnabled != nil {

		// query param atime_scrub.enabled
		var qrAtimeScrubEnabled bool

		if o.AtimeScrubEnabled != nil {
			qrAtimeScrubEnabled = *o.AtimeScrubEnabled
		}
		qAtimeScrubEnabled := swag.FormatBool(qrAtimeScrubEnabled)
		if qAtimeScrubEnabled != "" {

			if err := r.SetQueryParam("atime_scrub.enabled", qAtimeScrubEnabled); err != nil {
				return err
			}
		}
	}

	if o.AtimeScrubPeriod != nil {

		// query param atime_scrub.period
		var qrAtimeScrubPeriod int64

		if o.AtimeScrubPeriod != nil {
			qrAtimeScrubPeriod = *o.AtimeScrubPeriod
		}
		qAtimeScrubPeriod := swag.FormatInt64(qrAtimeScrubPeriod)
		if qAtimeScrubPeriod != "" {

			if err := r.SetQueryParam("atime_scrub.period", qAtimeScrubPeriod); err != nil {
				return err
			}
		}
	}

	if o.CifsChangeNotifyEnabled != nil {

		// query param cifs_change_notify.enabled
		var qrCifsChangeNotifyEnabled bool

		if o.CifsChangeNotifyEnabled != nil {
			qrCifsChangeNotifyEnabled = *o.CifsChangeNotifyEnabled
		}
		qCifsChangeNotifyEnabled := swag.FormatBool(qrCifsChangeNotifyEnabled)
		if qCifsChangeNotifyEnabled != "" {

			if err := r.SetQueryParam("cifs_change_notify.enabled", qCifsChangeNotifyEnabled); err != nil {
				return err
			}
		}
	}

	if o.ConstituentsPerAggregate != nil {

		// query param constituents_per_aggregate
		var qrConstituentsPerAggregate int64

		if o.ConstituentsPerAggregate != nil {
			qrConstituentsPerAggregate = *o.ConstituentsPerAggregate
		}
		qConstituentsPerAggregate := swag.FormatInt64(qrConstituentsPerAggregate)
		if qConstituentsPerAggregate != "" {

			if err := r.SetQueryParam("constituents_per_aggregate", qConstituentsPerAggregate); err != nil {
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

	if o.DrCache != nil {

		// query param dr_cache
		var qrDrCache bool

		if o.DrCache != nil {
			qrDrCache = *o.DrCache
		}
		qDrCache := swag.FormatBool(qrDrCache)
		if qDrCache != "" {

			if err := r.SetQueryParam("dr_cache", qDrCache); err != nil {
				return err
			}
		}
	}

	if o.GlobalFileLockingEnabled != nil {

		// query param global_file_locking_enabled
		var qrGlobalFileLockingEnabled bool

		if o.GlobalFileLockingEnabled != nil {
			qrGlobalFileLockingEnabled = *o.GlobalFileLockingEnabled
		}
		qGlobalFileLockingEnabled := swag.FormatBool(qrGlobalFileLockingEnabled)
		if qGlobalFileLockingEnabled != "" {

			if err := r.SetQueryParam("global_file_locking_enabled", qGlobalFileLockingEnabled); err != nil {
				return err
			}
		}
	}

	if o.GuaranteeType != nil {

		// query param guarantee.type
		var qrGuaranteeType string

		if o.GuaranteeType != nil {
			qrGuaranteeType = *o.GuaranteeType
		}
		qGuaranteeType := qrGuaranteeType
		if qGuaranteeType != "" {

			if err := r.SetQueryParam("guarantee.type", qGuaranteeType); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OriginsClusterName != nil {

		// query param origins.cluster.name
		var qrOriginsClusterName string

		if o.OriginsClusterName != nil {
			qrOriginsClusterName = *o.OriginsClusterName
		}
		qOriginsClusterName := qrOriginsClusterName
		if qOriginsClusterName != "" {

			if err := r.SetQueryParam("origins.cluster.name", qOriginsClusterName); err != nil {
				return err
			}
		}
	}

	if o.OriginsClusterUUID != nil {

		// query param origins.cluster.uuid
		var qrOriginsClusterUUID string

		if o.OriginsClusterUUID != nil {
			qrOriginsClusterUUID = *o.OriginsClusterUUID
		}
		qOriginsClusterUUID := qrOriginsClusterUUID
		if qOriginsClusterUUID != "" {

			if err := r.SetQueryParam("origins.cluster.uuid", qOriginsClusterUUID); err != nil {
				return err
			}
		}
	}

	if o.OriginsCreateTime != nil {

		// query param origins.create_time
		var qrOriginsCreateTime string

		if o.OriginsCreateTime != nil {
			qrOriginsCreateTime = *o.OriginsCreateTime
		}
		qOriginsCreateTime := qrOriginsCreateTime
		if qOriginsCreateTime != "" {

			if err := r.SetQueryParam("origins.create_time", qOriginsCreateTime); err != nil {
				return err
			}
		}
	}

	if o.OriginsIPAddress != nil {

		// query param origins.ip_address
		var qrOriginsIPAddress string

		if o.OriginsIPAddress != nil {
			qrOriginsIPAddress = *o.OriginsIPAddress
		}
		qOriginsIPAddress := qrOriginsIPAddress
		if qOriginsIPAddress != "" {

			if err := r.SetQueryParam("origins.ip_address", qOriginsIPAddress); err != nil {
				return err
			}
		}
	}

	if o.OriginsSize != nil {

		// query param origins.size
		var qrOriginsSize int64

		if o.OriginsSize != nil {
			qrOriginsSize = *o.OriginsSize
		}
		qOriginsSize := swag.FormatInt64(qrOriginsSize)
		if qOriginsSize != "" {

			if err := r.SetQueryParam("origins.size", qOriginsSize); err != nil {
				return err
			}
		}
	}

	if o.OriginsState != nil {

		// query param origins.state
		var qrOriginsState string

		if o.OriginsState != nil {
			qrOriginsState = *o.OriginsState
		}
		qOriginsState := qrOriginsState
		if qOriginsState != "" {

			if err := r.SetQueryParam("origins.state", qOriginsState); err != nil {
				return err
			}
		}
	}

	if o.OriginsSvmName != nil {

		// query param origins.svm.name
		var qrOriginsSvmName string

		if o.OriginsSvmName != nil {
			qrOriginsSvmName = *o.OriginsSvmName
		}
		qOriginsSvmName := qrOriginsSvmName
		if qOriginsSvmName != "" {

			if err := r.SetQueryParam("origins.svm.name", qOriginsSvmName); err != nil {
				return err
			}
		}
	}

	if o.OriginsSvmUUID != nil {

		// query param origins.svm.uuid
		var qrOriginsSvmUUID string

		if o.OriginsSvmUUID != nil {
			qrOriginsSvmUUID = *o.OriginsSvmUUID
		}
		qOriginsSvmUUID := qrOriginsSvmUUID
		if qOriginsSvmUUID != "" {

			if err := r.SetQueryParam("origins.svm.uuid", qOriginsSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.OriginsVolumeName != nil {

		// query param origins.volume.name
		var qrOriginsVolumeName string

		if o.OriginsVolumeName != nil {
			qrOriginsVolumeName = *o.OriginsVolumeName
		}
		qOriginsVolumeName := qrOriginsVolumeName
		if qOriginsVolumeName != "" {

			if err := r.SetQueryParam("origins.volume.name", qOriginsVolumeName); err != nil {
				return err
			}
		}
	}

	if o.OriginsVolumeUUID != nil {

		// query param origins.volume.uuid
		var qrOriginsVolumeUUID string

		if o.OriginsVolumeUUID != nil {
			qrOriginsVolumeUUID = *o.OriginsVolumeUUID
		}
		qOriginsVolumeUUID := qrOriginsVolumeUUID
		if qOriginsVolumeUUID != "" {

			if err := r.SetQueryParam("origins.volume.uuid", qOriginsVolumeUUID); err != nil {
				return err
			}
		}
	}

	if o.OverrideEncryption != nil {

		// query param override_encryption
		var qrOverrideEncryption bool

		if o.OverrideEncryption != nil {
			qrOverrideEncryption = *o.OverrideEncryption
		}
		qOverrideEncryption := swag.FormatBool(qrOverrideEncryption)
		if qOverrideEncryption != "" {

			if err := r.SetQueryParam("override_encryption", qOverrideEncryption); err != nil {
				return err
			}
		}
	}

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	if o.RelativeSizeEnabled != nil {

		// query param relative_size.enabled
		var qrRelativeSizeEnabled bool

		if o.RelativeSizeEnabled != nil {
			qrRelativeSizeEnabled = *o.RelativeSizeEnabled
		}
		qRelativeSizeEnabled := swag.FormatBool(qrRelativeSizeEnabled)
		if qRelativeSizeEnabled != "" {

			if err := r.SetQueryParam("relative_size.enabled", qRelativeSizeEnabled); err != nil {
				return err
			}
		}
	}

	if o.RelativeSizePercentage != nil {

		// query param relative_size.percentage
		var qrRelativeSizePercentage int64

		if o.RelativeSizePercentage != nil {
			qrRelativeSizePercentage = *o.RelativeSizePercentage
		}
		qRelativeSizePercentage := swag.FormatInt64(qrRelativeSizePercentage)
		if qRelativeSizePercentage != "" {

			if err := r.SetQueryParam("relative_size.percentage", qRelativeSizePercentage); err != nil {
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

	if o.Size != nil {

		// query param size
		var qrSize int64

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
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

	if o.UseTieredAggregate != nil {

		// query param use_tiered_aggregate
		var qrUseTieredAggregate bool

		if o.UseTieredAggregate != nil {
			qrUseTieredAggregate = *o.UseTieredAggregate
		}
		qUseTieredAggregate := swag.FormatBool(qrUseTieredAggregate)
		if qUseTieredAggregate != "" {

			if err := r.SetQueryParam("use_tiered_aggregate", qUseTieredAggregate); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if o.WritebackEnabled != nil {

		// query param writeback.enabled
		var qrWritebackEnabled bool

		if o.WritebackEnabled != nil {
			qrWritebackEnabled = *o.WritebackEnabled
		}
		qWritebackEnabled := swag.FormatBool(qrWritebackEnabled)
		if qWritebackEnabled != "" {

			if err := r.SetQueryParam("writeback.enabled", qWritebackEnabled); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
