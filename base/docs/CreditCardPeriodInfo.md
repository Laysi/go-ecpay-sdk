# CreditCardPeriodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號(由綠界提供)**  | [optional] 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。  | [optional] 
**TradeNo** | Pointer to **string** | **綠界的交易編號** 首次授權所產生的綠界交易編號  | [optional] 
**RtnCode** | Pointer to **int** | **交易狀態**    回傳值為 1 時代表授權成功，其餘為失敗，失敗代碼請參考交易訊息代碼一覽表  | [optional] 
**PeriodType** | Pointer to [**CreditPeriodTypeEnum**](CreditPeriodTypeEnum.md) |  | [optional] 
**Frequency** | Pointer to **int** | **執行頻率**  訂單建立時所設定的執行頻率  | [optional] 
**ExecTimes** | Pointer to **int** | **執行次數**  訂單建立時所設定的執行頻率  | [optional] 
**PeriodAmount** | Pointer to **int** | **每次授權金額**  訂單建立時的每次要授權金額  | [optional] 
**Amount** | Pointer to **int** | **授權金額** 所授權的金額  | [optional] 
**Gwsr** | Pointer to **int** | **授權交易單號** 所授權的交易單號   | [optional] 
**ProcessDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **授權成功處理時間** 格式為 yyyy/MM/dd HH:mm:ss   | [optional] 
**AuthCode** | Pointer to **string** | **授權碼** 授權碼  | [optional] 
**Card4no** | Pointer to **string** | **卡片的末 4 碼** 卡片的末四碼  | [optional] 
**Card6no** | Pointer to **string** | **卡片的前 6 碼** 卡片的前六碼  | [optional] 
**TotalSuccessTimes** | Pointer to **int** | **已成功授權次數合計** 目前已成功授權的次數  | [optional] 
**TotalSuccessAmount** | Pointer to **int** | **已成功授權總金額** 目前已成功授權的金額合計   | [optional] 
**ExecStatus** | Pointer to [**ExecStatusEnum**](ExecStatusEnum.md) |  | [optional] 
**ExecLogRecord** | Pointer to [**[]ExecLogRecord**](ExecLogRecord.md) |  | [optional] 

## Methods

### NewCreditCardPeriodInfo

`func NewCreditCardPeriodInfo() *CreditCardPeriodInfo`

NewCreditCardPeriodInfo instantiates a new CreditCardPeriodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardPeriodInfoWithDefaults

`func NewCreditCardPeriodInfoWithDefaults() *CreditCardPeriodInfo`

NewCreditCardPeriodInfoWithDefaults instantiates a new CreditCardPeriodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *CreditCardPeriodInfo) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *CreditCardPeriodInfo) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *CreditCardPeriodInfo) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.

### HasMerchantID

`func (o *CreditCardPeriodInfo) HasMerchantID() bool`

HasMerchantID returns a boolean if a field has been set.

### GetMerchantTradeNo

`func (o *CreditCardPeriodInfo) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *CreditCardPeriodInfo) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *CreditCardPeriodInfo) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.

### HasMerchantTradeNo

`func (o *CreditCardPeriodInfo) HasMerchantTradeNo() bool`

HasMerchantTradeNo returns a boolean if a field has been set.

### GetTradeNo

`func (o *CreditCardPeriodInfo) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *CreditCardPeriodInfo) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *CreditCardPeriodInfo) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *CreditCardPeriodInfo) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### GetRtnCode

`func (o *CreditCardPeriodInfo) GetRtnCode() int`

GetRtnCode returns the RtnCode field if non-nil, zero value otherwise.

### GetRtnCodeOk

`func (o *CreditCardPeriodInfo) GetRtnCodeOk() (*int, bool)`

GetRtnCodeOk returns a tuple with the RtnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnCode

`func (o *CreditCardPeriodInfo) SetRtnCode(v int)`

SetRtnCode sets RtnCode field to given value.

### HasRtnCode

`func (o *CreditCardPeriodInfo) HasRtnCode() bool`

