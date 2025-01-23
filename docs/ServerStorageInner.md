# ServerStorageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | ポート番号 | 
**Type** | **string** | 種別 | 
**SizeGibibytes** | **int32** | ストレージ容量(GiB) | 

## Methods

### NewServerStorageInner

`func NewServerStorageInner(port int32, type_ string, sizeGibibytes int32, ) *ServerStorageInner`

NewServerStorageInner instantiates a new ServerStorageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStorageInnerWithDefaults

`func NewServerStorageInnerWithDefaults() *ServerStorageInner`

NewServerStorageInnerWithDefaults instantiates a new ServerStorageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ServerStorageInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServerStorageInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServerStorageInner) SetPort(v int32)`

SetPort sets Port field to given value.


### GetType

`func (o *ServerStorageInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerStorageInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerStorageInner) SetType(v string)`

SetType sets Type field to given value.


### GetSizeGibibytes

`func (o *ServerStorageInner) GetSizeGibibytes() int32`

GetSizeGibibytes returns the SizeGibibytes field if non-nil, zero value otherwise.

### GetSizeGibibytesOk

`func (o *ServerStorageInner) GetSizeGibibytesOk() (*int32, bool)`

GetSizeGibibytesOk returns a tuple with the SizeGibibytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGibibytes

`func (o *ServerStorageInner) SetSizeGibibytes(v int32)`

SetSizeGibibytes sets SizeGibibytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


