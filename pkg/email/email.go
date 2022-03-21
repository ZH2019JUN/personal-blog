package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host string
	Port int
	IsSSL bool
	UserName string
	Password string
	From string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{info}
}

func (e *Email)SendEmail(to []string,subject,body string) error {
	m := gomail.NewMessage()
	//发送邮件信息为：发件人，收件人，主题和正文
	m.SetHeader("From",e.From)
	m.SetHeader("To",to...)
	m.SetHeader("Subject",subject)
	m.SetBody("text/html",body)

	//调用NewDialer方法新建一个SMTP拨号实例并设置拨号信息
	dialer := gomail.NewDialer(e.Host,e.Port,e.UserName,e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	//打开与SMTP服务器的连接并发送电子邮件
	return dialer.DialAndSend(m)
}