HasRtnCode returns a boolean if a field has been set.

### GetPeriodType

`func (o *CreditCardPeriodInfo) GetPeriodType() CreditPeriodTypeEnum`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *CreditCardPeriodInfo) GetPeriodTypeOk() (*CreditPeriodTypeEnum, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *CreditCardPeriodInfo) SetPeriodType(v CreditPeriodTypeEnum)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *CreditCardPeriodInfo) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### GetFrequency

`func (o *CreditCardPeriodInfo) GetFrequency() int`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CreditCardPeriodInfo) GetFrequencyOk() (*int, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CreditCardPeriodInfo) SetFrequency(v int)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CreditCardPeriodInfo) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetExecTimes

`func (o *CreditCardPeriodInfo) GetExecTimes() int`

GetExecTimes returns the ExecTimes field if non-nil, zero value otherwise.

### GetExecTimesOk

`func (o *CreditCardPeriodInfo) GetExecTimesOk() (*int, bool)`

GetExecTimesOk returns a tuple with the ExecTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecTimes

`func (o *CreditCardPeriodInfo) SetExecTimes(v int)`

SetExecTimes sets ExecTimes field to given value.

### HasExecTimes

`func (o *CreditCardPeriodInfo) HasExecTimes() bool`

HasExecTimes returns a boolean if a field has been set.

### GetPeriodAmount

`func (o *CreditCardPeriodInfo) GetPeriodAmount() int`

GetPeriodAmount returns the PeriodAmount field if non-nil, zero value otherwise.

### GetPeriodAmountOk

`func (o *CreditCardPeriodInfo) GetPeriodAmountOk() (*int, bool)`

GetPeriodAmountOk returns a tuple with the PeriodAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAmount

`func (o *CreditCardPeriodInfo) SetPeriodAmount(v int)`

SetPeriodAmount sets PeriodAmount field to given value.

### HasPeriodAmount

`func (o *CreditCardPeriodInfo) HasPeriodAmount() bool`

HasPeriodAmount returns a boolean if a field has been set.

### GetAmount

`func (o *CreditCardPeriodInfo) GetAmount() int`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditCardPeriodInfo) GetAmountOk() (*int, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditCardPeriodInfo) SetAmount(v int)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreditCardPeriodInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetGwsr

`func (o *CreditCardPeriodInfo) GetGwsr() int`

GetGwsr returns the Gwsr field if non-nil, zero value otherwise.

### GetGwsrOk

`func (o *CreditCardPeriodInfo) GetGwsrOk() (*int, bool)`

GetGwsrOk returns a tuple with the Gwsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwsr

`func (o *CreditCardPeriodInfo) SetGwsr(v int)`

SetGwsr sets Gwsr field to given value.

### HasGwsr

`func (o *CreditCardPeriodInfo) HasGwsr() bool`

HasGwsr returns a boolean if a field has been set.

### GetProcessDate

`func (o *CreditCardPeriodInfo) GetProcessDate() ECPayDateTime`

GetProcessDate returns the ProcessDate field if non-nil, zero value otherwise.

### GetProcessDateOk

`func (o *CreditCardPeriodInfo) GetProcessDateOk() (*ECPayDateTime, bool)`

GetProcessDateOk returns a tuple with the ProcessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDate

`func (o *CreditCardPeriodInfo) SetProcessDate(v ECPayDateTime)`

SetProcessDate sets ProcessDate field to given value.

### HasProcessDate

`func (o *CreditCardPeriodInfo) HasProcessDate() bool`

HasProcessDate returns a boolean if a field has been set.

### GetAuthCode

