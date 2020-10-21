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
client := ecpay.NewClient("<MERCHANT_ID>", "<HASH_KEY>", "<HASH_IV>", "<RETURN_URL>")
```
### for staging
```go
client := ecpay.NewStageClient(ecpay.WithReturnURL("https://example.com/path/to/ecpay/result"))
```
## Create Order

```go
html := client.CreateOrder("<TradeNO>", time.Now(), 1000, "<Description>", []string{"<ItemName1>", "<ItemName2>"}).
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
