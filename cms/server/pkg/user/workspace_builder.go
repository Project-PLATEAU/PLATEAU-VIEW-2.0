package user

type WorkspaceBuilder struct {
	t            *Workspace
	members      map[ID]Member
	integrations map[IntegrationID]Member
	personal     bool
}

func NewWorkspace() *WorkspaceBuilder {
	return &WorkspaceBuilder{t: &Workspace{}}
}

func (b *WorkspaceBuilder) Build() (*Workspace, error) {
	if b.t.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.members == nil {
		b.t.members = NewMembers()
	} else {
		b.t.members = NewMembersWith(b.members)
	}
	if b.integrations != nil {
		b.t.members.integrations = b.integrations
	}
	b.t.members.fixed = b.personal
	return b.t, nil
}

func (b *WorkspaceBuilder) MustBuild() *Workspace {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *WorkspaceBuilder) ID(id WorkspaceID) *WorkspaceBuilder {
	b.t.id = id
	return b
}

func (b *WorkspaceBuilder) NewID() *WorkspaceBuilder {
	b.t.id = NewWorkspaceID()
	return b
}

func (b *WorkspaceBuilder) Name(name string) *WorkspaceBuilder {
	b.t.name = name
	return b
}

func (b *WorkspaceBuilder) Members(members map[ID]Member) *WorkspaceBuilder {
	b.members = members
	return b
}

func (b *WorkspaceBuilder) Integrations(integrations map[IntegrationID]Member) *WorkspaceBuilder {
	b.integrations = integrations
	return b
}

func (b *WorkspaceBuilder) Personal(p bool) *WorkspaceBuilder {
	b.personal = p
	return b
}
