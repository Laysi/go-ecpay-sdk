# AioCheckOutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號(由綠界提供)**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號(由特店提供)**   特店交易編號均為唯一值，不可重複使用。   英數字大小寫混合   如何避免訂單編號重複請參考 FAQ   如有使用 &#x60;PlatformID&#x60; ，平台商底下所有商家之訂單編號亦不可重複。    | 
**StoreID** | Pointer to **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | [optional] 
**MerchantTradeDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **特店交易時間** 格式為 &#x60;yyyy/MM/dd HH:mm:ss&#x60;  | 
**PaymentType** | Pointer to [**AioCheckPaymentTypeEnum**](AioCheckPaymentTypeEnum.md) |  | [default to "aio"]
**TotalAmount** | Pointer to **int32** | **交易金額**   請帶整數，不可有小數點。   僅限新台幣。   各付款金額的限制，請參考 &lt;https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID&#x3D;3605&gt;  | 
**TradeDesc** | Pointer to **string** | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。  | 
**ItemName** | Pointer to **string** | **商品名稱**   1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。   2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。    | 
**ReturnURL** | Pointer to **string** | **付款完成通知回傳網址**   當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：    1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。   2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。    | 
**ChoosePayment** | Pointer to [**ChoosePaymentEnum**](ChoosePaymentEnum.md) |  | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 
**ClientBackURL** | Pointer to **string** | **Client端返回特店的按鈕連結**   消費者點選此按鈕後，會將頁面導回到此設定的網址   注意事項：   導回時不會帶付款結果到此網址，只是將頁面導回而已。   設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。   設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。   若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。    | [optional] 
**ItemURL** | Pointer to **string** | **商品銷售網址**  | [optional] 
**Remark** | Pointer to **string** | **備註欄位**  | [optional] 
**ChooseSubPayment** | Pointer to **string** | **付款子項目**   若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。   請參考付款方式一覽表    | [optional] 
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
**StoreExpireDate** | Pointer to **int32** | **超商繳費截止時間**   注意事項：   &#x60;CVS&#x60;:以分鐘為單位   &#x60;BARCODE&#x60;:以天為單位   若未設定此參數，&#x60;CVS&#x60; 預設為 &#x60;10080&#x60; 分鐘(&#x60;7&#x60; 天)；BARCODE 預設為 &#x60;7&#x60; 天。   若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 &#x60;86400&#x60; 分鐘，超過時一律以 &#x60;86400&#x60; 分鐘計(&#x60;60&#x60; 天)   例：&#x60;08/01&#x60; 的 &#x60;20:15&#x60; 分購買商品，繳費期限為 &#x60;7&#x60; 天，表示 &#x60;8/08&#x60; 的 &#x60;20:15&#x60; 分前您必須前往超商繳費。   範例: &#x60;CVS&#x60;&#x3D;&#x60;1440&#x60;(共 &#x60;1&#x60; 天)、&#x60;BARCODE&#x60;&#x3D;&#x60;7&#x60;(共 &#x60;7&#x60; 天)    | [optional] 
**Desc1** | Pointer to **string** | **交易描述1** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc2** | Pointer to **string** | **交易描述2** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc3** | Pointer to **string** | **交易描述3** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc4** | Pointer to **string** | **交易描述4** 會出現在超商繳費平台螢幕上  | [optional] 
**PaymentInfoURL** | Pointer to **string** | **Server 端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。   請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。    | [optional] 
**ClientRedirectURL** | Pointer to **string** | **Client端回傳付款方式相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。   請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。    | [optional] 
**BindingCard** | Pointer to [**BindingCardEnum**](BindingCardEnum.md) |  | [optional] 
**MerchantMemberID** | Pointer to **string** | **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 &#x60;MerchantID&#x60; + &#x60;廠商會員編號&#x60;)  | [optional] 
**Redeem** | Pointer to [**RedeemEnum**](RedeemEnum.md) |  | [optional] 
**UnionPay** | Pointer to [**UnionPayEnum**](UnionPayEnum.md) |  | [optional] [default to UNIONPAYENUM_NOT_SPECIFIED]
**CreditInstallment** | Pointer to **string** | **刷卡分期期數**    提供刷卡分期期數   信用卡分期可用參數為:3,6,12,18,24   注意事項：   使用的期數必須先透過申請開通後方能使用，並以申請開通的期數為主。    | [optional] 
**PeriodAmount** | Pointer to **int32** | **每次授權金額**   每次要授權(扣款)的金額。   注意事項：   綠界會依此次授權金額&#x60;PeriodAmount&#x60;所設定的金額做為之後固定授權的金額。   交易金額&#x60;TotalAmount&#x60;設定金額必須和授權金額&#x60;PeriodAmount&#x60;相同。   請帶整數，不可有小數點。僅限新台幣。  | [optional] 
**PeriodType** | Pointer to [**CreditPeriodTypeEnum**](CreditPeriodTypeEnum.md) |  | [optional] 
**Frequency** | Pointer to **int32** | **執行頻率**   此參數用來定義多久要執行一次   注意事項：   至少要大於等於 1 次以上。   當 &#x60;PeriodType&#x60; 設為 &#x60;D&#x60; 時，最多可設 &#x60;365&#x60; 次。    當 &#x60;PeriodType&#x60; 設為 &#x60;M&#x60; 時，最多可設 &#x60;12&#x60; 次。     當 &#x60;PeriodType&#x60; 設為 &#x60;Y&#x60; 時，最多可設 &#x60;1&#x60; 次。  | [optional] 
**ExecTimes** | Pointer to **int32** | **執行次數**   總共要執行幾次。   注意事項：   至少要大於 1 次以上。   當 &#x60;PeriodType&#x60; 設為 &#x60;D&#x60; 時，最多可設 &#x60;999&#x60;次。   當 &#x60;PeriodType&#x60; 設為 &#x60;M&#x60; 時，最多可設 &#x60;99&#x60;次。   當 &#x60;PeriodType&#x60; 設為 &#x60;Y&#x60; 時，最多可設 &#x60;9&#x60; 次。    例 1：   當信用卡定期定額扣款為每個月扣 1 次500 元，總共要扣 12次   &#x60;TotalAmount&#x60;參數請帶 &#x60;500&#x60; &#x60;PeriodAmount&#x60;&#x3D;&#x60;500&#x60;   &#x60;PeriodType&#x60;&#x3D;&#x60;M&#x60;   &#x60;Frequency&#x60;&#x3D;&#x60;1&#x60;   &#x60;ExecTimes&#x60;&#x3D;&#x60;12&#x60;    例 2：   當信用卡定期定額扣款為 6000 元，每 6 個月扣 1 次，總共要扣 2 次時    交易金額&#x60;TotalAmount&#x60;參數請帶 &#x60;6000&#x60;   &#x60;PeriodType&#x60;&#x3D;&#x60;M&#x60;   &#x60;Frequency&#x60;&#x3D;&#x60;6&#x60;   &#x60;ExecTimes&#x60;&#x3D;&#x60;2&#x60;    | [optional] 
**PeriodReturnURL** | Pointer to **string** | **定期定額的執行結果回應URL**   若交易是信用卡定期定額的方式，則每次執行授權完，會將授權結果回傳到這個設定的 URL。   回覆內容請參考。    | [optional] 
**RelateNumber** | Pointer to **string** | **特店自訂編號** 此為特店自訂編號，編號均為唯一值不可重複使用。  | [optional] 
**CustomerID** | Pointer to **string** | **客戶編號** 該參數有值時，僅接受『英文、數字、下底線』等字元。  | [optional] 
**CustomerIdentifier** | Pointer to **string** | **統一編號** 該參數有值時，請帶固定長度為數字 8 碼。  | [optional] 
**CustomerName** | Pointer to **string** | **客戶名稱**    當列印註記&#x60;Print&#x60;為 &#x60;1&#x60;(列印)時，則該參數必須有值。   該參數有值時，僅接受『中、英文及數字』等字元。   請將參數值做 UrlEncode 方式編碼。    | [optional] 
**CustomerAddr** | Pointer to **string** | **客戶地址**    當列印註記&#x60;Print&#x60;為 &#x60;1&#x60;(列印)時，則該參數必須有值。   當該參數有值時，請注意特殊字元轉換 。    請將參數值做 UrlEncode 方式編碼。  | [optional] 
**CustomerPhone** | Pointer to **string** | **客戶手機號碼**   當客戶電子信箱&#x60;CustomerEmail&#x60;為空字串時，則該參數必須有值。   當該參數有值時，則格式為數字。   注意事項：   請填手機號碼，不能填市話因為要收簡訊通知用    | [optional] 
**CustomerEmail** | Pointer to **string** | **客戶電子信箱**   當客戶手機號碼&#x60;CustomerPhone&#x60;為空字串時，則該參數必須有值。   當該參數有值時，則格式需符合 EMAIL格式。   請將參數值做 UrlEncode 方式編碼。    | [optional] 
**ClearanceMark** | Pointer to [**ClearanceMarkEnum**](ClearanceMarkEnum.md) |  | [optional] 
**TaxType** | Pointer to [**TaxTypeEnum**](TaxTypeEnum.md) |  | [optional] 
**CarruerType** | Pointer to [**CarruerTypeEnum**](CarruerTypeEnum.md) |  | [optional] 
**CarruerNum** | Pointer to **string** | **載具編號**   1. 當載具類別 &#x60;CarruerType&#x60;&#x3D;&#x60;&#x60;無載具)，請帶空字串。   2. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;1&#x60;(綠界科技電子發票載具)時，請帶空字串，系統會自動帶入值，為合作特店載具統一編號+自訂編號(RelateNumber)。   3. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;2&#x60;(買受人之自然人憑證)時，則請帶固定長度為16且格式 為2碼大寫英文字母加上14碼數字。   4. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;3&#x60;(買受人之手機條碼)時，則請帶固定長度為 8且格式為 1 碼斜線「/」加上由 7 碼數字及大寫英文字母及+-.符號組成。    注意事項：   1. 若手機條碼中有加號，可能在介接驗證時 發生錯誤，請將加號改為空白字元，產生 驗證碼。   2. 英文、數字、符號僅接受半形字   3. 若載具編號為手機條碼載具時，請先呼叫B2C電子發票介接技術文件手機條碼載驗證ＡＰＩ進行檢核    | [optional] 
**Donation** | Pointer to [**InvoiceDonationEunm**](InvoiceDonationEunm.md) |  | [optional] 
**LoveCode** | Pointer to **string** | **捐贈碼**   消費者選擇捐贈發票則於此欄位須填入受贈單位之捐贈碼。   1. 若捐贈註記 &#x60;Donation&#x60;&#x3D; &#x60;1&#x60; (捐贈)時，此欄位須有值。   2. 捐贈碼以阿拉伯數字為限，最少三碼，最多七碼。內容定位採「文字格式」，首位可以為零。     | [optional] 
**Print** | Pointer to [**InvoicePrintEnum**](InvoicePrintEnum.md) |  | [optional] 
**InvoiceItemName** | Pointer to **string** | **商品名稱**   預設不可為空字串且格式為 名稱 1 | 名稱 2 | 名稱 3 | … | 名稱 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。   將參數值以 UrlEncode 方式編碼。    | [optional] 
**InvoiceItemCount** | Pointer to **string** | **商品數量**   預設不可為空字串且格式為 數量 1 | 數量 2 | 數量 3 | … | 數量 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。  | [optional] 
**InvoiceItemWord** | Pointer to **string** | **商品單位**   商品單位若超過二筆以上請以「|」符號區隔單位最大長度為 6 碼。   請將參數做 UrlEncode 方式編碼。  | [optional] 
**InvoiceItemPrice** | Pointer to **string** | **商品價格**   預設不可為空字串且格式為 價格 1 | 價格 2 | 價格 3 | … | 價格 n，當含有二筆或以上的商品價格時，則以「|」符號區隔。  | [optional] 
**InvoiceItemTaxType** | Pointer to **string** | **商品課稅別**   1：應稅   2：零稅率   3：免稅   注意事項：   1. 預設為空字串，當課稅類別 [TaxType] &#x3D; 9 時，此欄位不可為空。   2. 格式為課稅 類別 1 | 課稅類別 2 | 課稅類別 3 | … | 課稅類別 n。當含有二筆或以上的商品課稅類別時，則以「|」符號區隔。   3. 課稅類別為混合稅率時，需含二筆或 以 上 的 商 品 課 稅   別[InvoiceItemTaxType]，且至少需有一筆商品課稅別為應稅及至少需有一筆商品課稅別為免稅或零稅率，即混稅發票只能 1.應稅+免稅 2.應稅+零稅率，免稅和零稅率發票不能同時開立。  | [optional] 
**InvoiceRemark** | Pointer to **string** | **備註** 當該參數有值時，請將參數值做UrlEncode 方式編碼。  | [optional] 
**DelayDay** | Pointer to **string** | **延遲天數**   本參數值請帶 0~15(天)，當天數為 0 時，則付款完成後立即開立發票。  | [optional] 
**InvType** | Pointer to **string** | **字軌類別**   若為一般稅額時，請帶 07。   預設值：07    | [optional] 

