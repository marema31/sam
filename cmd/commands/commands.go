package commands

import (
	"github.com/marema31/sam/cmd/context"
	"github.com/marema31/sam/cmd/option"
	"github.com/marema31/sam/cmd/project"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		context.NewContextCommand(),
		option.NewOptionCommand(),
		project.NewProjectCommand(),
		newActivateCommand(),
		newLoadCommand(),
		newPsCommand(),
		newStartCommand(),
		newStopCommand(),
		newRestartCommand(),
	)
}
