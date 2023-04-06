package user

type Theme string

const (
	ThemeDefault Theme = "default"
	ThemeLight   Theme = "light"
	ThemeDark    Theme = "dark"
)

func (t Theme) Ref() *Theme {
	return &t
}
