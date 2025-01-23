# SwitchExternalConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceCode** | **string** | サービスコード | [readonly] 
**Type** | **string** | 外部接続方式 | [readonly] 
**Services** | [**[]SwitchExternalConnectionServicesInner**](SwitchExternalConnectionServicesInner.md) |  | [readonly] 

## Methods

### NewSwitchExternalConnection

`func NewSwitchExternalConnection(serviceCode string, type_ string, services []SwitchExternalConnectionServicesInner, ) *SwitchExternalConnection`

NewSwitchExternalConnection instantiates a new SwitchExternalConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchExternalConnectionWithDefaults

`func NewSwitchExternalConnectionWithDefaults() *SwitchExternalConnection`

NewSwitchExternalConnectionWithDefaults instantiates a new SwitchExternalConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceCode

`func (o *SwitchExternalConnection) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *SwitchExternalConnection) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *SwitchExternalConnection) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetType

`func (o *SwitchExternalConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchExternalConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchExternalConnection) SetType(v string)`

SetType sets Type field to given value.


### GetServices

`func (o *SwitchExternalConnection) GetServices() []SwitchExternalConnectionServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *SwitchExternalConnection) GetServicesOk() (*[]SwitchExternalConnectionServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *SwitchExternalConnection) SetServices(v []SwitchExternalConnectionServicesInner)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


