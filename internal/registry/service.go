package registry

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/internal"
	"github.com/nadilas/moar/internal/storage"
)

type Reader interface {
	UriForModule(ctx context.Context, module string, version string) (string, error)
	GetModule(ctx context.Context, name string) (*internal.Module, error)
	Close() error
}

type Writer interface {
	PutModule(ctx context.Context, module *internal.Module) error
	RemoveModule(ctx context.Context, name string) error
	PutVersion(ctx context.Context, module string, version string, data []byte) error
	RemoveVersion(ctx context.Context, module string, version string) error
}

type Storage interface {
	Reader
	Writer
}

type Service struct {
	storage      Storage
	reverseProxy string
}

func New(storage Storage, reverseProxy string) *Service {
	return &Service{storage: storage, reverseProxy: reverseProxy}
}

func (s *Service) NewModule(ctx context.Context, name string, author string) error {
	m := &internal.Module{
		Name:     name,
		Author:   author,
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

func (s *Service) UploadVersion(ctx context.Context, module *internal.Module, version *semver.Version, data []byte) error {
	// upload version
	err := s.storage.PutVersion(ctx, module.Name, version.String(), data)
	if err != nil {
		return err
	}
	// update module manifest to include new version
	module.Versions = append(module.Versions, internal.NewVersion(version))
	return s.storage.PutModule(ctx, module)
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
	return s.storage.PutModule(ctx, module)
}

// UriForModule finds the public URI for a module. If a version has been previously selected, that version is used, otherwise
// this function defaults to the latest version of the module
func (s *Service) UriForModule(ctx context.Context, module *internal.Module) (string, error) {
	v := module.SelectedVersion()
	if v == nil {
		v = module.Latest()
	}
	uri, err := s.storage.UriForModule(ctx, module.Name, v.String())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", s.reverseProxy, uri), nil
}

func (s *Service) GetModule(ctx context.Context, name string) (*internal.Module, error) {
	m, err := s.storage.GetModule(ctx, name)
	if errors.Is(err, storage.ModuleNotFound) {
		return nil, ModuleNotFound
	}
	return m, err
}

func (s *Service) Close() error {
	return s.storage.Close()
}
