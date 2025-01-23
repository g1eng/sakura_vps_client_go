# GetServerMonitoringList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | データ総数 | 
**Next** | **NullableString** | 次のページへのURL | 
**Previous** | **NullableString** | 前のページへのURL | 
**Results** | [**[]ServerMonitoring**](ServerMonitoring.md) |  | 

## Methods

### NewGetServerMonitoringList200Response

`func NewGetServerMonitoringList200Response(count int32, next NullableString, previous NullableString, results []ServerMonitoring, ) *GetServerMonitoringList200Response`

NewGetServerMonitoringList200Response instantiates a new GetServerMonitoringList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerMonitoringList200ResponseWithDefaults

`func NewGetServerMonitoringList200ResponseWithDefaults() *GetServerMonitoringList200Response`

NewGetServerMonitoringList200ResponseWithDefaults instantiates a new GetServerMonitoringList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetServerMonitoringList200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetServerMonitoringList200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetServerMonitoringList200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *GetServerMonitoringList200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetServerMonitoringList200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetServerMonitoringList200Response) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *GetServerMonitoringList200Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *GetServerMonitoringList200Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *GetServerMonitoringList200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetServerMonitoringList200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetServerMonitoringList200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *GetServerMonitoringList200Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *GetServerMonitoringList200Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *GetServerMonitoringList200Response) GetResults() []ServerMonitoring`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetServerMonitoringList200Response) GetResultsOk() (*[]ServerMonitoring, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetServerMonitoringList200Response) SetResults(v []ServerMonitoring)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


