# AioCheckOutChoosePaymentCVSBARCODEOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreExpireDate** | Pointer to **int32** | **超商繳費截止時間** 注意事項： &#x60;CVS&#x60;:以分鐘為單位 &#x60;BARCODE&#x60;:以天為單位 若未設定此參數，&#x60;CVS&#x60; 預設為 &#x60;10080&#x60; 分鐘(&#x60;7&#x60; 天)；BARCODE 預設為 &#x60;7&#x60; 天。 若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 &#x60;86400&#x60; 分鐘，超過時一律以 &#x60;86400&#x60; 分鐘計(&#x60;60&#x60; 天) 例：&#x60;08/01&#x60; 的 &#x60;20:15&#x60; 分購買商品，繳費期限為 &#x60;7&#x60; 天，表示 &#x60;8/08&#x60; 的 &#x60;20:15&#x60; 分前您必須前往超商繳費。 範例: &#x60;CVS&#x60;&#x3D;&#x60;1440&#x60;(共 &#x60;1&#x60; 天)、&#x60;BARCODE&#x60;&#x3D;&#x60;7&#x60;(共 &#x60;7&#x60; 天)  | [optional] 
**Desc1** | Pointer to **string** | **交易描述1** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc2** | Pointer to **string** | **交易描述2** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc3** | Pointer to **string** | **交易描述3** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc4** | Pointer to **string** | **交易描述4** 會出現在超商繳費平台螢幕上  | [optional] 
**PaymentInfoURL** | Pointer to **string** | **Server 端回傳付款相關資訊** 若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。 請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.] 注意事項： 頁面將會停留在綠界，顯示繳費的相關資訊。 回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。  | [optional] 
**ClientRedirectURL** | Pointer to **string** | **Client端回傳付款方式相關資訊** 若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。 請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.] 注意事項： 若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。 若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。 回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。  | [optional] 

## Methods

### NewAioCheckOutChoosePaymentCVSBARCODEOption

`func NewAioCheckOutChoosePaymentCVSBARCODEOption() *AioCheckOutChoosePaymentCVSBARCODEOption`

NewAioCheckOutChoosePaymentCVSBARCODEOption instantiates a new AioCheckOutChoosePaymentCVSBARCODEOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutChoosePaymentCVSBARCODEOptionWithDefaults

`func NewAioCheckOutChoosePaymentCVSBARCODEOptionWithDefaults() *AioCheckOutChoosePaymentCVSBARCODEOption`

NewAioCheckOutChoosePaymentCVSBARCODEOptionWithDefaults instantiates a new AioCheckOutChoosePaymentCVSBARCODEOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreExpireDate

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetStoreExpireDate() int32`

GetStoreExpireDate returns the StoreExpireDate field if non-nil, zero value otherwise.

### GetStoreExpireDateOk

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetStoreExpireDateOk() (*int32, bool)`

GetStoreExpireDateOk returns a tuple with the StoreExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreExpireDate

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) SetStoreExpireDate(v int32)`

SetStoreExpireDate sets StoreExpireDate field to given value.

### HasStoreExpireDate

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) HasStoreExpireDate() bool`

HasStoreExpireDate returns a boolean if a field has been set.

### GetDesc1

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc1() string`

GetDesc1 returns the Desc1 field if non-nil, zero value otherwise.

### GetDesc1Ok

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc1Ok() (*string, bool)`

GetDesc1Ok returns a tuple with the Desc1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc1

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) SetDesc1(v string)`

SetDesc1 sets Desc1 field to given value.

### HasDesc1

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) HasDesc1() bool`

HasDesc1 returns a boolean if a field has been set.

### GetDesc2

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc2() string`

GetDesc2 returns the Desc2 field if non-nil, zero value otherwise.

### GetDesc2Ok

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc2Ok() (*string, bool)`

GetDesc2Ok returns a tuple with the Desc2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc2

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) SetDesc2(v string)`

SetDesc2 sets Desc2 field to given value.

### HasDesc2

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) HasDesc2() bool`

HasDesc2 returns a boolean if a field has been set.

### GetDesc3

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc3() string`

GetDesc3 returns the Desc3 field if non-nil, zero value otherwise.

### GetDesc3Ok

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc3Ok() (*string, bool)`

GetDesc3Ok returns a tuple with the Desc3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc3

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) SetDesc3(v string)`

SetDesc3 sets Desc3 field to given value.

### HasDesc3

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) HasDesc3() bool`

HasDesc3 returns a boolean if a field has been set.

### GetDesc4

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc4() string`

GetDesc4 returns the Desc4 field if non-nil, zero value otherwise.

### GetDesc4Ok

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetDesc4Ok() (*string, bool)`

GetDesc4Ok returns a tuple with the Desc4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc4

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) SetDesc4(v string)`

SetDesc4 sets Desc4 field to given value.

### HasDesc4

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) HasDesc4() bool`

HasDesc4 returns a boolean if a field has been set.

### GetPaymentInfoURL

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetPaymentInfoURL() string`

GetPaymentInfoURL returns the PaymentInfoURL field if non-nil, zero value otherwise.

### GetPaymentInfoURLOk

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetPaymentInfoURLOk() (*string, bool)`

GetPaymentInfoURLOk returns a tuple with the PaymentInfoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoURL

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) SetPaymentInfoURL(v string)`

SetPaymentInfoURL sets PaymentInfoURL field to given value.

### HasPaymentInfoURL

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) HasPaymentInfoURL() bool`

HasPaymentInfoURL returns a boolean if a field has been set.

### GetClientRedirectURL

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetClientRedirectURL() string`

GetClientRedirectURL returns the ClientRedirectURL field if non-nil, zero value otherwise.

### GetClientRedirectURLOk

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) GetClientRedirectURLOk() (*string, bool)`

GetClientRedirectURLOk returns a tuple with the ClientRedirectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRedirectURL

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) SetClientRedirectURL(v string)`

SetClientRedirectURL sets ClientRedirectURL field to given value.

### HasClientRedirectURL

`func (o *AioCheckOutChoosePaymentCVSBARCODEOption) HasClientRedirectURL() bool`

HasClientRedirectURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


