package interactor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearth-cms/server/pkg/model"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	r := repo.Container{}
	i := NewItem(&r, nil)
	assert.NotNil(t, i)
}

func TestItem_FindByID(t *testing.T) {
	sid := id.NewSchemaID()
	id1 := id.NewItemID()
	i1 := item.New().ID(id1).Schema(sid).Model(id.NewModelID()).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()
	id2 := id.NewItemID()
	i2 := item.New().ID(id2).Schema(sid).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()

	wid := id.NewWorkspaceID()
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	op := &usecase.Operator{
		User: lo.ToPtr(u.ID()),
	}

	tests := []struct {
		name  string
		seeds item.List
		args  struct {
			id       id.ItemID
			operator *usecase.Operator
		}
		want        *item.Item
		mockItemErr bool
		wantErr     error
	}{
		{
			name:  "find 1 of 2",
			seeds: item.List{i1, i2},
			args: struct {
				id       id.ItemID
				operator *usecase.Operator
			}{
				id:       id1,
				operator: op,
			},
			want:    i1,
			wantErr: nil,
		},
		{
			name:  "find 1 of 0",
			seeds: item.List{},
			args: struct {
				id       id.ItemID
				operator *usecase.Operator
			}{
				id:       id1,
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
			if tc.mockItemErr {
				memory.SetItemError(db.Item, tc.wantErr)
			}
			for _, p := range tc.seeds {
				err := db.Item.Save(ctx, p)
				assert.NoError(t, err)
			}
			itemUC := NewItem(db, nil)
			itemUC.ignoreEvent = true

			got, err := itemUC.FindByID(ctx, tc.args.id, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got.Value())
		})
	}
}

func TestItem_FindBySchema(t *testing.T) {
	uid := id.NewUserID()
	wid := id.NewWorkspaceID()
	pid := id.NewProjectID()
	sf1 := schema.NewField(schema.NewBool().TypeProperty()).NewID().Key(key.Random()).MustBuild()
	s1 := schema.New().NewID().Workspace(wid).Project(pid).Fields(schema.FieldList{sf1}).MustBuild()
	s2 := schema.New().NewID().Workspace(wid).Project(pid).MustBuild()
	restore := util.MockNow(time.Now().Truncate(time.Millisecond).UTC())
	i1 := item.New().NewID().
		Schema(s1.ID()).
		Model(id.NewModelID()).
		Project(pid).
		Fields([]*item.Field{
			item.NewField(sf1.ID(), value.TypeBool.Value(true).AsMultiple()),
		}).
		Thread(id.NewThreadID()).
		MustBuild()
	restore()
	restore = util.MockNow(time.Now().Truncate(time.Millisecond).Add(time.Second).UTC())
	i2 := item.New().NewID().
		Schema(s1.ID()).
		Model(id.NewModelID()).
		Project(pid).
		Fields([]*item.Field{
			item.NewField(sf1.ID(), value.TypeBool.Value(true).AsMultiple()),
		}).
		Thread(id.NewThreadID()).
		MustBuild()
	restore()
	restore = util.MockNow(time.Now().Truncate(time.Millisecond).Add(time.Second * 2).UTC())
	i3 := item.New().NewID().
		Schema(s2.ID()).
		Model(id.NewModelID()).
		Project(pid).Thread(id.NewThreadID()).MustBuild()
	restore()

	type args struct {
		schema     id.SchemaID
		operator   *usecase.Operator
		pagination *usecasex.Pagination
	}

	tests := []struct {
		name        string
		seedItems   item.List
		seedSchema  *schema.Schema
		args        args
		want        int
		wantErr     error
		mockItemErr bool
	}{
		{
			name:       "find 2 of 3",
			seedItems:  item.List{i1, i2, i3},
			seedSchema: s1,
			args: args{
				schema: s1.ID(),
				operator: &usecase.Operator{
					User:             &uid,
					ReadableProjects: []id.ProjectID{pid},
					WritableProjects: []id.ProjectID{pid},
				},
			},
			want:    2,
			wantErr: nil,
		},
		{
			name:       "items not found",
			seedItems:  item.List{},
			seedSchema: s1,
			args: args{
				schema: s1.ID(),
				operator: &usecase.Operator{
					User:             &uid,
					ReadableProjects: []id.ProjectID{pid},
					WritableProjects: []id.ProjectID{pid},
				},
			},
			want:    0,
			wantErr: nil,
		},
		{
			name:       "schema not found",
			seedItems:  item.List{i1, i2, i3},
			seedSchema: s2,
			args: args{
				schema: s1.ID(),
				operator: &usecase.Operator{
					User:             &uid,
					ReadableProjects: []id.ProjectID{pid},
					WritableProjects: []id.ProjectID{pid},
				},
			},
			want:    0,
			wantErr: rerror.ErrNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockItemErr {
				memory.SetItemError(db.Item, tc.wantErr)
			}

			for _, seed := range tc.seedItems {
				err := db.Item.Save(ctx, seed)
				assert.NoError(t, err)
			}
			if tc.seedSchema != nil {
				err := db.Schema.Save(ctx, tc.seedSchema)
				assert.NoError(t, err)
			}

			itemUC := NewItem(db, nil)
			itemUC.ignoreEvent = true

			got, _, err := itemUC.FindBySchema(ctx, tc.args.schema, nil, tc.args.pagination, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, len(got))
		})
	}
}

