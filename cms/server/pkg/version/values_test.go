package version

import (
	"testing"
	"time"

	"github.com/reearth/reearthx/util"
	"github.com/stretchr/testify/assert"
)

func TestNewValues(t *testing.T) {
	vx, vy, vz := New(), New(), New()
	v := NewValue(vx, nil, nil, time.Time{}, 0)
	v2 := NewValue(vx, nil, NewRefs("y"), time.Time{}, 1)
	v3 := NewValue(vy, nil, NewRefs("y"), time.Time{}, 1)
	v4 := NewValue(vy, NewVersions(vz), nil, time.Time{}, 0)
	assert.Equal(t, &Values[int]{
		inner: []*Value[int]{v},
	}, NewValues(v))
	assert.Nil(t, NewValues(v, v2))
	assert.Nil(t, NewValues(v2, v3))
	assert.Nil(t, NewValues(v4))
}

func TestValues_MustBeValues(t *testing.T) {
	vx := New()
	v := NewValue(vx, nil, nil, time.Time{}, 0)
	v2 := NewValue(vx, nil, nil, time.Time{}, 0)

	assert.Equal(t, &Values[int]{
		inner: []*Value[int]{v},
	}, MustBeValues(v))
	assert.Panics(t, func() { MustBeValues(v, v2) })
}

func TestValues_IsArchived(t *testing.T) {
	v := &Values[int]{}
	assert.False(t, v.IsArchived())
	assert.Same(t, v, v.SetArchived(true))
	assert.True(t, v.IsArchived())
	assert.Same(t, v, v.SetArchived(false))
	assert.False(t, v.IsArchived())
}

func TestValues_Get(t *testing.T) {
	vx, vy, vz := New(), New(), New()
	v := &Values[string]{
		inner: []*Value[string]{
			NewValue(vx, nil, nil, time.Time{}, "foo1"),
			NewValue(vy, nil, nil, time.Time{}, "foo2"),
			NewValue(vz, nil, NewRefs(Latest), time.Time{}, "foo3"),
		},
	}

	got := v.Get(vz.OrRef())
	assert.Equal(t, NewValue(vz, nil, NewRefs(Latest), time.Time{}, "foo3"), got)

	// cannot modify
	got.value = "d"
	assert.Equal(t, "foo3", v.Get(vz.OrRef()).Value())

	got = v.Get(Ref(Latest).OrVersion())
	assert.Equal(t, NewValue(vz, nil, NewRefs(Latest), time.Time{}, "foo3"), got)

	// cannot modify
	got.value = "d"
	assert.Equal(t, "foo3", v.Get(Ref(Latest).OrVersion()).Value())

	assert.Nil(t, v.Get(New().OrRef()))
	assert.Nil(t, v.Get(Ref("main2").OrVersion()))
}

func TestValues_Latest(t *testing.T) {
	vx, vy, vz := New(), New(), New()
	assert.Equal(
		t,
		NewValue(vx, nil, NewRefs("latest"), time.Time{}, ""),
		(&Values[string]{
			inner: []*Value[string]{
				NewValue(vx, nil, NewRefs("latest"), time.Time{}, ""),
				NewValue(vy, nil, nil, time.Time{}, ""),
				NewValue(vz, nil, nil, time.Time{}, ""),
			},
		}).Latest(),
	)
	assert.Nil(t, (&Values[string]{
		inner: []*Value[string]{
			NewValue(vx, nil, NewRefs("a"), time.Time{}, ""),
			NewValue(vy, nil, nil, time.Time{}, ""),
			NewValue(vz, nil, nil, time.Time{}, ""),
		},
	}).Latest())
	assert.Nil(t, (*Values[string])(nil).Latest())
}

func TestValues_LatestVersion(t *testing.T) {
	vx, vy, vz := New(), New(), New()
	assert.Equal(t, vx.Ref(), (&Values[string]{
		inner: []*Value[string]{
			NewValue(vx, nil, NewRefs("latest"), time.Time{}, ""),
			NewValue(vy, nil, nil, time.Time{}, ""),
			NewValue(vz, nil, nil, time.Time{}, ""),
		},
	}).LatestVersion())
	assert.Nil(t, (&Values[string]{
		inner: []*Value[string]{
			NewValue(vx, nil, NewRefs("a"), time.Time{}, ""),
			NewValue(vy, nil, nil, time.Time{}, ""),
			NewValue(vz, nil, nil, time.Time{}, ""),
		},
	}).LatestVersion())
	assert.Nil(t, (*Values[string])(nil).LatestVersion())
}

func TestValues_All(t *testing.T) {
	vx, vy, vz := New(), New(), New()
	v := &Values[string]{
		inner: []*Value[string]{
			NewValue(vx, nil, NewRefs("latest"), time.Time{}, "a"),
			NewValue(vy, nil, nil, time.Time{}, "b"),
			NewValue(vz, nil, nil, time.Time{}, "c"),
		},
	}
	got := v.All()
	assert.Equal(t, v.inner, got)
	assert.NotSame(t, v.inner, got)
	assert.Nil(t, (*Values[string])(nil).All())
}

