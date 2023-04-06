package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/schema"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestItem_FindByID(t *testing.T) {
	id1 := id.NewItemID()
	sid := id.NewSchemaID()
	pid := id.NewProjectID()
	sfid := schema.NewFieldID()
	fs := []*item.Field{item.NewField(sfid, value.TypeBool.Value(true).AsMultiple())}
	i1 := item.New().ID(id1).Fields(fs).Schema(sid).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(pid).MustBuild()
	tests := []struct {
		Name               string
		Input              id.ItemID
		RepoData, Expected *item.Item
		WantErr            bool
	}{
		{
			Name:     "must find a item",
			Input:    i1.ID(),
			RepoData: i1,
			Expected: i1,
		},
		{
			Name:     "must not find any item",
			Input:    id.NewItemID(),
			RepoData: i1,
			WantErr:  true,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))
			repo := NewItem(client)

			ctx := context.Background()
			err := repo.Save(ctx, tc.RepoData)
			assert.NoError(tt, err)

			got, err := repo.FindByID(ctx, tc.Input, nil)
			if tc.WantErr {
				assert.Equal(tt, err, rerror.ErrNotFound)
			} else {
				assert.Equal(tt, tc.Expected.ID(), got.Value().ID())
			}
		})
	}
}

func TestItem_FindAllVersionsByID(t *testing.T) {
	now1 := time.Now().Truncate(time.Millisecond).UTC()
	now2 := now1.Add(time.Second)
	nowid1 := primitive.NewObjectIDFromTimestamp(time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC)).Timestamp()
	nowid2 := primitive.NewObjectIDFromTimestamp(time.Date(2022, time.April, 2, 0, 0, 0, 0, time.UTC)).Timestamp()

	iid := id.NewItemID()
	sfid := schema.NewFieldID()
	pid := id.NewProjectID()
	fs := []*item.Field{item.NewField(sfid, value.TypeBool.Value(true).AsMultiple())}
	i1 := item.New().ID(iid).Fields(fs).Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).Timestamp(now1).MustBuild()
	i2 := item.New().ID(iid).Fields(fs).Schema(i1.Schema()).Model(id.NewModelID()).Project(i1.Project()).Thread(id.NewThreadID()).Timestamp(now2).MustBuild()

	init := mongotest.Connect(t)
	client := mongox.NewClientWithDatabase(init(t))

	r := NewItem(client)
	ctx := context.Background()
	defer util.MockNow(nowid1)()
	assert.NoError(t, r.Save(ctx, i1))

	got1, err := r.FindAllVersionsByID(ctx, iid)
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.NewValue(got1[0].Version(), nil, version.NewRefs(version.Latest), nowid1, i1),
	}, got1)

	defer util.MockNow(nowid2)()
	assert.NoError(t, r.Save(ctx, i2))

	got2, err := r.FindAllVersionsByID(ctx, iid)
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.NewValue(got2[0].Version(), nil, nil, nowid1, i1),
		version.NewValue(got2[1].Version(), version.NewVersions(got2[0].Version()), version.NewRefs(version.Latest), nowid2, i2),
	}, got2)

	r = r.Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{id.NewProjectID()},
	})
	got3, err := r.FindAllVersionsByID(ctx, iid)
	assert.Nil(t, got3)
	assert.NoError(t, err)
}

func TestItem_FindAllVersionsByIDs(t *testing.T) {
	now1 := time.Now().Truncate(time.Millisecond).UTC()
	now2 := now1.Add(time.Second)
	nowid1 := primitive.NewObjectIDFromTimestamp(time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC)).Timestamp()
	nowid2 := primitive.NewObjectIDFromTimestamp(time.Date(2022, time.April, 2, 0, 0, 0, 0, time.UTC)).Timestamp()

	iid1 := id.NewItemID()
	iid2 := id.NewItemID()
	sfid := schema.NewFieldID()
	pid := id.NewProjectID()
	fs := []*item.Field{item.NewField(sfid, value.TypeBool.Value(true).AsMultiple())}
	i1 := item.New().ID(iid1).Fields(fs).Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).Timestamp(now1).MustBuild()
	i2 := item.New().ID(iid2).Fields(fs).Schema(i1.Schema()).Model(id.NewModelID()).Project(i1.Project()).Thread(id.NewThreadID()).Timestamp(now2).MustBuild()

	init := mongotest.Connect(t)
	client := mongox.NewClientWithDatabase(init(t))

	r := NewItem(client)
	ctx := context.Background()
	defer util.MockNow(nowid1)()
	assert.NoError(t, r.Save(ctx, i1))

	got1, err := r.FindAllVersionsByIDs(ctx, id.ItemIDList{iid1})
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.NewValue(got1[0].Version(), nil, version.NewRefs(version.Latest), nowid1, i1),
	}, got1)

	defer util.MockNow(nowid2)()
	assert.NoError(t, r.Save(ctx, i2))

	got2, err := r.FindAllVersionsByIDs(ctx, id.ItemIDList{iid1, iid2})
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.NewValue(got2[0].Version(), nil, version.NewRefs(version.Latest), nowid1, i1),
		version.NewValue(got2[1].Version(), nil, version.NewRefs(version.Latest), nowid2, i2),
	}, got2)

	r = r.Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{id.NewProjectID()},
	})
	got3, err := r.FindAllVersionsByIDs(ctx, id.ItemIDList{iid1, iid2})
	assert.Nil(t, got3)
	assert.NoError(t, err)
}

