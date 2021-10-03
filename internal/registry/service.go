package registry

import (
	"context"
	"errors"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/internal"
	"github.com/nadilas/moar/internal/storage"
)

type Reader interface {
	ModuleResources(ctx context.Context, module string, version string, data bool) ([]internal.File, error)
	GetModule(ctx context.Context, name string, loadData bool) (*internal.Module, error)
	Close() error
}

type Writer interface {
	PutModule(ctx context.Context, module internal.Module) error
	RemoveModule(ctx context.Context, name string) error
	PutVersion(ctx context.Context, module string, version string, files []internal.File) error
	RemoveVersion(ctx context.Context, module string, version string) error
}

type Storage interface {
	Reader
	Writer
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) NewModule(ctx context.Context, name, author, language string) error {
	m := internal.Module{
		Name:     name,
		Author:   author,
		Language: language,
		Versions: nil,
	}
	return s.storage.PutModule(ctx, m)
}

func (s *Service) DeleteModule(ctx context.Context, module *internal.Module) error {
	// remove all versions
	for len(module.Versions) > 0 {
		version := module.Versions[0]
		err := s.DeleteVersion(ctx, module, version.Version())
		if err != nil {
			return err
		}
	}
	return s.storage.RemoveModule(ctx, module.Name)
}

func (s *Service) UploadVersion(ctx context.Context, module *internal.Module, version *semver.Version, files []internal.File) error {
	// upload version
	err := s.storage.PutVersion(ctx, module.Name, version.String(), files)
	if err != nil {
		return err
	}
	// update module manifest to include new version
	newVersion := internal.NewVersion(version, files)
	module.Versions = append(module.Versions, newVersion)
	return s.storage.PutModule(ctx, *module)
}

func (s *Service) DeleteVersion(ctx context.Context, module *internal.Module, version *semver.Version) error {
	err := s.storage.RemoveVersion(ctx, module.Name, version.String())
	if err != nil {
		return err
	}

	// remove version from module and update manifest
	idx := -1
	for i, v := range module.Versions {
		if v.Version().Equal(version) {
			idx = i
		}
	}
	if idx < 0 {
		return VersionNotFound
	}
	module.Versions = append(module.Versions[:idx], module.Versions[idx+1:]...)
	return s.storage.PutModule(ctx, *module)
}

func (s *Service) GetModule(ctx context.Context, name string, loadData bool) (*internal.Module, error) {
	m, err := s.storage.GetModule(ctx, name, loadData)
	if errors.Is(err, storage.ModuleNotFound) {
		return nil, ModuleNotFound
	}
	return m, err
}

func (s *Service) Close() error {
	return s.storage.Close()
}
