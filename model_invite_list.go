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

// InviteList List of active invites. Onece an invite has been redeemed it will be removed automatically. An administrator of the team can delete unused invites.
type InviteList struct {
	Invites *[]Invite `json:"invites,omitempty"`
}

// NewInviteList instantiates a new InviteList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteList() *InviteList {
	this := InviteList{}
	return &this
}

// NewInviteListWithDefaults instantiates a new InviteList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteListWithDefaults() *InviteList {
	this := InviteList{}
	return &this
}

// GetInvites returns the Invites field value if set, zero value otherwise.
func (o *InviteList) GetInvites() []Invite {
	if o == nil || o.Invites == nil {
		var ret []Invite
		return ret
	}
	return *o.Invites
}

// GetInvitesOk returns a tuple with the Invites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteList) GetInvitesOk() (*[]Invite, bool) {
	if o == nil || o.Invites == nil {
		return nil, false
	}
	return o.Invites, true
}

// HasInvites returns a boolean if a field has been set.
func (o *InviteList) HasInvites() bool {
	if o != nil && o.Invites != nil {
		return true
	}

	return false
}

// SetInvites gets a reference to the given []Invite and assigns it to the Invites field.
func (o *InviteList) SetInvites(v []Invite) {
	o.Invites = &v
}

func (o InviteList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invites != nil {
		toSerialize["invites"] = o.Invites
	}
	return json.Marshal(toSerialize)
}

type NullableInviteList struct {
	value *InviteList
	isSet bool
}

func (v NullableInviteList) Get() *InviteList {
	return v.value
}

func (v *NullableInviteList) Set(val *InviteList) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteList) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteList(val *InviteList) *NullableInviteList {
	return &NullableInviteList{value: val, isSet: true}
}

func (v NullableInviteList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
