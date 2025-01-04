package consumer

import (
	"errors"
	"log"
	"slices"

	"github.com/lj3954/ridgo/pkg/common"
)

type Sku struct {
	Id      string   `json:"Id"`
	Lang    Language `json:"Language"`
	LocLang Language `json:"LocalizedLanguage"`
}

type skus struct {
	Skus []Sku `json:"Skus"`
}

const (
	Release10 = "10"
	Release11 = "11"
)

func ReleaseFromString(s string) (common.Release, error) {
	switch s {
	case Release10:
		return common.Release(s), nil
	case Release11:
		return common.Release(s), nil
	default:
		return "", errors.New("Invalid release")
	}
}

var Releases = []common.Release{Release10, Release11}

type dlOptions struct {
	Options []struct {
		Uri    string      `json:"Uri"`
		DlType common.Arch `json:"DownloadType"`
	} `json:"ProductDownloadOptions"`
}

func SingleSearch(release common.Release, lang Language, arch common.Arch) (*common.UrlData, error) {
	s, err := NewSearch(release, arch)
	if err != nil {
		return nil, err
	}
	skus, err := s.FindSkus()
	if err != nil {
		log.Fatalln(err)
	}
	i := slices.IndexFunc(skus, func(s Sku) bool {
		return s.Lang == lang || s.LocLang == lang
	})
	if i < 0 {
		log.Fatalln("No matching language found")
	}
	sku := skus[i]
	s.SetSku(sku)
	return s.FindUrl()
}
