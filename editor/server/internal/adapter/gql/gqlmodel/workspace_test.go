package gqlmodel

import (
	"testing"

	"github.com/reearth/reearth/server/pkg/workspace"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestToRole(t *testing.T) {
	assert.Equal(t, Role(RoleOwner), ToRole(workspace.RoleOwner))
	assert.Equal(t, Role(RoleWriter), ToRole(workspace.RoleWriter))
	assert.Equal(t, Role(RoleReader), ToRole(workspace.RoleReader))
	assert.Equal(t, Role(""), ToRole(workspace.Role("unknown")))
}

func TestFromRole(t *testing.T) {
	assert.Equal(t, workspace.RoleOwner, FromRole(RoleOwner))
	assert.Equal(t, workspace.RoleWriter, FromRole(RoleWriter))
	assert.Equal(t, workspace.RoleReader, FromRole(RoleReader))
	assert.Equal(t, workspace.Role(""), FromRole("unknown"))
}

func TestToPolicy(t *testing.T) {
	assert.Equal(t, &Policy{
		ID:                    ID("x"),
		Name:                  "aaa",
		ProjectCount:          lo.ToPtr(1),
		MemberCount:           lo.ToPtr(2),
		PublishedProjectCount: lo.ToPtr(3),
		LayerCount:            lo.ToPtr(4),
		AssetStorageSize:      lo.ToPtr(int64(5)),
		DatasetCount:          lo.ToPtr(6),
		DatasetSchemaCount:    lo.ToPtr(7),
	}, ToPolicy(workspace.NewPolicy(workspace.PolicyOption{
		ID:                    workspace.PolicyID("x"),
		Name:                  "aaa",
		ProjectCount:          lo.ToPtr(1),
		MemberCount:           lo.ToPtr(2),
		PublishedProjectCount: lo.ToPtr(3),
		LayerCount:            lo.ToPtr(4),
		AssetStorageSize:      lo.ToPtr(int64(5)),
		DatasetCount:          lo.ToPtr(6),
		DatasetSchemaCount:    lo.ToPtr(7),
	})))
	assert.Nil(t, ToPolicy(nil))
}