## Methods

### NewAioCheckOutRequest

`func NewAioCheckOutRequest(merchantID string, merchantTradeNo string, merchantTradeDate ECPayDateTime, paymentType AioCheckPaymentTypeEnum, totalAmount int32, tradeDesc string, itemName string, returnURL string, choosePayment ChoosePaymentEnum, checkMacValue string, encryptType EncryptTypeEnum, ) *AioCheckOutRequest`

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

`func (o *AioCheckOutRequest) GetPaymentType() AioCheckPaymentTypeEnum`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *AioCheckOutRequest) GetPaymentTypeOk() (*AioCheckPaymentTypeEnum, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *AioCheckOutRequest) SetPaymentType(v AioCheckPaymentTypeEnum)`

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

`func (o *AioCheckOutRequest) GetChoosePayment() ChoosePaymentEnum`

GetChoosePayment returns the ChoosePayment field if non-nil, zero value otherwise.

### GetChoosePaymentOk

`func (o *AioCheckOutRequest) GetChoosePaymentOk() (*ChoosePaymentEnum, bool)`

GetChoosePaymentOk returns a tuple with the ChoosePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoosePayment

`func (o *AioCheckOutRequest) SetChoosePayment(v ChoosePaymentEnum)`

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

`func (o *AioCheckOutRequest) GetNeedExtraPaidInfo() NeedExtraPaidInfoEnum`

GetNeedExtraPaidInfo returns the NeedExtraPaidInfo field if non-nil, zero value otherwise.

### GetNeedExtraPaidInfoOk

`func (o *AioCheckOutRequest) GetNeedExtraPaidInfoOk() (*NeedExtraPaidInfoEnum, bool)`

GetNeedExtraPaidInfoOk returns a tuple with the NeedExtraPaidInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedExtraPaidInfo

`func (o *AioCheckOutRequest) SetNeedExtraPaidInfo(v NeedExtraPaidInfoEnum)`

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

`func (o *AioCheckOutRequest) GetInvoiceMark() InvoiceMarkEnum`

GetInvoiceMark returns the InvoiceMark field if non-nil, zero value otherwise.

### GetInvoiceMarkOk

`func (o *AioCheckOutRequest) GetInvoiceMarkOk() (*InvoiceMarkEnum, bool)`

GetInvoiceMarkOk returns a tuple with the InvoiceMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMark

`func (o *AioCheckOutRequest) SetInvoiceMark(v InvoiceMarkEnum)`

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

`func (o *AioCheckOutRequest) GetEncryptType() EncryptTypeEnum`

GetEncryptType returns the EncryptType field if non-nil, zero value otherwise.

### GetEncryptTypeOk

`func (o *AioCheckOutRequest) GetEncryptTypeOk() (*EncryptTypeEnum, bool)`

GetEncryptTypeOk returns a tuple with the EncryptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptType

`func (o *AioCheckOutRequest) SetEncryptType(v EncryptTypeEnum)`

SetEncryptType sets EncryptType field to given value.


### GetLanguage

`func (o *AioCheckOutRequest) GetLanguage() LanguageEnum`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AioCheckOutRequest) GetLanguageOk() (*LanguageEnum, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AioCheckOutRequest) SetLanguage(v LanguageEnum)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AioCheckOutRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetStoreExpireDate

`func (o *AioCheckOutRequest) GetStoreExpireDate() int32`

GetStoreExpireDate returns the StoreExpireDate field if non-nil, zero value otherwise.

### GetStoreExpireDateOk

`func (o *AioCheckOutRequest) GetStoreExpireDateOk() (*int32, bool)`

GetStoreExpireDateOk returns a tuple with the StoreExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreExpireDate

`func (o *AioCheckOutRequest) SetStoreExpireDate(v int32)`

SetStoreExpireDate sets StoreExpireDate field to given value.

### HasStoreExpireDate

`func (o *AioCheckOutRequest) HasStoreExpireDate() bool`

HasStoreExpireDate returns a boolean if a field has been set.

### GetDesc1

`func (o *AioCheckOutRequest) GetDesc1() string`

GetDesc1 returns the Desc1 field if non-nil, zero value otherwise.

### GetDesc1Ok

`func (o *AioCheckOutRequest) GetDesc1Ok() (*string, bool)`

GetDesc1Ok returns a tuple with the Desc1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc1

`func (o *AioCheckOutRequest) SetDesc1(v string)`

SetDesc1 sets Desc1 field to given value.

### HasDesc1

`func (o *AioCheckOutRequest) HasDesc1() bool`

HasDesc1 returns a boolean if a field has been set.

### GetDesc2

`func (o *AioCheckOutRequest) GetDesc2() string`

GetDesc2 returns the Desc2 field if non-nil, zero value otherwise.

### GetDesc2Ok

`func (o *AioCheckOutRequest) GetDesc2Ok() (*string, bool)`

GetDesc2Ok returns a tuple with the Desc2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc2

`func (o *AioCheckOutRequest) SetDesc2(v string)`

SetDesc2 sets Desc2 field to given value.

### HasDesc2

`func (o *AioCheckOutRequest) HasDesc2() bool`

HasDesc2 returns a boolean if a field has been set.

### GetDesc3

`func (o *AioCheckOutRequest) GetDesc3() string`

GetDesc3 returns the Desc3 field if non-nil, zero value otherwise.

### GetDesc3Ok

`func (o *AioCheckOutRequest) GetDesc3Ok() (*string, bool)`

GetDesc3Ok returns a tuple with the Desc3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc3

`func (o *AioCheckOutRequest) SetDesc3(v string)`

SetDesc3 sets Desc3 field to given value.

### HasDesc3

`func (o *AioCheckOutRequest) HasDesc3() bool`

HasDesc3 returns a boolean if a field has been set.

### GetDesc4

`func (o *AioCheckOutRequest) GetDesc4() string`

GetDesc4 returns the Desc4 field if non-nil, zero value otherwise.

### GetDesc4Ok

`func (o *AioCheckOutRequest) GetDesc4Ok() (*string, bool)`

GetDesc4Ok returns a tuple with the Desc4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc4

`func (o *AioCheckOutRequest) SetDesc4(v string)`

SetDesc4 sets Desc4 field to given value.

### HasDesc4

`func (o *AioCheckOutRequest) HasDesc4() bool`

HasDesc4 returns a boolean if a field has been set.

### GetPaymentInfoURL

`func (o *AioCheckOutRequest) GetPaymentInfoURL() string`

GetPaymentInfoURL returns the PaymentInfoURL field if non-nil, zero value otherwise.

### GetPaymentInfoURLOk

`func (o *AioCheckOutRequest) GetPaymentInfoURLOk() (*string, bool)`

GetPaymentInfoURLOk returns a tuple with the PaymentInfoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoURL

`func (o *AioCheckOutRequest) SetPaymentInfoURL(v string)`

SetPaymentInfoURL sets PaymentInfoURL field to given value.

### HasPaymentInfoURL

`func (o *AioCheckOutRequest) HasPaymentInfoURL() bool`

HasPaymentInfoURL returns a boolean if a field has been set.

### GetClientRedirectURL

`func (o *AioCheckOutRequest) GetClientRedirectURL() string`

GetClientRedirectURL returns the ClientRedirectURL field if non-nil, zero value otherwise.

### GetClientRedirectURLOk

`func (o *AioCheckOutRequest) GetClientRedirectURLOk() (*string, bool)`

GetClientRedirectURLOk returns a tuple with the ClientRedirectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRedirectURL

`func (o *AioCheckOutRequest) SetClientRedirectURL(v string)`

SetClientRedirectURL sets ClientRedirectURL field to given value.

### HasClientRedirectURL

`func (o *AioCheckOutRequest) HasClientRedirectURL() bool`

HasClientRedirectURL returns a boolean if a field has been set.

### GetBindingCard

`func (o *AioCheckOutRequest) GetBindingCard() BindingCardEnum`

GetBindingCard returns the BindingCard field if non-nil, zero value otherwise.

### GetBindingCardOk

`func (o *AioCheckOutRequest) GetBindingCardOk() (*BindingCardEnum, bool)`

GetBindingCardOk returns a tuple with the BindingCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingCard

`func (o *AioCheckOutRequest) SetBindingCard(v BindingCardEnum)`

SetBindingCard sets BindingCard field to given value.

### HasBindingCard

`func (o *AioCheckOutRequest) HasBindingCard() bool`

HasBindingCard returns a boolean if a field has been set.

### GetMerchantMemberID

`func (o *AioCheckOutRequest) GetMerchantMemberID() string`

GetMerchantMemberID returns the MerchantMemberID field if non-nil, zero value otherwise.

### GetMerchantMemberIDOk

`func (o *AioCheckOutRequest) GetMerchantMemberIDOk() (*string, bool)`

GetMerchantMemberIDOk returns a tuple with the MerchantMemberID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMemberID

`func (o *AioCheckOutRequest) SetMerchantMemberID(v string)`

SetMerchantMemberID sets MerchantMemberID field to given value.

### HasMerchantMemberID

`func (o *AioCheckOutRequest) HasMerchantMemberID() bool`

HasMerchantMemberID returns a boolean if a field has been set.

### GetRedeem

`func (o *AioCheckOutRequest) GetRedeem() RedeemEnum`

GetRedeem returns the Redeem field if non-nil, zero value otherwise.

### GetRedeemOk

`func (o *AioCheckOutRequest) GetRedeemOk() (*RedeemEnum, bool)`

GetRedeemOk returns a tuple with the Redeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeem

`func (o *AioCheckOutRequest) SetRedeem(v RedeemEnum)`

SetRedeem sets Redeem field to given value.

### HasRedeem

`func (o *AioCheckOutRequest) HasRedeem() bool`

HasRedeem returns a boolean if a field has been set.

### GetUnionPay

`func (o *AioCheckOutRequest) GetUnionPay() UnionPayEnum`

GetUnionPay returns the UnionPay field if non-nil, zero value otherwise.

### GetUnionPayOk

`func (o *AioCheckOutRequest) GetUnionPayOk() (*UnionPayEnum, bool)`

GetUnionPayOk returns a tuple with the UnionPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionPay

`func (o *AioCheckOutRequest) SetUnionPay(v UnionPayEnum)`

SetUnionPay sets UnionPay field to given value.

### HasUnionPay

`func (o *AioCheckOutRequest) HasUnionPay() bool`

HasUnionPay returns a boolean if a field has been set.

### GetCreditInstallment

`func (o *AioCheckOutRequest) GetCreditInstallment() string`

GetCreditInstallment returns the CreditInstallment field if non-nil, zero value otherwise.

### GetCreditInstallmentOk

`func (o *AioCheckOutRequest) GetCreditInstallmentOk() (*string, bool)`

GetCreditInstallmentOk returns a tuple with the CreditInstallment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditInstallment

`func (o *AioCheckOutRequest) SetCreditInstallment(v string)`

SetCreditInstallment sets CreditInstallment field to given value.

### HasCreditInstallment

`func (o *AioCheckOutRequest) HasCreditInstallment() bool`

HasCreditInstallment returns a boolean if a field has been set.

### GetPeriodAmount

`func (o *AioCheckOutRequest) GetPeriodAmount() int32`

GetPeriodAmount returns the PeriodAmount field if non-nil, zero value otherwise.

### GetPeriodAmountOk

`func (o *AioCheckOutRequest) GetPeriodAmountOk() (*int32, bool)`

GetPeriodAmountOk returns a tuple with the PeriodAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAmount

`func (o *AioCheckOutRequest) SetPeriodAmount(v int32)`

SetPeriodAmount sets PeriodAmount field to given value.

### HasPeriodAmount

`func (o *AioCheckOutRequest) HasPeriodAmount() bool`

HasPeriodAmount returns a boolean if a field has been set.

### GetPeriodType

`func (o *AioCheckOutRequest) GetPeriodType() CreditPeriodTypeEnum`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *AioCheckOutRequest) GetPeriodTypeOk() (*CreditPeriodTypeEnum, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *AioCheckOutRequest) SetPeriodType(v CreditPeriodTypeEnum)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *AioCheckOutRequest) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### GetFrequency

`func (o *AioCheckOutRequest) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AioCheckOutRequest) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AioCheckOutRequest) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *AioCheckOutRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetExecTimes

