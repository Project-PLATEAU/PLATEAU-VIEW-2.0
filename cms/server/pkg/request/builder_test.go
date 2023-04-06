package request

import (
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Build(t *testing.T) {
	i1, _ := NewItem(id.NewItemID())
	req := &Request{
		id:        NewID(),
		workspace: NewWorkspaceID(),
		project:   NewProjectID(),
		items: ItemList{{
			item:    NewItemID(),
			pointer: version.New().OrRef(),
		}},
		title:       "title",
		description: "desc",
		createdBy:   NewUserID(),
		reviewers:   []UserID{NewUserID()},
		thread:      NewThreadID(),
	}
	expected := &Request{
		id:          req.ID(),
		workspace:   req.Workspace(),
		project:     req.Project(),
		items:       req.Items(),
		title:       req.Title(),
		description: req.Description(),
		createdBy:   req.CreatedBy(),
		reviewers:   req.Reviewers(),
		thread:      req.thread,
		updatedAt:   req.ID().Timestamp(),
		state:       StateWaiting,
	}
	type fields struct {
		r *Request
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Request
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				r: req,
			},
			want: expected,
		},
		{
			name: "invalid ID",
			fields: fields{
				r: &Request{},
			},
			want:    req,
			wantErr: ErrInvalidID,
		},
		{
			name: "invalid project ID",
			fields: fields{
				r: &Request{
					id: NewID(),
				},
			},
			wantErr: ErrInvalidID,
		},
		{
			name: "invalid workspace ID",
			fields: fields{
				r: &Request{
					id:      NewID(),
					project: NewProjectID(),
				},
			},
			wantErr: ErrInvalidID,
		},
		{
			name: "invalid thread ID",
			fields: fields{
				r: &Request{
					id:        NewID(),
					project:   NewProjectID(),
					workspace: NewWorkspaceID(),
				},
			},
			wantErr: ErrInvalidID,
		},
		{
			name: "invalid user ID",
			fields: fields{
				r: &Request{
					id:        NewID(),
					project:   NewProjectID(),
					workspace: NewWorkspaceID(),
					thread:    NewThreadID(),
				},
			},
			wantErr: ErrInvalidID,
		},
		{
			name: "empty items",
			fields: fields{
				r: &Request{
					id:        NewID(),
					project:   NewProjectID(),
					workspace: NewWorkspaceID(),
					thread:    NewThreadID(),
					createdBy: NewUserID(),
				},
			},
			wantErr: ErrEmptyItems,
		},
		{
			name: "duplicated item",
			fields: fields{
				r: &Request{
					id:        NewID(),
					project:   NewProjectID(),
					workspace: NewWorkspaceID(),
					thread:    NewThreadID(),
					createdBy: NewUserID(),
					items:     ItemList{i1, i1},
				},
			},
			wantErr: ErrDuplicatedItem,
		},
		{
			name: "empty title",
			fields: fields{
				r: &Request{
					id:        NewID(),
					project:   NewProjectID(),
					workspace: NewWorkspaceID(),
					thread:    NewThreadID(),
					createdBy: NewUserID(),
					items: ItemList{{
						item:    NewItemID(),
						pointer: version.New().OrRef(),
					}},
				},
			},
			wantErr: ErrEmptyTitle,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			//tt.Parallel()
			b := &Builder{
				r: tc.fields.r,
			}
			got, err := b.Build()
			if tc.wantErr != nil {
				assert.Equal(tt, tc.wantErr, err)
				assert.Panics(tt, func() {
					_ = b.MustBuild()
				})
				return
			}
			assert.Equal(tt, tc.want, got)
			got = b.MustBuild()
			assert.Equal(tt, tc.want, got)
		})
	}
}

func TestBuilder_ID(t *testing.T) {
	b := &Builder{r: &Request{}}
	rid := NewID()
	b.ID(rid)
	assert.Equal(t, rid, b.r.ID())
}

func TestBuilder_ApprovedAt(t *testing.T) {
	b := &Builder{r: &Request{}}
	now := time.Now()
	b.ApprovedAt(&now)
	assert.Equal(t, &now, b.r.ApprovedAt())
}

func TestBuilder_ClosedAt(t *testing.T) {
	b := &Builder{r: &Request{}}
	now := time.Now()
	b.ClosedAt(&now)
	assert.Equal(t, &now, b.r.ClosedAt())
}

func TestBuilder_CreatedBy(t *testing.T) {
	b := &Builder{r: &Request{}}
	uid := NewUserID()
	b.CreatedBy(uid)
	assert.Equal(t, uid, b.r.CreatedBy())
}

func TestBuilder_Description(t *testing.T) {
	b := &Builder{r: &Request{}}
	desc := "foo"
	b.Description(desc)
	assert.Equal(t, desc, b.r.Description())
}

func TestBuilder_Items(t *testing.T) {
	b := &Builder{r: &Request{}}
	items := ItemList{{
		item:    NewItemID(),
		pointer: version.New().OrRef(),
	}}
	b.Items(items)
	assert.Equal(t, items, b.r.Items())
}

func TestBuilder_NewID(t *testing.T) {
	b := &Builder{r: &Request{}}
	b.NewID()
	assert.NotNil(t, b.r.ID())
}

func TestBuilder_Project(t *testing.T) {
	b := &Builder{r: &Request{}}
	pid := NewProjectID()
	b.Project(pid)
	assert.Equal(t, pid, b.r.Project())
}

func TestBuilder_Reviewers(t *testing.T) {
	b := &Builder{r: &Request{}}
	rev := UserIDList{NewUserID()}
	b.Reviewers(rev)
	assert.Equal(t, rev, b.r.Reviewers())
}

func TestBuilder_State(t *testing.T) {
	b := &Builder{r: &Request{}}
	s := StateWaiting
	b.State(s)
	assert.Equal(t, s, b.r.State())
}

func TestBuilder_Thread(t *testing.T) {
	b := &Builder{r: &Request{}}
	tid := NewThreadID()
	b.Thread(tid)
	assert.Equal(t, tid, b.r.Thread())
}

func TestBuilder_Title(t *testing.T) {
	b := &Builder{r: &Request{}}
	desc := "hoge"
	b.Title(desc)
	assert.Equal(t, desc, b.r.Title())
}

func TestBuilder_UpdatedAt(t *testing.T) {
	b := &Builder{r: &Request{}}
	now := time.Now()
	b.UpdatedAt(now)
	assert.Equal(t, now, b.r.UpdatedAt())
}

func TestBuilder_Workspace(t *testing.T) {
	b := &Builder{r: &Request{}}
	wid := NewWorkspaceID()
	b.Workspace(wid)
	assert.Equal(t, wid, b.r.Workspace())
}

func TestNew(t *testing.T) {
	b := New()
	assert.NotNil(t, b.r)
}
