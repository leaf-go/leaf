package utils

import (
	"errors"
	"github.com/go-gomail/gomail"
)

const ()

type EmailConfig struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	ServerPort int

	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	Toers []string
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCers []string
}

type email struct {
	config  *EmailConfig
	message *gomail.Message
}

type EmailConfigs func(*email)

func Email(configs ...EmailConfigs) *email {
	e := &email{}
	e.config = &EmailConfig{}
	e.message = gomail.NewMessage()
	e.applyConfigs(configs)

	//e.config.ServerHost = IF(e.config.ServerHost == "", Email_Def_ServerHost, e.config.ServerHost).(string)
	//e.config.ServerPort = IF(e.config.ServerPort == 0, Email_Def_ServerPort, e.config.ServerPort).(int)
	//e.config.FromEmail = IF(e.config.FromEmail == "", Email_Def_FromEmail, e.config.FromEmail).(string)
	//e.config.FromPasswd = IF(e.config.FromPasswd == "", Email_Def_FromPasswd, e.config.FromPasswd).(string)
	//e.config.Toers = IF(len(e.config.Toers) == 0, strings.Split(Email_Def_Toers, ","), e.config.Toers).([]string)

	return e
}

func SetEmailServer(host string, port int) EmailConfigs {
	return func(e *email) {
		e.config.ServerHost = host
		e.config.ServerPort = port
	}
}

func SetEmailFrom(user, password string) EmailConfigs {
	return func(e *email) {
		e.config.FromEmail = user
		e.config.FromPasswd = password
	}
}

func SetEmailCCers(ccers ...string) EmailConfigs {
	return func(e *email) {
		e.config.CCers = ccers
	}
}

func SetEmailToers(toers ...string) EmailConfigs {
	return func(e *email) {
		e.config.Toers = toers
	}
}

func (e *email) applyConfigs(configs []EmailConfigs) {
	for _, f := range configs {
		f(e)
	}
}

// Send body支持html格式字符串
func (this *email) Send(subject, body string, configs ...EmailConfigs) (err error) {
	this.applyConfigs(configs)

	if len(this.config.Toers) == 0 {
		err = errors.New("toers is empty...")
		return
	}
	// 收件人可以有多个，故用此方式
	this.message.SetHeader("To", this.config.Toers...)
	//抄送列表
	if len(this.config.CCers) != 0 {
		this.message.SetHeader("Cc", this.config.CCers...)
	}
	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	this.message.SetAddressHeader("From", this.config.FromEmail, "")
	// 主题
	this.message.SetHeader("Subject", subject)
	// 正文
	this.message.SetBody("text/html", body)

	d := gomail.NewDialer(this.config.ServerHost, this.config.ServerPort, this.config.FromEmail, this.config.FromPasswd)
	// 发送
	err = d.DialAndSend(this.message)
	return
}
