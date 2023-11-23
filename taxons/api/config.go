package api

import "github.com/jefersonf/taxons/cfg"

type Config struct {
	listenAddr  string
	appLanguage *cfg.Lang
}

func (c *Config) SetListenAddress(listenAddress string) {
	c.listenAddr = listenAddress
}

func (c *Config) SetLang(langCode cfg.LangCode) {
	c.appLanguage = cfg.GetLangByCode(langCode)
}
