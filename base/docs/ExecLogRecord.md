# ExecLogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RtnCode** | Pointer to **int** | **交易狀態**   回傳值為 1 時代表授權成功，其餘為失敗，失敗代碼請參考交易訊息代碼一覽  | [optional] 
**Amount** | Pointer to **int** | **本次授權金額**  所授權的金額  | [optional] 
**Gwsr** | Pointer to **int** | **授權交易單號**  所授權的交易單號  | [optional] 
**ProcessDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **處理時間** 格式為 yyyy/MM/dd HH:mm:ss  | [optional] 
**AuthCode** | Pointer to **string** | **授權碼** 授權碼  | [optional] 
**TradeNo** | Pointer to **string** | **綠界的交易編號** 請保存綠界的交易編號與特店交易編號 &#x60;MerchantTradeNo&#x60; 的關連。  | [optional] 

## Methods

### NewExecLogRecord

`func NewExecLogRecord() *ExecLogRecord`

NewExecLogRecord instantiates a new ExecLogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecLogRecordWithDefaults

`func NewExecLogRecordWithDefaults() *ExecLogRecord`

NewExecLogRecordWithDefaults instantiates a new ExecLogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRtnCode

`func (o *ExecLogRecord) GetRtnCode() int`

GetRtnCode returns the RtnCode field if non-nil, zero value otherwise.

### GetRtnCodeOk

`func (o *ExecLogRecord) GetRtnCodeOk() (*int, bool)`

GetRtnCodeOk returns a tuple with the RtnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnCode

`func (o *ExecLogRecord) SetRtnCode(v int)`

SetRtnCode sets RtnCode field to given value.

### HasRtnCode

`func (o *ExecLogRecord) HasRtnCode() bool`

HasRtnCode returns a boolean if a field has been set.

### GetAmount

`func (o *ExecLogRecord) GetAmount() int`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExecLogRecord) GetAmountOk() (*int, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExecLogRecord) SetAmount(v int)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExecLogRecord) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetGwsr

`func (o *ExecLogRecord) GetGwsr() int`

GetGwsr returns the Gwsr field if non-nil, zero value otherwise.

### GetGwsrOk

`func (o *ExecLogRecord) GetGwsrOk() (*int, bool)`

GetGwsrOk returns a tuple with the Gwsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwsr

`func (o *ExecLogRecord) SetGwsr(v int)`

SetGwsr sets Gwsr field to given value.

### HasGwsr

`func (o *ExecLogRecord) HasGwsr() bool`

HasGwsr returns a boolean if a field has been set.

### GetProcessDate

`func (o *ExecLogRecord) GetProcessDate() ECPayDateTime`

GetProcessDate returns the ProcessDate field if non-nil, zero value otherwise.

### GetProcessDateOk

`func (o *ExecLogRecord) GetProcessDateOk() (*ECPayDateTime, bool)`

GetProcessDateOk returns a tuple with the ProcessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDate

`func (o *ExecLogRecord) SetProcessDate(v ECPayDateTime)`

SetProcessDate sets ProcessDate field to given value.

### HasProcessDate

`func (o *ExecLogRecord) HasProcessDate() bool`

HasProcessDate returns a boolean if a field has been set.

### GetAuthCode

`func (o *ExecLogRecord) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ExecLogRecord) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ExecLogRecord) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ExecLogRecord) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetTradeNo

`func (o *ExecLogRecord) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ExecLogRecord) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ExecLogRecord) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *ExecLogRecord) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


