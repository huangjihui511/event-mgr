package notify

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Email", func() {
	Context("Work", func() {
		err := SendToEmail("717655909@qq.com", "test", emailMetaQQ, "test", "this is test")
		Expect(err).NotTo(HaveOccurred())
	})
})

func TestSendToEmail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "notify")
}