func TestItem_FindByIDs(t *testing.T) {
	defer util.MockNow(time.Now().Truncate(time.Millisecond).UTC())()

	sid := id.NewSchemaID()
	pid := id.NewProjectID()
	sfid := schema.NewFieldID()
	fs := []*item.Field{item.NewField(sfid, value.TypeBool.Value(true).AsMultiple())}
	i1 := item.New().NewID().Fields(fs).Schema(sid).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Fields(fs).Schema(sid).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).MustBuild()
	tests := []struct {
		Name               string
		Input              id.ItemIDList
		RepoData, Expected item.List
	}{
		{
			Name:     "must find two items",
			Input:    id.ItemIDList{i1.ID(), i2.ID()},
			RepoData: item.List{i1, i2},
			Expected: item.List{i1, i2},
		},
		{
			Name:     "must not find any item",
			Input:    id.ItemIDList{id.NewItemID()},
			RepoData: item.List{i1, i2},
			Expected: item.List{},
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewItem(client)
			ctx := context.Background()
			// @TODO create SaveAll func
			for _, i := range tc.RepoData {
				err := repo.Save(ctx, i)
				assert.NoError(tt, err)
			}

			got, _ := repo.FindByIDs(ctx, tc.Input, nil)
			assert.Equal(tt, tc.Expected, got.Unwrap())
		})
	}
}

func TestItem_FindBySchema(t *testing.T) {
	defer util.MockNow(time.Now().Truncate(time.Millisecond).UTC())()

	sid := id.NewSchemaID()
	pid := id.NewProjectID()
	sfid := schema.NewFieldID()
	fs := []*item.Field{item.NewField(sfid, value.TypeBool.Value(true).AsMultiple())}
	i1 := item.New().NewID().Fields(fs).Schema(sid).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(pid).MustBuild()
	i2 := item.New().NewID().Fields(fs).Schema(sid).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(pid).MustBuild()
	i3 := item.New().NewID().Fields(fs).Schema(sid).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(id.NewProjectID()).MustBuild()
	i4 := item.New().NewID().Fields(fs).Schema(id.NewSchemaID()).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(pid).MustBuild()

	tests := []struct {
		Name        string
		Input       id.SchemaID
		Seeds       item.List
		Expected    item.List
		ExpectedErr error
	}{
		{
			Name:     "must find two items (first 10)",
			Input:    sid,
			Seeds:    item.List{i1, i2, i3},
			Expected: item.List{i1, i2},
		},
		{
			Name:  "must not find any item",
			Input: sid,
			Seeds: item.List{i4},
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			client := mongox.NewClientWithDatabase(init(t))
			r := NewItem(client)
			ctx := context.Background()

			for _, i := range tc.Seeds {
				err := r.Save(ctx, i)
				assert.NoError(tt, err)
			}
			r = r.Filtered(repo.ProjectFilter{
				Readable: []id.ProjectID{pid},
				Writable: []id.ProjectID{pid},
			})

			got, _, err := r.FindBySchema(ctx, tc.Input, nil, nil, usecasex.CursorPagination{First: lo.ToPtr(int64(10))}.Wrap())
			assert.Equal(tt, tc.Expected, got.Unwrap())
			assert.Equal(tt, tc.ExpectedErr, err)
		})
	}
}

