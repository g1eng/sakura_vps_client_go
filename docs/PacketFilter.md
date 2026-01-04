# PacketFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | パケットフィルタのON/OFF * true ON * false OFF | 
**Rules** | Pointer to [**[]PacketFilterRule**](PacketFilterRule.md) | アクセスを許可するルールのリスト - 制限   - パケットフィルタOFF(enabled: false)の場合は指定できません   - 指定できるルール数には制限があります。全ルールのポート定義数の合計が40個まで定義できます。       ポート個別指定\&quot;80\&quot;とポート範囲指定\&quot;8080:8085\&quot;はそれぞれ1個とカウントします。 | [optional] 

## Methods

### NewPacketFilter

`func NewPacketFilter(enabled bool, ) *PacketFilter`

NewPacketFilter instantiates a new PacketFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketFilterWithDefaults

`func NewPacketFilterWithDefaults() *PacketFilter`

NewPacketFilterWithDefaults instantiates a new PacketFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PacketFilter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PacketFilter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PacketFilter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRules

`func (o *PacketFilter) GetRules() []PacketFilterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PacketFilter) GetRulesOk() (*[]PacketFilterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PacketFilter) SetRules(v []PacketFilterRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PacketFilter) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


