# AioCheckOutGeneralOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號(由綠界提供)**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號(由特店提供)**   特店交易編號均為唯一值，不可重複使用。   英數字大小寫混合   如何避免訂單編號重複請參考 FAQ   如有使用 &#x60;PlatformID&#x60; ，平台商底下所有商家之訂單編號亦不可重複。    | 
**StoreID** | Pointer to **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | [optional] 
**MerchantTradeDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **特店交易時間** 格式為 &#x60;yyyy/MM/dd HH:mm:ss&#x60;  | 
**PaymentType** | Pointer to [**AioCheckPaymentTypeEnum**](AioCheckPaymentTypeEnum.md) |  | [default to "aio"]
**TotalAmount** | Pointer to **int** | **交易金額**   請帶整數，不可有小數點。   僅限新台幣。   各付款金額的限制，請參考 &lt;https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID&#x3D;3605&gt;  | 
**TradeDesc** | Pointer to **string** | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。  | 
**ItemName** | Pointer to **string** | **商品名稱**   1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。   2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。    | 
**ReturnURL** | Pointer to **string** | **付款完成通知回傳網址**   當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：    1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。   2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。    | 
**ChoosePayment** | Pointer to [**ChoosePaymentEnum**](ChoosePaymentEnum.md) |  | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 
**ClientBackURL** | Pointer to **string** | **Client端返回特店的按鈕連結**   消費者點選此按鈕後，會將頁面導回到此設定的網址   注意事項：   導回時不會帶付款結果到此網址，只是將頁面導回而已。   設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。   設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。   若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。    | [optional] 
**ItemURL** | Pointer to **string** | **商品銷售網址**  | [optional] 
**Remark** | Pointer to **string** | **備註欄位**  | [optional] 
**ChooseSubPayment** | Pointer to [**ChooseSubPaymentEnum**](ChooseSubPaymentEnum.md) |  | [optional] 
**OrderResultURL** | Pointer to **string** | **Client端回傳付款結果網址**     當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：   1. 若與[ClientBackURL]同時設定，將會以此參數為主。   2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。    | [optional] 
**NeedExtraPaidInfo** | Pointer to [**NeedExtraPaidInfoEnum**](NeedExtraPaidInfoEnum.md) |  | [optional] 
**DeviceSource** | Pointer to **string** | **裝置來源** 請帶空值，由系統自動判定。  | [optional] 
**IgnorePayment** | Pointer to **string** | **隱藏付款**   當付款方式 &#x60;ChoosePayment&#x60; 為 &#x60;ALL&#x60; 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。   可用的參數值：   - &#x60;Credit&#x60;: 信用卡   - &#x60;WebATM&#x60;: 網路 ATM   - &#x60;ATM&#x60;: 自動櫃員機   - &#x60;CVS&#x60;: 超商代碼   - &#x60;BARCODE&#x60;: 超商條碼    | [optional] 
**PlatformID** | Pointer to **string** | **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 &#x60;MerchantID&#x60;。    | [optional] 
**InvoiceMark** | Pointer to [**InvoiceMarkEnum**](InvoiceMarkEnum.md) |  | [optional] 
**CustomField1** | Pointer to **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField2** | Pointer to **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField3** | Pointer to **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField4** | Pointer to **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**EncryptType** | Pointer to [**EncryptTypeEnum**](EncryptTypeEnum.md) |  | [default to ENCRYPTTYPEENUM_SHA256]
**Language** | Pointer to [**LanguageEnum**](LanguageEnum.md) |  | [optional] 

## Methods

### NewAioCheckOutGeneralOption

`func NewAioCheckOutGeneralOption(merchantID string, merchantTradeNo string, merchantTradeDate ECPayDateTime, paymentType AioCheckPaymentTypeEnum, totalAmount int, tradeDesc string, itemName string, returnURL string, choosePayment ChoosePaymentEnum, checkMacValue string, encryptType EncryptTypeEnum, ) *AioCheckOutGeneralOption`

NewAioCheckOutGeneralOption instantiates a new AioCheckOutGeneralOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutGeneralOptionWithDefaults

`func NewAioCheckOutGeneralOptionWithDefaults() *AioCheckOutGeneralOption`

NewAioCheckOutGeneralOptionWithDefaults instantiates a new AioCheckOutGeneralOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *AioCheckOutGeneralOption) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *AioCheckOutGeneralOption) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *AioCheckOutGeneralOption) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *AioCheckOutGeneralOption) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *AioCheckOutGeneralOption) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *AioCheckOutGeneralOption) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetStoreID

