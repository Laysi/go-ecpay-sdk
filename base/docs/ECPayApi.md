# \ECPayApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashierQueryCreditCardPeriodInfoPost**](ECPayApi.md#CashierQueryCreditCardPeriodInfoPost) | **Post** /Cashier/QueryCreditCardPeriodInfo | 



## CashierQueryCreditCardPeriodInfoPost

> CreditCardPeriodInfo CashierQueryCreditCardPeriodInfoPost(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).TimeStamp(timeStamp).CheckMacValue(checkMacValue).Execute()



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
    merchantTradeNo := "merchantTradeNo_example" // string | **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。 
    timeStamp := 987 // int | **驗證時間**   將當下的時間轉為UnixTimeStamp(見範例)用於驗證此次介接的時間區間。   綠界驗證時間區間暫訂為 3 分鐘內有效，超過則此次介接無效。   
    checkMacValue := "checkMacValue_example" // string | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ECPayApi.CashierQueryCreditCardPeriodInfoPost(context.Background(), merchantID, merchantTradeNo, timeStamp, checkMacValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ECPayApi.CashierQueryCreditCardPeriodInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CashierQueryCreditCardPeriodInfoPost`: CreditCardPeriodInfo
    fmt.Fprintf(os.Stdout, "Response from `ECPayApi.CashierQueryCreditCardPeriodInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashierQueryCreditCardPeriodInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantID** | **string** | **特店編號(由綠界提供)**  | 
 **merchantTradeNo** | **string** | **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。  | 
 **timeStamp** | **int** | **驗證時間**   將當下的時間轉為UnixTimeStamp(見範例)用於驗證此次介接的時間區間。   綠界驗證時間區間暫訂為 3 分鐘內有效，超過則此次介接無效。    | 
 **checkMacValue** | **string** | **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式  | 

### Return type

[**CreditCardPeriodInfo**](CreditCardPeriodInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

