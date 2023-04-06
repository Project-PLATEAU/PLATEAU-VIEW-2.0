package indexer

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/reearth/reearthx/log"
)

type FS interface {
	Open(name string) (io.ReadCloser, error)
}

type OutputFS interface {
	Open(name string) (WriteCloser, error)
}

type WriteCloser interface {
	io.Writer
	io.Closer
}

type NopCloser struct {
	w io.Writer
}

func NewNopCloser(w io.Writer) *NopCloser {
	if w == nil {
		return nil
	}
	return &NopCloser{w: w}
}

func (n *NopCloser) Write(p []byte) (int, error) {
	return n.w.Write(p)
}

func (n *NopCloser) Close() error {
	return nil
}

type FSFS struct {
	fs fs.FS
}

func NewFSFS(f fs.FS) *FSFS {
	return &FSFS{fs: f}
}

func (f *FSFS) Open(name string) (io.ReadCloser, error) {
	file, err := f.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

type ZipFS struct {
	z *zip.Reader
}

func NewZipFS(z *zip.Reader) *ZipFS {
	return &ZipFS{z: z}
}

func (f *ZipFS) Open(name string) (io.ReadCloser, error) {
	file, err := f.z.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(buf), nil
}

type OSOutputFS struct {
	base string
}

func NewOSOutputFS(base string) *OSOutputFS {
	return &OSOutputFS{base: base}
}

func (f *OSOutputFS) Open(name string) (w WriteCloser, err error) {
	if err := f.mkdir(); err != nil {
		return nil, err
	}
	return os.Create(filepath.Join(f.base, name))
}

func (f *OSOutputFS) mkdir() error {
	if f.base != "" {
		return os.MkdirAll(f.base, os.ModePerm)
	}
	return nil
}

type ZipOutputFS struct {
	base string
	w    *zip.Writer
}

func NewZipOutputFS(w *zip.Writer, base string) *ZipOutputFS {
	return &ZipOutputFS{base: base, w: w}
}

func (f *ZipOutputFS) Open(name string) (WriteCloser, error) {
	w, err := f.w.Create(path.Join(f.base, name))
	return NewNopCloser(w), err
}

type HTTPFS struct {
	c    *http.Client
	base string
}

func NewHTTPFS(c *http.Client, base string) *HTTPFS {
	if c == nil {
		c = http.DefaultClient
	}
	return &HTTPFS{c: c, base: base}
}

func (f *HTTPFS) Open(name string) (io.ReadCloser, error) {
	u, err := url.JoinPath(f.base, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get url from %s and %s: %w", f.base, name, err)
	}

	log.Debugf("indexer: http get: %s", u)
	res, err := f.c.Get(u)
	if err != nil {
		return nil, fmt.Errorf("failed to get from %s: %w", u, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is %d", res.StatusCode)
	}

	// It's storage but some HTTP requests for b3dm data fails without this line
	res2, err2 := f.c.Get(u)
	if err2 != nil {
		_ = res.Body.Close()
		return nil, fmt.Errorf("failed to get from %s: %w", u, err2)
	}
	defer func() {
		_ = res2.Body.Close()
	}()

	return res.Body, nil
}
