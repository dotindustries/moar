package rpc

import (
	"connectrpc.com/connect"
	"fmt"
	"github.com/dotindustries/moar/internal"
	"github.com/dotindustries/moar/moarpb/v1"
)

func requiredArgumentError(name string) error {
	return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("missing required argument: %s", name))
}

func moduleToDto(module *internal.Module) *moarpb.Module {
	var versions []*moarpb.Version
	for _, v := range module.Versions {
		var files []*moarpb.File
		for _, file := range v.Files {
			files = append(files, &moarpb.File{
				Name:     file.Name,
				MimeType: file.MimeType,
				Data:     file.Data,
			})
		}
		versions = append(versions, &moarpb.Version{
			Name:  v.Value,
			Files: files,
		})
	}
	mod := &moarpb.Module{
		Name:     module.Name,
		Author:   module.Author,
		Language: module.Language,
		Versions: versions,
	}
	return mod
}
