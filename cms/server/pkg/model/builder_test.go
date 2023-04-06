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

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Builder
	}{
		{
			name: "test",
			want: &Builder{
				model: &Model{},
				k:     key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, New())
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	mId := NewID()
	pId := id.NewProjectID()
	sId := id.NewSchemaID()
	now := time.Now()
	type fields struct {
		m *Model
		k key.Key
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Model
		wantErr error
	}{
		{
			name: "pass",
			fields: fields{
				m: &Model{
					id:          mId,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want: &Model{
				id:          mId,
				project:     pId,
				schema:      sId,
				name:        "m1",
				description: "m1 desc",
				key:         key.New("T123456"),
				public:      false,
				updatedAt:   now,
			},
			wantErr: nil,
		},
		{
			name: "pass with out updated at",
			fields: fields{
				m: &Model{
					id:          mId,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					// updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want: &Model{
				id:          mId,
				project:     pId,
				schema:      sId,
				name:        "m1",
				description: "m1 desc",
				key:         key.New("T123456"),
				public:      false,
				updatedAt:   mId.Timestamp(),
			},
			wantErr: nil,
		},
		{
			name: "fail 1",
			fields: fields{
				m: &Model{
					// id:          nil,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name: "fail 2",
			fields: fields{
				m: &Model{
					id:      mId,
					project: pId,
					// schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name: "fail 3",
			fields: fields{
				m: &Model{
					id:          mId,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				// k: key.New("T123456"),
			},
			want: nil,
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

			b := &Builder{
				model: tt.fields.m,
				k:     tt.fields.k,
			}
			got, err := b.Build()
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBuilder_MustBuild(t *testing.T) {
	mId := NewID()
	pId := id.NewProjectID()
	sId := id.NewSchemaID()
	now := time.Now()
	type fields struct {
		m *Model
		k key.Key
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Model
		wantErr error
	}{
		{
			name: "pass",
			fields: fields{
				m: &Model{
					id:          mId,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want: &Model{
				id:          mId,
				project:     pId,
				schema:      sId,
				name:        "m1",
				description: "m1 desc",
				key:         key.New("T123456"),
				public:      false,
				updatedAt:   now,
			},
			wantErr: nil,
		},
		{
			name: "pass with out updated at",
			fields: fields{
				m: &Model{
					id:          mId,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					// updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want: &Model{
				id:          mId,
				project:     pId,
				schema:      sId,
				name:        "m1",
				description: "m1 desc",
				key:         key.New("T123456"),
				public:      false,
				updatedAt:   mId.Timestamp(),
			},
			wantErr: nil,
		},
		{
			name: "fail 1",
			fields: fields{
				m: &Model{
					// id:          nil,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name: "fail 2",
			fields: fields{
				m: &Model{
					id:      mId,
					project: pId,
					// schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				k: key.New("T123456"),
			},
			want:    nil,
			wantErr: ErrInvalidID,
		},
		{
			name: "fail 3",
			fields: fields{
				m: &Model{
					id:          mId,
					project:     pId,
					schema:      sId,
					name:        "m1",
					description: "m1 desc",
					key:         key.New("T123456"),
					public:      false,
					updatedAt:   now,
				},
				// k: key.New("T123456"),
			},
			want: nil,
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

			b := &Builder{
				model: tt.fields.m,
				k:     tt.fields.k,
			}

			if tt.wantErr != nil {
				assert.PanicsWithError(t, tt.wantErr.Error(), func() {
					b.MustBuild()
				})
				return
			}
			assert.Equal(t, tt.want, b.MustBuild())
		})
	}
}

func TestBuilder_Description(t *testing.T) {
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				description: "d1",
			},
			want: &Builder{
				model: &Model{
					description: "d1",
				},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equal(t, tt.want, b.Description(tt.args.description))
		})
	}
}

func TestBuilder_ID(t *testing.T) {
	mId := NewID()
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		id ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				id: mId,
			},
			want: &Builder{
				model: &Model{
					id: mId,
				},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equal(t, tt.want, b.ID(tt.args.id))
		})
	}
}

func TestBuilder_Public(t *testing.T) {
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		public bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test 1",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				public: true,
			},
			want: &Builder{
				model: &Model{
					public: true,
				},
				k: key.Key{},
			},
		},
		{
			name: "test 2",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				public: false,
			},
			want: &Builder{
				model: &Model{
					public: false,
				},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equal(t, tt.want, b.Public(tt.args.public))
		})
	}
}

func TestBuilder_Key(t *testing.T) {
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		key key.Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test 1",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				key: key.New("123321"),
			},
			want: &Builder{
				model: &Model{},
				k:     key.New("123321"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equalf(t, tt.want, b.Key(tt.args.key), "Key(%v)", tt.args.key)
		})
	}
}

func TestBuilder_Name(t *testing.T) {
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				name: "n1",
			},
			want: &Builder{
				model: &Model{
					name: "n1",
				},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equal(t, tt.want, b.Name(tt.args.name))
		})
	}
}

func TestBuilder_NewID(t *testing.T) {
	type fields struct {
		p *Model
		k key.Key
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test",
			fields: fields{
				p: &Model{},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.p,
				k:     tt.fields.k,
			}
			b.NewID()
			assert.False(t, b.model.id.IsEmpty())
		})
	}
}

func TestBuilder_Project(t *testing.T) {
	pId := id.NewProjectID()
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		p id.ProjectID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				p: pId,
			},
			want: &Builder{
				model: &Model{
					project: pId,
				},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equal(t, tt.want, b.Project(tt.args.p))
		})
	}
}

func TestBuilder_RandomKey(t *testing.T) {
	type fields struct {
		p *Model
		k key.Key
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test",
			fields: fields{
				p: &Model{},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.p,
				k:     tt.fields.k,
			}
			b.RandomKey()
			assert.True(t, b.k.IsValid())
		})
	}
}

func TestBuilder_Schema(t *testing.T) {
	sId := id.NewSchemaID()
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		s id.SchemaID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				s: sId,
			},
			want: &Builder{
				model: &Model{
					schema: sId,
				},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equal(t, tt.want, b.Schema(tt.args.s))
		})
	}
}

func TestBuilder_UpdatedAt(t *testing.T) {
	now := time.Now()
	type fields struct {
		model *Model
		k     key.Key
	}
	type args struct {
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "test",
			fields: fields{
				model: &Model{},
				k:     key.Key{},
			},
			args: args{
				updatedAt: now,
			},
			want: &Builder{
				model: &Model{
					updatedAt: now,
				},
				k: key.Key{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := &Builder{
				model: tt.fields.model,
				k:     tt.fields.k,
			}
			assert.Equal(t, tt.want, b.UpdatedAt(tt.args.updatedAt))
		})
	}
}
