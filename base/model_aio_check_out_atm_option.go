/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.22
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ecpayBase

import (
	"encoding/json"
)

// AioCheckOutAtmOption struct for AioCheckOutAtmOption
type AioCheckOutAtmOption struct {
	// **允許繳費有效天數**   若需設定最長 60 天，最短 1 天。   未設定此參數則預設為 3 天   注意事項：以天為單位
	ExpireDate *int `json:"ExpireDate,omitempty"`
	// **Server端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：銀行代碼、繳費虛擬帳號繳費期限…等)。   請參考[ATM、CVS 或 BARCODE 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。
	PaymentInfoURL *string `json:"PaymentInfoURL,omitempty"`
	// **Client端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Client 端回傳消費者付款方式相關資訊(例：銀行代碼、繳費虛擬帳號繳費期限…等)且將頁面轉到特店指定的頁面。請參考[ATM、CVS 或 BARCODE 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。
	ClientRedirectURL *string `json:"ClientRedirectURL,omitempty"`
}

// NewAioCheckOutAtmOption instantiates a new AioCheckOutAtmOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutAtmOption() *AioCheckOutAtmOption {
	this := AioCheckOutAtmOption{}
	var expireDate int = 3
	this.ExpireDate = &expireDate
	return &this
}

// NewAioCheckOutAtmOptionWithDefaults instantiates a new AioCheckOutAtmOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutAtmOptionWithDefaults() *AioCheckOutAtmOption {
	this := AioCheckOutAtmOption{}
	var expireDate int = 3
	this.ExpireDate = &expireDate
	return &this
}

// GetExpireDate returns the ExpireDate field value if set, zero value otherwise.
func (o *AioCheckOutAtmOption) GetExpireDate() int {
	if o == nil || o.ExpireDate == nil {
		var ret int
		return ret
	}
	return *o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutAtmOption) GetExpireDateOk() (*int, bool) {
	if o == nil || o.ExpireDate == nil {
		return nil, false
	}
	return o.ExpireDate, true
}

// HasExpireDate returns a boolean if a field has been set.
func (o *AioCheckOutAtmOption) HasExpireDate() bool {
	if o != nil && o.ExpireDate != nil {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given int and assigns it to the ExpireDate field.
func (o *AioCheckOutAtmOption) SetExpireDate(v int) {
	o.ExpireDate = &v
}

// GetPaymentInfoURL returns the PaymentInfoURL field value if set, zero value otherwise.
func (o *AioCheckOutAtmOption) GetPaymentInfoURL() string {
	if o == nil || o.PaymentInfoURL == nil {
		var ret string
		return ret
	}
	return *o.PaymentInfoURL
}

// GetPaymentInfoURLOk returns a tuple with the PaymentInfoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutAtmOption) GetPaymentInfoURLOk() (*string, bool) {
	if o == nil || o.PaymentInfoURL == nil {
		return nil, false
	}
	return o.PaymentInfoURL, true
}

// HasPaymentInfoURL returns a boolean if a field has been set.
func (o *AioCheckOutAtmOption) HasPaymentInfoURL() bool {
	if o != nil && o.PaymentInfoURL != nil {
		return true
	}

	return false
}

// SetPaymentInfoURL gets a reference to the given string and assigns it to the PaymentInfoURL field.
func (o *AioCheckOutAtmOption) SetPaymentInfoURL(v string) {
	o.PaymentInfoURL = &v
}

// GetClientRedirectURL returns the ClientRedirectURL field value if set, zero value otherwise.
func (o *AioCheckOutAtmOption) GetClientRedirectURL() string {
	if o == nil || o.ClientRedirectURL == nil {
		var ret string
		return ret
	}
	return *o.ClientRedirectURL
}

// GetClientRedirectURLOk returns a tuple with the ClientRedirectURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutAtmOption) GetClientRedirectURLOk() (*string, bool) {
	if o == nil || o.ClientRedirectURL == nil {
		return nil, false
	}
	return o.ClientRedirectURL, true
}

// HasClientRedirectURL returns a boolean if a field has been set.
func (o *AioCheckOutAtmOption) HasClientRedirectURL() bool {
	if o != nil && o.ClientRedirectURL != nil {
		return true
	}

	return false
}

// SetClientRedirectURL gets a reference to the given string and assigns it to the ClientRedirectURL field.
func (o *AioCheckOutAtmOption) SetClientRedirectURL(v string) {
	o.ClientRedirectURL = &v
}

func (o AioCheckOutAtmOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpireDate != nil {
		toSerialize["ExpireDate"] = o.ExpireDate
	}
	if o.PaymentInfoURL != nil {
		toSerialize["PaymentInfoURL"] = o.PaymentInfoURL
	}
	if o.ClientRedirectURL != nil {
		toSerialize["ClientRedirectURL"] = o.ClientRedirectURL
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutAtmOption struct {
	value *AioCheckOutAtmOption
	isSet bool
}

func (v NullableAioCheckOutAtmOption) Get() *AioCheckOutAtmOption {
	return v.value
}

func (v *NullableAioCheckOutAtmOption) Set(val *AioCheckOutAtmOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutAtmOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutAtmOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutAtmOption(val *AioCheckOutAtmOption) *NullableAioCheckOutAtmOption {
	return &NullableAioCheckOutAtmOption{value: val, isSet: true}
}

func (v NullableAioCheckOutAtmOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutAtmOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
