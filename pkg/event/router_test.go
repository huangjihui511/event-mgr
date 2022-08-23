package event

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router", func() {
	It("not empty", func() {
		ri := RouterItem{
			EventName: "abc",
			Email:     "123",
		}
		RouterItems[ri] = true
		r := RouterItems[ri]
		Expect(r).Should(BeTrue())
	})
	It("empty", func() {
		ri := RouterItem{
			EventName: "abc",
			Email:     "12345678",
		}
		r := RouterItems[ri]
		Expect(r).Should(BeFalse())
	})

})
