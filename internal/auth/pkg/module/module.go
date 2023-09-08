package module

import (
	"go.uber.org/fx"

	"moke/internal/auth/pkg/mfx"
	public "moke/internal/auth/service"
)

// AuthModule Auth service module
var AuthModule = fx.Module("auth", fx.Options(
	mfx.SettingsModule,
	public.ServiceModule,
))

// AuthMiddlewareModule Auth service middleware for grpc
// if import this module, every grpc unary/stream will auth by {mfx.AuthCheckModule}
var AuthMiddlewareModule = fx.Module("auth_middleware", fx.Options(
	mfx.SettingsModule,
	mfx.AuthClientModule,
	mfx.AuthCheckModule,
))
