# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UserProfile**](UserProfile.md) |  | [optional] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *Member) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Member) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Member) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *Member) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetRole

`func (o *Member) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Member) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Member) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Member) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserId

`func (o *Member) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Member) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Member) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Member) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *Member) GetUser() UserProfile`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Member) GetUserOk() (*UserProfile, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Member) SetUser(v UserProfile)`

SetUser sets User field to given value.

### HasUser

`func (o *Member) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


