# MemberOfATeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UserProfile**](UserProfile.md) |  | [optional] 

## Methods

### NewMemberOfATeam

`func NewMemberOfATeam() *MemberOfATeam`

NewMemberOfATeam instantiates a new MemberOfATeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberOfATeamWithDefaults

`func NewMemberOfATeamWithDefaults() *MemberOfATeam`

NewMemberOfATeamWithDefaults instantiates a new MemberOfATeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *MemberOfATeam) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberOfATeam) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberOfATeam) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberOfATeam) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *MemberOfATeam) GetUser() UserProfile`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MemberOfATeam) GetUserOk() (*UserProfile, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MemberOfATeam) SetUser(v UserProfile)`

SetUser sets User field to given value.

### HasUser

`func (o *MemberOfATeam) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


