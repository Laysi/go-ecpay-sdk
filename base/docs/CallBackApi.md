# \CallBackApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackReturnURLPost**](CallBackApi.md#CallbackReturnURLPost) | **Post** /callback/ReturnURL | 



## CallbackReturnURLPost

> string CallbackReturnURLPost(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).StoreID(storeID).RtnCode(rtnCode).RtnMsg(rtnMsg).TradeNo(tradeNo).TradeAmt(tradeAmt).PaymentDate(paymentDate).PaymentType(paymentType).PaymentTypeChargeFee(paymentTypeChargeFee).TradeDate(tradeDate).CheckMacValue(checkMacValue).SimulatePaid(simulatePaid).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).CustomField5(customField5).Execute()





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
    merchantID := "merchantID_example" // string | **特店編號** 
    merchantTradeNo := "merchantTradeNo_example" // string | **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合 
    storeID := "storeID_example" // string | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。 
    rtnCode := 987 // int | **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。 
    rtnMsg := "rtnMsg_example" // string | **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded   
    tradeNo := "tradeNo_example" // string | **綠界的交易編號** 請保存綠界的交易編號與特店交易編號[MerchantTradeNo]的關連。 
    tradeAmt := 987 // int | **交易金額** 
    paymentDate := "paymentDate_example" // ECPayDateTime | **付款時間** 格式為 yyyy/MM/dd HH:mm:ss 
    paymentType := openapiclient.PaymentTypeEnum{} // PaymentTypeEnum | 
    paymentTypeChargeFee := 987 // int | **通路費** 
    tradeDate := "tradeDate_example" // ECPayDateTime | **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss 
    checkMacValue := "checkMacValue_example" // string | **檢查碼** 特店必須檢查檢查碼`CheckMacValue`來驗證，請參考附錄檢查碼機制。 
    simulatePaid := openapiclient.SimulatePaidEnum{} // SimulatePaidEnum | 
    customField1 := "customField1_example" // string | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位 
    customField2 := "customField2_example" // string | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位 
    customField3 := "customField3_example" // string | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位 
    customField4 := "customField4_example" // string | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位 
    customField5 := "customField5_example" // string | **自訂名稱欄位5**   提供合作廠商使用記錄用客製化使用欄位 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallBackApi.CallbackReturnURLPost(context.Background(), merchantID, merchantTradeNo, storeID, rtnCode, rtnMsg, tradeNo, tradeAmt, paymentDate, paymentType, paymentTypeChargeFee, tradeDate, checkMacValue, simulatePaid, customField1, customField2, customField3, customField4, customField5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallBackApi.CallbackReturnURLPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallbackReturnURLPost`: string
    fmt.Fprintf(os.Stdout, "Response from `CallBackApi.CallbackReturnURLPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallbackReturnURLPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantID** | **string** | **特店編號**  | 
 **merchantTradeNo** | **string** | **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合  | 
 **storeID** | **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
 **rtnCode** | **int** | **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。  | 
 **rtnMsg** | **string** | **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded    | 
 **tradeNo** | **string** | **綠界的交易編號** 請保存綠界的交易編號與特店交易編號[MerchantTradeNo]的關連。  | 
 **tradeAmt** | **int** | **交易金額**  | 
 **paymentDate** | **ECPayDateTime** | **付款時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
 **paymentType** | [**PaymentTypeEnum**](PaymentTypeEnum.md) |  | 
 **paymentTypeChargeFee** | **int** | **通路費**  | 
 **tradeDate** | **ECPayDateTime** | **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
 **checkMacValue** | **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 
 **simulatePaid** | [**SimulatePaidEnum**](SimulatePaidEnum.md) |  | 
 **customField1** | **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField2** | **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField3** | **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField4** | **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField5** | **string** | **自訂名稱欄位5**   提供合作廠商使用記錄用客製化使用欄位  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

