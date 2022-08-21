package zsb

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("api", func() {
	r, err := getZSExchangeRatio(context.TODO())
	Expect(err).ShouldNot(HaveOccurred())
	Expect(r.USDBuy).ShouldNot(BeZero())
	Expect(r.USDSell).ShouldNot(BeZero())
})

func TestZSB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "zsb")
}
