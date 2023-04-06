package memory

import (
	"time"

	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearthx/usecasex"
)

func New() *repo.Container {
	return &repo.Container{
		Asset:       NewAsset(),
		AssetFile:   NewAssetFile(),
		Lock:        NewLock(),
		User:        NewUser(),
		Request:     NewRequest(),
		Workspace:   NewWorkspace(),
		Project:     NewProject(),
		Model:       NewModel(),
		Item:        NewItem(),
		Schema:      NewSchema(),
		Integration: NewIntegration(),
		Thread:      NewThread(),
		Event:       NewEvent(),
		Transaction: &usecasex.NopTransaction{},
	}
}

func MockNow(r *repo.Container, t time.Time) func() {
	p := r.Project.(*Project).now.Mock(t)

	return func() {
		p()
	}
}
