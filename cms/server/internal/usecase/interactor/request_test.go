package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestRequest_FindByID(t *testing.T) {
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())
	wid := id.NewWorkspaceID()

	req1 := request.New().
		NewID().
		Workspace(wid).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()
	req2 := request.New().
		NewID().
		Workspace(wid).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("hoge").
		MustBuild()
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	op := &usecase.Operator{
		User: lo.ToPtr(u.ID()),
	}

	tests := []struct {
		name  string
		seeds request.List
		args  struct {
			id       id.RequestID
			operator *usecase.Operator
		}
		want           *request.Request
		mockRequestErr bool
		wantErr        error
	}{
		{
			name:  "find 1 of 2",
			seeds: request.List{req1, req2},
			args: struct {
				id       id.RequestID
				operator *usecase.Operator
			}{
				id:       req1.ID(),
				operator: op,
			},
			want:    req1,
			wantErr: nil,
		},
		{
			name:  "find 1 of 0",
			seeds: request.List{},
			args: struct {
				id       id.RequestID
				operator *usecase.Operator
			}{
				id:       req1.ID(),
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockRequestErr {
				memory.SetRequestError(db.Request, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Request.Save(ctx, p)
				assert.NoError(t, err)
			}
			requestUC := NewRequest(db, nil)

			got, err := requestUC.FindByID(ctx, tc.args.id, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRequest_FindByIDs(t *testing.T) {
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())
	wid := id.NewWorkspaceID()

	req1 := request.New().
		NewID().
		Workspace(wid).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()
	req2 := request.New().
		NewID().
		Workspace(wid).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("hoge").
		MustBuild()
	req3 := request.New().
		NewID().
		Workspace(wid).
		Project(pid).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("xxx").
		MustBuild()
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	op := &usecase.Operator{
		User: lo.ToPtr(u.ID()),
	}

	tests := []struct {
		name  string
		seeds request.List
		args  struct {
			ids      id.RequestIDList
			operator *usecase.Operator
		}
		want           int
		mockRequestErr bool
		wantErr        error
	}{
		{
			name:  "find 2 of 3",
			seeds: request.List{req1, req2, req3},
			args: struct {
				ids      id.RequestIDList
				operator *usecase.Operator
			}{
				ids:      id.RequestIDList{req1.ID(), req2.ID()},
				operator: op,
			},
			want: 2,
		},
		{
			name:  "find 0 of 3",
			seeds: request.List{req1, req2, req3},
			args: struct {
				ids      id.RequestIDList
				operator *usecase.Operator
			}{
				ids:      id.RequestIDList{id.NewRequestID()},
				operator: op,
			},
			want: 0,
		},
		{
			name:           "mock error",
			mockRequestErr: true,
			wantErr:        errors.New("test"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockRequestErr {
				memory.SetRequestError(db.Request, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Request.Save(ctx, p)
				assert.NoError(t, err)
			}
			requestUC := NewRequest(db, nil)

			got, err := requestUC.FindByIDs(ctx, tc.args.ids, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, len(got))
		})
	}
}

func TestRequest_FindByProject(t *testing.T) {
	pid := id.NewProjectID()
	item, _ := request.NewItemWithVersion(id.NewItemID(), version.New().OrRef())
	wid := id.NewWorkspaceID()

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
		State(request.StateDraft).
		Title("hoge").
		MustBuild()
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	op := &usecase.Operator{
		User: lo.ToPtr(u.ID()),
	}

	tests := []struct {
		name  string
		seeds request.List
		args  struct {
			pid      id.ProjectID
			filter   interfaces.RequestFilter
			operator *usecase.Operator
		}
		want           int
		mockRequestErr bool
		wantErr        error
	}{
		{
			name:  "must find 2",
			seeds: request.List{req1, req2},
			args: struct {
				pid      id.ProjectID
				filter   interfaces.RequestFilter
				operator *usecase.Operator
			}{
				pid:      pid,
				operator: op,
			},
			want: 2,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2},
			args: struct {
				pid      id.ProjectID
				filter   interfaces.RequestFilter
				operator *usecase.Operator
			}{
				pid: pid,
				filter: interfaces.RequestFilter{
					Keyword: lo.ToPtr("foo"),
				},
				operator: op,
			},
			want: 1,
		},
		{
			name:  "must find 1",
			seeds: request.List{req1, req2},
			args: struct {
				pid      id.ProjectID
				filter   interfaces.RequestFilter
				operator *usecase.Operator
			}{
				pid: pid,
				filter: interfaces.RequestFilter{
					State: []request.State{request.StateDraft},
				},
				operator: op,
			},
			want: 1,
		},
		{
			name:           "mock error",
			mockRequestErr: true,
			wantErr:        errors.New("test"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockRequestErr {
				memory.SetRequestError(db.Request, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Request.Save(ctx, p)
				assert.NoError(t, err)
			}
			requestUC := NewRequest(db, nil)

			got, _, err := requestUC.FindByProject(ctx, tc.args.pid, tc.args.filter, nil, nil, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, len(got))
		})
	}
}

func TestRequest_Approve(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()

	// TODO: add error cases
	prj := project.New().NewID().MustBuild()
	s := schema.New().NewID().Workspace(id.NewWorkspaceID()).Project(prj.ID()).MustBuild()
	m := model.New().NewID().Schema(s.ID()).RandomKey().MustBuild()
	i := item.New().NewID().Schema(s.ID()).Model(m.ID()).Project(prj.ID()).Thread(id.NewThreadID()).MustBuild()
	item, _ := request.NewItem(i.ID())
	wid := id.NewWorkspaceID()
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	req1 := request.New().
		NewID().
		Workspace(wid).
		Project(prj.ID()).
		Reviewers(id.UserIDList{u.ID()}).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{item}).
		Title("foo").
		MustBuild()
	op := &usecase.Operator{
		User:             lo.ToPtr(u.ID()),
		OwningWorkspaces: id.WorkspaceIDList{wid},
	}
	ctx := context.Background()

	db := memory.New()
	// if tc.mockRequestErr {
	//	memory.SetRequestError(db.Request, tc.wantErr)
	// }
	err := db.Project.Save(ctx, prj)
	assert.NoError(t, err)
	err = db.Request.Save(ctx, req1)
	assert.NoError(t, err)
	err = db.Schema.Save(ctx, s)
	assert.NoError(t, err)
	err = db.Model.Save(ctx, m)
	assert.NoError(t, err)
	err = db.Item.Save(ctx, i)
	assert.NoError(t, err)

	requestUC := NewRequest(db, nil)
	_, err = requestUC.Approve(ctx, req1.ID(), op)
	assert.NoError(t, err)

	itemUC := NewItem(db, nil)
	itm, err := itemUC.FindByID(ctx, i.ID(), op)
	assert.NoError(t, err)
	expected := version.MustBeValue(itm.Version(), nil, version.NewRefs(version.Public, version.Latest), now, i)
	assert.Equal(t, expected, itm)
}
