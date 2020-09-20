package ecpay_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoEcpaySdk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoEcpaySdk Suite")
}
