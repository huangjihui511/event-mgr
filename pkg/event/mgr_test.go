package event

import (
	"context"
	eventInterface "huangjihui511/event-mgr/pkg/event/interfaces"
	watcherInterface "huangjihui511/event-mgr/pkg/watcher/interfaces"

	"huangjihui511/event-mgr/pkg/utils"
	watcherMockInterface "huangjihui511/event-mgr/pkg/watcher/interfaces/mock_interfaces"
	"huangjihui511/event-mgr/pkg/watcher/scb"
	"time"

	"github.com/golang/mock/gomock"
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
	Context("trigger times", func() {
		It("Should many times", func() {
			ctr := gomock.NewController(GinkgoT())
			watcher := watcherMockInterface.NewMockInterface(ctr)
			result := watcherMockInterface.NewMockResultInterface(ctr)
			times := 0
			result.EXPECT().IsNotify().Return(false).MaxTimes(100)
			result.EXPECT().Msg().Return("").MaxTimes(100)
			result.EXPECT().Error().Return(nil).MaxTimes(100)
			watcher.EXPECT().Call(gomock.Any()).Do(func(_ context.Context) watcherInterface.ResultInterface {
				times++
				return result
			}).Return(result).MaxTimes(100)
			watcher.EXPECT().Name().Return("test watcher").MaxTimes(100)
			r := watcher.Call(context.TODO())
			Expect(r.Error()).NotTo(HaveOccurred())
			e := NewTimer(1*time.Second, watcher)
			startEvents(context.TODO(), []eventInterface.Interface{e})
			endTimer := time.After(time.Second * 10)
			<-endTimer
			Expect(times > 10).Should(BeTrue())
		})
	})
})
