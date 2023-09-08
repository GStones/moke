package module

import (
	"go.uber.org/fx"

	"moke/internal/buddy/internal/app/service/public"
	fx2 "moke/internal/buddy/pkg/bfx"
)

var BuddyModule = fx.Module("buddy",
	public.ServiceModule,
	fx2.BuddySettingsModule,
)
