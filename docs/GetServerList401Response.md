# GetServerList401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | エラー内容を示す簡潔な識別子 | [optional] 
**Message** | Pointer to **string** | エラーの内容 | [optional] 

## Methods

### NewGetServerList401Response

`func NewGetServerList401Response() *GetServerList401Response`

NewGetServerList401Response instantiates a new GetServerList401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerList401ResponseWithDefaults

`func NewGetServerList401ResponseWithDefaults() *GetServerList401Response`

NewGetServerList401ResponseWithDefaults instantiates a new GetServerList401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetServerList401Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetServerList401Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetServerList401Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetServerList401Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetServerList401Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetServerList401Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetServerList401Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetServerList401Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


