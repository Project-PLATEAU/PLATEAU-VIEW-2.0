package repo

import (
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
)

type Container struct {
	Asset       Asset
	AssetFile   AssetFile
	Lock        Lock
	User        User
	Workspace   Workspace
	Project     Project
	Model       Model
	Schema      Schema
	Item        Item
	Integration Integration
	Thread      Thread
	Event       Event
	Request     Request
	Transaction usecasex.Transaction
}

var (
	ErrOperationDenied = rerror.NewE(i18n.T("operation denied"))
)

func (c *Container) Filtered(workspace WorkspaceFilter, project ProjectFilter) *Container {
	if c == nil {
		return c
	}
	return &Container{
		Asset:       c.Asset.Filtered(project),
		AssetFile:   c.AssetFile,
		Lock:        c.Lock,
		Transaction: c.Transaction,
		Workspace:   c.Workspace,
		User:        c.User,
		Request:     c.Request,
		Item:        c.Item.Filtered(project),
		Project:     c.Project.Filtered(workspace),
		Model:       c.Model.Filtered(project),
		Schema:      c.Schema.Filtered(workspace),
		Thread:      c.Thread.Filtered(workspace),
		Integration: c.Integration,
		Event:       c.Event,
	}
}

type WorkspaceFilter struct {
	Readable user.WorkspaceIDList
	Writable user.WorkspaceIDList
}

func WorkspaceFilterFromOperator(o *usecase.Operator) WorkspaceFilter {
	return WorkspaceFilter{
		Readable: o.AllReadableWorkspaces(),
		Writable: o.AllWritableWorkspaces(),
	}
}

func (f WorkspaceFilter) Clone() WorkspaceFilter {
	return WorkspaceFilter{
		Readable: f.Readable.Clone(),
		Writable: f.Writable.Clone(),
	}
}

func (f WorkspaceFilter) Merge(g WorkspaceFilter) WorkspaceFilter {
	var r, w user.WorkspaceIDList
	if f.Readable != nil || g.Readable != nil {
		if f.Readable == nil {
			r = g.Readable.Clone()
		} else {
			r = append(f.Readable, g.Readable...)
		}
	}
	if f.Writable != nil || g.Writable != nil {
		if f.Writable == nil {
			w = g.Writable.Clone()
		} else {
			w = append(f.Writable, g.Writable...)
		}
	}
	return WorkspaceFilter{
		Readable: r,
		Writable: w,
	}
}

func (f WorkspaceFilter) CanRead(id user.WorkspaceID) bool {
	return f.Readable == nil || f.Readable.Has(id) || f.CanWrite(id)
}

func (f WorkspaceFilter) CanWrite(id user.WorkspaceID) bool {
	return f.Writable == nil || f.Writable.Has(id)
}

type ProjectFilter struct {
	Readable project.IDList
	Writable project.IDList
}

func ProjectFilterFromOperator(o *usecase.Operator) ProjectFilter {
	return ProjectFilter{
		Readable: o.AllReadableProjects(),
		Writable: o.AllWritableProjects(),
	}
}

func (f ProjectFilter) Clone() ProjectFilter {
	return ProjectFilter{
		Readable: f.Readable.Clone(),
		Writable: f.Writable.Clone(),
	}
}

func (f ProjectFilter) Merge(g ProjectFilter) ProjectFilter {
	var r, w project.IDList
	if f.Readable != nil || g.Readable != nil {
		if f.Readable == nil {
			r = g.Readable.Clone()
		} else {
			r = append(f.Readable, g.Readable...)
		}
	}
	if f.Writable != nil || g.Writable != nil {
		if f.Writable == nil {
			w = g.Writable.Clone()
		} else {
			w = append(f.Writable, g.Writable...)
		}
	}
	return ProjectFilter{
		Readable: r,
		Writable: w,
	}
}

func (f ProjectFilter) CanRead(id project.ID) bool {
	return f.Readable == nil || f.Readable.Has(id) || f.CanWrite(id)
}

func (f ProjectFilter) CanWrite(id project.ID) bool {
	return f.Writable == nil || f.Writable.Has(id)
}
