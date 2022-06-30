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

// AnAPIToken struct for AnAPIToken
type AnAPIToken struct {
	// The resource of the token.  The token applies to the specified resource and any resources below this so the resource `/` applies to the root resource and any resource below the root resource. In the same manner `/collections` will apply to all collectins while `/collections/{id}` will apply to that particular collection.
	Resource *string `json:"resource,omitempty"`
	// Write flag for token.  If this is set to `true` the token can be used for write operations in the API such as POST, DELETE and PATCH.
	Write *bool `json:"write,omitempty"`
	// Tags for the token.
	Tags *map[string]string `json:"tags,omitempty"`
	// The last time the token was used. Time in ms since epoch.
	LastUse *string `json:"lastUse,omitempty"`
	Uses    *string `json:"uses,omitempty"`
}

// NewAnAPIToken instantiates a new AnAPIToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnAPIToken() *AnAPIToken {
	this := AnAPIToken{}
	return &this
}

// NewAnAPITokenWithDefaults instantiates a new AnAPIToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnAPITokenWithDefaults() *AnAPIToken {
	this := AnAPIToken{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *AnAPIToken) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnAPIToken) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *AnAPIToken) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *AnAPIToken) SetResource(v string) {
	o.Resource = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *AnAPIToken) GetWrite() bool {
	if o == nil || o.Write == nil {
		var ret bool
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnAPIToken) GetWriteOk() (*bool, bool) {
	if o == nil || o.Write == nil {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *AnAPIToken) HasWrite() bool {
	if o != nil && o.Write != nil {
		return true
	}

	return false
}

// SetWrite gets a reference to the given bool and assigns it to the Write field.
func (o *AnAPIToken) SetWrite(v bool) {
	o.Write = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AnAPIToken) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnAPIToken) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AnAPIToken) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *AnAPIToken) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetLastUse returns the LastUse field value if set, zero value otherwise.
func (o *AnAPIToken) GetLastUse() string {
	if o == nil || o.LastUse == nil {
		var ret string
		return ret
	}
	return *o.LastUse
}

// GetLastUseOk returns a tuple with the LastUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnAPIToken) GetLastUseOk() (*string, bool) {
	if o == nil || o.LastUse == nil {
		return nil, false
	}
	return o.LastUse, true
}

// HasLastUse returns a boolean if a field has been set.
func (o *AnAPIToken) HasLastUse() bool {
	if o != nil && o.LastUse != nil {
		return true
	}

	return false
}

// SetLastUse gets a reference to the given string and assigns it to the LastUse field.
func (o *AnAPIToken) SetLastUse(v string) {
	o.LastUse = &v
}

// GetUses returns the Uses field value if set, zero value otherwise.
func (o *AnAPIToken) GetUses() string {
	if o == nil || o.Uses == nil {
		var ret string
		return ret
	}
	return *o.Uses
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnAPIToken) GetUsesOk() (*string, bool) {
	if o == nil || o.Uses == nil {
		return nil, false
	}
	return o.Uses, true
}

// HasUses returns a boolean if a field has been set.
func (o *AnAPIToken) HasUses() bool {
	if o != nil && o.Uses != nil {
		return true
	}

	return false
}

// SetUses gets a reference to the given string and assigns it to the Uses field.
func (o *AnAPIToken) SetUses(v string) {
	o.Uses = &v
}

func (o AnAPIToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Write != nil {
		toSerialize["write"] = o.Write
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.LastUse != nil {
		toSerialize["lastUse"] = o.LastUse
	}
	if o.Uses != nil {
		toSerialize["uses"] = o.Uses
	}
	return json.Marshal(toSerialize)
}

type NullableAnAPIToken struct {
	value *AnAPIToken
	isSet bool
}

func (v NullableAnAPIToken) Get() *AnAPIToken {
	return v.value
}

func (v *NullableAnAPIToken) Set(val *AnAPIToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAnAPIToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAnAPIToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnAPIToken(val *AnAPIToken) *NullableAnAPIToken {
	return &NullableAnAPIToken{value: val, isSet: true}
}

func (v NullableAnAPIToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnAPIToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
