package event

import (
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timer", func() {
	Context("Channel", func() {
		tt := Timer{
			watcherInterface: scb.WatcherExchangeRatio{},
			duration:         time.Second * 5,
		}
		c := tt.Chan()
		triggered := false
		endTimer := time.After(time.Second * 7)
		for {
			out := false
			select {
			case <-endTimer:
				out = true
			case <-c:
				triggered = true
			}
			if out {
				break
			}
		}
		Expect(triggered).Should(Equal(true))
	})
})

func TestTimer_Chan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "event")
}
