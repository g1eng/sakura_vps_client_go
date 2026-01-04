# PostServerStartMigrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerOnUponCompletion** | Pointer to **bool** | 移行完了時に自動で電源をONにする | [optional] [default to false]

## Methods

### NewPostServerStartMigrationRequest

`func NewPostServerStartMigrationRequest() *PostServerStartMigrationRequest`

NewPostServerStartMigrationRequest instantiates a new PostServerStartMigrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostServerStartMigrationRequestWithDefaults

`func NewPostServerStartMigrationRequestWithDefaults() *PostServerStartMigrationRequest`

NewPostServerStartMigrationRequestWithDefaults instantiates a new PostServerStartMigrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerOnUponCompletion

`func (o *PostServerStartMigrationRequest) GetPowerOnUponCompletion() bool`

GetPowerOnUponCompletion returns the PowerOnUponCompletion field if non-nil, zero value otherwise.

### GetPowerOnUponCompletionOk

`func (o *PostServerStartMigrationRequest) GetPowerOnUponCompletionOk() (*bool, bool)`

GetPowerOnUponCompletionOk returns a tuple with the PowerOnUponCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnUponCompletion

`func (o *PostServerStartMigrationRequest) SetPowerOnUponCompletion(v bool)`

SetPowerOnUponCompletion sets PowerOnUponCompletion field to given value.

### HasPowerOnUponCompletion

`func (o *PostServerStartMigrationRequest) HasPowerOnUponCompletion() bool`

HasPowerOnUponCompletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


