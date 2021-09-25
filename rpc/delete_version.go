package rpc

import (
	"context"
	"errors"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/internal/registry"
	"github.com/nadilas/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) DeleteVersion(ctx context.Context, request *moarpb.DeleteVersionRequest) (*moarpb.DeleteVersionResponse, error) {
	if request.Version == "" {
		return nil, twirp.RequiredArgumentError("version")
	}
	if request.ModuleName == "" {
		return nil, twirp.RequiredArgumentError("moduleName")
	}

	version, err := semver.NewVersion(request.Version)
	if err != nil {
		return nil, twirp.InvalidArgumentError("version", err.Error())
	}
	module, err := s.registry.GetModule(ctx, request.ModuleName)
	if errors.Is(err, registry.ModuleNotFound) {
		return nil, twirp.NotFoundError("module not found: " + request.ModuleName)
	} else if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	if !module.HasVersion(version) {
		return nil, twirp.NotFoundError("module version not found: " + request.Version)
	}
	err = s.registry.DeleteVersion(ctx, module, version)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &moarpb.DeleteVersionResponse{}, nil
}
