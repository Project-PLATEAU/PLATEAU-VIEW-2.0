package mongodoc

import (
	"net/url"
	"time"

	"github.com/reearth/reearth/server/pkg/id"
	"github.com/reearth/reearth/server/pkg/project"
	"github.com/reearth/reearth/server/pkg/visualizer"
	"github.com/reearth/reearthx/mongox"
)

type ProjectDocument struct {
	ID                string
	Archived          bool
	IsBasicAuthActive bool
	BasicAuthUsername string
	BasicAuthPassword string
	UpdatedAt         time.Time
	PublishedAt       time.Time
	Name              string
	Description       string
	Alias             string
	ImageURL          string
	PublicTitle       string
	PublicDescription string
	PublicImage       string
	PublicNoIndex     bool
	Team              string // DON'T CHANGE NAME'
	Visualizer        string
	PublishmentStatus string
}

type ProjectConsumer = mongox.SliceFuncConsumer[*ProjectDocument, *project.Project]

func NewProjectConsumer() *ProjectConsumer {
	return NewComsumer[*ProjectDocument, *project.Project]()
}

func NewProject(project *project.Project) (*ProjectDocument, string) {
	pid := project.ID().String()

	imageURL := ""
	if u := project.ImageURL(); u != nil {
		imageURL = u.String()
	}

	return &ProjectDocument{
		ID:                pid,
		Archived:          project.IsArchived(),
		IsBasicAuthActive: project.IsBasicAuthActive(),
		BasicAuthUsername: project.BasicAuthUsername(),
		BasicAuthPassword: project.BasicAuthPassword(),
		UpdatedAt:         project.UpdatedAt(),
		PublishedAt:       project.PublishedAt(),
		Name:              project.Name(),
		Description:       project.Description(),
		Alias:             project.Alias(),
		ImageURL:          imageURL,
		PublicTitle:       project.PublicTitle(),
		PublicDescription: project.PublicDescription(),
		PublicImage:       project.PublicImage(),
		PublicNoIndex:     project.PublicNoIndex(),
		Team:              project.Workspace().String(),
		Visualizer:        string(project.Visualizer()),
		PublishmentStatus: string(project.PublishmentStatus()),
	}, pid
}

func (d *ProjectDocument) Model() (*project.Project, error) {
	pid, err := id.ProjectIDFrom(d.ID)
	if err != nil {
		return nil, err
	}
	tid, err := id.WorkspaceIDFrom(d.Team)
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
		IsArchived(d.Archived).
		IsBasicAuthActive(d.IsBasicAuthActive).
		BasicAuthUsername(d.BasicAuthUsername).
		BasicAuthPassword(d.BasicAuthPassword).
		UpdatedAt(d.UpdatedAt).
		PublishedAt(d.PublishedAt).
		Name(d.Name).
		Description(d.Description).
		Alias(d.Alias).
		ImageURL(imageURL).
		PublicTitle(d.PublicTitle).
		PublicDescription(d.PublicDescription).
		PublicImage(d.PublicImage).
		PublicNoIndex(d.PublicNoIndex).
		Workspace(tid).
		Visualizer(visualizer.Visualizer(d.Visualizer)).
		PublishmentStatus(project.PublishmentStatus(d.PublishmentStatus)).
		Build()
}
