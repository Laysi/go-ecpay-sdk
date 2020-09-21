# PeriodReturnData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合  | 
**StoreID** | Pointer to **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
**RtnCode** | Pointer to **int** | **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。  | 
**RtnMsg** | Pointer to **string** | **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded    | 
**PeriodType** | Pointer to [**CreditPeriodTypeEnum**](CreditPeriodTypeEnum.md) |  | 
**Frequency** | Pointer to **int** | **執行頻率**  訂單建立時所設定的執行頻率  | 
**ExecTimes** | Pointer to **int** | **執行次數**  訂單建立時所設定的執行頻率  | 
**Amount** | Pointer to **int** | **本次授權金額**  此次所授權的金額  | 
**Gwsr** | Pointer to **int** | **授權交易單號**  此次所授權的交易單號  | 
**ProcessDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **處理時間** 處理時間 ( yyyy/MM/dd HH:mm:ss )  | 
**AuthCode** | Pointer to **string** | **授權碼** 授權碼  | 
**FirstAuthAmount** | Pointer to **int** | **初次授權金額** 定期定額交易的第一筆授權金額。  | 
**TotalSuccessTimes** | Pointer to **int** | **已執行成功次數**  目前已成功授權的次數。  | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 
**SimulatePaid** | Pointer to [**SimulatePaidEnum**](SimulatePaidEnum.md) |  | 
**CustomField1** | Pointer to **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField2** | Pointer to **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField3** | Pointer to **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField4** | Pointer to **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位  | 

## Methods

### NewPeriodReturnData

`func NewPeriodReturnData(merchantID string, merchantTradeNo string, storeID string, rtnCode int, rtnMsg string, periodType CreditPeriodTypeEnum, frequency int, execTimes int, amount int, gwsr int, processDate ECPayDateTime, authCode string, firstAuthAmount int, totalSuccessTimes int, checkMacValue string, simulatePaid SimulatePaidEnum, customField1 string, customField2 string, customField3 string, customField4 string, ) *PeriodReturnData`

NewPeriodReturnData instantiates a new PeriodReturnData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodReturnDataWithDefaults

`func NewPeriodReturnDataWithDefaults() *PeriodReturnData`

NewPeriodReturnDataWithDefaults instantiates a new PeriodReturnData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *PeriodReturnData) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *PeriodReturnData) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *PeriodReturnData) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *PeriodReturnData) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *PeriodReturnData) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *PeriodReturnData) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetStoreID

`func (o *PeriodReturnData) GetStoreID() string`

GetStoreID returns the StoreID field if non-nil, zero value otherwise.

### GetStoreIDOk

`func (o *PeriodReturnData) GetStoreIDOk() (*string, bool)`

GetStoreIDOk returns a tuple with the StoreID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreID

`func (o *PeriodReturnData) SetStoreID(v string)`

SetStoreID sets StoreID field to given value.


### GetRtnCode

`func (o *PeriodReturnData) GetRtnCode() int`

GetRtnCode returns the RtnCode field if non-nil, zero value otherwise.

### GetRtnCodeOk

`func (o *PeriodReturnData) GetRtnCodeOk() (*int, bool)`

GetRtnCodeOk returns a tuple with the RtnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnCode

`func (o *PeriodReturnData) SetRtnCode(v int)`

SetRtnCode sets RtnCode field to given value.


### GetRtnMsg

`func (o *PeriodReturnData) GetRtnMsg() string`

GetRtnMsg returns the RtnMsg field if non-nil, zero value otherwise.

### GetRtnMsgOk

`func (o *PeriodReturnData) GetRtnMsgOk() (*string, bool)`

GetRtnMsgOk returns a tuple with the RtnMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnMsg

`func (o *PeriodReturnData) SetRtnMsg(v string)`

SetRtnMsg sets RtnMsg field to given value.


### GetPeriodType

`func (o *PeriodReturnData) GetPeriodType() CreditPeriodTypeEnum`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *PeriodReturnData) GetPeriodTypeOk() (*CreditPeriodTypeEnum, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *PeriodReturnData) SetPeriodType(v CreditPeriodTypeEnum)`

SetPeriodType sets PeriodType field to given value.


### GetFrequency

`func (o *PeriodReturnData) GetFrequency() int`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PeriodReturnData) GetFrequencyOk() (*int, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PeriodReturnData) SetFrequency(v int)`

SetFrequency sets Frequency field to given value.


### GetExecTimes

`func (o *PeriodReturnData) GetExecTimes() int`

GetExecTimes returns the ExecTimes field if non-nil, zero value otherwise.

