package request

import (
	"time"

	"github.com/samber/lo"
)

type Builder struct {
	r *Request
}

func New() *Builder {
	return &Builder{r: &Request{}}
}

func (b *Builder) Build() (*Request, error) {
	if b.r.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.r.project.IsNil() {
		return nil, ErrInvalidID
	}
	if b.r.workspace.IsNil() {
		return nil, ErrInvalidID
	}
	if b.r.thread.IsNil() {
		return nil, ErrInvalidID
	}
	if b.r.createdBy.IsNil() {
		return nil, ErrInvalidID
	}
	if len(b.r.items) == 0 {
		return nil, ErrEmptyItems
	}
	if b.r.Items().HasDuplication() {
		return nil, ErrDuplicatedItem
	}
	if b.r.title == "" {
		return nil, ErrEmptyTitle
	}
	if b.r.state == "" {
		b.r.state = StateWaiting
	}
	if b.r.updatedAt.IsZero() {
		b.r.updatedAt = b.r.id.Timestamp()
	}
	return b.r, nil
}

func (b *Builder) MustBuild() *Request {
	return lo.Must(b.Build())
}

func (b *Builder) NewID() *Builder {
	b.r.id = NewID()
	return b
}

func (b *Builder) ID(id ID) *Builder {
	b.r.id = id
	return b
}

func (b *Builder) Title(t string) *Builder {
	b.r.title = t
	return b
}

func (b *Builder) Description(desc string) *Builder {
	b.r.description = desc
	return b
}

func (b *Builder) Items(items ItemList) *Builder {
	b.r.items = items
	return b
}

func (b *Builder) Project(p ProjectID) *Builder {
	b.r.project = p
	return b
}

func (b *Builder) Workspace(w WorkspaceID) *Builder {
	b.r.workspace = w
	return b
}

func (b *Builder) State(s State) *Builder {
	b.r.state = s
	return b
}

func (b *Builder) CreatedBy(u UserID) *Builder {
	b.r.createdBy = u
	return b
}

func (b *Builder) Reviewers(r UserIDList) *Builder {
	b.r.reviewers = r
	return b
}

func (b *Builder) Thread(t ThreadID) *Builder {
	b.r.thread = t
	return b
}

func (b *Builder) UpdatedAt(ua time.Time) *Builder {
	b.r.updatedAt = ua
	return b
}

func (b *Builder) ApprovedAt(a *time.Time) *Builder {
	b.r.approvedAt = a
	return b
}

func (b *Builder) ClosedAt(c *time.Time) *Builder {
	b.r.closedAt = c
	return b
}
