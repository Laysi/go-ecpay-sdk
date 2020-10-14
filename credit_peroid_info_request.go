package ecpay

import (
	"context"
	"github.com/Laysi/go-ecpay-sdk/base"
	"net/http"
	"time"
)

func defaultContext(c context.Context, contexts ...context.Context) context.Context {
	if len(contexts) == 0 {
		return c
	}
	return contexts[0]
}

func (e ECPayClient) QueryCreditCardPeriodInfo(merchantID, merchantTradeNo string, timeStamp time.Time, cs ...context.Context) (ecpayBase.CreditCardPeriodInfo, *http.Response, error) {
	params := ecpayBase.QueryCreditCardPeriodInfoRequest{
		MerchantID:      merchantID,
		MerchantTradeNo: merchantTradeNo,
		TimeStamp:       int(timeStamp.Unix()),
	}
	params.CheckMacValue = e.GenerateCheckMacValue(StructToParamsMap(params))

	return e.apiClient.ECPayApi.
		QueryCreditCardPeriodInfo(e.WithContext(defaultContext(context.Background(), cs...))).
		MerchantID(params.MerchantID).
		MerchantTradeNo(params.MerchantTradeNo).
		TimeStamp(params.TimeStamp).
		CheckMacValue(params.CheckMacValue).
		Execute()
}
