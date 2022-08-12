package scb

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("api", func() {
	Context("Works", func() {
		ratio, err := getSCExchangeRatio(context.TODO())
		Expect(err).NotTo(HaveOccurred())
		Expect(ratio.BuyRatio).ShouldNot(BeZero())
		Expect(ratio.SellRatio).ShouldNot(BeZero())
	})
})

func Test_getSCExchangeRatio(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "api")
}
