# AioCheckOutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | 特店編號(由綠界提供) | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號(由特店提供)** 特店交易編號均為唯一值，不可重複使用。 英數字大小寫混合 如何避免訂單編號重複請參考 FAQ 如有使用 &#x60;PlatformID&#x60; ，平台商底下所有商家之訂單編號亦不可重複。  | 
**StoreID** | Pointer to **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | [optional] 
**MerchantTradeDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **特店交易時間** 格式為 &#x60;yyyy/MM/dd HH:mm:ss&#x60;  | 
**PaymentType** | Pointer to **string** | **交易類型** 請固定填入 &#x60;aio&#x60;  | [default to "aio"]
**TotalAmount** | Pointer to **int32** | **交易金額** 請帶整數，不可有小數點。 僅限新台幣。 各付款金額的限制，請參考 &lt;https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID&#x3D;3605&gt;  | 
**TradeDesc** | Pointer to **string** | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。  | 
**ItemName** | Pointer to **string** | **商品名稱** 1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。 2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。  | 
**ReturnURL** | Pointer to **string** | **付款完成通知回傳網址** 當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。 2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。  | 
**ChoosePayment** | Pointer to **string** | **選擇預設付款方式** 綠界提供下列付款方式，請於建立訂單時傳送過來: - &#x60;Credit&#x60;: 信用卡及銀聯卡(需申請開通) - &#x60;WebATM&#x60;: 網路 ATM - &#x60;ATM&#x60;: 自動櫃員機 - &#x60;CVS&#x60;: 超商代碼 - &#x60;BARCODE&#x60;: 超商條碼 - &#x60;ALL&#x60;: 不指定付款方式，由綠界顯示付款方式選擇頁面。  注意事項： 1.若為手機版時不支援下列付款方式:   - WebATM:網路 ATM  2.如需要不透過綠界畫面取得 &#x60;ATM&#x60;、&#x60;CVS&#x60;、&#x60;BARCODE&#x60; 的繳費代碼，請參考 FAQ。  | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 
**ClientBackURL** | Pointer to **string** | **Client 端返回特店的按鈕連結** 消費者點選此按鈕後，會將頁面導回到此設定的網址 注意事項： 導回時不會帶付款結果到此網址，只是將頁面導回而已。 設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。 設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。 若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。 若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。  | [optional] 
**ItemURL** | Pointer to **string** | **商品銷售網址**  | [optional] 
**Remark** | Pointer to **string** | **備註欄位**  | [optional] 
**ChooseSubPayment** | Pointer to **string** | **付款子項目** 若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。 請參考付款方式一覽表  | [optional] 
**OrderResultURL** | Pointer to **string** | **Client端回傳付款結果網址** 當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 若與[ClientBackURL]同時設定，將會以此參數為主。 2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。  | [optional] 
**NeedExtraPaidInfo** | Pointer to **string** | **是否需要額外的付款資訊** 額外的付款資訊： 若不回傳額外的付款資訊時，參數值請傳：&#x60;Ｎ&#x60;； 若要回傳額外的付款資訊時，參數值請傳：&#x60;Ｙ&#x60;，付款完成後綠界會以 Server POST 方式回傳額外付款資訊。 注意事項： 回傳額外付款資訊參數請參考-額外回傳的參數  | [optional] [default to "N"]
**DeviceSource** | Pointer to **string** | **裝置來源** 請帶空值，由系統自動判定。  | [optional] 
**IgnorePayment** | Pointer to **string** | **隱藏付款** 當付款方式 &#x60;ChoosePayment&#x60; 為 &#x60;ALL&#x60; 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。 可用的參數值： - &#x60;Credit&#x60;: 信用卡 - &#x60;WebATM&#x60;: 網路 ATM - &#x60;ATM&#x60;: 自動櫃員機 - &#x60;CVS&#x60;: 超商代碼 - &#x60;BARCODE&#x60;: 超商條碼  | [optional] 
**PlatformID** | Pointer to **string** | **特約合作平台商代號(由綠界提供)** 為專案合作的平台商使用。 一般特店或平台商本身介接，則參數請帶放空值。 若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 &#x60;MerchantID&#x60;。  | [optional] 
**InvoiceMark** | Pointer to **string** | **電子發票開立註記** 此參數為付款完成後同時開立電子發票。 若要使用時，該參數須設定為「Y」，同時還要設定「電子發票介接相關參數」 注意事項： 正式環境欲使用電子發票功能，須與綠界申請開通，若未開通請致電客服中心 &#x60;(02) 2655-1775&#x60;。  | [optional] 
**CustomField1** | Pointer to **string** | **自訂名稱欄位1** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField2** | Pointer to **string** | **自訂名稱欄位2** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField3** | Pointer to **string** | **自訂名稱欄位3** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField4** | Pointer to **string** | **自訂名稱欄位4** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**EncryptType** | Pointer to **int32** | **CheckMacValue加密類型** 請固定填入 1，使用 SHA256 加密。  | [optional] [default to 1]
**Language** | Pointer to **string** | **語系設定** 預設語系為中文，若要變更語系參數值請帶： - 英語：&#x60;ENG&#x60; - 韓語：&#x60;KOR&#x60; - 日語：&#x60;JPN&#x60; - 簡體中文：&#x60;CHI&#x60;  | [optional] 

