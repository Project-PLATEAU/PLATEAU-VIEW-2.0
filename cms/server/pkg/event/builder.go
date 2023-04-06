package event

import (
	"time"

	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/samber/lo"
)

type Builder[T any] struct {
	i *Event[T]
}

func New[T any]() *Builder[T] {
	return &Builder[T]{i: &Event[T]{}}
}

func (b *Builder[T]) Build() (*Event[T], error) {
	if b.i.id.IsNil() || !b.i.operator.Validate() {
		return nil, ErrInvalidID
	}
	return b.i, nil
}

func (b *Builder[T]) MustBuild() *Event[T] {
	return lo.Must(b.Build())
}

func (b *Builder[T]) ID(id ID) *Builder[T] {
	b.i.id = id
	return b
}

func (b *Builder[T]) NewID() *Builder[T] {
	b.i.id = NewID()
	return b
}

func (b *Builder[T]) Timestamp(t time.Time) *Builder[T] {
	b.i.timestamp = t
	return b
}

func (b *Builder[T]) Type(t Type) *Builder[T] {
	b.i.ty = t
	return b
}

func (b *Builder[T]) Project(prj *Project) *Builder[T] {
	b.i.prj = prj.Clone()
	return b
}

func (b *Builder[T]) Operator(o operator.Operator) *Builder[T] {
	b.i.operator = o
	return b
}

func (b *Builder[T]) Object(o T) *Builder[T] {
	b.i.object = o
	return b
}
