package gqlmodel

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/reearth/reearth-cms/server/pkg/file"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestConvert_FromFile(t *testing.T) {
	buf := bytes.NewBufferString("aaa")
	buflen := int64(buf.Len())

	u1 := graphql.Upload{
		File:        strings.NewReader("aaa"),
		Filename:    "aaa.txt",
		Size:        buflen,
		ContentType: "text/plain",
	}
	want1 := file.File{
		Content:     io.NopCloser(strings.NewReader("aaa")),
		Path:        "aaa.txt",
		Size:        buflen,
		ContentType: "text/plain",
	}

	var u2 *graphql.Upload = nil
	want2 := (*file.File)(nil)

	tests := []struct {
		name string
		arg  *graphql.Upload
		want *file.File
	}{
		{
			name: "from file valid",
			arg:  &u1,
			want: &want1,
		},
		{
			name: "from file nil",
			arg:  u2,
			want: want2,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := FromFile(tc.arg)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestToPageInfo(t *testing.T) {
	tests := []struct {
		name string
		args *usecasex.PageInfo
		want *PageInfo
	}{
		{
			name: "nil",
			args: nil,
			want: &PageInfo{},
		},
		{
			name: "success",
			args: &usecasex.PageInfo{
				TotalCount:      0,
				StartCursor:     usecasex.CursorFromRef(lo.ToPtr("c1")),
				EndCursor:       nil,
				HasNextPage:     true,
				HasPreviousPage: false,
			},
			want: &PageInfo{
				StartCursor:     usecasex.CursorFromRef(lo.ToPtr("c1")),
				EndCursor:       nil,
				HasNextPage:     true,
				HasPreviousPage: false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, ToPageInfo(tt.args))
		})
	}
}

func TestPagination_Into(t *testing.T) {
	tests := []struct {
		name string
		args *Pagination
		want *usecasex.Pagination
	}{
		{
			name: "nil",
			args: nil,
			want: nil,
		},
		{
			name: "success",
			args: &Pagination{
				First:  nil,
				Last:   nil,
				After:  nil,
				Before: nil,
			},
			want: usecasex.CursorPagination{
				Before: nil,
				After:  nil,
				First:  nil,
				Last:   nil,
			}.Wrap(),
		},
		{
			name: "success 2",
			args: &Pagination{
				First:  lo.ToPtr(10),
				Last:   nil,
				After:  usecasex.CursorFromRef(lo.ToPtr("c1")),
				Before: nil,
			},
			want: usecasex.CursorPagination{
				Before: nil,
				After:  usecasex.CursorFromRef(lo.ToPtr("c1")),
				First:  lo.ToPtr(int64(10)),
				Last:   nil,
			}.Wrap(),
		},
		{
			name: "success offset",
			args: &Pagination{
				First:  lo.ToPtr(10),
				Offset: lo.ToPtr(50),
			},
			want: usecasex.OffsetPagination{
				Limit:  10,
				Offset: 50,
			}.Wrap(),
		},
		{
			name: "success offset",
			args: &Pagination{
				First:  nil,
				Offset: lo.ToPtr(50),
			},
			want: usecasex.OffsetPagination{
				Limit:  50,
				Offset: 50,
			}.Wrap(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.args.Into())
		})
	}
}
