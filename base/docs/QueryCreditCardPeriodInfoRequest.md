# QueryCreditCardPeriodInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號(由綠界提供)**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。  | 
**TimeStamp** | Pointer to **int** | **驗證時間**   將當下的時間轉為UnixTimeStamp(見範例)用於驗證此次介接的時間區間。   綠界驗證時間區間暫訂為 3 分鐘內有效，超過則此次介接無效。    | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 

## Methods

### NewQueryCreditCardPeriodInfoRequest

`func NewQueryCreditCardPeriodInfoRequest(merchantID string, merchantTradeNo string, timeStamp int, checkMacValue string, ) *QueryCreditCardPeriodInfoRequest`

NewQueryCreditCardPeriodInfoRequest instantiates a new QueryCreditCardPeriodInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCreditCardPeriodInfoRequestWithDefaults

`func NewQueryCreditCardPeriodInfoRequestWithDefaults() *QueryCreditCardPeriodInfoRequest`

NewQueryCreditCardPeriodInfoRequestWithDefaults instantiates a new QueryCreditCardPeriodInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *QueryCreditCardPeriodInfoRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *QueryCreditCardPeriodInfoRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *QueryCreditCardPeriodInfoRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *QueryCreditCardPeriodInfoRequest) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *QueryCreditCardPeriodInfoRequest) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *QueryCreditCardPeriodInfoRequest) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetTimeStamp

`func (o *QueryCreditCardPeriodInfoRequest) GetTimeStamp() int`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *QueryCreditCardPeriodInfoRequest) GetTimeStampOk() (*int, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *QueryCreditCardPeriodInfoRequest) SetTimeStamp(v int)`

SetTimeStamp sets TimeStamp field to given value.


### GetCheckMacValue

`func (o *QueryCreditCardPeriodInfoRequest) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *QueryCreditCardPeriodInfoRequest) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *QueryCreditCardPeriodInfoRequest) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


