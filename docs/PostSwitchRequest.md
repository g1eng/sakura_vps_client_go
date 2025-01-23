# PostSwitchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**ZoneCode** | **string** | ゾーンコード * tk2 東京第2 * tk3 東京第3 * os3 大阪第3 * is1 石狩第1 | 

## Methods

### NewPostSwitchRequest

`func NewPostSwitchRequest(name string, description string, zoneCode string, ) *PostSwitchRequest`

NewPostSwitchRequest instantiates a new PostSwitchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSwitchRequestWithDefaults

`func NewPostSwitchRequestWithDefaults() *PostSwitchRequest`

NewPostSwitchRequestWithDefaults instantiates a new PostSwitchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostSwitchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostSwitchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostSwitchRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostSwitchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostSwitchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostSwitchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetZoneCode

`func (o *PostSwitchRequest) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *PostSwitchRequest) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *PostSwitchRequest) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


