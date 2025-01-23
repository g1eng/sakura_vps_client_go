# NfsServerStorageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | 種別 | 
**SizeGibibytes** | **int32** | ストレージ容量(GiB) | 

## Methods

### NewNfsServerStorageInner

`func NewNfsServerStorageInner(type_ string, sizeGibibytes int32, ) *NfsServerStorageInner`

NewNfsServerStorageInner instantiates a new NfsServerStorageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsServerStorageInnerWithDefaults

`func NewNfsServerStorageInnerWithDefaults() *NfsServerStorageInner`

NewNfsServerStorageInnerWithDefaults instantiates a new NfsServerStorageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NfsServerStorageInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NfsServerStorageInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NfsServerStorageInner) SetType(v string)`

SetType sets Type field to given value.


### GetSizeGibibytes

`func (o *NfsServerStorageInner) GetSizeGibibytes() int32`

GetSizeGibibytes returns the SizeGibibytes field if non-nil, zero value otherwise.

### GetSizeGibibytesOk

`func (o *NfsServerStorageInner) GetSizeGibibytesOk() (*int32, bool)`

GetSizeGibibytesOk returns a tuple with the SizeGibibytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGibibytes

`func (o *NfsServerStorageInner) SetSizeGibibytes(v int32)`

SetSizeGibibytes sets SizeGibibytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