`func (o *AioCheckOutRequest) GetExecTimes() int32`

GetExecTimes returns the ExecTimes field if non-nil, zero value otherwise.

### GetExecTimesOk

`func (o *AioCheckOutRequest) GetExecTimesOk() (*int32, bool)`

GetExecTimesOk returns a tuple with the ExecTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecTimes

`func (o *AioCheckOutRequest) SetExecTimes(v int32)`

SetExecTimes sets ExecTimes field to given value.

### HasExecTimes

`func (o *AioCheckOutRequest) HasExecTimes() bool`

HasExecTimes returns a boolean if a field has been set.

### GetPeriodReturnURL

`func (o *AioCheckOutRequest) GetPeriodReturnURL() string`

GetPeriodReturnURL returns the PeriodReturnURL field if non-nil, zero value otherwise.

### GetPeriodReturnURLOk

`func (o *AioCheckOutRequest) GetPeriodReturnURLOk() (*string, bool)`

GetPeriodReturnURLOk returns a tuple with the PeriodReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodReturnURL

`func (o *AioCheckOutRequest) SetPeriodReturnURL(v string)`

SetPeriodReturnURL sets PeriodReturnURL field to given value.

### HasPeriodReturnURL

`func (o *AioCheckOutRequest) HasPeriodReturnURL() bool`

