package rpc

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"

	"github.com/dotindustries/moar/internal/registry"
	"github.com/dotindustries/moar/moarpb/v1"
)

func (s *Server) GetModule(ctx context.Context, c *connect.Request[moarpb.GetModuleRequest]) (*connect.Response[moarpb.GetModuleResponse], error) {
	request := c.Msg
	var modules []*moarpb.Module

	if request.ModuleName == "" {
		mods, err := s.registry.GetAllModules(ctx, false)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		for _, module := range mods {
			modules = append(modules, moduleToDto(module))
		}
	} else {
		module, err := s.registry.GetModule(ctx, request.ModuleName, false)
		if errors.Is(err, registry.ModuleNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module not found: "+request.ModuleName))
		} else if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		modules = append(modules, moduleToDto(module))
	}

	return connect.NewResponse(&moarpb.GetModuleResponse{Module: modules}), nil
}
