package client

import (
	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/utility/cshell"
)

func RunAuthCmd(url string) {
	sh := ishell.New()
	cshell.Info(sh, "interactive demo connect to "+url)

	if authCmd, err := NewAuthClient(url); err != nil {
		cshell.Die(sh, err)
		return
	} else {
		sh.AddCmd(authCmd.GetCmd())
		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				err := authCmd.Close()
				if err != nil {
					cshell.Die(c, err)
				}
				cshell.Done(c, "interrupted, press again to exit")
			}
		})
	}

	sh.Run()
}
