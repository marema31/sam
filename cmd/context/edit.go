package context

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newContextEditCommand() *cobra.Command {
	var editService bool

	cmd := &cobra.Command{
		Use:   "edit <project> <context>",
		Short: "Edit a context",
		Long: `Edit the specified context.
Without flags, environement will be edited.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return EditContext(args[0], args[1], editService)
		},
	}
	cmd.Flags().BoolVarP(&editService, "services", "s", false, "Edit the services")

	return cmd
}

func EditContext(projectName string, contextName string, editService bool) error {
	fmt.Println("edit called for context:", contextName, "in project:", projectName, "with service:", editService)

	return nil
}
