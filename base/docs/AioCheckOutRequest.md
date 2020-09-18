# AioCheckOutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | **string** | 特店編號(由綠界提供) | 
**MerchantTradeNo** | **string** | **特店交易編號(由特店提供)** 特店交易編號均為唯一值，不可重複使用。 英數字大小寫混合 如何避免訂單編號重複請參考 FAQ 如有使用 &#x60;PlatformID&#x60; ，平台商底下所有商家之訂單編號亦不可重複。  | 
**StoreID** | **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | [optional] 
**MerchantTradeDate** | [**ECPayDateTime**](ECPayDateTime.md) | **特店交易時間** 格式為 &#x60;yyyy/MM/dd HH:mm:ss&#x60;  | 
**PaymentType** | **string** | **交易類型** 請固定填入 &#x60;aio&#x60;  | [default to PAYMENT_TYPE_AIO]
**TotalAmount** | **int32** | **交易金額** 請帶整數，不可有小數點。 僅限新台幣。 各付款金額的限制，請參考 &lt;https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID&#x3D;3605&gt;  | 
**TradeDesc** | **string** | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。  | 
**ItemName** | **string** | **商品名稱** 1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。 2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。  | 
**ReturnURL** | **string** | **付款完成通知回傳網址** 當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。 2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。  | 
**ChoosePayment** | **string** | **選擇預設付款方式** 綠界提供下列付款方式，請於建立訂單時傳送過來: - &#x60;Credit&#x60;: 信用卡及銀聯卡(需申請開通) - &#x60;WebATM&#x60;: 網路 ATM - &#x60;ATM&#x60;: 自動櫃員機 - &#x60;CVS&#x60;: 超商代碼 - &#x60;BARCODE&#x60;: 超商條碼 - &#x60;ALL&#x60;: 不指定付款方式，由綠界顯示付款方式選擇頁面。  注意事項： 1.若為手機版時不支援下列付款方式:   - WebATM:網路 ATM  2.如需要不透過綠界畫面取得 &#x60;ATM&#x60;、&#x60;CVS&#x60;、&#x60;BARCODE&#x60; 的繳費代碼，請參考 FAQ。  | 
**CheckMacValue** | **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 
**ClientBackURL** | **string** | **Client端返回特店的按鈕連結** 消費者點選此按鈕後，會將頁面導回到此設定的網址 注意事項： 導回時不會帶付款結果到此網址，只是將頁面導回而已。 設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。 設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。 若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。 若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。  | [optional] 
**ItemURL** | **string** | **商品銷售網址**  | [optional] 
**Remark** | **string** | **備註欄位**  | [optional] 
**ChooseSubPayment** | **string** | **付款子項目** 若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。 請參考付款方式一覽表  | [optional] 
**OrderResultURL** | **string** | **Client端回傳付款結果網址** 當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 若與[ClientBackURL]同時設定，將會以此參數為主。 2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。  | [optional] 
**NeedExtraPaidInfo** | **string** | **是否需要額外的付款資訊** 額外的付款資訊： 若不回傳額外的付款資訊時，參數值請傳：&#x60;Ｎ&#x60;； 若要回傳額外的付款資訊時，參數值請傳：&#x60;Ｙ&#x60;，付款完成後綠界會以 Server POST 方式回傳額外付款資訊。 注意事項： 回傳額外付款資訊參數請參考-額外回傳的參數  | [optional] [default to NEED_EXTRA_PAID_INFO_N]
**DeviceSource** | **string** | **裝置來源** 請帶空值，由系統自動判定。  | [optional] 
**IgnorePayment** | **string** | **隱藏付款** 當付款方式 &#x60;ChoosePayment&#x60; 為 &#x60;ALL&#x60; 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。 可用的參數值： - &#x60;Credit&#x60;: 信用卡 - &#x60;WebATM&#x60;: 網路 ATM - &#x60;ATM&#x60;: 自動櫃員機 - &#x60;CVS&#x60;: 超商代碼 - &#x60;BARCODE&#x60;: 超商條碼  | [optional] 
**PlatformID** | **string** | **特約合作平台商代號(由綠界提供)** 為專案合作的平台商使用。 一般特店或平台商本身介接，則參數請帶放空值。 若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 &#x60;MerchantID&#x60;。  | [optional] 
**InvoiceMark** | **string** | **電子發票開立註記** 此參數為付款完成後同時開立電子發票。 若要使用時，該參數須設定為「Y」，同時還要設定「電子發票介接相關參數」 注意事項： 正式環境欲使用電子發票功能，須與綠界申請開通，若未開通請致電客服中心 &#x60;(02) 2655-1775&#x60;。  | [optional] 
**CustomField1** | **string** | **自訂名稱欄位1** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField2** | **string** | **自訂名稱欄位2** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField3** | **string** | **自訂名稱欄位3** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**CustomField4** | **string** | **自訂名稱欄位4** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | [optional] 
**EncryptType** | **int32** | **CheckMacValue加密類型** 請固定填入 &#x60;1&#x60;，使用 SHA256 加密。  | [default to ENCRYPT_TYPE__1]
**Language** | **string** | **語系設定** 預設語系為中文，若要變更語系參數值請帶： - 英語：&#x60;ENG&#x60; - 韓語：&#x60;KOR&#x60; - 日語：&#x60;JPN&#x60; - 簡體中文：&#x60;CHI&#x60;  | [optional] 
**StoreExpireDate** | **int32** | **超商繳費截止時間** 注意事項： &#x60;CVS&#x60;:以分鐘為單位 &#x60;BARCODE&#x60;:以天為單位 若未設定此參數，&#x60;CVS&#x60; 預設為 &#x60;10080&#x60; 分鐘(&#x60;7&#x60; 天)；BARCODE 預設為 &#x60;7&#x60; 天。 若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 &#x60;86400&#x60; 分鐘，超過時一律以 &#x60;86400&#x60; 分鐘計(&#x60;60&#x60; 天) 例：&#x60;08/01&#x60; 的 &#x60;20:15&#x60; 分購買商品，繳費期限為 &#x60;7&#x60; 天，表示 &#x60;8/08&#x60; 的 &#x60;20:15&#x60; 分前您必須前往超商繳費。 範例: &#x60;CVS&#x60;&#x3D;&#x60;1440&#x60;(共 &#x60;1&#x60; 天)、&#x60;BARCODE&#x60;&#x3D;&#x60;7&#x60;(共 &#x60;7&#x60; 天)  | [optional] 
**Desc1** | **string** | **交易描述1** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc2** | **string** | **交易描述2** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc3** | **string** | **交易描述3** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc4** | **string** | **交易描述4** 會出現在超商繳費平台螢幕上  | [optional] 
**PaymentInfoURL** | **string** | **Server 端回傳付款相關資訊** 若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。 請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.] 注意事項： 頁面將會停留在綠界，顯示繳費的相關資訊。 回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。  | [optional] 
**ClientRedirectURL** | **string** | **Client端回傳付款方式相關資訊** 若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。 請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.] 注意事項： 若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。 若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。 回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。  | [optional] 
**BindingCard** | **int32** | **記憶卡號** 使用記憶信用卡 使用：請傳 &#x60;1&#x60; 不使用：請傳 &#x60;0&#x60;  | [optional] 
**MerchantMemberID** | **string** | **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 &#x60;MerchantID&#x60; + &#x60;廠商會員編號&#x60;)  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


