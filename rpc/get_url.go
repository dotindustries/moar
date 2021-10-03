package rpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/internal"
	"github.com/nadilas/moar/internal/registry"
	"github.com/nadilas/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) GetUrl(ctx context.Context, request *moarpb.GetUrlRequest) (*moarpb.GetUrlResponse, error) {
	if request.ModuleName == "" {
		return nil, twirp.RequiredArgumentError("moduleName")
	}

	// 1. check if module exists in registry
	module, err := s.registry.GetModule(ctx, request.ModuleName, false)
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

	var version *internal.Version
	selectedVersionString := ""
	switch request.VersionSelector.(type) {
	case *moarpb.GetUrlRequest_Version:
		var err error
		versionString := request.GetVersion()
		if versionString == "latest" {
			selectedVersionString = "latest"
			version = module.Latest()
			break
		}
		v, err := semver.NewVersion(versionString)
		if err != nil {
			return nil, twirp.WrapError(twirp.InvalidArgumentError("version", "version invalid"), err)
		}
		if v != nil && !module.HasVersion(v) {
			return nil, twirp.NotFoundError("module version not found: " + request.ModuleName + "@" + v.String())
		}
		version = module.SetSelectedVersion(v)
	case *moarpb.GetUrlRequest_VersionConstraint:
		constraint, err := semver.NewConstraint(request.GetVersionConstraint())
		if err != nil {
			return nil, twirp.WrapError(twirp.InvalidArgumentError("versionConstraint", "constraint invalid"), err)
		}
		version = module.SelectVersion(constraint)
		selectedVersionString = version.Version().String()
	default:
		version = module.Latest()
		selectedVersionString = "latest"
		s.logger.Debugf("Module (%s) version is not specified in query, defaulting to latest", request.ModuleName)
	}

	var resources []*moarpb.VersionResource
	for _, file := range version.Files {
		resources = append(resources, &moarpb.VersionResource{
			Uri:         fmt.Sprintf("%s/%s", s.reverseProxy, file.Uri),
			Name:        file.Name,
			ContentType: file.MimeType,
		})
	}
	mod := moduleToDto(module)

	return &moarpb.GetUrlResponse{
		Resources:       resources,
		Module:          mod,
		SelectedVersion: selectedVersionString,
	}, nil
}

func moduleToDto(module *internal.Module) *moarpb.Module {
	var versions []*moarpb.Version
	for _, v := range module.Versions {
		var files []*moarpb.File
		for _, file := range v.Files {
			files = append(files, &moarpb.File{
				Name:     file.Name,
				MimeType: file.MimeType,
				Data:     file.Data,
			})
		}
		versions = append(versions, &moarpb.Version{
			Name:  v.Value,
			Files: files,
		})
	}
	mod := &moarpb.Module{
		Name:     module.Name,
		Author:   module.Author,
		Language: module.Language,
		Versions: versions,
	}
	return mod
}
