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

	if len(module.Versions) == 0 {
		return nil, twirp.NotFoundError("module has no versions")
	}

	var version *semver.Version
	switch request.VersionSelector.(type) {
	case *moarpb.GetUrlRequest_Version:
		var err error
		versionString := request.GetVersion()
		if versionString == "latest" {
			break
		}
		version, err = semver.NewVersion(versionString)
		if err != nil {
			return nil, twirp.WrapError(twirp.InvalidArgumentError("version", "version invalid"), err)
		}
		if version != nil && !module.HasVersion(version) {
			return nil, twirp.NotFoundError("module version not found: " + request.ModuleName + "@" + version.String())
		}
		module.SetSelectedVersion(version)
	case *moarpb.GetUrlRequest_VersionConstraint:
		constraint, err := semver.NewConstraint(request.GetVersionConstraint())
		if err != nil {
			return nil, twirp.WrapError(twirp.InvalidArgumentError("versionConstraint", "constraint invalid"), err)
		}
		version = module.SelectVersion(constraint)
	default:
		s.logger.Debug("Module (%s) version is not specified in query, defaulting to latest")
	}

	uri, err := s.registry.UriForModule(ctx, module)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &moarpb.GetUrlResponse{
		Uri: uri,
		Module: &moarpb.Module{
			Name:     module.Name,
			Author:   module.Author,
			Language: module.Language,
			Versions: module.VersionStrings(),
		},
		SelectedVersion: version.String(),
	}, nil
}