`func (o *AioCheckOutGeneralOption) GetStoreID() string`

GetStoreID returns the StoreID field if non-nil, zero value otherwise.

### GetStoreIDOk

`func (o *AioCheckOutGeneralOption) GetStoreIDOk() (*string, bool)`

GetStoreIDOk returns a tuple with the StoreID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreID

`func (o *AioCheckOutGeneralOption) SetStoreID(v string)`

SetStoreID sets StoreID field to given value.

### HasStoreID

`func (o *AioCheckOutGeneralOption) HasStoreID() bool`

HasStoreID returns a boolean if a field has been set.

### GetMerchantTradeDate

`func (o *AioCheckOutGeneralOption) GetMerchantTradeDate() ECPayDateTime`

GetMerchantTradeDate returns the MerchantTradeDate field if non-nil, zero value otherwise.

### GetMerchantTradeDateOk

`func (o *AioCheckOutGeneralOption) GetMerchantTradeDateOk() (*ECPayDateTime, bool)`

GetMerchantTradeDateOk returns a tuple with the MerchantTradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeDate

`func (o *AioCheckOutGeneralOption) SetMerchantTradeDate(v ECPayDateTime)`

SetMerchantTradeDate sets MerchantTradeDate field to given value.


### GetPaymentType

`func (o *AioCheckOutGeneralOption) GetPaymentType() AioCheckPaymentTypeEnum`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *AioCheckOutGeneralOption) GetPaymentTypeOk() (*AioCheckPaymentTypeEnum, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *AioCheckOutGeneralOption) SetPaymentType(v AioCheckPaymentTypeEnum)`

SetPaymentType sets PaymentType field to given value.


### GetTotalAmount

`func (o *AioCheckOutGeneralOption) GetTotalAmount() int`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *AioCheckOutGeneralOption) GetTotalAmountOk() (*int, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *AioCheckOutGeneralOption) SetTotalAmount(v int)`

SetTotalAmount sets TotalAmount field to given value.


### GetTradeDesc

`func (o *AioCheckOutGeneralOption) GetTradeDesc() string`

GetTradeDesc returns the TradeDesc field if non-nil, zero value otherwise.

### GetTradeDescOk

`func (o *AioCheckOutGeneralOption) GetTradeDescOk() (*string, bool)`

GetTradeDescOk returns a tuple with the TradeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDesc

`func (o *AioCheckOutGeneralOption) SetTradeDesc(v string)`

SetTradeDesc sets TradeDesc field to given value.


### GetItemName

`func (o *AioCheckOutGeneralOption) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *AioCheckOutGeneralOption) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *AioCheckOutGeneralOption) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetReturnURL

`func (o *AioCheckOutGeneralOption) GetReturnURL() string`

GetReturnURL returns the ReturnURL field if non-nil, zero value otherwise.

### GetReturnURLOk

`func (o *AioCheckOutGeneralOption) GetReturnURLOk() (*string, bool)`

GetReturnURLOk returns a tuple with the ReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnURL

`func (o *AioCheckOutGeneralOption) SetReturnURL(v string)`

SetReturnURL sets ReturnURL field to given value.


### GetChoosePayment

`func (o *AioCheckOutGeneralOption) GetChoosePayment() ChoosePaymentEnum`

GetChoosePayment returns the ChoosePayment field if non-nil, zero value otherwise.

### GetChoosePaymentOk

`func (o *AioCheckOutGeneralOption) GetChoosePaymentOk() (*ChoosePaymentEnum, bool)`

GetChoosePaymentOk returns a tuple with the ChoosePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoosePayment

`func (o *AioCheckOutGeneralOption) SetChoosePayment(v ChoosePaymentEnum)`

SetChoosePayment sets ChoosePayment field to given value.


### GetCheckMacValue

`func (o *AioCheckOutGeneralOption) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *AioCheckOutGeneralOption) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *AioCheckOutGeneralOption) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.


### GetClientBackURL

`func (o *AioCheckOutGeneralOption) GetClientBackURL() string`

GetClientBackURL returns the ClientBackURL field if non-nil, zero value otherwise.

### GetClientBackURLOk

`func (o *AioCheckOutGeneralOption) GetClientBackURLOk() (*string, bool)`

GetClientBackURLOk returns a tuple with the ClientBackURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBackURL

`func (o *AioCheckOutGeneralOption) SetClientBackURL(v string)`

SetClientBackURL sets ClientBackURL field to given value.

### HasClientBackURL

`func (o *AioCheckOutGeneralOption) HasClientBackURL() bool`

