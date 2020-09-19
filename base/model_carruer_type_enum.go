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

// CarruerTypeEnum **載具類別**   若為無載具時，則請帶空字串。   若為特店載具時，則請帶 `1`。   若為買受人之自然人憑證號碼時，則請帶 `2`。   若為買受人之手機條碼資料時，則請帶`3`。   若統一編號 `CustomerIdentifier` 有值時，則載具類別不可為特店載具或自然人憑證載具。   注意事項：當`Print`有值時，載具類別不得有值。
type CarruerTypeEnum string

// List of CarruerTypeEnum
const (
	CARRUERTYPEENUM_NO_CARRIER                         CarruerTypeEnum = ""
	CARRUERTYPEENUM_SHOP_CARRIER                       CarruerTypeEnum = "1"
	CARRUERTYPEENUM_NATURAL_PERSON_CERTIFICATE_CARRIER CarruerTypeEnum = "2"
	CARRUERTYPEENUM_PHONE_BARCODE_CARRIER              CarruerTypeEnum = "3"
)

// Ptr returns reference to CarruerTypeEnum value
func (v CarruerTypeEnum) Ptr() *CarruerTypeEnum {
	return &v
}

type NullableCarruerTypeEnum struct {
	value *CarruerTypeEnum
	isSet bool
}

func (v NullableCarruerTypeEnum) Get() *CarruerTypeEnum {
	return v.value
}

func (v *NullableCarruerTypeEnum) Set(val *CarruerTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCarruerTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCarruerTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarruerTypeEnum(val *CarruerTypeEnum) *NullableCarruerTypeEnum {
	return &NullableCarruerTypeEnum{value: val, isSet: true}
}

func (v NullableCarruerTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarruerTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}