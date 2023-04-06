package interactor

import (
	"context"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/memory"
	"github.com/reearth/reearth-cms/server/internal/usecase"
	"github.com/reearth/reearth-cms/server/internal/usecase/interfaces"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/integration"
	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/reearth/reearthx/rerror"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	Now        time.Time
	Op         *usecase.Operator
	Uri        *url.URL
	UId        id.UserID
	IId1, IId2 id.IntegrationID
	I1, I2     *integration.Integration
}

func testSuite() testData {
	now := time.Now().Truncate(time.Millisecond).UTC()
	wid := id.NewWorkspaceID()
	uId := id.NewUserID()
	u := user.New().Name("aaa").ID(uId).Email("aaa@bbb.com").Workspace(wid).MustBuild()
	op := &usecase.Operator{
		User:               lo.ToPtr(u.ID()),
		ReadableWorkspaces: nil,
		WritableWorkspaces: nil,
		OwningWorkspaces:   []id.WorkspaceID{wid},
	}
	uri := lo.Must(url.Parse("https://sub.hugo2.com/dir?p=1#test"))
	iId1 := id.NewIntegrationID()
	iId2 := id.NewIntegrationID()
	i1 := integration.New().ID(iId1).Name("i1").Developer(uId).Type(integration.TypePrivate).LogoUrl(uri).UpdatedAt(now).MustBuild()
	i2 := integration.New().ID(iId2).Name("i2").Developer(uId).Type(integration.TypePrivate).LogoUrl(uri).UpdatedAt(now).MustBuild()
	return testData{
		Now:  now,
		Op:   op,
		Uri:  uri,
		UId:  uId,
		IId1: iId1,
		IId2: iId2,
		I1:   i1,
		I2:   i2,
	}
}

func assertIntegrationEq(t *testing.T, expected, got *integration.Integration) {
	if expected == nil || got == nil {
		assert.Nil(t, expected)
		assert.Nil(t, got)
		return
	}
	assert.Equal(t, expected.Type(), got.Type())
	assert.Equal(t, expected.Name(), got.Name())
	assert.Equal(t, expected.Developer(), got.Developer())
	assert.Equal(t, expected.Description(), got.Description())
	assert.Equal(t, expected.LogoUrl(), got.LogoUrl())
	if expected.Webhooks() == nil || got.Webhooks() == nil {
		assert.Nil(t, expected.Webhooks())
		assert.Nil(t, got.Webhooks())
		return
	}
	for i, w := range got.Webhooks() {
		e := expected.Webhooks()[i]
		assertWebhookEq(t, e, w)
	}
	// assert.Equal(t, expected.ID(), got.ID())
	// assert.Equal(t, expected.UpdatedAt(), got.UpdatedAt())
}

func assertWebhookEq(t *testing.T, expected, got *integration.Webhook) {
	if expected == nil || got == nil {
		assert.Nil(t, expected)
		assert.Nil(t, got)
		return
	}
	assert.Equal(t, expected.Name(), got.Name())
	assert.Equal(t, expected.Active(), got.Active())
	assert.Equal(t, expected.URL(), got.URL())
	assert.Equal(t, expected.Trigger(), got.Trigger())
	// assert.Equal(t, expected.ID(), got.ID())
	// assert.Equal(t, expected.UpdatedAt(), got.UpdatedAt())
}

func TestIntegration_Create(t *testing.T) {
	ts := testSuite()

	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    interfaces.CreateIntegrationParam
		want    *integration.Integration
		wantErr error
	}{
		{
			name:  "create",
			seeds: []*integration.Integration{},
			args: interfaces.CreateIntegrationParam{
				Name:        "i1",
				Description: nil,
				Type:        integration.TypePrivate,
				Logo:        *ts.Uri,
			},
			want:    ts.I1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			got, err := i.Create(ctx, tt.args, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Regexp(t, regexp.MustCompile("secret_[a-zA-Z0-9]{43}"), got.Token())
			assert.False(t, got.ID().IsEmpty())
			assertIntegrationEq(t, tt.want, got)

			got, err = db.Integration.FindByID(ctx, got.ID())
			assert.NoError(t, err)
			assertIntegrationEq(t, tt.want, got)
		})
	}
}

