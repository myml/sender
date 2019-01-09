package sender

type Sender interface {
	SendSMS() error
}

func NewSender(s Sender) Sender {
	return s
}
