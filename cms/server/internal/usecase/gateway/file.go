package gateway

import (
	"context"
	"io"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/file"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var (
	ErrInvalidFile        error = rerror.NewE(i18n.T("invalid file"))
	ErrFailedToUploadFile error = rerror.NewE(i18n.T("failed to upload file"))
	ErrFileTooLarge       error = rerror.NewE(i18n.T("file too large"))
	ErrFailedToDeleteFile error = rerror.NewE(i18n.T("failed to delete file"))
	ErrFileNotFound       error = rerror.NewE(i18n.T("file not found"))
)

type FileEntry struct {
	Name string
	Size int64
}

type File interface {
	ReadAsset(context.Context, string, string) (io.ReadCloser, error)
	GetAssetFiles(context.Context, string) ([]FileEntry, error)
	UploadAsset(context.Context, *file.File) (string, int64, error)
	DeleteAsset(context.Context, string, string) error
	GetURL(*asset.Asset) string
}
