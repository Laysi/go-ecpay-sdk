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
			v.Set("Aa", "Aa")
			v.Set("Bbbbb", "Bbbbb")
			Expect(v.Encode()).To(Equal("aA=aA&Aa=Aa&Bbbbb=Bbbbb&Zzzz=Zzzz"))
		})
	})
})

var _ = Describe("Slice", func() {
	Context("Encoding", func() {
		It("should use lower case sorting to encoding", func() {
			v := []string{"aA", "Zzzz", "Aa", "Bbbbb"}
			sort.Sort(ecpay.LowerStringSlice(v))
			Expect(v).To(Equal([]string{"aA", "Aa", "Bbbbb", "Zzzz"}))
		})
	})
})
