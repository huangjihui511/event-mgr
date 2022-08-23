package notify

type RouterItem struct {
	EventName string
	Email     string
}

var (
	RouterItems = map[RouterItem]bool{
		{
			EventName: "",
			Email:     EmailJihui,
		}: true,
	}
	EmailJihui = "717655909@qq.com"
)

func init() {

}
