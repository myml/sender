package alicloud

import "log"

func NewAliCloud() *AliCloud {
	return &AliCloud{}
}

type AliCloud struct {
}

func (ali *AliCloud) SendSMS() error {
	log.Println("aliCloud send SMS")
	return nil
}
