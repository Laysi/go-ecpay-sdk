/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// CreditPeriodTypeEnum **週期種類**   可設定以下參數：   `D`：以天為週期   `M`：以月為週期   `Y`：以年為週期
type CreditPeriodTypeEnum string

// List of CreditPeriodTypeEnum
const (
	CREDITPERIODTYPEENUM_DAY   CreditPeriodTypeEnum = "D"
	CREDITPERIODTYPEENUM_MONTH CreditPeriodTypeEnum = "M"
	CREDITPERIODTYPEENUM_YEAR  CreditPeriodTypeEnum = "Y"
)

// Ptr returns reference to CreditPeriodTypeEnum value
func (v CreditPeriodTypeEnum) Ptr() *CreditPeriodTypeEnum {
	return &v
}

type NullableCreditPeriodTypeEnum struct {
	value *CreditPeriodTypeEnum
	isSet bool
}

func (v NullableCreditPeriodTypeEnum) Get() *CreditPeriodTypeEnum {
	return v.value
}

func (v *NullableCreditPeriodTypeEnum) Set(val *CreditPeriodTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPeriodTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPeriodTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPeriodTypeEnum(val *CreditPeriodTypeEnum) *NullableCreditPeriodTypeEnum {
	return &NullableCreditPeriodTypeEnum{value: val, isSet: true}
}

func (v NullableCreditPeriodTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPeriodTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
