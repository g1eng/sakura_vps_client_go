# ProblemDetails400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | エラー内容を示す簡潔な識別子 * &#x60;invalid&#x60; - 不正なリクエスト値,リクエスト値が妥当でない * &#x60;parse_error&#x60; - 不正な形式,リクエスト値を読み取ることができない * &#x60;bad_request&#x60; - リクエストの内容に何らかの問題がある | [optional] 
**Message** | Pointer to **string** | エラーの内容 | [optional] 
**Errors** | Pointer to [**NullableProblemDetails400Errors**](ProblemDetails400Errors.md) |  | [optional] 

## Methods

### NewProblemDetails400

`func NewProblemDetails400() *ProblemDetails400`

NewProblemDetails400 instantiates a new ProblemDetails400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemDetails400WithDefaults

`func NewProblemDetails400WithDefaults() *ProblemDetails400`

NewProblemDetails400WithDefaults instantiates a new ProblemDetails400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ProblemDetails400) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProblemDetails400) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProblemDetails400) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProblemDetails400) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ProblemDetails400) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProblemDetails400) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProblemDetails400) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProblemDetails400) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *ProblemDetails400) GetErrors() ProblemDetails400Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ProblemDetails400) GetErrorsOk() (*ProblemDetails400Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ProblemDetails400) SetErrors(v ProblemDetails400Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ProblemDetails400) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ProblemDetails400) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ProblemDetails400) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


