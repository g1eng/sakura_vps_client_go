# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | 
**Code** | **string** | ゾーンコード * tk1 東京第1 * tk2 東京第2 * tk3 東京第3 * os1 大阪第1 * os2 大阪第2 * os3 大阪第3 * is1 石狩第1 | 
**Name** | **string** | ゾーン名称 | 
**CanUseLocal** | **bool** | ローカルネットワーク接続が可能かどうか | 
**CanUseHybrid** | **bool** | ハイブリッド接続が利用可能かどうか | 

## Methods

### NewZone

`func NewZone(id int32, code string, name string, canUseLocal bool, canUseHybrid bool, ) *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Zone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v int32)`

SetId sets Id field to given value.


### GetCode

`func (o *Zone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Zone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Zone) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *Zone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Zone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Zone) SetName(v string)`

SetName sets Name field to given value.


### GetCanUseLocal

`func (o *Zone) GetCanUseLocal() bool`

GetCanUseLocal returns the CanUseLocal field if non-nil, zero value otherwise.

### GetCanUseLocalOk

`func (o *Zone) GetCanUseLocalOk() (*bool, bool)`

GetCanUseLocalOk returns a tuple with the CanUseLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseLocal

`func (o *Zone) SetCanUseLocal(v bool)`

SetCanUseLocal sets CanUseLocal field to given value.


### GetCanUseHybrid

`func (o *Zone) GetCanUseHybrid() bool`

GetCanUseHybrid returns the CanUseHybrid field if non-nil, zero value otherwise.

### GetCanUseHybridOk

`func (o *Zone) GetCanUseHybridOk() (*bool, bool)`

GetCanUseHybridOk returns a tuple with the CanUseHybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUseHybrid

`func (o *Zone) SetCanUseHybrid(v bool)`

SetCanUseHybrid sets CanUseHybrid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


