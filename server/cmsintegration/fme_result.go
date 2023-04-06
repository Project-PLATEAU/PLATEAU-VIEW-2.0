package cmsintegration

import (
	"sort"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type FMEResult struct {
	Type    string         `json:"type"`
	Status  string         `json:"status"`
	ID      string         `json:"id"`
	LogURL  string         `json:"logUrl"`
	Results map[string]any `json:"results"`
}

type FMEResultAssets struct {
	All  string
	Dic  string
	Bldg []string
	Tran []string
	Frn  string
	Veg  string
	Luse string
	Lsld string
	Urf  []string
	Fld  []string
	Tnm  []string
	Htd  []string
	Ifld []string
}

func (b FMEResult) GetResult() (r FMEResultAssets, unknown []string) {
	for k, v := range b.Results {
		if strings.HasPrefix(k, "bldg") {
			if v2, ok := v.(string); ok {
				r.Bldg = append(r.Bldg, v2)
			}
		} else if strings.HasPrefix(k, "tran") {
			if v2, ok := v.(string); ok {
				r.Tran = append(r.Tran, v2)
			}
		} else if strings.HasPrefix(k, "fld") {
			r.Fld = append(r.Fld, getFld(v)...)
		} else if strings.HasPrefix(k, "tnm") || strings.HasPrefix(k, "tnum") {
			r.Tnm = append(r.Tnm, getFld(v)...)
		} else if strings.HasPrefix(k, "htd") {
			r.Htd = append(r.Htd, getFld(v)...)
		} else if strings.HasPrefix(k, "ifld") {
			r.Ifld = append(r.Ifld, getFld(v)...)
		} else if k == "*" {
			if v2, ok := v.(string); ok {
				r.All = v2
			}
		} else if k == "_dic" {
			if v2, ok := v.(string); ok {
				r.Dic = v2
			}
		} else if k == "luse" {
			if v2, ok := v.(string); ok {
				r.Luse = v2
			}
		} else if k == "veg" {
			if v2, ok := v.(string); ok {
				r.Veg = v2
			}
		} else if k == "frn" {
			if v2, ok := v.(string); ok {
				r.Frn = v2
			}
		} else if k == "lsld" || k == "lsls" {
			if v2, ok := v.(string); ok {
				r.Lsld = v2
			}
		} else if strings.HasPrefix(k, "urf") {
			if v2, ok := v.(string); ok {
				r.Urf = append(r.Urf, v2)
			}
		} else {
			unknown = append(unknown, k)
		}
	}

	sort.Strings(r.Bldg)
	sort.Strings(r.Tran)
	sort.Strings(r.Fld)
	sort.Strings(r.Tnm)
	sort.Strings(r.Htd)
	sort.Strings(r.Ifld)
	sort.Strings(r.Urf)
	return
}

func (d FMEResult) GetDic() string {
	for k, v := range d.Results {
		if k == "_dic" {
			if v2, ok := v.(string); ok {
				return v2
			}
		}
	}
	return ""
}

func getFld(o any) (r []string) {
	switch p := o.(type) {
	case string:
		r = []string{p}
	case []any:
		r = lo.FilterMap(p, func(v any, _ int) (string, bool) {
			v2, ok := v.(string)
			return v2, ok
		})
	case []string:
		r = append(r, p...)
	case map[string]any:
		for _, v := range p {
			if v2, ok := v.(string); ok {
				r = append(r, v2)
			}
		}
		sort.Strings(r)
	}
	return
}

func (a FMEResultAssets) Entries() (s []lo.Entry[string, []string]) {
	if a.All != "" {
		s = append(s, lo.Entry[string, []string]{
			Key:   "all",
			Value: []string{a.All},
		})
	}
	if a.Dic != "" {
		s = append(s, lo.Entry[string, []string]{
			Key:   "dictionary",
			Value: []string{a.Dic},
		})
	}
	if len(a.Bldg) > 0 {
		s = append(s, lo.Entry[string, []string]{
			Key:   "bldg",
			Value: slices.Clone(a.Bldg),
		})
	}
	if len(a.Tran) > 0 {
		s = append(s, lo.Entry[string, []string]{
			Key:   "tran",
			Value: slices.Clone(a.Tran),
		})
	}
	if a.Frn != "" {
		s = append(s, lo.Entry[string, []string]{
			Key:   "frn",
			Value: []string{a.Frn},
		})
	}
	if a.Luse != "" {
		s = append(s, lo.Entry[string, []string]{
			Key:   "luse",
			Value: []string{a.Luse},
		})
	}
	if a.Veg != "" {
		s = append(s, lo.Entry[string, []string]{
			Key:   "veg",
			Value: []string{a.Veg},
		})
	}
	if a.Lsld != "" {
		s = append(s, lo.Entry[string, []string]{
			Key:   "lsld",
			Value: []string{a.Lsld},
		})
	}
	if len(a.Fld) > 0 {
		s = append(s, lo.Entry[string, []string]{
			Key:   "fld",
			Value: slices.Clone(a.Fld),
		})
	}
	if len(a.Tnm) > 0 {
		s = append(s, lo.Entry[string, []string]{
			Key:   "tnm",
			Value: slices.Clone(a.Tnm),
		})
	}
	if len(a.Htd) > 0 {
		s = append(s, lo.Entry[string, []string]{
			Key:   "htd",
			Value: slices.Clone(a.Htd),
		})
	}
	if len(a.Ifld) > 0 {
		s = append(s, lo.Entry[string, []string]{
			Key:   "ifld",
			Value: slices.Clone(a.Ifld),
		})
	}
	if len(a.Urf) > 0 {
		s = append(s, lo.Entry[string, []string]{
			Key:   "urf",
			Value: slices.Clone(a.Urf),
		})
	}
	return
}
