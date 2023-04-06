package gcp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/reearth/reearth-cms/worker/internal/usecase/gateway"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
)

const (
	gcsAssetBasePath string = "assets"
	cacheSize               = 5 * 1024 * 1024 * 1024 // 5GB
)

type fileRepo struct {
	bucketName   string
	base         *url.URL
	cacheControl string
}

type Buffer struct {
	b bytes.Reader
}

func NewFile(bucketName, base string, cacheControl string) (gateway.File, error) {
	if bucketName == "" {
		return nil, errors.New("bucket name is empty")
	}

	var u *url.URL
	if base == "" {
		base = fmt.Sprintf("https://storage.googleapis.com/%s", bucketName)
	}

	var err error
	u, _ = url.Parse(base)
	if err != nil {
		return nil, errors.New("invalid base URL")
	}

	return &fileRepo{
		bucketName:   bucketName,
		base:         u,
		cacheControl: cacheControl,
	}, nil
}

func (f *fileRepo) Read(ctx context.Context, path string) (gateway.ReadAtCloser, int64, error) {
	if path == "" {
		return nil, 0, rerror.ErrNotFound
	}
	objectName := getGCSObjectNameFromURL(gcsAssetBasePath, path)
	// return f.NewGCSReaderAt(ctx, objectName)
	return f.readAll(ctx, objectName)
}

// Upload is the function which allows this func's user to generate the function to upload asset to GCS dynamically
func (f *fileRepo) Upload(ctx context.Context, name string) (io.WriteCloser, error) {
	if name == "" {
		return nil, gateway.ErrInvalidFile
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: upload bucket err: %+v\n", err)
		return nil, rerror.ErrInternalBy(err)
	}

	name = filepath.Join(gcsAssetBasePath, name)

	object := bucket.Object(name)
	if err := object.Delete(ctx); err != nil && !errors.Is(err, storage.ErrObjectNotExist) {
		log.Errorf("gcs: upload delete err: %+v\n", err)
		return nil, gateway.ErrFailedToUploadFile
	}

	writer := object.NewWriter(ctx)
	writer.ObjectAttrs.CacheControl = f.cacheControl
	return writer, nil
}

// GCSReaderAt is a struct which implements io.ReadAt interface
func (f *fileRepo) NewGCSReaderAt(ctx context.Context, objectName string) (gateway.ReadAtCloser, int64, error) {
	rowReaderAt, size, err := f.newRawGCSReaderAt(ctx, objectName)
	if err != nil {
		log.Errorf("gcs: rawGCSReaderAt err: ObjectName=%s, err=%+v\n", objectName, err)
		return nil, 0, rerror.ErrInternalBy(err)
	}
	return rowReaderAt, size, nil
}

type rawGCSReaderAt struct {
	ctx context.Context
	obj *storage.ObjectHandle
}

func (f *fileRepo) readAll(ctx context.Context, objectName string) (gateway.ReadAtCloser, int64, error) {
	if objectName == "" {
		return nil, 0, rerror.ErrNotFound
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: read bucket err: %+v\n", err)
		return nil, 0, rerror.ErrInternalBy(err)
	}

	obj := bucket.Object(objectName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, 0, err
	}

	// read all data on memory
	objectData, err := io.ReadAll(r)
	if err != nil {
		return nil, 0, err
	}

	reader := bytes.NewReader(objectData)
	bufReader := Buffer{
		*reader,
	}

	return &bufReader, int64(len(objectData)), nil
}

func (b *Buffer) Close() error {
	return nil
}

func (b *Buffer) ReadAt(b2 []byte, off int64) (n int, err error) {
	return b.b.ReadAt(b2, off)
}

// newRawGCSReaderAt implements io.ReadAt
func (f *fileRepo) newRawGCSReaderAt(ctx context.Context, objectName string) (gateway.ReadAtCloser, int64, error) {
	if objectName == "" {
		return nil, 0, rerror.ErrNotFound
	}

	bucket, err := f.bucket(ctx)
	if err != nil {
		log.Errorf("gcs: read bucket err: %+v\n", err)
		return nil, 0, rerror.ErrInternalBy(err)
	}
	obj := bucket.Object(objectName)
	attr, err := obj.Attrs(ctx)
	if err != nil {
		return nil, 0, err
	}
	size := attr.Size
	return &rawGCSReaderAt{ctx, obj}, size, nil
}

func (g *rawGCSReaderAt) ReadAt(b []byte, off int64) (n int, err error) {
	rc, err := g.obj.NewRangeReader(g.ctx, off, int64(len(b)))
	if err != nil {
		return
	}
	defer rc.Close()

	return rc.Read(b)
}

func (g *rawGCSReaderAt) Close() error {
	return nil
}

// helpers
func (f *fileRepo) bucket(ctx context.Context) (*storage.BucketHandle, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	bucket := client.Bucket(f.bucketName)
	return bucket, nil
}

func getGCSObjectNameFromURL(assetBasePath string, path string) string {
	if path == "" {
		return ""
	}

	p := filepath.Join(assetBasePath, path)

	return p
}
