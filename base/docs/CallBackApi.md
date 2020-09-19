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
    merchantID := "merchantID_example" // string | 特店編號
    merchantTradeNo := "merchantTradeNo_example" // string | 特店交易編號
    storeID := "storeID_example" // string | 特店旗下店舖代號
    rtnCode := 987 // float32 | 交易狀態
    rtnMsg := "rtnMsg_example" // string | 交易訊息
    tradeNo := "tradeNo_example" // string | 綠界的交易編號
    tradeAmt := 987 // int32 | 交易金額
    paymentDate := "paymentDate_example" // string | 付款時間
    paymentType := openapiclient.PaymentTypeEnum{} // PaymentTypeEnum | 特店選擇的付款方式
    paymentTypeChargeFee := 987 // int32 | 通路費
    tradeDate := "tradeDate_example" // string | 訂單成立時間
    checkMacValue := "checkMacValue_example" // string | 檢查碼
    simulatePaid := openapiclient.ECPaySimulateEnum{} // ECPaySimulateEnum | 是否為模擬付款
    customField1 := "customField1_example" // string | 自訂名稱欄位 1 (optional)
    customField2 := "customField2_example" // string | 自訂名稱欄位 2 (optional)
    customField3 := "customField3_example" // string | 自訂名稱欄位 3 (optional)
    customField4 := "customField4_example" // string | 自訂名稱欄位 4 (optional)
    customField5 := "customField5_example" // string | 自訂名稱欄位 5 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallBackApi.CallbackReturnURLPost(context.Background(), merchantID, merchantTradeNo, storeID, rtnCode, rtnMsg, tradeNo, tradeAmt, paymentDate, paymentType, paymentTypeChargeFee, tradeDate, checkMacValue, simulatePaid).CustomField1(customField1).CustomField2(customField2).CustomField3(customField3).CustomField4(customField4).CustomField5(customField5).Execute()
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
 **merchantID** | **string** | 特店編號 | 
 **merchantTradeNo** | **string** | 特店交易編號 | 
 **storeID** | **string** | 特店旗下店舖代號 | 
 **rtnCode** | **float32** | 交易狀態 | 
 **rtnMsg** | **string** | 交易訊息 | 
 **tradeNo** | **string** | 綠界的交易編號 | 
 **tradeAmt** | **int32** | 交易金額 | 
 **paymentDate** | **string** | 付款時間 | 
 **paymentType** | [**PaymentTypeEnum**](.md) | 特店選擇的付款方式 | 
 **paymentTypeChargeFee** | **int32** | 通路費 | 
 **tradeDate** | **string** | 訂單成立時間 | 
 **checkMacValue** | **string** | 檢查碼 | 
 **simulatePaid** | [**ECPaySimulateEnum**](.md) | 是否為模擬付款 | 
 **customField1** | **string** | 自訂名稱欄位 1 | 
 **customField2** | **string** | 自訂名稱欄位 2 | 
 **customField3** | **string** | 自訂名稱欄位 3 | 
 **customField4** | **string** | 自訂名稱欄位 4 | 
 **customField5** | **string** | 自訂名稱欄位 5 | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

