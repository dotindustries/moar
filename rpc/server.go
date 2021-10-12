package rpc

import (
	"context"

	"github.com/Masterminds/semver"
	"github.com/dotindustries/moar/internal"
	"github.com/sirupsen/logrus"
)

type RegistryReader interface {
	GetModule(ctx context.Context, name string, loadData bool) (*internal.Module, error)
	Close() error
}

type RegistryWriter interface {
	NewModule(ctx context.Context, name string, author string, language string) error
	DeleteModule(ctx context.Context, module *internal.Module) error
	UploadVersion(ctx context.Context, module *internal.Module, version *semver.Version, files []internal.File) error
	DeleteVersion(ctx context.Context, module *internal.Module, version *semver.Version) error
}

type ModuleRegistry interface {
	RegistryReader
	RegistryWriter
}

type Server struct {
	registry                ModuleRegistry
	reverseProxy            string
	logger                  *logrus.Entry
	versionOverwriteEnabled bool
}

type Opts struct {
	VersionOverwriteEnabled bool
}

func NewServer(registry ModuleRegistry, reverseProxy string, opts Opts) *Server {
	return &Server{
		registry:                registry,
		logger:                  logrus.WithField("op", "server"),
		reverseProxy:            reverseProxy,
		versionOverwriteEnabled: opts.VersionOverwriteEnabled,
	}
}

func (s *Server) Shutdown() {
	err := s.registry.Close()
	if err != nil {
		panic(err)
	}
}
