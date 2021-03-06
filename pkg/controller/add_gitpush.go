package controller

import (
	"github.com/redhat-developer/git-service/pkg/controller/gitpush"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, gitpush.Add)
}
