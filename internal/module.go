package internal

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver"
)

type Module struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	// All available versions in the module
	Versions        []*Version `json:"versions,omitempty"`
	Language        string     `json:"language,omitempty"`
	selectedVersion *Version
}

// Prints the module name and its version selector
//
// If the module has a selected version, it is used instead of the default Latest()
func (m *Module) String() string {
	v := m.selectedVersion.Version()
	if v == nil {
		v = m.Latest()
	}
	if v == nil {
		return m.Name
	}
	return fmt.Sprintf("%s@%s", m.Name, v.String())
}

// HasVersion verifies if this module has a specific version
func (m *Module) HasVersion(version *semver.Version) bool {
	if version == nil {
		return false
	}
	for _, v := range m.Versions {
		if version.Equal(v.Version()) {
			return true
		}
	}
	return false
}

//VersionStrings serializes the versions to string
func (m *Module) VersionStrings() []string {
	l := make([]string, len(m.Versions))
	for i, v := range m.Versions {
		l[i] = v.Version().String()
	}
	return l
}

//SelectVersion selects the highest version number which matches constraint
func (m *Module) SelectVersion(constraint *semver.Constraints) *semver.Version {
	// descending sort of versions
	sort.SliceStable(m.Versions, func(i, j int) bool { return m.Versions[i].Version().GreaterThan(m.Versions[j].Version()) })
	for _, v := range m.Versions {
		if constraint.Check(v.Version()) {
			m.selectedVersion = v
			return v.Version()
		}
	}
	return nil
}

func (m *Module) SelectedVersion() *semver.Version {
	if m.selectedVersion == nil {
		return nil
	}
	return m.selectedVersion.Version()
}

func (m *Module) SetSelectedVersion(version *semver.Version) {
	for _, v := range m.Versions {
		if v.Version().Equal(version) {
			m.selectedVersion = v
		}
	}
}

// Latest returns the latest stable version, excludes pre-releases
func (m *Module) Latest() *semver.Version {
	if len(m.Versions) == 0 {
		return nil
	}
	sort.SliceStable(m.Versions, func(i, j int) bool { return m.Versions[i].Version().GreaterThan(m.Versions[j].Version()) })
	return m.Versions[0].Version()
}

func (m *Module) Init() {
	for _, version := range m.Versions {
		version.Init()
	}
}
