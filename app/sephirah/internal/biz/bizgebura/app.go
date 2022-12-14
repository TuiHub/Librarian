package bizgebura

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/logger"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (g *Gebura) CreateApp(ctx context.Context, app *App) (*App, *errors.Error) {
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	app.InternalID = resp.Id
	app.Source = AppSourceInternal
	app.SourceAppID = ""
	app.SourceURL = ""
	if _, err = g.mapper.InsertVertex(ctx, &mapper.InsertVertexRequest{
		VertexList: []*mapper.Vertex{{
			Vid:  app.InternalID,
			Type: mapper.VertexType_VERTEX_TYPE_ABSTRACT,
		}},
	}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if err = g.repo.CreateApp(ctx, app); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return app, nil
}

func (g *Gebura) UpdateApp(ctx context.Context, app *App) *errors.Error {
	app.Source = AppSourceInternal
	err := g.repo.UpdateApp(ctx, app)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (g *Gebura) UpsertApp(ctx context.Context, app []*App) ([]*App, *errors.Error) {
	for _, a := range app {
		resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
		if err != nil {
			logger.Infof("NewID failed: %s", err.Error())
			return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
		}
		a.InternalID = resp.Id
	}
	err := g.repo.UpsertApp(ctx, app)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return app, nil
}

func (g *Gebura) ListApp(
	ctx context.Context,
	paging Paging,
	sources []AppSource,
	types []AppType,
	ids []int64,
	containDetails bool,
) ([]*App, *errors.Error) {
	apps, err := g.repo.ListApp(ctx, paging, sources, types, ids, containDetails)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}

func (g *Gebura) BindApp(ctx context.Context, internal App, bind App) (*App, *errors.Error) {
	resp, err := g.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		logger.Infof("NewID failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	bind.InternalID = resp.Id
	if err = g.repo.UpsertApp(ctx, []*App{&bind}); err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return &bind, nil
}

func (g *Gebura) ListBindApp(ctx context.Context, id int64) ([]*App, *errors.Error) {
	app, err := g.repo.ListApp(ctx, Paging{
		PageSize: 1,
		PageNum:  1,
	}, nil, nil, []int64{id}, false)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if len(app) != 1 {
		return nil, pb.ErrorErrorReasonBadRequest("No such app")
	}
	resp, err := g.mapper.FetchEqualVertex(ctx, &mapper.FetchEqualVertexRequest{SrcVid: id})
	if err != nil {
		logger.Infof("Fetch Equal Vertex failed: %s", err.Error())
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	appids := make([]int64, len(resp.GetVertexList()))
	for i, v := range resp.GetVertexList() {
		appids[i] = v.GetVid()
	}
	apps, err := g.repo.ListApp(ctx, Paging{
		PageSize: 99, //nolint:gomnd //TODO
		PageNum:  1,
	}, nil, nil, appids, true)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return apps, nil
}
