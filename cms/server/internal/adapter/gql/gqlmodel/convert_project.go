package gqlmodel

import (
	"github.com/reearth/reearth-cms/server/pkg/project"
)

func ToProject(p *project.Project) *Project {
	if p == nil {
		return nil
	}

	return &Project{
		ID:          IDFrom(p.ID()),
		WorkspaceID: IDFrom(p.Workspace()),
		CreatedAt:   p.CreatedAt(),
		Alias:       p.Alias(),
		Name:        p.Name(),
		Description: p.Description(),
		UpdatedAt:   p.UpdatedAt(),
		Publication: ToProjectPublication(p.Publication()),
	}
}

func ToProjectPublication(p *project.Publication) *ProjectPublication {
	if p == nil {
		return nil
	}

	return &ProjectPublication{
		Scope:       ToProjectPublicationScope(p.Scope()),
		AssetPublic: p.AssetPublic(),
	}
}

func ToProjectPublicationScope(p project.PublicationScope) ProjectPublicationScope {
	switch p {
	case project.PublicationScopePublic:
		return ProjectPublicationScopePublic
	case project.PublicationScopeLimited:
		return ProjectPublicationScopeLimited
	}
	return ProjectPublicationScopePrivate
}

func FromProjectPublicationScope(p ProjectPublicationScope) project.PublicationScope {
	switch p {
	case ProjectPublicationScopePublic:
		return project.PublicationScopePublic
	case ProjectPublicationScopeLimited:
		return project.PublicationScopeLimited
	}
	return project.PublicationScopePrivate
}
