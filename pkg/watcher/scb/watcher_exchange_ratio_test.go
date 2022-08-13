package scb

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("api", func() {
	Context("Works", func() {
		err := WatcherExchangeRatio{}.Call(context.TODO()).Error()
		Expect(err).NotTo(HaveOccurred())
	})
})
