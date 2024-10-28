package enum

type NotificationType string

const (
	Assigned NotificationType = "assigned"
	Deleted  NotificationType = "deleted"
	Updated  NotificationType = "updated"
)
