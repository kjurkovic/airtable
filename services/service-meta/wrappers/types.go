package wrappers

type AuditType string

var (
	MetaCreated  AuditType = "META_CREATED"
	MetaModified AuditType = "META_MODIFIED"
	MetaDeleted  AuditType = "META_DELETED"
)
