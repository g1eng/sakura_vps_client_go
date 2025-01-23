# HealthCheckSsh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | 監視方法 | 
**Port** | **int32** | ポート番号 | 
**IntervalMinutes** | **int32** | チェック間隔(分) | 

## Methods

### NewHealthCheckSsh

`func NewHealthCheckSsh(protocol string, port int32, intervalMinutes int32, ) *HealthCheckSsh`

NewHealthCheckSsh instantiates a new HealthCheckSsh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckSshWithDefaults

`func NewHealthCheckSshWithDefaults() *HealthCheckSsh`

NewHealthCheckSshWithDefaults instantiates a new HealthCheckSsh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *HealthCheckSsh) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheckSsh) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HealthCheckSsh) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *HealthCheckSsh) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheckSsh) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HealthCheckSsh) SetPort(v int32)`

SetPort sets Port field to given value.


### GetIntervalMinutes

`func (o *HealthCheckSsh) GetIntervalMinutes() int32`

GetIntervalMinutes returns the IntervalMinutes field if non-nil, zero value otherwise.

### GetIntervalMinutesOk

`func (o *HealthCheckSsh) GetIntervalMinutesOk() (*int32, bool)`

GetIntervalMinutesOk returns a tuple with the IntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMinutes

`func (o *HealthCheckSsh) SetIntervalMinutes(v int32)`

SetIntervalMinutes sets IntervalMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


