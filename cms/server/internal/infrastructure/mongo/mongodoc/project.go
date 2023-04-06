package mongodoc

import (
	"net/url"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearthx/mongox"
)

type ProjectDocument struct {
	ID          string
	UpdatedAt   time.Time
	Name        string
	Description string
	Alias       string
	ImageURL    string
	Workspace   string
	Publication *ProjectPublicationDocument
}

type ProjectPublicationDocument struct {
	AssetPublic bool
	Scope       string
}

func NewProject(project *project.Project) (*ProjectDocument, string) {
	pid := project.ID().String()

	imageURL := ""
	if u := project.ImageURL(); u != nil {
		imageURL = u.String()
	}

	return &ProjectDocument{
		ID:          pid,
		UpdatedAt:   project.UpdatedAt(),
		Name:        project.Name(),
		Description: project.Description(),
		Alias:       project.Alias(),
		ImageURL:    imageURL,
		Workspace:   project.Workspace().String(),
		Publication: NewProjectPublication(project.Publication()),
	}, pid
}

func NewProjectPublication(p *project.Publication) *ProjectPublicationDocument {
	if p == nil {
		return nil
	}

	return &ProjectPublicationDocument{
		AssetPublic: p.AssetPublic(),
		Scope:       string(p.Scope()),
	}
}

func (d *ProjectDocument) Model() (*project.Project, error) {
	pid, err := id.ProjectIDFrom(d.ID)
	if err != nil {
		return nil, err
	}
	tid, err := id.WorkspaceIDFrom(d.Workspace)
	if err != nil {
		return nil, err
	}

	var imageURL *url.URL
	if d.ImageURL != "" {
		if imageURL, err = url.Parse(d.ImageURL); err != nil {
			imageURL = nil
		}
	}

	return project.New().
		ID(pid).
		UpdatedAt(d.UpdatedAt).
		Name(d.Name).
		Description(d.Description).
		Alias(d.Alias).
		Workspace(tid).
		ImageURL(imageURL).
		Publication(d.Publication.Model()).
		Build()
}

func (d *ProjectPublicationDocument) Model() *project.Publication {
	if d == nil {
		return nil
	}
	return project.NewPublication(project.PublicationScope(d.Scope), d.AssetPublic)
}

type ProjectConsumer = mongox.SliceFuncConsumer[*ProjectDocument, *project.Project]

func NewProjectConsumer() *ProjectConsumer {
	return NewComsumer[*ProjectDocument, *project.Project]()
}
