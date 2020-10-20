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

// PeriodOrderResult struct for PeriodOrderResult
type PeriodOrderResult struct {
	// **特店編號**
	MerchantID string `json:"MerchantID"`
	// **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合
	MerchantTradeNo string `json:"MerchantTradeNo"`
	// **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。
	StoreID string `json:"StoreID"`
	// **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。
	RtnCode int `json:"RtnCode"`
	// **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded
	RtnMsg     string               `json:"RtnMsg"`
	PeriodType CreditPeriodTypeEnum `json:"PeriodType"`
	// **執行頻率**  訂單建立時所設定的執行頻率
	Frequency int `json:"Frequency"`
	// **執行次數**  訂單建立時所設定的執行頻率
	ExecTimes int `json:"ExecTimes"`
	// **本次授權金額**  此次所授權的金額
	Amount int `json:"Amount"`
	// **授權交易單號**  此次所授權的交易單號
	Gwsr int `json:"Gwsr"`
	// **處理時間** 處理時間 ( yyyy/MM/dd HH:mm:ss )
	ProcessDate ECPayDateTime `json:"ProcessDate"`
	// **授權碼** 授權碼
	AuthCode string `json:"AuthCode"`
	// **初次授權金額** 定期定額交易的第一筆授權金額。
	FirstAuthAmount int `json:"FirstAuthAmount"`
	// **已執行成功次數**  目前已成功授權的次數。
	TotalSuccessTimes int `json:"TotalSuccessTimes"`
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

// NewPeriodOrderResult instantiates a new PeriodOrderResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriodOrderResult(merchantID string, merchantTradeNo string, storeID string, rtnCode int, rtnMsg string, periodType CreditPeriodTypeEnum, frequency int, execTimes int, amount int, gwsr int, processDate ECPayDateTime, authCode string, firstAuthAmount int, totalSuccessTimes int, checkMacValue string, simulatePaid SimulatePaidEnum, customField1 string, customField2 string, customField3 string, customField4 string) *PeriodOrderResult {
	this := PeriodOrderResult{}
	this.MerchantID = merchantID
	this.MerchantTradeNo = merchantTradeNo
	this.StoreID = storeID
	this.RtnCode = rtnCode
	this.RtnMsg = rtnMsg
	this.PeriodType = periodType
	this.Frequency = frequency
	this.ExecTimes = execTimes
	this.Amount = amount
	this.Gwsr = gwsr
	this.ProcessDate = processDate
	this.AuthCode = authCode
	this.FirstAuthAmount = firstAuthAmount
	this.TotalSuccessTimes = totalSuccessTimes
	this.CheckMacValue = checkMacValue
	this.SimulatePaid = simulatePaid
	this.CustomField1 = customField1
	this.CustomField2 = customField2
	this.CustomField3 = customField3
	this.CustomField4 = customField4
	return &this
}

// NewPeriodOrderResultWithDefaults instantiates a new PeriodOrderResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodOrderResultWithDefaults() *PeriodOrderResult {
	this := PeriodOrderResult{}
	return &this
}

// GetMerchantID returns the MerchantID field value
func (o *PeriodOrderResult) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *PeriodOrderResult) SetMerchantID(v string) {
	o.MerchantID = v
}

// GetMerchantTradeNo returns the MerchantTradeNo field value
func (o *PeriodOrderResult) GetMerchantTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantTradeNo
}

// GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetMerchantTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTradeNo, true
}

// SetMerchantTradeNo sets field value
func (o *PeriodOrderResult) SetMerchantTradeNo(v string) {
	o.MerchantTradeNo = v
}

// GetStoreID returns the StoreID field value
func (o *PeriodOrderResult) GetStoreID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoreID
}

// GetStoreIDOk returns a tuple with the StoreID field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetStoreIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreID, true
}

// SetStoreID sets field value
func (o *PeriodOrderResult) SetStoreID(v string) {
	o.StoreID = v
}

// GetRtnCode returns the RtnCode field value
func (o *PeriodOrderResult) GetRtnCode() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.RtnCode
}