HasPeriodReturnURL returns a boolean if a field has been set.

### GetRelateNumber

`func (o *AioCheckOutRequest) GetRelateNumber() string`

GetRelateNumber returns the RelateNumber field if non-nil, zero value otherwise.

### GetRelateNumberOk

`func (o *AioCheckOutRequest) GetRelateNumberOk() (*string, bool)`

GetRelateNumberOk returns a tuple with the RelateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelateNumber

`func (o *AioCheckOutRequest) SetRelateNumber(v string)`

SetRelateNumber sets RelateNumber field to given value.

### HasRelateNumber

`func (o *AioCheckOutRequest) HasRelateNumber() bool`

HasRelateNumber returns a boolean if a field has been set.

### GetCustomerID

`func (o *AioCheckOutRequest) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *AioCheckOutRequest) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *AioCheckOutRequest) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *AioCheckOutRequest) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetCustomerIdentifier

`func (o *AioCheckOutRequest) GetCustomerIdentifier() string`

GetCustomerIdentifier returns the CustomerIdentifier field if non-nil, zero value otherwise.

### GetCustomerIdentifierOk

`func (o *AioCheckOutRequest) GetCustomerIdentifierOk() (*string, bool)`

GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdentifier

`func (o *AioCheckOutRequest) SetCustomerIdentifier(v string)`

SetCustomerIdentifier sets CustomerIdentifier field to given value.

### HasCustomerIdentifier

`func (o *AioCheckOutRequest) HasCustomerIdentifier() bool`

HasCustomerIdentifier returns a boolean if a field has been set.

### GetCustomerName

`func (o *AioCheckOutRequest) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *AioCheckOutRequest) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *AioCheckOutRequest) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *AioCheckOutRequest) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerAddr

