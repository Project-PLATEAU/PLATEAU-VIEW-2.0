package id

import (
	"regexp"
	"strings"

	"github.com/goombaio/namegenerator"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

var keyRegexp = regexp.MustCompile("^[a-zA-Z0-9_-]{1,32}$")

var ngKeys = []string{"id"}

type Key struct {
	key string
}

func NewKey(key string) Key {
	if !keyRegexp.MatchString(key) {
		return Key{}
	}
	k := Key{key}
	return k
}

func RandomKey() Key {
	seed := util.Now().UTC().UnixNano()
	return NewKey(namegenerator.NewNameGenerator(seed).Generate())
}

func (k Key) IsValid() bool {
	return k.key != "" && !strings.HasPrefix(k.key, "_") && !strings.HasPrefix(k.key, "-") && !slices.Contains(ngKeys, k.key)
}

func (k Key) Ref() *Key {
	return &k
}

func (k Key) String() string {
	return k.key
}

func (k *Key) StringRef() *string {
	if k == nil {
		return nil
	}
	return lo.ToPtr(k.key)
}
