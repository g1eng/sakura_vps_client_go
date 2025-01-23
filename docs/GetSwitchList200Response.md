# GetSwitchList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | データ総数 | 
**Next** | **NullableString** | 次のページへのURL | 
**Previous** | **NullableString** | 前のページへのURL | 
**Results** | [**[]Switch**](Switch.md) |  | 

## Methods

### NewGetSwitchList200Response

`func NewGetSwitchList200Response(count int32, next NullableString, previous NullableString, results []Switch, ) *GetSwitchList200Response`

NewGetSwitchList200Response instantiates a new GetSwitchList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSwitchList200ResponseWithDefaults

`func NewGetSwitchList200ResponseWithDefaults() *GetSwitchList200Response`

NewGetSwitchList200ResponseWithDefaults instantiates a new GetSwitchList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetSwitchList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetSwitchList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetSwitchList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *GetSwitchList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetSwitchList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetSwitchList200Response) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *GetSwitchList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *GetSwitchList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *GetSwitchList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetSwitchList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetSwitchList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *GetSwitchList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *GetSwitchList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *GetSwitchList200Response) GetResults() []Switch`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSwitchList200Response) GetResultsOk() (*[]Switch, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSwitchList200Response) SetResults(v []Switch)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


