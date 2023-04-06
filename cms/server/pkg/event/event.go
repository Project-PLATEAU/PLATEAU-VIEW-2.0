package event

import (
	"time"

	"github.com/reearth/reearth-cms/server/pkg/operator"
)

type Type string

const (
	ItemCreate      = "item.create"
	ItemUpdate      = "item.update"
	ItemDelete      = "item.delete"
	ItemPublish     = "item.publish"
	ItemUnpublish   = "item.unpublish"
	AssetCreate     = "asset.create"
	AssetDecompress = "asset.decompress"
	AssetDelete     = "asset.delete"
)

type Event[T any] struct {
	id        ID
	timestamp time.Time
	operator  operator.Operator
	ty        Type
	prj       *Project
	object    T
}

func (e *Event[T]) ID() ID {
	return e.id
}

func (e *Event[T]) Type() Type {
	return e.ty
}

func (e *Event[T]) Timestamp() time.Time {
	return e.timestamp
}

func (e *Event[T]) Operator() operator.Operator {
	return e.operator
}

func (e *Event[T]) Project() *Project {
	return e.prj.Clone()
}

func (e *Event[T]) Object() any {
	return e.object
}

func (e *Event[T]) Clone() *Event[T] {
	if e == nil {
		return nil
	}
	return &Event[T]{
		id:        e.id.Clone(),
		timestamp: e.timestamp,
		operator:  e.operator,
		ty:        e.ty,
		prj:       e.prj.Clone(),
		object:    e.object,
	}
}

type Project struct {
	ID    string
	Alias string
}

func (p *Project) Clone() *Project {
	if p == nil {
		return nil
	}
	return &Project{
		ID:    p.ID,
		Alias: p.Alias,
	}
}
