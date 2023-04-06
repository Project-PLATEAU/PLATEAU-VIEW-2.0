package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
)

func TestModel_Clone(t *testing.T) {
	mId := NewID()
	pId := id.NewProjectID()
	sId := id.NewSchemaID()
	now := time.Now()
	tests := []struct {
		name  string
		model *Model
	}{
		{
			name: "test",
			model: &Model{
				id:          mId,
				project:     pId,
				schema:      sId,
				name:        "n1",
				description: "d1",
				key:         key.New("123456"),
				public:      false,
				updatedAt:   now,
			},
		},
		{
			name:  "test",
			model: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := tt.model.Clone()
			if tt.model == nil {
				assert.Nil(t, c)
				return
			}
			assert.Equal(t, tt.model, c)
			assert.NotSame(t, tt.model, c)
			assert.NotSame(t, tt.model.id, c.id)
			assert.NotSame(t, tt.model.schema, c.schema)
			assert.NotSame(t, tt.model.project, c.project)
			assert.NotSame(t, tt.model.name, c.name)
			assert.NotSame(t, tt.model.description, c.description)
			assert.NotSame(t, tt.model.key, c.key)
			assert.NotSame(t, tt.model.public, c.public)
			assert.NotSame(t, tt.model.updatedAt, c.updatedAt)
		})
	}
}

func TestModel_CreatedAt(t *testing.T) {
	mId := NewID()
	tests := []struct {
		name  string
		model Model
		want  time.Time
	}{
		{
			name: "test",
			model: Model{
				id: mId,
			},
			want: mId.Timestamp(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.CreatedAt())
		})
	}
}

func TestModel_Description(t *testing.T) {
	tests := []struct {
		name  string
		model Model
		want  string
	}{
		{
			name: "test",
			model: Model{
				description: "d1",
			},
			want: "d1",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.Description())
		})
	}
}

func TestModel_ID(t *testing.T) {
	mId := NewID()
	tests := []struct {
		name  string
		model Model
		want  ID
	}{
		{
			name: "test",
			model: Model{
				id: mId,
			},
			want: mId,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.ID())
		})
	}
}

func TestModel_Key(t *testing.T) {
	tests := []struct {
		name  string
		model Model
		want  key.Key
	}{
		{
			name: "test",
			model: Model{
				key: key.New("123456"),
			},
			want: key.New("123456"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.Key())
		})
	}
}

func TestModel_Name(t *testing.T) {
	tests := []struct {
		name  string
		model Model
		want  string
	}{
		{
			name: "test",
			model: Model{
				name: "n1",
			},
			want: "n1",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.Name())
		})
	}
}

func TestModel_Project(t *testing.T) {
	pId := id.NewProjectID()
	tests := []struct {
		name  string
		model Model
		want  ProjectID
	}{
		{
			name: "test",
			model: Model{
				project: pId,
			},
			want: pId,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.Project(), "Project()")
		})
	}
}

func TestModel_Public(t *testing.T) {
	tests := []struct {
		name  string
		model Model
		want  bool
	}{
		{
			name: "public true",
			model: Model{
				public: true,
			},
			want: true,
		},
		{
			name: "public false",
			model: Model{
				public: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.Public())
		})
	}
}

func TestModel_Schema(t *testing.T) {
	sId := id.NewSchemaID()
	tests := []struct {
		name  string
		model Model
		want  SchemaID
	}{
		{
			name: "test",
			model: Model{
				schema: sId,
			},
			want: sId,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.Schema())
		})
	}
}

func TestModel_SetDescription(t *testing.T) {
	type args struct {
		description string
	}
	tests := []struct {
		name string
		want Model
		args args
	}{
		{
			name: "test",
			args: args{
				description: "d1",
			},
			want: Model{
				description: "d1",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := Model{}
			m.SetDescription(tt.args.description)
			assert.Equal(t, tt.want, m)
		})
	}
}

func TestModel_SetKey(t *testing.T) {
	type args struct {
		key key.Key
	}
	tests := []struct {
		name    string
		args    args
		want    Model
		wantErr error
	}{
		{
			name:    "pass",
			args:    args{key: key.New("123456")},
			want:    Model{key: key.New("123456")},
			wantErr: nil,
		},
		{
			name: "fail",
			args: args{key: key.New("a")},
			want: Model{},
			wantErr: &rerror.Error{
				Label: ErrInvalidKey,
				Err:   fmt.Errorf("%s", "a"),
			},
		},
		{
			name: "fail 2",
			args: args{key: key.New("_aaaaaaaa")},
			want: Model{},
			wantErr: &rerror.Error{
				Label: ErrInvalidKey,
				Err:   fmt.Errorf("%s", "_aaaaaaaa"),
			},
		},
		{
			name: "fail 3",
			args: args{key: key.New("-aaaaaaaa")},
			want: Model{},
			wantErr: &rerror.Error{
				Label: ErrInvalidKey,
				Err:   fmt.Errorf("%s", "-aaaaaaaa"),
			},
		},
		{
			name: "fails assets",
			args: args{key: key.New("assets")},
			want: Model{},
			wantErr: &rerror.Error{
				Label: ErrInvalidKey,
				Err:   fmt.Errorf("%s", "assets"),
			},
		},
		{
			name: "fails items",
			args: args{key: key.New("items")},
			want: Model{},
			wantErr: &rerror.Error{
				Label: ErrInvalidKey,
				Err:   fmt.Errorf("%s", "items"),
			},
		},
		{
			name: "empty",
			args: args{key: key.New("")},
			want: Model{},
			wantErr: &rerror.Error{
				Label: ErrInvalidKey,
				Err:   fmt.Errorf("%s", ""),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := Model{}
			err := m.SetKey(tt.args.key)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, m)
		})
	}
}

func TestModel_SetName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		want Model
		args args
	}{
		{
			name: "test",
			args: args{
				name: "n1",
			},
			want: Model{
				name: "n1",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := Model{}
			m.SetName(tt.args.name)
			assert.Equal(t, tt.want, m)
		})
	}
}

func TestModel_SetPublic(t *testing.T) {
	m := &Model{public: false}
	m.SetPublic(true)
	assert.Equal(t, &Model{public: true}, m)

	m = &Model{public: true}
	m.SetPublic(false)
	assert.Equal(t, &Model{public: false}, m)
}

func TestModel_SetUpdatedAt(t *testing.T) {
	now := time.Now()
	type args struct {
		updateAt time.Time
	}
	tests := []struct {
		name string
		want Model
		args args
	}{
		{
			name: "test",
			args: args{
				updateAt: now,
			},
			want: Model{
				updatedAt: now,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := Model{}
			m.SetUpdatedAt(tt.args.updateAt)
			assert.Equal(t, tt.want, m)
		})
	}
}

func TestModel_UpdatedAt(t *testing.T) {
	now := time.Now()
	mId := NewID()
	tests := []struct {
		name  string
		model Model
		want  time.Time
	}{
		{
			name: "test",
			model: Model{
				updatedAt: now,
			},
			want: now,
		},
		{
			name: "test",
			model: Model{
				id: mId,
			},
			want: mId.Timestamp(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.model.UpdatedAt())
		})
	}
}
