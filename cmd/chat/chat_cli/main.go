package main

import (
	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"
)

const (
	defaultChat      = "localhost:8081"
	defaultAuth      = "localhost:8081"
	defaultAuthAppId = "test"
	defaultUsername  = "test"
	defaultSecure    = false
)

var options struct {
	chat     string
	auth     string
	appId    string
	username string
	secure   bool
}

func main() {

	rootCmd := &cobra.Command{
		Use:   "chat_cli",
		Short: "Run an interactive chat",
	}

	{
		shell := &cobra.Command{
			Use:   "shell",
			Short: "Run an interactive chat client",
			Run: func(cmd *cobra.Command, args []string) {
				shell()
			},
		}
		rootCmd.AddCommand(shell)
	}

	rootCmd.PersistentFlags().StringVar(
		&options.auth,
		"auth",
		defaultAuth,
		"authentication service (<host>:<port>)",
	)
	rootCmd.PersistentFlags().StringVar(
		&options.chat,
		"chat",
		defaultChat,
		"chat service (<host>:<port>)",
	)
	rootCmd.PersistentFlags().StringVar(
		&options.appId,
		"appId",
		defaultAuthAppId,
		"appId for authentication",
	)
	rootCmd.PersistentFlags().StringVar(
		&options.username,
		"username",
		defaultUsername,
		"username for authentication",
	)
	rootCmd.PersistentFlags().BoolVar(
		&options.secure,
		"secure",
		defaultSecure,
		"if provided, connect securely",
	)

	rootCmd.Execute()
}

func shell() {
	sh := ishell.New()

	sh.Println("Chat Interactive Shell")

	sh.Run()
}
