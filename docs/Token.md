# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** | The resource of the token.  The token applies to the specified resource and any resources below this so the resource &#x60;/&#x60; applies to the root resource and any resource below the root resource. In the same manner &#x60;/collections&#x60; will apply to all collectins while &#x60;/collections/{id}&#x60; will apply to that particular collection. | [optional] 
**Write** | Pointer to **bool** | Write flag for token.  If this is set to &#x60;true&#x60; the token can be used for write operations in the API such as POST, DELETE and PATCH. | [optional] 
**Token** | Pointer to **string** | Use this in the &#x60;X-API-Token&#x60; header when using the API. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags for the token. | [optional] 
**LastUse** | Pointer to **string** | The last time the token was used. Time in ms since epoch. | [optional] 
**Uses** | Pointer to **string** |  | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *Token) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Token) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Token) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Token) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetWrite

`func (o *Token) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *Token) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *Token) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *Token) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetToken

`func (o *Token) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Token) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Token) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Token) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTags

`func (o *Token) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Token) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Token) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Token) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastUse

`func (o *Token) GetLastUse() string`

GetLastUse returns the LastUse field if non-nil, zero value otherwise.

### GetLastUseOk

`func (o *Token) GetLastUseOk() (*string, bool)`

GetLastUseOk returns a tuple with the LastUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUse

`func (o *Token) SetLastUse(v string)`

SetLastUse sets LastUse field to given value.

### HasLastUse

`func (o *Token) HasLastUse() bool`

HasLastUse returns a boolean if a field has been set.

### GetUses

`func (o *Token) GetUses() string`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *Token) GetUsesOk() (*string, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *Token) SetUses(v string)`

SetUses sets Uses field to given value.

### HasUses

`func (o *Token) HasUses() bool`

HasUses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


