package any

import (
	"errors"

	"github.com/lj3954/ridgo/pkg/common"
	"github.com/lj3954/ridgo/pkg/consumer"
	"github.com/lj3954/ridgo/pkg/enterprise"
)

func Get(pRel, pLang, pArch string) (*common.UrlData, error) {
	arch, err := common.ArchFromString(pArch)
	if err != nil {
		return nil, err
	}
	if release, err := consumer.ReleaseFromString(pRel); err == nil {
		lang, err := consumer.LanguageFromString(pLang)
		if err != nil {
			return nil, err
		}
		return consumer.SingleSearch(release, lang, arch)
	} else if release, err := enterprise.ReleaseFromString(pRel); err == nil {
		lang, err := enterprise.LanguageFromString(pLang)
		if err != nil {
			return nil, err
		}
		return enterprise.GetUrl(release, lang, arch)
	} else {
		return nil, errors.New("Invalid release")
	}
}
