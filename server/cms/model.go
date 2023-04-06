package cms

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

const (
	AssetArchiveExtractionStatusDone = "done"
	tag                              = "cms"
)

type Asset struct {
	ID                      string `json:"id,omitempty"`
	ProjectID               string `json:"projectId,omitempty"`
	URL                     string `json:"url,omitempty"`
	ContentType             string `json:"contentType,omitempty"`
	ArchiveExtractionStatus string `json:"archiveExtractionStatus,omitempty"`
	File                    *File  `json:"file,omitempty"`
}

type File struct {
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	Path        string `json:"path"`
	Children    []File `json:"children"`
}

func (f File) Paths() []string {
	return filePaths(f)
}

func filePaths(f File) (p []string) {
	if len(f.Children) == 0 {
		p = append(p, f.Path)
	}
	p = append(p, lo.FlatMap(f.Children, func(f File, _ int) []string {
		return filePaths(f)
	})...)
	return p
}

func (a *Asset) Clone() *Asset {
	if a == nil {
		return nil
	}
	return &Asset{
		ID:                      a.ID,
		ProjectID:               a.ProjectID,
		URL:                     a.URL,
		ArchiveExtractionStatus: a.ArchiveExtractionStatus,
	}
}

func (a *Asset) ToPublic() *PublicAsset {
	if a == nil {
		return nil
	}
	return &PublicAsset{
		Type:                    "asset",
		ID:                      a.ID,
		URL:                     a.URL,
		ContentType:             a.ContentType,
		ArchiveExtractionStatus: a.ArchiveExtractionStatus,
		// Files: ,
	}
}

type Model struct {
	ID           string    `json:"id"`
	Key          string    `json:"key,omitempty"`
	LastModified time.Time `json:"lastModified,omitempty"`
}

type Items struct {
	Items      []Item `json:"items"`
	Page       int    `json:"page"`
	PerPage    int    `json:"perPage"`
	TotalCount int    `json:"totalCount"`
}

func (r Items) HasNext() bool {
	if r.PerPage == 0 {
		return false
	}
	return r.TotalCount > r.Page*r.PerPage
}

type Item struct {
	ID      string  `json:"id"`
	ModelID string  `json:"modelId"`
	Fields  []Field `json:"fields"`
}

func (i *Item) Clone() *Item {
	if i == nil {
		return nil
	}
	return &Item{
		ID:      i.ID,
		ModelID: i.ModelID,
		Fields:  slices.Clone(i.Fields),
	}
}

func (i Item) Field(id string) *Field {
	f, ok := lo.Find(i.Fields, func(f Field) bool { return f.ID == id })
	if ok {
		return &f
	}
	return nil
}

func (i Item) FieldByKey(key string) *Field {
	f, ok := lo.Find(i.Fields, func(f Field) bool { return f.Key == key })
	if ok {
		return &f
	}
	return nil
}

func (d Item) Unmarshal(i any) {
	if i == nil {
		return
	}

	v := reflect.ValueOf(i)
	if v.IsNil() {
		return
	}

	v = v.Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		key, _, _ := strings.Cut(tag, ",")
		if key == "" || key == "-" {
			continue
		}

		vf := v.FieldByName(f.Name)
		if !vf.CanSet() {
			continue
		}

		if key == "id" {
			if f.Type.Kind() == reflect.String {
				vf.SetString(d.ID)
			}
			continue
		}

		if itf := d.FieldByKey(key); itf != nil {
			if f.Type.Kind() == reflect.String {
				if itfv := itf.ValueString(); itfv != nil {
					vf.SetString(*itfv)
				}
			} else if f.Type.Kind() == reflect.Slice && f.Type.Elem().Kind() == reflect.String {
				if te := f.Type.Elem(); te.Name() == "string" {
					if itfv := itf.ValueStrings(); itfv != nil {
						vf.Set(reflect.ValueOf(itfv))
					}
				} else if itfv := itf.ValueStrings(); itfv != nil {
					s := reflect.MakeSlice(f.Type, 0, len(itfv))
					for _, v := range itfv {
						rv := reflect.ValueOf(v).Convert(te)
						s = reflect.Append(s, rv)
					}
					vf.Set(s)
				}
			} else if itf.Value != nil && reflect.TypeOf(itf.Value).AssignableTo(vf.Type()) {
				vf.Set(reflect.ValueOf(itf.Value))
			}
		}
	}
}

func Marshal(i any, item *Item) {
	if item == nil || i == nil {
		return
	}

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Pointer {
		if v.IsNil() {
			return
		}
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	ni := Item{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(tag)
		key, ty, _ := strings.Cut(tag, ",")
		if key == "" || key == "-" {
			continue
		}

		vf := v.FieldByName(f.Name)
		if key == "id" {
			ni.ID, _ = vf.Interface().(string)
			continue
		}

		vft := vf.Type()
		var i any
		if vft.Kind() == reflect.String {
			v := vf.Convert(reflect.TypeOf("")).Interface()
			if v != "" {
				i = v
			}
		} else if vft.Kind() == reflect.Slice && vft.Elem().Kind() == reflect.String && vf.Len() > 0 {
			st := reflect.TypeOf("")
			v := make([]string, 0, vf.Len())
			for i := 0; i < cap(v); i++ {
				vfs := vf.Index(i).Convert(st)
				v = append(v, vfs.String())
			}
			i = v
		}

		if i != nil {
			ni.Fields = append(ni.Fields, Field{
				Key:   key,
				Type:  ty,
				Value: i,
			})
		}
	}

	*item = ni
}

type Field struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type"`
	Value any    `json:"value"`
	Key   string `json:"key,omitempty"`
}

func (f *Field) ValueString() *string {
	if f == nil {
		return nil
	}

	if v, ok := f.Value.(string); ok {
		return &v
	}

	return nil
}

func (f *Field) ValueStrings() []string {
	if f == nil {
		return nil
	}

	if v, ok := f.Value.([]string); ok {
		return v
	}

	if v, ok := f.Value.([]any); ok {
		return lo.FilterMap(v, func(e any, _ int) (string, bool) {
			s, ok := e.(string)
			return s, ok
		})
	}

	return nil
}

func (f *Field) ValueInt() *int {
	if f == nil {
		return nil
	}

	if v, ok := f.Value.(int); ok {
		return &v
	}
	return nil
}

func (f *Field) ValueJSON() (any, error) {
	if f == nil {
		return nil, nil
	}
	s := f.ValueString()
	if s == nil {
		return nil, nil
	}

	var j any
	err := json.Unmarshal([]byte(*s), &j)
	return j, err
}

func (f *Field) Clone() *Field {
	if f == nil {
		return nil
	}
	return &Field{
		ID:    f.ID,
		Type:  f.Type,
		Value: f.Value,
		Key:   f.Key,
	}
}

type Schema struct {
	ID        string        `json:"id"`
	Fields    []SchemaField `json:"fields"`
	ProjectID string        `json:"projectId"`
}

func (d Schema) FieldIDByKey(k string) string {
	f, ok := lo.Find(d.Fields, func(f SchemaField) bool {
		return f.Key == k
	})
	if !ok {
		return ""
	}
	return f.ID
}

type SchemaField struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Key  string `json:"key"`
}
