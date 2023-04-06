package asset

import (
	"mime"
	"path"
	"strings"

	"golang.org/x/exp/slices"
)

type FileBuilder struct {
	f                 *File
	detectContentType bool
}

func NewFile() *FileBuilder {
	return &FileBuilder{
		f: &File{},
	}
}

func (b *FileBuilder) Name(name string) *FileBuilder {
	b.f.name = name
	return b
}

func (b *FileBuilder) ContentType(contentType string) *FileBuilder {
	b.f.contentType = contentType
	return b
}

func (b *FileBuilder) Path(path string) *FileBuilder {
	if !strings.HasPrefix(path, "/") && path != "" {
		path = "/" + path
	}

	b.f.path = path
	return b
}

func (b *FileBuilder) Size(size uint64) *FileBuilder {
	b.f.size = size
	return b
}

func (b *FileBuilder) Children(children []*File) *FileBuilder {
	b.f.children = slices.Clone(children)
	return b
}

func (b *FileBuilder) GuessContentType() *FileBuilder {
	b.detectContentType = true
	return b
}

func (b *FileBuilder) Dir() *FileBuilder {
	if b.f.children == nil {
		b.f.children = []*File{}
	}
	return b
}

func (b *FileBuilder) Build() *File {
	if b.detectContentType {
		b.f.contentType = mime.TypeByExtension(path.Ext(b.f.path))
	}

	return b.f
}
