package internal

import (
	"fmt"
	"github.com/hashicorp/go-version"
)

type VersionPart int

func ParseVersionPart(s string) (VersionPart, error) {
	switch s {
	case "major":
		return MajorLevel, nil
	case "minor":
		return MinorLevel, nil
	case "patch":
		return PatchLevel, nil
	default:
		return 0, fmt.Errorf("invalid version part: %s", s)
	}
}

const (
	MajorLevel VersionPart = iota + 1
	MinorLevel
	PatchLevel
)

type CurrentVersion struct {
	Major int
	Minor int
	Patch int
}

// NewCurrentVersion creates a new version object
func NewCurrentVersion(v *version.Version) *CurrentVersion {
	return &CurrentVersion{
		Major: v.Segments()[0],
		Minor: v.Segments()[1],
		Patch: v.Segments()[2],
	}
}

// Bump the version
func (v *CurrentVersion) Bump(part VersionPart) {
	switch part {
	case MajorLevel:
		v.Major++
		v.Minor = 0
		v.Patch = 0
	case MinorLevel:
		v.Minor++
		v.Patch = 0
	case PatchLevel:
		v.Patch++
	}
}

// String returns the version as a string
func (v *CurrentVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Version returns the version as a version object
func (v *CurrentVersion) Version() *version.Version {
	return version.Must(version.NewVersion(v.String()))

}
