# GetNfsServerList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | データ総数 | 
**Next** | **NullableString** | 次のページへのURL | 
**Previous** | **NullableString** | 前のページへのURL | 
**Results** | [**[]NfsServer**](NfsServer.md) |  | 

## Methods

### NewGetNfsServerList200Response

`func NewGetNfsServerList200Response(count int32, next NullableString, previous NullableString, results []NfsServer, ) *GetNfsServerList200Response`

NewGetNfsServerList200Response instantiates a new GetNfsServerList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNfsServerList200ResponseWithDefaults

`func NewGetNfsServerList200ResponseWithDefaults() *GetNfsServerList200Response`

NewGetNfsServerList200ResponseWithDefaults instantiates a new GetNfsServerList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetNfsServerList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetNfsServerList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetNfsServerList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *GetNfsServerList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetNfsServerList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetNfsServerList200Response) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *GetNfsServerList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *GetNfsServerList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *GetNfsServerList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetNfsServerList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetNfsServerList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *GetNfsServerList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *GetNfsServerList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *GetNfsServerList200Response) GetResults() []NfsServer`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetNfsServerList200Response) GetResultsOk() (*[]NfsServer, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetNfsServerList200Response) SetResults(v []NfsServer)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


