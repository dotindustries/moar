package rpc

import (
	"context"
	"errors"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/internal/registry"
	"github.com/nadilas/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) GetUrl(ctx context.Context, request *moarpb.GetUrlRequest) (*moarpb.GetUrlResponse, error) {
	if request.ModuleName == "" {
		return nil, twirp.RequiredArgumentError("moduleName")
	}

	// 1. check if module exists in registry
	module, err := s.registry.GetModule(ctx, request.ModuleName)
	if err != nil {
		if errors.Is(err, registry.ModuleNotFound) {
			return nil, twirp.NotFoundError("module not found: " + request.ModuleName)
		}
		s.logger.Warnf("Module not found: %s: %v", request.ModuleName, err)
		return nil, twirp.InternalError(err.Error())
	}

	var version *semver.Version
	switch request.VersionSelector.(type) {
	case *moarpb.GetUrlRequest_Version:
		// requesting specific version
		var err error
		version, err = semver.NewVersion(request.GetVersion())
		if err != nil {
			return nil, twirp.WrapError(twirp.InvalidArgumentError("version", "version invalid"), err)
		}
	case *moarpb.GetUrlRequest_VersionConstraint:
		constraint, err := semver.NewConstraint(request.GetVersionConstraint())
		if err != nil {
			return nil, twirp.WrapError(twirp.InvalidArgumentError("versionConstraint", "constraint invalid"), err)
		}
		version = module.SelectVersion(constraint)
	default:
		return nil, twirp.RequiredArgumentError("versionSelector")
	}

	if version == nil || len(module.Versions) == 0 {
		return nil, twirp.NotFoundError("module has no versions")
	}

	// 2. check if version exists in registry
	if !module.HasVersion(version) {
		return nil, twirp.NotFoundError("module version not found: " + request.ModuleName + "@" + version.String())
	}

	// 3. build GET url from storage service
	uri, err := s.registry.UriForModule(ctx, module)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &moarpb.GetUrlResponse{
		Uri: uri,
		Module: &moarpb.Module{
			Name:     module.Name,
			Versions: module.VersionStrings(),
		},
		SelectedVersion: version.String(),
	}, nil
}
