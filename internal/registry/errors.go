package registry

import "errors"

var (
	ModuleNotFound  = errors.New("module not found")
	VersionNotFound = errors.New("version not found")
)
