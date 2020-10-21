# ECPay SDK for Golang
[![license](https://img.shields.io/github/license/Laysi/go-ecpay-sdk)](https://github.com/Laysi/go-ecpay-sdk/blob/master/LICENSE)
![go-version](https://img.shields.io/github/go-mod/go-version/Laysi/go-ecpay-sdk)
![latest-version](https://img.shields.io/github/v/tag/Laysi/go-ecpay-sdk?label=Latest%20Version)
[![test-ci](https://github.com/Laysi/go-ecpay-sdk/workflows/Test%20CI/badge.svg?branch=master&event=push)](https://github.com/Laysi/go-ecpay-sdk/actions?query=workflow%3A%22Test+CI%22)
![stability-wip](https://img.shields.io/badge/Stability-work_in_progress-lightgrey.svg)
<!--![docs](https://img.shields.io/badge/Docs-outdated-red)-->

## Description
ECPay SDK for Golang,with some helpers package.

## Feature
- Configure
  - Merchant ID
  - Hash Key
  - Hash IV
  - Return URL
  - Period Return Url
- API
  - AioCheckOut Query Builder
  - QueryCreditCardPeriodInfo
  - QueryTradeInfo
- Gin CheckMacValue Validator Handler
- Predefined Model
- OpenAPI Definition

## Init client
```
import (
	"github.com/Laysi/go-ecpay-sdk"
	"github.com/Laysi/go-ecpay-sdk/base"
)
```
### for production
```go
client := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", "<RETURN_URL>")
```
### for staging
```go
client := ecpay.NewStageClient(ecpay.WithReturnURL("<RETURN_URL>"))
```

### options

```go
client = ecpay.NewStageClient(
    ecpay.WithReturnURL("<RETURN_URL>"), 
    ecpay.WithPeriodReturnURL("<PERIOD_RETURN_URL>"),
    ecpay.WithClientBackURL("<CLIENT_BACK_URL>"),
    ecpay.WithClientRedirectURL("<CLIENT_REDIRECT_URL>"),
    ecpay.WithOrderResultURL("<ORDER_RESULT_URL>"),
    ecpay.WithPaymentInfoURL("<PAYMENT_INFO_URL>"),
    ecpay.WithPlatformID("<PLATFORM_ID>"),
    ecpay.WithCtxFunc(func(c context.Context) context.Context {
        return context.WithValue(c, "<KEY>", "<VALUE>") // Add context operation for every request pass to the api client
    }),
    ecpay.WithDebug,  // Turn on debug log
)
```

## Usage
### Create Order

```go
html := client.CreateOrder("<MerchantTradeNo>", time.Now(), 1000, "<Description>", []string{"<ItemName1>", "<ItemName2>"}).
    WithOptional(ecpay.AioCheckOutGeneralOptional{
        StoreID:           "<StoreID>",
        ItemURL:           "<ItemURL>",
        Remark:            "<Remark>",
        ChooseSubPayment:  ecpayBase.CHOOSESUBPAYMENTENUM_EMPTY,
        NeedExtraPaidInfo: ecpayBase.NEEDEXTRAPAIDINFOENUM_N,
        PlatformID:        "<PlatformID>",
        CustomField1:      "<CustomField1>",
        CustomField2:      "<CustomField2>",
        CustomField3:      "<CustomField3>",
        CustomField4:      "<CustomField4>",
        Language:          ecpayBase.LANGUAGEENUM_ENG,
    }).
    SetAllPayment(ecpayBase.CHOOSEPAYMENTENUM_ATM, ecpayBase.CHOOSEPAYMENTENUM_CREDIT).
    WithCreditOptional(ecpay.AioCheckOutCreditOptional{
        BindingCard:      ecpayBase.BINDINGCARDENUM_BINDING,
        MerchantMemberID: "<MerchantMemberID>",
    }).
    WithCreditOnetimeOptional(ecpay.AioCheckOutCreditOnetimeOptional{
        Redeem:   ecpayBase.REDEEMENUM_Y,
        UnionPay: ecpayBase.UNIONPAYENUM_HIDDEN,
    }).
    WithCreditInstallmentOptional(3, 12, 18).
    WithCreditPeriodOptional(ecpayBase.CREDITPERIODTYPEENUM_DAY, 1, 0).
    WithAtmOptional(5).
    WithCvsBarcodeOptional(ecpay.AioCheckOutCvsBarcodeOptional{
        StoreExpireDate: 5,
        Desc1:           "<Desc1>",
        Desc2:           "<Desc1>",
        Desc3:           "<Desc1>",
        Desc4:           "<Desc1>",
    }).
    GenerateRequestHtml()
```

### QueryCreditCardPeriodInfo

```go
info, resp, err := client.QueryCreditCardPeriodInfo("<MerchantTradeNo>", time.Now())

```

### QueryTradeInfo

```go
info, resp, err := client.QueryTradeInfo("<MerchantTradeNo>", time.Now())
```

## Gin Helper
### Check Mac Validator Handler
This handler can use to check the data CheckMacValue from the ECPAY server,and will return `0|Error` with 400 status if validating failed.

```go
server := gin.Default()
var ecpayClient = ecpay.NewStageClient(ecpay.WithReturnURL("<RETURN_URL>"))
server.POST("/ecpay/return",ecpayGin.ECPayCheckMacValueHandler(ecpayClient), func(c *gin.Context) {...})
```

### Ecpay Data Binding Patch Helper
Due to [gin-gonic/gin#2510](https://github.com/gin-gonic/gin/issues/2510),there have some bug to binding the datetime field of data from ecpay server in gin's current version.
So there is a helper function to deal with it as a temporary workaround.


```go
server := gin.Default()
server.POST("/ecpay/return",ecpayCheckMacValueHandler, func(c *gin.Context) {
    data := ecpayBase.OrderResult{}
    err := ecpayGin.ResponseBodyDateTimePatchHelper(c)
    if err != nil {
        fmt.Println(err.Error())
        c.Status(http.StatusInternalServerError)
        return
    }
    if err = c.MustBindWith(&data, binding.FormPost); err != nil {
        fmt.Println(err.Error())
        return
    }
})

server.POST("/ecpay/period",ecpayCheckMacValueHandler, func(c *gin.Context) {
    data := ecpayBase.PeriodOrderResult{}
    err := ecpayGin.ResponseBodyDateTimePatchHelper(c)

    if err != nil {
        fmt.Println(err.Error())
        c.Status(http.StatusInternalServerError)
        return
    }
    if err := c.MustBindWith(&data, binding.FormPost); err != nil {
        fmt.Println(err.Error())
        return
    }
})
```