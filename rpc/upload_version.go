package rpc

import (
	"context"

	"github.com/Masterminds/semver"
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
	module, err := s.registry.GetModule(ctx, request.ModuleName)
	if err != nil {
		s.logger.Warnf("Module not found: %s: %v", request.ModuleName, err)
		return nil, twirp.WrapError(twirp.NotFoundError("module not found: "+request.ModuleName), err)
	}

	if module.HasVersion(newVersion) {
		return nil, twirp.InvalidArgumentError("version", "upload not possible: version already exists")
	}

	err = s.registry.UploadVersion(ctx, module, newVersion, request.FileData, request.StyleData)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	return &moarpb.UploadVersionResponse{}, nil
}