func TestValues_Clone(t *testing.T) {
	vx, vy, vz := New(), New(), New()
	v := &Values[string]{
		inner: []*Value[string]{
			NewValue(vx, nil, NewRefs("latest"), time.Time{}, "a"),
			NewValue(vy, nil, nil, time.Time{}, "b"),
			NewValue(vz, nil, nil, time.Time{}, "c"),
		},
		archived: true,
	}
	got := v.Clone()
	assert.Equal(t, v, got)
	assert.NotSame(t, v, got)
	assert.NotSame(t, v.inner, got.inner)
	assert.Nil(t, (*Values[string])(nil).Clone())
}

func TestValues_Add(t *testing.T) {
	now := util.Now()
	defer util.MockNow(now)()

	vx, vy := New(), New()
	v := &Values[string]{
		inner: []*Value[string]{
			NewValue(vx, NewVersions(vy), NewRefs(Latest), time.Time{}, "1"),
			NewValue(vy, nil, NewRefs("a"), time.Time{}, "2"),
		},
	}

	v.Add("3", Ref("a").OrVersion().Ref())
	vv := v.Get(Ref("a").OrVersion())
	assert.Equal(t, NewValue(vv.Version(), NewVersions(vy), NewRefs("a"), now, "3"), vv)
	assert.Equal(t, NewValue(vy, nil, nil, time.Time{}, "2"), v.Get(vy.OrRef()))
	assert.True(t, v.validate())

	v.Add("3", Ref("").OrVersion().Ref())
	assert.Nil(t, v.Get(Ref("").OrVersion()))

	v.archived = true
	v.Add("3", Ref("xxx").OrVersion().Ref())
	assert.Nil(t, v.Get(Ref("xxx").OrVersion()))
}

func TestValues_UpdateRef(t *testing.T) {
	vx, vy := New(), New()

	type args struct {
		ref Ref
		vr  *VersionOrRef
	}

	tests := []struct {
		name   string
		target *Values[string]
		args   args
		want   *Values[string]
	}{
		{
			name: "ref is not found",
			target: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
					NewValue(vy, nil, NewRefs("B"), time.Time{}, "b"),
				}},
			args: args{
				ref: "A",
				vr:  nil,
			},
			want: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
					NewValue(vy, nil, NewRefs("B"), time.Time{}, "b"),
				},
			},
		},
		{
			name: "ref should be deleted",
			target: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
					NewValue(vy, nil, NewRefs("B"), time.Time{}, "b"),
				},
			},
			args: args{
				ref: "B",
				vr:  nil,
			},
			want: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
					NewValue(vy, nil, nil, time.Time{}, "b"),
				},
			},
		},
		{
			name: "new ref should be set",
			target: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
					NewValue(vy, nil, NewRefs("B"), time.Time{}, "b"),
				},
			},
			args: args{
				ref: "A",
				vr:  vx.OrRef().Ref(),
			},
			want: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, NewRefs("A"), time.Time{}, "a"),
					NewValue(vy, nil, NewRefs("B"), time.Time{}, "b"),
				},
			},
		},
		{
			name: "ref should be moved",
			target: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, NewRefs("A"), time.Time{}, "a"),
					NewValue(vy, nil, NewRefs("B"), time.Time{}, "b"),
				},
			},
			args: args{
				ref: "B",
				vr:  Ref("A").OrVersion().Ref(),
			},
			want: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, NewRefs("A", "B"), time.Time{}, "a"),
					NewValue(vy, nil, nil, time.Time{}, "b"),
				},
			},
		},
		{
			name: "latest should not be updated",
			target: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
					NewValue(vy, NewVersions(vx), NewRefs(Latest), time.Time{}, "b"),
				},
			},
			args: args{
				ref: Latest,
				vr:  vx.OrRef().Ref(),
			},
			want: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
					NewValue(vy, NewVersions(vx), NewRefs(Latest), time.Time{}, "b"),
				},
			},
		},
		{
			name: "archived should not be updated",
			target: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
				},
				archived: true,
			},
			args: args{
				ref: "x",
				vr:  vx.OrRef().Ref(),
			},
			want: &Values[string]{
				inner: []*Value[string]{
					NewValue(vx, nil, nil, time.Time{}, "a"),
				},
				archived: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			target := tt.target
			target.UpdateRef(tt.args.ref, tt.args.vr)
			assert.Equal(t, tt.want, target)
			assert.True(t, target.validate())
		})
	}
}

func TestUnwrapValues(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, UnwrapValues([]*Value[int]{{value: 1}, {value: 2}, {value: 3}}))
	assert.Nil(t, UnwrapValues[int](nil))
}
