/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.23
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ecpayBase

import (
	"encoding/json"
)

// AioCheckOutCreditPeriodOption struct for AioCheckOutCreditPeriodOption
type AioCheckOutCreditPeriodOption struct {
	// **每次授權金額**   每次要授權(扣款)的金額。   注意事項：   綠界會依此次授權金額`PeriodAmount`所設定的金額做為之後固定授權的金額。   交易金額`TotalAmount`設定金額必須和授權金額`PeriodAmount`相同。   請帶整數，不可有小數點。僅限新台幣。
	PeriodAmount int                  `json:"PeriodAmount"`
	PeriodType   CreditPeriodTypeEnum `json:"PeriodType"`
	// **執行頻率**   此參數用來定義多久要執行一次   注意事項：   至少要大於等於 1 次以上。   當 `PeriodType` 設為 `D` 時，最多可設 `365` 次。    當 `PeriodType` 設為 `M` 時，最多可設 `12` 次。     當 `PeriodType` 設為 `Y` 時，最多可設 `1` 次。
	Frequency int `json:"Frequency"`
	// **執行次數**   總共要執行幾次。   注意事項：   至少要大於 1 次以上。   當 `PeriodType` 設為 `D` 時，最多可設 `999`次。   當 `PeriodType` 設為 `M` 時，最多可設 `99`次。   當 `PeriodType` 設為 `Y` 時，最多可設 `9` 次。    例 1：   當信用卡定期定額扣款為每個月扣 1 次500 元，總共要扣 12次   `TotalAmount`參數請帶 `500` `PeriodAmount`=`500`   `PeriodType`=`M`   `Frequency`=`1`   `ExecTimes`=`12`    例 2：   當信用卡定期定額扣款為 6000 元，每 6 個月扣 1 次，總共要扣 2 次時    交易金額`TotalAmount`參數請帶 `6000`   `PeriodType`=`M`   `Frequency`=`6`   `ExecTimes`=`2`
	ExecTimes int `json:"ExecTimes"`
	// **定期定額的執行結果回應URL**   若交易是信用卡定期定額的方式，則每次執行授權完，會將授權結果回傳到這個設定的 URL。   回覆內容請參考。
	PeriodReturnURL *string `json:"PeriodReturnURL,omitempty"`
}

// NewAioCheckOutCreditPeriodOption instantiates a new AioCheckOutCreditPeriodOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutCreditPeriodOption(periodAmount int, periodType CreditPeriodTypeEnum, frequency int, execTimes int) *AioCheckOutCreditPeriodOption {
	this := AioCheckOutCreditPeriodOption{}
	this.PeriodAmount = periodAmount
	this.PeriodType = periodType
	this.Frequency = frequency
	this.ExecTimes = execTimes
	return &this
}

// NewAioCheckOutCreditPeriodOptionWithDefaults instantiates a new AioCheckOutCreditPeriodOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutCreditPeriodOptionWithDefaults() *AioCheckOutCreditPeriodOption {
	this := AioCheckOutCreditPeriodOption{}
	return &this
}

// GetPeriodAmount returns the PeriodAmount field value
func (o *AioCheckOutCreditPeriodOption) GetPeriodAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.PeriodAmount
}

// GetPeriodAmountOk returns a tuple with the PeriodAmount field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditPeriodOption) GetPeriodAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodAmount, true
}

// SetPeriodAmount sets field value
func (o *AioCheckOutCreditPeriodOption) SetPeriodAmount(v int) {
	o.PeriodAmount = v
}

// GetPeriodType returns the PeriodType field value
func (o *AioCheckOutCreditPeriodOption) GetPeriodType() CreditPeriodTypeEnum {
	if o == nil {
		var ret CreditPeriodTypeEnum
		return ret
	}

	return o.PeriodType
}

// GetPeriodTypeOk returns a tuple with the PeriodType field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditPeriodOption) GetPeriodTypeOk() (*CreditPeriodTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodType, true
}

// SetPeriodType sets field value
func (o *AioCheckOutCreditPeriodOption) SetPeriodType(v CreditPeriodTypeEnum) {
	o.PeriodType = v
}

// GetFrequency returns the Frequency field value
func (o *AioCheckOutCreditPeriodOption) GetFrequency() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditPeriodOption) GetFrequencyOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *AioCheckOutCreditPeriodOption) SetFrequency(v int) {
	o.Frequency = v
}

// GetExecTimes returns the ExecTimes field value
func (o *AioCheckOutCreditPeriodOption) GetExecTimes() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.ExecTimes
}

// GetExecTimesOk returns a tuple with the ExecTimes field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditPeriodOption) GetExecTimesOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecTimes, true
}

// SetExecTimes sets field value
func (o *AioCheckOutCreditPeriodOption) SetExecTimes(v int) {
	o.ExecTimes = v
}

// GetPeriodReturnURL returns the PeriodReturnURL field value if set, zero value otherwise.
func (o *AioCheckOutCreditPeriodOption) GetPeriodReturnURL() string {
	if o == nil || o.PeriodReturnURL == nil {
		var ret string
		return ret
	}
	return *o.PeriodReturnURL
}

// GetPeriodReturnURLOk returns a tuple with the PeriodReturnURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCreditPeriodOption) GetPeriodReturnURLOk() (*string, bool) {
	if o == nil || o.PeriodReturnURL == nil {
		return nil, false
	}
	return o.PeriodReturnURL, true
}

// HasPeriodReturnURL returns a boolean if a field has been set.
func (o *AioCheckOutCreditPeriodOption) HasPeriodReturnURL() bool {
	if o != nil && o.PeriodReturnURL != nil {
		return true
	}

	return false
}

// SetPeriodReturnURL gets a reference to the given string and assigns it to the PeriodReturnURL field.
func (o *AioCheckOutCreditPeriodOption) SetPeriodReturnURL(v string) {
	o.PeriodReturnURL = &v
}

func (o AioCheckOutCreditPeriodOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["PeriodAmount"] = o.PeriodAmount
	}
	if true {
		toSerialize["PeriodType"] = o.PeriodType
	}
	if true {
		toSerialize["Frequency"] = o.Frequency
	}
	if true {
		toSerialize["ExecTimes"] = o.ExecTimes
	}
	if o.PeriodReturnURL != nil {
		toSerialize["PeriodReturnURL"] = o.PeriodReturnURL
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutCreditPeriodOption struct {
	value *AioCheckOutCreditPeriodOption
	isSet bool
}

func (v NullableAioCheckOutCreditPeriodOption) Get() *AioCheckOutCreditPeriodOption {
	return v.value
}

func (v *NullableAioCheckOutCreditPeriodOption) Set(val *AioCheckOutCreditPeriodOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutCreditPeriodOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutCreditPeriodOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutCreditPeriodOption(val *AioCheckOutCreditPeriodOption) *NullableAioCheckOutCreditPeriodOption {
	return &NullableAioCheckOutCreditPeriodOption{value: val, isSet: true}
}

func (v NullableAioCheckOutCreditPeriodOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutCreditPeriodOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
