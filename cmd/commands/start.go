package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start <project>",
		Short: "Start a project",
		Long: `Start the service for activated context 
and enabled options of provided project.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Start(args[0])
		},
	}

	return cmd
}

func Start(projectName string) error {
	fmt.Println("start called for project:", projectName)
	return nil
}
