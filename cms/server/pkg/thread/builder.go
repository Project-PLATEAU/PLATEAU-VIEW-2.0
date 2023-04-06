package thread

import (
	"golang.org/x/exp/slices"
)

type Builder struct {
	th *Thread
}

func New() *Builder {
	return &Builder{th: &Thread{}}
}

func (b *Builder) Build() (*Thread, error) {
	if b.th.id.IsNil() {
		return nil, ErrInvalidID
	}

	if b.th.workspace.IsNil() {
		return nil, ErrNoWorkspaceID
	}

	return b.th, nil
}

func (b *Builder) MustBuild() *Thread {
	th, err := b.Build()
	if err != nil {
		panic(err)
	}
	return th
}

func (b *Builder) ID(id ID) *Builder {
	b.th.id = id
	return b
}

func (b *Builder) Workspace(wid WorkspaceID) *Builder {
	b.th.workspace = wid
	return b
}

func (b *Builder) NewID() *Builder {
	b.th.id = NewID()
	return b
}

func (b *Builder) Comments(c []*Comment) *Builder {
	b.th.comments = slices.Clone(c)
	return b
}
