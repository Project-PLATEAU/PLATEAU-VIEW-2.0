package project

import (
	"net/url"
	"time"
)

type Builder struct {
	p *Project
}

func New() *Builder {
	return &Builder{p: &Project{}}
}

func (b *Builder) Build() (*Project, error) {
	if b.p.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.p.alias != "" && !CheckAliasPattern(b.p.alias) {
		return nil, ErrInvalidAlias
	}
	if b.p.updatedAt.IsZero() {
		b.p.updatedAt = b.p.CreatedAt()
	}
	return b.p, nil
}

func (b *Builder) MustBuild() *Project {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *Builder) ID(id ID) *Builder {
	b.p.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.p.id = NewID()
	return b
}

func (b *Builder) UpdatedAt(updatedAt time.Time) *Builder {
	b.p.updatedAt = updatedAt
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.p.name = name
	return b
}

func (b *Builder) Description(description string) *Builder {
	b.p.description = description
	return b
}

func (b *Builder) Alias(alias string) *Builder {
	b.p.alias = alias
	return b
}

func (b *Builder) ImageURL(imageURL *url.URL) *Builder {
	if imageURL == nil {
		b.p.imageURL = nil
	} else {
		// https://github.com/golang/go/issues/38351
		imageURL2 := *imageURL
		b.p.imageURL = &imageURL2
	}
	return b
}

func (b *Builder) Workspace(team WorkspaceID) *Builder {
	b.p.workspaceID = team
	return b
}

func (b *Builder) Publication(publication *Publication) *Builder {
	b.p.publication = publication
	return b
}
