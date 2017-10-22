package context

import (
	"github.com/spf13/cobra"
)

func NewContextCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "context",
		Short: "Manage contexts",
		Args:  cobra.NoArgs,
		//		RunE: cmd.ShowHelp(),
	}

	cmd.AddCommand(
		newContextCreateCommand(),
		newContextEditCommand(),
		newContextListCommand(),
		newContextRemoveCommand(),
	)
	return cmd
}