### GetExecTimesOk

`func (o *PeriodReturnData) GetExecTimesOk() (*int, bool)`

GetExecTimesOk returns a tuple with the ExecTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecTimes

`func (o *PeriodReturnData) SetExecTimes(v int)`

SetExecTimes sets ExecTimes field to given value.


### GetAmount

`func (o *PeriodReturnData) GetAmount() int`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PeriodReturnData) GetAmountOk() (*int, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PeriodReturnData) SetAmount(v int)`

SetAmount sets Amount field to given value.


### GetGwsr

`func (o *PeriodReturnData) GetGwsr() int`

GetGwsr returns the Gwsr field if non-nil, zero value otherwise.

### GetGwsrOk

`func (o *PeriodReturnData) GetGwsrOk() (*int, bool)`

GetGwsrOk returns a tuple with the Gwsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwsr

`func (o *PeriodReturnData) SetGwsr(v int)`

SetGwsr sets Gwsr field to given value.


### GetProcessDate

`func (o *PeriodReturnData) GetProcessDate() ECPayDateTime`

GetProcessDate returns the ProcessDate field if non-nil, zero value otherwise.

### GetProcessDateOk

`func (o *PeriodReturnData) GetProcessDateOk() (*ECPayDateTime, bool)`

GetProcessDateOk returns a tuple with the ProcessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDate

`func (o *PeriodReturnData) SetProcessDate(v ECPayDateTime)`

SetProcessDate sets ProcessDate field to given value.


### GetAuthCode

`func (o *PeriodReturnData) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PeriodReturnData) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PeriodReturnData) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetFirstAuthAmount

`func (o *PeriodReturnData) GetFirstAuthAmount() int`

GetFirstAuthAmount returns the FirstAuthAmount field if non-nil, zero value otherwise.

### GetFirstAuthAmountOk

`func (o *PeriodReturnData) GetFirstAuthAmountOk() (*int, bool)`

GetFirstAuthAmountOk returns a tuple with the FirstAuthAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAuthAmount

`func (o *PeriodReturnData) SetFirstAuthAmount(v int)`

SetFirstAuthAmount sets FirstAuthAmount field to given value.


### GetTotalSuccessTimes

`func (o *PeriodReturnData) GetTotalSuccessTimes() int`

GetTotalSuccessTimes returns the TotalSuccessTimes field if non-nil, zero value otherwise.

### GetTotalSuccessTimesOk

`func (o *PeriodReturnData) GetTotalSuccessTimesOk() (*int, bool)`

GetTotalSuccessTimesOk returns a tuple with the TotalSuccessTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSuccessTimes

`func (o *PeriodReturnData) SetTotalSuccessTimes(v int)`

SetTotalSuccessTimes sets TotalSuccessTimes field to given value.


### GetCheckMacValue

`func (o *PeriodReturnData) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *PeriodReturnData) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *PeriodReturnData) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.


### GetSimulatePaid

`func (o *PeriodReturnData) GetSimulatePaid() SimulatePaidEnum`

GetSimulatePaid returns the SimulatePaid field if non-nil, zero value otherwise.

### GetSimulatePaidOk

`func (o *PeriodReturnData) GetSimulatePaidOk() (*SimulatePaidEnum, bool)`

GetSimulatePaidOk returns a tuple with the SimulatePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatePaid

`func (o *PeriodReturnData) SetSimulatePaid(v SimulatePaidEnum)`

SetSimulatePaid sets SimulatePaid field to given value.


### GetCustomField1

`func (o *PeriodReturnData) GetCustomField1() string`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *PeriodReturnData) GetCustomField1Ok() (*string, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *PeriodReturnData) SetCustomField1(v string)`

SetCustomField1 sets CustomField1 field to given value.


### GetCustomField2

`func (o *PeriodReturnData) GetCustomField2() string`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *PeriodReturnData) GetCustomField2Ok() (*string, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *PeriodReturnData) SetCustomField2(v string)`

SetCustomField2 sets CustomField2 field to given value.


### GetCustomField3

`func (o *PeriodReturnData) GetCustomField3() string`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *PeriodReturnData) GetCustomField3Ok() (*string, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *PeriodReturnData) SetCustomField3(v string)`

SetCustomField3 sets CustomField3 field to given value.


### GetCustomField4

`func (o *PeriodReturnData) GetCustomField4() string`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *PeriodReturnData) GetCustomField4Ok() (*string, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *PeriodReturnData) SetCustomField4(v string)`

SetCustomField4 sets CustomField4 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


