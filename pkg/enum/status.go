package enum

type Status string

const (
	NotStarted Status = "not started"
	InProgress Status = "in progress"
	Completed  Status = "completed"
)