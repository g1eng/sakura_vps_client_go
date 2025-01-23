# PostSwitch201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [readonly] 
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**SwitchCode** | **string** | スイッチコード | [readonly] 
**Zone** | [**PostSwitch201ResponseZone**](PostSwitch201ResponseZone.md) |  | 
**ServerInterfaces** | **[]int32** | 接続されているサーバーのインターフェースid | [readonly] 
**NfsServerInterfaces** | **[]int32** | 接続されている追加ストレージ（NFS）のインターフェースid | [readonly] 
**ExternalConnection** | **map[string]interface{}** | 接続されている外部接続の情報 | [readonly] 

## Methods

### NewPostSwitch201Response

`func NewPostSwitch201Response(id int32, name string, description string, switchCode string, zone PostSwitch201ResponseZone, serverInterfaces []int32, nfsServerInterfaces []int32, externalConnection map[string]interface{}, ) *PostSwitch201Response`

NewPostSwitch201Response instantiates a new PostSwitch201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSwitch201ResponseWithDefaults

`func NewPostSwitch201ResponseWithDefaults() *PostSwitch201Response`

NewPostSwitch201ResponseWithDefaults instantiates a new PostSwitch201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostSwitch201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostSwitch201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostSwitch201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PostSwitch201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostSwitch201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostSwitch201Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostSwitch201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostSwitch201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostSwitch201Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSwitchCode

`func (o *PostSwitch201Response) GetSwitchCode() string`

GetSwitchCode returns the SwitchCode field if non-nil, zero value otherwise.

### GetSwitchCodeOk

`func (o *PostSwitch201Response) GetSwitchCodeOk() (*string, bool)`

GetSwitchCodeOk returns a tuple with the SwitchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchCode

`func (o *PostSwitch201Response) SetSwitchCode(v string)`

SetSwitchCode sets SwitchCode field to given value.


### GetZone

`func (o *PostSwitch201Response) GetZone() PostSwitch201ResponseZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PostSwitch201Response) GetZoneOk() (*PostSwitch201ResponseZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PostSwitch201Response) SetZone(v PostSwitch201ResponseZone)`

SetZone sets Zone field to given value.


### GetServerInterfaces

`func (o *PostSwitch201Response) GetServerInterfaces() []int32`

GetServerInterfaces returns the ServerInterfaces field if non-nil, zero value otherwise.

### GetServerInterfacesOk

`func (o *PostSwitch201Response) GetServerInterfacesOk() (*[]int32, bool)`

GetServerInterfacesOk returns a tuple with the ServerInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaces

`func (o *PostSwitch201Response) SetServerInterfaces(v []int32)`

SetServerInterfaces sets ServerInterfaces field to given value.


### GetNfsServerInterfaces

`func (o *PostSwitch201Response) GetNfsServerInterfaces() []int32`

GetNfsServerInterfaces returns the NfsServerInterfaces field if non-nil, zero value otherwise.

### GetNfsServerInterfacesOk

`func (o *PostSwitch201Response) GetNfsServerInterfacesOk() (*[]int32, bool)`

GetNfsServerInterfacesOk returns a tuple with the NfsServerInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsServerInterfaces

`func (o *PostSwitch201Response) SetNfsServerInterfaces(v []int32)`

SetNfsServerInterfaces sets NfsServerInterfaces field to given value.


### GetExternalConnection

`func (o *PostSwitch201Response) GetExternalConnection() map[string]interface{}`

GetExternalConnection returns the ExternalConnection field if non-nil, zero value otherwise.

### GetExternalConnectionOk

`func (o *PostSwitch201Response) GetExternalConnectionOk() (*map[string]interface{}, bool)`

GetExternalConnectionOk returns a tuple with the ExternalConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnection

`func (o *PostSwitch201Response) SetExternalConnection(v map[string]interface{})`

SetExternalConnection sets ExternalConnection field to given value.


### SetExternalConnectionNil

`func (o *PostSwitch201Response) SetExternalConnectionNil(b bool)`

 SetExternalConnectionNil sets the value for ExternalConnection to be an explicit nil

### UnsetExternalConnection
`func (o *PostSwitch201Response) UnsetExternalConnection()`

UnsetExternalConnection ensures that no value is present for ExternalConnection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


