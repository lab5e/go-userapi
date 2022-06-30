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

// ATeam struct for ATeam
type ATeam struct {
	IsPrivate *bool `json:"isPrivate,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Members *[]Member `json:"members,omitempty"`
}

// NewATeam instantiates a new ATeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewATeam() *ATeam {
	this := ATeam{}
	return &this
}

// NewATeamWithDefaults instantiates a new ATeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewATeamWithDefaults() *ATeam {
	this := ATeam{}
	return &this
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *ATeam) GetIsPrivate() bool {
	if o == nil || o.IsPrivate == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ATeam) GetIsPrivateOk() (*bool, bool) {
	if o == nil || o.IsPrivate == nil {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *ATeam) HasIsPrivate() bool {
	if o != nil && o.IsPrivate != nil {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *ATeam) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ATeam) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ATeam) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ATeam) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ATeam) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ATeam) GetMembers() []Member {
	if o == nil || o.Members == nil {
		var ret []Member
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ATeam) GetMembersOk() (*[]Member, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ATeam) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []Member and assigns it to the Members field.
func (o *ATeam) SetMembers(v []Member) {
	o.Members = &v
}

func (o ATeam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsPrivate != nil {
		toSerialize["isPrivate"] = o.IsPrivate
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableATeam struct {
	value *ATeam
	isSet bool
}

func (v NullableATeam) Get() *ATeam {
	return v.value
}

func (v *NullableATeam) Set(val *ATeam) {
	v.value = val
	v.isSet = true
}

func (v NullableATeam) IsSet() bool {
	return v.isSet
}

func (v *NullableATeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableATeam(val *ATeam) *NullableATeam {
	return &NullableATeam{value: val, isSet: true}
}

func (v NullableATeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableATeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

