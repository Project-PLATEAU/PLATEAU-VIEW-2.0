package app

import (
	"context"

	i18nFS "github.com/reearth/reearth-cms/server/i18n"
	"github.com/reearth/reearth-cms/server/internal/adapter"
	"github.com/reearth/reearthx/i18n"
	"golang.org/x/text/language"
)

var (
	// rerror
	_ = i18n.T("not found")
	_ = i18n.T("internal")
	_ = i18n.T("invalid params")
	_ = i18n.T("not implemented")

	localeFS = i18nFS.LocalsFS
	bundle   = i18n.NewBundle(language.English)
)

func init() {
	bundle.MustLoadFS(localeFS, "en.yml", "ja.yml")
}

func getI18nLocalizer(c context.Context) *i18n.Localizer {
	var lang string
	if o := adapter.Operator(c); o != nil {
		lang = o.Lang
	}
	return i18n.NewLocalizer(bundle, lang)
}
