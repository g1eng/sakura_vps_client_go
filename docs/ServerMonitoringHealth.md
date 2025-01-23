# ServerMonitoringHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCheckedAt** | **NullableTime** | 最終監視日時 | 
**Status** | **NullableString** | ステータス   * healthy 正常   * unhealthy 異常 | 
**LastStatusChangedAt** | **NullableTime** | 最終ステータス変更日時 | 

## Methods

### NewServerMonitoringHealth

`func NewServerMonitoringHealth(lastCheckedAt NullableTime, status NullableString, lastStatusChangedAt NullableTime, ) *ServerMonitoringHealth`

NewServerMonitoringHealth instantiates a new ServerMonitoringHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerMonitoringHealthWithDefaults

`func NewServerMonitoringHealthWithDefaults() *ServerMonitoringHealth`

NewServerMonitoringHealthWithDefaults instantiates a new ServerMonitoringHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCheckedAt

`func (o *ServerMonitoringHealth) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *ServerMonitoringHealth) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *ServerMonitoringHealth) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.


### SetLastCheckedAtNil

`func (o *ServerMonitoringHealth) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *ServerMonitoringHealth) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetStatus

`func (o *ServerMonitoringHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerMonitoringHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerMonitoringHealth) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ServerMonitoringHealth) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ServerMonitoringHealth) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastStatusChangedAt

`func (o *ServerMonitoringHealth) GetLastStatusChangedAt() time.Time`

GetLastStatusChangedAt returns the LastStatusChangedAt field if non-nil, zero value otherwise.

### GetLastStatusChangedAtOk

`func (o *ServerMonitoringHealth) GetLastStatusChangedAtOk() (*time.Time, bool)`

GetLastStatusChangedAtOk returns a tuple with the LastStatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusChangedAt

`func (o *ServerMonitoringHealth) SetLastStatusChangedAt(v time.Time)`

SetLastStatusChangedAt sets LastStatusChangedAt field to given value.


### SetLastStatusChangedAtNil

`func (o *ServerMonitoringHealth) SetLastStatusChangedAtNil(b bool)`

 SetLastStatusChangedAtNil sets the value for LastStatusChangedAt to be an explicit nil

### UnsetLastStatusChangedAt
`func (o *ServerMonitoringHealth) UnsetLastStatusChangedAt()`

UnsetLastStatusChangedAt ensures that no value is present for LastStatusChangedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


