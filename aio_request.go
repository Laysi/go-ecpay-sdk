package ecpay

import (
	"github.com/Laysi/go-ecpay-sdk/base"
	"github.com/go-errors/errors"
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
	err     error
	request *AioOrderRequest
	client  Client
}
type AioOrderRequestCreditWithClient struct {
	*AioOrderRequestWithClient
}

type AioOrderRequestAtmWithClient struct {
	*AioOrderRequestWithClient
}

type AioOrderRequestCvsBarcodeWithClient struct {
	*AioOrderRequestWithClient
}

type AioOrderRequestAllWithClient struct {
	*AioOrderRequestWithClient
}

type AioOrderRequestInvoiceWithClient struct {
	*AioOrderRequestWithClient
}

func (c Client) CreateOrder(tradeNo string, tradeDate time.Time, amount int, description string, itemNames []string) *AioOrderRequestWithClient {
	description = strings.ReplaceAll(description, "\n", " ")
	description = strings.ReplaceAll(description, "\r", " ")
	r := &AioOrderRequestWithClient{
		request: &AioOrderRequest{
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

	r.request.AioCheckOutGeneralOption.OrderResultURL = r.client.orderResultURL
	r.request.AioCheckOutGeneralOption.ClientBackURL = r.client.clientBackURL
	return r
}

func (r *AioOrderRequestWithClient) WithOptional(optional AioCheckOutGeneralOptional) *AioOrderRequestWithClient {
	if r.err != nil {
		return r
	}
	r.request.AioCheckOutGeneralOption.ChooseSubPayment = (*ecpayBase.ChooseSubPaymentEnum)(PtrNilString(string(optional.ChooseSubPayment)))
	r.request.AioCheckOutGeneralOption.CustomField1 = PtrNilString(optional.CustomField1)
	r.request.AioCheckOutGeneralOption.CustomField2 = PtrNilString(optional.CustomField2)
	r.request.AioCheckOutGeneralOption.CustomField3 = PtrNilString(optional.CustomField3)
	r.request.AioCheckOutGeneralOption.CustomField4 = PtrNilString(optional.CustomField4)
	r.request.AioCheckOutGeneralOption.ItemURL = PtrNilString(optional.ItemURL)
	r.request.AioCheckOutGeneralOption.PlatformID = PtrNilString(optional.PlatformID)
	r.request.AioCheckOutGeneralOption.Remark = PtrNilString(optional.Remark)
	r.request.AioCheckOutGeneralOption.StoreID = PtrNilString(optional.StoreID)
	r.request.AioCheckOutGeneralOption.Language = (*ecpayBase.LanguageEnum)(PtrNilString(string(optional.Language)))
	r.request.AioCheckOutGeneralOption.NeedExtraPaidInfo = (*ecpayBase.NeedExtraPaidInfoEnum)(PtrNilString(string(optional.NeedExtraPaidInfo)))
	return r
}

// Atm Payment

func (r *AioOrderRequestWithClient) SetAtmPayment() *AioOrderRequestAtmWithClient {
	if r.err != nil {
		return &AioOrderRequestAtmWithClient{r}
	}
	r.request.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_ATM
	r.request.AioCheckOutAtmOption = &ecpayBase.AioCheckOutAtmOption{}
	r.request.AioCheckOutAtmOption.ClientRedirectURL = r.client.clientRedirectURL
	r.request.AioCheckOutAtmOption.PaymentInfoURL = r.client.paymentInfoURL
	return &AioOrderRequestAtmWithClient{r}
}

// expireDate: **允許繳費有效天數**   若需設定最長 60 天，最短 1 天。   未設定此參數則預設為 3 天   注意事項：以天為單位
func (r *AioOrderRequestAtmWithClient) WithAtmOptional(expireDate int) *AioOrderRequestAtmWithClient {
	if r.err != nil {
		return r
	}
	r.request.AioCheckOutAtmOption.ExpireDate = ecpayBase.PtrInt(expireDate)
	return r
}

// WebATM Payment

func (r *AioOrderRequestWithClient) SetWebAtmPayment() *AioOrderRequestWithClient {
	if r.err != nil {
		return r
	}
	r.request.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_WEB_ATM
	return r
}

// Cvs Barcode Payment

func (r *AioOrderRequestWithClient) SetCvsPayment() *AioOrderRequestCvsBarcodeWithClient {
	if r.err != nil {
		return &AioOrderRequestCvsBarcodeWithClient{r}
	}
	r.request.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_CVS
	r.request.AioCheckOutCvsBarcodeOption = &ecpayBase.AioCheckOutCvsBarcodeOption{}
	r.request.AioCheckOutCvsBarcodeOption.PaymentInfoURL = r.client.paymentInfoURL
	r.request.AioCheckOutCvsBarcodeOption.ClientRedirectURL = r.client.clientRedirectURL
	return &AioOrderRequestCvsBarcodeWithClient{r}
}

func (r *AioOrderRequestWithClient) SetBarcodePayment() *AioOrderRequestCvsBarcodeWithClient {
	if r.err != nil {
		return &AioOrderRequestCvsBarcodeWithClient{r}
	}
	r.request.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_BARCODE
	r.request.AioCheckOutCvsBarcodeOption = &ecpayBase.AioCheckOutCvsBarcodeOption{}
	r.request.AioCheckOutCvsBarcodeOption.PaymentInfoURL = r.client.paymentInfoURL
	r.request.AioCheckOutCvsBarcodeOption.ClientRedirectURL = r.client.clientRedirectURL
	return &AioOrderRequestCvsBarcodeWithClient{r}
}

func (r *AioOrderRequestCvsBarcodeWithClient) WithCvsBarcodeOptional(option AioCheckOutCvsBarcodeOptional) *AioOrderRequestCvsBarcodeWithClient {
	if r.err != nil {
		return r
	}
	r.request.AioCheckOutCvsBarcodeOption.Desc1 = PtrNilString(option.Desc1)
	r.request.AioCheckOutCvsBarcodeOption.Desc2 = PtrNilString(option.Desc2)
	r.request.AioCheckOutCvsBarcodeOption.Desc3 = PtrNilString(option.Desc3)
	r.request.AioCheckOutCvsBarcodeOption.Desc4 = PtrNilString(option.Desc4)
	r.request.AioCheckOutCvsBarcodeOption.StoreExpireDate = PtrNilInt(option.StoreExpireDate)
	return r
}

// Credit Payment

func (r *AioOrderRequestWithClient) SetCreditPayment() *AioOrderRequestCreditWithClient {
	if r.err != nil {
		return &AioOrderRequestCreditWithClient{r}
	}
	r.request.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_CREDIT
	r.request.AioCheckOutCreditOption = &ecpayBase.AioCheckOutCreditOption{}
	return &AioOrderRequestCreditWithClient{r}
}

func (r *AioOrderRequestCreditWithClient) WithCreditOptional(option AioCheckOutCreditOptional) *AioOrderRequestCreditWithClient {
	if r.err != nil {
		return r
	}
	r.request.AioCheckOutCreditOption.BindingCard = (*ecpayBase.BindingCardEnum)(PtrNilInt(int(option.BindingCard)))
	r.request.AioCheckOutCreditOption.MerchantMemberID = PtrNilString(option.MerchantMemberID)
	return r
}

func (r *AioOrderRequestCreditWithClient) WithCreditOnetimeOptional(option AioCheckOutCreditOnetimeOptional) *AioOrderRequestCreditWithClient {
	if r.err != nil {
		return r
	}
	r.request.AioCheckOutCreditOnetimeOption = &ecpayBase.AioCheckOutCreditOnetimeOption{
		Redeem:   (*ecpayBase.RedeemEnum)(PtrNilString(string(option.Redeem))),
		UnionPay: (*ecpayBase.UnionPayEnum)(PtrNilInt(int(option.UnionPay))),
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
func (r *AioOrderRequestCreditWithClient) WithCreditInstallmentOptional(installments ...int) *AioOrderRequestCreditWithClient {
	if r.err != nil {
		return r
	}
	r.request.AioCheckOutCreditInstallmentOption = &ecpayBase.AioCheckOutCreditInstallmentOption{
		CreditInstallment: strings.Join(IntSliceConverter(installments).ToStringSlice(), ","),
	}
	return r
}

func (r *AioOrderRequestCreditWithClient) WithCreditPeriodOptional(periodType ecpayBase.CreditPeriodTypeEnum, frequency int, execTimes int) *AioOrderRequestCreditWithClient {
	if r.err != nil {
		return r
	}
	if r.client.periodReturnURL == nil {
		r.err = errors.New("Missing configure of PeriodReturnURL to use Period Credit Payment Option")
		return r
	}
	r.request.AioCheckOutCreditPeriodOption = &ecpayBase.AioCheckOutCreditPeriodOption{
		PeriodAmount:    r.request.TotalAmount,
		PeriodType:      periodType,
		Frequency:       frequency,
		ExecTimes:       execTimes,
		PeriodReturnURL: r.client.periodReturnURL,
	}
	return r
}

// all payment
func (r *AioOrderRequestWithClient) SetAllPayment(ignorePayment ...ecpayBase.ChoosePaymentEnum) *AioOrderRequestAllWithClient {
	if r.err != nil {
		return &AioOrderRequestAllWithClient{r}
	}
	r = r.SetCreditPayment().
		SetAtmPayment().
		SetCvsPayment().
		SetBarcodePayment().AioOrderRequestWithClient
	r.request.ChoosePayment = ecpayBase.CHOOSEPAYMENTENUM_ALL

	var ignorePaymentValues []string
	for _, v := range ignorePayment {
		ignorePaymentValues = append(ignorePaymentValues, string(v))
	}
	if len(ignorePaymentValues) != 0 {
		r.request.IgnorePayment = ecpayBase.PtrString(strings.Join(ignorePaymentValues, "#"))
	}
	return &AioOrderRequestAllWithClient{r}
}

// all options
func (r *AioOrderRequestAllWithClient) WithCreditOptional(option AioCheckOutCreditOptional) *AioOrderRequestAllWithClient {
	return &AioOrderRequestAllWithClient{
		(&AioOrderRequestCreditWithClient{r.AioOrderRequestWithClient}).
			WithCreditOptional(option).
			AioOrderRequestWithClient,
	}
}

func (r *AioOrderRequestAllWithClient) WithCreditOnetimeOptional(option AioCheckOutCreditOnetimeOptional) *AioOrderRequestAllWithClient {
	return &AioOrderRequestAllWithClient{
		(&AioOrderRequestCreditWithClient{r.AioOrderRequestWithClient}).
			WithCreditOnetimeOptional(option).
			AioOrderRequestWithClient,
	}
}

func (r *AioOrderRequestAllWithClient) WithCreditInstallmentOptional(installments ...int) *AioOrderRequestAllWithClient {
	return &AioOrderRequestAllWithClient{
		(&AioOrderRequestCreditWithClient{r.AioOrderRequestWithClient}).
			WithCreditInstallmentOptional(installments...).
			AioOrderRequestWithClient,
	}
}

func (r *AioOrderRequestAllWithClient) WithCreditPeriodOptional(periodType ecpayBase.CreditPeriodTypeEnum, frequency int, execTimes int) *AioOrderRequestAllWithClient {
	return &AioOrderRequestAllWithClient{
		(&AioOrderRequestCreditWithClient{r.AioOrderRequestWithClient}).
			WithCreditPeriodOptional(periodType, frequency, execTimes).
			AioOrderRequestWithClient,
	}
}

func (r *AioOrderRequestAllWithClient) WithAtmOptional(expireDate int) *AioOrderRequestAllWithClient {
	return &AioOrderRequestAllWithClient{
		(&AioOrderRequestAtmWithClient{r.AioOrderRequestWithClient}).
			WithAtmOptional(expireDate).
			AioOrderRequestWithClient,
	}
}

func (r *AioOrderRequestAllWithClient) WithCvsBarcodeOptional(option AioCheckOutCvsBarcodeOptional) *AioOrderRequestAllWithClient {
	return &AioOrderRequestAllWithClient{
		(&AioOrderRequestCvsBarcodeWithClient{r.AioOrderRequestWithClient}).
			WithCvsBarcodeOptional(option).
			AioOrderRequestWithClient,
	}
}

// Invoice

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

func (r *AioOrderRequestWithClient) SetInvoice(relateNumber string, taxType ecpayBase.TaxTypeEnum, donation ecpayBase.InvoiceDonationEunm, print ecpayBase.InvoicePrintEnum, items []InvoiceItem, delay int, invType string) *AioOrderRequestInvoiceWithClient {
	if r.err != nil {
		return &AioOrderRequestInvoiceWithClient{r}
	}
	r.request.AioCheckOutGeneralOption.InvoiceMark = ecpayBase.INVOICEMARKENUM_Y.Ptr()
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
	r.request.AioCheckOutInvoiceOption = &ecpayBase.AioCheckOutInvoiceOption{
		RelateNumber:     relateNumber,
		TaxType:          taxType,
		Donation:         donation,
		Print:            print,
		InvoiceItemName:  FormUrlEncode(strings.Join(itemsName, "|")),
		InvoiceItemCount: strings.Join(itemsCount, "|"),
		InvoiceItemWord:  FormUrlEncode(strings.Join(itemsWord, "|")),
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
		r.request.AioCheckOutInvoiceOption.InvoiceItemTaxType = ecpayBase.PtrString(strings.Join(itemsTaxType, "|"))

	}
	return &AioOrderRequestInvoiceWithClient{r}
}

func (r *AioOrderRequestInvoiceWithClient) WithInvoiceOptional(option AioCheckOutInvoiceOptional) *AioOrderRequestInvoiceWithClient {
	if r.err != nil {
		return r
	}

	//if r.AioCheckOutInvoiceOption == nil{
	//	r.err = errors.New("wrong must use SetInvoice before WithInvoiceOptional to set required parameter properly")
	//	return r
	//}
	r.request.AioCheckOutInvoiceOption.CustomerID = PtrNilString(option.CustomerID)
	r.request.AioCheckOutInvoiceOption.CustomerIdentifier = PtrNilString(option.CustomerIdentifier)
	r.request.AioCheckOutInvoiceOption.CustomerName = PtrNilString(FormUrlEncode(option.CustomerName))
	r.request.AioCheckOutInvoiceOption.CustomerAddr = PtrNilString(FormUrlEncode(option.CustomerAddr))
	r.request.AioCheckOutInvoiceOption.CustomerPhone = PtrNilString(option.CustomerPhone)
	r.request.AioCheckOutInvoiceOption.CustomerEmail = PtrNilString(FormUrlEncode(option.CustomerEmail))
	r.request.AioCheckOutInvoiceOption.ClearanceMark = (*ecpayBase.ClearanceMarkEnum)(PtrNilString(string(option.ClearanceMark)))
	r.request.AioCheckOutInvoiceOption.CarruerType = (*ecpayBase.CarruerTypeEnum)(PtrNilString(string(option.CarruerType)))
	r.request.AioCheckOutInvoiceOption.CarruerNum = PtrNilString(option.CarruerNum)
	r.request.AioCheckOutInvoiceOption.LoveCode = PtrNilString(option.LoveCode)
	//r.AioCheckOutInvoiceOption.InvoiceItemTaxType = PtrNilString(option.InvoiceItemTaxType)
	r.request.AioCheckOutInvoiceOption.InvoiceRemark = PtrNilString(FormUrlEncode(option.InvoiceRemark))
	return r
}

func (r *AioOrderRequestWithClient) GenerateUrlQuery() (string, error) {
	if r.err != nil {
		return "", r.err
	}
	params := StructToParamsMap(r.request)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.request.CheckMacValue = checkMac
	params = StructToParamsMap(r.request)
	return NewECPayValuesFromMap(params).Encode(), nil

}

func (r *AioOrderRequestWithClient) GenerateCheckMac() (string, error) {
	if r.err != nil {
		return "", r.err
	}
	params := StructToParamsMap(r.request)
	checkMac := r.client.GenerateCheckMacValue(params)
	return checkMac, nil
}

func (r *AioOrderRequestWithClient) GenerateRequestHtml() (string, error) {
	if r.err != nil {
		return "", r.err
	}
	params := StructToParamsMap(r.request)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.request.CheckMacValue = checkMac
	params = StructToParamsMap(r.request)
	return r.client.GenerateAutoSubmitHtmlForm(params, r.client.GetCurrentServer()+r.client.aioCheckOutPath), nil

}
