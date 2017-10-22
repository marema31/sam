package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newProjectRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm <project>",
		Short: "Remove a project",
		Long:  `Remove the specified project.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RemoveProject(args[0])
		},
	}

	return cmd
}

func RemoveProject(projectName string) error {
	fmt.Println("rm called for project:", projectName)
	return nil
}
