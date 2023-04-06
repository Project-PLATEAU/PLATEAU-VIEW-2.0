package mongo

import (
	"context"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestRequest_Filtered(t *testing.T) {
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
		Title("hoge").
		MustBuild()
	tests := []struct {
		name    string
		seeds   request.List
		args    repo.ProjectFilter
		wantErr error
	}{
		{
			name:  "operation denied",
			seeds: request.List{req1, req2},
			args: repo.ProjectFilter{
				Readable: []id.ProjectID{},
				Writable: []id.ProjectID{},
			},
			wantErr: repo.ErrOperationDenied,
		},
		{
			name:  "success",
			seeds: request.List{req1, req2},

			args: repo.ProjectFilter{
				Readable: []id.ProjectID{pid},
				Writable: []id.ProjectID{pid},
			},
		},
	}
	initDB := mongotest.Connect(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewRequest(client).Filtered(tt.args)
			ctx := context.Background()
			for _, p := range tt.seeds {
				err := r.Save(ctx, p)
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestRequest_FindByID(t *testing.T) {
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
		Title("hoge").
		MustBuild()
	tests := []struct {
		name    string
		seeds   request.List
		args    request.ID
		want    *request.Request
		wantErr error
	}{
		{
			name:  "success",
			seeds: request.List{req1, req2},
			args:  req1.ID(),
			want:  req1,
		},
		{
			name:    "not found",
			seeds:   request.List{req1, req2},
			args:    id.NewRequestID(),
			wantErr: rerror.ErrNotFound,
		},
	}
	initDB := mongotest.Connect(t)
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewRequest(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}
			got, err := r.FindByID(ctx, tc.args)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.want.ID(), got.ID())
		})
	}
}

func TestRequest_FindByIDs(t *testing.T) {
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
		Title("hoge").
		MustBuild()

	tests := []struct {
		name  string
		seeds request.List
		args  id.RequestIDList
		want  int
	}{
		{
			name:  "must find 2",
			seeds: request.List{req1, req2},
			args:  id.RequestIDList{req1.ID(), req2.ID()},
			want:  2,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2},
			args:  id.RequestIDList{id.NewRequestID(), req1.ID()},
			want:  1,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewRequest(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}
			got, _ := r.FindByIDs(ctx, tc.args)
			assert.Equal(t, tc.want, len(got))
		})
	}
}

func TestRequest_FindByProject(t *testing.T) {
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())
	reviewer := id.NewUserID()
	creator := id.NewUserID()
	req1 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(creator).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Reviewers(id.UserIDList{reviewer}).
		Title("foo").
		MustBuild()
	req2 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		State(request.StateDraft).
		Title("hoge").
		MustBuild()
	type args struct {
		projectID id.ProjectID
		repo.RequestFilter
	}
	tests := []struct {
		name  string
		seeds request.List
		args  args
		want  int
	}{
		{
			name:  "must find 2",
			seeds: request.List{req1, req2},
			args: args{
				projectID: pid,
			},
			want: 2,
		},
		{
			name:  "must find 0",
			seeds: request.List{req1, req2},
			args: args{
				projectID: id.NewProjectID(),
			},
			want: 0,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2},
			args: args{
				projectID: pid,
				RequestFilter: repo.RequestFilter{
					Keyword: lo.ToPtr("foo"),
				},
			},
			want: 1,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2},
			args: args{
				projectID: pid,
				RequestFilter: repo.RequestFilter{
					State: []request.State{request.StateDraft},
				},
			},
			want: 1,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2},
			args: args{
				projectID: pid,
				RequestFilter: repo.RequestFilter{
					Reviewer: reviewer.Ref(),
				},
			},
			want: 1,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2},
			args: args{
				projectID: pid,
				RequestFilter: repo.RequestFilter{
					CreatedBy: creator.Ref(),
				},
			},
			want: 1,
		},
		{
			name:  "must find 0",
			seeds: request.List{req1, req2},
			args: args{
				projectID: pid,
				RequestFilter: repo.RequestFilter{
					Keyword: lo.ToPtr("foo"),
					State:   []request.State{request.StateDraft},
				},
			},
			want: 0,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewRequest(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}
			got, _, _ := r.FindByProject(ctx, tc.args.projectID, tc.args.RequestFilter, &usecasex.Sort{}, usecasex.CursorPagination{First: lo.ToPtr(int64(10))}.Wrap())
			assert.Equal(t, tc.want, len(got))
		})
	}
}

func TestRequest_SaveAll(t *testing.T) {
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
		Title("hoge").
		MustBuild()
	req3 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("xxx").
		MustBuild()

	initDB := mongotest.Connect(t)

	client := mongox.NewClientWithDatabase(initDB(t))

	r := NewRequest(client)
	ctx := context.Background()
	err := r.SaveAll(ctx, pid, request.List{req1, req2, req3})
	assert.NoError(t, err)
	got, _ := r.FindByIDs(ctx, id.RequestIDList{req1.ID(), req2.ID(), req3.ID()})
	assert.Equal(t, 3, len(got))
}

func TestRequest_FindByItem(t *testing.T) {
	pid := id.NewProjectID()
	item1, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())
	item2, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())
	reviewer := id.NewUserID()
	creator := id.NewUserID()
	req1 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(creator).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item1}).
		Reviewers(id.UserIDList{reviewer}).
		Title("foo").
		MustBuild()
	req2 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item1, item2}).
		State(request.StateDraft).
		Title("hoge").
		MustBuild()
	req3 := request.New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item2}).
		Title("xxx").
		MustBuild()

	tests := []struct {
		name  string
		seeds request.List
		input id.ItemIDList
		want  int
	}{
		{
			name:  "must find 2",
			seeds: request.List{req1, req2, req3},
			input: id.ItemIDList{item1.Item(), item2.Item()},
			want:  3,
		},
		{
			name:  "must find 0",
			seeds: request.List{req1, req2, req3},
			input: id.ItemIDList{id.NewItemID()},
			want:  0,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2, req3},
			input: id.ItemIDList{item1.Item()},
			want:  2,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			client := mongox.NewClientWithDatabase(initDB(t))

			r := NewRequest(client)
			ctx := context.Background()
			for _, p := range tc.seeds {
				err := r.Save(ctx, p)
				assert.NoError(t, err)
			}
			got, _ := r.FindByItems(ctx, tc.input)
			assert.Equal(t, tc.want, len(got))
		})
	}
}