## Methods

### NewAioCheckOutRequest

`func NewAioCheckOutRequest(merchantID string, merchantTradeNo string, merchantTradeDate ECPayDateTime, paymentType string, totalAmount int32, tradeDesc string, itemName string, returnURL string, choosePayment string, checkMacValue string, ) *AioCheckOutRequest`

NewAioCheckOutRequest instantiates a new AioCheckOutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutRequestWithDefaults

`func NewAioCheckOutRequestWithDefaults() *AioCheckOutRequest`

NewAioCheckOutRequestWithDefaults instantiates a new AioCheckOutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *AioCheckOutRequest) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *AioCheckOutRequest) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *AioCheckOutRequest) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *AioCheckOutRequest) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *AioCheckOutRequest) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *AioCheckOutRequest) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetStoreID

`func (o *AioCheckOutRequest) GetStoreID() string`

GetStoreID returns the StoreID field if non-nil, zero value otherwise.

### GetStoreIDOk

`func (o *AioCheckOutRequest) GetStoreIDOk() (*string, bool)`

GetStoreIDOk returns a tuple with the StoreID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreID

`func (o *AioCheckOutRequest) SetStoreID(v string)`

SetStoreID sets StoreID field to given value.

### HasStoreID

`func (o *AioCheckOutRequest) HasStoreID() bool`

HasStoreID returns a boolean if a field has been set.

### GetMerchantTradeDate

`func (o *AioCheckOutRequest) GetMerchantTradeDate() ECPayDateTime`

GetMerchantTradeDate returns the MerchantTradeDate field if non-nil, zero value otherwise.

### GetMerchantTradeDateOk

`func (o *AioCheckOutRequest) GetMerchantTradeDateOk() (*ECPayDateTime, bool)`

GetMerchantTradeDateOk returns a tuple with the MerchantTradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeDate

`func (o *AioCheckOutRequest) SetMerchantTradeDate(v ECPayDateTime)`

SetMerchantTradeDate sets MerchantTradeDate field to given value.


### GetPaymentType

`func (o *AioCheckOutRequest) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *AioCheckOutRequest) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *AioCheckOutRequest) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetTotalAmount

`func (o *AioCheckOutRequest) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *AioCheckOutRequest) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *AioCheckOutRequest) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetTradeDesc

`func (o *AioCheckOutRequest) GetTradeDesc() string`

GetTradeDesc returns the TradeDesc field if non-nil, zero value otherwise.

### GetTradeDescOk

`func (o *AioCheckOutRequest) GetTradeDescOk() (*string, bool)`

GetTradeDescOk returns a tuple with the TradeDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDesc

`func (o *AioCheckOutRequest) SetTradeDesc(v string)`

SetTradeDesc sets TradeDesc field to given value.


### GetItemName

`func (o *AioCheckOutRequest) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *AioCheckOutRequest) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *AioCheckOutRequest) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetReturnURL

`func (o *AioCheckOutRequest) GetReturnURL() string`

GetReturnURL returns the ReturnURL field if non-nil, zero value otherwise.

### GetReturnURLOk

`func (o *AioCheckOutRequest) GetReturnURLOk() (*string, bool)`

GetReturnURLOk returns a tuple with the ReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnURL

`func (o *AioCheckOutRequest) SetReturnURL(v string)`

SetReturnURL sets ReturnURL field to given value.


### GetChoosePayment

`func (o *AioCheckOutRequest) GetChoosePayment() string`

GetChoosePayment returns the ChoosePayment field if non-nil, zero value otherwise.

### GetChoosePaymentOk

`func (o *AioCheckOutRequest) GetChoosePaymentOk() (*string, bool)`

GetChoosePaymentOk returns a tuple with the ChoosePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoosePayment

`func (o *AioCheckOutRequest) SetChoosePayment(v string)`

SetChoosePayment sets ChoosePayment field to given value.


### GetCheckMacValue

`func (o *AioCheckOutRequest) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *AioCheckOutRequest) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *AioCheckOutRequest) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.


