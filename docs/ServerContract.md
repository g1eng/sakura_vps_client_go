# ServerContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanCode** | **int32** | プランコード | 
**PlanName** | **string** | プラン名 | 
**ServiceCode** | **string** | サービスコード | 

## Methods

### NewServerContract

`func NewServerContract(planCode int32, planName string, serviceCode string, ) *ServerContract`

NewServerContract instantiates a new ServerContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerContractWithDefaults

`func NewServerContractWithDefaults() *ServerContract`

NewServerContractWithDefaults instantiates a new ServerContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanCode

`func (o *ServerContract) GetPlanCode() int32`

GetPlanCode returns the PlanCode field if non-nil, zero value otherwise.

### GetPlanCodeOk

`func (o *ServerContract) GetPlanCodeOk() (*int32, bool)`

GetPlanCodeOk returns a tuple with the PlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCode

`func (o *ServerContract) SetPlanCode(v int32)`

SetPlanCode sets PlanCode field to given value.


### GetPlanName

`func (o *ServerContract) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *ServerContract) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *ServerContract) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetServiceCode

`func (o *ServerContract) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *ServerContract) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *ServerContract) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