`func (o *AioCheckOutRequest) GetCustomerAddr() string`

GetCustomerAddr returns the CustomerAddr field if non-nil, zero value otherwise.

### GetCustomerAddrOk

`func (o *AioCheckOutRequest) GetCustomerAddrOk() (*string, bool)`

GetCustomerAddrOk returns a tuple with the CustomerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAddr

`func (o *AioCheckOutRequest) SetCustomerAddr(v string)`

SetCustomerAddr sets CustomerAddr field to given value.

### HasCustomerAddr

`func (o *AioCheckOutRequest) HasCustomerAddr() bool`

HasCustomerAddr returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *AioCheckOutRequest) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *AioCheckOutRequest) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *AioCheckOutRequest) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *AioCheckOutRequest) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *AioCheckOutRequest) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *AioCheckOutRequest) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *AioCheckOutRequest) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *AioCheckOutRequest) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetClearanceMark

`func (o *AioCheckOutRequest) GetClearanceMark() ClearanceMarkEnum`

GetClearanceMark returns the ClearanceMark field if non-nil, zero value otherwise.

### GetClearanceMarkOk

`func (o *AioCheckOutRequest) GetClearanceMarkOk() (*ClearanceMarkEnum, bool)`

GetClearanceMarkOk returns a tuple with the ClearanceMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearanceMark

