package integration

import (
	"net/url"
	"time"
)

type Builder struct {
	i *Integration
}

func New() *Builder {
	return &Builder{i: &Integration{}}
}

func (b *Builder) Build() (*Integration, error) {
	if b.i.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.i.updatedAt.IsZero() {
		b.i.updatedAt = b.i.CreatedAt()
	}
	return b.i, nil
}

func (b *Builder) MustBuild() *Integration {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *Builder) ID(id ID) *Builder {
	b.i.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.i.id = NewID()
	return b
}

func (b *Builder) UpdatedAt(updatedAt time.Time) *Builder {
	b.i.updatedAt = updatedAt
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.i.name = name
	return b
}

func (b *Builder) Description(description string) *Builder {
	b.i.description = description
	return b
}

func (b *Builder) Type(t Type) *Builder {
	b.i.iType = t
	return b
}

func (b *Builder) LogoUrl(logoURL *url.URL) *Builder {
	if logoURL == nil {
		b.i.logoUrl = nil
	} else {
		b.i.logoUrl, _ = url.Parse(logoURL.String())
	}
	return b
}

func (b *Builder) GenerateToken() *Builder {
	b.i.RandomToken()
	return b
}

func (b *Builder) Token(token string) *Builder {
	b.i.token = token
	return b
}

func (b *Builder) Developer(developer UserID) *Builder {
	b.i.developer = developer
	return b
}

func (b *Builder) Webhook(webhook []*Webhook) *Builder {
	b.i.webhooks = webhook
	return b
}
