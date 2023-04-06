package interactor

import (
	"github.com/reearth/reearth-cms/worker/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/worker/internal/usecase/repo"
)

type Usecase struct {
	gateways *gateway.Container
	webhook  repo.Webhook
}

func NewUsecase(g *gateway.Container, webhook repo.Webhook) *Usecase {
	return &Usecase{gateways: g, webhook: webhook}
}
