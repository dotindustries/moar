package rpc

import (
	"context"
	"errors"

	"github.com/nadilas/moar/internal/registry"
	"github.com/nadilas/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) GetModule(ctx context.Context, request *moarpb.GetModuleRequest) (*moarpb.GetModuleResponse, error) {
	if request.ModuleName == "" {
		return nil, twirp.RequiredArgumentError("moduleName")
	}
	module, err := s.registry.GetModule(ctx, request.ModuleName)
	if errors.Is(err, registry.ModuleNotFound) {
		return nil, twirp.NotFoundError("module not found: " + request.ModuleName)
	} else if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	mod := moduleToDto(module)

	return &moarpb.GetModuleResponse{Module: mod}, nil
}
