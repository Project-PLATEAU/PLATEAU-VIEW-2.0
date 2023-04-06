package interactor

import (
	"context"
	"errors"
	"io"
	"path"
	"strings"

	"github.com/reearth/reearth-cms/worker/pkg/asset"
	"github.com/reearth/reearth-cms/worker/pkg/decompressor"
	"github.com/reearth/reearthx/log"
	"github.com/samber/lo"
)

func (u *Usecase) Decompress(ctx context.Context, assetID, assetPath string) error {
	err := u.decompress(ctx, assetID, assetPath)
	if err != nil {
		log.Errorf("failed to notify to CMS, Asset=%s, Path=%s", assetID, assetPath)
		return u.gateways.CMS.NotifyAssetDecompressed(ctx, assetID, lo.ToPtr(asset.ArchiveExtractionStatusFailed))
	}
	return u.gateways.CMS.NotifyAssetDecompressed(ctx, assetID, lo.ToPtr(asset.ArchiveExtractionStatusDone))
}

func (u *Usecase) decompress(ctx context.Context, assetID, assetPath string) error {
	ext := strings.TrimPrefix(path.Ext(assetPath), ".")
	base := strings.TrimPrefix(strings.TrimSuffix(assetPath, "."+ext), "/")

	compressedFile, size, err := u.gateways.File.Read(ctx, assetPath)
	if err != nil {
		log.Errorf("failed to load zip file from storage, Asset=%s, Path=%s, Err=%s", assetID, assetPath, err.Error())
		return err
	}

	uploadFunc := func(name string) (io.WriteCloser, error) {
		w, err := u.gateways.File.Upload(ctx, path.Join(base, name))
		if err != nil {
			return nil, err
		}
		return w, nil
	}

	de, err := decompressor.New(compressedFile, size, ext, uploadFunc)
	if err != nil {
		if errors.Is(err, decompressor.ErrUnsupportedExtention) {
			log.Infof("unsupported extension: decompression skipped assetID=%s ext=%s", assetID, ext)
			return nil
		}
		return err
	}

	if err = de.Decompress(base); err != nil {
		return err
	}

	return nil
}
