// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/server"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(sephirah_Server *conf.Sephirah_Server, sephirah_Data *conf.Sephirah_Data, auth *conf.Auth) (*kratos.App, func(), error) {
	entClient, cleanup, err := data.NewSQLClient(sephirah_Data)
	if err != nil {
		return nil, nil, err
	}
	dataData := data.NewData(entClient)
	tipherethRepo := data.NewTipherethRepo(dataData)
	libauthAuth, err := libauth.NewAuth(auth)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianMapperServiceClient, err := client.NewMapperClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianSearcherServiceClient, err := client.NewSearcherClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	librarianPorterServiceClient, err := client.NewPorterClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tipherethUsecase := biz.NewTipherethUsecase(tipherethRepo, libauthAuth, librarianMapperServiceClient, librarianSearcherServiceClient, librarianPorterServiceClient)
	librarianSephirahServiceServer := service.NewLibrarianSephirahServiceService(tipherethUsecase)
	grpcServer := server.NewGRPCServer(sephirah_Server, librarianSephirahServiceServer)
	app := newApp(grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
