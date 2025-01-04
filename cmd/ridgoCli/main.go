package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lj3954/ridgo/pkg/common"
	"github.com/lj3954/ridgo/pkg/consumer"
	"github.com/lj3954/ridgo/pkg/enterprise"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("Usage: ridgo <release> <language> <arch>")
	}
	arch, err := common.ArchFromString(os.Args[3])
	if err != nil {
		log.Fatalln(err)
	}
	var url *common.UrlData
	if release, err := consumer.ReleaseFromString(os.Args[1]); err == nil {
		url, err = consumerUrl(release, os.Args[2], arch)
		if err != nil {
			log.Fatalln(err)
		}
	} else if release, err := enterprise.ReleaseFromString(os.Args[1]); err == nil {
		url, err = enterpriseUrl(release, os.Args[2], arch)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("Invalid release")
	}
	fmt.Println(url.URL)
}

func consumerUrl(release common.Release, lang string, arch common.Arch) (*common.UrlData, error) {
	l, err := consumer.LanguageFromString(lang)
	if err != nil {
		return nil, err
	}
	return consumer.SingleSearch(release, l, arch)
}

func enterpriseUrl(release common.Release, lang string, arch common.Arch) (*common.UrlData, error) {
	l, err := enterprise.LanguageFromString(lang)
	if err != nil {
		return nil, err
	}
	return enterprise.GetUrl(release, l, arch)
}
