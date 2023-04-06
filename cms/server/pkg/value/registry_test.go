package value

type tpmock struct {
	TypeProperty
}

func (*tpmock) ToValue(i any) (any, bool) {
	return i.(string) + "a", true
}

func (*tpmock) ToInterface(v any) (any, bool) {
	return v.(string) + "bar", true
}

func (*tpmock) Equal(a, b any) bool {
	return a.(string) == b.(string)
}

func (*tpmock) IsEmpty(v any) bool {
	return v.(string) == ""
}

func (*tpmock) Validate(v any) bool {
	_, ok := v.(string)
	return ok
}
