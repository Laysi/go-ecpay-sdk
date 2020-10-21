package ecpay

import (
	"github.com/Laysi/go-ecpay-sdk/base"
	"strconv"
	"strings"
	"time"
)

type AioOrderRequest struct {
	ecpayBase.AioCheckOutGeneralOption
	*ecpayBase.AioCheckOutAtmOption
	*ecpayBase.AioCheckOutCvsBarcodeOption
	*ecpayBase.AioCheckOutCreditOption
	*ecpayBase.AioCheckOutCreditOnetimeOption
	*ecpayBase.AioCheckOutCreditInstallmentOption
	*ecpayBase.AioCheckOutCreditPeriodOption
	*ecpayBase.AioCheckOutInvoiceOption
}

type AioOrderRequestWithClient struct {
	*AioOrderRequest
	client Client
}

func (c Client) CreateOrder(tradeNo string, tradeDate time.Time, amount int, description string, itemNames []string) *AioOrderRequestWithClient {
	r := &AioOrderRequestWithClient{
		AioOrderRequest: &AioOrderRequest{
			AioCheckOutGeneralOption: ecpayBase.AioCheckOutGeneralOption{
				MerchantID:        c.merchantID,
				EncryptType:       ecpayBase.ENCRYPTTYPEENUM_SHA256,
				MerchantTradeNo:   tradeNo,
				MerchantTradeDate: ecpayBase.ECPayDateTime(tradeDate),
				PaymentType:       ecpayBase.AIOCHECKPAYMENTTYPEENUM_AIO,
				TotalAmount:       amount,
				TradeDesc:         description,
				ItemName:          strings.Join(itemNames, "#"),
				ReturnURL:         c.returnURL,
				ChoosePayment:     "",
			},
		},
		client: c,
	}

	r.AioCheckOutGeneralOption.OrderResultURL = r.client.orderResultURL
	r.AioCheckOutGeneralOption.ClientBackURL = r.client.clientBackURL
	return r
}

func (r *AioOrderRequestWithClient) WithOptional(optional AioCheckOutGeneralOptional) *AioOrderRequestWithClient {
	r.AioCheckOutGeneralOption.ChooseSubPayment = (*ecpayBase.ChooseSubPaymentEnum)(PtrNilString(string(optional.ChooseSubPayment)))
	r.AioCheckOutGeneralOption.CustomField1 = PtrNilString(optional.CustomField1)
	r.AioCheckOutGeneralOption.CustomField2 = PtrNilString(optional.CustomField2)
	r.AioCheckOutGeneralOption.CustomField3 = PtrNilString(optional.CustomField3)
	r.AioCheckOutGeneralOption.CustomField4 = PtrNilString(optional.CustomField4)
	r.AioCheckOutGeneralOption.ItemURL = PtrNilString(optional.ItemURL)
	r.AioCheckOutGeneralOption.PlatformID = PtrNilString(optional.PlatformID)
	r.AioCheckOutGeneralOption.Remark = PtrNilString(optional.Remark)
	r.AioCheckOutGeneralOption.StoreID = PtrNilString(optional.StoreID)
	r.AioCheckOutGeneralOption.Language = (*ecpayBase.LanguageEnum)(PtrNilString(string(optional.Language)))
	r.AioCheckOutGeneralOption.NeedExtraPaidInfo = (*ecpayBase.NeedExtraPaidInfoEnum)(PtrNilString(string(optional.NeedExtraPaidInfo)))
	return r
}

func (r *AioOrderRequestWithClient) SetAtmPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_ATM
	r.AioCheckOutAtmOption = &ecpayBase.AioCheckOutAtmOption{}
	r.AioCheckOutAtmOption.ClientRedirectURL = r.client.clientRedirectURL
	r.AioCheckOutAtmOption.PaymentInfoURL = r.client.paymentInfoURL
	return r
}

// expireDate: **允許繳費有效天數**   若需設定最長 60 天，最短 1 天。   未設定此參數則預設為 3 天   注意事項：以天為單位
func (r *AioOrderRequestWithClient) WithAtmOptional(expireDate int) *AioOrderRequestWithClient {
	r.AioCheckOutAtmOption.ExpireDate = ecpayBase.PtrInt(expireDate)
	return r
}

