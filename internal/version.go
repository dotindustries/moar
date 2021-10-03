package internal

import "github.com/Masterminds/semver"

type VersionResources struct {
	ScriptUri string `json:"script_uri,omitempty"`
	StyleUri  string `json:"style_uri,omitempty"`
}

type Version struct {
	Value string `json:"value,omitempty"`
	v     *semver.Version
	Files []File `json:"files,omitempty"`
}

func NewVersion(ver *semver.Version, files []File) *Version {
	return &Version{
		Value: ver.String(),
		v:     ver,
		Files: files,
	}
}

func (v *Version) Init() {
	ver, err := semver.NewVersion(v.Value)
	if err != nil {
		panic(err)
	}
	v.v = ver
}

func (v *Version) Version() *semver.Version {
	return v.v
}
