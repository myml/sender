package sender

import (
	"math/rand"
	"os"
	"testing"

	sendcloud "bj.git.sndu.cn/deepinid/sender/send-cloud"
	share "bj.git.sndu.cn/deepinid/sender/share"
)

const (
	TestPhone    = ""
	TestTemplate = ""
)

var (
	SendCloud_SMS_KEY  = os.Getenv("SENDCLOUD_SMS_KEY")
	SendCloud_SMS_USER = os.Getenv("SENDCLOUD_SMS_USER")
)

func Test_SendCloud_SendSMS(t *testing.T) {
	sc := sendcloud.NewSendCloud(sendcloud.Account{
		SMS_KEY:  SendCloud_SMS_KEY,
		SMS_USER: SendCloud_SMS_USER,
	})
	send := NewSender(sc)

	opt := share.SendSMSOption{
		Phone: []string{TestPhone},
		TemplateOption: share.TemplateOption{
			Template: TestTemplate,
			TemplateVariable: map[string]interface{}{
				"code": rand.Intn(8999) + 1000,
			},
		},
	}
	err := send.SendSMS(&opt)
	if err != nil {
		t.Fatal(err)
	}
}
