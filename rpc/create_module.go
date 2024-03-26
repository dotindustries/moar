package rpc

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"

	"github.com/dotindustries/moar/internal/registry"
	"github.com/dotindustries/moar/moarpb/v1"
)

func (s *Server) CreateModule(ctx context.Context, c *connect.Request[moarpb.CreateModuleRequest]) (*connect.Response[moarpb.CreateModuleResponse], error) {
	request := c.Msg
	_, err := s.registry.GetModule(ctx, request.ModuleName, false)
	if err == nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("module already exists: %s", request.ModuleName))
	} else if err != nil && !errors.Is(err, registry.ModuleNotFound) {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if request.ModuleName == "" {
		return nil, requiredArgumentError("module_name")
	}

	err = s.registry.NewModule(ctx, request.ModuleName, request.Author, request.Language)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&moarpb.CreateModuleResponse{}), nil
}
