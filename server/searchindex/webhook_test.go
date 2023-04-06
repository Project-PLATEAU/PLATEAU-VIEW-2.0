package searchindex

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"math/rand"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/eukarya-inc/reearth-plateauview/server/cms/cmswebhook"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/rerror"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestWebhook_AssetAlreadyDecompressed(t *testing.T) {
	assert := assert.New(t)
	log := initLogger(t)

	itemsProject := "prj"
	itemsModel := "itemitem"
	storageProject := "sys"
	assetName := "bldg_lod1"
	assetName2 := "bldg2_lod1"
	assets := []*cms.Asset{
		{
			ID:                      "bldg",
			URL:                     "https://example.com/" + assetName + ".zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: cms.AssetArchiveExtractionStatusDone,
		},
		{
			ID:                      "bldg2",
			URL:                     "https://example.com/" + assetName2 + ".zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: cms.AssetArchiveExtractionStatusDone,
		},
		{
			ID:                      "bldg3",
			URL:                     "https://example.com/bldg_lod2.zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: cms.AssetArchiveExtractionStatusDone,
		},
	}
	items := []*cms.Item{
		{
			ID: "item",
			Fields: Item{
				Bldg:              []string{assets[0].ID, assets[1].ID, assets[2].ID},
				SearchIndex:       nil,
				SearchIndexStatus: StatusReady,
			}.Fields(),
			ModelID: itemsModel,
		},
	}
	c := newMockedCMS(t, itemsProject, itemsModel, storageProject, storageModel, items, assets)
	h := webhookHandler(c, Config{
		CMSModel:          itemsModel,
		CMSStorageProject: storageProject,
		skipIndexer:       true,
	})

	payload := &cmswebhook.Payload{
		Type: cmswebhook.EventItemUpdate,
		ItemData: &cmswebhook.ItemData{
			Item: items[0],
			Model: &cms.Model{
				Key: itemsModel,
			},
			Schema: &cms.Schema{
				ProjectID: itemsProject,
			},
		},
		Operator: cmswebhook.Operator{
			User: &cmswebhook.User{ID: "aaa"},
		},
	}

	assert.NoError(h(httptest.NewRequest("POST", "/", nil), payload))

	// assert logs
	assert.Contains(log(), "searchindex webhook: item: ")
	assert.Equal("searchindex webhook: start processing", log())
	assert.Equal("searchindex webhook: start processing for "+assetName, log())
	assert.Equal("searchindex webhook: start processing for "+assetName2, log())
	assert.Equal("searchindex webhook: done", log())

	// assert item
	item2, _ := c.items.Load(items[0].ID)
	assert.Equal(StatusOK, ItemFrom(*item2).SearchIndexStatus)

	// asset comments
	assert.Equal(
		[]string{"検索インデックスの構築を開始しました。", "検索インデックスの構築が完了しました。"},
		util.DR(c.comments.Load(items[0].ID)),
	)
}

