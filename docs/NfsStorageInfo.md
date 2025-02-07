# NfsStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeSpaceKib** | **int32** | ストレージの空き容量（KiB） | [readonly] 
**UsagePercentage** | **int32** | ストレージの使用率 | [readonly] 
**CapacityKib** | **int32** | ストレージの全容量（KiB） | [readonly] 
**UsageKib** | **int32** | ストレージの使用容量（KiB） | [readonly] 

## Methods

### NewNfsStorageInfo

`func NewNfsStorageInfo(freeSpaceKib int32, usagePercentage int32, capacityKib int32, usageKib int32, ) *NfsStorageInfo`

NewNfsStorageInfo instantiates a new NfsStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsStorageInfoWithDefaults

`func NewNfsStorageInfoWithDefaults() *NfsStorageInfo`

NewNfsStorageInfoWithDefaults instantiates a new NfsStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeSpaceKib

`func (o *NfsStorageInfo) GetFreeSpaceKib() int32`

GetFreeSpaceKib returns the FreeSpaceKib field if non-nil, zero value otherwise.

### GetFreeSpaceKibOk

`func (o *NfsStorageInfo) GetFreeSpaceKibOk() (*int32, bool)`

GetFreeSpaceKibOk returns a tuple with the FreeSpaceKib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceKib

`func (o *NfsStorageInfo) SetFreeSpaceKib(v int32)`

SetFreeSpaceKib sets FreeSpaceKib field to given value.


### GetUsagePercentage

`func (o *NfsStorageInfo) GetUsagePercentage() int32`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *NfsStorageInfo) GetUsagePercentageOk() (*int32, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *NfsStorageInfo) SetUsagePercentage(v int32)`

SetUsagePercentage sets UsagePercentage field to given value.


### GetCapacityKib

`func (o *NfsStorageInfo) GetCapacityKib() int32`

GetCapacityKib returns the CapacityKib field if non-nil, zero value otherwise.

### GetCapacityKibOk

`func (o *NfsStorageInfo) GetCapacityKibOk() (*int32, bool)`

GetCapacityKibOk returns a tuple with the CapacityKib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityKib

`func (o *NfsStorageInfo) SetCapacityKib(v int32)`

SetCapacityKib sets CapacityKib field to given value.


### GetUsageKib

`func (o *NfsStorageInfo) GetUsageKib() int32`

GetUsageKib returns the UsageKib field if non-nil, zero value otherwise.

### GetUsageKibOk

`func (o *NfsStorageInfo) GetUsageKibOk() (*int32, bool)`

GetUsageKibOk returns a tuple with the UsageKib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageKib

`func (o *NfsStorageInfo) SetUsageKib(v int32)`

SetUsageKib sets UsageKib field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


