package searchindex

import (
	"net/url"
	"path"
)

const modelKey = "plateau"

type Config struct {
	CMSBase           string
	CMSToken          string
	CMSStorageProject string
	// optioanl
	CMSStorageModel string
	CMSModel        string
	Delegate        bool
	DelegateURL     string
	Debug           bool
	// internal
	skipIndexer bool
}

func (c *Config) Default() {
	if c.CMSModel == "" {
		c.CMSModel = modelKey
	}
}

func getAssetBase(u *url.URL) string {
	u2 := *u
	b := path.Join(path.Dir(u.Path), pathFileName(u.Path))
	u2.Path = b
	return u2.String()
}
