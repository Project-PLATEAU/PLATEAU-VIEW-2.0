package asset

import (
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	a *Asset
}

func New() *Builder {
	return &Builder{a: &Asset{}}
}

func (b *Builder) Build() (*Asset, error) {
	if b.a.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.a.project.IsNil() {
		return nil, ErrNoProjectID
	}
	if b.a.user.IsNil() && b.a.integration.IsNil() {
		return nil, ErrNoUser
	}
	if b.a.thread.IsNil() {
		return nil, ErrNoThread
	}
	if b.a.size == 0 {
		return nil, ErrZeroSize
	}
	if b.a.uuid == "" {
		return nil, ErrNoUUID
	}
	if b.a.createdAt.IsZero() {
		b.a.createdAt = b.a.id.Timestamp()
	}
	return b.a, nil
}

func (b *Builder) MustBuild() *Asset {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *Builder) ID(id ID) *Builder {
	b.a.id = id
	return b
}

func (b *Builder) NewID() *Builder {
	b.a.id = NewID()
	return b
}

func (b *Builder) Project(pid ProjectID) *Builder {
	b.a.project = pid
	return b
}

func (b *Builder) CreatedAt(createdAt time.Time) *Builder {
	b.a.createdAt = createdAt
	return b
}

func (b *Builder) CreatedByUser(createdBy UserID) *Builder {
	b.a.user = &createdBy
	b.a.integration = nil
	return b
}

func (b *Builder) CreatedByIntegration(createdBy IntegrationID) *Builder {
	b.a.integration = &createdBy
	b.a.user = nil
	return b
}

func (b *Builder) FileName(name string) *Builder {
	b.a.fileName = name
	return b
}

func (b *Builder) Size(size uint64) *Builder {
	b.a.size = size
	return b
}

func (b *Builder) Type(t *PreviewType) *Builder {
	b.a.previewType = t
	return b
}

func (b *Builder) UUID(uuid string) *Builder {
	b.a.uuid = uuid
	return b
}

func (b *Builder) NewUUID() *Builder {
	b.a.uuid = uuid.NewString()
	return b
}

func (b *Builder) Thread(th ThreadID) *Builder {
	b.a.thread = th
	return b
}

func (b *Builder) ArchiveExtractionStatus(s *ArchiveExtractionStatus) *Builder {
	b.a.archiveExtractionStatus = s
	return b
}
