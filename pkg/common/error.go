package common

import (
	"fmt"
	"strings"
)

type InvalidReleaseError struct {
	ReleaseList []Release
	Invalid     Release
}

func (e InvalidReleaseError) Error() string {
	releases := make([]string, len(e.ReleaseList))
	for i, r := range e.ReleaseList {
		releases[i] = string(r)
	}
	return fmt.Sprintf("Invalid release: %s. Valid releases: %s", e.Invalid, strings.Join(releases, ", "))
}

type RejectedRequestError struct{}

func (RejectedRequestError) Error() string {
	return "Microsoft blocked the request based on your IP address"
}

type UnexpArchError struct {
	ArchList []Arch
	Invalid  Arch
}

func (e UnexpArchError) Error() string {
	archs := make([]string, len(e.ArchList))
	for i, a := range e.ArchList {
		archs[i] = a.String()
	}
	return fmt.Sprintf("Invalid architecture: %s. Valid architectures: %s", e.Invalid, strings.Join(archs, ", "))
}
