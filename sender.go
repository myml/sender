package sender

import "github.com/myml/sender/share"

type Sender interface {
	SendSMS(*share.SendSMSOption) error
}

func NewSender(s Sender) Sender {
	return s
}
