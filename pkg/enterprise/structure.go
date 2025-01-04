package enterprise

import (
	"errors"
	"strings"

	"github.com/lj3954/ridgo/pkg/common"
)

const (
	ReleaseTenEnterprise    = "10-enterprise"
	ReleaseTenLtsc          = "10-ltsc"
	ReleaseElevenEnterprise = "11-enterprise"
	ReleaseElevenLtsc       = "11-ltsc"
	ReleaseServer2012R2     = "server-2012-r2"
	ReleaseServer2016       = "server-2016"
	ReleaseServer2019       = "server-2019"
	ReleaseServer2022       = "server-2022"
	ReleaseServer2025       = "server-2025"
)

func fixedRelease(release common.Release) string {
	if strings.HasSuffix(string(release), "ltsc") {
		return string(release[:len(release)-4]) + "-enterprise"
	}
	return string(release)
}

func ReleaseFromString(s string) (common.Release, error) {
	switch s {
	case ReleaseTenEnterprise, ReleaseTenLtsc, ReleaseElevenEnterprise, ReleaseElevenLtsc:
		return common.Release(s), nil
	case ReleaseServer2012R2, ReleaseServer2016, ReleaseServer2019, ReleaseServer2022, ReleaseServer2025:
		return common.Release(s), nil
	default:
		return "", errors.New("Invalid release")
	}
}

func editionText(arch common.Arch) string {
	switch arch {
	case common.X86_64:
		return "64-bit"
	case common.I686:
		return "32-bit"
	case common.Arm64:
		return "ARM64-bit"
	default:
		panic("Unreachable")
	}
}