func TestItem_FindAllVersionsByID(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()

	sid := id.NewSchemaID()
	id1 := id.NewItemID()
	i1 := item.New().ID(id1).Project(id.NewProjectID()).Schema(sid).Model(id.NewModelID()).Thread(id.NewThreadID()).MustBuild()

	wid := id.NewWorkspaceID()
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	op := &usecase.Operator{
		User: lo.ToPtr(u.ID()),
	}
	ctx := context.Background()

	db := memory.New()
	err := db.Item.Save(ctx, i1)
	assert.NoError(t, err)

	itemUC := NewItem(db, nil)
	itemUC.ignoreEvent = true

	// first version
	res, err := itemUC.FindAllVersionsByID(ctx, id1, op)
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.NewValue(res[0].Version(), nil, version.NewRefs(version.Latest), now, i1),
	}, res)

	// second version
	err = db.Item.Save(ctx, i1)
	assert.NoError(t, err)

	res, err = itemUC.FindAllVersionsByID(ctx, id1, op)
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.NewValue(res[0].Version(), nil, nil, now, i1),
		version.NewValue(res[1].Version(), version.NewVersions(res[0].Version()), version.NewRefs(version.Latest), now, i1),
	}, res)

	// not found
	res, err = itemUC.FindAllVersionsByID(ctx, id.NewItemID(), op)
	assert.NoError(t, err)
	assert.Empty(t, res)

	// mock item error
	wantErr := errors.New("test")
	memory.SetItemError(db.Item, wantErr)
	item2, err := itemUC.FindAllVersionsByID(ctx, id1, op)
	assert.Nil(t, item2)
	assert.Equal(t, wantErr, err)
}

