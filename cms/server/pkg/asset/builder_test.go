package asset

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

type Tests []struct {
	name  string
	input Input
	want  *Asset
	err   error
}

type Input struct {
	id                      ID
	project                 ProjectID
	createdAt               time.Time
	createdByUser           UserID
	createdByIntegration    IntegrationID
	fileName                string
	size                    uint64
	previewType             *PreviewType
	uuid                    string
	thread                  ThreadID
	archiveExtractionStatus *ArchiveExtractionStatus
}

func TestBuilder_Build(t *testing.T) {
	var aid ID = NewID()
	pid := NewProjectID()
	uid := NewUserID()
	iid := NewIntegrationID()
	thid := NewThreadID()
	tim, _ := time.Parse(time.RFC3339, "2021-03-16T04:19:57.592Z")
	var size uint64 = 15

	tests := Tests{
		{
			name: "should create an asset",
			input: Input{
				id:                      aid,
				project:                 pid,
				createdAt:               tim,
				createdByUser:           uid,
				fileName:                "hoge",
				size:                    size,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			want: &Asset{
				id:                      aid,
				project:                 pid,
				createdAt:               tim,
				user:                    &uid,
				fileName:                "hoge",
				size:                    size,
				previewType:             PreviewTypeFromRef(lo.ToPtr("image")),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
		},
		{
			name: "fail: empty project id",
			input: Input{
				id:                      aid,
				createdByUser:           uid,
				fileName:                "hoge",
				size:                    size,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			err: ErrNoProjectID,
		},
		{
			name: "fail: empty id",
			input: Input{
				project:                 pid,
				createdByUser:           uid,
				fileName:                "hoge",
				size:                    size,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			err: ErrInvalidID,
		},
		{
			name: "fail: empty user",
			input: Input{
				id:                      aid,
				project:                 pid,
				fileName:                "hoge",
				size:                    size,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			err: ErrNoUser,
		},
		{
			name: "fail: zero size",
			input: Input{
				id:                      aid,
				project:                 pid,
				createdByUser:           uid,
				fileName:                "hoge",
				size:                    0,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			err: ErrZeroSize,
		},
		{
			name: "fail: invalid threadId",
			input: Input{
				id:                      aid,
				project:                 pid,
				createdByUser:           uid,
				fileName:                "hoge",
				size:                    size,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  ThreadID{},
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			err: ErrNoThread,
		},
		{
			name: "fail: no uuid",
			input: Input{
				id:            aid,
				project:       pid,
				createdByUser: uid,
				fileName:      "hoge",
				size:          size,
				previewType:   PreviewTypeFromRef(lo.ToPtr(PreviewTypeImage.String())),
				thread:        thid,
			},
			err: ErrNoUUID,
		},
		{
			name: "should create asset with id timestamp",
			input: Input{
				id:                      aid,
				project:                 pid,
				createdByUser:           uid,
				fileName:                "hoge",
				size:                    size,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			want: &Asset{
				id:                      aid,
				project:                 pid,
				createdAt:               aid.Timestamp(),
				user:                    &uid,
				fileName:                "hoge",
				size:                    size,
				previewType:             PreviewTypeFromRef(lo.ToPtr("image")),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
		},
		{
			name: "should create asset with id timestamp",
			input: Input{
				id:                      aid,
				project:                 pid,
				createdByIntegration:    iid,
				fileName:                "hoge",
				size:                    size,
				previewType:             lo.ToPtr(PreviewTypeImage),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
			want: &Asset{
				id:                      aid,
				project:                 pid,
				createdAt:               aid.Timestamp(),
				integration:             &iid,
				fileName:                "hoge",
				size:                    size,
				previewType:             PreviewTypeFromRef(lo.ToPtr("image")),
				uuid:                    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:                  thid,
				archiveExtractionStatus: lo.ToPtr(ArchiveExtractionStatusPending),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ab := New().
				ID(tt.input.id).
				Project(tt.input.project).
				CreatedAt(tt.input.createdAt).
				FileName(tt.input.fileName).
				Size(tt.input.size).
				Type(tt.input.previewType).
				UUID(tt.input.uuid).
				Thread(tt.input.thread).
				ArchiveExtractionStatus(tt.input.archiveExtractionStatus)
			if !tt.input.createdByUser.IsNil() {
				ab.CreatedByUser(tt.input.createdByUser)
			}
			if !tt.input.createdByIntegration.IsNil() {
				ab.CreatedByIntegration(tt.input.createdByIntegration)
			}

			got, err := ab.Build()
			if tt.err != nil {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestBuilder_MustBuild(t *testing.T) {
	var aid ID = NewID()
	pid := NewProjectID()
	uid := NewUserID()
	thid := NewThreadID()
	tim, _ := time.Parse(time.RFC3339, "2021-03-16T04:19:57.592Z")
	var size uint64 = 15

	tests := Tests{
		{
			name: "Valid asset",
			input: Input{
				id:            aid,
				project:       pid,
				createdAt:     tim,
				createdByUser: uid,
				fileName:      "hoge",
				size:          size,
				previewType:   PreviewTypeFromRef(lo.ToPtr("image")),
				uuid:          "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:        thid,
			},
			want: &Asset{
				id:          aid,
				project:     pid,
				createdAt:   tim,
				user:        &uid,
				fileName:    "hoge",
				size:        size,
				previewType: PreviewTypeFromRef(lo.ToPtr("image")),
				uuid:        "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:      thid,
			},
		},
		{
			name: "fail: Invalid Id",
			input: Input{
				id:            ID{},
				project:       pid,
				createdAt:     tim,
				createdByUser: uid,
				fileName:      "hoge",
				size:          size,
				previewType:   PreviewTypeFromRef(lo.ToPtr("image")),
				uuid:          "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				thread:        thid,
			},
			err: ErrInvalidID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			build := func() *Asset {
				t.Helper()
				return New().
					ID(tt.input.id).
					Project(tt.input.project).
					CreatedAt(tt.input.createdAt).
					CreatedByUser(tt.input.createdByUser).
					FileName(tt.input.fileName).
					Type(tt.input.previewType).
					Size(tt.input.size).
					UUID(tt.input.uuid).
					Thread(tt.input.thread).
					MustBuild()
			}
			if tt.err != nil {
				assert.PanicsWithValue(t, tt.err, func() { _ = build() })
			} else {
				assert.Equal(t, tt.want, build())
			}
		})
	}
}

func TestBuilder_NewID(t *testing.T) {
	pid := NewProjectID()
	uid := NewUserID()
	var size uint64 = 15
	a := New().NewID().Project(pid).CreatedByUser(uid).Size(size).Thread(NewThreadID()).NewUUID().MustBuild()
	assert.False(t, a.id.IsNil())
}
