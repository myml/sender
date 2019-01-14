package sender

import "bj.git.sndu.cn/deepinid/sender/share"

type Sender interface {
	SendSMS(*share.SendSMSOption) error
}

func NewSender(s Sender) Sender {
	return s
}
