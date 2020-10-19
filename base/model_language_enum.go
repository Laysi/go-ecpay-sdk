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

// LanguageEnum **語系設定**   預設語系為中文，若要變更語系參數值請帶：   - 英語：`ENG`   - 韓語：`KOR`   - 日語：`JPN`   - 簡體中文：`CHI`
type LanguageEnum string

// List of LanguageEnum
const (
	LANGUAGEENUM_ENG LanguageEnum = "ENG"
	LANGUAGEENUM_KOR LanguageEnum = "KOR"
	LANGUAGEENUM_JPN LanguageEnum = "JPN"
	LANGUAGEENUM_CHI LanguageEnum = "CHI"
)

// Ptr returns reference to LanguageEnum value
func (v LanguageEnum) Ptr() *LanguageEnum {
	return &v
}

type NullableLanguageEnum struct {
	value *LanguageEnum
	isSet bool
}

func (v NullableLanguageEnum) Get() *LanguageEnum {
	return v.value
}

func (v *NullableLanguageEnum) Set(val *LanguageEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageEnum(val *LanguageEnum) *NullableLanguageEnum {
	return &NullableLanguageEnum{value: val, isSet: true}
}

func (v NullableLanguageEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
