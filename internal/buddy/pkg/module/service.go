package module

import (
	"go.uber.org/fx"

	"moke/internal/buddy/service"

	fx2 "moke/internal/buddy/pkg/bfx"
)

var BuddyModule = fx.Module("buddy",
	service.Module,
	fx2.BuddySettingsModule,
)
