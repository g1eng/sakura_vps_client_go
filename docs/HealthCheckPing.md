# HealthCheckPing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | 監視方法 | 
**IntervalMinutes** | **int32** | チェック間隔(分) | 

## Methods

### NewHealthCheckPing

`func NewHealthCheckPing(protocol string, intervalMinutes int32, ) *HealthCheckPing`

NewHealthCheckPing instantiates a new HealthCheckPing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckPingWithDefaults

`func NewHealthCheckPingWithDefaults() *HealthCheckPing`

NewHealthCheckPingWithDefaults instantiates a new HealthCheckPing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *HealthCheckPing) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheckPing) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HealthCheckPing) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetIntervalMinutes

`func (o *HealthCheckPing) GetIntervalMinutes() int32`

GetIntervalMinutes returns the IntervalMinutes field if non-nil, zero value otherwise.

### GetIntervalMinutesOk

`func (o *HealthCheckPing) GetIntervalMinutesOk() (*int32, bool)`

GetIntervalMinutesOk returns a tuple with the IntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMinutes

`func (o *HealthCheckPing) SetIntervalMinutes(v int32)`

SetIntervalMinutes sets IntervalMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


