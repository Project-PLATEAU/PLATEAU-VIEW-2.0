package http

import (
	"context"

	"github.com/reearth/reearth-cms/worker/internal/usecase/interactor"
)

type DecompressController struct {
	usecase *interactor.Usecase
}

func NewDecompressController(u *interactor.Usecase) *DecompressController {
	return &DecompressController{
		usecase: u,
	}
}

type DecompressInput struct {
	AssetID string `json:"assetId"`
	Path    string `json:"path"`
}

func (c *DecompressController) Decompress(ctx context.Context, input DecompressInput) error {
	return c.usecase.Decompress(ctx, input.AssetID, input.Path)
}
