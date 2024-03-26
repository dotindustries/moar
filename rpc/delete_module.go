package rpc

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"

	"github.com/dotindustries/moar/internal/registry"
	"github.com/dotindustries/moar/moarpb/v1"
)

func (s *Server) DeleteModule(ctx context.Context, c *connect.Request[moarpb.DeleteModuleRequest]) (*connect.Response[moarpb.DeleteModuleResponse], error) {
	request := c.Msg
	if request.ModuleName == "" {
		return nil, requiredArgumentError("moduleName")
	}

	module, err := s.registry.GetModule(ctx, request.ModuleName, false)
	if errors.Is(err, registry.ModuleNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module not found: "+request.ModuleName))
	} else if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	err = s.registry.DeleteModule(ctx, module)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&moarpb.DeleteModuleResponse{}), nil
}
