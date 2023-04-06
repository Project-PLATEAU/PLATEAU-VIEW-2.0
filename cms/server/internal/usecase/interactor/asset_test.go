package interactor

import (
	"bytes"
	"context"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/reearth/reearth-cms/server/internal/infrastructure/fs"
	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/file"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/project"
	"github.com/reearth/reearth-cms/server/pkg/task"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/idx"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/usecasex"
	"github.com/samber/lo"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestAsset_FindByID(t *testing.T) {
	pid := id.NewProjectID()
	id1 := id.NewAssetID()
	uid1 := id.NewUserID()
	a1 := asset.New().
		ID(id1).
		Project(pid).
		CreatedByUser(uid1).
		Size(1000).
		Thread(id.NewThreadID()).
		NewUUID().
		MustBuild()

	op := &usecase.Operator{}

	type args struct {
		id       id.AssetID
		operator *usecase.Operator
	}

	tests := []struct {
		name    string
		seeds   []*asset.Asset
		args    args
		want    *asset.Asset
		wantErr error
	}{
		{
			name:  "Not found in empty db",
			seeds: []*asset.Asset{},
			args: args{
				id:       id.NewAssetID(),
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:  "Not found",
			seeds: []*asset.Asset{a1},
			args: args{
				id:       id.NewAssetID(),
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:  "Found 1",
			seeds: []*asset.Asset{a1},
			args: args{
				id:       id1,
				operator: op,
			},
			want:    a1,
			wantErr: nil,
		},
		{
			name: "Found 2",
			seeds: []*asset.Asset{
				a1,
				asset.New().
					NewID().
					Project(id.NewProjectID()).
					CreatedByUser(id.NewUserID()).
					Size(1000).
					Thread(id.NewThreadID()).
					NewUUID().
					MustBuild(),
				asset.New().
					NewID().
					Project(id.NewProjectID()).
					CreatedByUser(id.NewUserID()).
					Size(1000).
					Thread(id.NewThreadID()).
					NewUUID().
					MustBuild(),
			},
			args: args{
				id:       id1,
				operator: op,
			},
			want:    a1,
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			for _, a := range tc.seeds {
				err := db.Asset.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}
			assetUC := NewAsset(db, nil)

			got, err := assetUC.FindByID(ctx, tc.args.id, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAsset_FindByIDs(t *testing.T) {
	pid1 := id.NewProjectID()
	uid1 := id.NewUserID()
	id1 := id.NewAssetID()
	id2 := id.NewAssetID()
	tim, _ := time.Parse(time.RFC3339, "2021-03-16T04:19:57.592Z")
	a1 := asset.New().ID(id1).
		Project(pid1).
		CreatedAt(tim).
		CreatedByUser(uid1).
		Size(1000).
		Thread(id.NewThreadID()).
		NewUUID().
		MustBuild()
	a2 := asset.New().ID(id2).
		Project(pid1).
		CreatedAt(tim).
		CreatedByUser(uid1).
		Size(1000).
		Thread(id.NewThreadID()).
		NewUUID().
		MustBuild()

	tests := []struct {
		name    string
		seeds   asset.List
		arg     id.AssetIDList
		want    asset.List
		wantErr error
	}{
		{
			name:    "0 count in empty db",
			seeds:   asset.List{},
			arg:     []id.AssetID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with asset for another workspaces",
			seeds: asset.List{
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
			},
			arg:     []id.AssetID{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single asset",
			seeds: asset.List{
				a1,
			},
			arg:     []id.AssetID{id1},
			want:    asset.List{a1},
			wantErr: nil,
		},
		{
			name: "1 count with multi assets",
			seeds: asset.List{
				a1,
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
			},
			arg:     []id.AssetID{id1},
			want:    asset.List{a1},
			wantErr: nil,
		},
		{
			name: "2 count with multi assets",
			seeds: asset.List{
				a1,
				a2,
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
			},
			arg:     []id.AssetID{id1, id2},
			want:    asset.List{a1, a2},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			for _, a := range tc.seeds {
				err := db.Asset.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}
			assetUC := NewAsset(db, nil)

			got, err := assetUC.FindByIDs(ctx, tc.arg, &usecase.Operator{})
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAsset_FindByProject(t *testing.T) {
	pid := id.NewProjectID()
	aid1 := id.NewAssetID()
	uid1 := id.NewUserID()
	a1 := asset.New().ID(aid1).Project(pid).NewUUID().
		CreatedByUser(uid1).Size(1000).Thread(id.NewThreadID()).MustBuild()

	aid2 := id.NewAssetID()
	uid2 := id.NewUserID()
	a2 := asset.New().ID(aid2).Project(pid).NewUUID().
		CreatedByUser(uid2).Size(1000).Thread(id.NewThreadID()).MustBuild()

	op := &usecase.Operator{}

	type args struct {
		pid      id.ProjectID
		f        interfaces.AssetFilter
		operator *usecase.Operator
	}
	tests := []struct {
		name    string
		seeds   asset.List
		args    args
		want    asset.List
		wantErr error
	}{
		{
			name:  "0 count in empty db",
			seeds: asset.List{},
			args: args{
				pid:      id.NewProjectID(),
				operator: op,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "0 count with asset for another projects",
			seeds: asset.List{
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
			},
			args: args{
				pid:      id.NewProjectID(),
				operator: op,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "1 count with single asset",
			seeds: asset.List{
				a1,
			},
			args: args{
				pid: pid,
				f: interfaces.AssetFilter{
					Pagination: usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap(),
				},
				operator: op,
			},
			want:    asset.List{a1},
			wantErr: nil,
		},
		{
			name: "1 count with multi assets",
			seeds: asset.List{
				a1,
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
			},
			args: args{
				pid: pid,
				f: interfaces.AssetFilter{
					Pagination: usecasex.CursorPagination{First: lo.ToPtr(int64(1))}.Wrap(),
				},
				operator: op,
			},
			want:    asset.List{a1},
			wantErr: nil,
		},
		{
			name: "2 count with multi assets",
			seeds: asset.List{
				a1,
				a2,
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
				asset.New().NewID().Project(id.NewProjectID()).NewUUID().
					CreatedByUser(id.NewUserID()).Size(1000).Thread(id.NewThreadID()).MustBuild(),
			},
			args: args{
				pid: pid,
				f: interfaces.AssetFilter{
					Pagination: usecasex.CursorPagination{First: lo.ToPtr(int64(2))}.Wrap(),
				},
				operator: op,
			},
			want:    asset.List{a1, a2},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			for _, a := range tc.seeds {
				err := db.Asset.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}
			assetUC := NewAsset(db, nil)

			got, _, err := assetUC.FindByProject(ctx, tc.args.pid, tc.args.f, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAsset_Create(t *testing.T) {
	mocktime := time.Now()
	ws := user.NewWorkspace().NewID().MustBuild()

	pid1 := id.NewProjectID()
	p1 := project.New().ID(pid1).Workspace(ws.ID()).UpdatedAt(mocktime).MustBuild()

	u := user.New().NewID().Name("aaa").Email("aaa@bbb.com").Workspace(ws.ID()).MustBuild()
	op := &usecase.Operator{
		User:               lo.ToPtr(u.ID()),
		WritableWorkspaces: []id.WorkspaceID{ws.ID()},
	}

	buf := bytes.NewBufferString("Hello")
	buf2 := bytes.NewBufferString("Hello")
	af := asset.NewFile().Name("aaa.txt").Size(uint64(buf.Len())).Path("aaa.txt").Build()
	af2 := asset.NewFile().Name("aaa.txt").Size(uint64(buf2.Len())).Path("aaa.txt").Build()

	type args struct {
		cpp      interfaces.CreateAssetParam
		operator *usecase.Operator
	}
	tests := []struct {
		name     string
		seeds    []*asset.Asset
		args     args
		want     *asset.Asset
		wantFile *asset.File
		wantErr  error
	}{
		{
			name:  "Create",
			seeds: []*asset.Asset{},
			args: args{
				cpp: interfaces.CreateAssetParam{
					ProjectID: p1.ID(),
					File: &file.File{
						Path:    "aaa.txt",
						Content: io.NopCloser(buf),
						Size:    int64(buf.Len()),
					},
				},
				operator: op,
			},
			want: asset.New().
				NewID().
				Project(p1.ID()).
				CreatedByUser(u.ID()).
				FileName("aaa.txt").
				Size(uint64(buf.Len())).
				Type(asset.PreviewTypeUnknown.Ref()).
				Thread(id.NewThreadID()).
				NewUUID().
				ArchiveExtractionStatus(lo.ToPtr(asset.ArchiveExtractionStatusInProgress)).
				MustBuild(),
			wantFile: af,
			wantErr:  nil,
		},
		{
			name:  "Create skip decompress",
			seeds: []*asset.Asset{},
			args: args{
				cpp: interfaces.CreateAssetParam{
					ProjectID: p1.ID(),
					File: &file.File{
						Path:    "aaa.txt",
						Content: io.NopCloser(buf2),
						Size:    int64(buf2.Len()),
					},
					SkipDecompression: true,
				},
				operator: op,
			},
			want: asset.New().
				NewID().
				Project(p1.ID()).
				CreatedByUser(u.ID()).
				FileName("aaa.txt").
				Size(uint64(buf2.Len())).
				Type(asset.PreviewTypeUnknown.Ref()).
				Thread(id.NewThreadID()).
				NewUUID().
				ArchiveExtractionStatus(lo.ToPtr(asset.ArchiveExtractionStatusSkipped)).
				MustBuild(),
			wantFile: af2,
			wantErr:  nil,
		},
		{
			name:  "Create invalid file size",
			seeds: []*asset.Asset{},
			args: args{
				cpp: interfaces.CreateAssetParam{
					ProjectID: p1.ID(),
					File: &file.File{
						Path:    "aaa.txt",
						Content: io.NopCloser(buf),
						Size:    10*1024*1024*1024 + 1,
					},
				},
				operator: op,
			},
			want:     nil,
			wantFile: nil,
			wantErr:  gateway.ErrFileTooLarge,
		},
		{
			name:  "Create invalid file",
			seeds: []*asset.Asset{},
			args: args{
				cpp: interfaces.CreateAssetParam{
					ProjectID: p1.ID(),
					File:      nil,
				},
				operator: op,
			},
			want:     nil,
			wantFile: nil,
			wantErr:  interfaces.ErrFileNotIncluded,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			mfs := afero.NewMemMapFs()
			f, _ := fs.NewFile(mfs, "")
			runnerGw := NewMockRunner()

			err := db.User.Save(ctx, u)
			assert.NoError(t, err)

			err2 := db.Project.Save(ctx, p1.Clone())
			assert.Nil(t, err2)

			for _, a := range tc.seeds {
				err := db.Asset.Save(ctx, a.Clone())
				assert.NoError(t, err)
			}

			assetUC := Asset{
				repos: db,
				gateways: &gateway.Container{
					File:       f,
					TaskRunner: runnerGw,
				},
				ignoreEvent: true,
			}

			got, gotFile, err := assetUC.Create(ctx, tc.args.cpp, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)

			if strings.HasPrefix(got.PreviewType().String(), "image/") {
				assert.Equal(t, asset.PreviewTypeImage.Ref(), got.PreviewType())
			} else {
				assert.Equal(t, asset.PreviewTypeUnknown.Ref(), got.PreviewType())
			}

			assert.Equal(t, tc.want.Project(), got.Project())
			assert.Equal(t, tc.want.PreviewType(), got.PreviewType())
			assert.Equal(t, tc.want.ArchiveExtractionStatus(), got.ArchiveExtractionStatus())

			dbGot, err := db.Asset.FindByID(ctx, got.ID())
			assert.NoError(t, err)
			assert.Equal(t, tc.want.Project(), dbGot.Project())
			assert.Equal(t, tc.want.PreviewType(), dbGot.PreviewType())
			assert.Equal(t, tc.want.ArchiveExtractionStatus(), dbGot.ArchiveExtractionStatus())

			assert.Equal(t, gotFile, tc.wantFile)
		})
	}
}

func TestAsset_Update(t *testing.T) {
	uid := id.NewUserID()
	ws := user.NewWorkspace().NewID().MustBuild()
	pid1 := id.NewProjectID()
	p := project.New().ID(pid1).Workspace(ws.ID()).MustBuild()

	var pti asset.PreviewType = asset.PreviewTypeImage
	var ptg asset.PreviewType = asset.PreviewTypeGeo

	aid1 := id.NewAssetID()
	thid := id.NewThreadID()
	a1 := asset.New().ID(aid1).Project(pid1).NewUUID().
		CreatedByUser(uid).Size(1000).Thread(thid).MustBuild()
	a1Updated := asset.New().ID(aid1).Project(pid1).UUID(a1.UUID()).
		CreatedByUser(uid).Size(1000).Thread(thid).Type(&pti).MustBuild()

	pid2 := id.NewProjectID()
	aid2 := id.NewAssetID()
	a2 := asset.New().ID(aid2).Project(pid2).NewUUID().
		CreatedByUser(uid).Size(1000).Thread(id.NewThreadID()).MustBuild()

	op := &usecase.Operator{
		User:             &uid,
		Integration:      nil,
		OwningWorkspaces: []id.WorkspaceID{ws.ID()},
		OwningProjects:   []id.ProjectID{pid1},
	}

	type args struct {
		upp      interfaces.UpdateAssetParam
		operator *usecase.Operator
	}
	tests := []struct {
		name    string
		seeds   []*asset.Asset
		args    args
		want    *asset.Asset
		wantErr error
	}{
		{
			name:  "update",
			seeds: []*asset.Asset{a1, a2},
			args: args{
				upp: interfaces.UpdateAssetParam{
					AssetID:     aid1,
					PreviewType: &pti,
				},
				operator: op,
			},
			want:    a1Updated,
			wantErr: nil,
		},
		{
			name:  "update not found",
			seeds: []*asset.Asset{a1, a2},
			args: args{
				upp: interfaces.UpdateAssetParam{
					AssetID:     idx.ID[id.Asset]{},
					PreviewType: &ptg,
				},
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			ctx := context.Background()
			db := memory.New()

			err := db.Project.Save(ctx, p)
			assert.NoError(t, err)
			for _, p := range tc.seeds {
				err := db.Asset.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			assetUC := NewAsset(db, &gateway.Container{})

			got, err := assetUC.Update(ctx, tc.args.upp, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestAsset_UpdateFiles(t *testing.T) {
	uid := id.NewUserID()
	assetID1 := asset.NewID()
	assetID2 := asset.NewID()
	ws := user.NewWorkspace().NewID().MustBuild()
	proj := project.New().NewID().Workspace(ws.ID()).MustBuild()

	thid := id.NewThreadID()
	sp := lo.ToPtr(asset.ArchiveExtractionStatusPending)
	a1 := asset.New().
		ID(assetID1).
		Project(proj.ID()).
		CreatedByUser(uid).
		Size(1000).
		UUID("5130c89f-8f67-4766-b127-49ee6796d464").
		Thread(thid).
		ArchiveExtractionStatus(sp).
		MustBuild()
	a1f := asset.NewFile().Name("xxx").Path("/xxx.zip").GuessContentType().Build()
	a2 := asset.New().
		ID(assetID2).
		Project(proj.ID()).
		CreatedByUser(uid).
		Size(1000).
		UUID("5130c89f-8f67-4766-b127-49ee6796d464").
		Thread(id.NewThreadID()).
		ArchiveExtractionStatus(sp).
		MustBuild()
	a2f := asset.NewFile().Build()

	op := &usecase.Operator{
		User:             &uid,
		OwningWorkspaces: []id.WorkspaceID{ws.ID()},
		OwningProjects:   []id.ProjectID{proj.ID()},
	}

	tests := []struct {
		name            string
		seedAssets      []*asset.Asset
		seedFiles       map[asset.ID]*asset.File
		seedProjects    []*project.Project
		prepareFileFunc func() afero.Fs
		assetID         id.AssetID
		status          *asset.ArchiveExtractionStatus
		want            *asset.Asset
		wantFile        *asset.File
		wantErr         error
	}{
		{
			name: "update asset not found",
			prepareFileFunc: func() afero.Fs {
				return mockFs()
			},
			assetID: assetID1,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:       "update file not found",
			seedAssets: []*asset.Asset{a1.Clone(), a2.Clone()},
			seedFiles: map[asset.ID]*asset.File{
				a1.ID(): a1f,
				a2.ID(): a2f,
			},
			prepareFileFunc: func() afero.Fs {
				return afero.NewMemMapFs()
			},
			assetID: assetID1,
			want:    nil,
			wantErr: gateway.ErrFileNotFound,
		},
		{
			name:       "update",
			seedAssets: []*asset.Asset{a1.Clone(), a2.Clone()},
			seedFiles: map[asset.ID]*asset.File{
				a1.ID(): a1f,
				a2.ID(): a2f,
			},
			seedProjects: []*project.Project{proj},
			prepareFileFunc: func() afero.Fs {
				return mockFs()
			},
			assetID: assetID1,
			status:  sp,
			want: asset.New().
				ID(assetID1).
				Project(proj.ID()).
				CreatedByUser(uid).
				UUID("5130c89f-8f67-4766-b127-49ee6796d464").
				Size(1000).
				Thread(thid).
				ArchiveExtractionStatus(sp).
				MustBuild(),
			wantFile: asset.NewFile().Name("xxx").Path("/xxx.zip").GuessContentType().Children([]*asset.File{
				asset.NewFile().Name("xxx").Path("/xxx").Dir().Children([]*asset.File{
					asset.NewFile().Name("yyy").Path("/xxx/yyy").Dir().Children([]*asset.File{
						asset.NewFile().Name("hello.txt").Path("/xxx/yyy/hello.txt").GuessContentType().Build(),
					}).Build(),
					asset.NewFile().Name("zzz.txt").Path("/xxx/zzz.txt").GuessContentType().Build(),
				}).Build(),
			}).Build(),
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			ctx := context.Background()
			db := memory.New()

			fileGw := lo.Must(fs.NewFile(tc.prepareFileFunc(), ""))

			err := db.Project.Save(ctx, proj)
			assert.NoError(t, err)
			for _, p := range tc.seedAssets {
				err := db.Asset.Save(ctx, p.Clone())
				assert.Nil(t, err)
			}
			for id, f := range tc.seedFiles {
				err := db.AssetFile.Save(ctx, id, f.Clone())
				assert.Nil(t, err)
			}
			for _, p := range tc.seedProjects {
				err := db.Project.Save(ctx, p.Clone())
				assert.Nil(t, err)
			}

			assetUC := Asset{
				repos: db,
				gateways: &gateway.Container{
					File: fileGw,
				},
				ignoreEvent: true,
			}
			got, err := assetUC.UpdateFiles(ctx, tc.assetID, tc.status, op)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)

			if tc.wantErr != nil {
				gotf, err := db.AssetFile.FindByID(ctx, tc.assetID)
				assert.NoError(t, err)
				assert.Equal(t, tc.wantFile, gotf)
			}
		})
	}
}

func TestAsset_Delete(t *testing.T) {
	uid := id.NewUserID()

	ws := user.NewWorkspace().NewID().MustBuild()
	proj1 := project.New().NewID().Workspace(ws.ID()).MustBuild()
	aid1 := id.NewAssetID()
	a1 := asset.New().ID(aid1).Project(proj1.ID()).NewUUID().
		CreatedByUser(uid).Size(1000).Thread(id.NewThreadID()).MustBuild()

	proj2 := project.New().NewID().MustBuild()
	aid2 := id.NewAssetID()
	a2 := asset.New().ID(aid2).Project(proj2.ID()).NewUUID().
		CreatedByUser(uid).Size(1000).Thread(id.NewThreadID()).MustBuild()

	op := &usecase.Operator{
		User:             &uid,
		OwningWorkspaces: []id.WorkspaceID{ws.ID()},
		OwningProjects:   []id.ProjectID{proj1.ID()},
	}

	type args struct {
		id       id.AssetID
		operator *usecase.Operator
	}
	tests := []struct {
		name         string
		seedsAsset   []*asset.Asset
		seedsProject []*project.Project
		args         args
		want         []*asset.Asset
		mockAssetErr bool
		wantErr      error
	}{
		{
			name:         "delete",
			seedsAsset:   []*asset.Asset{a1, a2},
			seedsProject: []*project.Project{proj1, proj2},
			args: args{
				id:       aid1,
				operator: op,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name:       "delete not found",
			seedsAsset: []*asset.Asset{a1, a2},
			args: args{
				id:       id.NewAssetID(),
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:       "delete od",
			seedsAsset: []*asset.Asset{},
			args: args{
				id:       aid2,
				operator: op,
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()

			for _, p := range tc.seedsAsset {
				err := db.Asset.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			for _, p := range tc.seedsProject {
				err := db.Project.Save(ctx, p.Clone())
				assert.NoError(t, err)
			}
			assetUC := Asset{
				repos:       db,
				gateways:    &gateway.Container{},
				ignoreEvent: true,
			}
			id, err := assetUC.Delete(ctx, tc.args.id, tc.args.operator)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.args.id, id)
			assert.NoError(t, err)

			_, err = db.Asset.FindByID(ctx, tc.args.id)
			assert.Equal(t, rerror.ErrNotFound, err)
		})
	}
}

func TestAsset_getExtenalFile(t *testing.T) {
	URL := "https://cms.com/test.txt"
	f := lo.Must(os.Open("testdata/test.txt"))
	defer f.Close()
	z := lo.Must(io.ReadAll(f))

	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("GET", URL, func(r *http.Request) (*http.Response, error) {
		res := httpmock.NewBytesResponse(200, z)
		res.Header.Set("Content-Type", mime.TypeByExtension(path.Ext(URL)))
		return res, nil
	})

	expected := file.File{Path: "/test.txt", Content: f}

	got, err := getExternalFile(context.Background(), URL)
	assert.NoError(t, err)
	assert.Equal(t, expected.Path, got.Path)
	assert.Equal(t, z, lo.Must(io.ReadAll(got.Content)))

}

type file2 struct {
	gateway.File
}

func (f *file2) GetURL(*asset.Asset) string {
	return "xxx"
}

func TestAsset_GetURL(t *testing.T) {
	uc := &Asset{
		gateways: &gateway.Container{
			File: &file2{},
		},
	}
	assert.Equal(t, "xxx", uc.GetURL(nil))
}

func mockFs() afero.Fs {
	files := map[string]string{
		"assets/51/30c89f-8f67-4766-b127-49ee6796d464/xxx.zip":           "xxx",
		"assets/51/30c89f-8f67-4766-b127-49ee6796d464/xxx/zzz.txt":       "zzz",
		"assets/51/30c89f-8f67-4766-b127-49ee6796d464/xxx/yyy/hello.txt": "hello",
		"plugins/aaa~1.0.0/foo.js":                                       "bar",
		"published/s.json":                                               "{}",
	}

	fs := afero.NewMemMapFs()
	for name, content := range files {
		f, _ := fs.Create(name)
		_, _ = f.WriteString(content)
		_ = f.Close()
	}
	return fs
}

// mockRunner implements gateway.TaskRunner
type mockRunner struct{}

func NewMockRunner() gateway.TaskRunner {
	return &mockRunner{}
}

func (r *mockRunner) Run(context.Context, task.Payload) error {
	return nil
}

func Test_detectPreviewType(t *testing.T) {
	tests := []struct {
		name  string
		files []gateway.FileEntry
		want  *asset.PreviewType
	}{
		{
			name: "MVT",
			files: []gateway.FileEntry{
				{
					Name: "test/0/123.mvt",
					Size: 123,
				},
			},
			want: lo.ToPtr(asset.PreviewTypeGeoMvt),
		},
		{
			name: "3d tiles",
			files: []gateway.FileEntry{
				{
					Name: "test/tileset.json",
					Size: 123,
				},
			},
			want: lo.ToPtr(asset.PreviewTypeGeo3dTiles),
		},
		{
			name: "Unknown",
			files: []gateway.FileEntry{
				{
					Name: "test.jpg",
					Size: 123,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, detectPreviewType(tt.files))
		})
	}
}
