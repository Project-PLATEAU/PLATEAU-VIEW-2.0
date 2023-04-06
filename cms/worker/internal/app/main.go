package app

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"golang.org/x/net/http2"

	rhttp "github.com/reearth/reearth-cms/worker/internal/adapter/http"
	rmongo "github.com/reearth/reearth-cms/worker/internal/infrastructure/mongo"
	"github.com/reearth/reearth-cms/worker/internal/usecase/gateway"
	"github.com/reearth/reearth-cms/worker/internal/usecase/interactor"
	"github.com/reearth/reearthx/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start(debug bool, version string) {
	log.Infof("reearth-cms/worker %s", version)

	ctx := context.Background()

	// Load config
	conf, cerr := ReadConfig(debug)
	if cerr != nil {
		log.Fatal(cerr)
	}

	// mongo
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
	mongoWebhook := rmongo.NewWebhook(client.Database("reearth_cms"))
	lo.Must0(mongoWebhook.InitIndex(ctx))

	// gateways
	gateways := initReposAndGateways(ctx, conf, debug)

	// usecase
	uc := interactor.NewUsecase(gateways, mongoWebhook)
	ctrl := rhttp.NewController(uc)
	handler := NewHandler(ctrl)

	// start web server
	NewServer(ctx, &ServerConfig{
		Config:   conf,
		Debug:    debug,
		Gateways: gateways,
	}, handler).Run()
}

type WebServer struct {
	address   string
	appServer *echo.Echo
}

type ServerConfig struct {
	Config   *Config
	Debug    bool
	Gateways *gateway.Container
}

func NewServer(ctx context.Context, cfg *ServerConfig, handler *Handler) *WebServer {
	port := cfg.Config.Port
	if port == "" {
		port = "8080"
	}

	host := cfg.Config.ServerHost
	if host == "" {
		if cfg.Debug {
			host = "localhost"
		} else {
			host = "0.0.0.0"
		}
	}
	address := host + ":" + port

	w := &WebServer{
		address: address,
	}

	w.appServer = initEcho(ctx, cfg, handler)
	return w
}

func (w *WebServer) Run() {
	defer log.Infoln("Server shutdown")

	debugLog := ""
	if w.appServer.Debug {
		debugLog += " with debug mode"
	}
	log.Infof("server started%s at http://%s\n", debugLog, w.address)

	go func() {
		err := w.appServer.StartH2CServer(w.address, &http2.Server{})
		log.Fatalln(err.Error())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
