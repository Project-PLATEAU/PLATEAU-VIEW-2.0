package project

const (
	PublicationScopePrivate PublicationScope = "private"
	PublicationScopeLimited PublicationScope = "limited"
	PublicationScopePublic  PublicationScope = "public"
)

type PublicationScope string

type Publication struct {
	scope       PublicationScope
	assetPublic bool
}

func NewPublication(scope PublicationScope, assetPublic bool) *Publication {
	p := &Publication{}
	p.SetScope(scope)
	p.SetAssetPublic(assetPublic)
	return p
}

func (p *Publication) Scope() PublicationScope {
	if p.scope == "" {
		return PublicationScopePrivate
	}
	return p.scope
}

func (p *Publication) AssetPublic() bool {
	return p.assetPublic
}

func (p *Publication) SetScope(scope PublicationScope) {
	if scope != PublicationScopePrivate && scope != PublicationScopeLimited && scope != PublicationScopePublic {
		scope = PublicationScopePrivate
	}

	p.scope = scope
}

func (p *Publication) SetAssetPublic(assetPublic bool) {
	p.assetPublic = assetPublic
}

func (p *Publication) Clone() *Publication {
	if p == nil {
		return nil
	}

	return &Publication{
		scope:       p.scope,
		assetPublic: p.assetPublic,
	}
}
