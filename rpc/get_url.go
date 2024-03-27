package rpc

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/dotindustries/moar/internal"
	"github.com/dotindustries/moar/internal/registry"
	moarpb "github.com/dotindustries/moar/moarpb/v1"
)

func (s *Server) GetUrl(ctx context.Context, c *connect.Request[moarpb.GetUrlRequest]) (*connect.Response[moarpb.GetUrlResponse], error) {
	request := c.Msg
	if request.ModuleName == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid argument: %s", "module_name"))
	}

	// 1. check if module exists in registry
	module, err := s.registry.GetModule(ctx, request.ModuleName, false)
	if err != nil {
		if errors.Is(err, registry.ModuleNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module not found: "+request.ModuleName))
		}
		s.logger.Warnf("Module not found: %s: %v", request.ModuleName, err)
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if len(module.Versions) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module has no versions"))
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
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
		if v != nil && !module.HasVersion(v) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module version not found: "+request.ModuleName+"@"+v.String()))
		}
		version = module.SetSelectedVersion(v)
	case *moarpb.GetUrlRequest_VersionConstraint:
		constraint, err := semver.NewConstraint(request.GetVersionConstraint())
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
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

	return connect.NewResponse(&moarpb.GetUrlResponse{
		Resources:       resources,
		Module:          mod,
		SelectedVersion: selectedVersionString,
	}), nil
}
