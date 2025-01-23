# ServerInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [readonly] 
**Device** | **string** | NIC名称 | [readonly] 
**ConnectableToGlobalNetwork** | **bool** | グローバルネットワークと接続可能か | [readonly] 
**ConnectTo** | **NullableString** | インターフェースの接続先 | [readonly] 
**Mac** | **string** | MACアドレス | [readonly] 
**SwitchId** | **NullableInt32** | スイッチID | 

## Methods

### NewServerInterface

`func NewServerInterface(id int32, device string, connectableToGlobalNetwork bool, connectTo NullableString, mac string, switchId NullableInt32, ) *ServerInterface`

NewServerInterface instantiates a new ServerInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInterfaceWithDefaults

`func NewServerInterfaceWithDefaults() *ServerInterface`

NewServerInterfaceWithDefaults instantiates a new ServerInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInterface) SetId(v int32)`

SetId sets Id field to given value.


### GetDevice

`func (o *ServerInterface) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ServerInterface) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ServerInterface) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetConnectableToGlobalNetwork

`func (o *ServerInterface) GetConnectableToGlobalNetwork() bool`

GetConnectableToGlobalNetwork returns the ConnectableToGlobalNetwork field if non-nil, zero value otherwise.

### GetConnectableToGlobalNetworkOk

`func (o *ServerInterface) GetConnectableToGlobalNetworkOk() (*bool, bool)`

GetConnectableToGlobalNetworkOk returns a tuple with the ConnectableToGlobalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectableToGlobalNetwork

`func (o *ServerInterface) SetConnectableToGlobalNetwork(v bool)`

SetConnectableToGlobalNetwork sets ConnectableToGlobalNetwork field to given value.


### GetConnectTo

`func (o *ServerInterface) GetConnectTo() string`

GetConnectTo returns the ConnectTo field if non-nil, zero value otherwise.

### GetConnectToOk

`func (o *ServerInterface) GetConnectToOk() (*string, bool)`

GetConnectToOk returns a tuple with the ConnectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTo

`func (o *ServerInterface) SetConnectTo(v string)`

SetConnectTo sets ConnectTo field to given value.


### SetConnectToNil

`func (o *ServerInterface) SetConnectToNil(b bool)`

 SetConnectToNil sets the value for ConnectTo to be an explicit nil

### UnsetConnectTo
`func (o *ServerInterface) UnsetConnectTo()`

UnsetConnectTo ensures that no value is present for ConnectTo, not even an explicit nil
### GetMac

`func (o *ServerInterface) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ServerInterface) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ServerInterface) SetMac(v string)`

SetMac sets Mac field to given value.


### GetSwitchId

`func (o *ServerInterface) GetSwitchId() int32`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *ServerInterface) GetSwitchIdOk() (*int32, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *ServerInterface) SetSwitchId(v int32)`

SetSwitchId sets SwitchId field to given value.


### SetSwitchIdNil

`func (o *ServerInterface) SetSwitchIdNil(b bool)`

 SetSwitchIdNil sets the value for SwitchId to be an explicit nil

### UnsetSwitchId
`func (o *ServerInterface) UnsetSwitchId()`

UnsetSwitchId ensures that no value is present for SwitchId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


