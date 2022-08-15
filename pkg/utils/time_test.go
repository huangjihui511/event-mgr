package utils

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("time", func() {
	t := TimeNow()
	Expect(t).NotTo(Equal(time.Time{}))
})

func TestTimeNow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "time")
}
