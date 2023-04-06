package gateway

import (
	"context"

	"github.com/reearth/reearth-cms/worker/pkg/asset"
)

type CMS interface {
	NotifyAssetDecompressed(ctx context.Context, assetID string, status *asset.ArchiveExtractionStatus) error
}
