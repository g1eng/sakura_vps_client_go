# GetRoleList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | データ総数 | 
**Next** | **NullableString** | 次のページへのURL | 
**Previous** | **NullableString** | 前のページへのURL | 
**Results** | [**[]Role**](Role.md) |  | 

## Methods

### NewGetRoleList200Response

`func NewGetRoleList200Response(count int32, next NullableString, previous NullableString, results []Role, ) *GetRoleList200Response`

NewGetRoleList200Response instantiates a new GetRoleList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleList200ResponseWithDefaults

`func NewGetRoleList200ResponseWithDefaults() *GetRoleList200Response`

NewGetRoleList200ResponseWithDefaults instantiates a new GetRoleList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetRoleList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetRoleList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetRoleList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *GetRoleList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetRoleList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetRoleList200Response) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *GetRoleList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *GetRoleList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *GetRoleList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetRoleList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetRoleList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *GetRoleList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *GetRoleList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *GetRoleList200Response) GetResults() []Role`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetRoleList200Response) GetResultsOk() (*[]Role, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetRoleList200Response) SetResults(v []Role)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


