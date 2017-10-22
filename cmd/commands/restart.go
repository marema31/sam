package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRestartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart <project>",
		Short: "Restart a project",
		Long: `Restart the service for activated context 
and enabled options of provided project.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Restart(args[0])
		},
	}

	return cmd
}

func Restart(projectName string) error {
	fmt.Println("restart called for project:", projectName)
	return nil
}
