/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.18
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package base

import (
	"encoding/json"
)

// ChoosePaymentEnum **選擇預設付款方式**   綠界提供下列付款方式，請於建立訂單時傳送過來:   - `Credit`: 信用卡及銀聯卡(需申請開通)   - `WebATM`: 網路 ATM   - `ATM`: 自動櫃員機   - `CVS`: 超商代碼    - `BARCODE`: 超商條碼   - `ALL`: 不指定付款方式，由綠界顯示付款方式選擇頁面。    注意事項： 1.若為手機版時不支援下列付款方式:   - WebATM:網路 ATM  2.如需要不透過綠界畫面取得 `ATM`、`CVS`、`BARCODE` 的繳費代碼，請參考 FAQ。
type ChoosePaymentEnum string

// List of ChoosePaymentEnum
const (
	CHOOSEPAYMENTENUM_CREDIT  ChoosePaymentEnum = "Credit"
	CHOOSEPAYMENTENUM_WEB_ATM ChoosePaymentEnum = "WebATM"
	CHOOSEPAYMENTENUM_ATM     ChoosePaymentEnum = "ATM"
	CHOOSEPAYMENTENUM_CVS     ChoosePaymentEnum = "CVS"
	CHOOSEPAYMENTENUM_BARCODE ChoosePaymentEnum = "BARCODE"
	CHOOSEPAYMENTENUM_ALL     ChoosePaymentEnum = "ALL"
)

// Ptr returns reference to ChoosePaymentEnum value
func (v ChoosePaymentEnum) Ptr() *ChoosePaymentEnum {
	return &v
}

type NullableChoosePaymentEnum struct {
	value *ChoosePaymentEnum
	isSet bool
}

func (v NullableChoosePaymentEnum) Get() *ChoosePaymentEnum {
	return v.value
}

func (v *NullableChoosePaymentEnum) Set(val *ChoosePaymentEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableChoosePaymentEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableChoosePaymentEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChoosePaymentEnum(val *ChoosePaymentEnum) *NullableChoosePaymentEnum {
	return &NullableChoosePaymentEnum{value: val, isSet: true}
}

func (v NullableChoosePaymentEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChoosePaymentEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
