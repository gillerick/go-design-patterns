package main

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) setTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) setSubtitle(subtitle string) {
	nb.Subtitle = subtitle
}

func (nb *NotificationBuilder) setMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) setImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) setIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) setPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) setNotType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	return nil, nil
}
