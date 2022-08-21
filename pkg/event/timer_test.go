package event

import (
	"context"
	"huangjihui511/event-mgr/pkg/watcher/interfaces/mock_interfaces"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timer", func() {
	Context("Channel", func() {
		It("Should triggered", func() {
			watcher := mock_interfaces.NewMockInterface(gomock.NewController(GinkgoT()))
			triggered := 0
			watcher.EXPECT().Call(gomock.Any()).Return(nil)
			c := NewTimer(1*time.Second, watcher).Chan(context.TODO())
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

func TestTimer_Chan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "event")
}