HasClientBackURL returns a boolean if a field has been set.

### GetItemURL

`func (o *AioCheckOutGeneralOption) GetItemURL() string`

GetItemURL returns the ItemURL field if non-nil, zero value otherwise.

### GetItemURLOk

`func (o *AioCheckOutGeneralOption) GetItemURLOk() (*string, bool)`

GetItemURLOk returns a tuple with the ItemURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemURL

`func (o *AioCheckOutGeneralOption) SetItemURL(v string)`

SetItemURL sets ItemURL field to given value.

### HasItemURL

`func (o *AioCheckOutGeneralOption) HasItemURL() bool`

HasItemURL returns a boolean if a field has been set.

### GetRemark

`func (o *AioCheckOutGeneralOption) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *AioCheckOutGeneralOption) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *AioCheckOutGeneralOption) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *AioCheckOutGeneralOption) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetChooseSubPayment

`func (o *AioCheckOutGeneralOption) GetChooseSubPayment() ChooseSubPaymentEnum`

GetChooseSubPayment returns the ChooseSubPayment field if non-nil, zero value otherwise.

### GetChooseSubPaymentOk

`func (o *AioCheckOutGeneralOption) GetChooseSubPaymentOk() (*ChooseSubPaymentEnum, bool)`

GetChooseSubPaymentOk returns a tuple with the ChooseSubPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChooseSubPayment

`func (o *AioCheckOutGeneralOption) SetChooseSubPayment(v ChooseSubPaymentEnum)`

SetChooseSubPayment sets ChooseSubPayment field to given value.

### HasChooseSubPayment

`func (o *AioCheckOutGeneralOption) HasChooseSubPayment() bool`

HasChooseSubPayment returns a boolean if a field has been set.

### GetOrderResultURL

`func (o *AioCheckOutGeneralOption) GetOrderResultURL() string`

GetOrderResultURL returns the OrderResultURL field if non-nil, zero value otherwise.

### GetOrderResultURLOk

`func (o *AioCheckOutGeneralOption) GetOrderResultURLOk() (*string, bool)`

GetOrderResultURLOk returns a tuple with the OrderResultURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderResultURL

`func (o *AioCheckOutGeneralOption) SetOrderResultURL(v string)`

SetOrderResultURL sets OrderResultURL field to given value.

### HasOrderResultURL

`func (o *AioCheckOutGeneralOption) HasOrderResultURL() bool`

HasOrderResultURL returns a boolean if a field has been set.

### GetNeedExtraPaidInfo

`func (o *AioCheckOutGeneralOption) GetNeedExtraPaidInfo() NeedExtraPaidInfoEnum`

GetNeedExtraPaidInfo returns the NeedExtraPaidInfo field if non-nil, zero value otherwise.

### GetNeedExtraPaidInfoOk

`func (o *AioCheckOutGeneralOption) GetNeedExtraPaidInfoOk() (*NeedExtraPaidInfoEnum, bool)`

GetNeedExtraPaidInfoOk returns a tuple with the NeedExtraPaidInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedExtraPaidInfo

`func (o *AioCheckOutGeneralOption) SetNeedExtraPaidInfo(v NeedExtraPaidInfoEnum)`

SetNeedExtraPaidInfo sets NeedExtraPaidInfo field to given value.

### HasNeedExtraPaidInfo

`func (o *AioCheckOutGeneralOption) HasNeedExtraPaidInfo() bool`

HasNeedExtraPaidInfo returns a boolean if a field has been set.

### GetDeviceSource

`func (o *AioCheckOutGeneralOption) GetDeviceSource() string`

GetDeviceSource returns the DeviceSource field if non-nil, zero value otherwise.

### GetDeviceSourceOk

`func (o *AioCheckOutGeneralOption) GetDeviceSourceOk() (*string, bool)`

GetDeviceSourceOk returns a tuple with the DeviceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSource

`func (o *AioCheckOutGeneralOption) SetDeviceSource(v string)`

SetDeviceSource sets DeviceSource field to given value.

### HasDeviceSource

`func (o *AioCheckOutGeneralOption) HasDeviceSource() bool`

HasDeviceSource returns a boolean if a field has been set.

### GetIgnorePayment

`func (o *AioCheckOutGeneralOption) GetIgnorePayment() string`

GetIgnorePayment returns the IgnorePayment field if non-nil, zero value otherwise.

### GetIgnorePaymentOk

`func (o *AioCheckOutGeneralOption) GetIgnorePaymentOk() (*string, bool)`

GetIgnorePaymentOk returns a tuple with the IgnorePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePayment

`func (o *AioCheckOutGeneralOption) SetIgnorePayment(v string)`

SetIgnorePayment sets IgnorePayment field to given value.

### HasIgnorePayment

`func (o *AioCheckOutGeneralOption) HasIgnorePayment() bool`

HasIgnorePayment returns a boolean if a field has been set.

### GetPlatformID

`func (o *AioCheckOutGeneralOption) GetPlatformID() string`

GetPlatformID returns the PlatformID field if non-nil, zero value otherwise.

### GetPlatformIDOk

`func (o *AioCheckOutGeneralOption) GetPlatformIDOk() (*string, bool)`

GetPlatformIDOk returns a tuple with the PlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformID

`func (o *AioCheckOutGeneralOption) SetPlatformID(v string)`

SetPlatformID sets PlatformID field to given value.

### HasPlatformID

`func (o *AioCheckOutGeneralOption) HasPlatformID() bool`

HasPlatformID returns a boolean if a field has been set.

### GetInvoiceMark

`func (o *AioCheckOutGeneralOption) GetInvoiceMark() InvoiceMarkEnum`

GetInvoiceMark returns the InvoiceMark field if non-nil, zero value otherwise.

### GetInvoiceMarkOk

`func (o *AioCheckOutGeneralOption) GetInvoiceMarkOk() (*InvoiceMarkEnum, bool)`

GetInvoiceMarkOk returns a tuple with the InvoiceMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMark

`func (o *AioCheckOutGeneralOption) SetInvoiceMark(v InvoiceMarkEnum)`

SetInvoiceMark sets InvoiceMark field to given value.

### HasInvoiceMark

`func (o *AioCheckOutGeneralOption) HasInvoiceMark() bool`

HasInvoiceMark returns a boolean if a field has been set.

### GetCustomField1

`func (o *AioCheckOutGeneralOption) GetCustomField1() string`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *AioCheckOutGeneralOption) GetCustomField1Ok() (*string, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *AioCheckOutGeneralOption) SetCustomField1(v string)`