func TestItem_FindByProject(t *testing.T) {
	defer util.MockNow(time.Now().Truncate(time.Millisecond).UTC())()
	pid := id.NewProjectID()
	sid := id.NewSchemaID()
	i1 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(pid).MustBuild()
	i2 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(pid).MustBuild()
	tests := []struct {
		Name     string
		Input    id.ProjectID
		RepoData item.List
		Expected item.List
	}{
		{
			Name:     "must find two items (first 10)",
			Input:    pid,
			RepoData: item.List{i1, i2},
			Expected: item.List{i1, i2},
		},
		{
			Name:     "must not find any item",
			Input:    id.NewProjectID(),
			RepoData: item.List{i1, i2},
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewItem(client).Filtered(repo.ProjectFilter{
				Readable: []id.ProjectID{pid},
				Writable: []id.ProjectID{pid},
			})
			ctx := context.Background()
			for _, i := range tc.RepoData {
				err := repo.Save(ctx, i)
				assert.NoError(tt, err)
			}

			got, _, _ := repo.FindByProject(ctx, tc.Input, nil, usecasex.CursorPagination{First: lo.ToPtr(int64(10))}.Wrap())
			assert.Equal(tt, tc.Expected, got.Unwrap())
		})
	}
}

func TestItem_Remove(t *testing.T) {
	id1 := id.NewItemID()
	sid := id.NewSchemaID()
	pid := id.NewProjectID()
	sfid := schema.NewFieldID()
	fs := []*item.Field{item.NewField(sfid, value.TypeBool.Value(true).AsMultiple())}
	i1 := item.New().ID(id1).Fields(fs).Schema(sid).Model(id.NewModelID()).Project(pid).Project(pid).Thread(id.NewThreadID()).MustBuild()
	init := mongotest.Connect(t)
	client := mongox.NewClientWithDatabase(init(t))

	r := NewItem(client).Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	})
	ctx := context.Background()
	_ = r.Save(ctx, i1)

	err := r.Remove(ctx, i1.ID())
	assert.NoError(t, err)

	got, err := r.FindByID(ctx, i1.ID(), nil)
	assert.Nil(t, got)
	assert.Equal(t, rerror.ErrNotFound, err)
}

func TestItem_Archive(t *testing.T) {
	iid := id.NewItemID()
	pid := id.NewProjectID()
	pid2 := id.NewProjectID()
	init := mongotest.Connect(t)
	client := mongox.NewClientWithDatabase(init(t))
	ctx := context.Background()

	r := NewItem(client).Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	})

	res, err := r.IsArchived(ctx, iid)
	assert.NoError(t, err)
	assert.False(t, res)

	// failed to archive
	err = r.Archive(ctx, iid, pid2, true)
	assert.Same(t, repo.ErrOperationDenied, err)

	res, err = r.IsArchived(ctx, iid)
	assert.NoError(t, err)
	assert.False(t, res)

	// successfully archive
	err = r.Archive(ctx, iid, pid, true)
	assert.NoError(t, err)

	var d bson.M
	err = client.Database().Collection("item").FindOne(ctx, bson.M{
		"__": true,
		"id": iid.String(),
	}).Decode(&d)
	assert.NoError(t, err)
	assert.Equal(t, bson.M{
		"_id":      d["_id"],
		"__":       true,
		"id":       iid.String(),
		"project":  pid.String(),
		"archived": true,
	}, d)

	res, err = r.IsArchived(ctx, iid)
	assert.NoError(t, err)
	assert.True(t, res)

	// failed to unarchive
	err = r.Archive(ctx, iid, pid2, false)
	assert.Same(t, repo.ErrOperationDenied, err)

	res, err = r.IsArchived(ctx, iid)
	assert.NoError(t, err)
	assert.True(t, res)

	// successfully unarchived
	err = r.Archive(ctx, iid, pid, false)
	assert.NoError(t, err)

	res, err = r.IsArchived(ctx, iid)
	assert.NoError(t, err)
	assert.False(t, res)
}

func TestItem_Search(t *testing.T) {
	sid := id.NewSchemaID()
	sid2 := id.NewSchemaID()
	sf1 := id.NewFieldID()
	sf2 := id.NewFieldID()
	f1 := item.NewField(sf1, value.TypeText.Value("foo").AsMultiple())
	f2 := item.NewField(sf2, value.TypeInteger.Value(2).AsMultiple())
	pid := id.NewProjectID()
	i1 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i3 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Fields([]*item.Field{f2}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i4 := item.New().NewID().Schema(sid2).Model(id.NewModelID()).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	tests := []struct {
		Name     string
		Input    *item.Query
		RepoData item.List
		Expected int
	}{
		{
			Name:     "must find two items (first 10)",
			Input:    item.NewQuery(pid, nil, "foo", nil),
			RepoData: item.List{i1, i2, i3},
			Expected: 2,
		},
		{
			Name:     "must find all items",
			Input:    item.NewQuery(pid, nil, "", nil),
			RepoData: item.List{i1, i2, i3},
			Expected: 3,
		},
		{
			Name:     "must find one item",
			Input:    item.NewQuery(pid, sid2.Ref(), "foo", nil),
			RepoData: item.List{i1, i2, i3, i4},
			Expected: 1,
		},
	}

	init := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			// tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewItem(client)
			ctx := context.Background()
			for _, i := range tc.RepoData {
				err := repo.Save(ctx, i)
				assert.NoError(t, err)
			}

			got, _, _ := repo.Search(ctx, tc.Input, nil, usecasex.CursorPagination{First: lo.ToPtr(int64(10))}.Wrap())
			assert.Equal(t, tc.Expected, len(got))
		})
	}
}

