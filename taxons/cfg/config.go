package cfg

type Lang struct{}

type LangCode struct{}

func GetLangByCode(langCode LangCode) *Lang {
	return &Lang{}
}
