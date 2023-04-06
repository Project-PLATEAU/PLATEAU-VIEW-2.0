package fs

import (
	"context"
	"io"
	"net/url"
	"os"
	"path"

	"github.com/kennygrant/sanitize"
	"github.com/reearth/reearth-cms/worker/internal/usecase/gateway"
	"github.com/reearth/reearthx/rerror"
	"github.com/spf13/afero"
)

type fileRepo struct {
	fs      afero.Fs
	urlBase *url.URL
}

func NewFile(fs afero.Fs, urlBase string) (gateway.File, error) {
	var b *url.URL
	var err error
	b, err = url.Parse(urlBase)
	if err != nil {
		return nil, invalidBaseURLErr
	}

	return &fileRepo{
		fs:      fs,
		urlBase: b,
	}, nil
}

// Read implements gateway.File
func (f *fileRepo) Read(ctx context.Context, path string) (gateway.ReadAtCloser, int64, error) {
	if path == "" {
		return nil, 0, rerror.ErrNotFound
	}

	file, err := f.fs.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, 0, rerror.ErrNotFound
		}
		return nil, 0, rerror.ErrInternalBy(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, 0, rerror.ErrInternalBy(err)
	}
	return file, fileInfo.Size(), nil
}

// Upload implements gateway.File
func (f *fileRepo) Upload(ctx context.Context, name string) (io.WriteCloser, error) {
	if name == "" {
		return nil, gateway.ErrFailedToUploadFile
	}

	sanitizedName := sanitize.Path(name)

	if fnd := path.Dir(sanitizedName); fnd != "" {
		if err := f.fs.MkdirAll(fnd, 0755); err != nil {
			return nil, rerror.ErrInternalBy(err)
		}
	}

	dest, err := f.fs.Create(sanitizedName)
	if err != nil {
		return nil, rerror.ErrInternalBy(err)
	}
	return dest, nil
}
