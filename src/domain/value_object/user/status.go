package value_object

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusNonactive UserStatus = "nonactive"
	UserStatusBlocked   UserStatus = "blocked"
	UserStatusSuspended UserStatus = "suspended"
)
