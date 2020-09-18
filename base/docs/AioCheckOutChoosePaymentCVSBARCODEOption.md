# AioCheckOutChoosePaymentCvsBarcodeOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreExpireDate** | **int32** | **超商繳費截止時間** 注意事項： &#x60;CVS&#x60;:以分鐘為單位 &#x60;BARCODE&#x60;:以天為單位 若未設定此參數，&#x60;CVS&#x60; 預設為 &#x60;10080&#x60; 分鐘(&#x60;7&#x60; 天)；BARCODE 預設為 &#x60;7&#x60; 天。 若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 &#x60;86400&#x60; 分鐘，超過時一律以 &#x60;86400&#x60; 分鐘計(&#x60;60&#x60; 天) 例：&#x60;08/01&#x60; 的 &#x60;20:15&#x60; 分購買商品，繳費期限為 &#x60;7&#x60; 天，表示 &#x60;8/08&#x60; 的 &#x60;20:15&#x60; 分前您必須前往超商繳費。 範例: &#x60;CVS&#x60;&#x3D;&#x60;1440&#x60;(共 &#x60;1&#x60; 天)、&#x60;BARCODE&#x60;&#x3D;&#x60;7&#x60;(共 &#x60;7&#x60; 天)  | [optional] 
**Desc1** | **string** | **交易描述1** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc2** | **string** | **交易描述2** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc3** | **string** | **交易描述3** 會出現在超商繳費平台螢幕上  | [optional] 
**Desc4** | **string** | **交易描述4** 會出現在超商繳費平台螢幕上  | [optional] 
**PaymentInfoURL** | **string** | **Server 端回傳付款相關資訊** 若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。 請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.] 注意事項： 頁面將會停留在綠界，顯示繳費的相關資訊。 回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。  | [optional] 
**ClientRedirectURL** | **string** | **Client端回傳付款方式相關資訊** 若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。 請參考[&#x60;ATM&#x60;、&#x60;CVS&#x60; 或 &#x60;BARCODE&#x60; 的取號結果通知.] 注意事項： 若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。 若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。 回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


