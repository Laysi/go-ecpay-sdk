/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.21
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ecpayBase

import (
	"encoding/json"
)

// QueryTradeInfoRequest struct for QueryTradeInfoRequest
type QueryTradeInfoRequest struct {
	// **特店編號(由綠界提供)**
	MerchantID string `json:"MerchantID"`
	// **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。
	MerchantTradeNo string `json:"MerchantTradeNo"`
	// **驗證時間**   將當下的時間轉為UnixTimeStamp(見範例)用於驗證此次介接的時間區間。   綠界驗證時間區間暫訂為 3 分鐘內有效，超過則此次介接無效。
	TimeStamp int `json:"TimeStamp"`
	// **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 `MerchantID`。
	PlatformID *string `json:"PlatformID,omitempty"`
	// **檢查碼** 特店必須檢查檢查碼`CheckMacValue`來驗證，請參考附錄檢查碼機制。
	CheckMacValue string `json:"CheckMacValue"`
}

// NewQueryTradeInfoRequest instantiates a new QueryTradeInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryTradeInfoRequest(merchantID string, merchantTradeNo string, timeStamp int, checkMacValue string) *QueryTradeInfoRequest {
	this := QueryTradeInfoRequest{}
	this.MerchantID = merchantID
	this.MerchantTradeNo = merchantTradeNo
	this.TimeStamp = timeStamp
	this.CheckMacValue = checkMacValue
	return &this
}

// NewQueryTradeInfoRequestWithDefaults instantiates a new QueryTradeInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTradeInfoRequestWithDefaults() *QueryTradeInfoRequest {
	this := QueryTradeInfoRequest{}
	return &this
}

// GetMerchantID returns the MerchantID field value
func (o *QueryTradeInfoRequest) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *QueryTradeInfoRequest) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *QueryTradeInfoRequest) SetMerchantID(v string) {
	o.MerchantID = v
}

// GetMerchantTradeNo returns the MerchantTradeNo field value
func (o *QueryTradeInfoRequest) GetMerchantTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantTradeNo
}

// GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field value
// and a boolean to check if the value has been set.
func (o *QueryTradeInfoRequest) GetMerchantTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTradeNo, true
}

// SetMerchantTradeNo sets field value
func (o *QueryTradeInfoRequest) SetMerchantTradeNo(v string) {
	o.MerchantTradeNo = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *QueryTradeInfoRequest) GetTimeStamp() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *QueryTradeInfoRequest) GetTimeStampOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *QueryTradeInfoRequest) SetTimeStamp(v int) {
	o.TimeStamp = v
}

// GetPlatformID returns the PlatformID field value if set, zero value otherwise.
func (o *QueryTradeInfoRequest) GetPlatformID() string {
	if o == nil || o.PlatformID == nil {
		var ret string
		return ret
	}
	return *o.PlatformID
}

// GetPlatformIDOk returns a tuple with the PlatformID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTradeInfoRequest) GetPlatformIDOk() (*string, bool) {
	if o == nil || o.PlatformID == nil {
		return nil, false
	}
	return o.PlatformID, true
}

// HasPlatformID returns a boolean if a field has been set.
func (o *QueryTradeInfoRequest) HasPlatformID() bool {
	if o != nil && o.PlatformID != nil {
		return true
	}

	return false
}

// SetPlatformID gets a reference to the given string and assigns it to the PlatformID field.
func (o *QueryTradeInfoRequest) SetPlatformID(v string) {
	o.PlatformID = &v
}

// GetCheckMacValue returns the CheckMacValue field value
func (o *QueryTradeInfoRequest) GetCheckMacValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckMacValue
}

// GetCheckMacValueOk returns a tuple with the CheckMacValue field value
// and a boolean to check if the value has been set.
func (o *QueryTradeInfoRequest) GetCheckMacValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckMacValue, true
}

// SetCheckMacValue sets field value
func (o *QueryTradeInfoRequest) SetCheckMacValue(v string) {
	o.CheckMacValue = v
}

func (o QueryTradeInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MerchantID"] = o.MerchantID
	}
	if true {
		toSerialize["MerchantTradeNo"] = o.MerchantTradeNo
	}
	if true {
		toSerialize["TimeStamp"] = o.TimeStamp
	}
	if o.PlatformID != nil {
		toSerialize["PlatformID"] = o.PlatformID
	}
	if true {
		toSerialize["CheckMacValue"] = o.CheckMacValue
	}
	return json.Marshal(toSerialize)
}

type NullableQueryTradeInfoRequest struct {
	value *QueryTradeInfoRequest
	isSet bool
}

func (v NullableQueryTradeInfoRequest) Get() *QueryTradeInfoRequest {
	return v.value
}

func (v *NullableQueryTradeInfoRequest) Set(val *QueryTradeInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTradeInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTradeInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTradeInfoRequest(val *QueryTradeInfoRequest) *NullableQueryTradeInfoRequest {
	return &NullableQueryTradeInfoRequest{value: val, isSet: true}
}

func (v NullableQueryTradeInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryTradeInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
