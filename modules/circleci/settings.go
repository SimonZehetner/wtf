package circleci

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	Common *cfg.Common

	apiKey string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.circleci")

	settings := Settings{
		Common: cfg.NewCommonSettingsFromYAML(ymlConfig),
		apiKey: localConfig.UString("apiKey", os.Getenv("WTF_CIRCLE_API_KEY")),
	}

	return &settings
}
