package thread

import (
	"time"

	"github.com/reearth/reearth-cms/server/pkg/operator"
)

type Comment struct {
	id      CommentID
	author  operator.Operator
	content string
}

func NewComment(id CommentID, author operator.Operator, content string) *Comment {
	return &Comment{
		id:      id,
		author:  author,
		content: content,
	}
}

func (c *Comment) ID() CommentID {
	return c.id
}

func (c *Comment) Author() operator.Operator {
	return c.author
}

func (c *Comment) Content() string {
	return c.content
}

func (c *Comment) CreatedAt() time.Time {
	return c.id.Timestamp()
}

func (c *Comment) SetContent(content string) {
	c.content = content
}

func (c *Comment) Clone() *Comment {
	if c == nil {
		return nil
	}

	return &Comment{
		id:      c.id,
		author:  c.author,
		content: c.content,
	}
}
