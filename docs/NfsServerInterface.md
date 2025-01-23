# NfsServerInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [readonly] 
**Device** | **string** | NIC名称 | [readonly] 
**ConnectTo** | **NullableString** | インターフェースの接続先 | [readonly] 
**Mac** | **string** | MACアドレス | [readonly] 
**SwitchId** | **NullableInt32** | スイッチID | 

## Methods

### NewNfsServerInterface

`func NewNfsServerInterface(id int32, device string, connectTo NullableString, mac string, switchId NullableInt32, ) *NfsServerInterface`

NewNfsServerInterface instantiates a new NfsServerInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsServerInterfaceWithDefaults

`func NewNfsServerInterfaceWithDefaults() *NfsServerInterface`

NewNfsServerInterfaceWithDefaults instantiates a new NfsServerInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NfsServerInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NfsServerInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NfsServerInterface) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *NfsServerInterface) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NfsServerInterface) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NfsServerInterface) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetConnectTo

`func (o *NfsServerInterface) GetConnectTo() string`

GetConnectTo returns the ConnectTo field if non-nil, zero value otherwise.

### GetConnectToOk

`func (o *NfsServerInterface) GetConnectToOk() (*string, bool)`

GetConnectToOk returns a tuple with the ConnectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTo

`func (o *NfsServerInterface) SetConnectTo(v string)`

SetConnectTo sets ConnectTo field to given value.


### SetConnectToNil

`func (o *NfsServerInterface) SetConnectToNil(b bool)`

 SetConnectToNil sets the value for ConnectTo to be an explicit nil

### UnsetConnectTo
`func (o *NfsServerInterface) UnsetConnectTo()`

UnsetConnectTo ensures that no value is present for ConnectTo, not even an explicit nil
### GetMac

`func (o *NfsServerInterface) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NfsServerInterface) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NfsServerInterface) SetMac(v string)`

SetMac sets Mac field to given value.


### GetSwitchId

`func (o *NfsServerInterface) GetSwitchId() int32`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NfsServerInterface) GetSwitchIdOk() (*int32, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NfsServerInterface) SetSwitchId(v int32)`

SetSwitchId sets SwitchId field to given value.


### SetSwitchIdNil

`func (o *NfsServerInterface) SetSwitchIdNil(b bool)`

 SetSwitchIdNil sets the value for SwitchId to be an explicit nil

### UnsetSwitchId
`func (o *NfsServerInterface) UnsetSwitchId()`

UnsetSwitchId ensures that no value is present for SwitchId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


