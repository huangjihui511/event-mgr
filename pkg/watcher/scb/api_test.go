package scb

import (
	"context"
	"huangjihui511/event-mgr/pkg/utils"
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
	Context("Works time china", func() {
		BeforeEach(func() {
			utils.Location = time.Now().Local().Location()
		})
		It("open", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-12T15:04:05+08:00")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(true))
		})
		It("not open 0", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-12T9:04:05+08:00")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(false))
		})
		It("not open 1", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-12T18:04:05+08:00")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(false))
		})
		It("not open 2", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-13T10:04:05+08:00")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(false))
		})
	})
	Context("Works time in other location", func() {
		It("east us location", func() {
			in, err := time.Parse(time.RFC3339, "2022-08-17T22:04:05+20:00")
			Expect(err).NotTo(HaveOccurred())
			r := isSCBMarketOpen(in)
			Expect(r).To(Equal(true))
		})
	})
})

func TestSCB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "scb")
}