func TestWebhook_AssetNotDecompressed(t *testing.T) {
	assert := assert.New(t)
	log := initLogger(t)

	itemsProject := "prj"
	itemsModel := "itemitem"
	storageProject := "sys"
	assetName := "bldg_lod1"
	assetName2 := "bldg2_lod1"
	assets := []*cms.Asset{
		{
			ID:                      "bldg",
			URL:                     "https://example.com/" + assetName + ".zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: "",
		},
		{
			ID:                      "bldg2",
			URL:                     "https://example.com/" + assetName2 + ".zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: cms.AssetArchiveExtractionStatusDone,
		},
	}
	items := []*cms.Item{
		{
			ID: "item",
			Fields: Item{
				Bldg:              []string{assets[0].ID, assets[1].ID},
				SearchIndex:       nil,
				SearchIndexStatus: StatusReady,
			}.Fields(),
			ModelID: itemsModel,
		},
	}
	c := newMockedCMS(t, itemsProject, itemsModel, storageProject, storageModel, items, assets)
	h := webhookHandler(c, Config{
		CMSModel:          itemsModel,
		CMSStorageProject: storageProject,
		skipIndexer:       true,
	})

	// item.update
	payload := &cmswebhook.Payload{
		Type: cmswebhook.EventItemUpdate,
		ItemData: &cmswebhook.ItemData{
			Item: items[0],
			Model: &cms.Model{
				Key: itemsModel,
			},
			Schema: &cms.Schema{
				ProjectID: itemsProject,
			},
		},
		Operator: cmswebhook.Operator{
			User: &cmswebhook.User{ID: "aaa"},
		},
	}
	assert.NoError(h(httptest.NewRequest("POST", "/", nil), payload))

	// assert logs
	assert.Contains(log(), "searchindex webhook: item: ")
	assert.Equal("searchindex webhook: skipped: all assets are not decompressed or no lod1 bldg", log())

	// assert storage
	storage := c.storage.FindAll(func(_ string, _ *cms.Item) bool { return true })
	assert.Len(storage, 1)
	storageitem := StorageItem{}
	storage[0].Unmarshal(&storageitem)
	assert.Equal(StorageItem{
		ID:    storage[0].ID,
		Item:  items[0].ID,
		Asset: []string{assets[0].ID},
	}, storageitem)

	// asset comments
	assert.Empty(util.DR(c.comments.Load(items[0].ID)))

	// asset.decompressed
	assets[0].ArchiveExtractionStatus = cms.AssetArchiveExtractionStatusDone
	c.assets.Store(assets[0].ID, assets[0])
	payload = &cmswebhook.Payload{
		Type:      cmswebhook.EventAssetDecompress,
		AssetData: (*cmswebhook.AssetData)(assets[0]),
		Operator: cmswebhook.Operator{
			Machine: &cmswebhook.Machine{},
		},
	}
	assert.NoError(h(httptest.NewRequest("POST", "/", nil), payload))

	// assert logs
	assert.Contains(log(), "searchindex webhook: item: ")
	assert.Equal("searchindex webhook: start processing", log())
	assert.Equal("searchindex webhook: start processing for "+assetName, log())
	assert.Equal("searchindex webhook: start processing for "+assetName2, log())
	assert.Equal("searchindex webhook: done", log())

	// assert storage
	assert.Equal(0, c.storage.Len())

	// assert item
	item2, _ := c.items.Load(items[0].ID)
	assert.Equal(StatusOK, ItemFrom(*item2).SearchIndexStatus)

	// asset comments
	assert.Equal(
		[]string{"検索インデックスの構築を開始しました。", "検索インデックスの構築が完了しました。"},
		util.DR(c.comments.Load(items[0].ID)),
	)
}

func TestWebhook_AssetNotDecompressed_DoubleUpdate(t *testing.T) {
	assert := assert.New(t)
	_ = initLogger(t)

	itemsProject := "prj"
	itemsModel := "itemitem"
	storageProject := "sys"
	assetName := "bldg_lod1"
	assetName2 := "bldg2_lod1"
	assets := []*cms.Asset{
		{
			ID:                      "bldg",
			URL:                     "https://example.com/" + assetName + ".zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: "",
		},
		{
			ID:                      "bldg2",
			URL:                     "https://example.com/" + assetName2 + ".zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: cms.AssetArchiveExtractionStatusDone,
		},
	}
	items := []*cms.Item{
		{
			ID: "item",
			Fields: Item{
				Bldg:              []string{assets[0].ID, assets[1].ID},
				SearchIndex:       nil,
				SearchIndexStatus: StatusReady,
			}.Fields(),
			ModelID: itemsModel,
		},
	}
	c := newMockedCMS(t, itemsProject, itemsModel, storageProject, storageModel, items, assets)
	h := webhookHandler(c, Config{
		CMSModel:          itemsModel,
		CMSStorageProject: storageProject,
		skipIndexer:       true,
	})

	// item.update
	payload := &cmswebhook.Payload{
		Type: cmswebhook.EventItemUpdate,
		ItemData: &cmswebhook.ItemData{
			Item: items[0],
			Model: &cms.Model{
				Key: itemsModel,
			},
			Schema: &cms.Schema{
				ProjectID: itemsProject,
			},
		},
		Operator: cmswebhook.Operator{
			User: &cmswebhook.User{ID: "aaa"},
		},
	}
	assert.NoError(h(httptest.NewRequest("POST", "/", nil), payload)) // first
	assert.NoError(h(httptest.NewRequest("POST", "/", nil), payload)) // second

	// assert storage
	storage := c.storage.FindAll(func(_ string, _ *cms.Item) bool { return true })
	assert.Len(storage, 1)
	storageitem := StorageItem{}
	storage[0].Unmarshal(&storageitem)
	assert.Equal(StorageItem{
		ID:    storage[0].ID,
		Item:  items[0].ID,
		Asset: []string{assets[0].ID},
	}, storageitem)
}

