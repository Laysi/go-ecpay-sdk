package ecpay_test

import (
	"github.com/Laysi/go-ecpay-sdk"
	. "github.com/onsi/ginkgo"
	"net/url"

	. "github.com/onsi/gomega"
	//"github.com/Laysi/go-ecpay-sdk"
)

var _ = Describe("Client", func() {
	Context("CheckMac", func() {
		It("mac should calculate correct", func() {
			client := ecpay.NewStageECPayClient()
			values := url.Values{}
			values.Set("test", "test")
			mac := client.GenerateCheckMacValue(values)
			Expect(mac).To(Equal("B5947417B0FB8C1B8EBCA08E2FF3D66B7FB2329C8529C1331FE2268E9430903A"))
		})
	})

	Context("AutoPostTmpl", func() {
		It("shouldn't fail to generate auto post form html", func() {
			client := ecpay.NewStageECPayClient()
			values := url.Values{}
			values.Set("test", "test")
			values.Add("test", "B")
			values.Set("Test", "C")
			values.Set("D", "E")
			values.Set("F", "G")
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
})