func TestItem_FindByProject(t *testing.T) {
	pid1 := id.NewProjectID()
	pid2 := id.NewProjectID()
	wid := id.NewWorkspaceID()
	s1 := project.New().ID(pid1).Workspace(wid).MustBuild()
	s2 := project.New().ID(pid2).Workspace(wid).MustBuild()
	i1 := item.New().NewID().
		Project(pid1).
		Schema(id.NewSchemaID()).
		Model(id.NewModelID()).
		Thread(id.NewThreadID()).
		Timestamp(time.Now().Truncate(time.Millisecond).UTC()).
		MustBuild()
	i2 := item.New().NewID().
		Project(pid1).
		Schema(id.NewSchemaID()).
		Model(id.NewModelID()).
		Thread(id.NewThreadID()).
		Timestamp(time.Now().Truncate(time.Millisecond).Add(time.Second).UTC()).
		MustBuild()
	i3 := item.New().NewID().
		Project(pid2).
		Schema(id.NewSchemaID()).
		Model(id.NewModelID()).
		Thread(id.NewThreadID()).
		Timestamp(time.Now().Truncate(time.Millisecond).Add(time.Second * 2).UTC()).
		MustBuild()

	u := user.New().NewID().Email("aaa@bbb.com").Name("foo").Workspace(wid).MustBuild()
	op := &usecase.Operator{
		User:             lo.ToPtr(u.ID()),
		ReadableProjects: []id.ProjectID{pid1, pid2},
	}

	type args struct {
		id         id.ProjectID
		operator   *usecase.Operator
		pagination *usecasex.Pagination
	}

	tests := []struct {
		name        string
		seedItems   item.List
		seedProject *project.Project
		args        args
		want        int
		mockItemErr bool
		wantErr     error
	}{
		{
			name:        "find 2 of 3",
			seedItems:   item.List{i1, i2, i3},
			seedProject: s1,
			args: args{
				id:       pid1,
				operator: op,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name:        "items not found",
			seedItems:   item.List{},
			seedProject: s1,
			args: args{
				id:       pid1,
				operator: op,
			},
			want:    0,
			wantErr: nil,
		},
		{
			name:        "project not found",
			seedItems:   item.List{i1, i2, i3},
			seedProject: s2,
			args: args{
				id:       id.NewProjectID(),
				operator: op,
			},
			want:    0,
			wantErr: rerror.ErrNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockItemErr {
				memory.SetItemError(db.Item, tc.wantErr)
			}
			for _, seed := range tc.seedItems {
				err := db.Item.Save(ctx, seed)
				assert.NoError(t, err)
			}
			err := db.Project.Save(ctx, tc.seedProject)
			assert.NoError(t, err)
			itemUC := NewItem(db, nil)
			itemUC.ignoreEvent = true

			got, _, err := itemUC.FindByProject(ctx, tc.args.id, tc.args.pagination, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, len(got.Unwrap()))
		})
	}
}

func TestItem_Search(t *testing.T) {
	sid1 := id.NewSchemaID()
	sf1 := id.NewFieldID()
	sf2 := id.NewFieldID()
	f1 := item.NewField(sf1, value.TypeText.Value("foo").AsMultiple())
	f2 := item.NewField(sf2, value.TypeText.Value("hoge").AsMultiple())
	id1 := id.NewItemID()
	pid := id.NewProjectID()
	i1 := item.New().ID(id1).Schema(sid1).Model(id.NewModelID()).Project(pid).Fields([]*item.Field{f1}).Thread(id.NewThreadID()).MustBuild()
	id2 := id.NewItemID()
	i2 := item.New().ID(id2).Schema(sid1).Model(id.NewModelID()).Project(pid).Fields([]*item.Field{f1}).Thread(id.NewThreadID()).MustBuild()
	id3 := id.NewItemID()
	i3 := item.New().ID(id3).Schema(sid1).Model(id.NewModelID()).Project(pid).Fields([]*item.Field{f2}).Thread(id.NewThreadID()).MustBuild()

	wid := id.NewWorkspaceID()
	u := user.New().NewID().Email("aaa@bbb.com").Workspace(wid).Name("foo").MustBuild()
	op := &usecase.Operator{
		User: lo.ToPtr(u.ID()),
	}

	tests := []struct {
		name  string
		seeds struct {
			items item.List
		}
		args struct {
			query    *item.Query
			operator *usecase.Operator
		}
		want        int
		mockItemErr bool
		wantErr     error
	}{
		{
			name: "find 2 of 3",
			seeds: struct {
				items item.List
			}{
				items: item.List{i1, i2, i3},
			},
			args: struct {
				query    *item.Query
				operator *usecase.Operator
			}{
				query:    item.NewQuery(pid, nil, "foo", nil),
				operator: op,
			},
			want:    2,
			wantErr: nil,
		},
		{
			name: "find 1 of 3",
			seeds: struct {
				items item.List
			}{
				items: item.List{i1, i2, i3},
			},
			args: struct {
				query    *item.Query
				operator *usecase.Operator
			}{
				query:    item.NewQuery(pid, nil, "hoge", nil),
				operator: op,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "items not found",
			seeds: struct {
				items item.List
			}{
				items: item.List{i1, i2, i3},
			},
			args: struct {
				query    *item.Query
				operator *usecase.Operator
			}{
				query:    item.NewQuery(pid, nil, "xxx", nil),
				operator: op,
			},
			want:    0,
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			if tc.mockItemErr {
				memory.SetItemError(db.Item, tc.wantErr)
			}
			for _, seed := range tc.seeds.items {
				err := db.Item.Save(ctx, seed)
				assert.Nil(t, err)
			}
			itemUC := NewItem(db, nil)
			itemUC.ignoreEvent = true

			got, _, err := itemUC.Search(ctx, tc.args.query, nil, nil, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, len(got))

		})
	}
}

