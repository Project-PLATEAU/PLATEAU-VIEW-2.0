package interfaces

import (
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

type ListOperation string

const (
	ListOperationAdd    ListOperation = "add"
	ListOperationMove   ListOperation = "move"
	ListOperationRemove ListOperation = "remove"
)

var (
	ErrOperationDenied error = rerror.NewE(i18n.T("operation denied"))
	ErrInvalidOperator error = rerror.NewE(i18n.T("invalid operator"))
)

type Container struct {
	Asset       Asset
	Workspace   Workspace
	User        User
	Item        Item
	Project     Project
	Request     Request
	Model       Model
	Schema      Schema
	Integration Integration
	Thread      Thread
}
