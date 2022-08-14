package devops

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("born", func() {
	w := NewWatcherBorn()
	It("Should Notify", func() {
		i := w.Call(context.TODO()).IsNotify()
		Expect(i).Should(Equal(true))
	})
	It("Should Not Notify", func() {
		i := w.Call(context.TODO()).IsNotify()
		Expect(i).Should(Equal(false))
	})
})

func TestWatcherBorn_Call(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "scb")
}
