package rpc

import (
	"context"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/internal"
	"github.com/sirupsen/logrus"
)

type RegistryReader interface {
	GetModule(ctx context.Context, name string) (*internal.Module, error)
	UriForModule(ctx context.Context, module *internal.Module) (*internal.VersionResources, error)
	Close() error
}

type RegistryWriter interface {
	NewModule(ctx context.Context, name string, author string, language string) error
	DeleteModule(ctx context.Context, module *internal.Module) error
	UploadVersion(ctx context.Context, module *internal.Module, version *semver.Version, data []byte, styleData []byte) error
	DeleteVersion(ctx context.Context, module *internal.Module, version *semver.Version) error
}

type ModuleRegistry interface {
	RegistryReader
	RegistryWriter
}

type Server struct {
	registry ModuleRegistry
	logger   *logrus.Entry
}

func NewServer(registry ModuleRegistry) *Server {
	return &Server{registry: registry, logger: logrus.WithField("op", "server")}
}

func (s *Server) Shutdown() {
	err := s.registry.Close()
	if err != nil {
		panic(err)
	}
}
