# QueryTradeInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號(由綠界提供)**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。  | 
**TimeStamp** | Pointer to **int** | **驗證時間**   將當下的時間轉為UnixTimeStamp(見範例)用於驗證此次介接的時間區間。   綠界驗證時間區間暫訂為 3 分鐘內有效，超過則此次介接無效。   | 
**PlatformID** | Pointer to **string** | **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 &#x60;MerchantID&#x60;。    | [optional] 
**CheckMacValue** | Pointer to **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 

## Methods

### NewQueryTradeInfoRequest

`func NewQueryTradeInfoRequest(merchantID string, merchantTradeNo string, timeStamp int, checkMacValue string, ) *QueryTradeInfoRequest`

NewQueryTradeInfoRequest instantiates a new QueryTradeInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTradeInfoRequestWithDefaults

`func NewQueryTradeInfoRequestWithDefaults() *QueryTradeInfoRequest`

NewQueryTradeInfoRequestWithDefaults instantiates a new QueryTradeInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *QueryTradeInfoRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *QueryTradeInfoRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *QueryTradeInfoRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *QueryTradeInfoRequest) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *QueryTradeInfoRequest) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *QueryTradeInfoRequest) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetTimeStamp

`func (o *QueryTradeInfoRequest) GetTimeStamp() int`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *QueryTradeInfoRequest) GetTimeStampOk() (*int, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *QueryTradeInfoRequest) SetTimeStamp(v int)`

SetTimeStamp sets TimeStamp field to given value.


### GetPlatformID

`func (o *QueryTradeInfoRequest) GetPlatformID() string`

GetPlatformID returns the PlatformID field if non-nil, zero value otherwise.

### GetPlatformIDOk

`func (o *QueryTradeInfoRequest) GetPlatformIDOk() (*string, bool)`

GetPlatformIDOk returns a tuple with the PlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformID

`func (o *QueryTradeInfoRequest) SetPlatformID(v string)`

SetPlatformID sets PlatformID field to given value.

### HasPlatformID

`func (o *QueryTradeInfoRequest) HasPlatformID() bool`

HasPlatformID returns a boolean if a field has been set.

### GetCheckMacValue

`func (o *QueryTradeInfoRequest) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *QueryTradeInfoRequest) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *QueryTradeInfoRequest) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


