package memorygit

import (
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/util"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestVersionedSyncMap_Load(t *testing.T) {
	vx := version.New()
	vsm := &VersionedSyncMap[string, string]{
		m: util.SyncMapFrom(map[string]*version.Values[string]{
			"a": version.MustBeValues(
				version.NewValue(vx, nil, nil, time.Time{}, "A"),
			),
			"b": version.MustBeValues(
				version.NewValue(vx, nil, version.NewRefs("a"), time.Time{}, "B"),
			),
		}),
	}

	tests := []struct {
		name  string
		m     *VersionedSyncMap[string, string]
		input struct {
			key string
			vor version.VersionOrRef
		}
		want struct {
			output string
			ok     bool
		}
	}{
		{
			name: "should load by version",
			m:    vsm,
			input: struct {
				key string
				vor version.VersionOrRef
			}{
				key: "a",
				vor: vx.OrRef(),
			},
			want: struct {
				output string
				ok     bool
			}{
				output: "A",
				ok:     true,
			},
		},
		{
			name: "should load by ref",
			m:    vsm,
			input: struct {
				key string
				vor version.VersionOrRef
			}{
				key: "b",
				vor: version.Ref("a").OrVersion(),
			},
			want: struct {
				output string
				ok     bool
			}{
				output: "B",
				ok:     true,
			},
		},
		{
			name: "should fail to find ref",
			m:    vsm,
			input: struct {
				key string
				vor version.VersionOrRef
			}{
				key: "b",
				vor: version.Ref("xxxx").OrVersion(),
			},
			want: struct {
				output string
				ok     bool
			}{
				output: "",
				ok:     false,
			},
		},
		{
			name: "should fail to find version",
			m:    vsm,
			input: struct {
				key string
				vor version.VersionOrRef
			}{
				key: "a",
				vor: version.New().OrRef(),
			},
			want: struct {
				output string
				ok     bool
			}{
				output: "",
				ok:     false,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, ok := tc.m.Load(tc.input.key, tc.input.vor)
			if tc.want.output == "" {
				assert.Nil(t, got)
			} else {
				assert.Equal(t, tc.want.output, got.Value())
			}
			assert.Equal(t, tc.want.ok, ok)
		})
	}
}

func TestVersionedSyncMap_LoadAll(t *testing.T) {
	vx, vy := version.New(), version.New()
	vsm := &VersionedSyncMap[string, string]{m: util.SyncMapFrom(
		map[string]*version.Values[string]{
			"a": version.MustBeValues(
				version.NewValue(vx, nil, nil, time.Time{}, "A"),
			),
			"b": version.MustBeValues(
				version.NewValue(vx, nil, version.NewRefs("a"), time.Time{}, "B"),
			),
			"c": version.MustBeValues(
				version.NewValue(vx, nil, nil, time.Time{}, "C"),
			),
			"d": version.MustBeValues(
				version.NewValue(vy, nil, version.NewRefs("a"), time.Time{}, "D"),
			),
		},
	)}
	tests := []struct {
		name  string
		m     *VersionedSyncMap[string, string]
		input struct {
			keys []string
			vor  version.VersionOrRef
		}
		want []string
	}{
		{
			name: "should load by version",
			m:    vsm,
			input: struct {
				keys []string
				vor  version.VersionOrRef
			}{
				keys: []string{"a", "b"},
				vor:  vx.OrRef(),
			},
			want: []string{"A", "B"},
		},
		{
			name: "should load by ref",
			m:    vsm,
			input: struct {
				keys []string
				vor  version.VersionOrRef
			}{
				keys: []string{"b", "d"},
				vor:  version.Ref("a").OrVersion(),
			},
			want: []string{"B", "D"},
		},
		{
			name: "should not load",
			m:    vsm,
			input: struct {
				keys []string
				vor  version.VersionOrRef
			}{
				keys: []string{"d"},
				vor:  vx.OrRef(),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.m.LoadAll(tc.input.keys, &tc.input.vor)
			got2 := version.UnwrapValues(got)
			slices.Sort(got2)
			assert.Equal(t, tc.want, got2)
		})
	}
}

func TestVersionedSyncMap_LoadAllVersions(t *testing.T) {
	vx, vy, vz := version.New(), version.New(), version.New()
	vsm := &VersionedSyncMap[string, string]{m: util.SyncMapFrom(
		map[string]*version.Values[string]{
			"a": version.MustBeValues(
				version.NewValue(vx, nil, nil, time.Time{}, "A"),
				version.NewValue(vy, nil, nil, time.Time{}, "B"),
				version.NewValue(vz, nil, nil, time.Time{}, "C"),
			),
		},
	)}
	tests := []struct {
		name  string
		m     *VersionedSyncMap[string, string]
		input struct {
			key string
		}
		want *version.Values[string]
	}{
		{
			name: "should load by version",
			m:    vsm,
			input: struct {
				key string
			}{
				key: "a",
			},
			want: version.MustBeValues(
				version.NewValue(vx, nil, nil, time.Time{}, "A"),
				version.NewValue(vy, nil, nil, time.Time{}, "B"),
				version.NewValue(vz, nil, nil, time.Time{}, "C"),
			),
		},
		{
			name: "should not load",
			m:    vsm,
			input: struct {
				key string
			}{
				key: "d",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.m.LoadAllVersions(tc.input.key)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestVersionedSyncMap_Store(t *testing.T) {
	vm := &VersionedSyncMap[string, string]{
		m: util.SyncMapFrom(map[string]*version.Values[string]{}),
	}

	_, ok := vm.Load("a", version.Latest.OrVersion())
	assert.False(t, ok)

	vm.SaveOne("a", "b", nil)
	got, ok := vm.Load("a", version.Latest.OrVersion())
	assert.True(t, ok)
	assert.Equal(t, "b", got.Value())

	vm.SaveOne("a", "c", nil)
	got2, ok2 := vm.Load("a", version.Latest.OrVersion())
	assert.True(t, ok2)
	assert.Equal(t, "c", got2.Value())

	vm.SaveOne("a", "d", version.Latest.OrVersion().Ref())
	got3, ok3 := vm.Load("a", version.Latest.OrVersion())
	assert.True(t, ok3)
	assert.Equal(t, "d", got3.Value())
}

func TestVersionedSyncMap_UpdateRef(t *testing.T) {
	vx := version.New()

	type args struct {
		key string
		ref version.Ref
		vr  *version.VersionOrRef
	}
	tests := []struct {
		name   string
		target *VersionedSyncMap[string, string]
		args   args
		want   *version.Values[string]
	}{
		{
			name: "set ref",
			target: &VersionedSyncMap[string, string]{
				m: util.SyncMapFrom(
					map[string]*version.Values[string]{
						"1": version.MustBeValues(version.NewValue(vx, nil, nil, time.Time{}, "a")),
						"2": version.MustBeValues(version.NewValue(vx, nil, nil, time.Time{}, "a")),
					},
				),
			},
			args: args{
				key: "1",
				ref: "A",
				vr:  vx.OrRef().Ref(),
			},
			want: version.MustBeValues(
				version.NewValue(vx, nil, version.NewRefs("A"), time.Time{}, "a"),
			),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.target.UpdateRef(tt.args.key, tt.args.ref, tt.args.vr)

			f, _ := tt.target.m.Load(tt.args.key)
			assert.Equal(t, tt.want, f)
		})
	}
}