### GetClientBackURL

`func (o *AioCheckOutRequest) GetClientBackURL() string`

GetClientBackURL returns the ClientBackURL field if non-nil, zero value otherwise.

### GetClientBackURLOk

`func (o *AioCheckOutRequest) GetClientBackURLOk() (*string, bool)`

GetClientBackURLOk returns a tuple with the ClientBackURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBackURL

`func (o *AioCheckOutRequest) SetClientBackURL(v string)`

SetClientBackURL sets ClientBackURL field to given value.

### HasClientBackURL

`func (o *AioCheckOutRequest) HasClientBackURL() bool`

HasClientBackURL returns a boolean if a field has been set.

### GetItemURL

`func (o *AioCheckOutRequest) GetItemURL() string`

GetItemURL returns the ItemURL field if non-nil, zero value otherwise.

### GetItemURLOk

`func (o *AioCheckOutRequest) GetItemURLOk() (*string, bool)`

GetItemURLOk returns a tuple with the ItemURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemURL

`func (o *AioCheckOutRequest) SetItemURL(v string)`

SetItemURL sets ItemURL field to given value.

### HasItemURL

`func (o *AioCheckOutRequest) HasItemURL() bool`

HasItemURL returns a boolean if a field has been set.

### GetRemark

`func (o *AioCheckOutRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *AioCheckOutRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *AioCheckOutRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *AioCheckOutRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetChooseSubPayment

`func (o *AioCheckOutRequest) GetChooseSubPayment() string`

GetChooseSubPayment returns the ChooseSubPayment field if non-nil, zero value otherwise.

### GetChooseSubPaymentOk

`func (o *AioCheckOutRequest) GetChooseSubPaymentOk() (*string, bool)`

GetChooseSubPaymentOk returns a tuple with the ChooseSubPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChooseSubPayment

`func (o *AioCheckOutRequest) SetChooseSubPayment(v string)`

SetChooseSubPayment sets ChooseSubPayment field to given value.

### HasChooseSubPayment

`func (o *AioCheckOutRequest) HasChooseSubPayment() bool`

HasChooseSubPayment returns a boolean if a field has been set.

### GetOrderResultURL

`func (o *AioCheckOutRequest) GetOrderResultURL() string`

GetOrderResultURL returns the OrderResultURL field if non-nil, zero value otherwise.

### GetOrderResultURLOk

`func (o *AioCheckOutRequest) GetOrderResultURLOk() (*string, bool)`

GetOrderResultURLOk returns a tuple with the OrderResultURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderResultURL

`func (o *AioCheckOutRequest) SetOrderResultURL(v string)`

SetOrderResultURL sets OrderResultURL field to given value.

### HasOrderResultURL

`func (o *AioCheckOutRequest) HasOrderResultURL() bool`

HasOrderResultURL returns a boolean if a field has been set.

### GetNeedExtraPaidInfo

`func (o *AioCheckOutRequest) GetNeedExtraPaidInfo() string`

GetNeedExtraPaidInfo returns the NeedExtraPaidInfo field if non-nil, zero value otherwise.

### GetNeedExtraPaidInfoOk

`func (o *AioCheckOutRequest) GetNeedExtraPaidInfoOk() (*string, bool)`

GetNeedExtraPaidInfoOk returns a tuple with the NeedExtraPaidInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedExtraPaidInfo

`func (o *AioCheckOutRequest) SetNeedExtraPaidInfo(v string)`

SetNeedExtraPaidInfo sets NeedExtraPaidInfo field to given value.

### HasNeedExtraPaidInfo

`func (o *AioCheckOutRequest) HasNeedExtraPaidInfo() bool`

HasNeedExtraPaidInfo returns a boolean if a field has been set.

### GetDeviceSource

`func (o *AioCheckOutRequest) GetDeviceSource() string`

GetDeviceSource returns the DeviceSource field if non-nil, zero value otherwise.

### GetDeviceSourceOk

`func (o *AioCheckOutRequest) GetDeviceSourceOk() (*string, bool)`

GetDeviceSourceOk returns a tuple with the DeviceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSource

`func (o *AioCheckOutRequest) SetDeviceSource(v string)`

SetDeviceSource sets DeviceSource field to given value.

### HasDeviceSource

`func (o *AioCheckOutRequest) HasDeviceSource() bool`

HasDeviceSource returns a boolean if a field has been set.

### GetIgnorePayment

`func (o *AioCheckOutRequest) GetIgnorePayment() string`

GetIgnorePayment returns the IgnorePayment field if non-nil, zero value otherwise.

### GetIgnorePaymentOk

`func (o *AioCheckOutRequest) GetIgnorePaymentOk() (*string, bool)`

GetIgnorePaymentOk returns a tuple with the IgnorePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePayment

`func (o *AioCheckOutRequest) SetIgnorePayment(v string)`

SetIgnorePayment sets IgnorePayment field to given value.

### HasIgnorePayment

`func (o *AioCheckOutRequest) HasIgnorePayment() bool`

HasIgnorePayment returns a boolean if a field has been set.

### GetPlatformID

`func (o *AioCheckOutRequest) GetPlatformID() string`

GetPlatformID returns the PlatformID field if non-nil, zero value otherwise.

### GetPlatformIDOk

`func (o *AioCheckOutRequest) GetPlatformIDOk() (*string, bool)`

GetPlatformIDOk returns a tuple with the PlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformID

`func (o *AioCheckOutRequest) SetPlatformID(v string)`

SetPlatformID sets PlatformID field to given value.

### HasPlatformID

`func (o *AioCheckOutRequest) HasPlatformID() bool`

HasPlatformID returns a boolean if a field has been set.

### GetInvoiceMark

`func (o *AioCheckOutRequest) GetInvoiceMark() string`

GetInvoiceMark returns the InvoiceMark field if non-nil, zero value otherwise.

### GetInvoiceMarkOk

`func (o *AioCheckOutRequest) GetInvoiceMarkOk() (*string, bool)`

GetInvoiceMarkOk returns a tuple with the InvoiceMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMark

`func (o *AioCheckOutRequest) SetInvoiceMark(v string)`

SetInvoiceMark sets InvoiceMark field to given value.

### HasInvoiceMark

`func (o *AioCheckOutRequest) HasInvoiceMark() bool`

HasInvoiceMark returns a boolean if a field has been set.

### GetCustomField1

`func (o *AioCheckOutRequest) GetCustomField1() string`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *AioCheckOutRequest) GetCustomField1Ok() (*string, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *AioCheckOutRequest) SetCustomField1(v string)`

