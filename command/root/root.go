package root

import (
	"fmt"
	"os"

	"github.com/metaplayer-one/metad/command/backup"
	"github.com/metaplayer-one/metad/command/genesis"
	"github.com/metaplayer-one/metad/command/helper"
	"github.com/metaplayer-one/metad/command/ibft"
	"github.com/metaplayer-one/metad/command/license"
	"github.com/metaplayer-one/metad/command/monitor"
	"github.com/metaplayer-one/metad/command/peers"
	"github.com/metaplayer-one/metad/command/secrets"
	"github.com/metaplayer-one/metad/command/server"
	"github.com/metaplayer-one/metad/command/status"
	"github.com/metaplayer-one/metad/command/txpool"
	"github.com/metaplayer-one/metad/command/version"
	"github.com/metaplayer-one/metad/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Metad Chain is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
