package app

import (
	"context"

	"github.com/reearth/reearth-cms/worker/internal/infrastructure/gcp"
	"github.com/reearth/reearth-cms/worker/internal/usecase/gateway"
	"github.com/reearth/reearthx/log"
)

func initReposAndGateways(ctx context.Context, conf *Config, debug bool) *gateway.Container {
	gateways := &gateway.Container{
		CMS: gcp.NewPubSub(conf.PubSub.Topic, conf.GCP.Project),
	}

	if conf.GCS.BucketName != "" {
		log.Infof("file: GCS storage is used: %s\n", conf.GCS.BucketName)
		fileRepo, err := gcp.NewFile(conf.GCS.BucketName, conf.GCS.AssetBaseURL, conf.GCS.PublicationCacheControl)
		if err != nil {
			if debug {
				log.Warnf("file: failed to init GCS storage: %s\n", err.Error())
				err = nil
			}
		}
		gateways.File = fileRepo
	}

	return gateways
}
