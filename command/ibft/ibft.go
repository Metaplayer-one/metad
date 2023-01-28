package ibft

import (
	"github.com/metaplayer-one/metad/command/helper"
	"github.com/metaplayer-one/metad/command/ibft/candidates"
	"github.com/metaplayer-one/metad/command/ibft/propose"
	"github.com/metaplayer-one/metad/command/ibft/quorum"
	"github.com/metaplayer-one/metad/command/ibft/snapshot"
	"github.com/metaplayer-one/metad/command/ibft/status"
	_switch "github.com/metaplayer-one/metad/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
