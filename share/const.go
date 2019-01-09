package share

import "errors"

// phone length limit
const (
	MinPhoneLength = 8
	MaxPhoneLength = 15
)

// error list
var (
	ErrorNotImplemented = errors.New("Not Implemented")
)
