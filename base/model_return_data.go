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

// ReturnData struct for ReturnData
type ReturnData struct {
	// **特店編號**
	MerchantID string `json:"MerchantID"`
	// **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合
	MerchantTradeNo string `json:"MerchantTradeNo"`
	// **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。
	StoreID string `json:"StoreID"`
	// **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。
	RtnCode int `json:"RtnCode"`
	// **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded
	RtnMsg string `json:"RtnMsg"`
	// **綠界的交易編號** 請保存綠界的交易編號與特店交易編號[MerchantTradeNo]的關連。
	TradeNo string `json:"TradeNo"`
	// **交易金額**
	TradeAmt int `json:"TradeAmt"`
	// **付款時間** 格式為 yyyy/MM/dd HH:mm:ss
	PaymentDate ECPayDateTime         `json:"PaymentDate"`
	PaymentType ReturnPaymentTypeEnum `json:"PaymentType"`
	// **通路費**
	PaymentTypeChargeFee int `json:"PaymentTypeChargeFee"`
	// **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss
	TradeDate ECPayDateTime `json:"TradeDate"`
	// **檢查碼** 特店必須檢查檢查碼`CheckMacValue`來驗證，請參考附錄檢查碼機制。
	CheckMacValue string           `json:"CheckMacValue"`
	SimulatePaid  SimulatePaidEnum `json:"SimulatePaid"`
	// **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位
	CustomField1 string `json:"CustomField1"`
	// **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位
	CustomField2 string `json:"CustomField2"`
	// **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位
	CustomField3 string `json:"CustomField3"`
	// **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位
	CustomField4 string `json:"CustomField4"`
}

// NewReturnData instantiates a new ReturnData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnData(merchantID string, merchantTradeNo string, storeID string, rtnCode int, rtnMsg string, tradeNo string, tradeAmt int, paymentDate ECPayDateTime, paymentType ReturnPaymentTypeEnum, paymentTypeChargeFee int, tradeDate ECPayDateTime, checkMacValue string, simulatePaid SimulatePaidEnum, customField1 string, customField2 string, customField3 string, customField4 string) *ReturnData {
	this := ReturnData{}
	this.MerchantID = merchantID
	this.MerchantTradeNo = merchantTradeNo
	this.StoreID = storeID
	this.RtnCode = rtnCode
	this.RtnMsg = rtnMsg
	this.TradeNo = tradeNo
	this.TradeAmt = tradeAmt
	this.PaymentDate = paymentDate
	this.PaymentType = paymentType
	this.PaymentTypeChargeFee = paymentTypeChargeFee
	this.TradeDate = tradeDate
	this.CheckMacValue = checkMacValue
	this.SimulatePaid = simulatePaid
	this.CustomField1 = customField1
	this.CustomField2 = customField2
	this.CustomField3 = customField3
	this.CustomField4 = customField4
	return &this
}

// NewReturnDataWithDefaults instantiates a new ReturnData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnDataWithDefaults() *ReturnData {
	this := ReturnData{}
	return &this
}

// GetMerchantID returns the MerchantID field value
func (o *ReturnData) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *ReturnData) SetMerchantID(v string) {
	o.MerchantID = v
}

// GetMerchantTradeNo returns the MerchantTradeNo field value
func (o *ReturnData) GetMerchantTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantTradeNo
}

// GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetMerchantTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTradeNo, true
}

// SetMerchantTradeNo sets field value
func (o *ReturnData) SetMerchantTradeNo(v string) {
	o.MerchantTradeNo = v
}

// GetStoreID returns the StoreID field value
func (o *ReturnData) GetStoreID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoreID
}

// GetStoreIDOk returns a tuple with the StoreID field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetStoreIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreID, true
}

// SetStoreID sets field value
func (o *ReturnData) SetStoreID(v string) {
	o.StoreID = v
}

// GetRtnCode returns the RtnCode field value
func (o *ReturnData) GetRtnCode() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.RtnCode
}

// GetRtnCodeOk returns a tuple with the RtnCode field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetRtnCodeOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RtnCode, true
}

// SetRtnCode sets field value
func (o *ReturnData) SetRtnCode(v int) {
	o.RtnCode = v
}

// GetRtnMsg returns the RtnMsg field value
func (o *ReturnData) GetRtnMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RtnMsg
}

// GetRtnMsgOk returns a tuple with the RtnMsg field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetRtnMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RtnMsg, true
}

// SetRtnMsg sets field value
func (o *ReturnData) SetRtnMsg(v string) {
	o.RtnMsg = v
}

// GetTradeNo returns the TradeNo field value
func (o *ReturnData) GetTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeNo, true
}

// SetTradeNo sets field value
func (o *ReturnData) SetTradeNo(v string) {
	o.TradeNo = v
}

// GetTradeAmt returns the TradeAmt field value
func (o *ReturnData) GetTradeAmt() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TradeAmt
}

// GetTradeAmtOk returns a tuple with the TradeAmt field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetTradeAmtOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeAmt, true
}

// SetTradeAmt sets field value
func (o *ReturnData) SetTradeAmt(v int) {
	o.TradeAmt = v
}

