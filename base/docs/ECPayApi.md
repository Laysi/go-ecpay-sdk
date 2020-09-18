# \ECPayApi

All URIs are relative to *https://payment.ecpay.com.tw*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AioCheckOut**](ECPayApi.md#AioCheckOut) | **Post** /Cashier/AioCheckOut/V5 | 



## AioCheckOut

> AioCheckOut(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ECPayApi.AioCheckOut(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ECPayApi.AioCheckOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAioCheckOutRequest struct via the builder pattern


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

