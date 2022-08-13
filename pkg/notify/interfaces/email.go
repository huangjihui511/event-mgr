package interfaces

type EmailMeta struct {
	User     string
	Name     string
	Password string
	Host     string
}

var (
	EmailMeta163 = EmailMeta{
		User: "billEventRobot@163.com",
		Name: "jh",
		// password: "0sB-7M2-3im-Fpm",
		Password: "WVJZOZWLIMGDCEUGmermaid",
		Host:     "smtp.163.com:587",
	}
	EmailMetaQQ = EmailMeta{
		User:     "717655909@qq.com",
		Name:     "notify robot",
		Password: "!@#qweasd",
		Host:     "smtp.qq.com:587",
	}
)

//go:generate sh -c "mockgen --build_flags=--mod=mod huangjihui511/event-mgr/pkg/notify/interfaces EmailInterface > ./mock_interfaces/email_interface.go"

type EmailInterface interface {
	Send(to, subject, body string) error
}
