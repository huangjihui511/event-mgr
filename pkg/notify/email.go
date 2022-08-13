package notify

import (
	notifyInterface "huangjihui511/event-mgr/pkg/notify/interfaces"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

var _ notifyInterface.EmailInterface = EmailSender{}

type EmailSender struct {
	notifyInterface.EmailMeta
}

// Send implements notifyInterface.EmailInterface
func (e EmailSender) Send(to string, subject string, body string) error {
	return sendToEmail2("717655909@qq.com", "jh", e.EmailMeta, subject, body)
}

func sendToEmail2(toEmail, toName string, senderEmail notifyInterface.EmailMeta, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail.User)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	// msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer(strings.Split(senderEmail.Host, ":")[0], 587, senderEmail.User, senderEmail.Password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
	return nil
}
