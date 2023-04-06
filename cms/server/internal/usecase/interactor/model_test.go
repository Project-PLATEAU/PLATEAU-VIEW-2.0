package interactor

import (
	"context"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/stretchr/testify/assert"
)

func TestModel_CheckKey(t *testing.T) {
	mockTime := time.Now()
	pId := id.NewProjectID()
	type args struct {
		pId id.ProjectID
		s   string
	}
	type seeds struct {
		model   model.List
		project project.List
	}
	tests := []struct {
		name    string
		seeds   seeds
		args    args
		want    bool
		mockErr bool
		wantErr error
	}{
		{
			name:  "in empty db",
			seeds: seeds{},
			args: args{
				pId: id.NewProjectID(),
				s:   "test123",
			},
			want:    true,
			mockErr: false,
			wantErr: nil,
		},
		{
			name: "with different key",
			seeds: seeds{
				model: []*model.Model{
					model.New().NewID().Key(key.Random()).Project(pId).Schema(id.NewSchemaID()).MustBuild(),
				},
			},
			args: args{
				pId: pId,
				s:   "test123",
			},
			want:    true,
			mockErr: false,
			wantErr: nil,
		},
		{
			name: "with same key",
			seeds: seeds{
				model: []*model.Model{
					model.New().NewID().Key(key.New("test123")).Project(pId).Schema(id.NewSchemaID()).MustBuild(),
				},
			},
			args: args{
				pId: pId,
				s:   "test123",
			},
			want:    false,
			mockErr: false,
			wantErr: nil,
		},
		{
			name: "with same key different project",
			seeds: seeds{
				model: []*model.Model{
					model.New().NewID().Key(key.New("test123")).Project(pId).Schema(id.NewSchemaID()).MustBuild(),
				},
			},
			args: args{
				pId: id.NewProjectID(),
				s:   "test123",
			},
			want:    true,
			mockErr: false,
			wantErr: nil,
		},
		{
			name:  "with invalid key",
			seeds: seeds{},
			args: args{
				pId: id.NewProjectID(),
				s:   "12",
			},
			want:    false,
			mockErr: true,
			wantErr: model.ErrInvalidKey,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tt.mockErr {
				memory.SetModelError(db.Model, tt.wantErr)
			}
			defer memory.MockNow(db, mockTime)()
			for _, m := range tt.seeds.model {
				err := db.Model.Save(ctx, m.Clone())
				assert.NoError(t, err)
			}
			for _, p := range tt.seeds.project {
				err := db.Project.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			u := NewModel(db, nil)

			got, err := u.CheckKey(ctx, tt.args.pId, tt.args.s)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.False(t, got)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModel_Create(t *testing.T) {
	mockTime := time.Now()
	// mId := id.NewModelID()
	// sId := id.NewSchemaID()
	// wid1 := id.NewWorkspaceID()
	// wid2 := id.NewWorkspaceID()
	//
	// pid1 := id.NewProjectID()
	// p1 := project.New().ID(pid1).Workspace(wid1).UpdatedAt(mockTime).MustBuild()
	//
	// pid2 := id.NewProjectID()
	// p2 := project.New().ID(pid2).Workspace(wid2).UpdatedAt(mockTime).MustBuild()
	//
	// u := user.New().NewID().Email("aaa@bbb.com").Workspace(wid1).MustBuild()
	// op := &usecase.Operator{
	// 	User:               u.ID(),
	// 	ReadableWorkspaces: []id.WorkspaceID{wid1, wid2},
	// 	WritableWorkspaces: []id.WorkspaceID{wid1},
	// }

	type args struct {
		param    interfaces.CreateModelParam
		operator *usecase.Operator
	}

	type seeds struct {
		model   model.List
		project project.List
	}
	tests := []struct {
		name    string
		seeds   seeds
		args    args
		want    *model.Model
		mockErr bool
		wantErr error
	}{
		// TODO: fix
		// {
		// 	name: "add",
		// 	seeds: seeds{
		// 		model:   nil,
		// 		project: []*project.Project{p1, p2},
		// 	},
		// 	args: args{
		// 		param: interfaces.CreateModelParam{
		// 			ProjectId:   pid1,
		// 			Name:        lo.ToPtr("m1"),
		// 			Description: lo.ToPtr("m1"),
		// 			Key:         lo.ToPtr("k123456"),
		// 			Public:      lo.ToPtr(true),
		// 		},
		// 		operator: op,
		// 	},
		// 	want:    model.New().ID(mId).Schema(sId).Project(pid1).Name("m1").Description("m1").Key(key.New("k123456")).Public(true).UpdatedAt(mockTime).MustBuild(),
		// 	mockErr: false,
		// 	wantErr: nil,
		// },
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tt.mockErr {
				memory.SetModelError(db.Model, tt.wantErr)
			}
			defer memory.MockNow(db, mockTime)()
			for _, m := range tt.seeds.model {
				err := db.Model.Save(ctx, m.Clone())
				assert.NoError(t, err)
			}
			for _, p := range tt.seeds.project {
				err := db.Project.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			u := NewModel(db, nil)

			got, err := u.Create(ctx, tt.args.param, tt.args.operator)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.Nil(t, got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModel_Delete(t *testing.T) {
	mockTime := time.Now()
	type args struct {
		modelID  id.ModelID
		operator *usecase.Operator
	}
	type seeds struct {
		model   model.List
		project project.List
	}
	tests := []struct {
		name    string
		seeds   seeds
		args    args
		mockErr bool
		wantErr error
	}{
		// {},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tt.mockErr {
				memory.SetModelError(db.Model, tt.wantErr)
			}
			defer memory.MockNow(db, mockTime)()
			for _, m := range tt.seeds.model {
				err := db.Model.Save(ctx, m.Clone())
				assert.NoError(t, err)
			}
			for _, p := range tt.seeds.project {
				err := db.Project.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			u := NewModel(db, nil)

			assert.Equal(t, tt.wantErr, u.Delete(ctx, tt.args.modelID, tt.args.operator))
		})
	}
}

func TestModel_FindByIDs(t *testing.T) {
	mockTime := time.Now()
	type args struct {
		ids      []id.ModelID
		operator *usecase.Operator
	}
	type seeds struct {
		model   model.List
		project project.List
	}
	tests := []struct {
		name    string
		seeds   seeds
		args    args
		want    model.List
		mockErr bool
		wantErr error
	}{
		{},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tt.mockErr {
				memory.SetModelError(db.Model, tt.wantErr)
			}
			defer memory.MockNow(db, mockTime)()
			for _, m := range tt.seeds.model {
				err := db.Model.Save(ctx, m.Clone())
				assert.NoError(t, err)
			}
			for _, p := range tt.seeds.project {
				err := db.Project.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			u := NewModel(db, nil)
			got, err := u.FindByIDs(ctx, tt.args.ids, tt.args.operator)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModel_Publish(t *testing.T) {
	mockTime := time.Now()
	type args struct {
		modelID  id.ModelID
		b        bool
		operator *usecase.Operator
	}
	type seeds struct {
		model   model.List
		project project.List
	}
	tests := []struct {
		name    string
		seeds   seeds
		args    args
		want    bool
		mockErr bool
		wantErr error
	}{
		{},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tt.mockErr {
				memory.SetModelError(db.Model, tt.wantErr)
			}
			defer memory.MockNow(db, mockTime)()
			for _, m := range tt.seeds.model {
				err := db.Model.Save(ctx, m.Clone())
				assert.NoError(t, err)
			}
			for _, p := range tt.seeds.project {
				err := db.Project.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			u := NewModel(db, nil)

			got, err := u.Publish(ctx, tt.args.modelID, tt.args.b, tt.args.operator)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModel_Update(t *testing.T) {
	mockTime := time.Now()
	type args struct {
		param    interfaces.UpdateModelParam
		operator *usecase.Operator
	}
	type seeds struct {
		model   model.List
		project project.List
	}
	tests := []struct {
		name    string
		seeds   seeds
		args    args
		want    *model.Model
		mockErr bool
		wantErr error
	}{
		{},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tt.mockErr {
				memory.SetModelError(db.Model, tt.wantErr)
			}
			defer memory.MockNow(db, mockTime)()
			for _, m := range tt.seeds.model {
				err := db.Model.Save(ctx, m.Clone())
				assert.NoError(t, err)
			}
			for _, p := range tt.seeds.project {
				err := db.Project.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			u := NewModel(db, nil)

			got, err := u.Update(ctx, tt.args.param, tt.args.operator)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewModel(t *testing.T) {
	type args struct {
		r *repo.Container
	}
	tests := []struct {
		name string
		args args
		want interfaces.Model
	}{
		// {},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, NewModel(tt.args.r, nil))
		})
	}
}
