/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.13 bordering-jerilyn
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// Member struct for Member
type Member struct {
	TeamId *string      `json:"teamId,omitempty"`
	Role   *string      `json:"role,omitempty"`
	UserId *string      `json:"userId,omitempty"`
	User   *UserProfile `json:"user,omitempty"`
}

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember() *Member {
	this := Member{}
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *Member) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *Member) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *Member) SetTeamId(v string) {
	o.TeamId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Member) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Member) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Member) SetRole(v string) {
	o.Role = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Member) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Member) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Member) SetUserId(v string) {
	o.UserId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Member) GetUser() UserProfile {
	if o == nil || o.User == nil {
		var ret UserProfile
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetUserOk() (*UserProfile, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Member) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserProfile and assigns it to the User field.
func (o *Member) SetUser(v UserProfile) {
	o.User = &v
}

func (o Member) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TeamId != nil {
		toSerialize["teamId"] = o.TeamId
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