`func (o *AioCheckOutRequest) SetClearanceMark(v ClearanceMarkEnum)`

SetClearanceMark sets ClearanceMark field to given value.

### HasClearanceMark

`func (o *AioCheckOutRequest) HasClearanceMark() bool`

HasClearanceMark returns a boolean if a field has been set.

### GetTaxType

`func (o *AioCheckOutRequest) GetTaxType() TaxTypeEnum`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *AioCheckOutRequest) GetTaxTypeOk() (*TaxTypeEnum, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *AioCheckOutRequest) SetTaxType(v TaxTypeEnum)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *AioCheckOutRequest) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetCarruerType

`func (o *AioCheckOutRequest) GetCarruerType() CarruerTypeEnum`

GetCarruerType returns the CarruerType field if non-nil, zero value otherwise.

### GetCarruerTypeOk

`func (o *AioCheckOutRequest) GetCarruerTypeOk() (*CarruerTypeEnum, bool)`

GetCarruerTypeOk returns a tuple with the CarruerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarruerType

`func (o *AioCheckOutRequest) SetCarruerType(v CarruerTypeEnum)`

SetCarruerType sets CarruerType field to given value.

### HasCarruerType

`func (o *AioCheckOutRequest) HasCarruerType() bool`

HasCarruerType returns a boolean if a field has been set.

