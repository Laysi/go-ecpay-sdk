# \ECPayApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AioCheckOut**](ECPayApi.md#AioCheckOut) | **Post** /Cashier/AioCheckOut/V5 | 



## AioCheckOut

> AioCheckOut(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).MerchantTradeDate(merchantTradeDate).PaymentType(paymentType).TotalAmount(totalAmount).TradeDesc(tradeDesc).ItemName(itemName).ReturnURL(returnURL).ChoosePayment(choosePayment).CheckMacValue(checkMacValue).StoreID(storeID).ClientBackURL(clientBackURL).ItemURL(itemURL).Remark(remark).ChooseSubPayment(chooseSubPayment).OrderResultURL(orderResultURL).NeedExtraPaidInfo(needExtraPaidInfo).DeviceSource(deviceSource).IgnorePayment(ignorePayment).PlatformID(platformID).InvoiceMark(invoiceMark).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).EncryptType(encryptType).Language(language).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    merchantID := "merchantID_example" // string | 特店編號(由綠界提供)
    merchantTradeNo := "merchantTradeNo_example" // string | **特店交易編號(由特店提供)** 特店交易編號均為唯一值，不可重複使用。 英數字大小寫混合 如何避免訂單編號重複請參考 FAQ 如有使用 `PlatformID` ，平台商底下所有商家之訂單編號亦不可重複。 
    merchantTradeDate := "merchantTradeDate_example" // ECPayDateTime | **特店交易時間** 格式為 `yyyy/MM/dd HH:mm:ss` 
    paymentType := "paymentType_example" // string | **交易類型** 請固定填入 `aio`  (default to "aio")
    totalAmount := 987 // int32 | **交易金額** 請帶整數，不可有小數點。 僅限新台幣。 各付款金額的限制，請參考 <https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID=3605> 
    tradeDesc := "tradeDesc_example" // string | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。 
    itemName := "itemName_example" // string | **商品名稱** 1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。 2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。 
    returnURL := "returnURL_example" // string | **付款完成通知回傳網址** 當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。 2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。 
    choosePayment := "choosePayment_example" // string | **選擇預設付款方式** 綠界提供下列付款方式，請於建立訂單時傳送過來: - `Credit`: 信用卡及銀聯卡(需申請開通) - `WebATM`: 網路 ATM - `ATM`: 自動櫃員機 - `CVS`: 超商代碼 - `BARCODE`: 超商條碼 - `ALL`: 不指定付款方式，由綠界顯示付款方式選擇頁面。  注意事項： 1.若為手機版時不支援下列付款方式:   - WebATM:網路 ATM  2.如需要不透過綠界畫面取得 `ATM`、`CVS`、`BARCODE` 的繳費代碼，請參考 FAQ。 
    checkMacValue := "checkMacValue_example" // string | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式 
    storeID := "storeID_example" // string | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  (optional)
    clientBackURL := "clientBackURL_example" // string | **Client 端返回特店的按鈕連結** 消費者點選此按鈕後，會將頁面導回到此設定的網址 注意事項： 導回時不會帶付款結果到此網址，只是將頁面導回而已。 設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。 設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。 若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。 若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。  (optional)
    itemURL := "itemURL_example" // string | **商品銷售網址**  (optional)
    remark := "remark_example" // string | **備註欄位**  (optional)
    chooseSubPayment := "chooseSubPayment_example" // string | **付款子項目** 若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。 請參考付款方式一覽表  (optional)
    orderResultURL := "orderResultURL_example" // string | **Client端回傳付款結果網址** 當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 若與[ClientBackURL]同時設定，將會以此參數為主。 2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。  (optional)
    needExtraPaidInfo := "needExtraPaidInfo_example" // string | **是否需要額外的付款資訊** 額外的付款資訊： 若不回傳額外的付款資訊時，參數值請傳：`Ｎ`； 若要回傳額外的付款資訊時，參數值請傳：`Ｙ`，付款完成後綠界會以 Server POST 方式回傳額外付款資訊。 注意事項： 回傳額外付款資訊參數請參考-額外回傳的參數  (optional) (default to "N")
    deviceSource := "deviceSource_example" // string | **裝置來源** 請帶空值，由系統自動判定。  (optional)
    ignorePayment := "ignorePayment_example" // string | **隱藏付款** 當付款方式 `ChoosePayment` 為 `ALL` 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。 可用的參數值： - `Credit`: 信用卡 - `WebATM`: 網路 ATM - `ATM`: 自動櫃員機 - `CVS`: 超商代碼 - `BARCODE`: 超商條碼  (optional)
    platformID := "platformID_example" // string | **特約合作平台商代號(由綠界提供)** 為專案合作的平台商使用。 一般特店或平台商本身介接，則參數請帶放空值。 若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 `MerchantID`。  (optional)
    invoiceMark := "invoiceMark_example" // string | **電子發票開立註記** 此參數為付款完成後同時開立電子發票。 若要使用時，該參數須設定為「Y」，同時還要設定「電子發票介接相關參數」 注意事項： 正式環境欲使用電子發票功能，須與綠界申請開通，若未開通請致電客服中心 `(02) 2655-1775`。  (optional)
    customField1 := "customField1_example" // string | **自訂名稱欄位1** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    customField2 := "customField2_example" // string | **自訂名稱欄位2** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    customField3 := "customField3_example" // string | **自訂名稱欄位3** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    customField4 := "customField4_example" // string | **自訂名稱欄位4** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    encryptType := 987 // int32 | **CheckMacValue加密類型** 請固定填入 1，使用 SHA256 加密。  (optional) (default to 1)
    language := "language_example" // string | **語系設定** 預設語系為中文，若要變更語系參數值請帶： - 英語：`ENG` - 韓語：`KOR` - 日語：`JPN` - 簡體中文：`CHI`  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ECPayApi.AioCheckOut(context.Background(), merchantID, merchantTradeNo, merchantTradeDate, paymentType, totalAmount, tradeDesc, itemName, returnURL, choosePayment, checkMacValue).StoreID(storeID).ClientBackURL(clientBackURL).ItemURL(itemURL).Remark(remark).ChooseSubPayment(chooseSubPayment).OrderResultURL(orderResultURL).NeedExtraPaidInfo(needExtraPaidInfo).DeviceSource(deviceSource).IgnorePayment(ignorePayment).PlatformID(platformID).InvoiceMark(invoiceMark).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).EncryptType(encryptType).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ECPayApi.AioCheckOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAioCheckOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantID** | **string** | 特店編號(由綠界提供) | 
 **merchantTradeNo** | **string** | **特店交易編號(由特店提供)** 特店交易編號均為唯一值，不可重複使用。 英數字大小寫混合 如何避免訂單編號重複請參考 FAQ 如有使用 &#x60;PlatformID&#x60; ，平台商底下所有商家之訂單編號亦不可重複。  | 
 **merchantTradeDate** | **ECPayDateTime** | **特店交易時間** 格式為 &#x60;yyyy/MM/dd HH:mm:ss&#x60;  | 
 **paymentType** | **string** | **交易類型** 請固定填入 &#x60;aio&#x60;  | [default to &quot;aio&quot;]
 **totalAmount** | **int32** | **交易金額** 請帶整數，不可有小數點。 僅限新台幣。 各付款金額的限制，請參考 &lt;https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID&#x3D;3605&gt;  | 
 **tradeDesc** | **string** | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。  | 
 **itemName** | **string** | **商品名稱** 1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。 2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。  | 
 **returnURL** | **string** | **付款完成通知回傳網址** 當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。 2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。  | 
 **choosePayment** | **string** | **選擇預設付款方式** 綠界提供下列付款方式，請於建立訂單時傳送過來: - &#x60;Credit&#x60;: 信用卡及銀聯卡(需申請開通) - &#x60;WebATM&#x60;: 網路 ATM - &#x60;ATM&#x60;: 自動櫃員機 - &#x60;CVS&#x60;: 超商代碼 - &#x60;BARCODE&#x60;: 超商條碼 - &#x60;ALL&#x60;: 不指定付款方式，由綠界顯示付款方式選擇頁面。  注意事項： 1.若為手機版時不支援下列付款方式:   - WebATM:網路 ATM  2.如需要不透過綠界畫面取得 &#x60;ATM&#x60;、&#x60;CVS&#x60;、&#x60;BARCODE&#x60; 的繳費代碼，請參考 FAQ。  | 
 **checkMacValue** | **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 
 **storeID** | **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
 **clientBackURL** | **string** | **Client 端返回特店的按鈕連結** 消費者點選此按鈕後，會將頁面導回到此設定的網址 注意事項： 導回時不會帶付款結果到此網址，只是將頁面導回而已。 設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。 設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。 若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。 若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。  | 
 **itemURL** | **string** | **商品銷售網址**  | 
 **remark** | **string** | **備註欄位**  | 
 **chooseSubPayment** | **string** | **付款子項目** 若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。 請參考付款方式一覽表  | 
 **orderResultURL** | **string** | **Client端回傳付款結果網址** 當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。 詳細說明請參考付款結果通知 注意事項： 1. 若與[ClientBackURL]同時設定，將會以此參數為主。 2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。  | 
 **needExtraPaidInfo** | **string** | **是否需要額外的付款資訊** 額外的付款資訊： 若不回傳額外的付款資訊時，參數值請傳：&#x60;Ｎ&#x60;； 若要回傳額外的付款資訊時，參數值請傳：&#x60;Ｙ&#x60;，付款完成後綠界會以 Server POST 方式回傳額外付款資訊。 注意事項： 回傳額外付款資訊參數請參考-額外回傳的參數  | [default to &quot;N&quot;]
 **deviceSource** | **string** | **裝置來源** 請帶空值，由系統自動判定。  | 
 **ignorePayment** | **string** | **隱藏付款** 當付款方式 &#x60;ChoosePayment&#x60; 為 &#x60;ALL&#x60; 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。 可用的參數值： - &#x60;Credit&#x60;: 信用卡 - &#x60;WebATM&#x60;: 網路 ATM - &#x60;ATM&#x60;: 自動櫃員機 - &#x60;CVS&#x60;: 超商代碼 - &#x60;BARCODE&#x60;: 超商條碼  | 
 **platformID** | **string** | **特約合作平台商代號(由綠界提供)** 為專案合作的平台商使用。 一般特店或平台商本身介接，則參數請帶放空值。 若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 &#x60;MerchantID&#x60;。  | 
 **invoiceMark** | **string** | **電子發票開立註記** 此參數為付款完成後同時開立電子發票。 若要使用時，該參數須設定為「Y」，同時還要設定「電子發票介接相關參數」 注意事項： 正式環境欲使用電子發票功能，須與綠界申請開通，若未開通請致電客服中心 &#x60;(02) 2655-1775&#x60;。  | 
 **customField1** | **string** | **自訂名稱欄位1** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **customField2** | **string** | **自訂名稱欄位2** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **customField3** | **string** | **自訂名稱欄位3** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **customField4** | **string** | **自訂名稱欄位4** 提供合作廠商使用記錄用客製化使用欄位 注意事項： 特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **encryptType** | **int32** | **CheckMacValue加密類型** 請固定填入 1，使用 SHA256 加密。  | [default to 1]
 **language** | **string** | **語系設定** 預設語系為中文，若要變更語系參數值請帶： - 英語：&#x60;ENG&#x60; - 韓語：&#x60;KOR&#x60; - 日語：&#x60;JPN&#x60; - 簡體中文：&#x60;CHI&#x60;  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

