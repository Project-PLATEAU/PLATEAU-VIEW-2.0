package thread

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/stretchr/testify/assert"
)

func TestThread_Getters(t *testing.T) {
	thid := NewID()
	wid := NewWorkspaceID()
	c := []*Comment{}

	got := Thread{
		id:        thid,
		workspace: wid,
		comments:  c,
	}

	assert.Equal(t, thid, got.ID())
	assert.Equal(t, wid, got.Workspace())
	assert.Equal(t, c, got.Comments())
}

func TestThread_Comments(t *testing.T) {
	var got *Thread = nil
	assert.Nil(t, got.Comments())

	c := []*Comment{{id: NewCommentID()}}
	got = &Thread{
		comments: c,
	}
	assert.Equal(t, c, got.Comments())
}

func TestThread_HasComment(t *testing.T) {
	c := NewComment(NewCommentID(), operator.OperatorFromUser(NewUserID()), "test")
	thread := (&Thread{
		id:        NewID(),
		workspace: NewWorkspaceID(),
		comments: []*Comment{
			{id: NewCommentID()}, c,
		},
	})

	ok := thread.HasComment(c.id)
	assert.True(t, ok)

	ok = thread.HasComment(id.NewCommentID())
	assert.False(t, ok)

	thread = nil
	ok = thread.HasComment(c.id)
	assert.False(t, ok)
}

func TestThread_AddComment(t *testing.T) {
	thread := (&Thread{
		id:        NewID(),
		workspace: NewWorkspaceID(),
	})
	c := NewComment(NewCommentID(), operator.OperatorFromUser(NewUserID()), "test")
	err := thread.AddComment(c)
	assert.NoError(t, err)
	assert.True(t, thread.HasComment(c.id))

	err = thread.AddComment(c)
	assert.ErrorIs(t, err, ErrCommentAlreadyExist)
}

func TestThread_UpdateComment(t *testing.T) {
	c := NewComment(NewCommentID(), operator.OperatorFromUser(NewUserID()), "test")
	thread := (&Thread{
		id:        NewID(),
		workspace: NewWorkspaceID(),
		comments: []*Comment{
			{id: NewCommentID()}, c,
		},
	})

	err := thread.UpdateComment(NewCommentID(), "updated")
	assert.ErrorIs(t, err, ErrCommentDoesNotExist)

	err = thread.UpdateComment(c.id, "updated")
	assert.NoError(t, err)
	assert.Equal(t, "updated", c.content)

}

func TestThread_DeleteComment(t *testing.T) {
	c := NewComment(NewCommentID(), operator.OperatorFromUser(NewUserID()), "test")
	thread := (&Thread{
		id:        NewID(),
		workspace: NewWorkspaceID(),
		comments: []*Comment{
			{id: NewCommentID()}, c,
		},
	})

	err := thread.DeleteComment(NewCommentID())
	assert.ErrorIs(t, err, ErrCommentDoesNotExist)

	err = thread.DeleteComment(c.id)
	assert.NoError(t, err)
	assert.False(t, thread.HasComment(c.id))
}

func TestThread_Comment(t *testing.T) {
	c := NewComment(NewCommentID(), operator.OperatorFromUser(NewUserID()), "test")
	thread := (&Thread{
		id:        NewID(),
		workspace: NewWorkspaceID(),
		comments: []*Comment{
			{id: NewCommentID()}, c,
		},
	})

	cc := thread.Comment(c.id)
	assert.Equal(t, c, cc)

}

func TestThread_Clone(t *testing.T) {
	thread := (&Thread{
		id:        NewID(),
		workspace: NewWorkspaceID(),
		comments: []*Comment{
			{id: NewCommentID()},
		},
	})
	assert.Nil(t, (*Thread)(nil).Clone())
	assert.Equal(t, thread, thread.Clone())
	assert.NotSame(t, thread, thread.Clone())
}
