package wrappers

type AuditType string

var (
	WorkspaceCreated  AuditType = "WORKSPACE_CREATED"
	WorkspaceModified AuditType = "WORKSPACE_MODIFIED"
	WorkspaceDeleted  AuditType = "WORKSPACE_DELETED"
)