// GetPaymentDate returns the PaymentDate field value
func (o *ReturnData) GetPaymentDate() ECPayDateTime {
	if o == nil {
		var ret ECPayDateTime
		return ret
	}

	return o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetPaymentDateOk() (*ECPayDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentDate, true
}

// SetPaymentDate sets field value
func (o *ReturnData) SetPaymentDate(v ECPayDateTime) {
	o.PaymentDate = v
}

// GetPaymentType returns the PaymentType field value
func (o *ReturnData) GetPaymentType() ReturnPaymentTypeEnum {
	if o == nil {
		var ret ReturnPaymentTypeEnum
		return ret
	}

	return o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetPaymentTypeOk() (*ReturnPaymentTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentType, true
}

// SetPaymentType sets field value
func (o *ReturnData) SetPaymentType(v ReturnPaymentTypeEnum) {
	o.PaymentType = v
}

// GetPaymentTypeChargeFee returns the PaymentTypeChargeFee field value
func (o *ReturnData) GetPaymentTypeChargeFee() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.PaymentTypeChargeFee
}

// GetPaymentTypeChargeFeeOk returns a tuple with the PaymentTypeChargeFee field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetPaymentTypeChargeFeeOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentTypeChargeFee, true
}

// SetPaymentTypeChargeFee sets field value
func (o *ReturnData) SetPaymentTypeChargeFee(v int) {
	o.PaymentTypeChargeFee = v
}

// GetTradeDate returns the TradeDate field value
func (o *ReturnData) GetTradeDate() ECPayDateTime {
	if o == nil {
		var ret ECPayDateTime
		return ret
	}

	return o.TradeDate
}

// GetTradeDateOk returns a tuple with the TradeDate field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetTradeDateOk() (*ECPayDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeDate, true
}

// SetTradeDate sets field value
func (o *ReturnData) SetTradeDate(v ECPayDateTime) {
	o.TradeDate = v
}

// GetCheckMacValue returns the CheckMacValue field value
func (o *ReturnData) GetCheckMacValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckMacValue
}

// GetCheckMacValueOk returns a tuple with the CheckMacValue field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetCheckMacValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckMacValue, true
}

// SetCheckMacValue sets field value
func (o *ReturnData) SetCheckMacValue(v string) {
	o.CheckMacValue = v
}

// GetSimulatePaid returns the SimulatePaid field value
func (o *ReturnData) GetSimulatePaid() SimulatePaidEnum {
	if o == nil {
		var ret SimulatePaidEnum
		return ret
	}

	return o.SimulatePaid
}

// GetSimulatePaidOk returns a tuple with the SimulatePaid field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetSimulatePaidOk() (*SimulatePaidEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SimulatePaid, true
}

// SetSimulatePaid sets field value
func (o *ReturnData) SetSimulatePaid(v SimulatePaidEnum) {
	o.SimulatePaid = v
}

// GetCustomField1 returns the CustomField1 field value
func (o *ReturnData) GetCustomField1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField1
}

// GetCustomField1Ok returns a tuple with the CustomField1 field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetCustomField1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField1, true
}

// SetCustomField1 sets field value
func (o *ReturnData) SetCustomField1(v string) {
	o.CustomField1 = v
}

// GetCustomField2 returns the CustomField2 field value
func (o *ReturnData) GetCustomField2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField2
}

// GetCustomField2Ok returns a tuple with the CustomField2 field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetCustomField2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField2, true
}

// SetCustomField2 sets field value
func (o *ReturnData) SetCustomField2(v string) {
	o.CustomField2 = v
}

// GetCustomField3 returns the CustomField3 field value
func (o *ReturnData) GetCustomField3() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField3
}

// GetCustomField3Ok returns a tuple with the CustomField3 field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetCustomField3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField3, true
}

// SetCustomField3 sets field value
func (o *ReturnData) SetCustomField3(v string) {
	o.CustomField3 = v
}

// GetCustomField4 returns the CustomField4 field value
func (o *ReturnData) GetCustomField4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField4
}

// GetCustomField4Ok returns a tuple with the CustomField4 field value
// and a boolean to check if the value has been set.
func (o *ReturnData) GetCustomField4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField4, true
}

// SetCustomField4 sets field value
func (o *ReturnData) SetCustomField4(v string) {
	o.CustomField4 = v
}

func (o ReturnData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MerchantID"] = o.MerchantID
	}
	if true {
		toSerialize["MerchantTradeNo"] = o.MerchantTradeNo
	}
	if true {
		toSerialize["StoreID"] = o.StoreID
	}
	if true {
		toSerialize["RtnCode"] = o.RtnCode
	}
	if true {
		toSerialize["RtnMsg"] = o.RtnMsg
	}
	if true {
		toSerialize["TradeNo"] = o.TradeNo
	}
	if true {
		toSerialize["TradeAmt"] = o.TradeAmt
	}
	if true {
		toSerialize["PaymentDate"] = o.PaymentDate
	}
	if true {
		toSerialize["PaymentType"] = o.PaymentType
	}
	if true {
		toSerialize["PaymentTypeChargeFee"] = o.PaymentTypeChargeFee
	}
	if true {
		toSerialize["TradeDate"] = o.TradeDate
	}
	if true {
		toSerialize["CheckMacValue"] = o.CheckMacValue
	}
	if true {
		toSerialize["SimulatePaid"] = o.SimulatePaid
	}
	if true {
		toSerialize["CustomField1"] = o.CustomField1
	}
	if true {
		toSerialize["CustomField2"] = o.CustomField2
	}
	if true {
		toSerialize["CustomField3"] = o.CustomField3
	}
	if true {
		toSerialize["CustomField4"] = o.CustomField4
	}
	return json.Marshal(toSerialize)
}

type NullableReturnData struct {
	value *ReturnData
	isSet bool
}

func (v NullableReturnData) Get() *ReturnData {
	return v.value
}

func (v *NullableReturnData) Set(val *ReturnData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnData(val *ReturnData) *NullableReturnData {
	return &NullableReturnData{value: val, isSet: true}
}

func (v NullableReturnData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
