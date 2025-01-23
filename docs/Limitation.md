# Limitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuPerformanceLimit** | Pointer to **string** | CPUリソースの制限可否 | [optional] 
**NetworkBandwidthLimit** | Pointer to **string** | ネットワーク帯域の制限可否 | [optional] 
**OutboundPort25Blocking** | Pointer to **string** | OP25Bの可否 | [optional] 
**StorageIopsLimit** | Pointer to **string** | ストレージのIOPSの制限可否 | [optional] 

## Methods

### NewLimitation

`func NewLimitation() *Limitation`

NewLimitation instantiates a new Limitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitationWithDefaults

`func NewLimitationWithDefaults() *Limitation`

NewLimitationWithDefaults instantiates a new Limitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuPerformanceLimit

`func (o *Limitation) GetCpuPerformanceLimit() string`

GetCpuPerformanceLimit returns the CpuPerformanceLimit field if non-nil, zero value otherwise.

### GetCpuPerformanceLimitOk

`func (o *Limitation) GetCpuPerformanceLimitOk() (*string, bool)`

GetCpuPerformanceLimitOk returns a tuple with the CpuPerformanceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPerformanceLimit

`func (o *Limitation) SetCpuPerformanceLimit(v string)`

SetCpuPerformanceLimit sets CpuPerformanceLimit field to given value.

### HasCpuPerformanceLimit

`func (o *Limitation) HasCpuPerformanceLimit() bool`

HasCpuPerformanceLimit returns a boolean if a field has been set.

### GetNetworkBandwidthLimit

`func (o *Limitation) GetNetworkBandwidthLimit() string`

GetNetworkBandwidthLimit returns the NetworkBandwidthLimit field if non-nil, zero value otherwise.

### GetNetworkBandwidthLimitOk

`func (o *Limitation) GetNetworkBandwidthLimitOk() (*string, bool)`

GetNetworkBandwidthLimitOk returns a tuple with the NetworkBandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidthLimit

`func (o *Limitation) SetNetworkBandwidthLimit(v string)`

SetNetworkBandwidthLimit sets NetworkBandwidthLimit field to given value.

### HasNetworkBandwidthLimit

`func (o *Limitation) HasNetworkBandwidthLimit() bool`

HasNetworkBandwidthLimit returns a boolean if a field has been set.

### GetOutboundPort25Blocking

`func (o *Limitation) GetOutboundPort25Blocking() string`

GetOutboundPort25Blocking returns the OutboundPort25Blocking field if non-nil, zero value otherwise.

### GetOutboundPort25BlockingOk

`func (o *Limitation) GetOutboundPort25BlockingOk() (*string, bool)`

GetOutboundPort25BlockingOk returns a tuple with the OutboundPort25Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPort25Blocking

`func (o *Limitation) SetOutboundPort25Blocking(v string)`

SetOutboundPort25Blocking sets OutboundPort25Blocking field to given value.

### HasOutboundPort25Blocking

`func (o *Limitation) HasOutboundPort25Blocking() bool`

HasOutboundPort25Blocking returns a boolean if a field has been set.

### GetStorageIopsLimit

`func (o *Limitation) GetStorageIopsLimit() string`

GetStorageIopsLimit returns the StorageIopsLimit field if non-nil, zero value otherwise.

### GetStorageIopsLimitOk

`func (o *Limitation) GetStorageIopsLimitOk() (*string, bool)`

GetStorageIopsLimitOk returns a tuple with the StorageIopsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIopsLimit

`func (o *Limitation) SetStorageIopsLimit(v string)`

SetStorageIopsLimit sets StorageIopsLimit field to given value.

### HasStorageIopsLimit

`func (o *Limitation) HasStorageIopsLimit() bool`

HasStorageIopsLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