func TestItem_Create(t *testing.T) {
	prj := project.New().NewID().MustBuild()
	sf := schema.NewField(schema.NewText(lo.ToPtr(10)).TypeProperty()).NewID().Name("f").Unique(true).Key(key.Random()).MustBuild()
	s := schema.New().NewID().Workspace(id.NewWorkspaceID()).Project(prj.ID()).Fields(schema.FieldList{sf}).MustBuild()
	m := model.New().NewID().Schema(s.ID()).Key(key.Random()).Project(s.Project()).MustBuild()

	ctx := context.Background()
	db := memory.New()
	lo.Must0(db.Project.Save(ctx, prj))
	lo.Must0(db.Schema.Save(ctx, s))
	lo.Must0(db.Model.Save(ctx, m))
	itemUC := NewItem(db, nil)
	itemUC.ignoreEvent = true

	op := &usecase.Operator{
		User:               id.NewUserID().Ref(),
		ReadableProjects:   []id.ProjectID{s.Project()},
		WritableProjects:   []id.ProjectID{s.Project()},
		ReadableWorkspaces: []id.WorkspaceID{s.Workspace()},
		WritableWorkspaces: []id.WorkspaceID{s.Workspace()},
	}

	// ok
	item, err := itemUC.Create(ctx, interfaces.CreateItemParam{
		SchemaID: s.ID(),
		ModelID:  m.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "xxx",
			},
		},
	}, op)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, s.ID(), item.Value().Schema())

	it, err := db.Item.FindByID(ctx, item.Value().ID(), nil)
	assert.NoError(t, err)
	assert.Equal(t, item, it)
	assert.Equal(t, value.TypeText.Value("xxx").AsMultiple(), it.Value().Field(sf.ID()).Value())

	// validate fails
	item, err = itemUC.Create(ctx, interfaces.CreateItemParam{
		SchemaID: s.ID(),
		ModelID:  m.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "abcabcabcabc", // too long
			},
		},
	}, op)
	assert.ErrorContains(t, err, "it sholud be shorter than 10")
	assert.Nil(t, item)

	// duplicated
	item, err = itemUC.Create(ctx, interfaces.CreateItemParam{
		SchemaID: s.ID(),
		ModelID:  m.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "xxx", // duplicated
			},
		},
	}, op)
	assert.Equal(t, interfaces.ErrDuplicatedItemValue, err)
	assert.Nil(t, item)

	// required
	sf.SetRequired(true)
	s.RemoveField(sf.ID())
	s.AddField(sf)
	lo.Must0(db.Schema.Save(ctx, s))
	item, err = itemUC.Create(ctx, interfaces.CreateItemParam{
		SchemaID: s.ID(),
		ModelID:  m.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "",
			},
		},
	}, op)
	assert.ErrorIs(t, err, schema.ErrValueRequired)
	assert.Nil(t, item)

	// mock item error
	wantErr := errors.New("test")
	memory.SetItemError(db.Item, wantErr)
	item, err = itemUC.Create(ctx, interfaces.CreateItemParam{
		SchemaID: s.ID(),
		ModelID:  m.ID(),
		Fields:   nil,
	}, op)
	assert.Equal(t, wantErr, err)
	assert.Nil(t, item)
}

