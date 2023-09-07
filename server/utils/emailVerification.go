package utils

import (
	"fmt"
	"server/config"
	"server/setting"
	"strings"

	"gopkg.in/gomail.v2"
)

// 发送验证码
func SendEmailCode(toUserEmail, code string) error {
	v, _ := config.LoadConfigFromYaml()
	Mailbox := &setting.Mailbox{}
	err := v.UnmarshalKey("mailbox", Mailbox)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	mailTitle := "Atidotes注册验证码：" + code //邮件标题

	var build strings.Builder
	build.WriteString("<P>欢迎您注册此网站，您的验证码为")
	build.WriteString(code)
	build.WriteString(" ，感谢使用！</p>")
	fmt.Print(build.String())
	mailBody := build.String() // 邮件内容

	//接收者邮箱列表
	mailTo := []string{
		toUserEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", Mailbox.Sender) //发送者腾讯邮箱账号
	m.SetHeader("To", mailTo...)        //接收者邮箱列表
	m.SetHeader("Subject", mailTitle)   //邮件标题
	m.SetBody("text/html", mailBody)    //邮件内容,可以是html

	// //添加附件
	// zipPath := "./user/temp.zip"
	// m.Attach(zipPath)

	//发送邮件服务器、端口、发送者qq邮箱、qq邮箱授权码
	//服务器地址和端口是腾讯的
	d := gomail.NewDialer("smtp.qq.com", 587, Mailbox.Sender, Mailbox.AuthCode)
	err = d.DialAndSend(m)
	return err
}
