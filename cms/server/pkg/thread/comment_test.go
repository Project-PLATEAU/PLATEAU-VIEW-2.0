package thread

import (
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/stretchr/testify/assert"
)

func TestComment_CommentType(t *testing.T) {
	cid := NewCommentID()
	uid := NewUserID()
	c := "xxx"
	mocknow := time.Now().Truncate(time.Millisecond)

	got := Comment{
		id:      cid,
		author:  operator.OperatorFromUser(uid),
		content: c,
	}

	assert.Equal(t, cid, got.ID())
	assert.Equal(t, uid, *got.Author().User())
	assert.Equal(t, c, got.Content())
	assert.Equal(t, mocknow, got.CreatedAt())
}

func TestComment_SetContent(t *testing.T) {
	comment := Comment{}
	comment.SetContent("xxx")
	assert.Equal(t, "xxx", comment.content)
}

func TestComment_CreatedAt(t *testing.T) {
	got := &Comment{id: NewCommentID()}
	assert.Equal(t, got.id.Timestamp(), got.CreatedAt())
	mocknow := time.Now().Truncate(time.Millisecond)
	got = &Comment{id: NewCommentID()}
	assert.Equal(t, mocknow, got.CreatedAt())
}

func TestComment_Clone(t *testing.T) {
	comment := (&Comment{
		id:      NewCommentID(),
		author:  operator.OperatorFromUser(NewUserID()),
		content: "test",
	})
	assert.Nil(t, (*Comment)(nil).Clone())
	assert.Equal(t, comment, comment.Clone())
	assert.NotSame(t, comment, comment.Clone())
}
