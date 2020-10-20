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

// InvoicePrintEnum **列印註記**   若為不列印或捐贈註記`Donation`為1(捐贈)時，請帶：`0`。   若為列印或統一編號`CustomerIdentifier`有值時，請帶：`1`
type InvoicePrintEnum string

// List of InvoicePrintEnum
const (
	INVOICEPRINTENUM_PRINT     InvoicePrintEnum = "0"
	INVOICEPRINTENUM_NOT_PRINT InvoicePrintEnum = "1"
)

// Ptr returns reference to InvoicePrintEnum value
func (v InvoicePrintEnum) Ptr() *InvoicePrintEnum {
	return &v
}

type NullableInvoicePrintEnum struct {
	value *InvoicePrintEnum
	isSet bool
}

func (v NullableInvoicePrintEnum) Get() *InvoicePrintEnum {
	return v.value
}

func (v *NullableInvoicePrintEnum) Set(val *InvoicePrintEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicePrintEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicePrintEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicePrintEnum(val *InvoicePrintEnum) *NullableInvoicePrintEnum {
	return &NullableInvoicePrintEnum{value: val, isSet: true}
}

func (v NullableInvoicePrintEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicePrintEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
