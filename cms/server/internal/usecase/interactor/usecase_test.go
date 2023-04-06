package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/usecasex"
	"github.com/stretchr/testify/assert"
)

func TestUc_checkPermission(t *testing.T) {
	tid := id.NewWorkspaceID()

	tests := []struct {
		name               string
		op                 *usecase.Operator
		readableWorkspaces id.WorkspaceIDList
		writableWorkspaces id.WorkspaceIDList
		wantErr            bool
	}{
		{
			name:    "nil operator",
			wantErr: false,
		},
		{
			name:               "nil operator 2",
			readableWorkspaces: id.WorkspaceIDList{id.NewWorkspaceID()},
			wantErr:            false,
		},
		{
			name:               "can read a workspace",
			readableWorkspaces: id.WorkspaceIDList{tid},
			op: &usecase.Operator{
				ReadableWorkspaces: id.WorkspaceIDList{tid},
			},
			wantErr: false,
		},
		{
			name:               "cannot read a workspace",
			readableWorkspaces: id.WorkspaceIDList{id.NewWorkspaceID()},
			op: &usecase.Operator{
				ReadableWorkspaces: id.WorkspaceIDList{},
			},
			wantErr: true,
		},
		{
			name:               "can write a workspace",
			writableWorkspaces: id.WorkspaceIDList{tid},
			op: &usecase.Operator{
				WritableWorkspaces: id.WorkspaceIDList{tid},
			},
			wantErr: false,
		},
		{
			name:               "cannot write a workspace",
			writableWorkspaces: id.WorkspaceIDList{tid},
			op: &usecase.Operator{
				WritableWorkspaces: id.WorkspaceIDList{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &uc{
				readableWorkspaces: tt.readableWorkspaces,
				writableWorkspaces: tt.writableWorkspaces,
			}
			got := e.checkPermission(tt.op)
			if tt.wantErr {
				assert.Equal(t, interfaces.ErrOperationDenied, got)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestUc(t *testing.T) {
	workspaces := id.WorkspaceIDList{id.NewWorkspaceID(), id.NewWorkspaceID(), id.NewWorkspaceID()}
	assert.Equal(t, &uc{}, Usecase())
	assert.Equal(t, &uc{readableWorkspaces: workspaces}, (&uc{}).WithReadableWorkspaces(workspaces...))
	assert.Equal(t, &uc{writableWorkspaces: workspaces}, (&uc{}).WithWritableWorkspaces(workspaces...))
	assert.Equal(t, &uc{tx: true}, (&uc{}).Transaction())
}

func TestRun(t *testing.T) {
	ctx := context.Background()
	err := errors.New("test")
	a, b, c := &struct{}{}, &struct{}{}, &struct{}{}

	// regular1: without tx
	tr := &usecasex.NopTransaction{}
	r := &repo.Container{Transaction: tr}
	gota, gotb, gotc, goterr := Run3(
		ctx, nil, r,
		Usecase(),
		func(ctx context.Context) (any, any, any, error) {
			return a, b, c, nil
		},
	)
	assert.Same(t, a, gota)
	assert.Same(t, b, gotb)
	assert.Same(t, c, gotc)
	assert.Nil(t, goterr)
	assert.False(t, tr.IsCommitted())

	// regular2: with tx
	tr = &usecasex.NopTransaction{}
	r.Transaction = tr
	_ = Run0(
		ctx, nil, r,
		Usecase().Transaction(),
		func(ctx context.Context) error {
			return nil
		},
	)
	assert.True(t, tr.IsCommitted())

	// iregular1: the usecase returns an error
	tr = &usecasex.NopTransaction{}
	r.Transaction = tr
	goterr = Run0(
		ctx, nil, r,
		Usecase().Transaction(),
		func(ctx context.Context) error {
			return err
		},
	)
	assert.Same(t, err, goterr)
	assert.False(t, tr.IsCommitted())

	// iregular2: tx.Begin returns an error
	tr = &usecasex.NopTransaction{
		BeginError: err,
	}
	r.Transaction = tr
	goterr = Run0(
		ctx, nil, r,
		Usecase().Transaction(),
		func(ctx context.Context) error {
			return nil
		},
	)
	assert.Same(t, err, goterr)
	assert.False(t, tr.IsCommitted())

	// iregular3: tx.End returns an error
	tr = &usecasex.NopTransaction{
		CommitError: err,
	}
	r.Transaction = tr
	goterr = Run0(
		ctx, nil, r,
		Usecase().Transaction(),
		func(ctx context.Context) error {
			return nil
		},
	)
	assert.Same(t, err, goterr)
	assert.True(t, tr.IsCommitted())
}