func TestIntegration_Update(t *testing.T) {
	ts := testSuite()

	type args struct {
		id     integration.ID
		params interfaces.UpdateIntegrationParam
	}
	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    args
		want    *integration.Integration
		wantErr error
	}{
		{
			name:  "update",
			seeds: []*integration.Integration{ts.I1},
			args: args{
				id: ts.IId1,
				params: interfaces.UpdateIntegrationParam{
					Name:        nil,
					Description: nil,
					Logo:        nil,
				},
			},
			want:    ts.I1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			got, err := i.Update(ctx, tt.args.id, tt.args.params, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			// assert.Regexp(t, regexp.MustCompile("secret_[a-zA-Z0-9]{43}"), got.Token())
			assert.False(t, got.ID().IsEmpty())
			assertIntegrationEq(t, tt.want, got)

			got, err = db.Integration.FindByID(ctx, tt.args.id)
			assert.NoError(t, err)
			assertIntegrationEq(t, tt.want, got)
		})
	}
}

func TestIntegration_Delete(t *testing.T) {
	ts := testSuite()

	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    integration.ID
		wantErr error
	}{
		{
			name:    "delete",
			seeds:   []*integration.Integration{ts.I1, ts.I2},
			args:    ts.IId1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			err := i.Delete(ctx, tt.args, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)

			got, err := db.Integration.FindByID(ctx, tt.args)
			assert.Nil(t, got)
			assert.Equal(t, rerror.ErrNotFound, err)
		})
	}
}

func TestIntegration_FindByIDs(t *testing.T) {
	ts := testSuite()

	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    []integration.ID
		want    []*integration.Integration
		wantErr error
	}{
		{
			name:    "test",
			seeds:   []*integration.Integration{ts.I1},
			args:    []integration.ID{ts.IId1},
			want:    []*integration.Integration{ts.I1},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			got, err := i.FindByIDs(ctx, tt.args, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			for idx, in := range got {
				// assert.Regexp(t, regexp.MustCompile("secret_[a-zA-Z0-9]{43}"), got.Token())
				assert.False(t, in.ID().IsEmpty())
				assertIntegrationEq(t, tt.want[idx], in)
			}

		})
	}
}

func TestIntegration_FindByUser(t *testing.T) {
	ts := testSuite()

	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    []integration.ID
		want    []*integration.Integration
		wantErr error
	}{
		{
			name:    "test",
			seeds:   []*integration.Integration{ts.I1},
			args:    []integration.ID{ts.IId1},
			want:    []*integration.Integration{ts.I1},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			got, err := i.FindByIDs(ctx, tt.args, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			for idx, in := range got {
				// assert.Regexp(t, regexp.MustCompile("secret_[a-zA-Z0-9]{43}"), got.Token())
				assert.False(t, in.ID().IsEmpty())
				assertIntegrationEq(t, tt.want[idx], in)
			}

		})
	}
}

func TestIntegration_CreateWebhook(t *testing.T) {
	ts := testSuite()

	type args struct {
		id     integration.ID
		params interfaces.CreateWebhookParam
	}
	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    args
		want    *integration.Webhook
		wantErr error
	}{
		{
			name:  "create",
			seeds: []*integration.Integration{ts.I1},
			args: args{
				id: ts.IId1,
				params: interfaces.CreateWebhookParam{
					Name:    "w1",
					URL:     *ts.Uri,
					Active:  true,
					Trigger: &interfaces.WebhookTriggerParam{},
				},
			},
			want:    integration.NewWebhookBuilder().NewID().Name("w1").Url(ts.Uri).Active(true).Trigger(integration.WebhookTrigger{}).MustBuild(),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			got, err := i.CreateWebhook(ctx, tt.args.id, tt.args.params, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assertWebhookEq(t, tt.want, got)
		})
	}
}

