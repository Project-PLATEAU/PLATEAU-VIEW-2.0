package operator

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestOperator(t *testing.T) {
	uID := user.NewID()
	iID := NewIntegrationID()

	uOp := OperatorFromUser(uID)
	iOp := OperatorFromIntegration(iID)
	cmsOp := OperatorFromMachine()

	assert.NotNil(t, uOp)
	assert.NotNil(t, iOp)

	assert.Equal(t, uID, *uOp.User())
	assert.Nil(t, uOp.Integration())
	assert.False(t, uOp.Machine())

	assert.Equal(t, iID, *iOp.Integration())
	assert.Nil(t, iOp.User())
	assert.False(t, uOp.Machine())

	assert.True(t, cmsOp.Machine())
	assert.Nil(t, cmsOp.User())
	assert.Nil(t, cmsOp.Integration())

	assert.True(t, uOp.Validate())
	assert.True(t, iOp.Validate())
	assert.True(t, cmsOp.Validate())

}
