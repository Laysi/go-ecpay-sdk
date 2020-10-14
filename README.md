# ECPay SDK for Golang
[![license](https://img.shields.io/github/license/Laysi/go-ecpay-sdk)](https://github.com/Laysi/go-ecpay-sdk/blob/master/LICENSE)
![go-version](https://img.shields.io/github/go-mod/go-version/Laysi/go-ecpay-sdk)
![latest-version](https://img.shields.io/github/v/tag/Laysi/go-ecpay-sdk?label=Latest%20Version)
[![test-ci](https://github.com/Laysi/go-ecpay-sdk/workflows/Test%20CI/badge.svg?branch=master&event=push)](https://github.com/Laysi/go-ecpay-sdk/actions?query=workflow%3A%22Test+CI%22)
![stability-wip](https://img.shields.io/badge/Stability-work_in_progress-lightgrey.svg)


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
  - CreditCardPeriodInfo
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
client := ecpay.NewClient("<MERCHANT_ID>","<RETURN_URL>","<PERIOD_RETURN_URL>","<HASH_KEY>","<HASH_IV>")
```
### for staging
```go
client := ecpay.NewStageClient()
```
## Create Order

```go
html := client.CreateOrder("<TradeNO>", time.Now(), 1000, "<Description>", []string{"<ItemName1>","<ItemName2>"}).
    WithOptional(ecpay.AioCheckOutGeneralOptional{
        StoreID:           ecpayBase.PtrString("<StoreID>"),
        //ClientBackURL:     nil,
        ItemURL:           ecpayBase.PtrString("<ItemURL>"),
        Remark:            ecpayBase.PtrString("<Remark>"),
        ChooseSubPayment:  ecpayBase.CHOOSESUBPAYMENTENUM_EMPTY.Ptr(),
        //OrderResultURL:    nil,
        NeedExtraPaidInfo: ecpayBase.NEEDEXTRAPAIDINFOENUM_N.Ptr(),
        PlatformID:   ecpayBase.PtrString("<PlatformID>"),
        CustomField1: ecpayBase.PtrString("<CustomField1>"),
        CustomField2: ecpayBase.PtrString("<CustomField2>"),
        CustomField3: ecpayBase.PtrString("<CustomField3>"),
        CustomField4: ecpayBase.PtrString("<CustomField4>"),
        Language:     ecpayBase.LANGUAGEENUM_ENG.Ptr(),
    }).
    SetAllPayment(ecpayBase.CHOOSEPAYMENTENUM_ATM, ecpayBase.CHOOSEPAYMENTENUM_CREDIT).
    WithCreditOptional(ecpay.AioCheckOutCreditOptional{
        BindingCard:      ecpayBase.BINDINGCARDENUM_BINDING.Ptr(),
        MerchantMemberID: ecpayBase.PtrString("<MerchantMemberID>"),
    }).
    WithCreditOnetimeOptional(ecpay.AioCheckOutCreditOnetimeOptional{
        Redeem:   ecpayBase.REDEEMENUM_Y.Ptr(),
        UnionPay: ecpayBase.UNIONPAYENUM_HIDDEN.Ptr(),
    }).
    WithCreditInstallmentOptional([]int{4}).
    WithCreditPeriodOptional(ecpayBase.CREDITPERIODTYPEENUM_DAY, 1, 0).
    WithAtmOptional(ecpay.AioCheckOutAtmOptional{
        ExpireDate:        ecpayBase.PtrInt(5),
        //PaymentInfoURL:    nil,
        //ClientRedirectURL: nil,
    }).
    WithCvsBarcodeOptional(ecpay.AioCheckOutCvsBarcodeOptional{
        StoreExpireDate:   ecpayBase.PtrInt(5),
        Desc1:             ecpayBase.PtrString("<Desc1>"),
        Desc2:             ecpayBase.PtrString("<Desc1>"),
        Desc3:             ecpayBase.PtrString("<Desc1>"),
        Desc4:             ecpayBase.PtrString("<Desc1>"),
        //PaymentInfoURL:    nil,
        //ClientRedirectURL: nil,
    }).
    GenerateRequestHtml()
```

## TODO
- [ ] comments docs improved
- [ ] better return data binding solution
- [ ] some explain about ecpay order html usage and progress
- [ ] docs about gin mac validator 
- [ ] add other optional url fields configure setting
  - [ ] ClientBackURL
  - [ ] PaymentInfoURL
  - [ ] ClientRedirectURL
  - [ ] OrderResultURL