package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPublication(t *testing.T) {
	assert.Equal(t, &Publication{
		scope:       PublicationScopePrivate,
		assetPublic: false,
	}, NewPublication(PublicationScopePrivate, false))
	assert.Equal(t, &Publication{
		scope:       PublicationScopeLimited,
		assetPublic: true,
	}, NewPublication(PublicationScopeLimited, true))
	assert.Equal(t, &Publication{
		scope:       PublicationScopePublic,
		assetPublic: false,
	}, NewPublication(PublicationScopePublic, false))
	assert.Equal(t, &Publication{
		scope:       PublicationScopePrivate,
		assetPublic: true,
	}, NewPublication("", true))
}

func TestPublication_Scope(t *testing.T) {
	assert.Equal(t, PublicationScopePrivate, (&Publication{}).Scope())
	assert.Equal(t, PublicationScopePublic, (&Publication{scope: PublicationScopePublic}).Scope())
}

func TestPublication_AssetPublic(t *testing.T) {
	assert.True(t, (&Publication{assetPublic: true}).AssetPublic())
}

func TestPublication_SetScope(t *testing.T) {
	p := &Publication{
		scope: PublicationScopePublic,
	}
	p.SetScope(PublicationScopePrivate)
	assert.Equal(t, &Publication{
		scope: PublicationScopePrivate,
	}, p)

	p = &Publication{}
	p.SetScope(PublicationScopeLimited)
	assert.Equal(t, &Publication{
		scope: PublicationScopeLimited,
	}, p)

	p = &Publication{}
	p.SetScope(PublicationScopePublic)
	assert.Equal(t, &Publication{
		scope: PublicationScopePublic,
	}, p)

	p = &Publication{
		scope: PublicationScopePublic,
	}
	p.SetScope("")
	assert.Equal(t, &Publication{
		scope: PublicationScopePrivate,
	}, p)
}

func TestPublication_SetAssetPublic(t *testing.T) {
	p := &Publication{
		assetPublic: false,
	}
	p.SetAssetPublic(true)
	assert.Equal(t, &Publication{
		assetPublic: true,
	}, p)

	p = &Publication{
		assetPublic: true,
	}
	p.SetAssetPublic(false)
	assert.Equal(t, &Publication{
		assetPublic: false,
	}, p)
}

func TestPublication_Clone(t *testing.T) {
	p := &Publication{
		assetPublic: false,
		scope:       PublicationScopeLimited,
	}
	p2 := p.Clone()
	assert.Equal(t, p, p2)
	assert.NotSame(t, p, p2)
	assert.Nil(t, (*Publication)(nil).Clone())
}
