# \CallBackApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackPeriodReturnURLPost**](CallBackApi.md#CallbackPeriodReturnURLPost) | **Post** /callback/PeriodReturnURL | 
[**CallbackReturnURLPost**](CallBackApi.md#CallbackReturnURLPost) | **Post** /callback/ReturnURL | 



## CallbackPeriodReturnURLPost

> string CallbackPeriodReturnURLPost(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).StoreID(storeID).RtnCode(rtnCode).RtnMsg(rtnMsg).PeriodType(periodType).Frequency(frequency).ExecTimes(execTimes).Amount(amount).Gwsr(gwsr).ProcessDate(processDate).AuthCode(authCode).FirstAuthAmount(firstAuthAmount).TotalSuccessTimes(totalSuccessTimes).CheckMacValue(checkMacValue).SimulatePaid(simulatePaid).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).Execute()





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
    periodType := openapiclient.CreditPeriodTypeEnum{} // CreditPeriodTypeEnum | 
    frequency := 987 // int | **執行頻率**  訂單建立時所設定的執行頻率 
    execTimes := 987 // int | **執行次數**  訂單建立時所設定的執行頻率 
    amount := 987 // int | **本次授權金額**  此次所授權的金額 
    gwsr := 987 // int | **授權交易單號**  此次所授權的交易單號 
    processDate := "processDate_example" // ECPayDateTime | **處理時間** 處理時間 ( yyyy/MM/dd HH:mm:ss ) 
    authCode := "authCode_example" // string | **授權碼** 授權碼 
    firstAuthAmount := 987 // int | **初次授權金額** 定期定額交易的第一筆授權金額。 
    totalSuccessTimes := 987 // int | **已執行成功次數**  目前已成功授權的次數。 
    checkMacValue := "checkMacValue_example" // string | **檢查碼** 特店必須檢查檢查碼`CheckMacValue`來驗證，請參考附錄檢查碼機制。 
    simulatePaid := openapiclient.SimulatePaidEnum{} // SimulatePaidEnum | 
    customField1 := "customField1_example" // string | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位 
    customField2 := "customField2_example" // string | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位 
    customField3 := "customField3_example" // string | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位 
    customField4 := "customField4_example" // string | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallBackApi.CallbackPeriodReturnURLPost(context.Background(), merchantID, merchantTradeNo, storeID, rtnCode, rtnMsg, periodType, frequency, execTimes, amount, gwsr, processDate, authCode, firstAuthAmount, totalSuccessTimes, checkMacValue, simulatePaid, customField1, customField2, customField3, customField4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallBackApi.CallbackPeriodReturnURLPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallbackPeriodReturnURLPost`: string
    fmt.Fprintf(os.Stdout, "Response from `CallBackApi.CallbackPeriodReturnURLPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallbackPeriodReturnURLPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantID** | **string** | **特店編號**  | 
 **merchantTradeNo** | **string** | **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合  | 
 **storeID** | **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
 **rtnCode** | **int** | **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。  | 
 **rtnMsg** | **string** | **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded    | 
 **periodType** | [**CreditPeriodTypeEnum**](CreditPeriodTypeEnum.md) |  | 
 **frequency** | **int** | **執行頻率**  訂單建立時所設定的執行頻率  | 
 **execTimes** | **int** | **執行次數**  訂單建立時所設定的執行頻率  | 
 **amount** | **int** | **本次授權金額**  此次所授權的金額  | 
 **gwsr** | **int** | **授權交易單號**  此次所授權的交易單號  | 
 **processDate** | **ECPayDateTime** | **處理時間** 處理時間 ( yyyy/MM/dd HH:mm:ss )  | 
 **authCode** | **string** | **授權碼** 授權碼  | 
 **firstAuthAmount** | **int** | **初次授權金額** 定期定額交易的第一筆授權金額。  | 
 **totalSuccessTimes** | **int** | **已執行成功次數**  目前已成功授權的次數。  | 
 **checkMacValue** | **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 
 **simulatePaid** | [**SimulatePaidEnum**](SimulatePaidEnum.md) |  | 
 **customField1** | **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField2** | **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField3** | **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField4** | **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位  | 

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


## CallbackReturnURLPost

> string CallbackReturnURLPost(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).StoreID(storeID).RtnCode(rtnCode).RtnMsg(rtnMsg).TradeNo(tradeNo).TradeAmt(tradeAmt).PaymentDate(paymentDate).PaymentType(paymentType).PaymentTypeChargeFee(paymentTypeChargeFee).TradeDate(tradeDate).CheckMacValue(checkMacValue).SimulatePaid(simulatePaid).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).Execute()





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
    paymentType := openapiclient.ReturnPaymentTypeEnum{} // ReturnPaymentTypeEnum | 
    paymentTypeChargeFee := 987 // int | **通路費** 
    tradeDate := "tradeDate_example" // ECPayDateTime | **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss 
    checkMacValue := "checkMacValue_example" // string | **檢查碼** 特店必須檢查檢查碼`CheckMacValue`來驗證，請參考附錄檢查碼機制。 
    simulatePaid := openapiclient.SimulatePaidEnum{} // SimulatePaidEnum | 
    customField1 := "customField1_example" // string | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位 
    customField2 := "customField2_example" // string | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位 
    customField3 := "customField3_example" // string | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位 
    customField4 := "customField4_example" // string | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallBackApi.CallbackReturnURLPost(context.Background(), merchantID, merchantTradeNo, storeID, rtnCode, rtnMsg, tradeNo, tradeAmt, paymentDate, paymentType, paymentTypeChargeFee, tradeDate, checkMacValue, simulatePaid, customField1, customField2, customField3, customField4).Execute()
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
 **paymentType** | [**ReturnPaymentTypeEnum**](ReturnPaymentTypeEnum.md) |  | 
 **paymentTypeChargeFee** | **int** | **通路費**  | 
 **tradeDate** | **ECPayDateTime** | **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
 **checkMacValue** | **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 
 **simulatePaid** | [**SimulatePaidEnum**](SimulatePaidEnum.md) |  | 
 **customField1** | **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField2** | **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField3** | **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位  | 
 **customField4** | **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位  | 

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

