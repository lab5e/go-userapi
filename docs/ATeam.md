# ATeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Members** | Pointer to [**[]Member**](Member.md) |  | [optional] 

## Methods

### NewATeam

`func NewATeam() *ATeam`

NewATeam instantiates a new ATeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewATeamWithDefaults

`func NewATeamWithDefaults() *ATeam`

NewATeamWithDefaults instantiates a new ATeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrivate

`func (o *ATeam) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ATeam) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ATeam) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ATeam) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetTags

`func (o *ATeam) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ATeam) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ATeam) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ATeam) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMembers

`func (o *ATeam) GetMembers() []Member`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ATeam) GetMembersOk() (*[]Member, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ATeam) SetMembers(v []Member)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ATeam) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


