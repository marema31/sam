package project

import (
	"github.com/spf13/cobra"
)

func NewProjectCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
		Args:  cobra.NoArgs,
		//		RunE: cmd.ShowHelp(),
	}

	cmd.AddCommand(
		newProjectCreateCommand(),
		newProjectEditCommand(),
		newProjectListCommand(),
		newProjectRemoveCommand(),
	)
	return cmd
}
