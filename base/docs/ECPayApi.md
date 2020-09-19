# \ECPayApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AioCheckOut**](ECPayApi.md#AioCheckOut) | **Post** /Cashier/AioCheckOut/V5 | 



## AioCheckOut

> AioCheckOut(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).MerchantTradeDate(merchantTradeDate).PaymentType(paymentType).TotalAmount(totalAmount).TradeDesc(tradeDesc).ItemName(itemName).ReturnURL(returnURL).ChoosePayment(choosePayment).CheckMacValue(checkMacValue).EncryptType(encryptType).StoreID(storeID).ClientBackURL(clientBackURL).ItemURL(itemURL).Remark(remark).ChooseSubPayment(chooseSubPayment).OrderResultURL(orderResultURL).NeedExtraPaidInfo(needExtraPaidInfo).DeviceSource(deviceSource).IgnorePayment(ignorePayment).PlatformID(platformID).InvoiceMark(invoiceMark).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).Language(language).StoreExpireDate(storeExpireDate).Desc1(desc1).Desc2(desc2).Desc3(desc3).Desc4(desc4).PaymentInfoURL(paymentInfoURL).ClientRedirectURL(clientRedirectURL).BindingCard(bindingCard).MerchantMemberID(merchantMemberID).Redeem(redeem).UnionPay(unionPay).CreditInstallment(creditInstallment).PeriodAmount(periodAmount).PeriodType(periodType).Frequency(frequency).ExecTimes(execTimes).PeriodReturnURL(periodReturnURL).RelateNumber(relateNumber).CustomerID(customerID).CustomerIdentifier(customerIdentifier).CustomerName(customerName).CustomerAddr(customerAddr).CustomerPhone(customerPhone).CustomerEmail(customerEmail).ClearanceMark(clearanceMark).TaxType(taxType).CarruerType(carruerType).CarruerNum(carruerNum).Donation(donation).LoveCode(loveCode).Print(print).InvoiceItemName(invoiceItemName).InvoiceItemCount(invoiceItemCount).InvoiceItemWord(invoiceItemWord).InvoiceItemPrice(invoiceItemPrice).InvoiceItemTaxType(invoiceItemTaxType).InvoiceRemark(invoiceRemark).DelayDay(delayDay).InvType(invType).Execute()



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
    merchantID := "merchantID_example" // string | **特店編號(由綠界提供)** 
    merchantTradeNo := "merchantTradeNo_example" // string | **特店交易編號(由特店提供)**   特店交易編號均為唯一值，不可重複使用。   英數字大小寫混合   如何避免訂單編號重複請參考 FAQ   如有使用 `PlatformID` ，平台商底下所有商家之訂單編號亦不可重複。   
    merchantTradeDate := "merchantTradeDate_example" // ECPayDateTime | **特店交易時間** 格式為 `yyyy/MM/dd HH:mm:ss` 
    paymentType := openapiclient.AioCheckPaymentTypeEnum{} // AioCheckPaymentTypeEnum |  (default to "aio")
    totalAmount := 987 // int32 | **交易金額**   請帶整數，不可有小數點。   僅限新台幣。   各付款金額的限制，請參考 <https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID=3605> 
    tradeDesc := "tradeDesc_example" // string | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。 
    itemName := "itemName_example" // string | **商品名稱**   1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。   2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。   
    returnURL := "returnURL_example" // string | **付款完成通知回傳網址**   當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：    1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。   2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。   
    choosePayment := openapiclient.ChoosePaymentEnum{} // ChoosePaymentEnum | 
    checkMacValue := "checkMacValue_example" // string | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式 
    encryptType := openapiclient.EncryptTypeEnum{} // EncryptTypeEnum |  (default to 1)
    storeID := "storeID_example" // string | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  (optional)
    clientBackURL := "clientBackURL_example" // string | **Client端返回特店的按鈕連結**   消費者點選此按鈕後，會將頁面導回到此設定的網址   注意事項：   導回時不會帶付款結果到此網址，只是將頁面導回而已。   設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。   設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。   若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。    (optional)
    itemURL := "itemURL_example" // string | **商品銷售網址**  (optional)
    remark := "remark_example" // string | **備註欄位**  (optional)
    chooseSubPayment := "chooseSubPayment_example" // string | **付款子項目**   若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。   請參考付款方式一覽表    (optional)
    orderResultURL := "orderResultURL_example" // string | **Client端回傳付款結果網址**     當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：   1. 若與[ClientBackURL]同時設定，將會以此參數為主。   2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。    (optional)
    needExtraPaidInfo := openapiclient.NeedExtraPaidInfoEnum{} // NeedExtraPaidInfoEnum |  (optional)
    deviceSource := "deviceSource_example" // string | **裝置來源** 請帶空值，由系統自動判定。  (optional)
    ignorePayment := "ignorePayment_example" // string | **隱藏付款**   當付款方式 `ChoosePayment` 為 `ALL` 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。   可用的參數值：   - `Credit`: 信用卡   - `WebATM`: 網路 ATM   - `ATM`: 自動櫃員機   - `CVS`: 超商代碼   - `BARCODE`: 超商條碼    (optional)
    platformID := "platformID_example" // string | **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 `MerchantID`。    (optional)
    invoiceMark := openapiclient.InvoiceMarkEnum{} // InvoiceMarkEnum |  (optional)
    customField1 := "customField1_example" // string | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    customField2 := "customField2_example" // string | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    customField3 := "customField3_example" // string | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    customField4 := "customField4_example" // string | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`  (optional)
    language := openapiclient.LanguageEnum{} // LanguageEnum |  (optional)
    storeExpireDate := 987 // int32 | **超商繳費截止時間**   注意事項：   `CVS`:以分鐘為單位   `BARCODE`:以天為單位   若未設定此參數，`CVS` 預設為 `10080` 分鐘(`7` 天)；BARCODE 預設為 `7` 天。   若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 `86400` 分鐘，超過時一律以 `86400` 分鐘計(`60` 天)   例：`08/01` 的 `20:15` 分購買商品，繳費期限為 `7` 天，表示 `8/08` 的 `20:15` 分前您必須前往超商繳費。   範例: `CVS`=`1440`(共 `1` 天)、`BARCODE`=`7`(共 `7` 天)    (optional)
    desc1 := "desc1_example" // string | **交易描述1** 會出現在超商繳費平台螢幕上  (optional)
    desc2 := "desc2_example" // string | **交易描述2** 會出現在超商繳費平台螢幕上  (optional)
    desc3 := "desc3_example" // string | **交易描述3** 會出現在超商繳費平台螢幕上  (optional)
    desc4 := "desc4_example" // string | **交易描述4** 會出現在超商繳費平台螢幕上  (optional)
    paymentInfoURL := "paymentInfoURL_example" // string | **Server 端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。   請參考[`ATM`、`CVS` 或 `BARCODE` 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。    (optional)
    clientRedirectURL := "clientRedirectURL_example" // string | **Client端回傳付款方式相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。   請參考[`ATM`、`CVS` 或 `BARCODE` 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。    (optional)
    bindingCard := openapiclient.BindingCardEnum{} // BindingCardEnum |  (optional)
    merchantMemberID := "merchantMemberID_example" // string | **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 `MerchantID` + `廠商會員編號`)  (optional)
    redeem := openapiclient.RedeemEnum{} // RedeemEnum |  (optional)
    unionPay := openapiclient.UnionPayEnum{} // UnionPayEnum |  (optional) (default to 0)
    creditInstallment := "creditInstallment_example" // string | **刷卡分期期數**    提供刷卡分期期數   信用卡分期可用參數為:3,6,12,18,24   注意事項：   使用的期數必須先透過申請開通後方能使用，並以申請開通的期數為主。    (optional)
    periodAmount := 987 // int32 | **每次授權金額**   每次要授權(扣款)的金額。   注意事項：   綠界會依此次授權金額`PeriodAmount`所設定的金額做為之後固定授權的金額。   交易金額`TotalAmount`設定金額必須和授權金額`PeriodAmount`相同。   請帶整數，不可有小數點。僅限新台幣。  (optional)
    periodType := openapiclient.CreditPeriodTypeEnum{} // CreditPeriodTypeEnum |  (optional)
    frequency := 987 // int32 | **執行頻率**   此參數用來定義多久要執行一次   注意事項：   至少要大於等於 1 次以上。   當 `PeriodType` 設為 `D` 時，最多可設 `365` 次。    當 `PeriodType` 設為 `M` 時，最多可設 `12` 次。     當 `PeriodType` 設為 `Y` 時，最多可設 `1` 次。  (optional)
    execTimes := 987 // int32 | **執行次數**   總共要執行幾次。   注意事項：   至少要大於 1 次以上。   當 `PeriodType` 設為 `D` 時，最多可設 `999`次。   當 `PeriodType` 設為 `M` 時，最多可設 `99`次。   當 `PeriodType` 設為 `Y` 時，最多可設 `9` 次。    例 1：   當信用卡定期定額扣款為每個月扣 1 次500 元，總共要扣 12次   `TotalAmount`參數請帶 `500` `PeriodAmount`=`500`   `PeriodType`=`M`   `Frequency`=`1`   `ExecTimes`=`12`    例 2：   當信用卡定期定額扣款為 6000 元，每 6 個月扣 1 次，總共要扣 2 次時    交易金額`TotalAmount`參數請帶 `6000`   `PeriodType`=`M`   `Frequency`=`6`   `ExecTimes`=`2`    (optional)
    periodReturnURL := "periodReturnURL_example" // string | **定期定額的執行結果回應URL**   若交易是信用卡定期定額的方式，則每次執行授權完，會將授權結果回傳到這個設定的 URL。   回覆內容請參考。    (optional)
    relateNumber := "relateNumber_example" // string | **特店自訂編號** 此為特店自訂編號，編號均為唯一值不可重複使用。  (optional)
    customerID := "customerID_example" // string | **客戶編號** 該參數有值時，僅接受『英文、數字、下底線』等字元。  (optional)
    customerIdentifier := "customerIdentifier_example" // string | **統一編號** 該參數有值時，請帶固定長度為數字 8 碼。  (optional)
    customerName := "customerName_example" // string | **客戶名稱**    當列印註記`Print`為 `1`(列印)時，則該參數必須有值。   該參數有值時，僅接受『中、英文及數字』等字元。   請將參數值做 UrlEncode 方式編碼。    (optional)
    customerAddr := "customerAddr_example" // string | **客戶地址**    當列印註記`Print`為 `1`(列印)時，則該參數必須有值。   當該參數有值時，請注意特殊字元轉換 。    請將參數值做 UrlEncode 方式編碼。  (optional)
    customerPhone := "customerPhone_example" // string | **客戶手機號碼**   當客戶電子信箱`CustomerEmail`為空字串時，則該參數必須有值。   當該參數有值時，則格式為數字。   注意事項：   請填手機號碼，不能填市話因為要收簡訊通知用    (optional)
    customerEmail := "customerEmail_example" // string | **客戶電子信箱**   當客戶手機號碼`CustomerPhone`為空字串時，則該參數必須有值。   當該參數有值時，則格式需符合 EMAIL格式。   請將參數值做 UrlEncode 方式編碼。    (optional)
    clearanceMark := openapiclient.ClearanceMarkEnum{} // ClearanceMarkEnum |  (optional)
    taxType := openapiclient.TaxTypeEnum{} // TaxTypeEnum |  (optional)
    carruerType := openapiclient.CarruerTypeEnum{} // CarruerTypeEnum |  (optional)
    carruerNum := "carruerNum_example" // string | **載具編號**   1. 當載具類別 `CarruerType`=``無載具)，請帶空字串。   2. 當載具類別`CarruerType`=`1`(綠界科技電子發票載具)時，請帶空字串，系統會自動帶入值，為合作特店載具統一編號+自訂編號(RelateNumber)。   3. 當載具類別`CarruerType`=`2`(買受人之自然人憑證)時，則請帶固定長度為16且格式 為2碼大寫英文字母加上14碼數字。   4. 當載具類別`CarruerType`=`3`(買受人之手機條碼)時，則請帶固定長度為 8且格式為 1 碼斜線「/」加上由 7 碼數字及大寫英文字母及+-.符號組成。    注意事項：   1. 若手機條碼中有加號，可能在介接驗證時 發生錯誤，請將加號改為空白字元，產生 驗證碼。   2. 英文、數字、符號僅接受半形字   3. 若載具編號為手機條碼載具時，請先呼叫B2C電子發票介接技術文件手機條碼載驗證ＡＰＩ進行檢核    (optional)
    donation := openapiclient.InvoiceDonationEunm{} // InvoiceDonationEunm |  (optional)
    loveCode := "loveCode_example" // string | **捐贈碼**   消費者選擇捐贈發票則於此欄位須填入受贈單位之捐贈碼。   1. 若捐贈註記 `Donation`= `1` (捐贈)時，此欄位須有值。   2. 捐贈碼以阿拉伯數字為限，最少三碼，最多七碼。內容定位採「文字格式」，首位可以為零。     (optional)
    print := openapiclient.InvoicePrintEnum{} // InvoicePrintEnum |  (optional)
    invoiceItemName := "invoiceItemName_example" // string | **商品名稱**   預設不可為空字串且格式為 名稱 1 | 名稱 2 | 名稱 3 | … | 名稱 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。   將參數值以 UrlEncode 方式編碼。    (optional)
    invoiceItemCount := "invoiceItemCount_example" // string | **商品數量**   預設不可為空字串且格式為 數量 1 | 數量 2 | 數量 3 | … | 數量 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。  (optional)
    invoiceItemWord := "invoiceItemWord_example" // string | **商品單位**   商品單位若超過二筆以上請以「|」符號區隔單位最大長度為 6 碼。   請將參數做 UrlEncode 方式編碼。  (optional)
    invoiceItemPrice := "invoiceItemPrice_example" // string | **商品價格**   預設不可為空字串且格式為 價格 1 | 價格 2 | 價格 3 | … | 價格 n，當含有二筆或以上的商品價格時，則以「|」符號區隔。  (optional)
    invoiceItemTaxType := "invoiceItemTaxType_example" // string | **商品課稅別**   1：應稅   2：零稅率   3：免稅   注意事項：   1. 預設為空字串，當課稅類別 [TaxType] = 9 時，此欄位不可為空。   2. 格式為課稅 類別 1 | 課稅類別 2 | 課稅類別 3 | … | 課稅類別 n。當含有二筆或以上的商品課稅類別時，則以「|」符號區隔。   3. 課稅類別為混合稅率時，需含二筆或 以 上 的 商 品 課 稅   別[InvoiceItemTaxType]，且至少需有一筆商品課稅別為應稅及至少需有一筆商品課稅別為免稅或零稅率，即混稅發票只能 1.應稅+免稅 2.應稅+零稅率，免稅和零稅率發票不能同時開立。  (optional)
    invoiceRemark := "invoiceRemark_example" // string | **備註** 當該參數有值時，請將參數值做UrlEncode 方式編碼。  (optional)
    delayDay := "delayDay_example" // string | **延遲天數**   本參數值請帶 0~15(天)，當天數為 0 時，則付款完成後立即開立發票。  (optional)
    invType := "invType_example" // string | **字軌類別**   若為一般稅額時，請帶 07。   預設值：07    (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ECPayApi.AioCheckOut(context.Background(), merchantID, merchantTradeNo, merchantTradeDate, paymentType, totalAmount, tradeDesc, itemName, returnURL, choosePayment, checkMacValue, encryptType).StoreID(storeID).ClientBackURL(clientBackURL).ItemURL(itemURL).Remark(remark).ChooseSubPayment(chooseSubPayment).OrderResultURL(orderResultURL).NeedExtraPaidInfo(needExtraPaidInfo).DeviceSource(deviceSource).IgnorePayment(ignorePayment).PlatformID(platformID).InvoiceMark(invoiceMark).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).Language(language).StoreExpireDate(storeExpireDate).Desc1(desc1).Desc2(desc2).Desc3(desc3).Desc4(desc4).PaymentInfoURL(paymentInfoURL).ClientRedirectURL(clientRedirectURL).BindingCard(bindingCard).MerchantMemberID(merchantMemberID).Redeem(redeem).UnionPay(unionPay).CreditInstallment(creditInstallment).PeriodAmount(periodAmount).PeriodType(periodType).Frequency(frequency).ExecTimes(execTimes).PeriodReturnURL(periodReturnURL).RelateNumber(relateNumber).CustomerID(customerID).CustomerIdentifier(customerIdentifier).CustomerName(customerName).CustomerAddr(customerAddr).CustomerPhone(customerPhone).CustomerEmail(customerEmail).ClearanceMark(clearanceMark).TaxType(taxType).CarruerType(carruerType).CarruerNum(carruerNum).Donation(donation).LoveCode(loveCode).Print(print).InvoiceItemName(invoiceItemName).InvoiceItemCount(invoiceItemCount).InvoiceItemWord(invoiceItemWord).InvoiceItemPrice(invoiceItemPrice).InvoiceItemTaxType(invoiceItemTaxType).InvoiceRemark(invoiceRemark).DelayDay(delayDay).InvType(invType).Execute()
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
 **merchantID** | **string** | **特店編號(由綠界提供)**  | 
 **merchantTradeNo** | **string** | **特店交易編號(由特店提供)**   特店交易編號均為唯一值，不可重複使用。   英數字大小寫混合   如何避免訂單編號重複請參考 FAQ   如有使用 &#x60;PlatformID&#x60; ，平台商底下所有商家之訂單編號亦不可重複。    | 
 **merchantTradeDate** | **ECPayDateTime** | **特店交易時間** 格式為 &#x60;yyyy/MM/dd HH:mm:ss&#x60;  | 
 **paymentType** | [**AioCheckPaymentTypeEnum**](AioCheckPaymentTypeEnum.md) |  | [default to &quot;aio&quot;]
 **totalAmount** | **int32** | **交易金額**   請帶整數，不可有小數點。   僅限新台幣。   各付款金額的限制，請參考 &lt;https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID&#x3D;3605&gt;  | 
 **tradeDesc** | **string** | **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。  | 
 **itemName** | **string** | **商品名稱**   1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。   2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。    | 
 **returnURL** | **string** | **付款完成通知回傳網址**   當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：    1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。   2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。    | 
 **choosePayment** | [**ChoosePaymentEnum**](ChoosePaymentEnum.md) |  | 
 **checkMacValue** | **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 
 **encryptType** | [**EncryptTypeEnum**](EncryptTypeEnum.md) |  | [default to 1]
 **storeID** | **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
 **clientBackURL** | **string** | **Client端返回特店的按鈕連結**   消費者點選此按鈕後，會將頁面導回到此設定的網址   注意事項：   導回時不會帶付款結果到此網址，只是將頁面導回而已。   設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。   設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。   若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。    | 
 **itemURL** | **string** | **商品銷售網址**  | 
 **remark** | **string** | **備註欄位**  | 
 **chooseSubPayment** | **string** | **付款子項目**   若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。   請參考付款方式一覽表    | 
 **orderResultURL** | **string** | **Client端回傳付款結果網址**     當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：   1. 若與[ClientBackURL]同時設定，將會以此參數為主。   2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。    | 
 **needExtraPaidInfo** | [**NeedExtraPaidInfoEnum**](NeedExtraPaidInfoEnum.md) |  | 
 **deviceSource** | **string** | **裝置來源** 請帶空值，由系統自動判定。  | 
 **ignorePayment** | **string** | **隱藏付款**   當付款方式 &#x60;ChoosePayment&#x60; 為 &#x60;ALL&#x60; 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。   可用的參數值：   - &#x60;Credit&#x60;: 信用卡   - &#x60;WebATM&#x60;: 網路 ATM   - &#x60;ATM&#x60;: 自動櫃員機   - &#x60;CVS&#x60;: 超商代碼   - &#x60;BARCODE&#x60;: 超商條碼    | 
 **platformID** | **string** | **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 &#x60;MerchantID&#x60;。    | 
 **invoiceMark** | [**InvoiceMarkEnum**](InvoiceMarkEnum.md) |  | 
 **customField1** | **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **customField2** | **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **customField3** | **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **customField4** | **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 &#x60;,.#()$[];%{}:/?&amp;@&lt;&gt;!&#x60;  | 
 **language** | [**LanguageEnum**](LanguageEnum.md) |  | 
 **storeExpireDate** | **int32** | **超商繳費截止時間**   注意事項：   &#x60;CVS&#x60;:以分鐘為單位   &#x60;BARCODE&#x60;:以天為單位   若未設定此參數，&#x60;CVS&#x60; 預設為 &#x60;10080&#x60; 分鐘(&#x60;7&#x60; 天)；BARCODE 預設為 &#x60;7&#x60; 天。   若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 &#x60;86400&#x60; 分鐘，超過時一律以 &#x60;86400&#x60; 分鐘計(&#x60;60&#x60; 天)   例：&#x60;08/01&#x60; 的 &#x60;20:15&#x60; 分購買商品，繳費期限為 &#x60;7&#x60; 天，表示 &#x60;8/08&#x60; 的 &#x60;20:15&#x60; 分前您必須前往超商繳費。   範例: &#x60;CVS&#x60;&#x3D;&#x60;1440&#x60;(共 &#x60;1&#x60; 天)、&#x60;BARCODE&#x60;&#x3D;&#x60;7&#x60;(共 &#x60;7&#x60; 天)    | 
 **desc1** | **string** | **交易描述1** 會出現在超商繳費平台螢幕上  | 
 **desc2** | **string** | **交易描述2** 會出現在超商繳費平台螢幕上  | 
 **desc3** | **string** | **交易描述3** 會出現在超商繳費平台螢幕上  | 
 **desc4** | **string** | **交易描述4** 會出現在超商繳費平台螢幕上  | 
 **paymentInfoURL** | **string** | **Server 端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。   請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。    | 
 **clientRedirectURL** | **string** | **Client端回傳付款方式相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。   請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。    | 
 **bindingCard** | [**BindingCardEnum**](BindingCardEnum.md) |  | 
 **merchantMemberID** | **string** | **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 &#x60;MerchantID&#x60; + &#x60;廠商會員編號&#x60;)  | 
 **redeem** | [**RedeemEnum**](RedeemEnum.md) |  | 
 **unionPay** | [**UnionPayEnum**](UnionPayEnum.md) |  | [default to 0]
 **creditInstallment** | **string** | **刷卡分期期數**    提供刷卡分期期數   信用卡分期可用參數為:3,6,12,18,24   注意事項：   使用的期數必須先透過申請開通後方能使用，並以申請開通的期數為主。    | 
 **periodAmount** | **int32** | **每次授權金額**   每次要授權(扣款)的金額。   注意事項：   綠界會依此次授權金額&#x60;PeriodAmount&#x60;所設定的金額做為之後固定授權的金額。   交易金額&#x60;TotalAmount&#x60;設定金額必須和授權金額&#x60;PeriodAmount&#x60;相同。   請帶整數，不可有小數點。僅限新台幣。  | 
 **periodType** | [**CreditPeriodTypeEnum**](CreditPeriodTypeEnum.md) |  | 
 **frequency** | **int32** | **執行頻率**   此參數用來定義多久要執行一次   注意事項：   至少要大於等於 1 次以上。   當 &#x60;PeriodType&#x60; 設為 &#x60;D&#x60; 時，最多可設 &#x60;365&#x60; 次。    當 &#x60;PeriodType&#x60; 設為 &#x60;M&#x60; 時，最多可設 &#x60;12&#x60; 次。     當 &#x60;PeriodType&#x60; 設為 &#x60;Y&#x60; 時，最多可設 &#x60;1&#x60; 次。  | 
 **execTimes** | **int32** | **執行次數**   總共要執行幾次。   注意事項：   至少要大於 1 次以上。   當 &#x60;PeriodType&#x60; 設為 &#x60;D&#x60; 時，最多可設 &#x60;999&#x60;次。   當 &#x60;PeriodType&#x60; 設為 &#x60;M&#x60; 時，最多可設 &#x60;99&#x60;次。   當 &#x60;PeriodType&#x60; 設為 &#x60;Y&#x60; 時，最多可設 &#x60;9&#x60; 次。    例 1：   當信用卡定期定額扣款為每個月扣 1 次500 元，總共要扣 12次   &#x60;TotalAmount&#x60;參數請帶 &#x60;500&#x60; &#x60;PeriodAmount&#x60;&#x3D;&#x60;500&#x60;   &#x60;PeriodType&#x60;&#x3D;&#x60;M&#x60;   &#x60;Frequency&#x60;&#x3D;&#x60;1&#x60;   &#x60;ExecTimes&#x60;&#x3D;&#x60;12&#x60;    例 2：   當信用卡定期定額扣款為 6000 元，每 6 個月扣 1 次，總共要扣 2 次時    交易金額&#x60;TotalAmount&#x60;參數請帶 &#x60;6000&#x60;   &#x60;PeriodType&#x60;&#x3D;&#x60;M&#x60;   &#x60;Frequency&#x60;&#x3D;&#x60;6&#x60;   &#x60;ExecTimes&#x60;&#x3D;&#x60;2&#x60;    | 
 **periodReturnURL** | **string** | **定期定額的執行結果回應URL**   若交易是信用卡定期定額的方式，則每次執行授權完，會將授權結果回傳到這個設定的 URL。   回覆內容請參考。    | 
 **relateNumber** | **string** | **特店自訂編號** 此為特店自訂編號，編號均為唯一值不可重複使用。  | 
 **customerID** | **string** | **客戶編號** 該參數有值時，僅接受『英文、數字、下底線』等字元。  | 
 **customerIdentifier** | **string** | **統一編號** 該參數有值時，請帶固定長度為數字 8 碼。  | 
 **customerName** | **string** | **客戶名稱**    當列印註記&#x60;Print&#x60;為 &#x60;1&#x60;(列印)時，則該參數必須有值。   該參數有值時，僅接受『中、英文及數字』等字元。   請將參數值做 UrlEncode 方式編碼。    | 
 **customerAddr** | **string** | **客戶地址**    當列印註記&#x60;Print&#x60;為 &#x60;1&#x60;(列印)時，則該參數必須有值。   當該參數有值時，請注意特殊字元轉換 。    請將參數值做 UrlEncode 方式編碼。  | 
 **customerPhone** | **string** | **客戶手機號碼**   當客戶電子信箱&#x60;CustomerEmail&#x60;為空字串時，則該參數必須有值。   當該參數有值時，則格式為數字。   注意事項：   請填手機號碼，不能填市話因為要收簡訊通知用    | 
 **customerEmail** | **string** | **客戶電子信箱**   當客戶手機號碼&#x60;CustomerPhone&#x60;為空字串時，則該參數必須有值。   當該參數有值時，則格式需符合 EMAIL格式。   請將參數值做 UrlEncode 方式編碼。    | 
 **clearanceMark** | [**ClearanceMarkEnum**](ClearanceMarkEnum.md) |  | 
 **taxType** | [**TaxTypeEnum**](TaxTypeEnum.md) |  | 
 **carruerType** | [**CarruerTypeEnum**](CarruerTypeEnum.md) |  | 
 **carruerNum** | **string** | **載具編號**   1. 當載具類別 &#x60;CarruerType&#x60;&#x3D;&#x60;&#x60;無載具)，請帶空字串。   2. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;1&#x60;(綠界科技電子發票載具)時，請帶空字串，系統會自動帶入值，為合作特店載具統一編號+自訂編號(RelateNumber)。   3. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;2&#x60;(買受人之自然人憑證)時，則請帶固定長度為16且格式 為2碼大寫英文字母加上14碼數字。   4. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;3&#x60;(買受人之手機條碼)時，則請帶固定長度為 8且格式為 1 碼斜線「/」加上由 7 碼數字及大寫英文字母及+-.符號組成。    注意事項：   1. 若手機條碼中有加號，可能在介接驗證時 發生錯誤，請將加號改為空白字元，產生 驗證碼。   2. 英文、數字、符號僅接受半形字   3. 若載具編號為手機條碼載具時，請先呼叫B2C電子發票介接技術文件手機條碼載驗證ＡＰＩ進行檢核    | 
 **donation** | [**InvoiceDonationEunm**](InvoiceDonationEunm.md) |  | 
 **loveCode** | **string** | **捐贈碼**   消費者選擇捐贈發票則於此欄位須填入受贈單位之捐贈碼。   1. 若捐贈註記 &#x60;Donation&#x60;&#x3D; &#x60;1&#x60; (捐贈)時，此欄位須有值。   2. 捐贈碼以阿拉伯數字為限，最少三碼，最多七碼。內容定位採「文字格式」，首位可以為零。     | 
 **print** | [**InvoicePrintEnum**](InvoicePrintEnum.md) |  | 
 **invoiceItemName** | **string** | **商品名稱**   預設不可為空字串且格式為 名稱 1 | 名稱 2 | 名稱 3 | … | 名稱 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。   將參數值以 UrlEncode 方式編碼。    | 
 **invoiceItemCount** | **string** | **商品數量**   預設不可為空字串且格式為 數量 1 | 數量 2 | 數量 3 | … | 數量 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。  | 
 **invoiceItemWord** | **string** | **商品單位**   商品單位若超過二筆以上請以「|」符號區隔單位最大長度為 6 碼。   請將參數做 UrlEncode 方式編碼。  | 
 **invoiceItemPrice** | **string** | **商品價格**   預設不可為空字串且格式為 價格 1 | 價格 2 | 價格 3 | … | 價格 n，當含有二筆或以上的商品價格時，則以「|」符號區隔。  | 
 **invoiceItemTaxType** | **string** | **商品課稅別**   1：應稅   2：零稅率   3：免稅   注意事項：   1. 預設為空字串，當課稅類別 [TaxType] &#x3D; 9 時，此欄位不可為空。   2. 格式為課稅 類別 1 | 課稅類別 2 | 課稅類別 3 | … | 課稅類別 n。當含有二筆或以上的商品課稅類別時，則以「|」符號區隔。   3. 課稅類別為混合稅率時，需含二筆或 以 上 的 商 品 課 稅   別[InvoiceItemTaxType]，且至少需有一筆商品課稅別為應稅及至少需有一筆商品課稅別為免稅或零稅率，即混稅發票只能 1.應稅+免稅 2.應稅+零稅率，免稅和零稅率發票不能同時開立。  | 
 **invoiceRemark** | **string** | **備註** 當該參數有值時，請將參數值做UrlEncode 方式編碼。  | 
 **delayDay** | **string** | **延遲天數**   本參數值請帶 0~15(天)，當天數為 0 時，則付款完成後立即開立發票。  | 
 **invType** | **string** | **字軌類別**   若為一般稅額時，請帶 07。   預設值：07    | 

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

