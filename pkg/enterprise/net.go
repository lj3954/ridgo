package enterprise

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/lj3954/ridgo/pkg/common"
)

func GetUrl(release common.Release, lang Language, arch common.Arch) (*common.UrlData, error) {
	if strings.HasPrefix(string(release), "1") {
		if lang == Russian {
			return nil, errors.New("Invalid language for edition")
		}
	} else {
		if lang == BrazilianPortuguese || lang == EnglishGB || lang == Korean || lang == TraditionalChinese {
			return nil, errors.New("Invalid language for edition")
		}
	}
	url := fmt.Sprintf("https://www.microsoft.com/en-us/evalcenter/download-windows-%s", fixedRelease(release))
	html, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer html.Body.Close()
	body, err := io.ReadAll(html.Body)
	if err != nil {
		return nil, err
	}

	var p string
	if strings.HasSuffix(string(release), "ltsc") {
		p = "LTSC "
	}
	pattern := fmt.Sprintf(`%s%s \(%s\).*?href="([^"]+)"`, p, editionText(arch), lang)
	re := regexp.MustCompile(pattern)
	match := re.FindSubmatch(body)
	if len(match) != 2 {
		return nil, errors.New("Could not find URL")
	}
	r := strings.NewReplacer(`&amp;`, `&`)
	return &common.UrlData{URL: r.Replace(string(match[1]))}, nil
}
