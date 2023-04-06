package fs

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/kennygrant/sanitize"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/file"
	"github.com/reearth/reearthx/rerror"
	"github.com/spf13/afero"
)

type fileRepo struct {
	fs      afero.Fs
	urlBase *url.URL
}

func NewFile(fs afero.Fs, urlBase string) (gateway.File, error) {
	var b *url.URL
	if urlBase == "" {
		urlBase = defaultBase
	}

	var err error
	b, err = url.Parse(urlBase)
	if err != nil {
		return nil, ErrInvalidBaseURL
	}

	return &fileRepo{
		fs:      fs,
		urlBase: b,
	}, nil
}

func (f *fileRepo) ReadAsset(ctx context.Context, u string, fn string) (io.ReadCloser, error) {
	if u == "" || fn == "" {
		return nil, rerror.ErrNotFound
	}

	p := getFSObjectPath(u, fn)
	sn := sanitize.Path(p)

	return f.read(ctx, sn)
}

func (f *fileRepo) GetAssetFiles(ctx context.Context, u string) ([]gateway.FileEntry, error) {
	if u == "" {
		return nil, rerror.ErrNotFound
	}

	p := getFSObjectPath(u, "")
	var fileEntries []gateway.FileEntry
	err := afero.Walk(f.fs, p, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileEntries = append(fileEntries, gateway.FileEntry{
			Name: strings.TrimPrefix(strings.TrimPrefix(path, p), "/"),
			Size: info.Size(),
		})
		return nil
	})
	if err != nil {
		if errors.Is(err, afero.ErrFileNotFound) {
			return nil, gateway.ErrFileNotFound
		} else {
			return nil, rerror.ErrInternalBy(err)
		}
	}

	if len(fileEntries) == 0 {
		return nil, gateway.ErrFileNotFound
	}

	return fileEntries, nil
}

func (f *fileRepo) UploadAsset(ctx context.Context, file *file.File) (string, int64, error) {
	if file == nil {
		return "", 0, gateway.ErrInvalidFile
	}
	if file.Size >= fileSizeLimit {
		return "", 0, gateway.ErrFileTooLarge
	}

	uuid := newUUID()

	p := getFSObjectPath(uuid, file.Path)

	if err := f.upload(ctx, p, file.Content); err != nil {
		return "", 0, err
	}

	return uuid, file.Size, nil
}

func (f *fileRepo) DeleteAsset(ctx context.Context, u string, fn string) error {
	if u == "" || fn == "" {
		return gateway.ErrInvalidFile
	}

	p := getFSObjectPath(u, fn)
	sn := sanitize.Path(p)

	return f.delete(ctx, sn)
}

func (f *fileRepo) GetURL(a *asset.Asset) string {
	uuid := a.UUID()
	return f.urlBase.JoinPath(assetDir, uuid[:2], uuid[2:], url.PathEscape(a.FileName())).String()
}

// helpers

func (f *fileRepo) read(ctx context.Context, filename string) (io.ReadCloser, error) {
	if filename == "" {
		return nil, rerror.ErrNotFound
	}

	file, err := f.fs.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, rerror.ErrNotFound
		}
		return nil, rerror.ErrInternalBy(err)
	}
	return file, nil
}

func (f *fileRepo) upload(ctx context.Context, filename string, content io.Reader) error {
	if filename == "" || content == nil {
		return gateway.ErrFailedToUploadFile
	}

	if fnd := path.Dir(filename); fnd != "" {
		if err := f.fs.MkdirAll(fnd, 0755); err != nil {
			return rerror.ErrInternalBy(err)
		}
	}

	dest, err := f.fs.Create(filename)
	if err != nil {
		return rerror.ErrInternalBy(err)
	}
	defer func() {
		_ = dest.Close()
	}()

	if _, err := io.Copy(dest, content); err != nil {
		return gateway.ErrFailedToUploadFile
	}

	return nil
}

func (f *fileRepo) delete(ctx context.Context, filename string) error {
	if filename == "" {
		return gateway.ErrFailedToUploadFile
	}

	if err := f.fs.RemoveAll(filename); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return rerror.ErrInternalBy(err)
	}
	return nil
}

func getFSObjectPath(uuid, objectName string) string {
	if uuid == "" || !IsValidUUID(uuid) {
		return ""
	}

	p := path.Join(assetDir, uuid[:2], uuid[2:], objectName)
	return sanitize.Path(p)
}

func newUUID() string {
	return uuid.New().String()
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
