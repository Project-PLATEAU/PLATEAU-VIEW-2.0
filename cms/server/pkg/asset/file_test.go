package asset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_FileType(t *testing.T) {
	c := NewFile().Build()
	f := NewFile().Name("aaa.txt").Path("/aaa.txt").Size(10).GuessContentType().Children([]*File{c}).Build()

	assert.Equal(t, "aaa.txt", f.Name())
	assert.Equal(t, uint64(10), f.Size())
	assert.Equal(t, "text/plain; charset=utf-8", f.ContentType())
	assert.Equal(t, "/aaa.txt", f.Path())
	assert.Equal(t, []*File{c}, f.Children())

	f.SetName("bbb")
	assert.Equal(t, "bbb", f.Name())

	c2 := NewFile().Build()
	f.AppendChild(c2)
	assert.Equal(t, []*File{c, c2}, f.Children())

	dir := NewFile().Name("dir").Path("/aaa").Children([]*File{c}).Build()
	assert.True(t, dir.IsDir())
}

func TestFile_Children(t *testing.T) {
	// nil file should return nil children
	var got *File = nil
	assert.Nil(t, got.Children())

	// file.Children() should return file.children
	c := []*File{}
	got = &File{
		children: c,
	}
	assert.Equal(t, c, got.Children())
}

func TestFile_Files(t *testing.T) {
	f := &File{
		path: "aaa",
		children: []*File{
			{
				path: "aaa/a",
				children: []*File{
					{
						path: "aaa/a/a.txt",
					},
				},
			},
			{
				path: "aaa/b.txt",
			},
		},
	}

	assert.Equal(t, []*File{
		{
			path: "aaa/a/a.txt",
		},
		{
			path: "aaa/b.txt",
		},
	}, f.Files())
}

func Test_FoldFiles(t *testing.T) {
	assert.Equal(t,
		&File{
			name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip",
			children: []*File{
				{name: "a.txt", path: "/a.txt", size: 10, contentType: "text/plain"},
				{name: "b.txt", path: "/b.txt", size: 20, contentType: "text/plain"},
			},
		},
		FoldFiles(
			[]*File{
				{name: "a.txt", path: "/a.txt", size: 10, contentType: "text/plain"},
				{name: "b.txt", path: "/b.txt", size: 20, contentType: "text/plain"},
			},
			&File{name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip"},
		),
	)
	assert.Equal(t,
		&File{
			name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip",
			children: []*File{
				{name: "hello", path: "/hello", size: 0, contentType: "", children: []*File{
					{name: "a.txt", path: "/hello/a.txt", size: 10, contentType: "text/plain"},
					{name: "b.txt", path: "/hello/b.txt", size: 20, contentType: "text/plain"},
				}},
			},
		},
		FoldFiles(
			[]*File{
				{name: "a.txt", path: "/hello/a.txt", size: 10, contentType: "text/plain"},
				{name: "b.txt", path: "/hello/b.txt", size: 20, contentType: "text/plain"},
			},
			&File{name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip"},
		),
	)

	assert.Equal(t,
		&File{
			name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip",
			children: []*File{
				{name: "hello", path: "/hello", size: 0, contentType: "", children: []*File{
					{name: "c.txt", path: "/hello/c.txt", size: 20, contentType: "text/plain"},
					{name: "good", path: "/hello/good", size: 0, contentType: "", children: []*File{
						{name: "a.txt", path: "/hello/good/a.txt", size: 10, contentType: "text/plain"},
						{name: "b.txt", path: "/hello/good/b.txt", size: 10, contentType: "text/plain"},
					}},
				}},
			},
		},
		FoldFiles(
			[]*File{
				{name: "a.txt", path: "/hello/good/a.txt", size: 10, contentType: "text/plain"},
				{name: "b.txt", path: "/hello/good/b.txt", size: 10, contentType: "text/plain"},
				{name: "c.txt", path: "/hello/c.txt", size: 20, contentType: "text/plain"},
			},
			&File{name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip"},
		),
	)
	assert.Equal(t,
		&File{
			name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip",
			children: []*File{
				{name: "hello", path: "/hello", size: 0, contentType: "", children: []*File{
					{name: "hello", path: "/hello/hello", children: []*File{
						{name: "a.txt", path: "/hello/hello/a.txt", size: 10, contentType: "text/plain"},
						{name: "b.txt", path: "/hello/hello/b.txt", size: 10, contentType: "text/plain"},
						{name: "c", path: "/hello/hello/c", children: []*File{
							{name: "d.txt", path: "/hello/hello/c/d.txt", size: 20, contentType: "text/plain"},
						}},
					}},
				},
				},
			},
		},
		FoldFiles(
			[]*File{
				{name: "a.txt", path: "/hello/hello/a.txt", size: 10, contentType: "text/plain"},
				{name: "b.txt", path: "/hello/hello/b.txt", size: 10, contentType: "text/plain"},
				{name: "d.txt", path: "/hello/hello/c/d.txt", size: 20, contentType: "text/plain"},
			},
			&File{name: "hello.zip", path: "/hello.zip", size: 100, contentType: "application/zip"},
		),
	)
}

func Test_File_RootPath(t *testing.T) {
	assert.Equal(t, "xx/xxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/hoge.zip", (&File{path: "hoge.zip"}).RootPath("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"))
}
