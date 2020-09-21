package ecpay

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Laysi/go-ecpay-sdk/base"
	"html/template"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ECPayClient struct {
	MerchantID     string
	AioCheckOutUrl string
	ReturnURL      string
	HashKey        string
	HashIV         string
}

func NewECPayClient(merchantID string, returnURL string, hashIV string, hashKey string) *ECPayClient {
	base.NewAPIClient(base.NewConfiguration())

	return &ECPayClient{MerchantID: merchantID, ReturnURL: returnURL, HashKey: hashIV, HashIV: hashKey}
}

func NewStageECPayClient() *ECPayClient {
	return &ECPayClient{MerchantID: "2000132", AioCheckOutUrl: "https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5", HashKey: "5294y06JbISpM5x9", HashIV: "v77hoKGq4kWxNNIS"}
}

var OrderTemplateText = `<form id="order_form" action="{{.Action}}" method="post">
 {{range $key,$element := .Values -}} 
	<input type="hidden" name="{{$key}}" id="{{$key}}" value="{{$element}}" /> 
{{end -}}
</form>
<script>$("#order_form").submit();</script>`

type OrderTmplArgs struct {
	Values map[string]string
	Action string
}

var OrderTmpl = template.Must(template.New("AutoPostOrder").Parse(OrderTemplateText))

func (c ECPayClient) GenerateAutoSubmitHtmlForm(params map[string]string, targetUrl string) string {

	var result bytes.Buffer
	err := OrderTmpl.Execute(&result, OrderTmplArgs{
		Values: params,
		Action: targetUrl,
	})
	if err != nil {
		panic(err)
	}
	return result.String()
}

type AioOrderRequest struct {
	base.AioCheckOutGeneralOption
	*base.AioCheckOutAtmOption
	*base.AioCheckOutCvsBarcodeOption
	*base.AioCheckOutCreditOption
	*base.AioCheckOutCreditOnetimeOption
	*base.AioCheckOutCreditInstallmentOption
	*base.AioCheckOutCreditPeriodOption
	*base.AioCheckOutInvoiceOption
}

type AioOrderRequestWithClient struct {
	*AioOrderRequest
	client ECPayClient
}

func (c ECPayClient) CreateOrder(tradeNo string, tradeDate time.Time, amount int, description string, itemName string) *AioOrderRequestWithClient {
	return &AioOrderRequestWithClient{
		AioOrderRequest: &AioOrderRequest{
			AioCheckOutGeneralOption: base.AioCheckOutGeneralOption{
				MerchantID:        c.MerchantID,
				EncryptType:       base.ENCRYPTTYPEENUM_SHA256,
				MerchantTradeNo:   tradeNo,
				MerchantTradeDate: base.ECPayDateTime(tradeDate),
				PaymentType:       base.AIOCHECKPAYMENTTYPEENUM_AIO,
				TotalAmount:       amount,
				TradeDesc:         description,
				ItemName:          itemName,
				ReturnURL:         c.ReturnURL,
				ChoosePayment:     "",
			},
		},
		client: c,
	}
}

func (r *AioOrderRequestWithClient) WithOptional(optional AioCheckOutGeneralOptional) *AioOrderRequestWithClient {
	r.AioCheckOutGeneralOption.ChooseSubPayment = optional.ChooseSubPayment
	r.AioCheckOutGeneralOption.ClientBackURL = optional.ClientBackURL
	r.AioCheckOutGeneralOption.CustomField1 = optional.CustomField1
	r.AioCheckOutGeneralOption.CustomField2 = optional.CustomField2
	r.AioCheckOutGeneralOption.CustomField3 = optional.CustomField3
	r.AioCheckOutGeneralOption.CustomField4 = optional.CustomField4
	r.AioCheckOutGeneralOption.IgnorePayment = optional.IgnorePayment
	r.AioCheckOutGeneralOption.ItemURL = optional.ItemURL
	r.AioCheckOutGeneralOption.OrderResultURL = optional.OrderResultURL
	r.AioCheckOutGeneralOption.PlatformID = optional.PlatformID
	r.AioCheckOutGeneralOption.Remark = optional.Remark
	r.AioCheckOutGeneralOption.StoreID = optional.StoreID
	r.AioCheckOutGeneralOption.Language = optional.Language
	r.AioCheckOutGeneralOption.NeedExtraPaidInfo = optional.NeedExtraPaidInfo
	return r
}

