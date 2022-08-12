package scb

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("api", func() {
	Context("Works", func() {
		err := ExchangeRatioWatcher{}.Call(context.TODO()).Error()
		Expect(err).NotTo(HaveOccurred())
	})
})

func TestExchangeRatioResult_Error(t *testing.T) {
}
