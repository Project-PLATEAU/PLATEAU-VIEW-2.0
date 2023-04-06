package workspace

type List []*Workspace

func (l List) FilterByID(ids ...ID) List {
	if l == nil {
		return nil
	}

	res := make(List, 0, len(l))
	for _, id := range ids {
		var t2 *Workspace
		for _, t := range l {
			if t.ID() == id {
				t2 = t
				break
			}
		}
		if t2 != nil {
			res = append(res, t2)
		}
	}
	return res
}
func (l List) FilterByUserRole(u UserID, r Role) List {
	if l == nil || u.IsNil() || r == "" {
		return nil
	}

	res := make(List, 0, len(l))
	for _, t := range l {
		tr := t.Members().GetRole(u)
		if tr == r {
			res = append(res, t)
		}
	}
	return res
}

func (l List) FilterByUserRoleIncluding(u UserID, r Role) List {
	if l == nil || u.IsNil() || r == "" {
		return nil
	}

	res := make(List, 0, len(l))
	for _, t := range l {
		tr := t.Members().GetRole(u)
		if tr.Includes(r) {
			res = append(res, t)
		}
	}
	return res
}

func (l List) IDs() []ID {
	if l == nil {
		return nil
	}

	res := make([]ID, 0, len(l))
	for _, t := range l {
		res = append(res, t.ID())
	}
	return res
}
