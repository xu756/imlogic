package msg

import (
	"imlogic/common/config"
	"imlogic/internal/email"
	"imlogic/internal/xerr"
	"log"
	"net/smtp"
)

// 发送 邮箱验证码 用于注册
var e *email.Email

func SendEmail(to string, content string) error {
	e = email.NewEmail()
	e.From = config.RunData.EmailConfig.From
	e.To = []string{to}
	e.Subject = "验证码"
	e.Text = []byte("验证码为：" + content)
	err := e.Send(config.RunData.EmailConfig.SMTPAddr, smtp.PlainAuth("", config.RunData.EmailConfig.From, config.RunData.EmailConfig.Password, config.RunData.EmailConfig.Host))
	if err != nil {
		log.Print(err)
		return xerr.ErrMsg(xerr.SendEmailErr)
	}
	return nil
}
