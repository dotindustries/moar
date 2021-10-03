package storage

import "errors"

var (
	ModuleNotFound = errors.New("module not found")
	FileNotFound   = errors.New("file not found")
)