func TestItem_Update(t *testing.T) {
	uId := id.NewUserID().Ref()
	prj := project.New().NewID().MustBuild()
	sf := schema.NewField(schema.NewText(lo.ToPtr(10)).TypeProperty()).NewID().Name("f").Unique(true).Key(key.Random()).MustBuild()
	s := schema.New().NewID().Workspace(id.NewWorkspaceID()).Project(prj.ID()).Fields(schema.FieldList{sf}).MustBuild()
	m := model.New().NewID().Schema(s.ID()).Key(key.Random()).Project(s.Project()).MustBuild()
	i := item.New().NewID().User(*uId).Model(m.ID()).Project(s.Project()).Schema(s.ID()).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().User(*uId).Model(m.ID()).Project(s.Project()).Schema(s.ID()).Thread(id.NewThreadID()).MustBuild()
	i3 := item.New().NewID().User(id.NewUserID()).Model(m.ID()).Project(s.Project()).Schema(s.ID()).Thread(id.NewThreadID()).MustBuild()

	ctx := context.Background()
	db := memory.New()
	lo.Must0(db.Project.Save(ctx, prj))
	lo.Must0(db.Schema.Save(ctx, s))
	lo.Must0(db.Model.Save(ctx, m))
	lo.Must0(db.Item.Save(ctx, i))
	lo.Must0(db.Item.Save(ctx, i2))
	lo.Must0(db.Item.Save(ctx, i3))
	itemUC := NewItem(db, nil)
	itemUC.ignoreEvent = true

	op := &usecase.Operator{
		User:             uId,
		ReadableProjects: []id.ProjectID{s.Project()},
		WritableProjects: []id.ProjectID{s.Project()},
	}

	// ok
	item, err := itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "xxx",
			},
		},
	}, op)
	assert.NoError(t, err)
	assert.Equal(t, i.ID(), item.Value().ID())
	assert.Equal(t, s.ID(), item.Value().Schema())

	it, err := db.Item.FindByID(ctx, item.Value().ID(), nil)
	assert.NoError(t, err)
	assert.Equal(t, item.Value(), it.Value())
	assert.Equal(t, value.TypeText.Value("xxx").AsMultiple(), it.Value().Field(sf.ID()).Value())

	// ok with key
	item, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Key:   sf.Key().Ref(),
				Type:  value.TypeText,
				Value: "yyy",
			},
		},
	}, op)
	assert.NoError(t, err)
	assert.Equal(t, i.ID(), item.Value().ID())
	assert.Equal(t, s.ID(), item.Value().Schema())

	it, err = db.Item.FindByID(ctx, item.Value().ID(), nil)
	assert.NoError(t, err)
	assert.Equal(t, item.Value(), it.Value())
	assert.Equal(t, value.TypeText.Value("yyy").AsMultiple(), it.Value().Field(sf.ID()).Value())

	// validate fails
	item, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "abcabcabcabc", // too long
			},
		},
	}, op)
	assert.ErrorContains(t, err, "it sholud be shorter than 10")
	assert.Nil(t, item)

	// update same item is not a duplicate
	item, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "xxx", // duplicated
			},
		},
	}, op)
	assert.NoError(t, err)
	assert.Equal(t, i.ID(), item.Value().ID())
	assert.Equal(t, s.ID(), item.Value().Schema())

	// update no permission
	_, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i3.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "xxx",
			},
		},
	}, op)
	assert.Equal(t, interfaces.ErrOperationDenied, err)

	// duplicate
	item, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i2.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "xxx", // duplicated
			},
		},
	}, op)
	assert.Equal(t, interfaces.ErrDuplicatedItemValue, err)
	assert.Nil(t, item)

	// no fields
	item, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i.ID(),
		Fields: []interfaces.ItemFieldParam{},
	}, op)
	assert.Equal(t, interfaces.ErrItemFieldRequired, err)
	assert.Nil(t, item)

	// required
	sf.SetRequired(true)
	s.RemoveField(sf.ID())
	s.AddField(sf)
	lo.Must0(db.Schema.Save(ctx, s))
	item, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "",
			},
		},
	}, op)
	assert.ErrorIs(t, err, schema.ErrValueRequired)
	assert.Nil(t, item)

	// mock item error
	wantErr := errors.New("test")
	memory.SetItemError(db.Item, wantErr)
	item, err = itemUC.Update(ctx, interfaces.UpdateItemParam{
		ItemID: i.ID(),
		Fields: []interfaces.ItemFieldParam{
			{
				Field: sf.ID().Ref(),
				Type:  value.TypeText,
				Value: "a",
			},
		},
	}, op)
	assert.Equal(t, wantErr, err)
	assert.Nil(t, item)
}

