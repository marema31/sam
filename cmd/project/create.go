package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newProjectCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <project>",
		Short: "Create a project",
		Long:  `Create the specified project.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return CreateProject(args[0])
		},
	}

	return cmd
}

func CreateProject(projectName string) error {
	fmt.Println("create called for project:", projectName)
	return nil
}
