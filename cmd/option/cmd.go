package option

import (
	"github.com/spf13/cobra"
)

func NewOptionCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "option",
		Short: "Manage options",
		Args:  cobra.NoArgs,
		//		RunE: cmd.ShowHelp(),
	}

	cmd.AddCommand(
		newOptionCreateCommand(),
		newOptionDisableCommand(),
		newOptionEditCommand(),
		newOptionEnableCommand(),
		newOptionListCommand(),
		newOptionRemoveCommand(),
	)
	return cmd
}