func TestWebhook_NoLod1Bldg(t *testing.T) {
	assert := assert.New(t)
	log := initLogger(t)

	itemsProject := "prj"
	itemsModel := "itemitem"
	storageProject := "sys"
	assetName := "bldg_lod2"
	assets := []*cms.Asset{
		{
			ID:                      "bldg",
			URL:                     "https://example.com/" + assetName + ".zip",
			ProjectID:               itemsProject,
			ArchiveExtractionStatus: cms.AssetArchiveExtractionStatusDone,
		},
	}
	items := []*cms.Item{
		{
			ID: "item",
			Fields: Item{
				Bldg:              []string{assets[0].ID},
				SearchIndex:       nil,
				SearchIndexStatus: StatusReady,
			}.Fields(),
			ModelID: itemsModel,
		},
	}
	c := newMockedCMS(t, itemsProject, itemsModel, storageProject, storageModel, items, assets)
	h := webhookHandler(c, Config{
		CMSModel:          itemsModel,
		CMSStorageProject: storageProject,
		skipIndexer:       true,
	})

	payload := &cmswebhook.Payload{
		Type: cmswebhook.EventItemUpdate,
		ItemData: &cmswebhook.ItemData{
			Item: items[0],
			Model: &cms.Model{
				Key: itemsModel,
			},
			Schema: &cms.Schema{
				ProjectID: itemsProject,
			},
		},
		Operator: cmswebhook.Operator{
			User: &cmswebhook.User{ID: "aaa"},
		},
	}

	assert.NoError(h(httptest.NewRequest("POST", "/", nil), payload))

	// assert logs
	assert.Contains(log(), "searchindex webhook: item: ")
	assert.Equal("searchindex webhook: skipped: all assets are not decompressed or no lod1 bldg", log())

	// assert item
	item2, _ := c.items.Load(items[0].ID)
	assert.Equal(StatusReady, ItemFrom(*item2).SearchIndexStatus)

	// asset comments
	assert.Empty(
		util.DR(c.comments.Load(items[0].ID)),
	)
}

