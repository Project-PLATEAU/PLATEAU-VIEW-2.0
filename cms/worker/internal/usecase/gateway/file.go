package gateway

import (
	"context"
	"errors"
	"io"
)

var (
	ErrInvalidFile        error = errors.New("invalid file")
	ErrNotFound           error = errors.New("not found")
	ErrFailedToUploadFile error = errors.New("failed to upload file")
	ErrFileTooLarge       error = errors.New("file too large")
	ErrFailedToRemoveFile error = errors.New("failed to remove file")
)

type ReadAtCloser interface {
	io.ReaderAt
	io.Closer
}

type File interface {
	Read(ctx context.Context, path string) (ReadAtCloser, int64, error)
	Upload(ctx context.Context, name string) (io.WriteCloser, error)
}
