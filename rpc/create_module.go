package rpc

import (
	"context"
	"errors"

	"github.com/nadilas/moar/internal/registry"
	"github.com/nadilas/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) CreateModule(ctx context.Context, request *moarpb.CreateModuleRequest) (*moarpb.CreateModuleResponse, error) {
	_, err := s.registry.GetModule(ctx, request.ModuleName)
	if err == nil {
		return nil, twirp.NewError(twirp.AlreadyExists, request.ModuleName)
	} else if err != nil && !errors.Is(err, registry.ModuleNotFound) {
		return nil, twirp.InternalErrorWith(err)
	}

	if request.ModuleName == "" {
		return nil, twirp.RequiredArgumentError("moduleName")
	}

	err = s.registry.NewModule(ctx, request.ModuleName, request.Author, request.Language)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &moarpb.CreateModuleResponse{}, nil
}
