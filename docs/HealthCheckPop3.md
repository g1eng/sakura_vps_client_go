# HealthCheckPop3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | 監視方法 | 
**Port** | **int32** | ポート番号 | 
**IntervalMinutes** | **int32** | チェック間隔(分) | 

## Methods

### NewHealthCheckPop3

`func NewHealthCheckPop3(protocol string, port int32, intervalMinutes int32, ) *HealthCheckPop3`

NewHealthCheckPop3 instantiates a new HealthCheckPop3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckPop3WithDefaults

`func NewHealthCheckPop3WithDefaults() *HealthCheckPop3`

NewHealthCheckPop3WithDefaults instantiates a new HealthCheckPop3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *HealthCheckPop3) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheckPop3) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HealthCheckPop3) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *HealthCheckPop3) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheckPop3) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HealthCheckPop3) SetPort(v int32)`

SetPort sets Port field to given value.


### GetIntervalMinutes

`func (o *HealthCheckPop3) GetIntervalMinutes() int32`

GetIntervalMinutes returns the IntervalMinutes field if non-nil, zero value otherwise.

### GetIntervalMinutesOk

`func (o *HealthCheckPop3) GetIntervalMinutesOk() (*int32, bool)`

GetIntervalMinutesOk returns a tuple with the IntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMinutes

`func (o *HealthCheckPop3) SetIntervalMinutes(v int32)`

SetIntervalMinutes sets IntervalMinutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


