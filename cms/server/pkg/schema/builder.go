package schema

type Builder struct {
	s *Schema
}

func New() *Builder {
	return &Builder{s: &Schema{}}
}

func (b *Builder) Build() (*Schema, error) {
	if b.s.id.IsNil() {
		return nil, ErrInvalidID
	}
	if b.s.workspace.IsNil() {
		return nil, ErrInvalidID
	}
	if b.s.project.IsNil() {
		return nil, ErrInvalidID
	}
	return b.s, nil
}

func (b *Builder) MustBuild() *Schema {
	r, err := b.Build()
	if err != nil {
		panic(err)
	}
	return r
}

func (b *Builder) ID(id ID) *Builder {
	b.s.id = id.Clone()
	return b
}

func (b *Builder) NewID() *Builder {
	b.s.id = NewID()
	return b
}

func (b *Builder) Workspace(workspace WorkspaceID) *Builder {
	b.s.workspace = workspace.Clone()
	return b
}

func (b *Builder) Project(project ProjectID) *Builder {
	b.s.project = project.Clone()
	return b
}

func (b *Builder) Fields(fields FieldList) *Builder {
	b.s.fields = fields.Clone()
	return b
}