func TestItem_FindByModelAndValue(t *testing.T) {
	init := mongotest.Connect(t)
	sid := id.NewSchemaID()
	sf1 := id.NewFieldID()
	sf2 := id.NewFieldID()
	f1 := item.NewField(sf1, value.TypeText.Value("foo").AsMultiple())
	f2 := item.NewField(sf2, value.TypeText.Value("hoge").AsMultiple())
	pid := id.NewProjectID()
	mid := id.NewModelID()
	i1 := item.New().NewID().Schema(sid).Model(mid).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Fields([]*item.Field{f2}).Project(pid).Thread(id.NewThreadID()).MustBuild()

	type args struct {
		model  id.ModelID
		fields []repo.FieldAndValue
	}
	tests := []struct {
		Name     string
		Input    args
		Seeds    item.List
		Expected int
		WantErr  error
	}{
		{
			Name: "must not find any item",
			Input: args{
				model: mid,
				fields: []repo.FieldAndValue{
					{
						Field: f2.FieldID(),
						Value: f2.Value(),
					},
				},
			},
			Seeds:    item.List{i1, i2},
			Expected: 0,
		},
		{
			Name: "must find one item",
			Input: args{
				model: mid,
				fields: []repo.FieldAndValue{
					{
						Field: f1.FieldID(),
						Value: f1.Value(),
					},
				},
			},
			Seeds:    item.List{i1, i2},
			Expected: 1,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewItem(client)
			ctx := context.Background()
			for _, i := range tc.Seeds {
				err := repo.Save(ctx, i)
				assert.NoError(tt, err)
			}

			got, err := repo.FindByModelAndValue(ctx, tc.Input.model, tc.Input.fields, nil)
			assert.Equal(tt, tc.WantErr, err)
			assert.Equal(tt, tc.Expected, len(got))
		})
	}
}

func TestItem_UpdateRef(t *testing.T) {
	vx := version.Ref("xxx")
	ctx := context.Background()
	i := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()
	init := mongotest.Connect(t)
	client := mongox.NewClientWithDatabase(init(t))
	r := NewItem(client)
	_ = r.Save(ctx, i)
	v, _ := r.FindByID(ctx, i.ID(), nil)
	err := r.UpdateRef(ctx, i.ID(), vx, v.Version().OrRef().Ref())
	assert.NoError(t, err)
	v2, _ := r.FindByID(ctx, i.ID(), nil)
	assert.Equal(t, version.NewRefs(vx, version.Latest), v2.Refs())
}

func TestItem_FindByAssets(t *testing.T) {
	init := mongotest.Connect(t)
	sid := id.NewSchemaID()
	aid1 := id.NewAssetID()
	aid2 := id.NewAssetID()
	sf1 := id.NewFieldID()
	f1 := item.NewField(sf1, value.TypeAsset.Value(aid1.String()).AsMultiple())
	pid := id.NewProjectID()
	mid := id.NewModelID()
	i1 := item.New().NewID().Schema(sid).Model(mid).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()

	tests := []struct {
		Name     string
		Input    id.AssetIDList
		Seeds    item.List
		Expected int
		WantErr  error
	}{
		{
			Name:     "must find 1 item",
			Input:    id.AssetIDList{aid1, aid2},
			Seeds:    item.List{i1},
			Expected: 1,
		},
		{
			Name:     "must not find any item",
			Input:    id.AssetIDList{},
			Seeds:    item.List{i1},
			Expected: 0,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			client := mongox.NewClientWithDatabase(init(t))

			repo := NewItem(client)
			ctx := context.Background()
			for _, i := range tc.Seeds {
				err := repo.Save(ctx, i)
				assert.NoError(tt, err)
			}

			got, err := repo.FindByAssets(ctx, tc.Input, nil)
			assert.Equal(tt, tc.WantErr, err)
			assert.Equal(tt, tc.Expected, len(got))
		})
	}
}
