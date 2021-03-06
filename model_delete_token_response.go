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

// DeleteTokenResponse struct for DeleteTokenResponse
type DeleteTokenResponse struct {
	Token *Token `json:"token,omitempty"`
}

// NewDeleteTokenResponse instantiates a new DeleteTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTokenResponse() *DeleteTokenResponse {
	this := DeleteTokenResponse{}
	return &this
}

// NewDeleteTokenResponseWithDefaults instantiates a new DeleteTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTokenResponseWithDefaults() *DeleteTokenResponse {
	this := DeleteTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DeleteTokenResponse) GetToken() Token {
	if o == nil || o.Token == nil {
		var ret Token
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTokenResponse) GetTokenOk() (*Token, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DeleteTokenResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given Token and assigns it to the Token field.
func (o *DeleteTokenResponse) SetToken(v Token) {
	o.Token = &v
}

func (o DeleteTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteTokenResponse struct {
	value *DeleteTokenResponse
	isSet bool
}

func (v NullableDeleteTokenResponse) Get() *DeleteTokenResponse {
	return v.value
}

func (v *NullableDeleteTokenResponse) Set(val *DeleteTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTokenResponse(val *DeleteTokenResponse) *NullableDeleteTokenResponse {
	return &NullableDeleteTokenResponse{value: val, isSet: true}
}

func (v NullableDeleteTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
