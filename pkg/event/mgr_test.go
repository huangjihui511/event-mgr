package event

import (
	"context"
	"huangjihui511/event-mgr/pkg/utils"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SCB Events", func() {
	Context("Is Notify", func() {
		AfterEach(func() {
			utils.TimeNow = time.Now
		})
		It("Should Not Send Email", func() {
			utils.TimeNow = func() time.Time {
				in, err := time.Parse(time.RFC3339, "2022-08-12T9:04:05+08:00")
				Expect(err).NotTo(HaveOccurred())
				return in
			}
			event := Timer{
				watcherInterface: scb.NewWatcherExchangeRatio("SCB Watcher"),
				duration:         time.Minute * 30,
			}
			isNotify := event.Watcher().Call(context.TODO()).IsNotify()
			Expect(isNotify).To(Equal(false))
		})
		It("Should Send Email", func() {
			utils.TimeNow = func() time.Time {
				in, err := time.Parse(time.RFC3339, "2022-08-12T9:34:05+08:00")
				Expect(err).NotTo(HaveOccurred())
				return in
			}
			event := Timer{
				watcherInterface: scb.NewWatcherExchangeRatio("SCB Watcher"),
				duration:         time.Minute * 30,
			}
			isNotify := event.Watcher().Call(context.TODO()).IsNotify()
			Expect(isNotify).To(Equal(true))
		})
	})
})
