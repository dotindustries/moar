package rpc

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/dotindustries/moar/internal/registry"
	"github.com/dotindustries/moar/moarpb/v1"
)

func (s *Server) DeleteVersion(ctx context.Context, c *connect.Request[moarpb.DeleteVersionRequest]) (*connect.Response[moarpb.DeleteVersionResponse], error) {
	request := c.Msg
	if request.Version == "" {
		return nil, requiredArgumentError("version")
	}
	if request.ModuleName == "" {
		return nil, requiredArgumentError("module_name")
	}

	version, err := semver.NewVersion(request.Version)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("version: %w", err))
	}
	module, err := s.registry.GetModule(ctx, request.ModuleName, false)
	if errors.Is(err, registry.ModuleNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module not found: "+request.ModuleName))
	} else if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if !module.HasVersion(version) {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module version not found: "+request.Version))
	}
	err = s.registry.DeleteVersion(ctx, module, version)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&moarpb.DeleteVersionResponse{}), nil
}
