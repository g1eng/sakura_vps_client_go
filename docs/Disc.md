# Disc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | 
**Name** | **string** | 名前 | 
**Description** | **string** | 説明 | 
**LicenseRequired** | **bool** | ライセンスが必要かどうか | 

## Methods

### NewDisc

`func NewDisc(id int32, name string, description string, licenseRequired bool, ) *Disc`

NewDisc instantiates a new Disc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscWithDefaults

`func NewDiscWithDefaults() *Disc`

NewDiscWithDefaults instantiates a new Disc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Disc) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disc) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disc) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Disc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Disc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Disc) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Disc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Disc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Disc) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLicenseRequired

`func (o *Disc) GetLicenseRequired() bool`

GetLicenseRequired returns the LicenseRequired field if non-nil, zero value otherwise.

### GetLicenseRequiredOk

`func (o *Disc) GetLicenseRequiredOk() (*bool, bool)`

GetLicenseRequiredOk returns a tuple with the LicenseRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRequired

`func (o *Disc) SetLicenseRequired(v bool)`

SetLicenseRequired sets LicenseRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


