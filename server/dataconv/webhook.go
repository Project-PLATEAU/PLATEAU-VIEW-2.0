package dataconv

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/eukarya-inc/reearth-plateauview/server/cms/cmswebhook"
	geojson "github.com/paulmach/go.geojson"
	"github.com/reearth/reearthx/log"
	"github.com/spkg/bom"
)

const defaultCMSModel = "dataset"

type Config struct {
	Disable  bool
	CMSBase  string
	CMSToken string
	// optional
	CMSModel string
}

type Item struct {
	ID         string   `json:"id" cms:"id"`
	Type       string   `json:"type" cms:"type"`
	DataFormat string   `json:"data_format" cms:"data_format,select"`
	Data       string   `json:"data" cms:"data,asset"`
	DataConv   string   `json:"data_conv" cms:"data_conv,select"`
	DataOrig   []string `json:"data_orig" cms:"data_orig,asset"`
}

func (i Item) Fields() []cms.Field {
	i2 := &cms.Item{}
	cms.Marshal(i, i2)
	return i2.Fields
}

func WebhookHandler(conf Config) (cmswebhook.Handler, error) {
	if conf.Disable {
		return nil, nil
	}

	c, err := cms.New(conf.CMSBase, conf.CMSToken)
	if err != nil {
		return nil, err
	}

	return func(req *http.Request, w *cmswebhook.Payload) error {
		return webhookHandler(req.Context(), w, conf, c)
	}, nil
}

func webhookHandler(ctx context.Context, w *cmswebhook.Payload, conf Config, c cms.Interface) error {
	if conf.CMSModel == "" {
		conf.CMSModel = defaultCMSModel
	}

	pid := w.ProjectID()
	if w.Type != cmswebhook.EventItemCreate && w.Type != cmswebhook.EventItemUpdate ||
		pid == "" ||
		w.ItemData.Item == nil ||
		w.ItemData.Model == nil ||
		w.ItemData.Model.Key != conf.CMSModel ||
		w.Operator.User == nil {
		var key string
		if w.ItemData.Model != nil {
			key = w.ItemData.Model.Key
		}
		// skipped
		log.Debugf("dataconv: skipped: invalid webhook: type=%s, projectid=%s, model=%s", w.Type, pid, key)
		return nil
	}

	var i Item
	w.ItemData.Item.Unmarshal(&i)
	landmark := strings.Contains(i.Type, "ランドマーク") || strings.Contains(i.Type, "鉄道駅")
	border := strings.Contains(i.Type, "行政界")
	if i.DataConv == "変換しない" || i.Data == "" || (!landmark && !border) {
		log.Debugf("dataconv: skipped: invalid item: %#v", i)
		return nil
	}

	a, err := c.Asset(ctx, i.Data)
	if err != nil || a.URL == "" {
		log.Debugf("dataconv: skipped: failed to load asset: %s", i.Data)
		return nil
	}

	u, err := url.Parse(a.URL)
	if b := fileName(u.Path); err != nil ||
		path.Ext(u.Path) != ".geojson" ||
		border && !strings.HasSuffix(b, "_border") ||
		landmark && !strings.HasSuffix(b, "_landmark") && !strings.HasSuffix(b, "_station") {
		log.Debugf("dataconv: skipped: invalid URL or ext is not geojson: %s", u)
		return nil
	}

	g, err := getGeoJSON(ctx, a.URL)
	if err != nil {
		return nil
	}

	id := strings.TrimSuffix(path.Base(u.Path), path.Ext(u.Path))
	var res any
	if landmark {
		res, err = ConvertLandmark(g, id)
	} else {
		res, err = ConvertBorder(g, id)
	}
	if err != nil {
		log.Errorf("dataconv: failed to convert %s: %v", id, err)
		return nil
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Errorf("dataconv: failed to marshal result (%s): %v", id, err)
		return nil
	}

	aid, err := c.UploadAssetDirectly(ctx, pid, id+".czml", bytes.NewReader(b))
	if err != nil {
		log.Errorf("dataconv: failed to upload asset (%s): %v", id, err)
		return nil
	}

	if _, err := c.UpdateItem(ctx, i.ID, Item{
		Data:       aid,
		DataFormat: "CZML",
		DataOrig:   []string{a.ID},
	}.Fields()); err != nil {
		log.Errorf("dataconv: failed to update item (%s): %v", id, err)
		return nil
	}

	log.Infof("dataconv: done: %s, %s", i.ID, id)

	return nil
}

func getGeoJSON(ctx context.Context, u string) (*geojson.FeatureCollection, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		log.Errorf("dataconv: failed to create a request: %v", err)
		return nil, nil
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("dataconv: failed to get asset: %v", err)
		return nil, nil
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		log.Errorf("dataconv: failed to get asset: status code is %d", res.StatusCode)
		return nil, nil
	}

	f := geojson.FeatureCollection{}
	if err := json.NewDecoder(bom.NewReader(res.Body)).Decode(&f); err != nil {
		log.Errorf("dataconv: invalid geojson: %v", err)
	}

	return &f, nil
}

func fileName(p string) string {
	return strings.TrimSuffix(path.Base(p), path.Ext(p))
}