func TestIntegration_UpdateWebhook(t *testing.T) {
	ts := testSuite()

	wId := id.NewWebhookID()
	ts.I2.SetWebhook([]*integration.Webhook{integration.NewWebhookBuilder().ID(wId).MustBuild()})
	type args struct {
		iId    integration.ID
		wId    integration.WebhookID
		params interfaces.UpdateWebhookParam
	}
	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    args
		want    *integration.Webhook
		wantErr error
	}{
		{
			name:  "create",
			seeds: []*integration.Integration{ts.I2},
			args: args{
				iId: ts.IId2,
				wId: wId,
				params: interfaces.UpdateWebhookParam{
					Name:    lo.ToPtr("w1"),
					URL:     ts.Uri,
					Active:  lo.ToPtr(true),
					Trigger: &interfaces.WebhookTriggerParam{},
				},
			},
			want:    integration.NewWebhookBuilder().ID(wId).Name("w1").Url(ts.Uri).Active(true).Trigger(integration.WebhookTrigger{}).MustBuild(),
			wantErr: nil,
		},
		{
			name:  "update item not found",
			seeds: []*integration.Integration{},
			args: args{
				iId: ts.IId1,
				params: interfaces.UpdateWebhookParam{
					Name:    lo.ToPtr("w1"),
					URL:     ts.Uri,
					Active:  lo.ToPtr(true),
					Trigger: &interfaces.WebhookTriggerParam{},
				},
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:  "update item not found",
			seeds: []*integration.Integration{ts.I1},
			args: args{
				iId: ts.IId1,
				params: interfaces.UpdateWebhookParam{
					Name:    lo.ToPtr("w1"),
					URL:     ts.Uri,
					Active:  lo.ToPtr(true),
					Trigger: &interfaces.WebhookTriggerParam{},
				},
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			got, err := i.UpdateWebhook(ctx, tt.args.iId, tt.args.wId, tt.args.params, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assertWebhookEq(t, tt.want, got)
		})
	}
}

func TestIntegration_DeleteWebhook(t *testing.T) {
	ts := testSuite()

	wId := id.NewWebhookID()
	ts.I2.SetWebhook([]*integration.Webhook{integration.NewWebhookBuilder().ID(wId).MustBuild()})
	type args struct {
		iId    integration.ID
		wId    integration.WebhookID
		params interfaces.UpdateWebhookParam
	}
	tests := []struct {
		name    string
		seeds   []*integration.Integration
		args    args
		want    *integration.Webhook
		wantErr error
	}{
		{
			name:  "create",
			seeds: []*integration.Integration{ts.I2},
			args: args{
				iId: ts.IId2,
				wId: wId,
				params: interfaces.UpdateWebhookParam{
					Name:    lo.ToPtr("w1"),
					URL:     ts.Uri,
					Active:  lo.ToPtr(true),
					Trigger: &interfaces.WebhookTriggerParam{},
				},
			},
			want:    integration.NewWebhookBuilder().ID(wId).Name("w1").Url(ts.Uri).Active(true).MustBuild(),
			wantErr: nil,
		},
		{
			name:  "update item not found",
			seeds: []*integration.Integration{},
			args: args{
				iId: ts.IId1,
				params: interfaces.UpdateWebhookParam{
					Name:    lo.ToPtr("w1"),
					URL:     ts.Uri,
					Active:  lo.ToPtr(true),
					Trigger: &interfaces.WebhookTriggerParam{},
				},
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:  "update item not found",
			seeds: []*integration.Integration{ts.I1},
			args: args{
				iId: ts.IId1,
				params: interfaces.UpdateWebhookParam{
					Name:    lo.ToPtr("w1"),
					URL:     ts.Uri,
					Active:  lo.ToPtr(true),
					Trigger: &interfaces.WebhookTriggerParam{},
				},
			},
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			db := memory.New()
			defer memory.MockNow(db, ts.Now)
			for _, s := range tt.seeds {
				err := db.Integration.Save(ctx, s.Clone())
				assert.NoError(t, err)
			}

			i := Integration{
				repos: db,
			}
			err := i.DeleteWebhook(ctx, tt.args.iId, tt.args.wId, ts.Op)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestNewIntegration(t *testing.T) {
	r := memory.New()
	assert.Equal(t, &Integration{repos: r}, NewIntegration(r, nil))
}
