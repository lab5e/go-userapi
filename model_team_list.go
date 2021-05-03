/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.10 constant-champ
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// TeamList struct for TeamList
type TeamList struct {
	Teams *[]Team `json:"teams,omitempty"`
}

// NewTeamList instantiates a new TeamList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamList() *TeamList {
	this := TeamList{}
	return &this
}

// NewTeamListWithDefaults instantiates a new TeamList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamListWithDefaults() *TeamList {
	this := TeamList{}
	return &this
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *TeamList) GetTeams() []Team {
	if o == nil || o.Teams == nil {
		var ret []Team
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamList) GetTeamsOk() (*[]Team, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *TeamList) HasTeams() bool {
	if o != nil && o.Teams != nil {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []Team and assigns it to the Teams field.
func (o *TeamList) SetTeams(v []Team) {
	o.Teams = &v
}

func (o TeamList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	return json.Marshal(toSerialize)
}

type NullableTeamList struct {
	value *TeamList
	isSet bool
}

func (v NullableTeamList) Get() *TeamList {
	return v.value
}

func (v *NullableTeamList) Set(val *TeamList) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamList) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamList(val *TeamList) *NullableTeamList {
	return &NullableTeamList{value: val, isSet: true}
}

func (v NullableTeamList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


