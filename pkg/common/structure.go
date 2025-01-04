package common

import (
	"errors"
	"strings"
)

type UrlData struct {
	URL string
}

type Arch int

//go:generate stringer -type=Arch
const (
	I686 = Arch(iota)
	X86_64
	Arm64
)

func ArchFromString(s string) (Arch, error) {
	switch strings.ToLower(s) {
	case "i686", "x86", "x32":
		return I686, nil
	case "amd64", "x86_64", "x64":
		return X86_64, nil
	case "arm64", "aarch64":
		return Arm64, nil
	}
	return 0, errors.New("Unexpected architecture")
}

type Release string