func (r *AioOrderRequestWithClient) SetCvsPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_CVS
	r.AioCheckOutCvsBarcodeOption = &ecpayBase.AioCheckOutCvsBarcodeOption{}
	r.AioCheckOutCvsBarcodeOption.PaymentInfoURL = r.client.paymentInfoURL
	r.AioCheckOutCvsBarcodeOption.ClientRedirectURL = r.client.clientRedirectURL
	return r
}

func (r *AioOrderRequestWithClient) SetBarcodePayment() *AioOrderRequestWithClient {
	r.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_BARCODE
	r.AioCheckOutCvsBarcodeOption = &ecpayBase.AioCheckOutCvsBarcodeOption{}
	r.AioCheckOutCvsBarcodeOption.PaymentInfoURL = r.client.paymentInfoURL
	r.AioCheckOutCvsBarcodeOption.ClientRedirectURL = r.client.clientRedirectURL
	return r
}

func (r *AioOrderRequestWithClient) WithCvsBarcodeOptional(option AioCheckOutCvsBarcodeOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCvsBarcodeOption.Desc1 = option.Desc1
	r.AioCheckOutCvsBarcodeOption.Desc2 = option.Desc2
	r.AioCheckOutCvsBarcodeOption.Desc3 = option.Desc3
	r.AioCheckOutCvsBarcodeOption.Desc4 = option.Desc4
	r.AioCheckOutCvsBarcodeOption.StoreExpireDate = option.StoreExpireDate
	return r
}

func (r *AioOrderRequestWithClient) SetCreditPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_CREDIT
	r.AioCheckOutCreditOption = &ecpayBase.AioCheckOutCreditOption{}
	return r
}

func (r *AioOrderRequestWithClient) WithCreditOptional(option AioCheckOutCreditOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCreditOption.BindingCard = (*ecpayBase.BindingCardEnum)(PtrNilInt(int(option.BindingCard)))
	r.AioCheckOutCreditOption.MerchantMemberID = PtrNilString(option.MerchantMemberID)
	return r
}

func (r *AioOrderRequestWithClient) WithCreditOnetimeOptional(option AioCheckOutCreditOnetimeOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCreditOnetimeOption = &ecpayBase.AioCheckOutCreditOnetimeOption{
		Redeem:   option.Redeem,
		UnionPay: option.UnionPay,
	}
	return r
}

type IntSliceConverter []int

func (i IntSliceConverter) ToStringSlice() []string {
	var result []string
	for _, val := range i {
		result = append(result, strconv.Itoa(val))
	}
	return result
}

// 信用卡分期可用參數為:3,6,12,18,24
func (r *AioOrderRequestWithClient) WithCreditInstallmentOptional(installments ...int) *AioOrderRequestWithClient {
	r.AioCheckOutCreditInstallmentOption = &ecpayBase.AioCheckOutCreditInstallmentOption{
		CreditInstallment: strings.Join(IntSliceConverter(installments).ToStringSlice(), ","),
	}
	return r
}

func (r *AioOrderRequestWithClient) WithCreditPeriodOptional(periodType ecpayBase.CreditPeriodTypeEnum, frequency int, execTimes int) *AioOrderRequestWithClient {
	r.AioCheckOutCreditPeriodOption = &ecpayBase.AioCheckOutCreditPeriodOption{
		PeriodAmount:    r.TotalAmount,
		PeriodType:      periodType,
		Frequency:       frequency,
		ExecTimes:       execTimes,
		PeriodReturnURL: r.client.periodReturnURL,
	}
	return r
}

func (r *AioOrderRequestWithClient) SetAllPayment(ignorePayment ...ecpayBase.ChoosePaymentEnum) *AioOrderRequestWithClient {
	r = r.SetCreditPayment().
		SetAtmPayment().
		SetCvsPayment().
		SetBarcodePayment()
	r.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_ALL

	var ignorePaymentValues []string
	for _, v := range ignorePayment {
		ignorePaymentValues = append(ignorePaymentValues, string(v))
	}
	r.IgnorePayment = ecpayBase.PtrString(strings.Join(ignorePaymentValues, "#"))
	return r
}

