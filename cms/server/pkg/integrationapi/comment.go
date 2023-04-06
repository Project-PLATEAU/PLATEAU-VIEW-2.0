package integrationapi

import (
	"github.com/reearth/reearth-cms/server/pkg/thread"
	"github.com/samber/lo"
)

func NewComment(c *thread.Comment) *Comment {
	if c == nil {
		return nil
	}

	var authorID any
	var authorType CommentAuthorType
	if c.Author().User() != nil {
		authorID = c.Author().User().Ref()
		authorType = User
	}
	if c.Author().Integration() != nil {
		authorID = c.Author().Integration().Ref()
		authorType = Integrtaion
	}
	return &Comment{
		Id:         c.ID().Ref(),
		AuthorId:   &authorID,
		AuthorType: &authorType,
		Content:    lo.ToPtr(c.Content()),
		CreatedAt:  lo.ToPtr(c.CreatedAt()),
	}
}
