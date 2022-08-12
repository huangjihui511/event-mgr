package notify

import (
	"net/smtp"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

type emailMeta struct {
	user     string
	name     string
	password string
	host     string
}

var (
	emailMeta163 = emailMeta{
		user: "billEventRobot@163.com",
		name: "jh",
		// password: "0sB-7M2-3im-Fpm",
		password: "WVJZOZWLIMGDCEUGmermaid",
		host:     "smtp.163.com:587",
	}
	emailMetaQQ = emailMeta{
		user:     "717655909@qq.com",
		name:     "notify robot",
		password: "!@#qweasd",
		host:     "smtp.qq.com:587",
	}
)

func SendToEmail(subject, body string) error {
	return sendToEmail2("717655909@qq.com", "jh", emailMetaQQ, subject, body)
}

func sendToEmail(toEmail, toName string, senderEmail emailMeta, subject, body string) error {
	auth := smtp.PlainAuth("", senderEmail.user, senderEmail.password, senderEmail.host)
	contentType := "Content-Type: multipart/alternative; "
	msg := []byte("To: " + toEmail + "\r\nFrom: " + senderEmail.name + "<" + senderEmail.user + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	send_to := strings.Split(toEmail, ";")
	err := smtp.SendMail(senderEmail.host, auth, senderEmail.user, send_to, msg)
	return err
}

func sendToEmail2(toEmail, toName string, senderEmail emailMeta, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail.user)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	// msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer(strings.Split(senderEmail.host, ":")[0], 587, senderEmail.user, senderEmail.password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
	return nil
}