// GetRtnCodeOk returns a tuple with the RtnCode field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetRtnCodeOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RtnCode, true
}

// SetRtnCode sets field value
func (o *PeriodOrderResult) SetRtnCode(v int) {
	o.RtnCode = v
}

// GetRtnMsg returns the RtnMsg field value
func (o *PeriodOrderResult) GetRtnMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RtnMsg
}

// GetRtnMsgOk returns a tuple with the RtnMsg field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetRtnMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RtnMsg, true
}

// SetRtnMsg sets field value
func (o *PeriodOrderResult) SetRtnMsg(v string) {
	o.RtnMsg = v
}

// GetPeriodType returns the PeriodType field value
func (o *PeriodOrderResult) GetPeriodType() CreditPeriodTypeEnum {
	if o == nil {
		var ret CreditPeriodTypeEnum
		return ret
	}

	return o.PeriodType
}

// GetPeriodTypeOk returns a tuple with the PeriodType field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetPeriodTypeOk() (*CreditPeriodTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodType, true
}

// SetPeriodType sets field value
func (o *PeriodOrderResult) SetPeriodType(v CreditPeriodTypeEnum) {
	o.PeriodType = v
}

// GetFrequency returns the Frequency field value
func (o *PeriodOrderResult) GetFrequency() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetFrequencyOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *PeriodOrderResult) SetFrequency(v int) {
	o.Frequency = v
}

// GetExecTimes returns the ExecTimes field value
func (o *PeriodOrderResult) GetExecTimes() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.ExecTimes
}

// GetExecTimesOk returns a tuple with the ExecTimes field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetExecTimesOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecTimes, true
}

// SetExecTimes sets field value
func (o *PeriodOrderResult) SetExecTimes(v int) {
	o.ExecTimes = v
}

// GetAmount returns the Amount field value
func (o *PeriodOrderResult) GetAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PeriodOrderResult) SetAmount(v int) {
	o.Amount = v
}

// GetGwsr returns the Gwsr field value
func (o *PeriodOrderResult) GetGwsr() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Gwsr
}

// GetGwsrOk returns a tuple with the Gwsr field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetGwsrOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gwsr, true
}

// SetGwsr sets field value
func (o *PeriodOrderResult) SetGwsr(v int) {
	o.Gwsr = v
}

// GetProcessDate returns the ProcessDate field value
func (o *PeriodOrderResult) GetProcessDate() ECPayDateTime {
	if o == nil {
		var ret ECPayDateTime
		return ret
	}

	return o.ProcessDate
}

// GetProcessDateOk returns a tuple with the ProcessDate field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetProcessDateOk() (*ECPayDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessDate, true
}

// SetProcessDate sets field value
func (o *PeriodOrderResult) SetProcessDate(v ECPayDateTime) {
	o.ProcessDate = v
}

// GetAuthCode returns the AuthCode field value
func (o *PeriodOrderResult) GetAuthCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthCode, true
}

// SetAuthCode sets field value
func (o *PeriodOrderResult) SetAuthCode(v string) {
	o.AuthCode = v
}

// GetFirstAuthAmount returns the FirstAuthAmount field value
func (o *PeriodOrderResult) GetFirstAuthAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.FirstAuthAmount
}

// GetFirstAuthAmountOk returns a tuple with the FirstAuthAmount field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetFirstAuthAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstAuthAmount, true
}

// SetFirstAuthAmount sets field value
func (o *PeriodOrderResult) SetFirstAuthAmount(v int) {
	o.FirstAuthAmount = v
}

// GetTotalSuccessTimes returns the TotalSuccessTimes field value
func (o *PeriodOrderResult) GetTotalSuccessTimes() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TotalSuccessTimes
}

// GetTotalSuccessTimesOk returns a tuple with the TotalSuccessTimes field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetTotalSuccessTimesOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSuccessTimes, true
}

