package indexer

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = &Config{
	IdProperty: "gml_id",
	Indexes: map[string]Index{
		"名称":           {Kind: "enum"},
		"用途":           {Kind: "enum"},
		"住所":           {Kind: "enum"},
		"建物利用現況（大分類）":  {Kind: "enum"},
		"建物利用現況（中分類）":  {Kind: "enum"},
		"建物利用現況（小分類）":  {Kind: "enum"},
		"建物利用現況（詳細分類）": {Kind: "enum"},
		"構造種別":         {Kind: "enum"},
		"構造種別（自治体独自）":  {Kind: "enum"},
		"耐火構造種別":       {Kind: "enum"},
	},
}

func TestIndexer(t *testing.T) {
	dir, err := os.Stat("testdata")
	if err != nil || !dir.IsDir() {
		t.Skip()
		return
	}

	b := bytes.NewBuffer(nil)
	zw := zip.NewWriter(b)

	input := NewFSFS(os.DirFS("testdata"))
	output := NewZipOutputFS(zw, "")
	indexer := NewIndexer(config, input, output, true)

	err = indexer.BuildAndWrite()
	assert.NoError(t, err)
	err = zw.Close()
	assert.NoError(t, err)

	br := bytes.NewReader(b.Bytes())
	zr, err := zip.NewReader(br, int64(b.Len()))
	assert.NoError(t, err)
	err = extractZip(zr, filepath.Join("testdata", "result"))
	assert.NoError(t, err)
}

func TestIndexerWithHTTPFS(t *testing.T) {
	u, _ := os.LookupEnv("REEARTH_PLATEAUVIEW_TESTURL")
	if u == "" {
		t.Skip()
		return
	}

	t.Log(u)
	b := bytes.NewBuffer(nil)
	zw := zip.NewWriter(b)

	input := NewHTTPFS(nil, u)
	output := NewZipOutputFS(zw, "")
	indexer := NewIndexer(config, input, output, true)

	err := indexer.BuildAndWrite()
	assert.NoError(t, err)
	err = zw.Close()
	assert.NoError(t, err)

	br := bytes.NewReader(b.Bytes())
	zr, err := zip.NewReader(br, int64(b.Len()))
	assert.NoError(t, err)
	err = extractZip(zr, filepath.Join("testdata", "result"))
	assert.NoError(t, err)
}

func extractZip(zr *zip.Reader, base string) error {
	_ = os.MkdirAll(base, os.ModePerm)
	for _, f := range zr.File {
		f := func() error {
			r, err := f.Open()
			if err != nil {
				return err
			}
			defer r.Close()
			f, err := os.Create(filepath.Join(base, f.Name))
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(f, r)
			return err
		}
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
