package thread

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Thread struct {
	id        ID
	workspace WorkspaceID
	comments  []*Comment
}

func (th *Thread) ID() ID {
	return th.id
}

func (th *Thread) Workspace() WorkspaceID {
	return th.workspace
}

func (th *Thread) Comments() []*Comment {
	if th == nil {
		return nil
	}
	return slices.Clone(th.comments)
}

func (th *Thread) HasComment(cid CommentID) bool {
	if th == nil {
		return false
	}
	return lo.SomeBy(th.comments, func(c *Comment) bool { return c.ID() == cid })
}

func (th *Thread) AddComment(c *Comment) error {
	if th.comments == nil {
		th.comments = []*Comment{}
	}
	if th.HasComment(c.ID()) {
		return ErrCommentAlreadyExist
	}

	th.comments = append(th.comments, c)
	return nil
}

func (th *Thread) UpdateComment(cid id.CommentID, content string) error {
	c, _ := lo.Find(th.comments, func(c *Comment) bool { return c.ID() == cid })
	if c == nil {
		return ErrCommentDoesNotExist
	}
	c.SetContent(content)
	return nil
}

func (th *Thread) DeleteComment(cid id.CommentID) error {
	i := slices.IndexFunc(th.Comments(), func(c *Comment) bool { return c.ID() == cid })
	if i < 0 {
		return ErrCommentDoesNotExist
	}

	comments := append(th.Comments()[:i], th.Comments()[i+1:]...)
	th.SetComments(comments...)
	return nil
}

func (th *Thread) Comment(cid id.CommentID) *Comment {
	c, _ := lo.Find(th.comments, func(c *Comment) bool { return c.ID() == cid })
	return c
}

func (th *Thread) SetComments(comments ...*Comment) {
	th.comments = slices.Clone(comments)
}

func (th *Thread) Clone() *Thread {
	if th == nil {
		return nil
	}

	comments := util.Map(th.comments, func(c *Comment) *Comment {
		return c.Clone()
	})

	return &Thread{
		id:        th.id.Clone(),
		workspace: th.workspace.Clone(),
		comments:  comments,
	}
}
