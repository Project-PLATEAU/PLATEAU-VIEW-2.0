package app

import (
	"context"
	"fmt"
	"time"

	"github.com/reearth/reearth-cms/server/internal/infrastructure/auth0"
	"github.com/reearth/reearth-cms/server/internal/infrastructure/fs"
	"github.com/reearth/reearth-cms/server/internal/infrastructure/gcp"
	mongorepo "github.com/reearth/reearth-cms/server/internal/infrastructure/mongo"
	"github.com/reearth/reearth-cms/server/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/server/internal/usecase/repo"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/mongox"
	"github.com/spf13/afero"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

const databaseName = "reearth_cms"

func initReposAndGateways(ctx context.Context, conf *Config, debug bool) (*repo.Container, *gateway.Container) {
	gateways := &gateway.Container{}

	// Mongo
	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI(conf.DB).
			SetConnectTimeout(time.Second*10).
			SetMonitor(otelmongo.NewMonitor()),
	)
	if err != nil {
		log.Fatalf("repo initialization error: %+v\n", err)
	}

	repos, err := mongorepo.New(ctx, client, databaseName, mongox.IsTransactionAvailable(conf.DB))
	if err != nil {
		log.Fatalf("Failed to init mongo: %+v\n", err)
	}

	// File
	var fileRepo gateway.File
	if conf.GCS.BucketName == "" {
		log.Infoln("file: local storage is used")
		datafs := afero.NewBasePathFs(afero.NewOsFs(), "data")
		fileRepo, err = fs.NewFile(datafs, conf.AssetBaseURL)
	} else {
		log.Infof("file: GCS storage is used: %s", conf.GCS.BucketName)
		fileRepo, err = gcp.NewFile(conf.GCS.BucketName, conf.AssetBaseURL, conf.GCS.PublicationCacheControl)
		if err != nil {
			log.Fatalf("file: failed to init GCS storage: %s\n", err.Error())
		}
	}
	if err != nil {
		log.Fatalln(fmt.Sprintf("file: init error: %+v", err))
	}
	gateways.File = fileRepo

	// Auth0
	gateways.Authenticator = auth0.New(conf.Auth0.Domain, conf.Auth0.ClientID, conf.Auth0.ClientSecret)

	// CloudTasks
	if conf.Task.GCPProject != "" && conf.Task.GCPRegion != "" || conf.Task.QueueName != "" {
		conf.Task.GCSHost = conf.Host
		taskRunner, err := gcp.NewTaskRunner(ctx, &conf.Task)
		if err != nil {
			log.Fatalln(fmt.Sprintf("task runner: init error: %+v", err))
		}
		gateways.TaskRunner = taskRunner
	} else {
		log.Infof("task runner: not used")
	}

	return repos, gateways
}
