package main

import (
	"gopkg.in/gomail.v2"
	"gvb_server/core"
	"gvb_server/global"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

func NewNote() Api {
	return Api{
		Subject: Note,
	}
}

func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

// send 邮件发送 发给谁 主题 正文
func send(name, subject, body string) error {
	e := global.Config.Email
	return sendEmail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

func sendEmail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(userName, sendName)) //发件人邮箱， 发件人名字
	message.SetHeader("To", mailTo)                                      //发送给谁
	message.SetHeader("Subject", subject)                                //主题
	message.SetBody("text/html", body)
	dialer := gomail.NewDialer(host, port, userName, authCode)
	err := dialer.DialAndSend(message)
	return err
}

func main() {
	core.InitConf()
	core.InitLogger()
	NewCode().Send("1390337762@qq.com", "yanzhengma")
}
