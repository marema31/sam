package context

import (
	"fmt"

	"github.com/marema31/sam/manage"
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

func ListContext(project string) error {
	fmt.Println("ls called for context for project:", project)
	contexts, err := manage.ListContexts(project)
	if err != nil {
		return err
	}
	for _, contexts := range contexts {
		fmt.Println("    ", contexts)
	}
	return nil
}