type InvoiceItem struct {
	// Name **商品名稱**
	Name string
	// Count **商品數量**
	Count int
	// Word **商品單位**
	Word string
	// Price **商品價格**
	Price int
	// TaxType **商品課稅別** 當課稅類別 [TaxType] = 9 (MIXED) 時，此欄位不可為空
	TaxType *ecpayBase.TaxTypeEnum
}

func (r *AioOrderRequestWithClient) SetInvoice(relateNumber string, taxType ecpayBase.TaxTypeEnum, donation ecpayBase.InvoiceDonationEunm, print ecpayBase.InvoicePrintEnum, items []InvoiceItem, delay int, invType string) *AioOrderRequestWithClient {
	r.AioCheckOutGeneralOption.InvoiceMark = ecpayBase.INVOICEMARKENUM_Y.Ptr()
	var itemsName []string
	var itemsCount []string
	var itemsWord []string
	var itemsPrice []string
	for _, item := range items {
		itemsName = append(itemsName, item.Name)
		itemsCount = append(itemsCount, strconv.Itoa(item.Count))
		itemsWord = append(itemsWord, item.Word)
		itemsPrice = append(itemsPrice, strconv.Itoa(item.Price))
	}
	r.AioCheckOutInvoiceOption = &ecpayBase.AioCheckOutInvoiceOption{
		RelateNumber:     relateNumber,
		TaxType:          taxType,
		Donation:         donation,
		Print:            print,
		InvoiceItemName:  strings.Join(itemsName, "|"),
		InvoiceItemCount: strings.Join(itemsCount, "|"),
		InvoiceItemWord:  strings.Join(itemsWord, "|"),
		InvoiceItemPrice: strings.Join(itemsPrice, "|"),
		DelayDay:         delay,
		InvType:          invType,
	}
	if taxType == ecpayBase.TAXTYPEENUM_MIXED {
		var itemsTaxType []string
		for _, item := range items {
			if taxType == ecpayBase.TAXTYPEENUM_MIXED {
				itemsTaxType = append(itemsTaxType, string(*item.TaxType))

			}
		}
		r.AioCheckOutInvoiceOption.InvoiceItemTaxType = ecpayBase.PtrString(strings.Join(itemsTaxType, "|"))

	}
	return r
}

func (r *AioOrderRequestWithClient) WithInvoiceOptional(option AioCheckOutInvoiceOptional) *AioOrderRequestWithClient {
	r.AioCheckOutInvoiceOption.CustomerID = option.CustomerID
	r.AioCheckOutInvoiceOption.CustomerIdentifier = option.CustomerIdentifier
	r.AioCheckOutInvoiceOption.CustomerName = option.CustomerName
	r.AioCheckOutInvoiceOption.CustomerAddr = option.CustomerAddr
	r.AioCheckOutInvoiceOption.CustomerPhone = option.CustomerPhone
	r.AioCheckOutInvoiceOption.CustomerEmail = option.CustomerEmail
	r.AioCheckOutInvoiceOption.ClearanceMark = option.ClearanceMark
	r.AioCheckOutInvoiceOption.CarruerType = option.CarruerType
	r.AioCheckOutInvoiceOption.CarruerNum = option.CarruerNum
	r.AioCheckOutInvoiceOption.LoveCode = option.LoveCode
	//r.AioCheckOutInvoiceOption.InvoiceItemTaxType = option.InvoiceItemTaxType
	r.AioCheckOutInvoiceOption.InvoiceRemark = option.InvoiceRemark
	return r
}

func (r *AioOrderRequestWithClient) GenerateUrlQuery() string {
	params := StructToParamsMap(r.AioOrderRequest)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.CheckMacValue = checkMac
	params = StructToParamsMap(r.AioOrderRequest)
	return NewECPayValuesFromMap(params).Encode()

}

func (r *AioOrderRequestWithClient) GenerateCheckMac() string {
	params := StructToParamsMap(r.AioOrderRequest)
	checkMac := r.client.GenerateCheckMacValue(params)
	return checkMac
}

func (r *AioOrderRequestWithClient) GenerateRequestHtml() string {
	params := StructToParamsMap(r.AioOrderRequest)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.CheckMacValue = checkMac
	params = StructToParamsMap(r.AioOrderRequest)
	return r.client.GenerateAutoSubmitHtmlForm(params, r.client.GetCurrentServer()+r.client.aioCheckOutPath)

}
