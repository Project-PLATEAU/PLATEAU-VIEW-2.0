package gqlmodel

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/request"
	"github.com/reearth/reearth-cms/server/pkg/version"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestToRequest(t *testing.T) {
	itm, _ := request.NewItem(id.NewItemID())
	req := request.New().
		NewID().
		Project(id.NewProjectID()).
		Workspace(id.NewWorkspaceID()).
		Items(request.ItemList{itm}).
		Title("foo").
		Description("xxx").
		State(request.StateClosed).
		Thread(id.NewThreadID()).
		Reviewers(id.UserIDList{id.NewUserID()}).
		CreatedBy(id.NewUserID()).
		ClosedAt(lo.ToPtr(util.Now())).
		ApprovedAt(lo.ToPtr(util.Now())).
		UpdatedAt(util.Now()).
		MustBuild()
	assert.Equal(t, &Request{
		ID: IDFrom(req.ID()),
		Items: []*RequestItem{{
			ItemID: IDFrom(itm.Item()),
			Ref:    lo.ToPtr(version.Public.String()),
		}},
		Title:       "foo",
		Description: lo.ToPtr("xxx"),
		CreatedByID: IDFrom(req.CreatedBy()),
		WorkspaceID: IDFrom(req.Workspace()),
		ProjectID:   IDFrom(req.Project()),
		ThreadID:    IDFrom(req.Thread()),
		ReviewersID: []ID{IDFrom(req.Reviewers()[0])},
		State:       RequestStateClosed,
		CreatedAt:   req.CreatedAt(),
		UpdatedAt:   req.UpdatedAt(),
		ApprovedAt:  req.ApprovedAt(),
		ClosedAt:    req.ClosedAt(),
	}, ToRequest(req))
}

func TestToRequestState(t *testing.T) {
	assert.Equal(t, RequestStateClosed, ToRequestState(request.StateClosed))
	assert.Equal(t, RequestStateWaiting, ToRequestState(request.StateWaiting))
	assert.Equal(t, RequestStateApproved, ToRequestState(request.StateApproved))
	assert.Equal(t, RequestStateDraft, ToRequestState(request.StateDraft))
	assert.Equal(t, RequestState(""), ToRequestState("xxx"))
}
