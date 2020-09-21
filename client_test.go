package ecpay_test

import (
	"github.com/Laysi/go-ecpay-sdk"
	"github.com/Laysi/go-ecpay-sdk/base"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
	//"github.com/Laysi/go-ecpay-sdk"
)

var _ = Describe("Client", func() {
	Context("CheckMac", func() {
		It("mac should calculate correct", func() {
			client := ecpay.NewStageECPayClient()
			values := map[string]string{"test": "test"}
			mac := client.GenerateCheckMacValue(values)
			Expect(mac).To(Equal("B5947417B0FB8C1B8EBCA08E2FF3D66B7FB2329C8529C1331FE2268E9430903A"))
		})
	})

	Context("AutoPostTmpl", func() {
		It("shouldn't fail to generate auto post form html", func() {
			client := ecpay.NewStageECPayClient()
			values := map[string]string{
				"test": "test",
				"Test": "C",
				"D":    "E",
				"F":    "G",
			}
			html := client.GenerateAutoSubmitHtmlForm(values, "http://test.test.test/test")
			//print(html)
			Expect(html).To(Equal(`<form id="order_form" action="http://test.test.test/test" method="post">
 <input type="hidden" name="D" id="D" value="E" /> 
<input type="hidden" name="F" id="F" value="G" /> 
<input type="hidden" name="Test" id="Test" value="C" /> 
<input type="hidden" name="test" id="test" value="test" /> 
</form>
<script>$("#order_form").submit();</script>`))
		})
	})

	Context("Order", func() {
		It("should success to create a normal order request", func() {
			client := ecpay.NewStageECPayClient()
			html := client.CreateOrder("r44q2g423gq", time.Time(base.MustParseECPayDateTime("2020/09/21 15:02:01")), 400, "世界好", "你好").
				//WithOptional(ecpay.AioCheckOutGeneralOptional{
				//	PlatformID: base.PtrString(""),
				//}).
				SetCreditPayment().
				GenerateRequestHtml()
			//print(html)
			Expect(html).To(Equal(`<form id="order_form" action="" method="post">
 <input type="hidden" name="CheckMacValue" id="CheckMacValue" value="9EB83EF0F2CB264252E1BCEEE7630D99BF6129941E63BF7BE5862ECA39F9062C" /> 
<input type="hidden" name="ChoosePayment" id="ChoosePayment" value="Credit" /> 
<input type="hidden" name="EncryptType" id="EncryptType" value="1" /> 
<input type="hidden" name="ItemName" id="ItemName" value="你好" /> 
<input type="hidden" name="MerchantID" id="MerchantID" value="2000132" /> 
<input type="hidden" name="MerchantTradeDate" id="MerchantTradeDate" value="2020/09/21 15:02:01" /> 
<input type="hidden" name="MerchantTradeNo" id="MerchantTradeNo" value="r44q2g423gq" /> 
<input type="hidden" name="PaymentType" id="PaymentType" value="aio" /> 
<input type="hidden" name="ReturnURL" id="ReturnURL" value="" /> 
<input type="hidden" name="TotalAmount" id="TotalAmount" value="400" /> 
<input type="hidden" name="TradeDesc" id="TradeDesc" value="世界好" /> 
</form>
<script>$("#order_form").submit();</script>`))
		})
	})
})