SetCustomField1 sets CustomField1 field to given value.

### HasCustomField1

`func (o *AioCheckOutRequest) HasCustomField1() bool`

HasCustomField1 returns a boolean if a field has been set.

### GetCustomField2

`func (o *AioCheckOutRequest) GetCustomField2() string`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *AioCheckOutRequest) GetCustomField2Ok() (*string, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *AioCheckOutRequest) SetCustomField2(v string)`

SetCustomField2 sets CustomField2 field to given value.

### HasCustomField2

`func (o *AioCheckOutRequest) HasCustomField2() bool`

HasCustomField2 returns a boolean if a field has been set.

### GetCustomField3

`func (o *AioCheckOutRequest) GetCustomField3() string`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *AioCheckOutRequest) GetCustomField3Ok() (*string, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *AioCheckOutRequest) SetCustomField3(v string)`

SetCustomField3 sets CustomField3 field to given value.

### HasCustomField3

`func (o *AioCheckOutRequest) HasCustomField3() bool`

HasCustomField3 returns a boolean if a field has been set.

### GetCustomField4

`func (o *AioCheckOutRequest) GetCustomField4() string`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *AioCheckOutRequest) GetCustomField4Ok() (*string, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *AioCheckOutRequest) SetCustomField4(v string)`

SetCustomField4 sets CustomField4 field to given value.

### HasCustomField4

`func (o *AioCheckOutRequest) HasCustomField4() bool`

HasCustomField4 returns a boolean if a field has been set.

### GetEncryptType

`func (o *AioCheckOutRequest) GetEncryptType() int32`

GetEncryptType returns the EncryptType field if non-nil, zero value otherwise.

### GetEncryptTypeOk

`func (o *AioCheckOutRequest) GetEncryptTypeOk() (*int32, bool)`

GetEncryptTypeOk returns a tuple with the EncryptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptType

`func (o *AioCheckOutRequest) SetEncryptType(v int32)`

SetEncryptType sets EncryptType field to given value.

### HasEncryptType

`func (o *AioCheckOutRequest) HasEncryptType() bool`

HasEncryptType returns a boolean if a field has been set.

### GetLanguage

`func (o *AioCheckOutRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AioCheckOutRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AioCheckOutRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AioCheckOutRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

