package domain

type Role string

const (
	// Role
	Admin      Role = "admin"
	Supervisor Role = "supervisor"
	Staff      Role = "staff"
)

var AccessLevel = map[Role]int{
	Staff:      1,
	Supervisor: 2,
	Admin:      3,
}

func IsValidRole(r Role) bool {
	switch r {
	case Admin, Supervisor, Staff:
		return true
	default:
		return false
	}
}
