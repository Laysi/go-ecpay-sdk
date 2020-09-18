# \CallBackApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallbackReturnURLPost**](CallBackApi.md#CallbackReturnURLPost) | **Post** /callback/ReturnURL | 



## CallbackReturnURLPost

> string CallbackReturnURLPost(ctx, merchantID, merchantTradeNo, storeID, rtnCode, rtnMsg, tradeNo, tradeAmt, paymentDate, paymentType, paymentTypeChargeFee, tradeDate, checkMacValue, simulatePaid, optional)



付款結果通知

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantID** | **string**| 特店編號 | 
**merchantTradeNo** | **string**| 特店交易編號 | 
**storeID** | **string**| 特店旗下店舖代號 | 
**rtnCode** | **float32**| 交易狀態 | 
**rtnMsg** | **string**| 交易訊息 | 
**tradeNo** | **string**| 綠界的交易編號 | 
**tradeAmt** | **int32**| 交易金額 | 
**paymentDate** | **string**| 付款時間 | 
**paymentType** | [**PaymentType**](.md)| 特店選擇的付款方式 | 
**paymentTypeChargeFee** | **int32**| 通路費 | 
**tradeDate** | **string**| 訂單成立時間 | 
**checkMacValue** | **string**| 檢查碼 | 
**simulatePaid** | [**EcPaySimulateEnum**](.md)| 是否為模擬付款 | 
 **optional** | ***CallbackReturnURLPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CallbackReturnURLPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------













 **customField1** | **optional.String**| 自訂名稱欄位 1 | 
 **customField2** | **optional.String**| 自訂名稱欄位 2 | 
 **customField3** | **optional.String**| 自訂名稱欄位 3 | 
 **customField4** | **optional.String**| 自訂名稱欄位 4 | 
 **customField5** | **optional.String**| 自訂名稱欄位 5 | 

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

