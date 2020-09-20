/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.10
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// SimulatePaidEnum **是否為模擬付款**   回傳值：   若為 1 時，代表此交易為模擬付款，請勿出貨。   若為 0 時，代表此交易非模擬付款。     注意事項：   特店可透過廠商後台網站來針對單筆訂單模擬綠界回傳付款通知，以方便介接 API 的測試。
type SimulatePaidEnum int32

// List of SimulatePaidEnum
const (
	SIMULATEPAIDENUM_REAL_PAID SimulatePaidEnum = 0
	SIMULATEPAIDENUM_SIMULATED SimulatePaidEnum = 1
)

// Ptr returns reference to SimulatePaidEnum value
func (v SimulatePaidEnum) Ptr() *SimulatePaidEnum {
	return &v
}

type NullableSimulatePaidEnum struct {
	value *SimulatePaidEnum
	isSet bool
}

func (v NullableSimulatePaidEnum) Get() *SimulatePaidEnum {
	return v.value
}

func (v *NullableSimulatePaidEnum) Set(val *SimulatePaidEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulatePaidEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulatePaidEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulatePaidEnum(val *SimulatePaidEnum) *NullableSimulatePaidEnum {
	return &NullableSimulatePaidEnum{value: val, isSet: true}
}

func (v NullableSimulatePaidEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulatePaidEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
