package untils

import (
	"github.com/go-ini/ini"
)

func ReadCon(section string,key string) string {
	cfg, err := ini.Load(
		"./app.conf",
	)
	if err != nil {
		return ""
	}
	return cfg.Section(section).Key(key).String()
}
