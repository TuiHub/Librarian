package main

import (
	"flag"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libzap"

	"github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func newApp(gs *grpc.Server, hs *http.Server, mq *libmq.MQ) *kratos.App {
	metadata := libapp.GetAppMetadata()
	return kratos.New(
		kratos.ID(metadata.ID),
		kratos.Name(metadata.Name),
		kratos.Version(metadata.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, hs, mq),
	)
}

func main() {
	// flagconf is the config flag.
	var flagconf string
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.Parse()
	metadata := libapp.GetAppMetadata()
	logger := log.With(zap.NewLogger(libzap.NewDefaultLogger()),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", metadata.ID,
		"service.name", metadata.Name,
		"service.version", metadata.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(logger)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Librarian
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(
		bc.Sephirah.Server,
		bc.Sephirah.Data,
		bc.Mapper.Data,
		bc.Searcher.Data,
		bc.Porter.Data,
		bc.Sephirah.Auth,
	)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