func TestItem_Delete(t *testing.T) {
	wid := id.NewWorkspaceID()
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	sid := id.NewSchemaID()
	id1 := id.NewItemID()
	i1 := item.New().ID(id1).User(u.ID()).Schema(sid).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()

	op := &usecase.Operator{
		User:             lo.ToPtr(u.ID()),
		WritableProjects: id.ProjectIDList{i1.Project()},
	}
	ctx := context.Background()

	db := memory.New()
	err := db.Item.Save(ctx, i1)
	assert.NoError(t, err)

	itemUC := NewItem(db, nil)
	itemUC.ignoreEvent = true
	err = itemUC.Delete(ctx, id1, op)
	assert.NoError(t, err)

	_, err = itemUC.FindByID(ctx, id1, op)
	assert.Error(t, err)

	// mock item error
	wantErr := rerror.ErrNotFound
	err = itemUC.Delete(ctx, id.NewItemID(), op)
	assert.Equal(t, wantErr, err)
}

func TestWorkFlow(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()

	wid := id.NewWorkspaceID()
	prj := project.New().NewID().Workspace(wid).MustBuild()
	s := schema.New().NewID().Workspace(id.NewWorkspaceID()).Project(prj.ID()).MustBuild()
	m := model.New().NewID().Project(prj.ID()).Schema(s.ID()).RandomKey().MustBuild()
	i := item.New().NewID().Schema(s.ID()).Model(m.ID()).Project(prj.ID()).Thread(id.NewThreadID()).MustBuild()
	ri, _ := request.NewItem(i.ID())
	u := user.New().Name("aaa").NewID().Email("aaa@bbb.com").Workspace(wid).MustBuild()
	req1 := request.New().
		NewID().
		Workspace(wid).
		Project(prj.ID()).
		Reviewers(id.UserIDList{u.ID()}).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(request.ItemList{ri}).
		Title("foo").
		MustBuild()
	op := &usecase.Operator{
		User:             lo.ToPtr(u.ID()),
		OwningWorkspaces: id.WorkspaceIDList{wid},
	}
	ctx := context.Background()

	db := memory.New()
	err := db.Project.Save(ctx, prj)
	assert.NoError(t, err)
	err = db.Schema.Save(ctx, s)
	assert.NoError(t, err)
	err = db.Model.Save(ctx, m)
	assert.NoError(t, err)
	err = db.Item.Save(ctx, i)
	assert.NoError(t, err)

	itemUC := NewItem(db, nil)

	status, err := itemUC.ItemStatus(ctx, id.ItemIDList{i.ID()}, op)
	assert.NoError(t, err)
	assert.Equal(t, map[id.ItemID]item.Status{i.ID(): item.StatusDraft}, status)

	err = db.Request.Save(ctx, req1)
	assert.NoError(t, err)

	status, err = itemUC.ItemStatus(ctx, id.ItemIDList{i.ID()}, op)
	assert.NoError(t, err)
	assert.Equal(t, map[id.ItemID]item.Status{i.ID(): item.StatusReview}, status)

	requestUC := NewRequest(db, nil)
	_, err = requestUC.Approve(ctx, req1.ID(), op)
	assert.NoError(t, err)

	status, err = itemUC.ItemStatus(ctx, id.ItemIDList{i.ID()}, op)
	assert.NoError(t, err)
	assert.Equal(t, map[id.ItemID]item.Status{i.ID(): item.StatusPublic}, status)

	_, err = itemUC.Unpublish(ctx, id.ItemIDList{i.ID()}, op)
	assert.NoError(t, err)

	status, err = itemUC.ItemStatus(ctx, id.ItemIDList{i.ID()}, op)
	assert.NoError(t, err)
	assert.Equal(t, map[id.ItemID]item.Status{i.ID(): item.StatusDraft}, status)
}
