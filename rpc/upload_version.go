package rpc

import (
	"connectrpc.com/connect"
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/dotindustries/moar/internal"
	moarpb "github.com/dotindustries/moar/moarpb/v1"
)

func (s *Server) UploadVersion(ctx context.Context, c *connect.Request[moarpb.UploadVersionRequest]) (*connect.Response[moarpb.UploadVersionResponse], error) {
	request := c.Msg
	if request.Version == "" {
		return nil, requiredArgumentError("version")
	}
	newVersion, err := semver.NewVersion(request.Version)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("version: %w", err))
	}
	module, err := s.registry.GetModule(ctx, request.ModuleName, false)
	if err != nil {
		s.logger.Warnf("Module not found: %s: %v", request.ModuleName, err)
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("module not found '%s': %w", request.ModuleName, err))
	}

	if module.HasVersion(newVersion) && !s.versionOverwriteEnabled {
		return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("version: overwrite disabled: version already exists"))
	}
	var files []internal.File
	for _, file := range request.Files {
		files = append(files, internal.File{
			Name:     file.Name,
			MimeType: file.MimeType,
			Data:     file.Data,
		})
	}
	err = s.registry.UploadVersion(ctx, module, newVersion, files)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&moarpb.UploadVersionResponse{}), nil
}
