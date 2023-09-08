package main

import (
	"github.com/gstones/moke-kit/fxmain"

	"moke/internal/buddy/pkg/module"
)

func main() {
	fxmain.Main(
		module.BuddyModule,
	)
}
