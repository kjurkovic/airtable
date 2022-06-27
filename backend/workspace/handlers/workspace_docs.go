// Package classification Workspace API
//
// Documentation for Workspace API
// 	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
// swagger:meta

package handlers

import (
	"workspace/models"

	"github.com/google/uuid"
)

// Swagger docs types

// A list of workspaces
// swagger:response workspacesResponse
type workspacesResponse struct {
	// User workspaces
	// in: body
	Body []models.Workspace
}

// swagger:parameters deleteWorkspace
type workspaceIDParameterWrapper struct {
	// The id of the workspace to delete
	// in: path
	// required: true
	ID uuid.UUID `json:"id"`
}

// swagger:response noContent
type workspaceNoContent struct {
}
