package daily

import (
	"context"
	"huangjihui511/event-mgr/pkg/watcher/interfaces/mock_interfaces"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Daily", func() {
	Context("cron", func() {
		It("should work", func() {
			triggered := false
			c := newWithSeconds()
			_, err := c.AddFunc("* * * * * ?", func() {
				triggered = true
			})
			c.Start()
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(5 * time.Second)
			Expect(triggered).Should(BeTrue())
		})
	})
	Context("channel", func() {
		It("should work", func() {
			watcher := mock_interfaces.NewMockInterface(gomock.NewController(GinkgoT()))
			triggered := 0
			watcher.EXPECT().Call(gomock.Any()).Return(nil)
			c := NewDaily("* * * * * ?", watcher).Chan(context.TODO())
			endTimer := time.After(time.Second * 10)
			for {
				end := false
				select {
				case <-c:
					triggered++
				case <-endTimer:
					end = true
				}
				if end {
					break
				}
			}
			Expect(triggered >= 9).Should(BeTrue())
		})
	})
})

func TestDaily(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "daily")
}
