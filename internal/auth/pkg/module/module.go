package module

import (
	"go.uber.org/fx"

	fx2 "moke/internal/auth/pkg/mfx"
	public "moke/internal/auth/service"
)

var AuthModule = fx.Module("auth", fx.Options(
	fx2.SettingsModule,
	public.ServiceModule,
))
