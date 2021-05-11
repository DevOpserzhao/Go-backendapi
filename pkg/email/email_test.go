package email

import "testing"

func TestEmail_Send(t *testing.T) {
	config := &EmailConfig{
		Subject:  "Test",
		SMTPHost: "smtp.qq.com",
		SMTP:     25,
		Sender:   "931883200@qq.com",
		Receiver: []string{"liyaoo1995@163.com"},
		Code:     "fdegdnvvnibhbdcf",
		Call:     "<a href='http://127.0.0.1/api/v1/verifyemail'>%s</a>",
	}
	e := New(config)
	err := e.Send("", "验证邮箱")
	if err != nil {
		t.Error(err.Error())
	}
}
