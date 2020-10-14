# \ECPayApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryCreditCardPeriodInfo**](ECPayApi.md#QueryCreditCardPeriodInfo) | **Post** /Cashier/QueryCreditCardPeriodInfo | 
[**QueryTradeInfo**](ECPayApi.md#QueryTradeInfo) | **Post** /Cashier/QueryTradeInfo/V5 | 



## QueryCreditCardPeriodInfo

> CreditCardPeriodInfo QueryCreditCardPeriodInfo(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).TimeStamp(timeStamp).CheckMacValue(checkMacValue).Execute()





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
    resp, r, err := api_client.ECPayApi.QueryCreditCardPeriodInfo(context.Background(), merchantID, merchantTradeNo, timeStamp, checkMacValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ECPayApi.QueryCreditCardPeriodInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryCreditCardPeriodInfo`: CreditCardPeriodInfo
    fmt.Fprintf(os.Stdout, "Response from `ECPayApi.QueryCreditCardPeriodInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryCreditCardPeriodInfoRequest struct via the builder pattern


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


## QueryTradeInfo

> TradeInfo QueryTradeInfo(ctx).MerchantID(merchantID).MerchantTradeNo(merchantTradeNo).TimeStamp(timeStamp).CheckMacValue(checkMacValue).PlatformID(platformID).Execute()





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
    checkMacValue := "checkMacValue_example" // string | **檢查碼** 特店必須檢查檢查碼`CheckMacValue`來驗證，請參考附錄檢查碼機制。 
    platformID := "platformID_example" // string | **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 `MerchantID`。    (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ECPayApi.QueryTradeInfo(context.Background(), merchantID, merchantTradeNo, timeStamp, checkMacValue).PlatformID(platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ECPayApi.QueryTradeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTradeInfo`: TradeInfo
    fmt.Fprintf(os.Stdout, "Response from `ECPayApi.QueryTradeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTradeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantID** | **string** | **特店編號(由綠界提供)**  | 
 **merchantTradeNo** | **string** | **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。  | 
 **timeStamp** | **int** | **驗證時間**   將當下的時間轉為UnixTimeStamp(見範例)用於驗證此次介接的時間區間。   綠界驗證時間區間暫訂為 3 分鐘內有效，超過則此次介接無效。   | 
 **checkMacValue** | **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 
 **platformID** | **string** | **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 &#x60;MerchantID&#x60;。    | 

### Return type

[**TradeInfo**](TradeInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

