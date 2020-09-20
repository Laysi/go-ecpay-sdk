# AioCheckOutAtmOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpireDate** | Pointer to **int32** | **允許繳費有效天數**   若需設定最長 60 天，最短 1 天。   未設定此參數則預設為 3 天   注意事項：以天為單位  | [optional] 
**PaymentInfoURL** | Pointer to **string** | **Server端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：銀行代碼、繳費虛擬帳號繳費期限…等)。   請參考[ATM、CVS 或 BARCODE 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。  | [optional] 
**ClientRedirectURL** | Pointer to **string** | **Client端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Client 端回傳消費者付款方式相關資訊(例：銀行代碼、繳費虛擬帳號繳費期限…等)且將頁面轉到特店指定的頁面。請參考[ATM、CVS 或 BARCODE 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。    | [optional] 

## Methods

### NewAioCheckOutAtmOption

`func NewAioCheckOutAtmOption() *AioCheckOutAtmOption`

NewAioCheckOutAtmOption instantiates a new AioCheckOutAtmOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutAtmOptionWithDefaults

`func NewAioCheckOutAtmOptionWithDefaults() *AioCheckOutAtmOption`

NewAioCheckOutAtmOptionWithDefaults instantiates a new AioCheckOutAtmOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpireDate

`func (o *AioCheckOutAtmOption) GetExpireDate() int32`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *AioCheckOutAtmOption) GetExpireDateOk() (*int32, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *AioCheckOutAtmOption) SetExpireDate(v int32)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *AioCheckOutAtmOption) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetPaymentInfoURL

`func (o *AioCheckOutAtmOption) GetPaymentInfoURL() string`

GetPaymentInfoURL returns the PaymentInfoURL field if non-nil, zero value otherwise.

### GetPaymentInfoURLOk

`func (o *AioCheckOutAtmOption) GetPaymentInfoURLOk() (*string, bool)`

GetPaymentInfoURLOk returns a tuple with the PaymentInfoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoURL

`func (o *AioCheckOutAtmOption) SetPaymentInfoURL(v string)`

SetPaymentInfoURL sets PaymentInfoURL field to given value.

### HasPaymentInfoURL

`func (o *AioCheckOutAtmOption) HasPaymentInfoURL() bool`

HasPaymentInfoURL returns a boolean if a field has been set.

### GetClientRedirectURL

`func (o *AioCheckOutAtmOption) GetClientRedirectURL() string`

GetClientRedirectURL returns the ClientRedirectURL field if non-nil, zero value otherwise.

### GetClientRedirectURLOk

`func (o *AioCheckOutAtmOption) GetClientRedirectURLOk() (*string, bool)`

GetClientRedirectURLOk returns a tuple with the ClientRedirectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRedirectURL

`func (o *AioCheckOutAtmOption) SetClientRedirectURL(v string)`

SetClientRedirectURL sets ClientRedirectURL field to given value.

### HasClientRedirectURL

`func (o *AioCheckOutAtmOption) HasClientRedirectURL() bool`

HasClientRedirectURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


