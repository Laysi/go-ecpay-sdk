/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.16
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// AioCheckOutCreditOption struct for AioCheckOutCreditOption
type AioCheckOutCreditOption struct {
	BindingCard *BindingCardEnum `json:"BindingCard,omitempty"`
	// **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 `MerchantID` + `廠商會員編號`)
	MerchantMemberID *string `json:"MerchantMemberID,omitempty"`
}

// NewAioCheckOutCreditOption instantiates a new AioCheckOutCreditOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutCreditOption() *AioCheckOutCreditOption {
	this := AioCheckOutCreditOption{}
	return &this
}

// NewAioCheckOutCreditOptionWithDefaults instantiates a new AioCheckOutCreditOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutCreditOptionWithDefaults() *AioCheckOutCreditOption {
	this := AioCheckOutCreditOption{}
	return &this
}

// GetBindingCard returns the BindingCard field value if set, zero value otherwise.
func (o *AioCheckOutCreditOption) GetBindingCard() BindingCardEnum {
	if o == nil || o.BindingCard == nil {
		var ret BindingCardEnum
		return ret
	}
	return *o.BindingCard
}

// GetBindingCardOk returns a tuple with the BindingCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditOption) GetBindingCardOk() (*BindingCardEnum, bool) {
	if o == nil || o.BindingCard == nil {
		return nil, false
	}
	return o.BindingCard, true
}

// HasBindingCard returns a boolean if a field has been set.
func (o *AioCheckOutCreditOption) HasBindingCard() bool {
	if o != nil && o.BindingCard != nil {
		return true
	}

	return false
}

// SetBindingCard gets a reference to the given BindingCardEnum and assigns it to the BindingCard field.
func (o *AioCheckOutCreditOption) SetBindingCard(v BindingCardEnum) {
	o.BindingCard = &v
}

// GetMerchantMemberID returns the MerchantMemberID field value if set, zero value otherwise.
func (o *AioCheckOutCreditOption) GetMerchantMemberID() string {
	if o == nil || o.MerchantMemberID == nil {
		var ret string
		return ret
	}
	return *o.MerchantMemberID
}

// GetMerchantMemberIDOk returns a tuple with the MerchantMemberID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditOption) GetMerchantMemberIDOk() (*string, bool) {
	if o == nil || o.MerchantMemberID == nil {
		return nil, false
	}
	return o.MerchantMemberID, true
}

// HasMerchantMemberID returns a boolean if a field has been set.
func (o *AioCheckOutCreditOption) HasMerchantMemberID() bool {
	if o != nil && o.MerchantMemberID != nil {
		return true
	}

	return false
}

// SetMerchantMemberID gets a reference to the given string and assigns it to the MerchantMemberID field.
func (o *AioCheckOutCreditOption) SetMerchantMemberID(v string) {
	o.MerchantMemberID = &v
}

func (o AioCheckOutCreditOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BindingCard != nil {
		toSerialize["BindingCard"] = o.BindingCard
	}
	if o.MerchantMemberID != nil {
		toSerialize["MerchantMemberID"] = o.MerchantMemberID
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutCreditOption struct {
	value *AioCheckOutCreditOption
	isSet bool
}

func (v NullableAioCheckOutCreditOption) Get() *AioCheckOutCreditOption {
	return v.value
}

func (v *NullableAioCheckOutCreditOption) Set(val *AioCheckOutCreditOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutCreditOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutCreditOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutCreditOption(val *AioCheckOutCreditOption) *NullableAioCheckOutCreditOption {
	return &NullableAioCheckOutCreditOption{value: val, isSet: true}
}

func (v NullableAioCheckOutCreditOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutCreditOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
