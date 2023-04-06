package value

type Type string

var TypeUnknown = Type("")

func (t Type) Default() bool {
	return defaultTypes.Get(t) != nil
}

func (t Type) None() *Optional {
	if t == TypeUnknown {
		return nil
	}
	return &Optional{t: t}
}

func (t Type) Value(i any) *Value {
	return t.ValueFrom(i, nil)
}

func (t Type) ValueFrom(i any, p TypeRegistry) *Value {
	if v, ok := p.ToValue(t, i); ok {
		return &Value{p: p, v: v, t: t}
	}
	return nil
}