### GetCarruerNum

`func (o *AioCheckOutRequest) GetCarruerNum() string`

GetCarruerNum returns the CarruerNum field if non-nil, zero value otherwise.

### GetCarruerNumOk

`func (o *AioCheckOutRequest) GetCarruerNumOk() (*string, bool)`

GetCarruerNumOk returns a tuple with the CarruerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarruerNum

`func (o *AioCheckOutRequest) SetCarruerNum(v string)`

SetCarruerNum sets CarruerNum field to given value.

### HasCarruerNum

`func (o *AioCheckOutRequest) HasCarruerNum() bool`

HasCarruerNum returns a boolean if a field has been set.

### GetDonation

`func (o *AioCheckOutRequest) GetDonation() InvoiceDonationEunm`

GetDonation returns the Donation field if non-nil, zero value otherwise.

### GetDonationOk

`func (o *AioCheckOutRequest) GetDonationOk() (*InvoiceDonationEunm, bool)`

GetDonationOk returns a tuple with the Donation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonation

`func (o *AioCheckOutRequest) SetDonation(v InvoiceDonationEunm)`

SetDonation sets Donation field to given value.

### HasDonation

`func (o *AioCheckOutRequest) HasDonation() bool`

HasDonation returns a boolean if a field has been set.

### GetLoveCode

`func (o *AioCheckOutRequest) GetLoveCode() string`

GetLoveCode returns the LoveCode field if non-nil, zero value otherwise.

### GetLoveCodeOk

`func (o *AioCheckOutRequest) GetLoveCodeOk() (*string, bool)`

GetLoveCodeOk returns a tuple with the LoveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoveCode

`func (o *AioCheckOutRequest) SetLoveCode(v string)`

SetLoveCode sets LoveCode field to given value.

### HasLoveCode

`func (o *AioCheckOutRequest) HasLoveCode() bool`

HasLoveCode returns a boolean if a field has been set.

### GetPrint

`func (o *AioCheckOutRequest) GetPrint() InvoicePrintEnum`

GetPrint returns the Print field if non-nil, zero value otherwise.

### GetPrintOk

`func (o *AioCheckOutRequest) GetPrintOk() (*InvoicePrintEnum, bool)`

GetPrintOk returns a tuple with the Print field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrint

`func (o *AioCheckOutRequest) SetPrint(v InvoicePrintEnum)`

SetPrint sets Print field to given value.

### HasPrint

`func (o *AioCheckOutRequest) HasPrint() bool`

HasPrint returns a boolean if a field has been set.

### GetInvoiceItemName

`func (o *AioCheckOutRequest) GetInvoiceItemName() string`

GetInvoiceItemName returns the InvoiceItemName field if non-nil, zero value otherwise.

### GetInvoiceItemNameOk

`func (o *AioCheckOutRequest) GetInvoiceItemNameOk() (*string, bool)`

GetInvoiceItemNameOk returns a tuple with the InvoiceItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemName

`func (o *AioCheckOutRequest) SetInvoiceItemName(v string)`

SetInvoiceItemName sets InvoiceItemName field to given value.

### HasInvoiceItemName

`func (o *AioCheckOutRequest) HasInvoiceItemName() bool`

HasInvoiceItemName returns a boolean if a field has been set.

### GetInvoiceItemCount

