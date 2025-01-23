# Keymap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | **string** | 指定したいキー配列の名称 | 

## Methods

### NewKeymap

`func NewKeymap(layout string, ) *Keymap`

NewKeymap instantiates a new Keymap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeymapWithDefaults

`func NewKeymapWithDefaults() *Keymap`

NewKeymapWithDefaults instantiates a new Keymap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *Keymap) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Keymap) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Keymap) SetLayout(v string)`

SetLayout sets Layout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


