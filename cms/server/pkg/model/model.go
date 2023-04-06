package model

import (
	"fmt"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
	"golang.org/x/exp/slices"
)

var (
	ErrInvalidKey = rerror.NewE(i18n.T("invalid key"))
	ngKeys        = []string{"assets", "schemas", "models", "items"}
)

type Model struct {
	id          ID
	project     ProjectID
	schema      SchemaID
	name        string
	description string
	key         key.Key
	public      bool
	updatedAt   time.Time
}

func (p *Model) ID() ID {
	return p.id
}

func (p *Model) Schema() SchemaID {
	return p.schema
}

func (p *Model) Project() ProjectID {
	return p.project
}

func (p *Model) Name() string {
	return p.name
}

func (p *Model) SetName(name string) {
	p.name = name
}

func (p *Model) Description() string {
	return p.description
}

func (p *Model) SetDescription(description string) {
	p.description = description
}

func (p *Model) Key() key.Key {
	return p.key
}

func (p *Model) SetKey(key key.Key) error {
	if !validateModelKey(key) {
		return &rerror.Error{
			Label: ErrInvalidKey,
			Err:   fmt.Errorf("%s", key.String()),
		}
	}
	p.key = key
	return nil
}

func (p *Model) Public() bool {
	return p.public
}

func (p *Model) SetPublic(public bool) {
	p.public = public
}

func (p *Model) UpdatedAt() time.Time {
	if p.updatedAt.IsZero() {
		return p.id.Timestamp()
	}
	return p.updatedAt
}

func (p *Model) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}

func (p *Model) CreatedAt() time.Time {
	return p.id.Timestamp()
}

func (p *Model) Clone() *Model {
	if p == nil {
		return nil
	}

	return &Model{
		id:          p.id.Clone(),
		project:     p.project.Clone(),
		schema:      p.schema.Clone(),
		name:        p.name,
		description: p.description,
		key:         p.Key(),
		public:      p.public,
		updatedAt:   p.updatedAt,
	}
}

func validateModelKey(k key.Key) bool {
	// assets is used as an API endpoint
	return k.IsValid() && len(k.String()) > 2 && !slices.Contains(ngKeys, k.String())
}
