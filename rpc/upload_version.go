package rpc

import (
	"context"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/internal"
	"github.com/nadilas/moar/moarpb"
	"github.com/twitchtv/twirp"
)

func (s *Server) UploadVersion(ctx context.Context, request *moarpb.UploadVersionRequest) (*moarpb.UploadVersionResponse, error) {
	if request.Version == "" {
		return nil, twirp.RequiredArgumentError("version")
	}
	newVersion, err := semver.NewVersion(request.Version)
	if err != nil {
		return nil, twirp.InvalidArgumentError("version", err.Error())
	}
	module, err := s.registry.GetModule(ctx, request.ModuleName, false)
	if err != nil {
		s.logger.Warnf("Module not found: %s: %v", request.ModuleName, err)
		return nil, twirp.WrapError(twirp.NotFoundError("module not found: "+request.ModuleName), err)
	}

	if module.HasVersion(newVersion) {
		return nil, twirp.InvalidArgumentError("version", "upload not possible: version already exists")
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
		return nil, twirp.InternalErrorWith(err)
	}
	return &moarpb.UploadVersionResponse{}, nil
}