func initLogger(t *testing.T) func() string {
	t.Helper()
	buf := bytes.NewBuffer(nil)
	// log.SetOutput(io.MultiWriter(log.DefaultOutput, buf))
	log.SetOutput(buf)
	t.Cleanup(func() { log.SetOutput(log.DefaultOutput) })

	scanner := bufio.NewScanner(buf)
	return func() string {
		if scanner.Scan() {
			t := scanner.Text()
			_, l, _ := strings.Cut(t, "\t")
			_, l, found := strings.Cut(l, "\t")
			if !found {
				return t
			}
			return l
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		return ""
	}
}

type mockedCMS struct {
	cms.Interface
	t                 *testing.T
	itemsprojectkey   string
	storageprojectkey string
	storagekey        string
	itemskey          string
	comments          *util.SyncMap[string, []string]
	storage           *util.SyncMap[string, *cms.Item]
	items             *util.SyncMap[string, *cms.Item]
	assets            *util.SyncMap[string, *cms.Asset]
}

func newMockedCMS(t *testing.T, itemsprojectkey, itemskey, storageprojectkey, storagekey string, items []*cms.Item, assets []*cms.Asset) *mockedCMS {
	return &mockedCMS{
		t:                 t,
		itemsprojectkey:   itemsprojectkey,
		storageprojectkey: storageprojectkey,
		storagekey:        storagekey,
		itemskey:          itemskey,
		comments:          util.SyncMapFrom[string, []string](nil),
		storage:           util.SyncMapFrom[string, *cms.Item](nil),
		items: util.SyncMapFrom(lo.SliceToMap(items, func(i *cms.Item) (string, *cms.Item) {
			return i.ID, i.Clone()
		})),
		assets: util.SyncMapFrom(lo.SliceToMap(assets, func(a *cms.Asset) (string, *cms.Asset) {
			return a.ID, a.Clone()
		})),
	}
}

func (c *mockedCMS) GetItem(ctx context.Context, itemID string, asset bool) (*cms.Item, error) {
	i, _ := c.l(itemID)
	if i == nil {
		return nil, rerror.ErrNotFound
	}
	return i, nil
}

func (c *mockedCMS) GetItemsByKey(ctx context.Context, projectIDOrAlias, modelIDOrKey string, asset bool) (*cms.Items, error) {
	m := c.m(projectIDOrAlias, modelIDOrKey)
	if m == nil {
		return nil, rerror.ErrNotFound
	}

	items := lo.Map(
		m.FindAll(func(_ string, _ *cms.Item) bool { return true }),
		func(i *cms.Item, _ int) cms.Item {
			return *i.Clone()
		},
	)

	return &cms.Items{
		Items:      items,
		Page:       1,
		PerPage:    len(items),
		TotalCount: len(items),
	}, nil
}

func (c *mockedCMS) CreateItemByKey(ctx context.Context, projectID, modelID string, fields []cms.Field) (*cms.Item, error) {
	m := c.m(projectID, modelID)
	if m == nil {
		return nil, rerror.ErrNotFound
	}

	id := randSeq(12)
	item := &cms.Item{
		ID:      id,
		ModelID: modelID,
		Fields: lo.Map(fields, func(f cms.Field, _ int) cms.Field {
			f.ID = randSeq(12)
			return f
		}),
	}
	m.Store(id, item)
	return item, nil
}

func (c *mockedCMS) UpdateItem(ctx context.Context, itemID string, fields []cms.Field) (*cms.Item, error) {
	i, m := c.l(itemID)
	if i == nil {
		return nil, rerror.ErrNotFound
	}
	newFields := make([]cms.Field, 0, len(i.Fields)+len(fields))
	for _, f := range fields {
		_, j, found := lo.FindIndexOf(i.Fields, func(g cms.Field) bool { return f.ID == g.ID })
		if !found {
			newFields = append(newFields, f)
		} else {
			i.Fields[j] = f
		}
	}
	i.Fields = newFields
	m.Store(itemID, i)
	return i, nil
}

func (c *mockedCMS) DeleteItem(ctx context.Context, itemID string) error {
	c.items.Delete(itemID)
	c.storage.Delete(itemID)
	return nil
}

func (c *mockedCMS) CommentToItem(ctx context.Context, itemID, comment string) error {
	i, _ := c.l(itemID)
	if i == nil {
		return rerror.ErrNotFound
	}
	comments, _ := c.comments.Load(itemID)
	comments = append(comments, comment)
	c.comments.Store(itemID, comments)
	return nil
}

func (c *mockedCMS) m(p, k string) *util.SyncMap[string, *cms.Item] {
	switch k {
	case c.itemskey:
		if p == c.itemsprojectkey {
			return c.items
		}
	case c.storagekey:
		if p == c.storageprojectkey {
			return c.storage
		}
	}
	return nil
}

func (c *mockedCMS) l(id string) (*cms.Item, *util.SyncMap[string, *cms.Item]) {
	l, ok := c.items.Load(id)
	if ok {
		return l.Clone(), c.items
	}

	l, ok = c.storage.Load(id)
	if ok {
		return l.Clone(), c.storage
	}

	return nil, nil
}

func (c *mockedCMS) Asset(ctx context.Context, id string) (*cms.Asset, error) {
	a, _ := c.assets.Load(id)
	if a == nil {
		return nil, rerror.ErrNotFound
	}
	return a, nil
}

func (c *mockedCMS) UploadAssetDirectly(ctx context.Context, projectID, name string, data io.Reader) (string, error) {
	if projectID != c.itemsprojectkey {
		return "", rerror.ErrNotFound
	}
	_, _ = io.Copy(io.Discard, data)
	a := &cms.Asset{
		ID:        randSeq(12),
		ProjectID: c.itemsprojectkey,
		URL:       "https://example.com",
	}
	c.assets.Store(a.ID, a)
	return a.ID, nil
}

func TestPathFileName(t *testing.T) {
	assert.Equal(t, "bbb", pathFileName("aaaa/bbb.txt"))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
