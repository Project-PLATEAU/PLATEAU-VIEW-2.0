package value

type Optional struct {
	t Type
	v *Value
}

func NewOptional(t Type, v *Value) *Optional {
	if v == nil {
		return t.None()
	}
	if t != v.Type() {
		return nil
	}
	return v.Some()
}

func OptionalFrom(v *Value) *Optional {
	if v.Type() == TypeUnknown {
		return nil
	}
	return &Optional{
		t: v.Type(),
		v: v,
	}
}

func (ov *Optional) Type() Type {
	if ov == nil {
		return TypeUnknown
	}
	return ov.t
}

func (ov *Optional) IsSome() bool {
	return !ov.IsNone()
}

func (ov *Optional) IsNone() bool {
	return ov == nil || ov.v == nil
}

func (ov *Optional) IsEmpty() bool {
	return ov.IsNone() || ov.v.IsEmpty()
}

func (ov *Optional) Value() *Value {
	if ov == nil || ov.t == TypeUnknown || ov.v == nil {
		return nil
	}
	return ov.v
}

func (ov *Optional) TypeAndValue() (Type, *Value) {
	return ov.Type(), ov.Value()
}

func (ov *Optional) SetValue(v *Value) {
	if ov == nil || ov.t == TypeUnknown || (v != nil && ov.t != v.Type()) {
		return
	}
	ov.v = v.Clone()
}

func (ov *Optional) Clone() *Optional {
	if ov == nil {
		return nil
	}
	return &Optional{
		t: ov.t,
		v: ov.v.Clone(),
	}
}

// Cast tries to convert the value to the new type and generates a new Optional.
func (ov *Optional) Cast(t Type) *Optional {
	if ov == nil || ov.t == TypeUnknown {
		return nil
	}

	if ov.v == nil {
		return t.None()
	}

	return &Optional{
		t: t,
		v: ov.v.Cast(t),
	}
}
