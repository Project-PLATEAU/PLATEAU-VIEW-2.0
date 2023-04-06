package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/item"
	"github.com/reearth/reearth-cms/server/pkg/value"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/stretchr/testify/assert"
)

func TestItem_FindByID(t *testing.T) {
	ctx := context.Background()
	i := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()
	r := NewItem()
	_ = r.Save(ctx, i)

	out, err := r.FindByID(ctx, i.ID(), nil)
	assert.NoError(t, err)
	assert.Equal(t, i, out.Value())

	out2, err := r.FindByID(ctx, id.ItemID{}, nil)
	assert.Nil(t, out2)
	assert.Same(t, rerror.ErrNotFound, err)
}

func TestItem_Remove(t *testing.T) {
	ctx := context.Background()
	pid, pid2 := id.NewProjectID(), id.NewProjectID()
	i1 := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i3 := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(pid2).Thread(id.NewThreadID()).MustBuild()

	r := NewItem()
	_ = r.Save(ctx, i1)
	_ = r.Save(ctx, i2)
	_ = r.Save(ctx, i3)
	r = r.Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	})

	err := r.Remove(ctx, i1.ID())
	assert.NoError(t, err)
	data, _ := r.FindByIDs(ctx, id.ItemIDList{i1.ID(), i2.ID()}, nil)
	assert.Equal(t, item.List{i2}, data.Unwrap())

	err = r.Remove(ctx, i1.ID())
	assert.Equal(t, rerror.ErrNotFound, err)

	err = r.Remove(ctx, i3.ID())
	assert.Equal(t, repo.ErrOperationDenied, err)

	wantErr := errors.New("test")
	SetItemError(r, wantErr)
	assert.Same(t, wantErr, r.Remove(ctx, i1.ID()))
}

func TestItem_Save(t *testing.T) {
	ctx := context.Background()
	i := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(id.NewProjectID()).MustBuild()
	i2 := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(id.NewProjectID()).MustBuild()
	pf := repo.ProjectFilter{
		Readable: []id.ProjectID{i.Project()},
		Writable: []id.ProjectID{i.Project()},
	}
	r := NewItem().Filtered(pf)

	_ = r.Save(ctx, i)
	got, _ := r.FindByID(ctx, i.ID(), nil)
	assert.Equal(t, i, got.Value())

	err := r.Save(ctx, i2)
	assert.Equal(t, repo.ErrOperationDenied, err)

	wantErr := errors.New("test")
	SetItemError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, i))
}

func TestItem_FindByIDs(t *testing.T) {
	ctx := context.Background()
	i := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(id.NewProjectID()).MustBuild()
	i2 := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Thread(id.NewThreadID()).Project(id.NewProjectID()).MustBuild()
	r := NewItem()
	_ = r.Save(ctx, i)
	_ = r.Save(ctx, i2)

	ids := id.ItemIDList{i.ID()}
	il := item.List{i}
	out, err := r.FindByIDs(ctx, ids, nil)
	assert.NoError(t, err)
	assert.Equal(t, il, out.Unwrap())
}

func TestItem_FindAllVersionsByID(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()
	ctx := context.Background()
	i := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()
	r := NewItem()
	_ = r.Save(ctx, i)

	v, err := r.FindAllVersionsByID(ctx, i.ID())
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.MustBeValue(v[0].Version(), nil, version.NewRefs(version.Latest), now, i),
	}, v)

	r = r.Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{},
		Writable: []id.ProjectID{},
	})
	res, err := r.FindAllVersionsByID(ctx, i.ID())
	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestItem_FindAllVersionsByIDs(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()
	ctx := context.Background()
	iid1, iid2 := item.NewID(), item.NewID()
	i1 := item.New().ID(iid1).Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().ID(iid2).Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()
	r := NewItem()
	_ = r.Save(ctx, i1)

	v, err := r.FindAllVersionsByIDs(ctx, id.ItemIDList{i1.ID()})
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.MustBeValue(v[0].Version(), nil, version.NewRefs(version.Latest), now, i1),
	}, v)

	_ = r.Save(ctx, i2)
	v, err = r.FindAllVersionsByIDs(ctx, id.ItemIDList{i1.ID(), i2.ID()})
	assert.NoError(t, err)
	assert.Equal(t, item.VersionedList{
		version.MustBeValue(v[0].Version(), nil, version.NewRefs(version.Latest), now, i1),
		version.MustBeValue(v[1].Version(), nil, version.NewRefs(version.Latest), now, i2),
	}, v)

	r = r.Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{},
		Writable: []id.ProjectID{},
	})
	res, err := r.FindAllVersionsByIDs(ctx, id.ItemIDList{i1.ID(), i2.ID()})
	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestItem_FindBySchema(t *testing.T) {
	ctx := context.Background()
	sid1, sid2 := id.NewSchemaID(), id.NewSchemaID()
	pid1, pid2 := id.NewProjectID(), id.NewProjectID()
	i1 := item.New().NewID().Schema(sid1).Project(pid1).Model(id.NewModelID()).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Schema(sid2).Project(pid2).Model(id.NewModelID()).Thread(id.NewThreadID()).MustBuild()

	r := NewItem().Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{pid1},
		Writable: []id.ProjectID{pid1},
	})
	_ = r.Save(ctx, i1)
	_ = r.Save(ctx, i2)

	got, _, err := r.FindBySchema(ctx, sid1, nil, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, item.List{i1}, got.Unwrap())

	got, _, err = r.FindBySchema(ctx, sid2, nil, nil, nil)
	assert.NoError(t, err)
	assert.Nil(t, got)
}

