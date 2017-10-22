package context

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newContextListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls <project>",
		Short: "List all contexts of project",
		Long:  `List all contexts of provided project.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListContext(args[0])
		},
	}

	return cmd
}

func ListContext(projectName string) error {
	fmt.Println("ls called for context for project:", projectName)
	return nil
}