// SetTotalSuccessTimes sets field value
func (o *PeriodOrderResult) SetTotalSuccessTimes(v int) {
	o.TotalSuccessTimes = v
}

// GetCheckMacValue returns the CheckMacValue field value
func (o *PeriodOrderResult) GetCheckMacValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckMacValue
}

// GetCheckMacValueOk returns a tuple with the CheckMacValue field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetCheckMacValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckMacValue, true
}

// SetCheckMacValue sets field value
func (o *PeriodOrderResult) SetCheckMacValue(v string) {
	o.CheckMacValue = v
}

// GetSimulatePaid returns the SimulatePaid field value
func (o *PeriodOrderResult) GetSimulatePaid() SimulatePaidEnum {
	if o == nil {
		var ret SimulatePaidEnum
		return ret
	}

	return o.SimulatePaid
}

// GetSimulatePaidOk returns a tuple with the SimulatePaid field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetSimulatePaidOk() (*SimulatePaidEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SimulatePaid, true
}

// SetSimulatePaid sets field value
func (o *PeriodOrderResult) SetSimulatePaid(v SimulatePaidEnum) {
	o.SimulatePaid = v
}

// GetCustomField1 returns the CustomField1 field value
func (o *PeriodOrderResult) GetCustomField1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField1
}

// GetCustomField1Ok returns a tuple with the CustomField1 field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetCustomField1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField1, true
}

// SetCustomField1 sets field value
func (o *PeriodOrderResult) SetCustomField1(v string) {
	o.CustomField1 = v
}

// GetCustomField2 returns the CustomField2 field value
func (o *PeriodOrderResult) GetCustomField2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField2
}

// GetCustomField2Ok returns a tuple with the CustomField2 field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetCustomField2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField2, true
}

// SetCustomField2 sets field value
func (o *PeriodOrderResult) SetCustomField2(v string) {
	o.CustomField2 = v
}

// GetCustomField3 returns the CustomField3 field value
func (o *PeriodOrderResult) GetCustomField3() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField3
}

// GetCustomField3Ok returns a tuple with the CustomField3 field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetCustomField3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField3, true
}

// SetCustomField3 sets field value
func (o *PeriodOrderResult) SetCustomField3(v string) {
	o.CustomField3 = v
}

// GetCustomField4 returns the CustomField4 field value
func (o *PeriodOrderResult) GetCustomField4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomField4
}

// GetCustomField4Ok returns a tuple with the CustomField4 field value
// and a boolean to check if the value has been set.
func (o *PeriodOrderResult) GetCustomField4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField4, true
}

// SetCustomField4 sets field value
func (o *PeriodOrderResult) SetCustomField4(v string) {
	o.CustomField4 = v
}

func (o PeriodOrderResult) MarshalJSON() ([]byte, error) {
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
		toSerialize["PeriodType"] = o.PeriodType
	}
	if true {
		toSerialize["Frequency"] = o.Frequency
	}
	if true {
		toSerialize["ExecTimes"] = o.ExecTimes
	}
	if true {
		toSerialize["Amount"] = o.Amount
	}
	if true {
		toSerialize["Gwsr"] = o.Gwsr
	}
	if true {
		toSerialize["ProcessDate"] = o.ProcessDate
	}
	if true {
		toSerialize["AuthCode"] = o.AuthCode
	}
	if true {
		toSerialize["FirstAuthAmount"] = o.FirstAuthAmount
	}
	if true {
		toSerialize["TotalSuccessTimes"] = o.TotalSuccessTimes
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

type NullablePeriodOrderResult struct {
	value *PeriodOrderResult
	isSet bool
}

func (v NullablePeriodOrderResult) Get() *PeriodOrderResult {
	return v.value
}

func (v *NullablePeriodOrderResult) Set(val *PeriodOrderResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodOrderResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodOrderResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodOrderResult(val *PeriodOrderResult) *NullablePeriodOrderResult {
	return &NullablePeriodOrderResult{value: val, isSet: true}
}

func (v NullablePeriodOrderResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodOrderResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
