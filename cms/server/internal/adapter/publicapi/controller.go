package publicapi

import (
	"context"
	"errors"

	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearthx/i18n"
	"github.com/reearth/reearthx/rerror"
)

var ErrInvalidProject = rerror.NewE(i18n.T("invalid project"))

type Controller struct {
	project          repo.Project
	usecases         *interfaces.Container
	assetUrlResolver asset.URLResolver
}

func NewController(project repo.Project, usecases *interfaces.Container, aur asset.URLResolver) *Controller {
	return &Controller{
		project:          project,
		usecases:         usecases,
		assetUrlResolver: aur,
	}
}

func (c *Controller) checkProject(ctx context.Context, prj string) (*project.Project, error) {
	pr, err := c.project.FindByPublicName(ctx, prj)
	if err != nil {
		if errors.Is(err, rerror.ErrNotFound) {
			return nil, rerror.ErrNotFound
		}
		return nil, ErrInvalidProject
	}

	if p := pr.Publication(); p == nil || p.Scope() != project.PublicationScopePublic {
		return nil, rerror.ErrNotFound
	}

	// TODO: check token if the scope is limited
	// if pr.Publication().Scope() == project.PublicationScopeLimited {
	// 	t := pr.Publication().Token()
	// 	if op := adapter.Operator(ctx); op == nil || t == "" || op.PublicAPIToken != t {
	// 		return nil, ErrInvalidProject
	// 	}
	// }

	return pr, nil
}
