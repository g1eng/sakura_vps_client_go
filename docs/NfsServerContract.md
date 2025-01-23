# NfsServerContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanCode** | **int32** | プランコード | 
**PlanName** | **string** | プラン名 | 
**ServiceCode** | **string** | サービスコード | 

## Methods

### NewNfsServerContract

`func NewNfsServerContract(planCode int32, planName string, serviceCode string, ) *NfsServerContract`

NewNfsServerContract instantiates a new NfsServerContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsServerContractWithDefaults

`func NewNfsServerContractWithDefaults() *NfsServerContract`

NewNfsServerContractWithDefaults instantiates a new NfsServerContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanCode

`func (o *NfsServerContract) GetPlanCode() int32`

GetPlanCode returns the PlanCode field if non-nil, zero value otherwise.

### GetPlanCodeOk

`func (o *NfsServerContract) GetPlanCodeOk() (*int32, bool)`

GetPlanCodeOk returns a tuple with the PlanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCode

`func (o *NfsServerContract) SetPlanCode(v int32)`

SetPlanCode sets PlanCode field to given value.


### GetPlanName

`func (o *NfsServerContract) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *NfsServerContract) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *NfsServerContract) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetServiceCode

`func (o *NfsServerContract) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *NfsServerContract) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *NfsServerContract) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


