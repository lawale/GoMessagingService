package entities

type PushNotificationDevice struct {
	BaseEntity
	UserId      string
	DeviceToken string
}
