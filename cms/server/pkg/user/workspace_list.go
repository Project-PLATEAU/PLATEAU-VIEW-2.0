package user

type WorkspaceList []*Workspace

func (l WorkspaceList) FilterByID(ids ...WorkspaceID) WorkspaceList {
	if l == nil {
		return nil
	}

	res := make(WorkspaceList, 0, len(l))
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

func (l WorkspaceList) FilterByUserRole(u ID, r Role) WorkspaceList {
	if l == nil || u.IsNil() || r == "" {
		return nil
	}

	res := make(WorkspaceList, 0, len(l))
	for _, t := range l {
		tr := t.Members().UserRole(u)
		if tr == r {
			res = append(res, t)
		}
	}
	return res
}

func (l WorkspaceList) FilterByIntegrationRole(i IntegrationID, r Role) WorkspaceList {
	if l == nil || i.IsNil() || r == "" {
		return nil
	}

	res := make(WorkspaceList, 0, len(l))
	for _, t := range l {
		tr := t.Members().IntegrationRole(i)
		if tr == r {
			res = append(res, t)
		}
	}
	return res
}

func (l WorkspaceList) FilterByUserRoleIncluding(u ID, r Role) WorkspaceList {
	if l == nil || u.IsNil() || r == "" {
		return nil
	}

	res := make(WorkspaceList, 0, len(l))
	for _, t := range l {
		tr := t.Members().UserRole(u)
		if tr.Includes(r) {
			res = append(res, t)
		}
	}
	return res
}

func (l WorkspaceList) IDs() WorkspaceIDList {
	if l == nil {
		return nil
	}

	res := make([]WorkspaceID, 0, len(l))
	for _, t := range l {
		res = append(res, t.ID())
	}
	return res
}
