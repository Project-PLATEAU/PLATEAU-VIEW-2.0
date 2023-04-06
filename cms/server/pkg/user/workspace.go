package user

type Workspace struct {
	id      WorkspaceID
	name    string
	members *Members
}

func (t *Workspace) ID() WorkspaceID {
	return t.id
}

func (t *Workspace) Name() string {
	return t.name
}

func (t *Workspace) Members() *Members {
	return t.members
}

func (t *Workspace) IsPersonal() bool {
	return t.members.Fixed()
}

func (t *Workspace) Rename(name string) {
	t.name = name
}
