package consumer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/google/uuid"

	"github.com/lj3954/ridgo/pkg/common"
)

type Search struct {
	arch      common.Arch
	client    *http.Client
	userAgent string
	url       string
	sessionId string
	skuId     string
}

const (
	profile          = "606624d44113"
	ff124ReleaseTime = 1710806400
	fourWeeks        = 2419200
)

func createUserAgent() string {
	// We'll select an expected latest version of Firefox based on its 4 week release cycle
	now := time.Now().Unix()
	ffRelease := 124 + (now-ff124ReleaseTime)/fourWeeks
	return fmt.Sprintf("Mozilla/5.0 (X11, Linux x86_64; rv:%d.0) Gecko/20100101 Firefox/%d.0", ffRelease, ffRelease)
}

func NewSearch(release common.Release, arch common.Arch) (*Search, error) {
	var s *Search
	switch {
	case release == Release10:
		if arch != common.X86_64 && arch != common.I686 {
			return nil, common.UnexpArchError{ArchList: []common.Arch{common.I686, common.X86_64}, Invalid: arch}
		}
		s = &Search{url: "https://microsoft.com/en-us/software-download/windows10ISO"}
	case release == Release11:
		switch arch {
		case common.X86_64:
			s = &Search{url: "https://microsoft.com/en-us/software-download/windows11"}
		case common.Arm64:
			s = &Search{url: "https://microsoft.com/en-us/software-download/windows11ARM64"}
		default:
			return nil, common.UnexpArchError{ArchList: []common.Arch{common.Arm64, common.X86_64}, Invalid: arch}
		}
	default:
		return nil, common.InvalidReleaseError{ReleaseList: Releases, Invalid: release}
	}
	s.arch = arch
	s.client = &http.Client{}
	s.userAgent = createUserAgent()
	s.sessionId = uuid.NewString()
	return s, nil
}

func (s *Search) findProductEditionId() (string, error) {
	req, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", s.userAgent)
	req.Header.Set("Accept", "")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(`option value="(\d+)"`)
	match := r.FindSubmatch(body)
	if len(match) == 0 {
		return "", errors.New("Could not find product edition id")
	}
	return string(match[1]), nil
}

func (s *Search) permitSession() error {
	url := fmt.Sprintf("https://vlscppe.microsoft.com/tags?org_id=y6jn8c31&session_id=%s", s.sessionId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", s.userAgent)
	req.Header.Set("Accept", "")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (s *Search) FindSkus() ([]Sku, error) {
	productEditionId, err := s.findProductEditionId()
	if err != nil {
		return nil, err
	}
	fmt.Println(productEditionId)
	if err := s.permitSession(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://www.microsoft.com/software-download-connector/api/getskuinformationbyproductedition?profile=%s&ProductEditionId=%s&SKU=undefined&friendlyFileName=undefined&Locale=en-US&sessionID=%s", profile, productEditionId, s.sessionId)
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var skus skus
	if err = json.Unmarshal(body, &skus); err != nil {
		return nil, err
	}
	return skus.Skus, nil
}

func (s *Search) SetSkuId(skuid string) {
	s.skuId = skuid
}

func (s *Search) SetSku(sku Sku) {
	s.SetSkuId(sku.Id)
}

func (s *Search) FindUrl() (*common.UrlData, error) {
	if s.skuId == "" {
		return nil, errors.New("You must set the skuid before finding URLs")
	}
	url := fmt.Sprintf("https://www.microsoft.com/software-download-connector/api/GetProductDownloadLinksBySku?profile=%s&productEditionId=undefined&SKU=%s&friendlyFileName=undefined&Locale=en-US&sessionID=%s", profile, s.skuId, s.sessionId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", s.url)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if bytes.Contains(body, []byte("Sentinel marked this request as rejected")) {
		return nil, common.RejectedRequestError{}
	}

	var urls dlOptions
	if err = json.Unmarshal(body, &urls); err != nil {
		return nil, err
	}

	for _, url := range urls.Options {
		if url.DlType == s.arch {
			return &common.UrlData{URL: url.Uri}, nil
		}
	}
	return nil, errors.New("Could not find URL for the specified architecture")
}