SetCustomField1 sets CustomField1 field to given value.

### HasCustomField1

`func (o *AioCheckOutGeneralOption) HasCustomField1() bool`

HasCustomField1 returns a boolean if a field has been set.

### GetCustomField2

`func (o *AioCheckOutGeneralOption) GetCustomField2() string`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *AioCheckOutGeneralOption) GetCustomField2Ok() (*string, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *AioCheckOutGeneralOption) SetCustomField2(v string)`

SetCustomField2 sets CustomField2 field to given value.

### HasCustomField2

`func (o *AioCheckOutGeneralOption) HasCustomField2() bool`

HasCustomField2 returns a boolean if a field has been set.

### GetCustomField3

`func (o *AioCheckOutGeneralOption) GetCustomField3() string`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *AioCheckOutGeneralOption) GetCustomField3Ok() (*string, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *AioCheckOutGeneralOption) SetCustomField3(v string)`

SetCustomField3 sets CustomField3 field to given value.

### HasCustomField3

`func (o *AioCheckOutGeneralOption) HasCustomField3() bool`

HasCustomField3 returns a boolean if a field has been set.

### GetCustomField4

`func (o *AioCheckOutGeneralOption) GetCustomField4() string`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *AioCheckOutGeneralOption) GetCustomField4Ok() (*string, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *AioCheckOutGeneralOption) SetCustomField4(v string)`

SetCustomField4 sets CustomField4 field to given value.

### HasCustomField4

`func (o *AioCheckOutGeneralOption) HasCustomField4() bool`

HasCustomField4 returns a boolean if a field has been set.

### GetEncryptType

`func (o *AioCheckOutGeneralOption) GetEncryptType() EncryptTypeEnum`

GetEncryptType returns the EncryptType field if non-nil, zero value otherwise.

### GetEncryptTypeOk

`func (o *AioCheckOutGeneralOption) GetEncryptTypeOk() (*EncryptTypeEnum, bool)`

GetEncryptTypeOk returns a tuple with the EncryptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptType

`func (o *AioCheckOutGeneralOption) SetEncryptType(v EncryptTypeEnum)`

SetEncryptType sets EncryptType field to given value.


### GetLanguage

`func (o *AioCheckOutGeneralOption) GetLanguage() LanguageEnum`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AioCheckOutGeneralOption) GetLanguageOk() (*LanguageEnum, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AioCheckOutGeneralOption) SetLanguage(v LanguageEnum)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AioCheckOutGeneralOption) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


