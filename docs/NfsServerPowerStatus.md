# NfsServerPowerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | 電源ステータス * power_on 電源ON * in_shutdown シャットダウン中 * power_off 電源OFF * unknown 不明（電源状態を取得できない） | 

## Methods

### NewNfsServerPowerStatus

`func NewNfsServerPowerStatus(status string, ) *NfsServerPowerStatus`

NewNfsServerPowerStatus instantiates a new NfsServerPowerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsServerPowerStatusWithDefaults

`func NewNfsServerPowerStatusWithDefaults() *NfsServerPowerStatus`

NewNfsServerPowerStatusWithDefaults instantiates a new NfsServerPowerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NfsServerPowerStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NfsServerPowerStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NfsServerPowerStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


