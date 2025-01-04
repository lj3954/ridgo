package consumer

import "fmt"

type Language string

const (
	Arabic               Language = "Arabic"
	BrazilianPortuguese  Language = "Brazilian Portuguese"
	Bulgarian            Language = "Bulgarian"
	Croatian             Language = "Croatian"
	Czech                Language = "Czech"
	Danish               Language = "Danish"
	Dutch                Language = "Dutch"
	EnglishInternational Language = "English International"
	EnglishUS            Language = "English (United States)"
	Estonian             Language = "Estonian"
	Finnish              Language = "Finnish"
	French               Language = "French"
	FrenchCanadian       Language = "French Canadian"
	German               Language = "German"
	Greek                Language = "Greek"
	Hebrew               Language = "Hebrew"
	Hungarian            Language = "Hungarian"
	Italian              Language = "Italian"
	Japanese             Language = "Japanese"
	Korean               Language = "Korean"
	Latvian              Language = "Latvian"
	Lithuanian           Language = "Lithuanian"
	MexicanSpanish       Language = "Spanish (Mexico)"
	Norwegian            Language = "Norwegian"
	Polish               Language = "Polish"
	Portuguese           Language = "Portuguese"
	Romanian             Language = "Romanian"
	Russian              Language = "Russian"
	SerbianLatin         Language = "Serbian Latin"
	SimplifiedChinese    Language = "Chinese (Simplified)"
	Slovak               Language = "Slovak"
	Slovenian            Language = "Slovenian"
	Spanish              Language = "Spanish"
	Swedish              Language = "Swedish"
	Thai                 Language = "Thai"
	TraditionalChinese   Language = "Chinese (Traditional)"
	Turkish              Language = "Turkish"
	Ukrainian            Language = "Ukrainian"
)

func LanguageFromString(s string) (Language, error) {
	switch s {
	case "Arabic":
		return Arabic, nil
	case "Brazilian Portuguese":
		return BrazilianPortuguese, nil
	case "Bulgarian":
		return Bulgarian, nil
	case "Croatian":
		return Croatian, nil
	case "Czech":
		return Czech, nil
	case "Danish":
		return Danish, nil
	case "Dutch":
		return Dutch, nil
	case "English International":
		return EnglishInternational, nil
	case "English (United States)", "English":
		return EnglishUS, nil
	case "Estonian":
		return Estonian, nil
	case "Finnish":
		return Finnish, nil
	case "French":
		return French, nil
	case "French Canadian":
		return FrenchCanadian, nil
	case "German":
		return German, nil
	case "Greek":
		return Greek, nil
	case "Hebrew":
		return Hebrew, nil
	case "Hungarian":
		return Hungarian, nil
	case "Italian":
		return Italian, nil
	case "Japanese":
		return Japanese, nil
	case "Korean":
		return Korean, nil
	case "Latvian":
		return Latvian, nil
	case "Lithuanian":
		return Lithuanian, nil
	case "Spanish (Mexico)":
		return MexicanSpanish, nil
	case "Norwegian":
		return Norwegian, nil
	case "Polish":
		return Polish, nil
	case "Portuguese":
		return Portuguese, nil
	case "Romanian":
		return Romanian, nil
	case "Russian":
		return Russian, nil
	case "Serbian Latin":
		return SerbianLatin, nil
	case "Chinese (Simplified)", "Chinese Simplified":
		return SimplifiedChinese, nil
	case "Slovak":
		return Slovak, nil
	case "Slovenian":
		return Slovenian, nil
	case "Spanish":
		return Spanish, nil
	case "Swedish":
		return Swedish, nil
	case "Thai":
		return Thai, nil
	case "Chinese (Traditional)", "Chinese Traditional":
		return TraditionalChinese, nil
	case "Turkish":
		return Turkish, nil
	case "Ukrainian":
		return Ukrainian, nil
	default:
		return "", fmt.Errorf("invalid language: %s", s)
	}
}
