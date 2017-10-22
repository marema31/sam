package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newProjectEditCommand() *cobra.Command {
	var editService bool

	cmd := &cobra.Command{
		Use:   "edit <project>",
		Short: "Edit a project",
		Long: `Edit the specified project.
Without flags, environement will be edited.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return EditProject(args[0], editService)
		},
	}
	cmd.Flags().BoolVarP(&editService, "services", "s", false, "Edit the services")

	return cmd
}

func EditProject(projectName string, editService bool) error {
	fmt.Println("edit called for project:", projectName, "with service:", editService)

	return nil
}
