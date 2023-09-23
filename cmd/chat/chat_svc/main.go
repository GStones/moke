package main

import (
	"github.com/gstones/moke-kit/fxmain"

	"moke/internal/chat/pkg/module"
)

func main() {
	fxmain.Main(
		module.ChatModule,
	)
}
