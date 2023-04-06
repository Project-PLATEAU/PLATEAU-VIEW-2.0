package indexer

import (
	"archive/zip"
	"bytes"
	"io"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestZipOutputFS(t *testing.T) {
	b := bytes.NewBuffer(nil)
	zw := zip.NewWriter(b)
	zofs := NewZipOutputFS(zw, "")

	f, err := zofs.Open("test.txt")
	assert.NoError(t, err)
	_, err = f.Write([]byte("hello world!"))
	assert.NoError(t, err)
	err = f.Close()
	assert.NoError(t, err)

	err = zw.Close()
	assert.NoError(t, err)

	br := bytes.NewReader(b.Bytes())
	zr, err := zip.NewReader(br, int64(b.Len()))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(zr.File))
	assert.Equal(t, "test.txt", zr.File[0].Name)
	zf, err := zr.File[0].Open()
	assert.NoError(t, err)
	assert.Equal(t, "hello world!", string(lo.Must(io.ReadAll(zf))))
	assert.NoError(t, zf.Close())
}
