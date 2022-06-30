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

// MemberOfATeam struct for MemberOfATeam
type MemberOfATeam struct {
	Role *string      `json:"role,omitempty"`
	User *UserProfile `json:"user,omitempty"`
}

// NewMemberOfATeam instantiates a new MemberOfATeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberOfATeam() *MemberOfATeam {
	this := MemberOfATeam{}
	return &this
}

// NewMemberOfATeamWithDefaults instantiates a new MemberOfATeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberOfATeamWithDefaults() *MemberOfATeam {
	this := MemberOfATeam{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MemberOfATeam) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberOfATeam) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MemberOfATeam) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MemberOfATeam) SetRole(v string) {
	o.Role = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MemberOfATeam) GetUser() UserProfile {
	if o == nil || o.User == nil {
		var ret UserProfile
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberOfATeam) GetUserOk() (*UserProfile, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MemberOfATeam) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserProfile and assigns it to the User field.
func (o *MemberOfATeam) SetUser(v UserProfile) {
	o.User = &v
}

func (o MemberOfATeam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableMemberOfATeam struct {
	value *MemberOfATeam
	isSet bool
}

func (v NullableMemberOfATeam) Get() *MemberOfATeam {
	return v.value
}

func (v *NullableMemberOfATeam) Set(val *MemberOfATeam) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberOfATeam) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberOfATeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberOfATeam(val *MemberOfATeam) *NullableMemberOfATeam {
	return &NullableMemberOfATeam{value: val, isSet: true}
}

func (v NullableMemberOfATeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberOfATeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
