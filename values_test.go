package ecpay_test

import (
	"github.com/Laysi/go-ecpay-sdk"
	. "github.com/onsi/ginkgo"
	"net/url"
	"sort"

	. "github.com/onsi/gomega"
)

var _ = Describe("Values", func() {
	Context("Encoding", func() {
		It("should use lower case sorting to encoding", func() {
			v := ecpay.ECPayValues{url.Values{}}
			v.Set("aA", "aA")
			v.Set("Zzzz", "Zzzz")
			v.Set("Ab", "Ab")
			v.Set("Bbbbb", "Bbbbb")
			Expect(v.Encode()).To(Equal("aA=aA&Ab=Ab&Bbbbb=Bbbbb&Zzzz=Zzzz"))
		})
	})
})

var _ = Describe("Slice", func() {
	Context("Encoding", func() {
		It("should use lower case sorting to encoding", func() {
			v := []string{"aA", "Zzzz", "Ab", "Bbbbb"}
			sort.Sort(ecpay.LowerStringSlice(v))
			Expect(v).To(Equal([]string{"aA", "Ab", "Bbbbb", "Zzzz"}))
		})
	})
})
