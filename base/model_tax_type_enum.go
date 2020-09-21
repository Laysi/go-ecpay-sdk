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

// TaxTypeEnum **課稅類別** 若為應稅，請帶 `1`。   若為零稅率，請帶 `2`。   若為免稅，請帶 `3`。   若為混合應稅與免稅或零稅率時(限收銀機發票無法分辨時使用，且需通過申請核可)，則請帶 `9`。
type TaxTypeEnum string

// List of TaxTypeEnum
const (
	TAXTYPEENUM_TAXABLE       TaxTypeEnum = "1"
	TAXTYPEENUM_ZERO_TAX_RATE TaxTypeEnum = "2"
	TAXTYPEENUM_DUTY_FREE     TaxTypeEnum = "3"
	TAXTYPEENUM_MIXED         TaxTypeEnum = "9"
)

// Ptr returns reference to TaxTypeEnum value
func (v TaxTypeEnum) Ptr() *TaxTypeEnum {
	return &v
}

type NullableTaxTypeEnum struct {
	value *TaxTypeEnum
	isSet bool
}

func (v NullableTaxTypeEnum) Get() *TaxTypeEnum {
	return v.value
}

func (v *NullableTaxTypeEnum) Set(val *TaxTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxTypeEnum(val *TaxTypeEnum) *NullableTaxTypeEnum {
	return &NullableTaxTypeEnum{value: val, isSet: true}
}

func (v NullableTaxTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