`func (o *AioCheckOutRequest) GetInvoiceItemCount() string`

GetInvoiceItemCount returns the InvoiceItemCount field if non-nil, zero value otherwise.

### GetInvoiceItemCountOk

`func (o *AioCheckOutRequest) GetInvoiceItemCountOk() (*string, bool)`

GetInvoiceItemCountOk returns a tuple with the InvoiceItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemCount

`func (o *AioCheckOutRequest) SetInvoiceItemCount(v string)`

SetInvoiceItemCount sets InvoiceItemCount field to given value.

### HasInvoiceItemCount

`func (o *AioCheckOutRequest) HasInvoiceItemCount() bool`

HasInvoiceItemCount returns a boolean if a field has been set.

### GetInvoiceItemWord

`func (o *AioCheckOutRequest) GetInvoiceItemWord() string`

GetInvoiceItemWord returns the InvoiceItemWord field if non-nil, zero value otherwise.

### GetInvoiceItemWordOk

`func (o *AioCheckOutRequest) GetInvoiceItemWordOk() (*string, bool)`

GetInvoiceItemWordOk returns a tuple with the InvoiceItemWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemWord

`func (o *AioCheckOutRequest) SetInvoiceItemWord(v string)`

SetInvoiceItemWord sets InvoiceItemWord field to given value.

### HasInvoiceItemWord

`func (o *AioCheckOutRequest) HasInvoiceItemWord() bool`

HasInvoiceItemWord returns a boolean if a field has been set.

### GetInvoiceItemPrice

`func (o *AioCheckOutRequest) GetInvoiceItemPrice() string`

GetInvoiceItemPrice returns the InvoiceItemPrice field if non-nil, zero value otherwise.

### GetInvoiceItemPriceOk

`func (o *AioCheckOutRequest) GetInvoiceItemPriceOk() (*string, bool)`

GetInvoiceItemPriceOk returns a tuple with the InvoiceItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemPrice

`func (o *AioCheckOutRequest) SetInvoiceItemPrice(v string)`

SetInvoiceItemPrice sets InvoiceItemPrice field to given value.

### HasInvoiceItemPrice

`func (o *AioCheckOutRequest) HasInvoiceItemPrice() bool`

HasInvoiceItemPrice returns a boolean if a field has been set.

### GetInvoiceItemTaxType

`func (o *AioCheckOutRequest) GetInvoiceItemTaxType() string`

GetInvoiceItemTaxType returns the InvoiceItemTaxType field if non-nil, zero value otherwise.

### GetInvoiceItemTaxTypeOk

`func (o *AioCheckOutRequest) GetInvoiceItemTaxTypeOk() (*string, bool)`

GetInvoiceItemTaxTypeOk returns a tuple with the InvoiceItemTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemTaxType

`func (o *AioCheckOutRequest) SetInvoiceItemTaxType(v string)`

SetInvoiceItemTaxType sets InvoiceItemTaxType field to given value.

### HasInvoiceItemTaxType

`func (o *AioCheckOutRequest) HasInvoiceItemTaxType() bool`

HasInvoiceItemTaxType returns a boolean if a field has been set.

### GetInvoiceRemark

`func (o *AioCheckOutRequest) GetInvoiceRemark() string`

GetInvoiceRemark returns the InvoiceRemark field if non-nil, zero value otherwise.

### GetInvoiceRemarkOk

`func (o *AioCheckOutRequest) GetInvoiceRemarkOk() (*string, bool)`

GetInvoiceRemarkOk returns a tuple with the InvoiceRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceRemark

`func (o *AioCheckOutRequest) SetInvoiceRemark(v string)`

SetInvoiceRemark sets InvoiceRemark field to given value.

### HasInvoiceRemark

`func (o *AioCheckOutRequest) HasInvoiceRemark() bool`

HasInvoiceRemark returns a boolean if a field has been set.

### GetDelayDay

`func (o *AioCheckOutRequest) GetDelayDay() string`

GetDelayDay returns the DelayDay field if non-nil, zero value otherwise.

### GetDelayDayOk

`func (o *AioCheckOutRequest) GetDelayDayOk() (*string, bool)`

GetDelayDayOk returns a tuple with the DelayDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDay

`func (o *AioCheckOutRequest) SetDelayDay(v string)`

SetDelayDay sets DelayDay field to given value.

### HasDelayDay

`func (o *AioCheckOutRequest) HasDelayDay() bool`

HasDelayDay returns a boolean if a field has been set.

### GetInvType

`func (o *AioCheckOutRequest) GetInvType() string`

GetInvType returns the InvType field if non-nil, zero value otherwise.

### GetInvTypeOk

`func (o *AioCheckOutRequest) GetInvTypeOk() (*string, bool)`

GetInvTypeOk returns a tuple with the InvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvType

`func (o *AioCheckOutRequest) SetInvType(v string)`

SetInvType sets InvType field to given value.

### HasInvType

`func (o *AioCheckOutRequest) HasInvType() bool`

HasInvType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


