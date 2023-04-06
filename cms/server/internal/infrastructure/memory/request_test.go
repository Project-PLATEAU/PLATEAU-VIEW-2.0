package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestRequest_Filtered(t *testing.T) {
	r := &Request{}
	pid := id.NewProjectID()

	assert.Equal(t, &Request{
		f: repo.ProjectFilter{
			Readable: id.ProjectIDList{pid},
			Writable: nil,
		},
	}, r.Filtered(repo.ProjectFilter{
		Readable: id.ProjectIDList{pid},
		Writable: nil,
	}))
}

func TestRequest_FindByID(t *testing.T) {
	ctx := context.Background()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())

	req := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(id.NewProjectID()).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()
	r := NewRequest()
	_ = r.Save(ctx, req)

	out, err := r.FindByID(ctx, req.ID())
	assert.NoError(t, err)
	assert.Equal(t, req, out)

	out2, err := r.FindByID(ctx, id.RequestID{})
	assert.Nil(t, out2)
	assert.Same(t, rerror.ErrNotFound, err)
}

func TestRequest_SaveAll(t *testing.T) {
	ctx := context.Background()
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())

	req1 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()

	req2 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()

	r := NewRequest()
	err := r.SaveAll(ctx, pid, request.List{req1, req2})
	assert.NoError(t, err)
	r = r.Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	})

	data, _ := r.FindByIDs(ctx, id.RequestIDList{req1.ID(), req2.ID()})
	assert.Equal(t, 2, len(data))

	wantErr := errors.New("test")
	SetRequestError(r, wantErr)
	assert.Same(t, wantErr, r.SaveAll(ctx, pid, request.List{}))
}

func TestRequest_Save(t *testing.T) {
	ctx := context.Background()
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())

	req1 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()

	req2 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(id.NewProjectID()).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()
	pf := repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	}
	r := NewRequest().Filtered(pf)

	_ = r.Save(ctx, req1)
	got, _ := r.FindByID(ctx, req1.ID())
	assert.Equal(t, req1, got)

	err := r.Save(ctx, req2)
	assert.Equal(t, repo.ErrOperationDenied, err)

	wantErr := errors.New("test")
	SetRequestError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, req1))
}

func TestRequest_FindByIDs(t *testing.T) {
	ctx := context.Background()
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())

	req1 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()

	req2 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()
	pf := repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	}
	r := NewRequest().Filtered(pf)
	_ = r.Save(ctx, req1)
	_ = r.Save(ctx, req2)

	got, err := r.FindByIDs(ctx, id.RequestIDList{req1.ID(), req2.ID()})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(got))
}

func TestRequest_FindByProject(t *testing.T) {
	ctx := context.Background()
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())

	req1 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()

	req2 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("xxx").
		State(request.StateDraft).
		MustBuild()

	req3 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(id.NewProjectID()).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()
	pf := repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	}
	r := NewRequest().Filtered(pf)
	_ = r.Save(ctx, req1)
	_ = r.Save(ctx, req2)
	_ = r.Save(ctx, req3)
	type args struct {
		id     id.ProjectID
		filter repo.RequestFilter
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find by project id only (find 2)",
			args: args{
				id: pid,
			},
			want: 2,
		},
		{
			name: "find by stat (find 1)",
			args: args{
				id: pid,
				filter: repo.RequestFilter{
					State: []request.State{request.StateDraft},
				},
			},
			want: 1,
		},
		{
			name: "find by title (find 1)",
			args: args{
				id: pid,
				filter: repo.RequestFilter{
					Keyword: lo.ToPtr("foo"),
				},
			},
			want: 1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			//tt.Parallel()
			got, _, _ := r.FindByProject(ctx, tc.args.id, tc.args.filter, nil, nil)

			assert.Equal(t, tc.want, len(got))
		})
	}

	wantErr := errors.New("test")
	SetRequestError(r, wantErr)
	_, _, err := r.FindByProject(ctx, pid, repo.RequestFilter{}, nil, nil)
	assert.Same(t, wantErr, err)
}

func TestRequest_FindByItem(t *testing.T) {
	ctx := context.Background()
	pid := id.NewProjectID()
	item1, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())
	item2, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())

	req1 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item1}).
		Title("foo").
		MustBuild()

	req2 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item1}).
		Title("xxx").
		State(request.StateDraft).
		MustBuild()

	req3 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item2}).
		Title("foo").
		MustBuild()
	pf := repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	}
	r := NewRequest().Filtered(pf)
	_ = r.Save(ctx, req1)
	_ = r.Save(ctx, req2)
	_ = r.Save(ctx, req3)
	tests := []struct {
		name  string
		input id.ItemIDList
		want  int
	}{
		{
			name:  "must find 2",
			input: id.ItemIDList{item1.Item()},
			want:  2,
		},
		{
			name:  "must find 1",
			input: id.ItemIDList{item2.Item()},
			want:  1,
		},
		{
			name:  "must find 0",
			input: id.ItemIDList{id.NewItemID()},
			want:  0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()
			got, _ := r.FindByItems(ctx, tc.input)

			assert.Equal(t, tc.want, len(got))
		})
	}

	wantErr := errors.New("test")
	SetRequestError(r, wantErr)
	_, err := r.FindByItems(ctx, id.ItemIDList{item1.Item()})
	assert.Same(t, wantErr, err)
}
