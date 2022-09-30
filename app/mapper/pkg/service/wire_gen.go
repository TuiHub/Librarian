// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/app/mapper/internal/data"
	"github.com/tuihub/librarian/app/mapper/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/protos/pkg/librarian/mapper/v1"
)

// Injectors from wire.go:

func NewMapperService(mapper_Data *conf.Mapper_Data) (v1.LibrarianMapperServiceServer, func(), error) {
	db, cleanup := data.NewNebula(mapper_Data)
	handle, cleanup2 := data.NewCayley(mapper_Data)
	mapperRepo, err := data.NewMapperRepo(db, handle)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mapperUseCase := biz.NewMapperUseCase(mapperRepo)
	librarianMapperServiceServer := service.NewLibrarianMapperServiceService(mapperUseCase)
	return librarianMapperServiceServer, func() {
		cleanup2()
		cleanup()
	}, nil
}
