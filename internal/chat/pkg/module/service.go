package module

import (
	"go.uber.org/fx"

	"moke/internal/chat/internal/service"
	"moke/internal/chat/pkg/cfx"
)

var ChatModule = fx.Module("chat",
	service.ChatService,
	cfx.ChatSettingsModule,
)
