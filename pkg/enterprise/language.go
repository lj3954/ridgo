package enterprise

import "fmt"

type Language string

const (
	BrazilianPortuguese Language = "pt-BR"
	EnglishGB           Language = "en-GB"
	EnglishUS           Language = "en-US"
	French              Language = "fr-FR"
	German              Language = "de-DE"
	Italian             Language = "it-IT"
	Japanese            Language = "ja-JP"
	Korean              Language = "ko-KR"
	Russian             Language = "ru-RU"
	SimplifiedChinese   Language = "zh-CN"
	Spanish             Language = "es-ES"
	TraditionalChinese  Language = "zh-TW"
)

func LanguageFromString(s string) (Language, error) {
	switch s {
	case "Portuguese (Brazil)":
		return BrazilianPortuguese, nil
	case "English (Great Britain)", "English International":
		return EnglishGB, nil
	case "English (United States)", "English":
		return EnglishUS, nil
	case "French":
		return French, nil
	case "German":
		return German, nil
	case "Italian":
		return Italian, nil
	case "Japanese":
		return Japanese, nil
	case "Korean":
		return Korean, nil
	case "Russian":
		return Russian, nil
	case "Chinese (Simplified)", "Chinese Simplified":
		return SimplifiedChinese, nil
	case "Spanish":
		return Spanish, nil
	case "Chinese (Traditional)", "Chinese Traditional":
		return TraditionalChinese, nil
	default:
		return "", fmt.Errorf("invalid language: %s", s)
	}
}
