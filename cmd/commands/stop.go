package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newStopCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop <project>",
		Short: "Stop a project",
		Long: `Stop the service for activated context 
and enabled options of provided project.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Stop(args[0])
		},
	}

	return cmd
}

func Stop(projectName string) error {
	fmt.Println("stop called for project:", projectName)
	return nil
}
