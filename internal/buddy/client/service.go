package client

import (
	"github.com/abiosoft/ishell"
	"github.com/gstones/moke-kit/utility/cshell"
)

func RunBuddyCmd(url string) {
	sh := ishell.New()
	cshell.Info(sh, "interactive buddy connect to "+url)

	if buddyCmd, err := NewBuddyClient(url); err != nil {
		cshell.Die(sh, err)
		return
	} else {
		sh.AddCmd(buddyCmd.GetCmd())
		sh.Interrupt(func(c *ishell.Context, count int, input string) {
			if count >= 2 {
				c.Stop()
			}
			if count == 1 {
				err := buddyCmd.Close()
				if err != nil {
					cshell.Die(c, err)
				}
				cshell.Done(c, "interrupted, press again to exit")
			}
		})
	}

	sh.Run()
}
