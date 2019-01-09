package sender

import (
	"testing"

	AliCloud "github.com/myml/sender/driver/ali-cloud"
	SendCloud "github.com/myml/sender/driver/send-cloud"
)

func Test_AliCloud_SendSMS(t *testing.T) {
	ali := AliCloud.NewAliCloud()
	send := NewSender(ali)
	send.SendSMS()
}

func Test_SendCloud_SendSMS(t *testing.T) {
	sc := SendCloud.NewSendCloud()
	send := NewSender(sc)
	send.SendSMS()
}