func TestItem_FindByProject(t *testing.T) {
	ctx := context.Background()
	pid := id.NewProjectID()
	i1 := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(pid).Thread(id.NewThreadID()).MustBuild()
	r := NewItem().Filtered(repo.ProjectFilter{
		Readable: []id.ProjectID{pid},
		Writable: []id.ProjectID{pid},
	})
	_ = r.Save(ctx, i1)
	_ = r.Save(ctx, i2)

	got, _, err := r.FindByProject(ctx, pid, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(got.Unwrap()))
}

func TestItem_FindByFieldValue(t *testing.T) {
	ctx := context.Background()
	sid := id.NewSchemaID()
	sf1 := id.NewFieldID()
	sf2 := id.NewFieldID()
	pid := id.NewProjectID()
	f1 := item.NewField(sf1, value.TypeText.Value("foo").AsMultiple())
	f2 := item.NewField(sf2, value.TypeText.Value("hoge").AsMultiple())
	i := item.New().NewID().Schema(sid).Model(id.NewModelID()).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i3 := item.New().NewID().Schema(sid).Model(id.NewModelID()).Fields([]*item.Field{f2}).Project(pid).Thread(id.NewThreadID()).MustBuild()

	r := NewItem()
	_ = r.Save(ctx, i)
	_ = r.Save(ctx, i2)
	_ = r.Save(ctx, i3)
	q := item.NewQuery(pid, sid.Ref(), "foo", nil)
	got, _, _ := r.Search(ctx, q, nil, nil)
	assert.Equal(t, 2, len(got))

	wantErr := errors.New("test")
	SetItemError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, i))
}

func TestItem_FindByModelAndValue(t *testing.T) {
	ctx := context.Background()
	sid := id.NewSchemaID()
	sf1 := id.NewFieldID()
	sf2 := id.NewFieldID()
	pid := id.NewProjectID()
	f1 := item.NewField(sf1, value.TypeText.Value("foo").AsMultiple())
	f2 := item.NewField(sf2, value.TypeText.Value("hoge").AsMultiple())
	mid := id.NewModelID()
	i := item.New().NewID().Schema(sid).Model(mid).Fields([]*item.Field{f1}).Project(pid).Thread(id.NewThreadID()).MustBuild()
	i2 := item.New().NewID().Schema(sid).Model(mid).Fields([]*item.Field{f2}).Project(pid).Thread(id.NewThreadID()).MustBuild()

	r := NewItem()
	_ = r.Save(ctx, i)
	_ = r.Save(ctx, i2)
	got, _ := r.FindByModelAndValue(ctx, mid, []repo.FieldAndValue{{
		Field: f1.FieldID(),
		Value: f1.Value(),
	}}, nil)
	assert.Equal(t, 1, len(got))

	wantErr := errors.New("test")
	SetItemError(r, wantErr)
	assert.Same(t, wantErr, r.Save(ctx, i))
}

func TestItem_UpdateRef(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()

	vx := version.Ref("xxx")
	ctx := context.Background()
	i := item.New().NewID().Schema(id.NewSchemaID()).Model(id.NewModelID()).Project(id.NewProjectID()).Thread(id.NewThreadID()).MustBuild()
	r := NewItem()
	_ = r.Save(ctx, i)
	v, _ := r.FindByID(ctx, i.ID(), nil)
	_ = r.UpdateRef(ctx, i.ID(), vx, v.Version().OrRef().Ref())
	v2, _ := r.FindByID(ctx, i.ID(), nil)
	assert.Equal(t, version.MustBeValue(v.Version(), nil, version.NewRefs(vx, version.Latest), now, i), v2)

	wantErr := errors.New("test")
	SetItemError(r, wantErr)
	assert.Same(t, wantErr, r.UpdateRef(ctx, i.ID(), vx, v.Version().OrRef().Ref()))
}
