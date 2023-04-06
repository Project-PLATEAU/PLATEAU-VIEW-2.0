package request

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestRequest_SetDescription(t *testing.T) {
	req := &Request{
		description: "xxx",
	}
	req.SetDescription("foo")
	assert.Equal(t, "foo", req.Description())
}

func TestRequest_SetItems(t *testing.T) {
	req := &Request{}
	i1, _ := NewItem(id.NewItemID())
	items1 := ItemList{i1}
	items2 := ItemList{i1, i1}
	err := req.SetItems(items1)
	assert.NoError(t, err)
	assert.Equal(t, items1, req.Items())
	err = req.SetItems(items2)
	assert.Same(t, ErrDuplicatedItem, err)
}

func TestRequest_SetReviewers(t *testing.T) {
	req := &Request{}
	reviewers := id.UserIDList{id.NewUserID()}
	req.SetReviewers(reviewers)
	assert.Equal(t, reviewers, req.Reviewers())
}

func TestRequest_SetState(t *testing.T) {
	req := &Request{
		description: "xxx",
	}
	req.SetDescription("foo")
	assert.Equal(t, "foo", req.Description())
}

func TestRequest_SetTitle(t *testing.T) {
	req := &Request{
		title: "xxx",
	}

	err := req.SetTitle("foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", req.Title())

	err = req.SetTitle("")
	assert.Equal(t, ErrEmptyTitle, err)

}

func TestRequest_SetState1(t *testing.T) {
	item, _ := NewItem(id.NewItemID())

	req1 := New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(id.NewProjectID()).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(ItemList{item}).
		Title("foo").
		MustBuild()

	req2 := New().
		NewID().
		Workspace(id.NewWorkspaceID()).
		Project(id.NewProjectID()).
		CreatedBy(id.NewUserID()).
		Thread(id.NewThreadID()).
		Items(ItemList{item}).
		Title("hoge").
		MustBuild()
	req1.SetState(StateClosed)
	assert.Equal(t, StateClosed, req1.State())
	assert.NotNil(t, req1.ClosedAt())

	req2.SetState(StateApproved)
	assert.Equal(t, StateApproved, req2.State())
	assert.NotNil(t, req2.ApprovedAt())
}
