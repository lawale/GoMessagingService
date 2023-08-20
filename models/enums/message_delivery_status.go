package enums

type MessageDeliveryStatus int

const (
	Pending = iota
	Delivered
	Failed
)
