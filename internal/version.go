package internal

import "github.com/Masterminds/semver"

type Version struct {
	Value string `json:"value,omitempty"`
	v     *semver.Version
}

func NewVersion(ver *semver.Version) *Version {
	return &Version{
		Value: ver.String(),
		v:     ver,
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