func (r *AioOrderRequestWithClient) SetAtmPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_ATM
	r.AioCheckOutAtmOption = &base.AioCheckOutAtmOption{}
	return r
}

func (r *AioOrderRequestWithClient) WithAtmOptional(option AioCheckOutAtmOptional) *AioOrderRequestWithClient {
	r.AioCheckOutAtmOption.ExpireDate = option.ExpireDate
	r.AioCheckOutAtmOption.ClientRedirectURL = option.ClientRedirectURL
	r.AioCheckOutAtmOption.PaymentInfoURL = option.PaymentInfoURL
	return r
}

func (r *AioOrderRequestWithClient) SetCvsPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_CVS
	r.AioCheckOutCvsBarcodeOption = &base.AioCheckOutCvsBarcodeOption{}
	return r
}

func (r *AioOrderRequestWithClient) SetBarcodePayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_BARCODE
	r.AioCheckOutCvsBarcodeOption = &base.AioCheckOutCvsBarcodeOption{}
	return r
}

func (r *AioOrderRequestWithClient) WithCvsBarcodeOptional(option AioCheckOutCvsBarcodeOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCvsBarcodeOption.PaymentInfoURL = option.PaymentInfoURL
	r.AioCheckOutCvsBarcodeOption.ClientRedirectURL = option.ClientRedirectURL
	r.AioCheckOutCvsBarcodeOption.Desc1 = option.Desc1
	r.AioCheckOutCvsBarcodeOption.Desc2 = option.Desc2
	r.AioCheckOutCvsBarcodeOption.Desc3 = option.Desc3
	r.AioCheckOutCvsBarcodeOption.Desc4 = option.Desc4
	r.AioCheckOutCvsBarcodeOption.StoreExpireDate = option.StoreExpireDate
	return r
}

func (r *AioOrderRequestWithClient) SetCreditPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_CREDIT
	r.AioCheckOutCreditOption = &base.AioCheckOutCreditOption{}
	return r
}

func (r *AioOrderRequestWithClient) WithCreditOptional(option AioCheckOutCreditOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCreditOption.BindingCard = option.BindingCard
	r.AioCheckOutCreditOption.MerchantMemberID = option.MerchantMemberID
	return r
}

