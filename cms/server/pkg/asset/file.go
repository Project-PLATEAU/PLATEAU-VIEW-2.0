package asset

import (
	"path"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type File struct {
	name        string
	size        uint64
	contentType string
	path        string
	children    []*File
}

func (f *File) Name() string {
	if f == nil {
		return ""
	}
	return f.name
}

func (f *File) SetName(n string) {
	f.name = n
}

func (f *File) Size() uint64 {
	if f == nil {
		return 0
	}
	return f.size
}

func (f *File) ContentType() string {
	if f == nil {
		return ""
	}
	return f.contentType
}

func (f *File) Path() string {
	if f == nil {
		return ""
	}
	return f.path
}

func (f *File) Children() []*File {
	if f == nil {
		return nil
	}
	return slices.Clone(f.children)
}

func (f *File) IsDir() bool {
	return f != nil && f.children != nil
}

func (f *File) AppendChild(c *File) {
	if f == nil {
		return
	}
	f.children = append(f.children, c)
}

func (f *File) Clone() *File {
	if f == nil {
		return nil
	}

	var children []*File
	if f.children != nil {
		children = lo.Map(f.children, func(f *File, _ int) *File { return f.Clone() })
	}

	return &File{
		name:        f.name,
		size:        f.size,
		contentType: f.contentType,
		path:        f.path,
		children:    children,
	}
}

func (f *File) Files() (res []*File) {
	if f == nil {
		return nil
	}
	if len(f.children) > 0 {
		for _, c := range f.children {
			res = append(res, c.Files()...)
		}
	} else {
		res = append(res, f)
	}
	return
}

func (f *File) RootPath(uuid string) string {
	if f == nil {
		return ""
	}
	return path.Join(uuid[:2], uuid[2:], f.path)
}

// TODO:improve perfomance later
func FoldFiles(files []*File, parent *File) *File {
	files = slices.Clone(files)
	slices.SortFunc(files, func(a, b *File) bool {
		return strings.Compare(a.Path(), b.Path()) < 0
	})

	var skipIndexes []int
	for i := range files {
		if slices.Contains(skipIndexes, i) {
			continue
		}

		// parent: /a/b
		// file: /a/b/c/d.txt
		// diff: c
		parentDir := strings.TrimPrefix(parent.Path(), "/")
		fileDir := strings.TrimPrefix(path.Dir(files[i].Path()), "/")
		diff := strings.TrimPrefix(strings.TrimPrefix(fileDir, parentDir), "/")

		var parents []string
		if diff != "" {
			parents = strings.Split(diff, "/")
		}

		if len(parents) == 0 {
			parent.AppendChild(files[i])
			continue
		}
		if !parent.IsDir() {
			// when parent is root
			parentDir = path.Dir(parent.Path())
		}

		dir := NewFile().Name(parents[0]).Dir().Path(path.Join(parentDir, parents[0])).Build()

		var childrenFiles []*File
		lo.ForEach(files, func(file *File, j int) {
			// a file whose path contains parent's path is folded under its parent
			if strings.HasPrefix(file.Path(), dir.Path()) {
				childrenFiles = append(childrenFiles, file)
				_, index, _ := lo.FindIndexOf(files, func(f *File) bool {
					return f == file
				})
				// a file which is folded here will be skipped above for loop
				skipIndexes = append(skipIndexes, index)
			}
		})

		foldedDir := FoldFiles(childrenFiles, dir)
		parent.AppendChild(foldedDir)
	}

	return parent
}
