package sdk

import (
	"net/http"

	"github.com/eukarya-inc/reearth-plateauview/server/cms/cmswebhook"
	"github.com/eukarya-inc/reearth-plateauview/server/fme"
	"github.com/reearth/reearthx/log"
)

var (
	modelKey = "plateau"
)

func WebhookHandler(conf Config) (cmswebhook.Handler, error) {
	s, err := NewServices(conf)
	if err != nil {
		return nil, err
	}

	return func(req *http.Request, w *cmswebhook.Payload) error {
		if !w.Operator.IsUser() && w.Operator.IsIntegrationBy(conf.CMSIntegration) {
			log.Debugf("sdk webhook: invalid event operator: %+v", w.Operator)
			return nil
		}

		if w.Type != cmswebhook.EventItemCreate && w.Type != cmswebhook.EventItemUpdate {
			log.Debugf("sdk webhook: invalid event type: %s", w.Type)
			return nil
		}

		if w.ItemData == nil || w.ItemData.Item == nil || w.ItemData.Model == nil {
			log.Debugf("sdk webhook: invalid event data: %+v", w.Data)
			return nil
		}

		if w.ItemData.Model.Key != modelKey {
			log.Debugf("sdk webhook: invalid model id: %s, key: %s", w.ItemData.Item.ModelID, w.ItemData.Model.Key)
			return nil
		}

		item := ItemFrom(*w.ItemData.Item)

		if item.MaxLODStatus != "" && item.MaxLODStatus != StatusReady {
			log.Debugf("sdk webhook: skipped: %s", item.MaxLODStatus)
			return nil
		}

		if item.CityGML == "" {
			log.Debugf("sdk webhook: skipped: no citygml")
			return nil
		}

		log.Infof("sdk webhook: item: %+v", item)

		ctx := req.Context()
		citygml, err := s.CMS.Asset(ctx, item.CityGML)
		if err != nil {
			log.Errorf("sdk webhook: failed to get citygml asset: %s", err)
			return nil
		}

		if err := s.FME.Request(ctx, fme.MaxLODRequest{
			ID: fme.ID{
				ItemID:    item.ID,
				AssetID:   citygml.ID,
				ProjectID: w.ItemData.Schema.ProjectID,
			}.String(conf.Secret),
			Target: citygml.URL,
		}); err != nil {
			log.Errorf("sdk webhook: failed to send request to FME: %s", err)
			return nil
		}

		if _, err := s.CMS.UpdateItem(ctx, item.ID, Item{
			MaxLODStatus: StatusProcessing,
		}.Fields()); err != nil {
			log.Errorf("sdk webhook: failed to update item: %w", err)
		}

		log.Infof("sdk webhook: done")
		return nil
	}, nil
}