func (r *AioOrderRequestWithClient) WithCreditOnetimeOptional(option AioCheckOutCreditOnetimeOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCreditOnetimeOption = &base.AioCheckOutCreditOnetimeOption{
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

func (r *AioOrderRequestWithClient) WithCreditInstallmentOptional(installments []int) *AioOrderRequestWithClient {
	r.AioCheckOutCreditInstallmentOption = &base.AioCheckOutCreditInstallmentOption{
		CreditInstallment: strings.Join(IntSliceConverter(installments).ToStringSlice(), ","),
	}
	return r
}

func (r *AioOrderRequestWithClient) WithCreditPeriodOptional(periodAmount int, periodType base.CreditPeriodTypeEnum, frequency int, execTimes int, option AioCheckOutCreditPeriodOptional) *AioOrderRequestWithClient {
	r.AioCheckOutCreditPeriodOption = &base.AioCheckOutCreditPeriodOption{
		PeriodAmount:    periodAmount,
		PeriodType:      periodType,
		Frequency:       frequency,
		ExecTimes:       execTimes,
		PeriodReturnURL: option.PeriodReturnURL,
	}
	return r
}

func (r *AioOrderRequestWithClient) SetAllPayment() *AioOrderRequestWithClient {
	r.ChoosePayment = base.CHOOSEPAYMENTENUM_ALL
	return r
}

type InvoiceItem struct {
	Name    string
	Count   int
	Word    string
	Price   int
	TaxType *base.TaxTypeEnum
}

func (r *AioOrderRequestWithClient) SetInvoice(relateNumber string, taxType base.TaxTypeEnum, donation base.InvoiceDonationEunm, print base.InvoicePrintEnum, items []InvoiceItem, delay int, invType string) *AioOrderRequestWithClient {
	r.AioCheckOutGeneralOption.InvoiceMark = base.INVOICEMARKENUM_Y.Ptr()
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
	r.AioCheckOutInvoiceOption = &base.AioCheckOutInvoiceOption{
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
	if taxType == base.TAXTYPEENUM_MIXED {
		var itemsTaxType []string
		for _, item := range items {
			if taxType == base.TAXTYPEENUM_MIXED {
				itemsTaxType = append(itemsTaxType, string(*item.TaxType))

			}
		}
		r.AioCheckOutInvoiceOption.InvoiceItemTaxType = base.PtrString(strings.Join(itemsTaxType, "|"))

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

func (c ECPayClient) GenerateCheckMacValue(params map[string]string) string {
	delete(params, "CheckMacValue")
	encodedParams := NewECPayValuesFromMap(params).Encode()
	encodedParams = fmt.Sprintf("HashKey=%s&%s&HashIV=%s", c.HashKey, encodedParams, c.HashIV)
	encodedParams = url.QueryEscape(encodedParams)
	encodedParams = strings.ReplaceAll(encodedParams, "%2d", "-")
	encodedParams = strings.ReplaceAll(encodedParams, "%5f", "_")
	encodedParams = strings.ReplaceAll(encodedParams, "%2e", ".")
	encodedParams = strings.ReplaceAll(encodedParams, "%21", "!")
	encodedParams = strings.ReplaceAll(encodedParams, "%2a", "*")
	encodedParams = strings.ReplaceAll(encodedParams, "%28", "(")
	encodedParams = strings.ReplaceAll(encodedParams, "%29", ")")
	encodedParams = strings.ToLower(encodedParams)
	//fmt.Println(encodedParams)
	sum := sha256.Sum256([]byte(encodedParams))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))
	return checkMac
}

func (r *AioOrderRequestWithClient) GenerateRequestHtml() string {
	params := StructToParamsMap(r.AioOrderRequest)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.CheckMacValue = checkMac
	params = StructToParamsMap(r.AioOrderRequest)
	return r.client.GenerateAutoSubmitHtmlForm(params, "")

}

func (r *AioOrderRequestWithClient) GenerateUrlQuery() string {
	params := StructToParamsMap(r.AioOrderRequest)
	checkMac := r.client.GenerateCheckMacValue(params)
	r.CheckMacValue = checkMac
	params = StructToParamsMap(r.AioOrderRequest)
	return NewECPayValuesFromMap(params).Encode()

}

func StructToParamsMap(data interface{}) map[string]string {
	params := map[string]string{}
	iVal := reflect.ValueOf(data)
	iTyp := reflect.TypeOf(data)
	if iVal.Kind() == reflect.Ptr {
		iVal = iVal.Elem()
	}
	if iTyp.Kind() == reflect.Ptr {
		iTyp = iTyp.Elem()
	}

	//stringerInterface := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	marshalerInterface := reflect.TypeOf((*json.Marshaler)(nil)).Elem()
	//typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		ft := iTyp.Field(i)

		if f.Kind() == reflect.Ptr && f.IsNil() {
			continue
		}
		if f.Kind() == reflect.Ptr {
			f = f.Elem()

		}

		if ft.Anonymous {
			nestedParams := StructToParamsMap(f.Interface())
			for key, val := range nestedParams {
				params[key] = val
			}
			continue
		}

		var v string
		//f.Anonymous
		switch realVal := f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(realVal)
		case string:
			v = realVal
		//case time.Time:
		//	v = realVal.Format(time.RFC3339)
		//case base.ECPayDateTime:
		//	v = realVal.String()
		default:
			switch {
			case f.Type().Implements(marshalerInterface):
				data, err := f.Interface().(json.Marshaler).MarshalJSON()
				if err != nil {
					panic(err)
				}
				unquoteData, err := strconv.Unquote(string(data))
				if err != nil {
					v = string(data)
				}
				v = unquoteData
			//case f.Type().Implements(stringerInterface):
			//	v = f.Interface().(fmt.Stringer).String()
			default:
				switch f.Kind() {
				case reflect.String:
					v = f.String()
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					v = strconv.FormatInt(f.Int(), 10)
				default:
					panic(fmt.Sprintf("Unknown type %T during transfer struct to map!\n", f.Interface()))

				}
			}
		}
		params[ft.Name] = v
	}
	return params
}
