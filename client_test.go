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
})
