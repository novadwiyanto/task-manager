package enum

type Role string

const (
	Admin    Role = "admin"
	Member   Role = "member"
	Guest    Role = "guest"
)