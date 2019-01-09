package sendcloud

import "log"

func NewSendCloud() *SendCloud {
	return &SendCloud{}
}

type SendCloud struct {
}

func (sc *SendCloud) SendSMS() error {
	log.Println("SendCloud send SMS")
	return nil
}
