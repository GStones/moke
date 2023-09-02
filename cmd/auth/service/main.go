package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"moke/internal/auth/pkg/module"
)

func main() {
	fxmain.Main(
		module.AuthModule,
	)
}
