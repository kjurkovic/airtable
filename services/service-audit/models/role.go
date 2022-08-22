package models

type Role int

const (
	ADMIN Role = iota
	MANAGER
	CLIENT
	UNDEFINED
)

func (role Role) toString() string {
	switch role {
	case ADMIN:
		return "admin"
	case MANAGER:
		return "manager"
	case CLIENT:
		return "client"
	default:
		return "undefined"
	}
}
