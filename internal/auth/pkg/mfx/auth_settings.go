package mfx

import (
	"github.com/gstones/moke-kit/utility/uconfig"
	"go.uber.org/fx"
)

type AuthSettingParams struct {
	fx.In

	AuthUrl       string `name:"AuthUrl"`
	AuthStoreName string `name:"AuthStoreName"`

	JwtTokenSecret string `name:"JwtTokenSecret"`
}

type AuthSettingsResult struct {
	fx.Out

	AuthStoreName  string `name:"AuthStoreName" envconfig:"AUTH_STORE_NAME" default:"auth"`
	AuthUrl        string `name:"AuthUrl" envconfig:"AUTH_URL" default:"localhost:8081"`
	JwtTokenSecret string `name:"JwtTokenSecret" default:"" envconfig:"JWT_TOKEN_SECRET"`
}

func (g *AuthSettingsResult) LoadFromEnv() (err error) {
	err = uconfig.Load(g)
	return
}

var SettingsModule = fx.Provide(
	func() (out AuthSettingsResult, err error) {
		err = out.LoadFromEnv()
		return
	},
)
