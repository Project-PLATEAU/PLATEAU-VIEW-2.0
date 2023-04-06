package mongo

import (
	"context"
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/asset"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearthx/mongox"
	"github.com/reearth/reearthx/mongox/mongotest"
	"github.com/reearth/reearthx/rerror"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAssetFileRepo_FindByID(t *testing.T) {
	aid := asset.NewID()

	tests := []struct {
		name    string
		seeds   map[asset.ID]*asset.File
		arg     id.AssetID
		want    *asset.File
		wantErr error
	}{
		{
			name:    "Not found in empty db",
			seeds:   nil,
			arg:     id.NewAssetID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Not found",
			seeds: map[asset.ID]*asset.File{
				asset.NewID(): asset.NewFile().Name("aaa.txt").Path("/aaa.txt").Size(100).Build(),
			},
			arg:     id.NewAssetID(),
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name:    "Empty",
			seeds:   nil,
			arg:     aid,
			want:    nil,
			wantErr: rerror.ErrNotFound,
		},
		{
			name: "Found 1",
			seeds: map[asset.ID]*asset.File{
				aid:           asset.NewFile().Name("aaa.txt").Path("/aaa.txt").Size(100).Build(),
				asset.NewID(): asset.NewFile().Name("aaa.txt").Path("/aaa.txt").Size(100).Build(),
			},
			arg:     aid,
			want:    asset.NewFile().Name("aaa.txt").Path("/aaa.txt").Size(100).Build(),
			wantErr: nil,
		},
	}

	initDB := mongotest.Connect(t)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			db := initDB(t)
			ctx := context.Background()
			_, _ = db.Collection("asset").InsertOne(ctx, bson.M{
				"id":   aid.String(),
				"hoge": "bar",
			})
			client := mongox.NewClientWithDatabase(db)

			r := NewAssetFile(client)
			for id, a := range tc.seeds {
				err := r.Save(ctx, id, a)
				assert.NoError(t, err)
			}

			got, err := r.FindByID(ctx, tc.arg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.want, got)

			c, _ := db.Collection("asset").CountDocuments(ctx, bson.M{"id": aid.String(), "hoge": "bar"})
			assert.Equal(t, int64(1), c)
		})
	}
}
