package rpc

import (
	"context"
	"errors"

	"github.com/dotindustries/moar/internal/registry"
	"github.com/dotindustries/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) DeleteModule(ctx context.Context, request *moarpb.DeleteModuleRequest) (*moarpb.DeleteModuleResponse, error) {
	if request.ModuleName == "" {
		return nil, twirp.RequiredArgumentError("moduleName")
	}

	module, err := s.registry.GetModule(ctx, request.ModuleName, false)
	if errors.Is(err, registry.ModuleNotFound) {
		return nil, twirp.NotFoundError("module not found: " + request.ModuleName)
	} else if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	err = s.registry.DeleteModule(ctx, module)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &moarpb.DeleteModuleResponse{}, nil
}
