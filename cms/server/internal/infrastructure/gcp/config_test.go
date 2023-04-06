package gcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloudTasksConfig_buildQueueUrl(t *testing.T) {
	type fields struct {
		GCPProject string
		GCPRegion  string
		QueueName  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr error
	}{
		{
			name: "should return property queue URL",
			fields: fields{
				GCPProject: "a",
				GCPRegion:  "b",
				QueueName:  "c",
			},
			want: "projects/a/locations/b/queues/c",
		},
		{
			name: "should return error if settings isn't provided",
			fields: fields{
				GCPProject: "a",
				GCPRegion:  "b",
			},
			wantErr: ErrMissignConfig,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &TaskConfig{
				GCPProject: tt.fields.GCPProject,
				GCPRegion:  tt.fields.GCPRegion,
				QueueName:  tt.fields.QueueName,
			}
			got, err := c.buildQueueUrl()
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