`func (o *CreditCardPeriodInfo) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *CreditCardPeriodInfo) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *CreditCardPeriodInfo) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *CreditCardPeriodInfo) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetCard4no

`func (o *CreditCardPeriodInfo) GetCard4no() string`

GetCard4no returns the Card4no field if non-nil, zero value otherwise.

### GetCard4noOk

`func (o *CreditCardPeriodInfo) GetCard4noOk() (*string, bool)`

GetCard4noOk returns a tuple with the Card4no field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard4no

`func (o *CreditCardPeriodInfo) SetCard4no(v string)`

SetCard4no sets Card4no field to given value.

### HasCard4no

`func (o *CreditCardPeriodInfo) HasCard4no() bool`

HasCard4no returns a boolean if a field has been set.

### GetCard6no

`func (o *CreditCardPeriodInfo) GetCard6no() string`

GetCard6no returns the Card6no field if non-nil, zero value otherwise.

### GetCard6noOk

`func (o *CreditCardPeriodInfo) GetCard6noOk() (*string, bool)`

GetCard6noOk returns a tuple with the Card6no field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard6no

`func (o *CreditCardPeriodInfo) SetCard6no(v string)`

SetCard6no sets Card6no field to given value.

### HasCard6no

`func (o *CreditCardPeriodInfo) HasCard6no() bool`

HasCard6no returns a boolean if a field has been set.

### GetTotalSuccessTimes

`func (o *CreditCardPeriodInfo) GetTotalSuccessTimes() int`

GetTotalSuccessTimes returns the TotalSuccessTimes field if non-nil, zero value otherwise.

### GetTotalSuccessTimesOk

`func (o *CreditCardPeriodInfo) GetTotalSuccessTimesOk() (*int, bool)`

GetTotalSuccessTimesOk returns a tuple with the TotalSuccessTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSuccessTimes

`func (o *CreditCardPeriodInfo) SetTotalSuccessTimes(v int)`

SetTotalSuccessTimes sets TotalSuccessTimes field to given value.

### HasTotalSuccessTimes

`func (o *CreditCardPeriodInfo) HasTotalSuccessTimes() bool`

HasTotalSuccessTimes returns a boolean if a field has been set.

### GetTotalSuccessAmount

`func (o *CreditCardPeriodInfo) GetTotalSuccessAmount() int`

GetTotalSuccessAmount returns the TotalSuccessAmount field if non-nil, zero value otherwise.

### GetTotalSuccessAmountOk

`func (o *CreditCardPeriodInfo) GetTotalSuccessAmountOk() (*int, bool)`

GetTotalSuccessAmountOk returns a tuple with the TotalSuccessAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSuccessAmount

`func (o *CreditCardPeriodInfo) SetTotalSuccessAmount(v int)`

SetTotalSuccessAmount sets TotalSuccessAmount field to given value.

### HasTotalSuccessAmount

`func (o *CreditCardPeriodInfo) HasTotalSuccessAmount() bool`

HasTotalSuccessAmount returns a boolean if a field has been set.

### GetExecStatus

`func (o *CreditCardPeriodInfo) GetExecStatus() ExecStatusEnum`

GetExecStatus returns the ExecStatus field if non-nil, zero value otherwise.

### GetExecStatusOk

`func (o *CreditCardPeriodInfo) GetExecStatusOk() (*ExecStatusEnum, bool)`

GetExecStatusOk returns a tuple with the ExecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecStatus

`func (o *CreditCardPeriodInfo) SetExecStatus(v ExecStatusEnum)`

SetExecStatus sets ExecStatus field to given value.

### HasExecStatus

`func (o *CreditCardPeriodInfo) HasExecStatus() bool`

HasExecStatus returns a boolean if a field has been set.

### GetExecLogRecord

`func (o *CreditCardPeriodInfo) GetExecLogRecord() []ExecLogRecord`

GetExecLogRecord returns the ExecLogRecord field if non-nil, zero value otherwise.

### GetExecLogRecordOk

`func (o *CreditCardPeriodInfo) GetExecLogRecordOk() (*[]ExecLogRecord, bool)`

GetExecLogRecordOk returns a tuple with the ExecLogRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecLogRecord

`func (o *CreditCardPeriodInfo) SetExecLogRecord(v []ExecLogRecord)`

SetExecLogRecord sets ExecLogRecord field to given value.

### HasExecLogRecord

`func (o *CreditCardPeriodInfo) HasExecLogRecord() bool`

HasExecLogRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


