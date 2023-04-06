package usecase

import (
	"testing"

	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestOperator_EventOperator(t *testing.T) {
	uId := id.NewUserID()
	op := Operator{
		User:        &uId,
		Integration: nil,
	}

	eOp := op.Operator()

	assert.NotNil(t, eOp.User())
	assert.Nil(t, eOp.Integration())
	assert.Equal(t, &uId, eOp.User())

	iId := id.NewIntegrationID()
	op = Operator{
		User:        nil,
		Integration: &iId,
	}

	eOp = op.Operator()

	assert.Nil(t, eOp.User())
	assert.NotNil(t, eOp.Integration())
	assert.Equal(t, &iId, eOp.Integration())
}
