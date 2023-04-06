package schema

import (
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/key"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Schema struct {
	id        ID
	project   ProjectID
	workspace id.WorkspaceID
	fields    []*Field
}

func (s *Schema) ID() ID {
	return s.id
}

func (s *Schema) Workspace() id.WorkspaceID {
	return s.workspace
}

func (s *Schema) Project() ProjectID {
	return s.project
}

func (s *Schema) SetWorkspace(workspace id.WorkspaceID) {
	s.workspace = workspace
}

func (s *Schema) HasField(f FieldID) bool {
	return lo.SomeBy(s.fields, func(g *Field) bool { return g.ID() == f })
}

func (s *Schema) HasFieldByKey(k string) bool {
	return lo.SomeBy(s.fields, func(g *Field) bool { return g.Key().String() == k })
}

func (s *Schema) AddField(f *Field) {
	if s.HasField(f.ID()) {
		return
	}
	if s.Fields().Count() == 0 {
		f.order = 0
	} else {
		//get the biggest order
		f.order = s.Fields().Ordered()[s.Fields().Count()-1].Order() + 1
	}
	s.fields = append(s.fields, f)
}

func (s *Schema) Field(fId FieldID) *Field {
	f, _ := lo.Find(s.fields, func(f *Field) bool { return f.id == fId })
	return f
}

func (s *Schema) FieldByIDOrKey(fId *FieldID, key *key.Key) *Field {
	f, _ := lo.Find(s.fields, func(f *Field) bool {
		return fId != nil && f.id == *fId || key != nil && key.IsValid() && f.key == *key
	})
	return f
}

func (s *Schema) Fields() FieldList {
	var fl FieldList = slices.Clone(s.fields)
	return fl.Ordered()
}

func (s *Schema) RemoveField(fid FieldID) {
	for i, field := range s.fields {
		if field.id == fid {
			s.fields = slices.Delete(s.fields, i, i+1)
			return
		}
	}
}

func (s *Schema) Clone() *Schema {
	if s == nil {
		return nil
	}

	return &Schema{
		id:        s.ID(),
		workspace: s.Workspace().Clone(),
		fields:    slices.Clone(s.fields),
	}
}
