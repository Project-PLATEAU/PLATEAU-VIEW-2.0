package schema

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/stretchr/testify/assert"
)

func TestSchema_AddField(t *testing.T) {
	fid := NewFieldID()
	tests := []struct {
		name string
		s    *Schema
		f    *Field
		want *Schema
	}{
		{
			name: "add on empty array",
			s:    &Schema{},
			f:    &Field{name: "f1"},
			want: &Schema{fields: []*Field{{name: "f1", order: 0}}},
		},
		{
			name: "add on not empty array",
			s:    &Schema{fields: []*Field{{id: fid, name: "f1", order: 1}}},
			f:    &Field{name: "f2"},
			want: &Schema{fields: []*Field{{id: fid, name: "f1", order: 1}, {name: "f2", order: 2}}},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid, name: "f1"}}},
			f:    &Field{id: fid, name: "f2"},
			want: &Schema{fields: []*Field{{id: fid, name: "f1"}}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.s.AddField(tc.f)
			assert.Equal(t, tc.want, tc.s)
		})
	}
}

func TestSchema_HasField(t *testing.T) {
	fid1 := NewFieldID()
	fid2 := NewFieldID()
	fid3 := NewFieldID()
	tests := []struct {
		name string
		s    *Schema
		fid  FieldID
		want bool
	}{
		{
			name: "add on empty array",
			s:    &Schema{},
			fid:  fid1,
			want: false,
		},
		{
			name: "add on not empty array",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}}},
			fid:  fid1,
			want: true,
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  fid1,
			want: true,
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  NewFieldID(),
			want: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.s.HasField(tc.fid))
		})
	}
}

func TestSchema_RemoveField(t *testing.T) {
	fid1 := NewFieldID()
	fid2 := NewFieldID()
	fid3 := NewFieldID()
	tests := []struct {
		name string
		s    *Schema
		fid  FieldID
		want *Schema
	}{
		{
			name: "add on empty array",
			s:    &Schema{},
			fid:  fid1,
			want: &Schema{},
		},
		{
			name: "add on not empty array",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}}},
			fid:  fid1,
			want: &Schema{fields: []*Field{}},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  fid1,
			want: &Schema{fields: []*Field{{id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  fid2,
			want: &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid3, name: "f3"}}},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  fid3,
			want: &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.s.RemoveField(tc.fid)
			assert.Equal(t, tc.want, tc.s)
		})
	}
}

func TestSchema_Field(t *testing.T) {
	fid1 := NewFieldID()
	fid2 := NewFieldID()
	fid3 := NewFieldID()
	tests := []struct {
		name string
		s    *Schema
		fid  FieldID
		want *Field
	}{
		{
			name: "add on empty array",
			s:    &Schema{},
			fid:  fid1,
			want: nil,
		},
		{
			name: "add on not empty array",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}}},
			fid:  fid1,
			want: &Field{id: fid1, name: "f1"},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  fid1,
			want: &Field{id: fid1, name: "f1"},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  fid2,
			want: &Field{id: fid2, name: "f2"},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fid:  fid3,
			want: &Field{id: fid3, name: "f3"},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.s.Field(tc.fid))
		})
	}
}

func TestSchema_FieldByIDOrKey(t *testing.T) {
	f1 := &Field{id: NewFieldID(), name: "f1"}
	f2 := &Field{id: NewFieldID(), name: "f2"}
	f3 := &Field{id: NewFieldID(), name: "f3", key: key.New("KEY")}
	f4 := &Field{id: NewFieldID(), name: "f4", key: key.New("id")}
	s := &Schema{fields: []*Field{f1, f2, f3, f4}}

	assert.Equal(t, f1, s.FieldByIDOrKey(f1.ID().Ref(), nil))
	assert.Equal(t, f2, s.FieldByIDOrKey(f2.ID().Ref(), nil))
	assert.Equal(t, f3, s.FieldByIDOrKey(f3.ID().Ref(), nil))
	assert.Equal(t, f4, s.FieldByIDOrKey(f4.ID().Ref(), nil))
	assert.Equal(t, f3, s.FieldByIDOrKey(nil, f3.Key().Ref()))
	assert.Equal(t, f1, s.FieldByIDOrKey(f1.ID().Ref(), f3.Key().Ref()))
	assert.Nil(t, s.FieldByIDOrKey(id.NewFieldID().Ref(), nil))
	assert.Nil(t, s.FieldByIDOrKey(nil, key.New("").Ref()))
	assert.Nil(t, s.FieldByIDOrKey(nil, key.New("x").Ref()))
	assert.Nil(t, s.FieldByIDOrKey(nil, key.New("id").Ref()))
}

func TestSchema_Fields(t *testing.T) {
	fid1 := NewFieldID()
	fid2 := NewFieldID()
	fid3 := NewFieldID()
	tests := []struct {
		name string
		s    *Schema
		want FieldList
	}{
		{
			name: "add on not empty array",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}}},
			want: []*Field{{id: fid1, name: "f1"}},
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			want: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.s.Fields())
		})
	}
}

func TestSchema_ID(t *testing.T) {
	sid := NewID()
	tests := []struct {
		name string
		s    Schema
		want ID
	}{
		{
			name: "id",
			want: ID{},
		},
		{
			name: "id",
			s: Schema{
				id: sid,
			},
			want: sid,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.s.ID())
		})
	}
}

func TestSchema_SetWorkspace(t *testing.T) {
	wid := id.NewWorkspaceID()
	tests := []struct {
		name string
		wid  id.WorkspaceID
		want *Schema
	}{
		{
			name: "id",
			wid:  wid,
			want: &Schema{workspace: wid},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Schema{}
			s.SetWorkspace(tt.wid)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSchema_Workspace(t *testing.T) {
	wId := id.NewWorkspaceID()
	tests := []struct {
		name string
		s    Schema
		want id.WorkspaceID
	}{
		{
			name: "id",
			want: id.WorkspaceID{},
		},
		{
			name: "id",
			s: Schema{
				workspace: wId,
			},
			want: wId,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.s.Workspace())
		})
	}
}

func TestSchema_Clone(t *testing.T) {
	s := &Schema{id: NewID()}
	c := s.Clone()
	assert.Equal(t, s, c)
	assert.NotSame(t, s, c)

	s = nil
	c = s.Clone()
	assert.Nil(t, c)
}

func TestSchema_HasFieldByKey(t *testing.T) {
	fid1 := NewFieldID()
	fid2 := NewFieldID()
	fid3 := NewFieldID()
	tests := []struct {
		name string
		s    *Schema
		fKey string
		want bool
	}{
		{
			name: "add on empty array",
			s:    &Schema{},
			fKey: "K123123",
			want: false,
		},
		{
			name: "add on not empty array",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1", key: key.New("K123123")}}},
			fKey: "K123123",
			want: true,
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1", key: key.New("K123123")}, {id: fid2, name: "f2", key: key.New("K111222")}, {id: fid3, name: "f3", key: key.New("K123111")}}},
			fKey: "K123123",
			want: true,
		},
		{
			name: "add duplicated field",
			s:    &Schema{fields: []*Field{{id: fid1, name: "f1"}, {id: fid2, name: "f2"}, {id: fid3, name: "f3"}}},
			fKey: "K123123",
			want: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.want, tc.s.HasFieldByKey(tc.fKey))
		})
	}
}
