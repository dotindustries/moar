package rpc

import (
	"context"
	"errors"

	"github.com/dotindustries/moar/internal/registry"
	"github.com/dotindustries/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) GetModule(ctx context.Context, request *moarpb.GetModuleRequest) (*moarpb.GetModuleResponse, error) {
	var modules []*moarpb.Module

	if request.ModuleName == "" {
		mods, err := s.registry.GetAllModules(ctx, false)
		if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
		for _, module := range mods {
			modules = append(modules, moduleToDto(module))
		}
	} else {
		module, err := s.registry.GetModule(ctx, request.ModuleName, false)
		if errors.Is(err, registry.ModuleNotFound) {
			return nil, twirp.NotFoundError("module not found: " + request.ModuleName)
		} else if err != nil {
			return nil, twirp.InternalErrorWith(err)
		}
		modules = append(modules, moduleToDto(module))
	}

	return &moarpb.GetModuleResponse{Module: modules}, nil
}
