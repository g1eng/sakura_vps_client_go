# Switch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [readonly] 
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**SwitchCode** | **string** | スイッチコード | [readonly] 
**Zone** | [**SwitchZone**](SwitchZone.md) |  | 
**ServerInterfaces** | **[]int32** | 接続されているサーバーのインターフェースid | [readonly] 
**NfsServerInterfaces** | **[]int32** | 接続されている追加ストレージ（NFS）のインターフェースid | [readonly] 
**ExternalConnection** | [**NullableSwitchExternalConnection**](SwitchExternalConnection.md) |  | 

## Methods

### NewSwitch

`func NewSwitch(id int32, name string, description string, switchCode string, zone SwitchZone, serverInterfaces []int32, nfsServerInterfaces []int32, externalConnection NullableSwitchExternalConnection, ) *Switch`

NewSwitch instantiates a new Switch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchWithDefaults

`func NewSwitchWithDefaults() *Switch`

NewSwitchWithDefaults instantiates a new Switch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Switch) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Switch) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Switch) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Switch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Switch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Switch) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Switch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Switch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Switch) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSwitchCode

`func (o *Switch) GetSwitchCode() string`

GetSwitchCode returns the SwitchCode field if non-nil, zero value otherwise.

### GetSwitchCodeOk

`func (o *Switch) GetSwitchCodeOk() (*string, bool)`

GetSwitchCodeOk returns a tuple with the SwitchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchCode

`func (o *Switch) SetSwitchCode(v string)`

SetSwitchCode sets SwitchCode field to given value.


### GetZone

`func (o *Switch) GetZone() SwitchZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Switch) GetZoneOk() (*SwitchZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Switch) SetZone(v SwitchZone)`

SetZone sets Zone field to given value.


### GetServerInterfaces

`func (o *Switch) GetServerInterfaces() []int32`

GetServerInterfaces returns the ServerInterfaces field if non-nil, zero value otherwise.

### GetServerInterfacesOk

`func (o *Switch) GetServerInterfacesOk() (*[]int32, bool)`

GetServerInterfacesOk returns a tuple with the ServerInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaces

`func (o *Switch) SetServerInterfaces(v []int32)`

SetServerInterfaces sets ServerInterfaces field to given value.


### GetNfsServerInterfaces

`func (o *Switch) GetNfsServerInterfaces() []int32`

GetNfsServerInterfaces returns the NfsServerInterfaces field if non-nil, zero value otherwise.

### GetNfsServerInterfacesOk

`func (o *Switch) GetNfsServerInterfacesOk() (*[]int32, bool)`

GetNfsServerInterfacesOk returns a tuple with the NfsServerInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsServerInterfaces

`func (o *Switch) SetNfsServerInterfaces(v []int32)`

SetNfsServerInterfaces sets NfsServerInterfaces field to given value.


### GetExternalConnection

`func (o *Switch) GetExternalConnection() SwitchExternalConnection`

GetExternalConnection returns the ExternalConnection field if non-nil, zero value otherwise.

### GetExternalConnectionOk

`func (o *Switch) GetExternalConnectionOk() (*SwitchExternalConnection, bool)`

GetExternalConnectionOk returns a tuple with the ExternalConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnection

`func (o *Switch) SetExternalConnection(v SwitchExternalConnection)`

SetExternalConnection sets ExternalConnection field to given value.


### SetExternalConnectionNil

`func (o *Switch) SetExternalConnectionNil(b bool)`

 SetExternalConnectionNil sets the value for ExternalConnection to be an explicit nil

### UnsetExternalConnection
`func (o *Switch) UnsetExternalConnection()`

UnsetExternalConnection ensures that no value is present for ExternalConnection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


