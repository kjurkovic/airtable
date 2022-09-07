package wrappers

type AuditType string

var (
	Registration       AuditType = "REGISTRATION"
	Login              AuditType = "LOGIN"
	DeleteUser         AuditType = "DELETE_USER"
	UpdateUserPassword AuditType = "UPDATE_USER_PASSWORD"
	UpdateUser         AuditType = "UPDATE_USER"
)
