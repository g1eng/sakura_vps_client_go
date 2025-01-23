# GetServerList403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | エラー内容を示す簡潔な識別子 | [optional] 
**Message** | Pointer to **string** | エラーの内容 | [optional] 

## Methods

### NewGetServerList403Response

`func NewGetServerList403Response() *GetServerList403Response`

NewGetServerList403Response instantiates a new GetServerList403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerList403ResponseWithDefaults

`func NewGetServerList403ResponseWithDefaults() *GetServerList403Response`

NewGetServerList403ResponseWithDefaults instantiates a new GetServerList403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetServerList403Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetServerList403Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetServerList403Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetServerList403Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetServerList403Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetServerList403Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetServerList403Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetServerList403Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


