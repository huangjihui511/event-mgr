package scb

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("api", func() {
	Context("Works api", func() {
		ratio, err := getSCExchangeRatio(context.TODO())
		Expect(err).NotTo(HaveOccurred())
		Expect(ratio.BuyRatio).ShouldNot(BeZero())
		Expect(ratio.SellRatio).ShouldNot(BeZero())
	})
	Context("Works time", func() {
		It("open", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-12T15:04:05Z")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(true))
		})
		It("not open 0", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-12T9:04:05Z")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(false))
		})
		It("not open 1", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-12T18:04:05Z")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(false))
		})
		It("not open 2", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-13T10:04:05Z")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(false))
		})
	})
})

func TestSCB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "scb")
}
