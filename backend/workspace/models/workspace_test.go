package models

import "testing"

func TestCheckValidation(t *testing.T) {
	workspace := &Workspace{
		Name: "Name",
	}

	err := workspace.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
