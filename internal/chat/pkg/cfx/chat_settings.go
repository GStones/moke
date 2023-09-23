package cfx

import (
	"github.com/gstones/moke-kit/utility/uconfig"
	"go.uber.org/fx"
)

type ChatSettingParams struct {
	fx.In
	Name         string `name:"Name"`
	ChatInterval int    `name:"ChatInterval"`
}

type ChatSettingResult struct {
	fx.Out
	Name         string `name:"Name" envconfig:"NAME" default:"chat"`
	ChatInterval int    `name:"ChatInterval" envconfig:"WORLD_CHAT_INTERVAL" default:"5"`
}

func (l *ChatSettingResult) LoadFromEnv() (err error) {
	err = uconfig.Load(l)
	return
}

var ChatSettingsModule = fx.Provide(
	func() (out ChatSettingResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
