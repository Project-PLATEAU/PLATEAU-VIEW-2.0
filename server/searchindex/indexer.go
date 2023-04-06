package searchindex

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/dustin/go-humanize"
	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/eukarya-inc/reearth-plateauview/server/searchindex/indexer"
	"github.com/reearth/reearthx/log"
)

var builtinConfig = &indexer.Config{
	IdProperty: "gml_id",
	Indexes: map[string]indexer.Index{
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

type Indexer struct {
	base   *url.URL
	config *indexer.Config
	cms    cms.Interface
	pid    string
	debug  bool
	// true -> faster but uses more memory
	zipMode bool
	// true -> more stable but uses more memory
	bufferMode bool
}

func NewIndexer(cms cms.Interface, pid string, base *url.URL, debug bool) *Indexer {
	return &Indexer{
		base:       base,
		config:     builtinConfig,
		cms:        cms,
		pid:        pid,
		debug:      debug,
		zipMode:    false,
		bufferMode: false,
	}
}

func NewZipIndexer(cms cms.Interface, pid string, base *url.URL, debug bool) *Indexer {
	i := NewIndexer(cms, pid, base, debug)
	i.zipMode = true
	return i
}

func (i *Indexer) BuildIndex(ctx context.Context, name string) (string, error) {
	indfs, err := i.fs()
	if err != nil {
		return "", fmt.Errorf("インデックスを作成できませんでした。%w", err)
	}

	ind := indexer.NewIndexer(builtinConfig, indfs, nil, i.debug)
	res, err := ind.Build()
	if err != nil {
		return "", fmt.Errorf("インデックスを作成できませんでした。%w", err)
	}

	log.Infof("indexer webhook: suceeded to build indexes for %s", name)

	if i.bufferMode {
		return i.uploadWithBuffer(ctx, name, res)
	}
	return i.uploadWithPipe(ctx, name, res)
}

func (i *Indexer) uploadWithPipe(ctx context.Context, name string, res indexer.Result) (string, error) {
	pr, pw := io.Pipe()

	aids := make(chan string)
	errs := make(chan error)
	go func() {
		aid, err := i.cms.UploadAssetDirectly(ctx, i.pid, fmt.Sprintf("%s_index.zip", name), pr)
		aids <- aid
		errs <- err
	}()

	zw := zip.NewWriter(pw)
	zfs := indexer.NewZipOutputFS(zw, "")
	if err := indexer.NewWriter(i.config, zfs).Write(res); err != nil {
		return "", fmt.Errorf("failed to save files to zip: %w", err)
	}

	if err := zw.Close(); err != nil {
		return "", fmt.Errorf("failed to close zip: %w", err)
	}

	if err := pw.Close(); err != nil {
		return "", fmt.Errorf("結果のアップロードに失敗しました。(2) %w", err)
	}

	log.Debugf("indexer webhook: waiting for finishing asset upload of indexes for %s", name)

	aid := <-aids
	err := <-errs
	if err != nil {
		return "", fmt.Errorf("結果のアップロードに失敗しました。(3) %w", err)
	}

	log.Debugf("indexer webhook: succeeded to zip indexes for %s", name)
	return aid, nil
}

func (i *Indexer) uploadWithBuffer(ctx context.Context, name string, res indexer.Result) (string, error) {
	b := &bytes.Buffer{}

	zw := zip.NewWriter(b)
	zfs := indexer.NewZipOutputFS(zw, "")
	if err := indexer.NewWriter(i.config, zfs).Write(res); err != nil {
		return "", fmt.Errorf("failed to save files to zip: %w", err)
	}

	if err := zw.Close(); err != nil {
		return "", fmt.Errorf("failed to close zip: %w", err)
	}

	log.Debugf("indexer webhook: succeeded to zip indexes for %s", name)

	aid, err := i.cms.UploadAssetDirectly(ctx, i.pid, fmt.Sprintf("%s_index.zip", name), b)
	if err != nil {
		return "", fmt.Errorf("結果のアップロードに失敗しました。(3) %w", err)
	}

	log.Debugf("indexer webhook: succeeded to upload indexes for %s", name)
	return aid, nil
}

func (i *Indexer) fs() (indexer.FS, error) {
	if i.zipMode {
		u := i.base.String()
		log.Infof("indexer webhook: zip indexer donwloads %s", u)
		res, err := http.DefaultClient.Get(u)
		if err != nil {
			return nil, fmt.Errorf("3D TilesのZipファイルのダウンロードに失敗しました。%w", err)
		}

		defer func() {
			_ = res.Body.Close()
		}()

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("3D TilesのZipファイルのダウンロードに失敗しました。ステータスコードが%dでした。", res.StatusCode)
		}

		b := bytes.NewBuffer(nil)
		size, err := io.Copy(b, res.Body)
		if err != nil {
			return nil, fmt.Errorf("3D TilesのZipファイルのダウンロードに失敗しました。%w", err)
		}

		log.Infof("indexer webhook: zip indexer donwloaded %s from %s", humanize.Bytes(uint64(size)), u)

		z, err := zip.NewReader(bytes.NewReader(b.Bytes()), size)
		if err != nil {
			return nil, fmt.Errorf("3D TilesのZipファイルが不正なフォーマットです。%w", err)
		}

		return indexer.NewZipFS(z), nil
	}

	return indexer.NewHTTPFS(nil, getAssetBase(i.base)), nil
}

type OutputFS struct {
	c   cms.Interface
	cb  func(assetID string, err error)
	ctx context.Context
	pid string
}

func NewOutputFS(ctx context.Context, c cms.Interface, projectID string, cb func(assetID string, err error)) *OutputFS {
	return &OutputFS{
		c:   c,
		cb:  cb,
		ctx: ctx,
		pid: projectID,
	}
}

func (f *OutputFS) Open(p string) (indexer.WriteCloser, error) {
	pr, pw := io.Pipe()

	go func() {
		f.cb(f.c.UploadAssetDirectly(f.ctx, f.pid, path.Base(p), pr))
	}()

	return pw, nil
}
