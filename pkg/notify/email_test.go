package notify

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Email", func() {
	Context("Work", func() {

	})
})

func TestSendToEmail(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "notify")
}
