package enums

type MessageType int

const (
	InApp = 1 << iota
	Push
	Sms
)
