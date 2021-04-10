/*
 * The User API
 *
 * API to manage teams, members and tokens
 *
 * API version: 1.3.4 breezy-leafy
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package userapi

import (
	"encoding/json"
)

// InviteRequest struct for InviteRequest
type InviteRequest struct {
	TeamId *string `json:"teamId,omitempty"`
	// The invite code.
	Code *string `json:"code,omitempty"`
}

// NewInviteRequest instantiates a new InviteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteRequest() *InviteRequest {
	this := InviteRequest{}
	return &this
}

// NewInviteRequestWithDefaults instantiates a new InviteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteRequestWithDefaults() *InviteRequest {
	this := InviteRequest{}
	return &this
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *InviteRequest) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteRequest) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *InviteRequest) HasTeamId() bool {
	if o != nil && o.TeamId != nil {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *InviteRequest) SetTeamId(v string) {
	o.TeamId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InviteRequest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteRequest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InviteRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InviteRequest) SetCode(v string) {
	o.Code = &v
}

func (o InviteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TeamId != nil {
		toSerialize["teamId"] = o.TeamId
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableInviteRequest struct {
	value *InviteRequest
	isSet bool
}

func (v NullableInviteRequest) Get() *InviteRequest {
	return v.value
}

func (v *NullableInviteRequest) Set(val *InviteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteRequest(val *InviteRequest) *NullableInviteRequest {
	return &NullableInviteRequest{value: val, isSet: true}
}

func (v NullableInviteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}